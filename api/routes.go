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
