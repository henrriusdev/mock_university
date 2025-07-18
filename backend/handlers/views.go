package handlers

import (
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
