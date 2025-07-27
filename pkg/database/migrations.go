package database

import (
	"context"
	"errors"
	"log"
	"mocku/pkg/model"
	"mocku/pkg/repository"
	"mocku/pkg/service"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// InsertDefaultData inserts default data into the database using services
func InsertDefaultData(ctx context.Context, services *service.Services) error {
	log.Println("Inserting default data...")

	if err := insertDefaultRoles(ctx, services); err != nil {
		return err
	}

	if err := insertDefaultUsers(ctx, services); err != nil {
		return err
	}

	activeCycle, err := insertDefaultCycle(ctx, services)
	if err != nil {
		return err
	}

	if err := insertDefaultConfig(ctx, services, activeCycle); err != nil {
		return err
	}

	log.Println("Default data insertion completed successfully")
	return nil
}

// insertDefaultRoles inserts default roles if they don't exist
func insertDefaultRoles(ctx context.Context, services *service.Services) error {
	// Define the roles
	defaultRoles := []model.Role{
		{Name: "Directive", Description: "Directive role"},
		{Name: "Cashier", Description: "Cashier role"},
		{Name: "Study control", Description: "Study control role"},
		{Name: "Professor", Description: "Professor role"},
		{Name: "Profesor leader", Description: "Profesor leader role"},
		{Name: "Student", Description: "Student role"},
	}

	// Insert roles if they don't exist
	for _, role := range defaultRoles {
		_, err := services.Role.GetByName(ctx, role.Name)
		if err == nil {
			// Role exists, continue to next role
			continue
		}

		// Role doesn't exist, create it
		_, err = services.Role.Create(ctx, role)
		if err != nil {
			log.Printf("Error creating role: %s", role.Name)
			return err
		}
		log.Printf("Created role: %s", role.Name)
	}

	return nil
}

// insertDefaultUsers inserts default admin user if it doesn't exist
func insertDefaultUsers(ctx context.Context, services *service.Services) error {
	// Check if the master admin user exists
	_, err := services.Users.GetByUsername(ctx, "admin_master")
	if err == nil {
		// User exists, nothing to do
		return nil
	}

	// If error is not "not found", return it
	if !errors.Is(err, repository.ErrNotFound) {
		return err
	}

	// Get the admin role
	adminRole, err := services.Role.GetByName(ctx, "Directivo")
	if err != nil {
		if !errors.Is(err, repository.ErrNotFound) {
			return err
		}
		log.Println("Admin role not found, skipping admin user creation")
		return nil
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create admin user
	adminUser := model.Users{
		Username: "admin_master",
		Password: string(hashedPassword),
		Email:    "admin@example.com",
		Name:     "Administrador máster",
		Avatar:   "",
		IsActive: true,
		RoleID:   adminRole.ID,
	}

	_, err = services.Users.Create(ctx, adminUser)
	if err != nil {
		return err
	}

	log.Println("Created admin user: admin_master")
	return nil
}

// insertDefaultCycle inserts a default active cycle if none exists
func insertDefaultCycle(ctx context.Context, services *service.Services) (model.Cycle, error) {
	// Check if an active cycle exists
	activeCycle, err := services.Cycle.GetActiveCycle(ctx)
	if err == nil {
		// Active cycle exists, return it
		return activeCycle, nil
	}

	// If error is not "not found", return it
	if !errors.Is(err, repository.ErrNotFound) {
		return model.Cycle{}, err
	}

	// Create a new active cycle
	newCycle := model.Cycle{
		Name:      "2024-2",
		Active:    true,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 4, 0), // 4 months from now
	}

	activeCycle, err = services.Cycle.Create(ctx, newCycle)
	if err != nil {
		return model.Cycle{}, err
	}

	log.Println("Created default cycle: 2024-2")
	return activeCycle, nil
}

// insertDefaultConfig inserts a default configuration for the active cycle if none exists
func insertDefaultConfig(ctx context.Context, services *service.Services, activeCycle model.Cycle) error {
	_, err := services.Configuration.GetActiveCycleConfiguration(ctx)
	if err == nil {
		// Configuration exists, nothing to do
		return nil
	}

	if !errors.Is(err, repository.ErrNotFound) {
		log.Println("Error getting active cycle configuration:", err)
		return err
	}

	// Create a new configuration for the active cycle
	config := model.Configuration{
		StartRegistrationSubjects: time.Now(),
		EndRegistrationSubjects:   time.Now().AddDate(0, 1, 0), // 1 month from now
		BlockNotPayInscription:    true,
		NumberFees:                0,
		NumberNotes:               0,
		FeeDates:                  []time.Time{},
		NotesPercentages:          []float64{},
		CycleID:                   activeCycle.ID,
	}

	_, err = services.Configuration.Create(ctx, config)
	if err != nil {
		return err
	}

	log.Println("Created default configuration for cycle:", activeCycle.Name)
	return nil
}
