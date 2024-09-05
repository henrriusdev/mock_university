package handlers

import (
	"mocku/backend/common"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/subject"
	"mocku/backend/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func (h *Handler) LoginPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		var formData common.LoginDto

		err := c.Bind(&formData)
		if err != nil {
			h.Logger.Printf("Error binding form data: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}

		if err = c.Validate(formData); err != nil {
			h.Logger.Printf("Error validating form data: %v", err)
			common.HandleBadRequest(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}

		user, err := h.Repo.GetByEmail(formData.Email, i, c.Response().Writer, c.Request())
		if err != nil {
			return nil
		}

		careersArray, err := h.Repo.GetCareers(i, c.Response().Writer, c.Request())
		if err != nil {
			return nil
		}
		if !utils.CheckPassword(user.Password, formData.Password) {
			err = i.Render(c.Response().Writer, c.Request(), "Auth/Login", inertia.Props{
				"careers": careersArray,
				"error":   "Credenciales incorrectas",
			})
			if err != nil {
				h.Logger.Printf("Error rendering login page: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
				return nil
			}

			return nil
		}

		utils.LoginRedirect(user.Edges.Role.ID, http.StatusSeeOther, c.Response().Writer, c.Request(), i)

		return nil
	}

	return fn
}

func (h *Handler) SettingsNotesPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.MethodNotAllowed
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
			return common.MethodNotAllowed
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
			return common.MethodNotAllowed
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
			return common.MethodNotAllowed
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

		notes := make([]float64, config.NumberNotes)
		if config.NumberNotes > 0 {
			notes, err = utils.ToPercentage(config.NumberNotes, r)
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
			return common.MethodNotAllowed
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
			return common.MethodNotAllowed
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

func (h *Handler) StudentPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		var studentRequest common.StudentRequestDto
		if err := c.Bind(&studentRequest); err != nil {
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		handler, err := c.FormFile("avatar")
		if err != nil {
			if err.Error() != "http: no such file" {
				h.Logger.Printf("Error getting file: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			handler = nil
		}

		// Guarda el archivo si se ha subido
		filePath, err := utils.UploadAvatar(studentRequest.Username, handler)
		if err != nil {
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		hashedPassword, err := utils.HashPassword(studentRequest.IdentityCard)
		if err != nil {
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		user, err := h.Repo.CreateUser(studentRequest, hashedPassword, filePath, i, w, r)
		if err != nil {
			return nil
		}
		err = h.Repo.CreateStudent(studentRequest, user, i, w, r)

		http.Redirect(w, r, "/directive/students", http.StatusSeeOther)
		return nil
	}

	return fn
}

func (h *Handler) Career(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.MethodNotAllowed
		}

		var careerRequest common.CareerRequestDto
		if err := c.Bind(&careerRequest); err != nil {
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		if err := h.Repo.CreateCareer(careerRequest, i, w, r); err != nil {
			return nil
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
			return common.MethodNotAllowed
		}

		var professorRequest common.ProfessorRequestDto
		if err := c.Bind(&professorRequest); err != nil {
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}
		handler, err := c.FormFile("avatar")
		if err != nil {
			h.Logger.Printf("Error getting file: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
		}

		// Guarda el archivo si se ha subido
		filePath, err := utils.UploadAvatar(professorRequest.Username, handler)
		if err != nil {
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		hashedPassword, err := utils.HashPassword(professorRequest.IdentityCard)
		if err != nil {
			h.Logger.Printf("Error hashing password: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		user, err := h.Repo.CreateUser(professorRequest, hashedPassword, filePath, i, w, r)
		if err != nil {
			return nil
		}

		err = h.Repo.CreateProfessor(professorRequest, user, i, w, r)
		if err != nil {
			return nil
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
			return common.MethodNotAllowed
		}

		err := r.ParseForm()
		if err != nil {
			h.Logger.Printf("Error parsing form: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
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
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		semesterInt, err := strconv.Atoi(semester)
		if err != nil {
			h.Logger.Printf("Error parsing semester: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		practiceHoursInt, err := strconv.Atoi(practiceHours)
		if err != nil {
			h.Logger.Printf("Error parsing practice hours: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		theoryHoursInt, err := strconv.Atoi(theoryHours)
		if err != nil {
			h.Logger.Printf("Error parsing theory hours: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		labHoursInt, err := strconv.Atoi(labHours)
		if err != nil {
			h.Logger.Printf("Error parsing lab hours: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		totalHoursInt, err := strconv.Atoi(totalHours)
		if err != nil {
			h.Logger.Printf("Error parsing total hours: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		professorIdInt, err := strconv.Atoi(professorId)
		if err != nil {
			h.Logger.Printf("Error parsing professor id: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		classScheduleMap := make(map[string][]string)
		if classSchedule != "" {
			err = utils.Unmarshal(classSchedule, &classScheduleMap)
			if err != nil {
				h.Logger.Printf("Error unmarshaling class schedule: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		}

		careersIdsSlice := strings.Split(careersIds, ",")
		careers, err := utils.StringSliceToIntSlice(careersIdsSlice)
		if err != nil {
			h.Logger.Printf("Error parsing careers ids: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		if id == "" {
			professor, err := h.DB.Professor.Query().
				Where(professor.ID(professorIdInt)).
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying professor: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
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
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		} else {
			subjectId, err := strconv.Atoi(id)
			if err != nil {
				h.Logger.Printf("Error parsing subject id: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			subj, err := h.DB.Subject.Query().
				Where(subject.ID(subjectId)).
				WithProfessor().
				WithCareer().
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying subject: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			professor, err := h.DB.Professor.Query().
				Where(professor.ID(professorIdInt)).
				Only(r.Context())
			if err != nil {
				h.Logger.Printf("Error querying professor: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
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
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}
		}

		i.Redirect(w, r, "/directive/subjects", 302)

		return nil
	}

	return fn
}
