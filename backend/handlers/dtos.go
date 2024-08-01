package handlers

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StudentDto struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}
