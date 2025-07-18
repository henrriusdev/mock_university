package handlers

import (
	"net/http"
	"strconv"

	"mocku/backend/common"
	"mocku/backend/utils"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func (h *Handler) SettingsNotesPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		notesNumber, err := strconv.Atoi(r.FormValue("notes"))
		if err != nil {
			h.Logger.Printf("Error parsing notes number: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = h.Repo.UpdateNumberNotes(notesNumber, i, w, r)
		if err != nil {
			return nil
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}

func (h *Handler) SettingsDates(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		startRegistrationSubjects, err := utils.ParseDate(r.FormValue("start_registration_subjects"))
		if err != nil {
			h.Logger.Printf("Error parsing start registration subjects: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		endRegistrationSubjects, err := utils.ParseDate(r.FormValue("end_registration_subjects"))
		if err != nil {
			h.Logger.Printf("Error parsing end registration subjects: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		cycleStart, err := utils.ParseDate(r.FormValue("cycle_start"))
		if err != nil {
			h.Logger.Printf("Error parsing cycle start: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		cycleEnd, err := utils.ParseDate(r.FormValue("cycle_end"))
		if err != nil {
			h.Logger.Printf("Error parsing cycle end: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = h.Repo.UpdateDates(startRegistrationSubjects, endRegistrationSubjects, cycleStart, cycleEnd, i, w, r)
		if err != nil {
			return nil
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}

func (h *Handler) SettingsPayments(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		numberFees, err := strconv.Atoi(r.FormValue("payments"))
		if err != nil {
			h.Logger.Printf("Error parsing number of fees: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = h.Repo.UpdateNumberFees(numberFees, i, w, r)
		if err != nil {
			return nil
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}

func (h *Handler) SettingsNotesPercentage(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		config, err := h.Repo.GetConfiguration(i, w, r)
		if err != nil {
			return nil
		}

		if config.NumberNotes > 0 {
			notes, err := utils.ToPercentage(config.NumberNotes, r)
			if err != nil {
				h.Logger.Printf("Error converting to percentage: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			err = h.Repo.UpdateNotesPercentages(notes, i, w, r)
			if err != nil {
				return nil
			}
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}

func (h *Handler) SettingsPaymentsDates(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		config, err := h.Repo.GetConfiguration(i, w, r)
		if err != nil {
			return nil
		}

		payments, err := utils.ParseFeeDates(config.NumberFees, r)
		if err != nil {
			h.Logger.Printf("Error parsing fee dates: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = h.Repo.UpdateFeeDates(payments, i, w, r)
		if err != nil {
			h.Logger.Printf("Error updating configuration: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		i.Redirect(w, r, "/settings", 302)
		return nil
	}

	return fn
}

func (h *Handler) SettingsCycle(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		currentCycle, err := h.Repo.GetCurrentCycle(i, w, r)
		if err != nil {
			return nil
		}

		newCycle := utils.SplitCycle(currentCycle.Name)

		err = h.Repo.InactivateCycle(i, w, r)
		if err != nil {
			return nil
		}

		currentCycle, err = h.Repo.NewCycle(newCycle, i, w, r)
		if err != nil {
			return nil
		}

		err = h.Repo.NewConfiguration(currentCycle, i, w, r)
		if err != nil {
			return nil
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}
