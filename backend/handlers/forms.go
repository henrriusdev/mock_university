package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"mocku/backend/ent/careers"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/student"
	"mocku/backend/ent/subject"
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
			h.Logger.Printf("Error binding form data: %v", err)
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}

		if err = c.Validate(formData); err != nil {
			h.Logger.Printf("Error validating form data: %v", err)
			HandleBadRequest(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}

		user, err := h.DB.Users.Query().
			Where(users.EmailEQ(formData.Email)).
			WithRole().
			First(c.Request().Context())
		if err != nil {
			h.Logger.Printf("Error querying user: %v", err)
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}

		careers, err := h.DB.Careers.Query().
			All(c.Request().Context())
		if err != nil {
			h.Logger.Printf("Error querying careers: %v", err)
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}
		if !utils.CheckPassword(user.Password, formData.Password) {
			err = i.Render(c.Response().Writer, c.Request(), "Auth/Login", inertia.Props{
				"careers": careers,
				"error":   "Credenciales incorrectas",
			})
			if err != nil {
				h.Logger.Printf("Error rendering login page: %v", err)
				HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
				return nil
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
			h.Logger.Printf("Error parsing form: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		notesNumber, err := strconv.Atoi(r.FormValue("notes"))
		if err != nil {
			h.Logger.Printf("Error parsing notes number: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.Active(true))).
			SetNumberNotes(notesNumber).
			Save(r.Context())
		if err != nil {
			h.Logger.Printf("Error updating configuration: %v", err)
			HandleServerErr(i, err)
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
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		startRegistrationSubjects, err := utils.ParseDate(r.FormValue("start_registration_subjects"))
		if err != nil {
			h.Logger.Printf("Error parsing start registration subjects: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		endRegistrationSubjects, err := utils.ParseDate(r.FormValue("end_registration_subjects"))
		if err != nil {
			h.Logger.Printf("Error parsing end registration subjects: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		cycleStart, err := utils.ParseDate(r.FormValue("cycle_start"))
		if err != nil {
			h.Logger.Printf("Error parsing cycle start: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		cycleEnd, err := utils.ParseDate(r.FormValue("cycle_end"))
		if err != nil {
			h.Logger.Printf("Error parsing cycle end: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		_, err = h.DB.Cycle.Update().
			Where(cycle.ActiveEQ(true)).
			SetStartDate(cycleStart).
			SetEndDate(cycleEnd).
			Save(r.Context())
		if err != nil {
			h.Logger.Printf("Error updating cycle: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
			SetStartRegistrationSubjects(startRegistrationSubjects).
			SetEndRegistrationSubjects(endRegistrationSubjects).
			Save(r.Context())
		if err != nil {
			h.Logger.Printf("Error updating configuration: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
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
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		numberFees, err := strconv.Atoi(r.FormValue("payments"))
		if err != nil {
			h.Logger.Printf("Error parsing number of fees: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
			SetNumberFees(numberFees).
			Save(r.Context())
		if err != nil {
			h.Logger.Printf("Error updating configuration: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
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
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		config := h.DB.Configuration.Query().
			WithCycle().Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).
			OnlyX(r.Context())

		notes := make([]float64, config.NumberNotes)
		if config.NumberNotes > 0 {
			for j := 0; j < config.NumberNotes; j++ {
				note, err := strconv.Atoi(r.FormValue(fmt.Sprintf("note-%d", j+1)))
				if err != nil {
					h.Logger.Printf("Error parsing note: %v", err)
					HandleServerErr(i, err).ServeHTTP(w, r)
					return nil
				}

				notes[j] = float64(note) / 100
			}

			_, err = h.DB.Configuration.Update().
				Where(configuration.HasCycleWith(cycle.Active(true))).
				SetNotesPercentages(notes).
				Save(r.Context())
			if err != nil {
				h.Logger.Printf("Error updating configuration: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
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
			HandleNotFound(i).ServeHTTP(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
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
				h.Logger.Printf("Error parsing payment date: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			fmt.Println(payment)
			payments[j] = payment
		}

		_, err = h.DB.Configuration.Update().
			Where(configuration.HasCycleWith(cycle.Active(true))).
			SetFeeDates(payments).
			Save(r.Context())
		if err != nil {
			h.Logger.Printf("Error updating configuration: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
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
			h.Logger.Printf("Error updating cycle: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		currentCycle, err = h.DB.Cycle.Create().
			SetStartDate(time.Now()).
			SetEndDate(time.Now()).
			SetActive(true).
			SetName(newCycle).
			Save(r.Context())
		if err != nil {
			h.Logger.Printf("Error creating cycle: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		_, err = h.DB.Configuration.Create().
			SetNumberNotes(0).
			SetNumberFees(0).
			SetStartRegistrationSubjects(time.Now()).
			SetEndRegistrationSubjects(time.Now()).
			SetCycle(currentCycle).Save(r.Context())
		if err != nil {
			h.Logger.Printf("Error creating configuration: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
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
			h.Logger.Printf("Error parsing form: %v", err)
			err = errors.New(err.Error() + " parse")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
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
					h.Logger.Printf("Error creating directory: %v", err)
					err = errors.New(err.Error() + " file 637")
					HandleServerErr(i, err).ServeHTTP(w, r)
					return nil
				}
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				h.Logger.Printf("Error copying file: %v", err)
				err = errors.New(err.Error() + " file 646")
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		}

		hashedPassword, err := utils.HashPassword(identityCard)
		if err != nil {
			h.Logger.Printf("Error hashing password: %v", err)
			err = errors.New(err.Error() + " password")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		postalCodeInt, err := strconv.Atoi(postalCode)
		if err != nil {
			h.Logger.Printf("Error parsing postal code: %v", err)
			err = errors.New(err.Error() + " postalCode")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		birthDateTime, err := utils.ParseDate(birthDate)
		if err != nil {
			h.Logger.Printf("Error parsing birth date: %v", err)
			err = errors.New(err.Error() + " birthDate")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		creditUnitsAccumulated, err := strconv.Atoi(cuAccumulated)
		if err != nil {
			h.Logger.Printf("Error parsing credit units accumulated: %v", err)
			err = errors.New(err.Error() + " creditUnitsAccumulated")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		semesterInt, err := strconv.Atoi(semester)
		if err != nil {
			h.Logger.Printf("Error parsing semester: %v", err)
			err = errors.New(err.Error() + " semester")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		totalAverageFloat, err := strconv.ParseFloat(totalAverage, 64)
		if err != nil {
			h.Logger.Printf("Error parsing total average: %v", err)
			err = errors.New(err.Error() + " totalAverage")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		careerId, err := strconv.Atoi(career)
		if err != nil {
			h.Logger.Printf("Error parsing career: %v", err)
			err = errors.New(err.Error() + " career")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
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
				h.Logger.Printf("Error creating user: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
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
				h.Logger.Printf("Error creating student: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		} else {
			studentId, err := strconv.Atoi(id)
			if err != nil {
				h.Logger.Printf("Error parsing student id: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			student, err := h.DB.Student.Query().
				Where(student.ID(studentId)).
				WithUser().
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying student: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			_, err = h.DB.Users.UpdateOne(student.Edges.User).
				SetEmail(email).
				SetUsername(username).
				SetName(name).
				SetAvatar(filePath).
				Save(r.Context())
			if err != nil {
				h.Logger.Printf("Error updating user: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
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
				h.Logger.Printf("Error updating student: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
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
			h.Logger.Printf("Error parsing form: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		name := r.FormValue("name")
		description := r.FormValue("description")
		leaderId := r.FormValue("leader")
		id := r.FormValue("id")

		if id == "" {
			var leader int
			if leaderId != "" {
				leader, err = strconv.Atoi(leaderId)
				if err != nil {
					h.Logger.Printf("Error parsing leader id: %v", err)
					HandleServerErr(i, err).ServeHTTP(w, r)
					return nil
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
				h.Logger.Printf("Error creating career: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		} else {
			careerId, err := strconv.Atoi(id)
			if err != nil {
				h.Logger.Printf("Error parsing career id: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			oldCareer, err := h.DB.Careers.Query().
				Where(careers.ID(careerId)).
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying career: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			var leader int
			if leaderId != "" {
				leader, err = strconv.Atoi(leaderId)
				if err != nil {
					h.Logger.Printf("Error parsing leader id: %v", err)
					HandleServerErr(i, err).ServeHTTP(w, r)
					return nil
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
				h.Logger.Printf("Error updating career: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
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
			h.Logger.Printf("Error parsing form: %v", err)
			err = errors.New(err.Error() + " parse")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
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
					h.Logger.Printf("Error creating directory: %v", err)
					err = errors.New(err.Error() + " file 637")
					HandleServerErr(i, err).ServeHTTP(w, r)
					return nil
				}
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				h.Logger.Printf("Error copying file: %v", err)
				err = errors.New(err.Error() + " file 646")
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		}

		hashedPassword, err := utils.HashPassword(identityCard)
		if err != nil {
			h.Logger.Printf("Error hashing password: %v", err)
			err = errors.New(err.Error() + " password")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		birthDateTime, err := utils.ParseDate(birthDate)
		if err != nil {
			h.Logger.Printf("Error parsing birth date: %v", err)
			err = errors.New(err.Error() + " birthDate")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
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
				return nil
			}

			_, err = h.DB.Professor.Create().
				SetPhone(phone).
				SetIdentityCard(identityCard).
				SetAddress(address).
				SetBirthDate(birthDateTime).
				SetUser(user).
				Save(r.Context())
			if err != nil {
				h.Logger.Printf("Error creating professor: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		} else {
			professorId, err := strconv.Atoi(id)
			if err != nil {
				h.Logger.Printf("Error parsing professor id: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			professor, err := h.DB.Professor.Query().
				Where(professor.ID(professorId)).
				WithUser().
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying professor: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			_, err = h.DB.Users.UpdateOne(professor.Edges.User).
				SetEmail(email).
				SetUsername(username).
				SetName(name).
				SetAvatar(filePath).
				Save(r.Context())
			if err != nil {
				h.Logger.Printf("Error updating user: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			_, err = h.DB.Professor.UpdateOne(professor).
				SetPhone(phone).
				SetIdentityCard(identityCard).
				SetAddress(address).
				SetBirthDate(birthDateTime).
				Save(r.Context())
			if err != nil {
				h.Logger.Printf("Error updating professor: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		}

		http.Redirect(w, r, "/directive/professors", http.StatusSeeOther)

		return nil
	}

	return fn
}

func (h *Handler) SubjectPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return methodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		name := r.FormValue("name")
		description := r.FormValue("description")
		creditUnits := r.FormValue("creditUnits")
		semester := r.FormValue("semester")
		code := r.FormValue("code")
		practiceHours := r.FormValue("practiceHours")
		theoryHours := r.FormValue("theoryHours")
		labHours := r.FormValue("labHours")
		totalHours := r.FormValue("totalHours")
		classSchedule := r.FormValue("classSchedule")
		professorId := r.FormValue("professor")
		careersIds := r.FormValue("careers")
		id := r.FormValue("id")

		creditUnitsInt, err := strconv.Atoi(creditUnits)
		if err != nil {
			h.Logger.Printf("Error parsing credit units: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		semesterInt, err := strconv.Atoi(semester)
		if err != nil {
			h.Logger.Printf("Error parsing semester: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		practiceHoursInt, err := strconv.Atoi(practiceHours)
		if err != nil {
			h.Logger.Printf("Error parsing practice hours: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		theoryHoursInt, err := strconv.Atoi(theoryHours)
		if err != nil {
			h.Logger.Printf("Error parsing theory hours: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		labHoursInt, err := strconv.Atoi(labHours)
		if err != nil {
			h.Logger.Printf("Error parsing lab hours: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		totalHoursInt, err := strconv.Atoi(totalHours)
		if err != nil {
			h.Logger.Printf("Error parsing total hours: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		professorIdInt, err := strconv.Atoi(professorId)
		if err != nil {
			h.Logger.Printf("Error parsing professor id: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		classScheduleMap := make(map[string][]string)
		if classSchedule != "" {
			err = utils.Unmarshal(classSchedule, &classScheduleMap)
			if err != nil {
				h.Logger.Printf("Error unmarshaling class schedule: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		}

		careersIdsSlice := strings.Split(careersIds, ",")
		careers, err := utils.StringSliceToIntSlice(careersIdsSlice)
		if err != nil {
			h.Logger.Printf("Error parsing careers ids: %v", err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		if id == "" {
			professor, err := h.DB.Professor.Query().
				Where(professor.ID(professorIdInt)).
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying professor: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			_, err = h.DB.Subject.Create().
				SetName(name).
				SetDescription(description).
				SetCreditUnits(creditUnitsInt).
				SetSemester(semesterInt).
				SetCode(code).
				SetPracticeHours(practiceHoursInt).
				SetTheoryHours(theoryHoursInt).
				SetLabHours(labHoursInt).
				SetTotalHours(totalHoursInt).
				SetClassSchedule(classScheduleMap).
				SetProfessor(professor).
				AddCareerIDs(careers...).
				Save(r.Context())
			if err != nil {
				h.Logger.Printf("Error creating subject: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		} else {
			subjectId, err := strconv.Atoi(id)
			if err != nil {
				h.Logger.Printf("Error parsing subject id: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			subj, err := h.DB.Subject.Query().
				Where(subject.ID(subjectId)).
				WithProfessor().
				WithCareer().
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying subject: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			professor, err := h.DB.Professor.Query().
				Where(professor.ID(professorIdInt)).
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying professor: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			_, err = h.DB.Subject.UpdateOne(subj).
				SetName(name).
				SetDescription(description).
				SetCreditUnits(creditUnitsInt).
				SetSemester(semesterInt).
				SetCode(code).
				SetPracticeHours(practiceHoursInt).
				SetTheoryHours(theoryHoursInt).
				SetLabHours(labHoursInt).
				SetTotalHours(totalHoursInt).
				SetClassSchedule(classScheduleMap).
				SetProfessor(professor).
				AddCareerIDs(careers...).
				Save(r.Context())
			if err != nil {
				h.Logger.Printf("Error creating subject: %v", err)
				HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		}

		i.Redirect(w, r, "/directive/subjects", 302)

		return nil
	}

	return fn
}
