package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func ParseDate(date string) (time.Time, error) {
	// Parsea la fecha a time.Time
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func FormatDate(date time.Time) string {
	// Formatea la fecha a string
	return date.Format("2006-01-02")
}

func SplitCycle(cycle string) string {
	year := cycle[:4]
	semester := cycle[5:]

	if semester == "1" {
		return fmt.Sprintf("%s-%s", year, "2")
	}

	yearInt, _ := strconv.Atoi(year)
	yearInt++

	return fmt.Sprintf("%d-%s", yearInt, "1")
}

func StructToJson(dto interface{}) (map[string]interface{}, error) {
	// Convertir el DTO a JSON
	jsonData, err := json.Marshal(dto)
	if err != nil {
		return nil, fmt.Errorf("error al convertir a JSON: %w", err)
	}

	// Convertir el JSON a map[string]interface{}
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, fmt.Errorf("error al convertir JSON a map: %w", err)
	}

	return result, nil
}

func StructArrayToJson(dtos []interface{}) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	for _, dto := range dtos {
		res, err := StructToJson(dto)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}

// Unmarshal takes a JSON string and unmarshals it into the provided target variable.
func Unmarshal(data string, target interface{}) error {
	return json.Unmarshal([]byte(data), target)
}

func StringSliceToIntSlice(slice []string) ([]int, error) {
	var result []int

	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		result = append(result, i)
	}

	return result, nil
}

func ToPercentage(numberNotes int, r *http.Request) ([]float64, error) {
	notes := make([]float64, numberNotes)
	for j := 0; j < numberNotes; j++ {
		note, err := strconv.Atoi(r.FormValue(fmt.Sprintf("note-%d", j+1)))
		if err != nil {
			return nil, err
		}

		notes[j] = float64(note) / 100
	}

	return notes, nil
}

func ParseFeeDates(numberFees int, r *http.Request) ([]time.Time, error) {
	var feesDates []time.Time

	for i := 0; i < numberFees; i++ {
		feeDate, err := ParseDate(r.FormValue(fmt.Sprintf("payment-%d", i+1)))
		if err != nil {
			return nil, err
		}

		feesDates = append(feesDates, feeDate)
	}

	return feesDates, nil
}

func UploadAvatar(username string, handler *multipart.FileHeader) (string, error) {
	filePath := "./uploads/" + username + "_avatar" + filepath.Ext(handler.Filename)
	f, err := os.Create(filePath)
	if err != nil {
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	defer f.Close()

	file, err := handler.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func Average(notes []float64, percentages []float64) float64 {
	var sum float64
	for i, note := range notes {
		sum += note * percentages[i]
	}

	return sum
}
