package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"mocku/backend/ent"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/professor"
	"mocku/backend/ent/student"
	"mocku/backend/utils"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func (h *Handler) Home(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		careers, err := h.DB.Careers.Query().All(c.Request().Context())
		if err != nil {
			HandleServerErr(i, err)
			return err
		}

		err = i.Render(c.Response().Writer, c.Request(), "Home/Index", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) Login(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		careers, err := h.DB.Careers.Query().All(c.Request().Context())
		if err != nil {
			HandleServerErr(i, err)
			return err
		}

		err = i.Render(c.Response().Writer, c.Request(), "Auth/Login", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err)
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) DirectiveDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		careers, err := h.DB.Careers.Query().All(c.Request().Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return err
		}

		err = i.Render(c.Response().Writer, c.Request(), "Directive/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(c.Response().Writer, c.Request())
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) PaymentsDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		err = i.Render(w, r, "Payment/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) ControlDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		err = i.Render(w, r, "Control/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) ProfessorDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		err = i.Render(w, r, "Professor/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) StudentDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		err = i.Render(w, r, "Student/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) Settings(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		config := h.DB.Configuration.Query().
			WithCycle().
			Where(configuration.HasCycleWith(cycle.Active(true))).
			OnlyX(r.Context())

		if config == nil {
			HandleServerErr(i, fmt.Errorf("config not found")).ServeHTTP(w, r)
			return nil
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
			return err
		}

		return nil
	}

	return fn
}

func (h *Handler) Students(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		students, err := h.DB.Student.Query().WithUser().WithCareer().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
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
			return err
		}

		fmt.Println(studentsPayload)

		err = i.Render(w, r, "Directive/Students/Home", inertia.Props{
			"students": studentsPayload,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}
	return fn
}

func (h *Handler) Student(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		// the id is a string, if it is "add" it means that the user wants to add a new student
		id := r.URL.Query().Get("id")
		var studentDto StudentDto
		var userDto UserDto
		if id != "add" {
			studentId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			student, err := h.DB.Student.Query().
				Where(student.ID(studentId)).
				WithUser().WithCareer().
				Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
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
			return err
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
			return err
		}

		return nil
	}
	return fn
}

func (h *Handler) Careers(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		careers, err := h.DB.Careers.Query().
			WithLeader(func(query *ent.ProfessorQuery) { query.WithUser() }).
			All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
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

		professors := h.DB.Professor.Query().WithUser().AllX(r.Context())
		professorsDto := make([]SelectDto, len(professors))
		for i, professor := range professors {
			professorsDto[i] = SelectDto{
				ID:   professor.ID,
				Name: professor.Edges.User.Name,
			}
		}

		err = i.Render(w, r, "Directive/Careers", inertia.Props{
			"careers":    careerDtos,
			"professors": professorsDto,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}
	return fn
}

func (h *Handler) Professors(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		professors, err := h.DB.Professor.Query().WithUser().All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		professorDtos := make([]ProfessorDto, len(professors))
		for i, professor := range professors {
			professorDtos[i] = ProfessorDto{
				ID:           professor.ID,
				Name:         professor.Edges.User.Name,
				Email:        professor.Edges.User.Email,
				Avatar:       strings.Replace(professor.Edges.User.Avatar, "./", "/", 1),
				IdentityCard: professor.IdentityCard,
				Phone:        professor.Phone,
			}
		}

		err = i.Render(w, r, "Directive/Professor/Home", inertia.Props{
			"professors": professorDtos,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}
	return fn
}

func (h *Handler) Professor(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()
		id := r.URL.Query().Get("id")
		var professorDto ProfessorDto
		var userDto UserDto
		if id != "add" {
			professorId, err := strconv.Atoi(id)
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			professor, err := h.DB.Professor.Query().Where(professor.ID(professorId)).WithUser().Only(r.Context())
			if err != nil {
				HandleServerErr(i, err).ServeHTTP(w, r)
				return err
			}

			professorDto = ProfessorDto{
				ID:           professor.ID,
				IdentityCard: professor.IdentityCard,
				Phone:        professor.Phone,
			}

			userDto = UserDto{
				ID:       0,
				Name:     professor.Edges.User.Name,
				Email:    professor.Edges.User.Email,
				Username: professor.Edges.User.Username,
				Avatar:   strings.Replace(professor.Edges.User.Avatar, "./", "/", 1),
				Active:   professor.Edges.User.IsActive,
			}
		}

		// get all professors that don't have a boss, without BossID, use Boss Edge
		professors, err := h.DB.Professor.Query().
			WithUser().
			Where(professor.Not(professor.HasBoss())).
			All(r.Context())
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		bossesDto := make([]SelectDto, len(professors))
		for i, professor := range professors {
			bossesDto[i] = SelectDto{
				ID:   professor.ID,
				Name: professor.Edges.User.Name,
			}
		}

		err = i.Render(w, r, "Directive/Professor/Upsert", inertia.Props{
			"professor": professorDto,
			"user":      userDto,
			"bosses":    bossesDto,
		})
		if err != nil {
			HandleServerErr(i, err).ServeHTTP(w, r)
			return err
		}

		return nil
	}
	return fn
}
