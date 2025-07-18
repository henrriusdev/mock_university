package handlers

import (
	"strconv"
	"strings"

	"mocku/backend/common"

	"mocku/backend/utils"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

func (h *Handler) PaymentsDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		careers, err := h.Repo.GetCareers(i, w, r)
		if err != nil {
			h.Logger.Printf("Error getting careers: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = i.Render(w, r, "Payment/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			h.Logger.Printf("Error rendering payment dash: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}

	return fn
}

func (h *Handler) ControlDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		careers, err := h.Repo.GetCareers(i, w, r)
		if err != nil {
			h.Logger.Printf("Error getting careers: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = i.Render(w, r, "Control/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			h.Logger.Printf("Error rendering control dash: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}

	return fn
}

func (h *Handler) ProfessorDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		careers, err := h.Repo.GetCareers(i, w, r)
		if err != nil {
			h.Logger.Printf("Error getting careers: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		err = i.Render(w, r, "Professor/Dash", inertia.Props{
			"careers": careers,
		})
		if err != nil {
			h.Logger.Printf("Error rendering professor dash: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}

	return fn
}

func (h *Handler) StudentDash(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		user := c.Get("user_id").(float64)
		userID := int(user)

		userName := c.Get("name").(string)

		config, err := h.Repo.GetConfiguration(i, w, r)
		if err != nil {
			return nil
		}

		notes, err := h.Repo.GetStudentNotes(userID, i, w, r)
		if err != nil {
			return nil
		}

		notesDto := make([]common.NoteDto, len(notes))
		for i, note := range notes {
			notesDto[i] = common.NoteDto{
				ID:      note.ID,
				Subject: note.Edges.Subject.Name,
				Notes:   note.Notes,
				Average: utils.Average(note.Notes, config.NotesPercentages),
			}
		}

		err = i.Render(w, r, "Student/Dash", inertia.Props{
			"notes":                     notesDto,
			"notesNumber":               config.NumberNotes,
			"scheduleRegistrationStart": utils.FormatDate(config.StartRegistrationSubjects),
			"scheduleRegistrationEnd":   utils.FormatDate(config.EndRegistrationSubjects),
			"userName":                  userName,
		})
		if err != nil {
			h.Logger.Printf("Error rendering student dash: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}

	return fn
}

func (h *Handler) Settings(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		config, err := h.Repo.GetConfiguration(i, w, r)
		if err != nil {
			return nil
		}

		err = i.Render(w, r, "Directive/Settings/Home", inertia.Props{
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
			h.Logger.Printf("Error rendering settings: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}

	return fn
}

func (h *Handler) Careers(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		careers, err := h.Repo.GetCareersWithLeader(i, w, r)
		if err != nil {
			return nil
		}

		careerDtos := make([]common.CareerDto, len(careers))
		for i, career := range careers {
			if career.Edges.Leader == nil {
				careerDtos[i] = common.CareerDto{
					ID:          career.ID,
					Name:        career.Name,
					Description: career.Description,
					LeaderName:  "",
					LeaderId:    0,
				}
				continue
			}
			careerDtos[i] = common.CareerDto{
				ID:          career.ID,
				Name:        career.Name,
				Description: career.Description,
				LeaderName:  career.Edges.Leader.Edges.User.Name,
				LeaderId:    career.Edges.Leader.Edges.User.ID,
			}
		}

		professors, err := h.Repo.GetProfessors(i, w, r)
		if err != nil {
			return nil
		}

		professorsDto := make([]common.SelectDto, len(professors))
		for i, professor := range professors {
			professorsDto[i] = common.SelectDto{
				ID:   professor.ID,
				Name: professor.Edges.User.Name,
			}
		}

		err = i.Render(w, r, "Directive/Careers", inertia.Props{
			"careers":    careerDtos,
			"professors": professorsDto,
		})
		if err != nil {
			h.Logger.Printf("Error rendering careers: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}
	return fn
}

func (h *Handler) Professors(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		professors, err := h.Repo.GetProfessors(i, w, r)
		if err != nil {
			h.Logger.Printf("Error getting professors: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		professorDtos := make([]common.ProfessorDto, len(professors))
		for i, professor := range professors {
			professorDtos[i] = common.ProfessorDto{
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
			h.Logger.Printf("Error rendering professors: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}
	return fn
}

func (h *Handler) Professor(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		id := r.URL.Query().Get("id")

		var professorDto common.ProfessorDto
		var userDto common.UserDto

		if id != "add" {
			professorId, err := strconv.Atoi(id)
			if err != nil {
				h.Logger.Printf("Error converting id to int: %v", err)
				common.HandleServerErr(i, err).ServeHTTP(w, r)
				return nil
			}

			professor, err := h.Repo.GetProfessorById(professorId, i, w, r)
			if err != nil {
				return nil
			}

			professorDto = common.ProfessorDto{
				ID:           professor.ID,
				IdentityCard: professor.IdentityCard,
				Phone:        professor.Phone,
			}

			userDto = common.UserDto{
				ID:       0,
				Name:     professor.Edges.User.Name,
				Email:    professor.Edges.User.Email,
				Username: professor.Edges.User.Username,
				Avatar:   strings.Replace(professor.Edges.User.Avatar, "./", "/", 1),
				Active:   professor.Edges.User.IsActive,
			}
		}

		// get all professors that don't have a boss, without BossID, use Boss Edge
		professors, err := h.Repo.GetProfessorsWithoutBoss(i, w, r)
		if err != nil {
			return nil
		}

		bossesDto := make([]common.SelectDto, len(professors))
		for i, professor := range professors {
			bossesDto[i] = common.SelectDto{
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
			h.Logger.Printf("Error rendering professor: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}
	return fn
}

func (h *Handler) Subjects(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		subjects, err := h.Repo.GetSubjects(i, w, r)
		if err != nil {
			return nil
		}

		subjectDtos := make([]common.SubjectDto, len(subjects))
		for i, subj := range subjects {
			subjectDtos[i] = common.SubjectDto{
				ID:            subj.ID,
				Name:          subj.Name,
				Description:   subj.Description,
				Code:          subj.Code,
				CreditUnits:   subj.CreditUnits,
				Semester:      subj.Semester,
				PracticeHours: subj.PracticeHours,
				TheoryHours:   subj.TheoryHours,
				LabHours:      subj.LabHours,
				TotalHours:    subj.TotalHours,
				ClassSchedule: subj.ClassSchedule,
				ProfessorId:   subj.Edges.Professor.ID,
				ProfessorName: subj.Edges.Professor.Edges.User.Name,
				Careers:       nil,
			}

			for _, career := range subj.Edges.Career {
				subjectDtos[i].Careers = append(subjectDtos[i].Careers, common.SelectDto{
					ID:   career.ID,
					Name: career.Name,
				})
			}
		}

		err = i.Render(w, r, "Directive/Subjects/Home", inertia.Props{
			"subjects": subjectDtos,
		})
		if err != nil {
			h.Logger.Printf("Error rendering subjects: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}
	return fn
}

func (h *Handler) Subject(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		id := r.URL.Query().Get("id")

		var subjectDto common.SubjectDto

		if id != "add" {
			subjectId, _ := strconv.Atoi(id)

			subj, err := h.Repo.GetSubjectById(subjectId, i, w, r)
			if err != nil {
				return nil
			}

			subjectDto = common.SubjectDto{
				ID:            subj.ID,
				Name:          subj.Name,
				Description:   subj.Description,
				Code:          subj.Code,
				CreditUnits:   subj.CreditUnits,
				Semester:      subj.Semester,
				PracticeHours: subj.PracticeHours,
				TheoryHours:   subj.TheoryHours,
				LabHours:      subj.LabHours,
				TotalHours:    subj.TotalHours,
				ClassSchedule: subj.ClassSchedule,
				ProfessorId:   subj.Edges.Professor.ID,
				ProfessorName: subj.Edges.Professor.Edges.User.Name,
			}

			subjectDto.Careers = common.FillSelectDto(subj.Edges.Career, "ID", "Name")

			subjectDto.Prerequisites = common.FillSelectDto(subj.Edges.Prerequisites, "ID", "Name")
		}

		careers, _ := h.Repo.GetCareers(i, w, r)
		careerDtos := common.FillSelectDto(careers, "ID", "Name")

		professors, _ := h.Repo.GetProfessors(i, w, r)
		professorDtos := common.FillSelectDto(professors, "ID", "Edges.User.Name")

		subjects, _ := h.Repo.GetSubjectsNotID(subjectDto.ID, i, w, r)
		subjectDtos := common.FillSelectDtoSubject(subjects)

		err := i.Render(w, r, "Directive/Subjects/Upsert", inertia.Props{
			"subject":    subjectDto,
			"professors": professorDtos,
			"careers":    careerDtos,
			"subjects":   subjectDtos,
		})
		if err != nil {
			h.Logger.Printf("Error rendering subject: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}

	return fn
}

func (h *Handler) StudentSchedule(i *inertia.Inertia) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		w, r := c.Response().Writer, c.Request()

		subjects, err := h.Repo.GetSubjects(i, w, r)
		if err != nil {
			return nil
		}

		var subjectsDto []common.ScheduleSubjectDto
		for _, subj := range subjects {
			subjectsDto = append(subjectsDto, common.ScheduleSubjectDto{
				ID:            subj.ID,
				Name:          subj.Name,
				Description:   subj.Description,
				Code:          subj.Code,
				Credits:       subj.CreditUnits,
				Semester:      subj.Semester,
				PHours:        subj.PracticeHours,
				THours:        subj.TheoryHours,
				LHours:        subj.LabHours,
				ClassSchedule: subj.ClassSchedule,
			})
		}

		err = i.Render(w, r, "Student/Schedule/Register", inertia.Props{
			"subjects": subjectsDto,
		})
		if err != nil {
			h.Logger.Printf("Error rendering schedules: %v", err)
			common.HandleServerErr(i, err).ServeHTTP(w, r)
			return nil
		}

		return nil
	}
	return fn
}
