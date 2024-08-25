package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"mocku/backend/ent/careers"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/student"
	"mocku/backend/ent/users"
	"mocku/backend/utils"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func (h *Handler) LoginPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		if c.Request().Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(c.Response().Writer, c.Request())
			return methodNotAllowed
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
			return methodNotAllowed
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
			return methodNotAllowed
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

func (h *Handler) SettingsPaymentsPostHandler(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		numberFees, err := strconv.Atoi(r.FormValue("payments"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
			SetNumberFees(numberFees).
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

func (h *Handler) SettingsNotesPercentagePostHandler(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		config := h.DB.Configuration.Query().
			WithCycle().Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
			OnlyX(r.Context())

		notes := make([]float64, config.NumberNotes)
		if config.NumberNotes > 0 {
			for j := 0; j < config.NumberNotes; j++ {
				note, err := strconv.Atoi(r.FormValue(fmt.Sprintf("note-%d", j+1)))
				if err != nil {
					HandleServerErr(i, err).ServeHTTP(w, r)
					return err
				}

				notes[j] = float64(note) / 100
			}

			_, err = h.DB.Configuration.Update().
				Where(configuration.HasCycleWith(cycle.Active(true))).
				SetNotesPercentages(notes).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}

func (h *Handler) SettingsPaymentsDatesPostHandler(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		config := h.DB.Configuration.Query().
			WithCycle().
			Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
			OnlyX(r.Context())

		payments := make([]time.Time, config.NumberFees)
		for j := 0; j < config.NumberFees; j++ {
			date := r.FormValue(fmt.Sprintf("payment-%d", j+1))
			payment, err := utils.ParseDate(date)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			fmt.Println(payment)
			payments[j] = payment
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.Active(true))).
			SetFeeDates(payments).
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

func (h *Handler) SettingsCyclePostHandler(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		currentCycle := h.DB.Cycle.Query().Where(cycle.ActiveEQ(true)).OnlyX(r.Context())

		newCycle := utils.SplitCycle(currentCycle.Name)

		_, err := h.DB.Cycle.Update().
			Where(cycle.ID(currentCycle.ID)).
			SetActive(false).
			Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		currentCycle, err = h.DB.Cycle.Create().
			SetStartDate(time.Now()).
			SetEndDate(time.Now()).
			SetActive(true).
			SetName(newCycle).
			Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		_, err = h.DB.Configuration.Create().
			SetNumberNotes(0).
			SetNumberFees(0).
			SetStartRegistrationSubjects(time.Now()).
			SetEndRegistrationSubjects(time.Now()).
			SetCycle(currentCycle).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		i.Redirect(w, r, "/settings", 302)

		return nil
	}

	return fn
}

func (h *Handler) StudentPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return methodNotAllowed
		}

		err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
		if err != nil {
			err = errors.New(err.Error() + " parse")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		// Accediendo a los valores enviados por el formulario
		id := r.FormValue("id")
		phone := r.FormValue("phone")
		district := r.FormValue("district")
		city := r.FormValue("city")
		postalCode := r.FormValue("postalCode")
		address := r.FormValue("address")
		identityCard := r.FormValue("identityCard")
		birthDate := r.FormValue("birthDate")
		cuAccumulated := r.FormValue("creditUnitsAccumulated")
		semester := r.FormValue("semester")
		totalAverage := r.FormValue("totalAverage")
		name := r.FormValue("name")
		email := r.FormValue("email")
		username := r.FormValue("username")
		career := r.FormValue("career")

		filePath := ""
		file, handler, err := r.FormFile("avatar")
		if err == nil {
			defer file.Close()

			// Guarda el archivo si se ha subido
			filePath = "./uploads/" + username + "_avatar" + filepath.Ext(handler.Filename)
			f, err := os.Create(filePath)
			if err != nil {
				err = os.MkdirAll("./uploads", os.ModePerm)
				if err != nil {
					err = errors.New(err.Error() + " file 637")
					HandleServerErr(i, err).ServeHTTP(w, r)
					return err
				}
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				err = errors.New(err.Error() + " file 646")
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		}

		hashedPassword, err := utils.HashPassword(identityCard)
		if err != nil {
			err = errors.New(err.Error() + " password")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		postalCodeInt, err := strconv.Atoi(postalCode)
		if err != nil {
			err = errors.New(err.Error() + " postalCode")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		birthDateTime, err := utils.ParseDate(birthDate)
		if err != nil {
			err = errors.New(err.Error() + " birthDate")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		creditUnitsAccumulated, err := strconv.Atoi(cuAccumulated)
		if err != nil {
			err = errors.New(err.Error() + " creditUnitsAccumulated")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		semesterInt, err := strconv.Atoi(semester)
		if err != nil {
			err = errors.New(err.Error() + " semester")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		totalAverageFloat, err := strconv.ParseFloat(totalAverage, 64)
		if err != nil {
			err = errors.New(err.Error() + " totalAverage")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		careerId, err := strconv.Atoi(career)
		if err != nil {
			err = errors.New(err.Error() + " career")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		if id == "" {
			user, err := h.DB.Users.Create().
				SetEmail(email).
				SetUsername(username).
				SetPassword(hashedPassword).
				SetName(name).
				SetAvatar(filePath).
				SetIsActive(true).
				SetRoleID(6).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			_, err = h.DB.Student.Create().
				SetPhone(phone).
				SetDistrict(district).
				SetCity(city).
				SetPostalCode(postalCodeInt).
				SetAddress(address).
				SetIdentityCard(identityCard).
				SetBirthDate(birthDateTime).
				SetCreditUnitsAccumulated(creditUnitsAccumulated).
				SetSemester(semesterInt).
				SetTotalAverage(totalAverageFloat).
				SetUser(user).
				SetCareerID(careerId).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		} else {
			studentId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			student, err := h.DB.Student.Query().
				Where(student.ID(studentId)).
				WithUser().
				Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			_, err = h.DB.Users.UpdateOne(student.Edges.User).
				SetEmail(email).
				SetUsername(username).
				SetName(name).
				SetAvatar(filePath).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			_, err = h.DB.Student.UpdateOne(student).
				SetPhone(phone).
				SetDistrict(district).
				SetCity(city).
				SetPostalCode(postalCodeInt).
				SetAddress(address).
				SetIdentityCard(identityCard).
				SetBirthDate(birthDateTime).
				SetCreditUnitsAccumulated(creditUnitsAccumulated).
				SetSemester(semesterInt).
				SetTotalAverage(totalAverageFloat).
				SetCareerID(careerId).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		}

		http.Redirect(w, r, "/directive/students", http.StatusSeeOther)

		return nil
	}

	return fn
}

func (h *Handler) Career(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		name := r.FormValue("name")
		description := r.FormValue("description")
		leaderId := r.FormValue("leader")
		id := r.FormValue("id")
		fmt.Println(leaderId)

		if id == "" {
			var leader int
			if leaderId != "" {
				leader, err = strconv.Atoi(leaderId)
				if err != nil {
					HandleServerErr(i, err).ServeHTTP(w, r)
					return err
				}

				_, err = h.DB.Careers.Create().
					SetName(name).
					SetDescription(description).
					SetLeaderID(leader).
					Save(r.Context())
			} else {
				_, err = h.DB.Careers.Create().
					SetName(name).
					SetDescription(description).
					Save(r.Context())
			}

			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		} else {
			careerId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			oldCareer, err := h.DB.Careers.Query().
				Where(careers.ID(careerId)).
				Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			var leader int
			if leaderId != "" {
				leader, err = strconv.Atoi(leaderId)
				if err != nil {
					HandleServerErr(i, err).ServeHTTP(w, r)
					return err
				}

				_, err = h.DB.Careers.UpdateOne(oldCareer).
					SetName(name).
					SetDescription(description).
					SetLeaderID(leader).
					Save(r.Context())
			} else {
				_, err = h.DB.Careers.UpdateOne(oldCareer).
					SetName(name).
					SetDescription(description).
					Save(r.Context())
			}

			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		}

		i.Redirect(w, r, "/directive/careers", 302)

		return nil
	}
	return fn
}

func (h *Handler) ProfessorPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return methodNotAllowed
		}

		err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
		if err != nil {
			err = errors.New(err.Error() + " parse")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		// Accediendo a los valores enviados por el formulario
		id := r.FormValue("id")
		phone := r.FormValue("phone")
		identityCard := r.FormValue("identityCard")
		name := r.FormValue("name")
		email := r.FormValue("email")
		username := r.FormValue("username")
		birthDate := r.FormValue("birthDate")
		address := r.FormValue("address")

		filePath := ""
		file, handler, err := r.FormFile("avatar")
		if err == nil {
			defer file.Close()

			// Guarda el archivo si se ha subido
			filePath = "./uploads/" + username + "_avatar" + filepath.Ext(handler.Filename)
			f, err := os.Create(filePath)
			if err != nil {
				err = os.MkdirAll("./uploads", os.ModePerm)
				if err != nil {
					err = errors.New(err.Error() + " file 637")
					HandleServerErr(i, err).ServeHTTP(w, r)
					return err
				}
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				err = errors.New(err.Error() + " file 646")
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		}

		hashedPassword, err := utils.HashPassword(identityCard)
		if err != nil {
			err = errors.New(err.Error() + " password")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		birthDateTime, err := utils.ParseDate(birthDate)
		if err != nil {
			err = errors.New(err.Error() + " birthDate")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		if id == "" {
			user, err := h.DB.Users.Create().
				SetEmail(email).
				SetUsername(username).
				SetPassword(hashedPassword).
				SetName(name).
				SetAvatar(filePath).
				SetIsActive(true).
				SetRoleID(4).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			_, err = h.DB.Professor.Create().
				SetPhone(phone).
				SetIdentityCard(identityCard).
				SetAddress(address).
				SetBirthDate(birthDateTime).
				SetUser(user).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		} else {
			professorId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			professor, err := h.DB.Professor.Query().
				Where(professor.ID(professorId)).
				WithUser().
				Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			_, err = h.DB.Users.UpdateOne(professor.Edges.User).
				SetEmail(email).
				SetUsername(username).
				SetName(name).
				SetAvatar(filePath).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			_, err = h.DB.Professor.UpdateOne(professor).
				SetPhone(phone).
				SetIdentityCard(identityCard).
				SetAddress(address).
				SetBirthDate(birthDateTime).
				Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}
		}

		http.Redirect(w, r, "/directive/professors", http.StatusSeeOther)

		return nil
	}

	return fn
}
