package handlers

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentDtoI struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Avatar       string  `json:"avatar"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Career       string  `json:"career"`
	TotalAverage float64 `json:"total_average"`
}
