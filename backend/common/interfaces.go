package common

type User interface {
	GetID() int
	GetName() string
	GetEmail() string
	GetUsername() string
}

func (s StudentRequestDto) GetID() int {
	return s.ID
}

func (s StudentRequestDto) GetName() string {
	return s.Name
}

func (s StudentRequestDto) GetEmail() string {
	return s.Email
}

func (s StudentRequestDto) GetUsername() string {
	return s.Username
}

func (p ProfessorRequestDto) GetID() int {
	return p.ID
}

func (p ProfessorRequestDto) GetName() string {
	return p.Name
}

func (p ProfessorRequestDto) GetEmail() string {
	return p.Email
}

func (p ProfessorRequestDto) GetUsername() string {
	return p.Username
}
