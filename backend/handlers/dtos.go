package handlers

type LoginDto struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
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

type SelectDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CareerDto struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LeaderName  string `json:"leader"`
	LeaderId    int    `json:"leaderId"`
}

type ProfessorDto struct {
	ID           int    `json:"id"`
	IdentityCard string `json:"identityCard"`
	Avatar       string `json:"avatar"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}
