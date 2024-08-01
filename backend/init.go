package backend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"mocku/backend/database"
	"mocku/backend/ent"
	"mocku/backend/handlers"
	"mocku/backend/utils"

	inertia "github.com/romsar/gonertia"

	_ "github.com/lib/pq"
)

// MountApp mounts the application, initialize inertia.js and database,  and starts the server
func MountApp() {
	i := initInertia()
	mux := http.NewServeMux()

	client := initDatabase()
	handler := handlers.Handler{
		DB: client,
	}

	// Routes
	mux.Handle("/", i.Middleware(handler.HomeHandler(i)))
	mux.Handle("/login", i.Middleware(handler.LoginHandler(i)))
	mux.Handle("/login_post", i.Middleware(handler.LoginPostHandler(i)))

	// Dashboard routes
	mux.Handle("/directive", i.Middleware(handler.DirectiveDashHandler(i)))
	mux.Handle("/payment", i.Middleware(handler.PaymentsDashHandler(i)))
	mux.Handle("/student", i.Middleware(handler.StudentDashHandler(i)))
	mux.Handle("/professor", i.Middleware(handler.ProfessorDashHandler(i)))
	mux.Handle("/control", i.Middleware(handler.ControlDashHandler(i)))

	// Directives routes
	mux.Handle("/settings", i.Middleware(handler.SettingsHandler(i)))
	mux.Handle("/settings/notes", i.Middleware(handler.SettingsNotesPostHandler(i)))
	mux.Handle("/settings/notes/percentages", i.Middleware(handler.SettingsNotesPercentagePostHandler(i)))
	mux.Handle("/settings/payment", i.Middleware(handler.SettingsPaymentsPostHandler(i)))
	mux.Handle("/settings/dates", i.Middleware(handler.SettingsDatesPostHandler(i)))
	mux.Handle("/settings/payment/dates", i.Middleware(handler.SettingsPaymentsDatesPostHandler(i)))
	mux.Handle("/settings/cycle", i.Middleware(handler.SettingsCyclePostHandler(i)))

	// API routes
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	http.ListenAndServe(":3000", mux)
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
	manifestPath := "./public/build/manifest.json"

	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		err := os.Rename("./public/build/.vite/manifest.json", "./public/build/manifest.json")
		if err != nil {
			return nil
		}
	}

	i, err := inertia.NewFromFile(
		"resources/views/root.html",
		inertia.WithVersionFromFile(manifestPath),
		inertia.WithSSR(),
	)
	if err != nil {
		log.Fatal(err)
	}

	i.ShareTemplateFunc("vite", utils.Vite(manifestPath, "/build/"))

	return i
}
