package common

func FillSelectDto(array interface{}, idField, nameField string) []SelectDto {
	var result []SelectDto
	for _, item := range array.([]*interface{}) {
		if m, ok := (*item).(map[string]interface{}); ok {
			result = append(result, SelectDto{
				ID:   m[idField].(int),
				Name: m[nameField].(string),
			})
		}
	}
	return result
}
