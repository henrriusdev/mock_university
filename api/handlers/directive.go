package handlers

import (
	"mocku/pkg/common"
	"mocku/pkg/model"
	"mocku/pkg/service"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/romsar/gonertia"
)

type Directive struct {
	i        *gonertia.Inertia
	services *service.Services
}

func NewDirective(i *gonertia.Inertia, services *service.Services) *Directive {
	return &Directive{
		i:        i,
		services: services,
	}
}

func (d *Directive) RegisterRoutes(e *echo.Echo) {
	e.GET("/", d.DirectiveDash)
	e.GET("/students", d.Students)
	e.GET("/students/view", d.Student)
	e.POST("/students/view/submit", d.StudentPost)
	e.GET("/careers", d.Careers)
	e.POST("/careers/submit", d.Career)
	e.GET("/professors", d.Professors)
	e.GET("/professors/view", d.Professor)
	e.POST("/professors/view/submit", d.ProfessorPost)
	e.GET("/subjects", d.Subjects)
	e.GET("/subjects/view", d.Subject)
	e.POST("/subjects/view/submit", d.SubjectPost)
}

func (d *Directive) DirectiveDash(c echo.Context) error {
	careers, err := d.services.Careers.GetAll(c.Request().Context())
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Dash", gonertia.Props{
		"careers": careers,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Students(c echo.Context) error {
	students, err := d.services.Student.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	studentDtos := make([]model.StudentsTableResponse, len(students))
	for i, student := range students {
		studentDtos[i] = model.StudentsTableResponse{
			ID:           int(student.ID),
			Name:         student.User.Name,
			Avatar:       student.User.Avatar,
			Email:        student.User.Email,
			Phone:        student.Phone,
			Career:       student.Career.Name,
			TotalAverage: student.TotalAverage,
		}
	}

	var dtos []interface{}
	for _, studentDto := range studentDtos {
		dtos = append(dtos, studentDto)
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Students/Home", gonertia.Props{
		"students": dtos,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}

func (d *Directive) Student(c echo.Context) error {
	id := c.QueryParam("id")

	var studentDto model.StudentResponse
	var userDto model.UserResponse

	if id != "add" {
		studentId, _ := strconv.Atoi(id)

		student, err := d.services.Student.GetByID(c.Request().Context(), studentId)
		if err != nil {
			return nil
		}

		studentDto = model.StudentResponse{
			ID:                     studentId,
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

		userDto = model.UserResponse{
			ID:       int(student.User.ID),
			Name:     student.User.Name,
			Email:    student.User.Email,
			Username: student.User.Username,
			Avatar:   strings.Replace(student.User.Avatar, "./", "/", 1),
			Active:   student.User.IsActive,
		}
	}

	careers, err := d.services.Careers.GetAll(c.Request().Context())
	if err != nil {
		return nil
	}

	var careerDtos []model.SelectResponse
	for _, career := range careers {
		careerDtos = append(careerDtos, model.SelectResponse{
			ID:   int(career.ID),
			Name: career.Name,
		})
	}

	err = d.i.Render(c.Response().Writer, c.Request(), "Directive/Students/Upsert", gonertia.Props{
		"student": studentDto,
		"user":    userDto,
		"careers": careerDtos,
	})
	if err != nil {
		common.HandleServerErr(d.i, err).ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}

	return nil
}
