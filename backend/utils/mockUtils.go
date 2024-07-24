package utils

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func HashPassword(password string) (string, error) {
	// Genera un hash de la contraseña utilizando bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(hashedPassword, password string) bool {
	// Compara la contraseña en texto plano con el hash
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func ValidatePassword(password string) []string {
	// Valida que la contraseña tenga al menos 6 caracteres, 1 mayúscula, 1 minúscula, 1 número, 1 caracter especial, minimo 8 caracteres de longitud y no contenga espacios
	var errors []string

	if len(password) < 8 {
		errors = append(errors, "La contraseña debe tener al menos 8 caracteres")
	}

	if !regexp.MustCompile(`[A-Z]`).Match([]byte(password)) {
		errors = append(errors, "La contraseña debe tener al menos una mayúscula")
	}

	if !regexp.MustCompile(`[a-z]`).Match([]byte(password)) {
		errors = append(errors, "La contraseña debe tener al menos una minúscula")
	}

	if !regexp.MustCompile(`[0-9]`).Match([]byte(password)) {
		errors = append(errors, "La contraseña debe tener al menos un número")
	}

	if !regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).Match([]byte(password)) {
		errors = append(errors, "La contraseña debe tener al menos un caracter especial")
	}

	if regexp.MustCompile(`\s`).Match([]byte(password)) {
		errors = append(errors, "La contraseña no debe contener espacios")
	}

	return errors
}

func ValidateEmail(email string) bool {
	// Valida que el email tenga un formato válido con regex
	ok, err := regexp.Match(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, []byte(email))
	if err != nil {
		return false
	}

	return ok
}

func ValidatePhone(phone string) bool {
	// Valida que el teléfono tenga un formato válido con regex
	ok, err := regexp.Match(`^\+?[\d]{1,3}[\d]{9,15}$`, []byte(phone))
	if err != nil {
		return false
	}

	return ok
}
