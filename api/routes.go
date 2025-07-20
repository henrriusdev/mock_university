package api

import (
	"mocku/api/handlers"
	"mocku/pkg/service"

	"github.com/labstack/echo/v4"
	"github.com/romsar/gonertia"
)

func authRoutes(e *echo.Echo, i *gonertia.Inertia, services *service.Services) {
	auth := handlers.NewAuth(services.Users, services.Careers, i)
	auth.RegisterRoutes(e.Group(""))
}

func directiveRoutes(e *echo.Echo, i *gonertia.Inertia, services *service.Services) {
	directive := handlers.NewDirective(i, services)
	directive.RegisterRoutes(e.Group("/directive"))
}

func settingRoutes(e *echo.Echo, i *gonertia.Inertia, services *service.Services) {
	setting := handlers.NewSetting(i, services)
	setting.RegisterRoutes(e.Group("/settings"))
}

func studentRoutes(e *echo.Echo, i *gonertia.Inertia, services *service.Services) {
	student := handlers.NewStudent(services.Student, services.Configuration, services.Subject, i)
	student.RegisterRoutes(e.Group("/student"))
}

// func (h *Handler) PaymentsDash(i *inertia.Inertia) echo.HandlerFunc {
// 	fn := func(c echo.Context) error {
// 		w, r := c.Response().Writer, c.Request()

// 		careers, err := h.Repo.GetCareers(i, w, r)
// 		if err != nil {
// 			h.Logger.Printf("Error getting careers: %v", err)
// 			common.HandleServerErr(i, err).ServeHTTP(w, r)
// 			return nil
// 		}

// 		err = i.Render(w, r, "Payment/Dash", inertia.Props{
// 			"careers": careers,
// 		})
// 		if err != nil {
// 			h.Logger.Printf("Error rendering payment dash: %v", err)
// 			common.HandleServerErr(i, err).ServeHTTP(w, r)
// 			return nil
// 		}

// 		return nil
// 	}

// 	return fn
// }

// func (h *Handler) ControlDash(i *inertia.Inertia) echo.HandlerFunc {
// 	fn := func(c echo.Context) error {
// 		w, r := c.Response().Writer, c.Request()

// 		careers, err := h.Repo.GetCareers(i, w, r)
// 		if err != nil {
// 			h.Logger.Printf("Error getting careers: %v", err)
// 			common.HandleServerErr(i, err).ServeHTTP(w, r)
// 			return nil
// 		}

// 		err = i.Render(w, r, "Control/Dash", inertia.Props{
// 			"careers": careers,
// 		})
// 		if err != nil {
// 			h.Logger.Printf("Error rendering control dash: %v", err)
// 			common.HandleServerErr(i, err).ServeHTTP(w, r)
// 			return nil
// 		}

// 		return nil
// 	}

// 	return fn
// }

// func (h *Handler) ProfessorDash(i *inertia.Inertia) echo.HandlerFunc {
// 	fn := func(c echo.Context) error {
// 		w, r := c.Response().Writer, c.Request()

// 		careers, err := h.Repo.GetCareers(i, w, r)
// 		if err != nil {
// 			h.Logger.Printf("Error getting careers: %v", err)
// 			common.HandleServerErr(i, err).ServeHTTP(w, r)
// 			return nil
// 		}

// 		err = i.Render(w, r, "Professor/Dash", inertia.Props{
// 			"careers": careers,
// 		})
// 		if err != nil {
// 			h.Logger.Printf("Error rendering professor dash: %v", err)
// 			common.HandleServerErr(i, err).ServeHTTP(w, r)
// 			return nil
// 		}

// 		return nil
// 	}

// 	return fn
// }
