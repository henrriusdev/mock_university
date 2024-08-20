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

	"mocku/backend/ent"
	"mocku/backend/ent/careers"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/student"
	"mocku/backend/ent/users"
	"mocku/backend/utils"

	inertia "github.com/romsar/gonertia"
)

func (h *Handler) Home(i *inertia.Inertia) http.Handler {
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

func (h *Handler) Login(i *inertia.Inertia) http.Handler {
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

func (h *Handler) LoginPost(i *inertia.Inertia) http.Handler {
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

func (h *Handler) DirectiveDash(i *inertia.Inertia) http.Handler {
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

func (h *Handler) PaymentsDash(i *inertia.Inertia) http.Handler {
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

func (h *Handler) ControlDash(i *inertia.Inertia) http.Handler {
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

func (h *Handler) ProfessorDash(i *inertia.Inertia) http.Handler {
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

func (h *Handler) StudentDash(i *inertia.Inertia) http.Handler {
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

func (h *Handler) Settings(i *inertia.Inertia) http.Handler {
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

func (h *Handler) SettingsNotesPost(i *inertia.Inertia) http.Handler {
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

func (h *Handler) SettingsDatesPost(i *inertia.Inertia) http.Handler {
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
				Avatar:   strings.Replace(student.Edges.User.Avatar, "./", "/", 1),
				Active:   student.Edges.User.IsActive,
			}
		}

		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		var careerDtos []SelectDto
		for _, career := range careers {
			careerDtos = append(careerDtos, SelectDto{
				ID:   career.ID,
				Name: career.Name,
			})
		}

		err = i.Render(w, r, "Directive/Students/Upsert", inertia.Props{
			"student": studentDto,
			"user":    userDto,
			"careers": careerDtos,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func (h *Handler) StudentPost(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
		if err != nil {
			err = errors.New(err.Error() + " parse")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
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
					return
				}
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				err = errors.New(err.Error() + " file 646")
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}
		}

		hashedPassword, err := utils.HashPassword(identityCard)
		if err != nil {
			err = errors.New(err.Error() + " password")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		postalCodeInt, err := strconv.Atoi(postalCode)
		if err != nil {
			err = errors.New(err.Error() + " postalCode")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		birthDateTime, err := utils.ParseDate(birthDate)
		if err != nil {
			err = errors.New(err.Error() + " birthDate")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		creditUnitsAccumulated, err := strconv.Atoi(cuAccumulated)
		if err != nil {
			err = errors.New(err.Error() + " creditUnitsAccumulated")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		semesterInt, err := strconv.Atoi(semester)
		if err != nil {
			err = errors.New(err.Error() + " semester")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		totalAverageFloat, err := strconv.ParseFloat(totalAverage, 64)
		if err != nil {
			err = errors.New(err.Error() + " totalAverage")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		careerId, err := strconv.Atoi(career)
		if err != nil {
			err = errors.New(err.Error() + " career")
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		if id == "" {
			user, err := h.DB.Users.Create().SetEmail(email).SetUsername(username).SetPassword(hashedPassword).SetName(name).SetAvatar(filePath).SetIsActive(true).SetRoleID(6).Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			_, err = h.DB.Student.Create().SetPhone(phone).SetDistrict(district).SetCity(city).SetPostalCode(postalCodeInt).SetAddress(address).SetIdentityCard(identityCard).SetBirthDate(birthDateTime).SetCreditUnitsAccumulated(creditUnitsAccumulated).SetSemester(semesterInt).SetTotalAverage(totalAverageFloat).SetUser(user).SetCareerID(careerId).Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}
		} else {
			studentId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			student, err := h.DB.Student.Query().Where(student.ID(studentId)).WithUser().Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			_, err = h.DB.Users.UpdateOne(student.Edges.User).SetEmail(email).SetUsername(username).SetName(name).SetAvatar(filePath).Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			_, err = h.DB.Student.UpdateOne(student).SetPhone(phone).SetDistrict(district).SetCity(city).SetPostalCode(postalCodeInt).SetAddress(address).SetIdentityCard(identityCard).SetBirthDate(birthDateTime).SetCreditUnitsAccumulated(creditUnitsAccumulated).SetSemester(semesterInt).SetTotalAverage(totalAverageFloat).SetCareerID(careerId).Save(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}
		}

		http.Redirect(w, r, "/directive/students", http.StatusSeeOther)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) Careers(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().WithLeader(func(query *ent.ProfessorQuery) { query.WithUser() }).All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		careerDtos := make([]CareerDto, len(careers))
		for i, career := range careers {
			if career.Edges.Leader == nil {
				careerDtos[i] = CareerDto{
					ID:          career.ID,
					Name:        career.Name,
					Description: career.Description,
					LeaderName:  "",
					LeaderId:    0,
				}
				continue
			}
			careerDtos[i] = CareerDto{
				ID:          career.ID,
				Name:        career.Name,
				Description: career.Description,
				LeaderName:  career.Edges.Leader.Edges.User.Name,
				LeaderId:    career.Edges.Leader.Edges.User.ID,
			}
		}

		err = i.Render(w, r, "Directive/Careers", inertia.Props{
			"careers": careerDtos,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func (h *Handler) Career(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}

		err := r.ParseForm()
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return
		}

		name := r.FormValue("name")
		description := r.FormValue("description")
		leaderId := r.FormValue("leader")
		id := r.FormValue("id")
		fmt.Println(id)

		if id == "" {
			var leader int
			if leaderId != "" {
				leader, err = strconv.Atoi(leaderId)
				if err != nil {
					HandleServerErr(i, err).ServeHTTP(w, r)
					return
				}

				_, err = h.DB.Careers.Create().SetName(name).SetDescription(description).SetLeaderID(leader).Save(r.Context())
			} else {
				_, err = h.DB.Careers.Create().SetName(name).SetDescription(description).Save(r.Context())
			}

			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}
		} else {
			careerId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			oldCareer, err := h.DB.Careers.Query().Where(careers.ID(careerId)).Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}

			var leader int
			if leaderId != "" {
				leader, err = strconv.Atoi(leaderId)
				if err != nil {
					HandleServerErr(i, err).ServeHTTP(w, r)
					return
				}

				_, err = h.DB.Careers.UpdateOne(oldCareer).SetName(name).SetDescription(description).SetLeaderID(leader).Save(r.Context())
			} else {
				_, err = h.DB.Careers.UpdateOne(oldCareer).SetName(name).SetDescription(description).Save(r.Context())
			}

			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return
			}
		}

		i.Redirect(w, r, "/directive/careers", 302)
	}
	return http.HandlerFunc(fn)
}
