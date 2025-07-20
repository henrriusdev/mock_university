package handlers

import (
	"mocku/pkg/common"
	"mocku/pkg/service"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/romsar/gonertia"
)

type Setting struct {
	i        *gonertia.Inertia
	services *service.Services
}

func NewSetting(i *gonertia.Inertia, services *service.Services) *Setting {
	return &Setting{
		i:        i,
		services: services,
	}
}

func (s *Setting) RegisterRoutes(e *echo.Group) {
	e.GET("/", s.Settings)
	e.POST("/notes", s.SettingsNotesPost)
	e.POST("/notes/percentages", s.SettingsNotesPercentage)
	e.POST("/payment", s.SettingsPayments)
	e.POST("/payment/dates", s.SettingsPaymentsDates)
	e.POST("/cycle", s.SettingsCycle)
	e.POST("/dates", s.SettingsDates)
}

func (s *Setting) Settings(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	config, err := s.services.Configuration.GetByID(c.Request().Context(), uint(1))
	if err != nil {
		return nil
	}

	err = s.i.Render(w, r, "Directive/Settings/Home", gonertia.Props{
		"notesNumber":   config.NumberNotes,
		"paymentNumber": config.NumberFees,
		"startRegSubj":  common.FormatDate(config.StartRegistrationSubjects),
		"endRegSubj":    common.FormatDate(config.EndRegistrationSubjects),
		"cycleStart":    common.FormatDate(config.Cycle.StartDate),
		"cycleEnd":      common.FormatDate(config.Cycle.EndDate),
		"percentages":   config.NotesPercentages,
		"paymentDates":  config.FeeDates,
	})
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	return nil
}

func (s *Setting) SettingsNotesPost(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	err := r.ParseForm()
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	notesNumber, err := strconv.Atoi(r.FormValue("notes"))
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	err = s.services.Configuration.UpdateNumberNotes(c.Request().Context(), notesNumber)
	if err != nil {
		return nil
	}

	s.i.Redirect(w, r, "/settings", 302)

	return nil
}

func (s *Setting) SettingsDates(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	err := r.ParseForm()
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	startRegistrationSubjects, err := common.ParseDate(r.FormValue("start_registration_subjects"))
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	endRegistrationSubjects, err := common.ParseDate(r.FormValue("end_registration_subjects"))
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	cycleStart, err := common.ParseDate(r.FormValue("cycle_start"))
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	cycleEnd, err := common.ParseDate(r.FormValue("cycle_end"))
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	err = s.services.Configuration.UpdateDates(c.Request().Context(), startRegistrationSubjects, endRegistrationSubjects, cycleStart, cycleEnd)
	if err != nil {
		return nil
	}

	s.i.Redirect(w, r, "/settings", 302)

	return nil
}

func (s *Setting) SettingsPayments(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	err := r.ParseForm()
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	numberFees, err := strconv.Atoi(r.FormValue("payments"))
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	err = s.services.Configuration.UpdateNumberFees(c.Request().Context(), numberFees)
	if err != nil {
		return nil
	}

	s.i.Redirect(w, r, "/settings", 302)

	return nil
}

func (s *Setting) SettingsNotesPercentage(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	err := r.ParseForm()
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	config, err := s.services.Configuration.GetByID(c.Request().Context(), uint(1))
	if err != nil {
		return nil
	}

	if config.NumberNotes > 0 {
		notes, err := common.ToPercentage(config.NumberNotes, r)
		if err != nil {
			common.HandleServerErr(s.i, err).ServeHTTP(w, r)
			return nil
		}

		err = s.services.Configuration.UpdateNotesPercentages(c.Request().Context(), notes)
		if err != nil {
			return nil
		}
	}

	s.i.Redirect(w, r, "/settings", 302)

	return nil
}

func (s *Setting) SettingsPaymentsDates(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	err := r.ParseForm()
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	config, err := s.services.Configuration.GetByID(c.Request().Context(), uint(1))
	if err != nil {
		return nil
	}

	payments, err := common.ParseFeeDates(config.NumberFees, r)
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	err = s.services.Configuration.UpdateFeeDates(c.Request().Context(), payments)
	if err != nil {
		common.HandleServerErr(s.i, err).ServeHTTP(w, r)
		return nil
	}

	s.i.Redirect(w, r, "/settings", 302)
	return nil
}

func (s *Setting) SettingsCycle(c echo.Context) error {
	w, r := c.Response().Writer, c.Request()

	currentCycle, err := s.services.Configuration.GetCurrentCycle(c.Request().Context())
	if err != nil {
		return nil
	}

	newCycle := common.SplitCycle(currentCycle.Name)

	err = s.services.Configuration.InactivateCycle(c.Request().Context())
	if err != nil {
		return nil
	}

	currentCycle, err = s.services.Configuration.NewCycle(c.Request().Context(), newCycle)
	if err != nil {
		return nil
	}

	err = s.services.Configuration.NewConfiguration(c.Request().Context(), currentCycle)
	if err != nil {
		return nil
	}

	s.i.Redirect(w, r, "/settings", 302)

	return nil
}
