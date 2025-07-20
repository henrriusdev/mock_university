package handlers

import (
	"mocku/pkg/common"
	"mocku/pkg/model"
	"mocku/pkg/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"
)

type Auth struct {
	service service.Users
	careers service.Careers
	i       *inertia.Inertia
}

func NewAuth(service service.Users, careers service.Careers, i *inertia.Inertia) *Auth {
	return &Auth{service: service, careers: careers, i: i}
}

func (a *Auth) RegisterRoutes(echo *echo.Group) {
	echo.GET("/login", a.Login)
	echo.POST("/login_post", a.LoginPost)
	echo.GET("", a.Home)
}

func (a *Auth) Home(c echo.Context) error {
	careers, err := a.careers.GetAll(c.Request().Context())
	if err != nil {
		common.HandleServerErr(a.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	err = a.i.Render(c.Response().Writer, c.Request(), "Home/Index", inertia.Props{
		"careers": careers,
	})
	if err != nil {
		return nil
	}

	return nil
}

func (a *Auth) Login(c echo.Context) error {
	careers, err := a.careers.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	errorQuery := c.QueryParam("error")

	err = a.i.Render(c.Response().Writer, c.Request(), "Auth/Login", inertia.Props{
		"careers": careers,
		"error":   errorQuery,
	})
	if err != nil {
		return nil
	}

	return nil
}

func (a *Auth) LoginPost(c echo.Context) error {
	var formData model.LoginRequest
	if c.Request().Method != http.MethodPost {
		common.HandleNotFound(a.i).ServeHTTP(c.Response().Writer, c.Request())
		return common.ErrMethodNotAllowed
	}

	if err := c.Bind(&formData); err != nil {
		common.HandleServerErr(a.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	if err := c.Validate(formData); err != nil {
		common.HandleBadRequest(a.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	user, err := a.service.GetByEmail(c.Request().Context(), formData.Email)
	if err != nil {
		common.HandleServerErr(a.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	if !common.CheckPassword(user.Password, formData.Password) {
		common.HandleUnauthorized(a.i).ServeHTTP(c.Response().Writer, c.Request())
		return common.ErrMethodNotAllowed
	}

	tokenString, err := common.GenerateJWT(int(user.ID), user.Name, user.Email, user.Role.Name)
	if err != nil {
		common.HandleServerErr(a.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	// Almacenar el JWT en una cookie segura
	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	// Almacenar el JWT en una cookie segura
	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(5 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	common.LoginRedirect(int(user.Role.ID), c.Response().Writer, c.Request(), a.i)

	return nil
}
