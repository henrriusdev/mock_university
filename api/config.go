package api

import (
	"context"
	"fmt"
	"log"
	"mocku/pkg/repository"
	"mocku/pkg/service"
	"mocku/pkg/store"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if cv.validator == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Validator is not initialized")
	}
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func MountApp() {
	i := initInertia()
	db := initDatabase()

	repos := NewRepositories(db)
	services := NewServices(repos)

	app := echo.New()

	inertiaMiddl := echo.WrapMiddleware(i.Middleware)
	app.Use(inertiaMiddl)
	app.Use(middleware.Recover())

	app.Validator = &CustomValidator{validator: validator.New()}

	// Routes
	app.GET("/", handler.Home(i))
	app.GET("/login", handler.Login(i))
	app.Any("/login_post", handler.LoginPost(i))

	// Directives routes
	directive := app.Group("/directive", inertiaMiddl, mocku.JWTMiddleware(), mocku.RoleMiddleware("directive"))
	directive.GET("", handler.DirectiveDash(i))
	directive.GET("/students", handler.Students(i))
	directive.GET("/students/view", handler.Student(i))
	directive.Any("/students/view/submit", handler.StudentPost(i))
	directive.GET("/careers", handler.Careers(i))
	directive.Any("/careers/submit", handler.Career(i))
	directive.GET("/professors", handler.Professors(i))
	directive.GET("/professors/view", handler.Professor(i))
	directive.Any("/professors/view/submit", handler.ProfessorPost(i))
	directive.GET("/subjects", handler.Subjects(i))
	directive.GET("/subjects/view", handler.Subject(i))
	directive.Any("/subjects/view/submit", handler.SubjectPost(i))

	// Settings routes
	settings := app.Group("/settings", inertiaMiddl, mocku.JWTMiddleware(), mocku.RoleMiddleware("directive"))
	settings.GET("", handler.Settings(i))
	settings.Any("/notes", handler.SettingsNotesPost(i))
	settings.Any("/notes/percentages", handler.SettingsNotesPercentage(i))
	settings.Any("/payment", handler.SettingsPayments(i))
	settings.Any("/payment/dates", handler.SettingsPaymentsDates(i))
	settings.Any("/cycle", handler.SettingsCycle(i))
	settings.Any("/dates", handler.SettingsDates(i))

	// Students routes
	student := app.Group("/student", inertiaMiddl, mocku.JWTMiddleware(), mocku.RoleMiddleware("student"))
	student.GET("", handler.StudentDash(i))
	student.GET("/schedule", handler.StudentSchedule(i))

	// Dashboard routes
	// mux.Handle("/payment", i.Middleware(handler.PaymentsDash(i)))
	// mux.Handle("/professor", i.Middleware(handler.ProfessorDash(i)))
	// mux.Handle("/control", i.Middleware(handler.ControlDash(i)))

	// // API routes
	app.Static("/build", "./public/build")
	app.Static("/uploads", "./uploads")

	// Start server
	app.Start(":3000")
}

// initDatabase initializes the database connection
func initDatabase() store.Queryable {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DATABASE"))

	db, err := store.NewConnection(dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// TODO: Replace with migrations or schema creation using sqlx
	// For now, we'll assume the schema already exists

	// TODO: Replace with direct inserts using repositories
	// Insert default data if needed
	// insertDefaultData(context.Background(), db)

	return db
}

// insertDefaultData inserts default data into the database
func insertDefaultData(ctx context.Context, db store.Queryable, repos *repository.Repositories, services *service.Services) {
	// TODO: Implement default data insertion using repositories
	// Example:
	// Insert default roles
	// Insert default users
	// Insert default cycle
	// Insert default configurations
}

func NewRepositories(db store.Queryable) *repository.Repositories {
	return &repository.Repositories{
		Users:         repository.NewUsers(db),
		Role:          repository.NewRole(db),
		Cycle:         repository.NewCycle(db),
		Configuration: repository.NewConfiguration(db),
		Student:       repository.NewStudent(db),
		Professor:     repository.NewProfessor(db),
		Subject:       repository.NewSubject(db),
		Module:        repository.NewModule(db),
		Blog:          repository.NewBlog(db),
		Payment:       repository.NewPayment(db),
		PaymentMethod: repository.NewPaymentMethod(db),
		Careers:       repository.NewCareers(db),
		Note:          repository.NewNote(db),
	}
}
