package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/student"
	"mocku/backend/ent/users"
	"mocku/backend/utils"

	inertia "github.com/romsar/gonertia"
)

func (h *Handler) HomeHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Home/Index", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Auth/Login", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		var formData LoginDto
		// get formData from request instead of json
		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err)
			return
		}

		formData.Email = r.FormValue("email")
		formData.Password = r.FormValue("password")
		user, err := h.DB.Users.Query().Where(users.EmailEQ(formData.Email)).WithRole().First(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}
		if !utils.CheckPassword(user.Password, formData.Password) {
			err = i.Render(w, r, "Auth/Login", inertia.Props{
				"careers": careers,
				"error":   "Credenciales incorrectas",
			})
			if err != nil {
				HandleServerErr(i, err)
				return
			}

			return
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

		println(view)

		i.Redirect(w, r, view, 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) DirectiveDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		fmt.Println(err)
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Directive/Dash", inertia.Props{
			"careers": careers,
		})
		fmt.Println(err)
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) PaymentsDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Payment/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) ControlDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Control/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) ProfessorDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Professor/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) StudentDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Student/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Get configuration, if there is no configuration with the active cycle, return empty
		config := h.DB.Configuration.Query().WithCycle().Where(configuration.HasCycleWith(cycle.Active(true))).OnlyX(r.Context())
		if config == nil {
			HandleServerErr(i, fmt.Errorf("config not found")).ServeHTTP(w, r)
			return
		}

		err := i.Render(w, r, "Directive/Settings/Home", inertia.Props{
			"notesNumber":   config.NumberNotes,
			"paymentNumber": config.NumberFees,
			"startRegSubj":  utils.FormatDate(config.StartRegistrationSubjects),
			"endRegSubj":    utils.FormatDate(config.EndRegistrationSubjects),
			"cycleStart":    utils.FormatDate(config.Edges.Cycle.StartDate),
			"cycleEnd":      utils.FormatDate(config.Edges.Cycle.EndDate),
			"percentages":   config.NotesPercentages,
			"paymentDates":  config.FeeDates,
		})
		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsNotesPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		notesNumber, err := strconv.Atoi(r.FormValue("notes"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Configuration.Update().Where(configuration.HasCycleWith(cycle.Active(true))).SetNumberNotes(notesNumber).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsDatesPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		startRegistrationSubjects, err := utils.ParseDate(r.FormValue("start_registration_subjects"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		endRegistrationSubjects, err := utils.ParseDate(r.FormValue("end_registration_subjects"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		cycleStart, err := utils.ParseDate(r.FormValue("cycle_start"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		cycleEnd, err := utils.ParseDate(r.FormValue("cycle_end"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Cycle.Update().Where(cycle.ActiveEQ(true)).SetStartDate(cycleStart).SetEndDate(cycleEnd).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Configuration.Update().Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).SetStartRegistrationSubjects(startRegistrationSubjects).SetEndRegistrationSubjects(endRegistrationSubjects).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsPaymentsPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		numberFees, err := strconv.Atoi(r.FormValue("payments"))
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Configuration.Update().Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).SetNumberFees(numberFees).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsNotesPercentagePostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		config := h.DB.Configuration.Query().WithCycle().Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).OnlyX(r.Context())

		notes := make([]float64, config.NumberNotes)
		if config.NumberNotes > 0 {
			for j := 0; j < config.NumberNotes; j++ {
				note, err := strconv.Atoi(r.FormValue(fmt.Sprintf("note-%d", j+1)))
				if err != nil {
					HandleServerErr(i, err).ServeHTTP(w, r)
					return
				}

				notes[j] = float64(note) / 100
			}

			_, err = h.DB.Configuration.Update().Where(configuration.HasCycleWith(cycle.Active(true))).SetNotesPercentages(notes).Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsPaymentsDatesPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			println(err)
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		config := h.DB.Configuration.Query().WithCycle().Where(configuration.HasCycleWith(cycle.ActiveEQ(true))).OnlyX(r.Context())

		payments := make([]time.Time, config.NumberFees)
		for j := 0; j < config.NumberFees; j++ {
			date := r.FormValue(fmt.Sprintf("payment-%d", j+1))
			payment, err := utils.ParseDate(date)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			fmt.Println(payment)
			payments[j] = payment
		}

		_, err = h.DB.Configuration.Update().Where(configuration.HasCycleWith(cycle.Active(true))).SetFeeDates(payments).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) SettingsCyclePostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		currentCycle := h.DB.Cycle.Query().Where(cycle.ActiveEQ(true)).OnlyX(r.Context())

		newCycle := utils.SplitCycle(currentCycle.Name)

		_, err := h.DB.Cycle.Update().Where(cycle.ID(currentCycle.ID)).SetActive(false).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		currentCycle, err = h.DB.Cycle.Create().SetStartDate(time.Now()).SetEndDate(time.Now()).SetActive(true).SetName(newCycle).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		_, err = h.DB.Configuration.Create().SetNumberNotes(0).SetNumberFees(0).SetStartRegistrationSubjects(time.Now()).SetEndRegistrationSubjects(time.Now()).SetCycle(currentCycle).Save(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		i.Redirect(w, r, "/settings", 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) Students(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		students, err := h.DB.Student.Query().WithUser().WithCareer().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		studentDtos := make([]StudentsTableDto, len(students))
		for i, student := range students {
			studentDtos[i] = StudentsTableDto{
				ID:           student.ID,
				Name:         student.Edges.User.Name,
				Avatar:       student.Edges.User.Avatar,
				Email:        student.Edges.User.Email,
				Phone:        student.Phone,
				Career:       student.Edges.Career.Name,
				TotalAverage: student.TotalAverage,
			}
		}

		var dtos []interface{}
		for _, studentDto := range studentDtos {
			dtos = append(dtos, studentDto)
		}

		studentsPayload, err := utils.StructArrayToJson(dtos)
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		fmt.Println(studentsPayload)

		err = i.Render(w, r, "Directive/Students/Home", inertia.Props{
			"students": studentsPayload,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func (h *Handler) Student(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// the id is a string, if it is "add" it means that the user wants to add a new student
		id := r.URL.Query().Get("id")
		var studentDto StudentDto
		var userDto UserDto
		if id != "add" {
			studentId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			student, err := h.DB.Student.Query().Where(student.ID(studentId)).WithUser().WithCareer().Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			studentDto = StudentDto{
				ID:                     student.ID,
				Phone:                  student.Phone,
				Address:                student.Address,
				District:               student.District,
				City:                   student.City,
				PostalCode:             student.PostalCode,
				IdentityCard:           student.IdentityCard,
				BirthDate:              student.BirthDate.Format("2006-01-02"),
				CreditUnitsAccumulated: student.CreditUnitsAccumulated,
				TotalAverage:           student.TotalAverage,
			}

			userDto = UserDto{
				ID:       0,
				Name:     student.Edges.User.Name,
				Email:    student.Edges.User.Email,
				Username: student.Edges.User.Username,
				Avatar:   student.Edges.User.Avatar,
				Active:   student.Edges.User.IsActive,
			}
		}

		err := i.Render(w, r, "Directive/Students/Upsert", inertia.Props{
			"student": studentDto,
			"user":    userDto,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}
	}
	return http.HandlerFunc(fn)
}
