package api

import (
	"context"
	"fmt"
	"log"
	"mocku/backend/database"
	"mocku/backend/ent"
	"mocku/backend/handlers"
	"mocku/backend/repos"
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
	client := initDatabase()

	repo := repos.NewRepo(client, log.New(os.Stdout, "mocku_repo: ", log.LstdFlags))

	handler := handlers.Handler{
		Repo:   repo,
		Logger: log.New(os.Stdout, "mocku: ", log.LstdFlags),
	}

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

// initDatabase initializes the database connection and creates the schema
func initDatabase() *ent.Client {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DATABASE"))

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := database.InsertDefaultRoles(context.Background(), client); err != nil {
		log.Fatalf("failed inserting default roles: %v", err)
	}

	// Insert default users if they don't exist.
	if err := database.InsertDefaultUsers(context.Background(), client); err != nil {
		log.Fatalf("failed inserting default users: %v", err)
	}

	if err := database.InsertDefaultCycle(context.Background(), client); err != nil {
		log.Fatalf("failed inserting default cycle: %v", err)
	}

	if err := database.InsertDefaultConfig(context.Background(), client); err != nil {
		log.Fatalf("failed inserting default config: %v", err)
	}

	return client
}

func NewServices() {

}

func NewRepos() {

}
