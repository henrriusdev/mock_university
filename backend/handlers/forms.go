package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"mocku/backend/common"
	"mocku/backend/utils"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func (h *Handler) LoginPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		var formData common.LoginDto
		if c.Request().Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(c.Response().Writer, c.Request())
			return common.ErrMethodNotAllowed
		}

		if err := c.Bind(&formData); err != nil {
			h.Logger.Printf("Error binding form data: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}

		if err := c.Validate(formData); err != nil {
			h.Logger.Printf("Error validating form data: %v", err)
			common.HandleBadRequest(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return nil
		}

		user, err := h.Repo.GetByEmail(formData.Email, i, c.Response().Writer, c.Request())
		if err != nil {
			h.Logger.Printf("Error getting user by email: %v", err)
			return nil
		}

		if !utils.CheckPassword(user.Password, formData.Password) {
			i.Redirect(c.Response().Writer, c.Request(), "/login?error=invalid credentials", 302)
		}

		tokenString, err := utils.GenerateJWT(user.ID, user.Name, user.Email, user.Edges.Role.Name)
		if err != nil {
			i.Redirect(c.Response().Writer, c.Request(), "/login?error=error generating token", 302)
			return nil
		}

		// Almacenar el JWT en una cookie segura
		c.SetCookie(&http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})

		// Almacenar el JWT en una cookie segura
		c.SetCookie(&http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Expires:  time.Now().Add(5 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})

		utils.LoginRedirect(user.Edges.Role.ID, c.Response().Writer, c.Request(), i)

		return nil
	}

	return fn
}

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

func (h *Handler) StudentPost(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		if r.Method != http.MethodPost {
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		var studentRequest common.StudentRequestDto
		if err := c.Bind(&studentRequest); err != nil {
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		handler, err := c.FormFile("avatar")
		if err != nil {
			h.Logger.Printf("Error getting file: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
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
		if err != nil {
			return nil
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
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
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
			return common.ErrMethodNotAllowed
		}

		var professorRequest common.ProfessorRequestDto
		if err := c.Bind(&professorRequest); err != nil {
			h.Logger.Printf("Error binding form data: %v", err)
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
			common.HandleNotFound(i).ServeHTTP(w, r)
			return common.ErrMethodNotAllowed
		}

		var subjectRequest common.SubjectRequestDto
		if err := c.Bind(&subjectRequest); err != nil {
			h.Logger.Printf("Error binding form data: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		careersIdsSlice := strings.Split(subjectRequest.CareerIds, ",")
		careers, err := utils.StringSliceToIntSlice(careersIdsSlice)
		if err != nil {
			h.Logger.Printf("Error parsing careers ids: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		var classSchedule map[string][]string
		if err := utils.Unmarshal(subjectRequest.ClassSchedule, &classSchedule); err != nil {
			h.Logger.Printf("Error unmarshalling class schedule: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		subjectsIdsSlice := strings.Split(subjectRequest.PreqIds, ",")
		subjects, err := utils.StringSliceToIntSlice(subjectsIdsSlice)
		if err != nil {
			h.Logger.Printf("Error parsing prerequisites ids: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = h.Repo.CreateSubject(subjectRequest, careers, subjects, classSchedule, i, w, r)
		if err != nil {
			return nil
		}

		i.Redirect(w, r, "/directive/subjects", 302)

		return nil
	}

	return fn
}
