package common

import "time"

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

type SubjectDto struct {
	ID            int                 `json:"id"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	CreditUnits   int                 `json:"creditUnits"`
	Semester      int                 `json:"semester"`
	Code          string              `json:"code"`
	PracticeHours int                 `json:"practiceHours"`
	TheoryHours   int                 `json:"theoryHours"`
	LabHours      int                 `json:"labHours"`
	TotalHours    int                 `json:"totalHours"`
	ClassSchedule map[string][]string `json:"classSchedule"`
	ProfessorId   int                 `json:"professorId"`
	ProfessorName string              `json:"professorName"`
	Careers       []SelectDto         `json:"careers"`
}

type StudentRequestDto struct {
	ID                     int       `form:"id"`
	Phone                  string    `form:"phone"`
	District               string    `form:"district"`
	City                   string    `form:"city"`
	Address                string    `form:"address"`
	IdentityCard           string    `form:"identityCard"`
	PostalCode             int       `form:"postalCode"`
	CreditUnitsAccumulated int       `form:"creditUnitsAccumulated"`
	Semester               int       `form:"semester"`
	TotalAverage           float64   `form:"totalAverage"`
	BirthDate              time.Time `form:"birthDate"`
	CareerId               int       `form:"career"`
	Name                   string    `form:"name"`
	Email                  string    `form:"email"`
	Username               string    `form:"username"`
}

type CareerRequestDto struct {
	ID          int    `form:"id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	LeaderId    *int   `form:"leaderId"`
}
