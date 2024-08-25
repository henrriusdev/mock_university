package handlers

import (
	"fmt"

	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
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
