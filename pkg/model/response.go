package model

type (
	StudentResponse struct {
		ID                     string  `json:"id"`
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

	StudentsTableResponse struct {
		ID           string  `json:"id"`
		IdentityCard string  `json:"identityCard"`
		Avatar       string  `json:"avatar"`
		Name         string  `json:"name"`
		Email        string  `json:"email"`
		Phone        string  `json:"phone"`
		Career       string  `json:"career"`
		TotalAverage float64 `json:"totalAverage"`
	}

	UserResponse struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Active   bool   `json:"active"`
	}

	SelectResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	SelectResponseSubject struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		Semester int    `json:"semester"`
	}

	CareerResponse struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		LeaderName  string `json:"leader"`
		LeaderId    string `json:"leaderId"`
	}

	ProfessorResponse struct {
		ID           string `json:"id"`
		IdentityCard string `json:"identityCard"`
		Avatar       string `json:"avatar"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
	}

	SubjectResponse struct {
		ID            string                  `json:"id"`
		Name          string                  `json:"name"`
		Description   string                  `json:"description"`
		CreditUnits   int                     `json:"creditUnits"`
		Semester      int                     `json:"semester"`
		Code          string                  `json:"code"`
		PracticeHours int                     `json:"practiceHours"`
		TheoryHours   int                     `json:"theoryHours"`
		LabHours      int                     `json:"labHours"`
		TotalHours    int                     `json:"totalHours"`
		ClassSchedule map[string][]string     `json:"classSchedule"`
		ProfessorId   string                  `json:"professorId"`
		ProfessorName string                  `json:"professorName"`
		Careers       []SelectResponse        `json:"careers"`
		Prerequisites []SelectResponseSubject `json:"prerequisites"`
	}

	NoteResponse struct {
		ID      string    `json:"id"`
		Subject string    `json:"subject"`
		Notes   []float64 `json:"notes"`
		Average float64   `json:"avg"`
	}

	ScheduleSubjectResponse struct {
		ID            string              `json:"id"`
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
