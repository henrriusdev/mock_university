package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"mocku/backend/ent"
	"reflect"
	"strings"
	"time"
)

// getNestedFieldValue obtiene el valor de un campo anidado (como "User.Name") usando reflection
func getNestedFieldValue(item interface{}, fieldPath string) (reflect.Value, bool) {
	fields := strings.Split(fieldPath, ".")
	value := reflect.ValueOf(item)

	// Desreferencia el puntero si es necesario
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for _, field := range fields {
		if value.Kind() == reflect.Ptr {
			value = value.Elem() // Desreferencia el puntero si es necesario
		}

		// Obtenemos el campo por nombre
		value = value.FieldByName(field)

		// Verificamos si el valor es válido
		if !value.IsValid() {
			return reflect.Value{}, false
		}
	}

	return value, true
}

func FillSelectDto[T any](array []*T, idField, nameField string) []SelectDto {
	var result []SelectDto
	for _, item := range array {
		if item == nil {
			continue // Si el item es nil, saltamos esta iteración
		}

		// Obtener los valores de los campos, incluyendo los campos anidados
		idFieldValue, idOk := getNestedFieldValue(item, idField)
		nameFieldValue, nameOk := getNestedFieldValue(item, nameField)

		// Verificamos que ambos campos existan y no sean nil
		if !idOk || !nameOk || idFieldValue.IsZero() || nameFieldValue.IsZero() {
			continue // Si alguno de los campos no es válido o es nil, lo saltamos
		}

		// Agregamos los valores al slice de SelectDto
		result = append(result, SelectDto{
			ID:   idFieldValue.Interface().(int),
			Name: nameFieldValue.Interface().(string),
		})
	}
	return result
}

func FillSelectDtoSubject(array []*ent.Subject) []SelectDtoSubject {
	var result []SelectDtoSubject
	for _, item := range array {
		if item == nil {
			continue // Si el item es nil, saltamos esta iteración
		}

		// Agregamos los valores al slice de SelectDto
		result = append(result, SelectDtoSubject{
			ID:       item.ID,
			Name:     item.Name,
			Code:     item.Code,
			Semester: item.Semester,
		})
	}
	return result
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d *Date) UnmarshalParam(param string) error {
	// Parse the incoming date string in the format YYYY-MM-DD
	t, err := time.Parse("2006-01-02", param)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("type assertion to time.Time failed")
	}
	d.Time = t
	return nil
}

func (d Date) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}
