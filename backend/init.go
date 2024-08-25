package backend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"mocku/backend/database"
	"mocku/backend/ent"
	"mocku/backend/handlers"
	"mocku/backend/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	inertia "github.com/romsar/gonertia"

	_ "github.com/lib/pq"
)

type CustomValidator struct {
	validator *validator.Validate
}

// MountApp mounts the application, initialize inertia.js and database,  and starts the server
func MountApp() {
	i := initInertia()
	client := initDatabase()

	handler := handlers.Handler{
		DB:     client,
		Logger: log.New(os.Stdout, "mocku: ", log.LstdFlags),
	}

	app := echo.New()
	middleware := echo.WrapMiddleware(i.Middleware)
	app.Use(middleware)

	// Routes
	app.GET("/", handler.Home(i))
	app.GET("/login", handler.Login(i))
	app.POST("/login_post", handler.LoginPost(i))

	// Directives routes
	directive := app.Group("/directive", nil)
	directive.GET("/students", handler.Students(i))
	directive.GET("/students/view", handler.Student(i))
	directive.POST("/students/view/submit", handler.StudentPost(i))
	directive.GET("/careers", handler.Careers(i))
	directive.POST("/careers/submit", handler.Career(i))
	directive.GET("/professors", handler.Professors(i))
	directive.GET("/professors/view", handler.Professor(i))
	directive.POST("/professors/view/submit", handler.ProfessorPost(i))

	// Settings routes
	settings := app.Group("/settings", nil)
	settings.GET("", handler.Settings(i))
	settings.POST("/notes", handler.SettingsNotesPost(i))
	settings.POST("/notes/percentages", handler.SettingsNotesPercentage(i))
	settings.POST("/payment", handler.SettingsPayments(i))
	settings.POST("/payment/dates", handler.SettingsPaymentsDates(i))
	settings.POST("/cycle", handler.SettingsCycle(i))
	settings.POST("/dates", handler.SettingsDates(i))

	// Dashboard routes
	// mux.Handle("/payment", i.Middleware(handler.PaymentsDash(i)))
	// mux.Handle("/student", i.Middleware(handler.StudentDash(i)))
	// mux.Handle("/professor", i.Middleware(handler.ProfessorDash(i)))
	// mux.Handle("/control", i.Middleware(handler.ControlDash(i)))

	// // API routes
	app.Static("/build", "./public/build")
	app.Static("/uploads", "./uploads")

	// Start server
	app.Start(":3000")
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
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

// initInertia initializes the inertia.js instance
func initInertia() *inertia.Inertia {
	viteHotFile := "./public/hot"
	rootViewFile := "resources/views/root.html"

	// check if laravel-vite-plugin is running in dev mode (it puts a "hot" file in the public folder)
	_, err := os.Stat(viteHotFile)
	if err == nil {
		i, err := inertia.NewFromFile(
			rootViewFile,
			inertia.WithSSR(),
		)
		if err != nil {
			log.Fatal(err)
		}
		i.ShareTemplateFunc("vite", func(entry string) (string, error) {
			content, err := os.ReadFile(viteHotFile)
			if err != nil {
				return "", err
			}
			url := strings.TrimSpace(string(content))
			if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
				url = url[strings.Index(url, ":")+1:]
			} else {
				url = "//localhost:8080"
			}

			if entry != "" && !strings.HasPrefix(entry, "/") {
				entry = "/" + entry
			}
			return url + entry, nil
		})
		return i
	}

	// laravel-vite-plugin not running in dev mode, use build manifest file
	manifestPath := "./public/build/manifest.json"

	// check if the manifest file exists, if not, rename it
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		err := os.Rename("./public/build/.vite/manifest.json", "./public/build/manifest.json")
		if err != nil {
			return nil
		}
	}

	i, err := inertia.NewFromFile(
		rootViewFile,
		inertia.WithVersionFromFile(manifestPath),
		inertia.WithSSR(),
	)
	if err != nil {
		log.Fatal(err)
	}

	i.ShareTemplateFunc("vite", utils.Vite(manifestPath, "/build/"))

	return i
}
