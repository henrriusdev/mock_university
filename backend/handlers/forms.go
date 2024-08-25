package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/users"
	"mocku/backend/utils"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func (h *Handler) LoginPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		if c.Request().Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(c.Response().Writer, c.Request())
			return errors.New("method not allowed")
		}

		var formData LoginDto

		err := c.Bind(&formData)
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return err
		}

		if err = c.Validate(formData); err != nil {
			HandleBadRequest(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return err
		}

		user, err := h.DB.Users.Query().
			Where(users.EmailEQ(formData.Email)).
			WithRole().
			First(c.Request().Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return err
		}

		careers, err := h.DB.Careers.Query().
			All(c.Request().Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return err
		}
		if !utils.CheckPassword(user.Password, formData.Password) {
			err = i.Render(c.Response().Writer, c.Request(), "Auth/Login", inertia.Props{
				"careers": careers,
				"error":   "Credenciales incorrectas",
			})
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
				return err
			}

			return nil
		}

		var view string
		role := user.Edges.Role.ID
		switch role {
		case 1:
			view = "/directive"
		case 2:
			view = "/payments"
		case 3:
			view = "/control"
		case 4, 5:
			view = "/professor"
		case 6:
			view = "/student"
		}

		i.Redirect(c.Response().Writer, c.Request(), view, 302)

		return nil
	}

	return fn
}

func (h *Handler) SettingsNotesPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return errors.New("method not allowed")
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		notesNumber, err := strconv.Atoi(r.FormValue("notes"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.Active(true))).
			SetNumberNotes(notesNumber).
			Save(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return err
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}

func (h *Handler) SettingsDatesPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return errors.New("method not allowed")
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		startRegistrationSubjects, err := utils.ParseDate(r.FormValue("start_registration_subjects"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		endRegistrationSubjects, err := utils.ParseDate(r.FormValue("end_registration_subjects"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		cycleStart, err := utils.ParseDate(r.FormValue("cycle_start"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		cycleEnd, err := utils.ParseDate(r.FormValue("cycle_end"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		_, err = h.DB.Cycle.Update().
			Where(cycle.ActiveEQ(true)).
			SetStartDate(cycleStart).
			SetEndDate(cycleEnd).
			Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
			SetStartRegistrationSubjects(startRegistrationSubjects).
			SetEndRegistrationSubjects(endRegistrationSubjects).
			Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}
