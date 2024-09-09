package common

import "time"

type (
	LoginDto struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
	}

	StudentDto struct {
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

	StudentsTableDto struct {
		ID           int     `json:"id"`
		IdentityCard string  `json:"identityCard"`
		Avatar       string  `json:"avatar"`
		Name         string  `json:"name"`
		Email        string  `json:"email"`
		Phone        string  `json:"phone"`
		Career       string  `json:"career"`
		TotalAverage float64 `json:"totalAverage"`
	}

	UserDto struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Active   bool   `json:"active"`
	}

	SelectDto struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	SelectDtoSubject struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		Semester int    `json:"semester"`
	}

	CareerDto struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		LeaderName  string `json:"leader"`
		LeaderId    int    `json:"leaderId"`
	}

	ProfessorDto struct {
		ID           int    `json:"id"`
		IdentityCard string `json:"identityCard"`
		Avatar       string `json:"avatar"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
	}

	SubjectDto struct {
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
		Prerequisites []SelectDto         `json:"prerequisites"`
	}

	StudentRequestDto struct {
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

	CareerRequestDto struct {
		ID          int    `form:"id"`
		Name        string `form:"name"`
		Description string `form:"description"`
		LeaderId    *int   `form:"leaderId"`
	}

	ProfessorRequestDto struct {
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

	SubjectRequestDto struct {
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
		CareerIds     string `form:"careers"`
		PreqIds       string `form:"prerequisites"`
	}

	Date struct {
		time.Time
	}

	NoteDto struct {
		ID      int       `json:"id"`
		Subject string    `json:"subject"`
		Notes   []float64 `json:"notes"`
		Average float64   `json:"avg"`
	}

	ScheduleSubjectDto struct {
		ID            int                 `json:"id"`
		Name          string              `json:"name"`
		Description   string              `json:"description"`
		Code          string              `json:"code"`
		Semester      int                 `json:"semester"`
		Credits       int                 `json:"credits"`
		PHours        int                 `json:"pHours"`
		THours        int                 `json:"tHours"`
		LHours        int                 `json:"lHours"`
		ClassSchedule map[string][]string `json:"classSchedule"`
	}
)
