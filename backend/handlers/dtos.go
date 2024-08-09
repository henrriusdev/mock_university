package handlers

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentDto struct {
	ID                     int     `json:"id"`
	IdentityCard           string  `json:"identityCard"`
	BirthDate              string  `json:"birthDate"`
	Phone                  string  `json:"phone"`
	Address                string  `json:"address"`
	District               string  `json:"district"`
	City                   string  `json:"city"`
	PostalCode             int     `json:"postalCode"`
	CreditUnitsAccumulated int     `json:"creditUnitsAccumulated"`
	TotalAverage           float64 `json:"totalAverage"`
}

type StudentsTableDto struct {
	ID           int     `json:"id"`
	IdentityCard string  `json:"identityCard"`
	Avatar       string  `json:"avatar"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Career       string  `json:"career"`
	TotalAverage float64 `json:"totalAverage"`
}

type UserDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Active   bool   `json:"active"`
}
