package utils

import "golang.org/x/crypto/bcrypt"

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
