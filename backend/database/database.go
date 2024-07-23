package database

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"mocku/backend/ent"
	"mocku/backend/ent/role"
	"mocku/backend/ent/users"
	"time"
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
	exists, err := client.Users.Query().Where(users.Username("director")).Exist(ctx)
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
