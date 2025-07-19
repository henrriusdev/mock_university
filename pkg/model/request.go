package model

import "time"

type (
	LoginRequest struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
	}

	StudentRequest struct {
		ID                     int     `form:"id"`
		Phone                  string  `form:"phone"`
		District               string  `form:"district"`
		City                   string  `form:"city"`
		Address                string  `form:"address"`
		IdentityCard           string  `form:"identityCard"`
		PostalCode             int     `form:"postalCode"`
		CreditUnitsAccumulated int     `form:"creditUnitsAccumulated"`
		Semester               int     `form:"semester"`
		TotalAverage           float64 `form:"totalAverage"`
		BirthDate              Date    `form:"birthDate"`
		CareerId               int     `form:"career"`
		Name                   string  `form:"name"`
		Email                  string  `form:"email"`
		Username               string  `form:"username"`
	}

	CareerRequest struct {
		ID          int    `form:"id"`
		Name        string `form:"name"`
		Description string `form:"description"`
		LeaderId    *int   `form:"leaderId"`
	}

	ProfessorRequest struct {
		ID           int    `form:"id"`
		IdentityCard string `form:"identityCard"`
		Name         string `form:"name"`
		Email        string `form:"email"`
		Phone        string `form:"phone"`
		Username     string `form:"username"`
		BirthDate    Date   `form:"birthDate"`
		Address      string `form:"address"`
		BossId       *int   `form:"bossId"`
	}

	SubjectRequest struct {
		ID            int    `form:"id"`
		Name          string `form:"name"`
		Description   string `form:"description"`
		CreditUnits   int    `form:"creditUnits"`
		Semester      int    `form:"semester"`
		Code          string `form:"code"`
		PracticeHours int    `form:"practiceHours"`
		TheoryHours   int    `form:"theoryHours"`
		LabHours      int    `form:"labHours"`
		TotalHours    int    `form:"totalHours"`
		ClassSchedule string `form:"classSchedule"`
		ProfessorId   int    `form:"professorId"`
		CareerId      int    `form:"careerId"`
		PreqIds       string `form:"prerequisites"`
	}

	Date struct {
		time.Time
	}
)
