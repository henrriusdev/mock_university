package model

type User interface {
	GetID() int
	GetName() string
	GetEmail() string
	GetUsername() string
}

func (s StudentRequest) GetID() int {
	return s.ID
}

func (s StudentRequest) GetName() string {
	return s.Name
}

func (s StudentRequest) GetEmail() string {
	return s.Email
}

func (s StudentRequest) GetUsername() string {
	return s.Username
}

func (p ProfessorRequest) GetID() int {
	return p.ID
}

func (p ProfessorRequest) GetName() string {
	return p.Name
}

func (p ProfessorRequest) GetEmail() string {
	return p.Email
}

func (p ProfessorRequest) GetUsername() string {
	return p.Username
}
