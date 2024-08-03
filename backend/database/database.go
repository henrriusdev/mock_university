package database

import (
	"context"
	"mocku/backend/ent"
	"mocku/backend/ent/configuration"
	"mocku/backend/ent/cycle"
	"mocku/backend/ent/role"
	"mocku/backend/ent/users"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type roles struct {
	Name        string
	Description string
}

func InsertDefaultRoles(ctx context.Context, client *ent.Client) error {
	// Define the roles
	defaultRoles := []roles{
		{"Directivo", "Rol de directivo"},
		{"Cajero", "Rol de cajero"},
		{"Control de estudio", "Rol de control de estudio"},
		{"Profesor", "Rol de profesor"},
		{"Profesor líder", "Rol de jefe de escuela"},
		{"Estudiante", "Rol de estudiante"},
	}

	// Insert roles if they don't exist
	for _, defaultRole := range defaultRoles {
		exists, err := client.Role.Query().Where(role.Name(defaultRole.Name)).Exist(ctx)
		if err != nil {
			return err
		}

		if !exists {
			if _, err := client.Role.Create().SetName(defaultRole.Name).SetDescription(defaultRole.Description).Save(ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func InsertDefaultUsers(ctx context.Context, client *ent.Client) error {
	// Check if the master admin user exists
	exists, err := client.Users.Query().Where(users.Username("admin_master")).Exist(ctx)
	if err != nil {
		return err
	}

	if !exists {
		adminRole, err := client.Role.Query().Where(role.Name("Directivo")).Only(ctx)
		if err != nil {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		if _, err := client.Users.Create().
			SetUsername("admin_master").
			SetPassword(string(hashedPassword)).
			SetEmail("admin@example.com").
			SetName("Administrador máster").
			SetAvatar("").
			SetIsActive(true).
			SetCreatedAt(time.Now()).
			SetRole(adminRole).
			Save(ctx); err != nil {
			return err
		}
	}

	return nil
}

func InsertDefaultCycle(ctx context.Context, client *ent.Client) error {
	// Check if the master admin user exists
	exists, err := client.Cycle.Query().Where(cycle.ActiveEQ(true)).Exist(ctx)
	if err != nil {
		return err
	}

	if !exists {
		if _, err := client.Cycle.Create().
			SetName("2024-2").SetActive(true).
			Save(ctx); err != nil {
			return err
		}
	}

	return nil
}

func InsertDefaultConfig(ctx context.Context, client *ent.Client) error {
	// Check if the master admin user exists
	exists, err := client.Configuration.Query().Where(configuration.HasCycleWith(cycle.Name("2024-2"))).Exist(ctx)
	if err != nil {
		return err
	}

	if !exists {
		cycle := client.Cycle.Query().Where(cycle.Active(true)).OnlyX(ctx)

		if _, err := client.Configuration.Create().
			SetCycle(cycle).
			SetBlockNotPayInscription(true).
			SetNumberFees(0).
			SetNumberNotes(0).
			SetStartRegistrationSubjects(time.Now()).
			SetEndRegistrationSubjects(time.Now()).
			Save(ctx); err != nil {
			return err
		}
	}

	return nil
}
