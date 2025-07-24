package handlers

import (
	"mocku/pkg/common"
	"mocku/pkg/model"
	"mocku/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/romsar/gonertia"
)

type Student struct {
	service service.Student
	config  service.Configuration
	subject service.Subject
	i       *gonertia.Inertia
}

func NewStudent(service service.Student, config service.Configuration, subject service.Subject, i *gonertia.Inertia) *Student {
	return &Student{service, config, subject, i}
}

func (s *Student) RegisterRoutes(e *echo.Group) {
	e.GET("", s.Dashboard)
	e.GET("/schedule", s.Schedule)
}

func (s *Student) Dashboard(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	user := c.Get("user_id").(string)
	userName := c.Get("name").(string)

	config, err := s.config.GetByID(c.Request().Context(), "1")
	if err != nil {
		return nil
	}

	notes, err := s.service.GetStudentNotes(c.Request().Context(), user)
	if err != nil {
		return nil
	}

	notesDto := make([]model.NoteResponse, len(notes))
	for i, note := range notes {
		notesDto[i] = model.NoteResponse{
			ID:      note.ID,
			Subject: note.Subject.Name,
			Notes:   note.Notes,
			Average: common.Average(note.Notes, config.NotesPercentages),
		}
	}

	err = s.i.Render(w, r, "Student/Dash", gonertia.Props{
		"notes":                     notesDto,
		"notesNumber":               config.NumberNotes,
		"scheduleRegistrationStart": common.FormatDate(config.StartRegistrationSubjects),
		"scheduleRegistrationEnd":   common.FormatDate(config.EndRegistrationSubjects),
		"userName":                  userName,
	})
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	return nil
}

func (s *Student) Schedule(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	subjects, err := s.subject.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	var subjectsDto []model.ScheduleSubjectResponse
	for _, subj := range subjects {
		subjectsDto = append(subjectsDto, model.ScheduleSubjectResponse{
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

	err = s.i.Render(w, r, "Student/Schedule/Register", gonertia.Props{
		"subjects": subjectsDto,
	})
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	return nil
}
