package service

import (
	"mocku/pkg/repository"
)

type Services struct {
	Users         Users
	Student       Student
	Professor     Professor
	Subject       Subject
	Note          Note
	Role          Role
	Permission    Permission
	Careers       Careers
	Cycle         Cycle
	Blog          Blog
	Request       Request
	Notification  Notification
	Activity      Activity
	Payment       Payment
	PaymentMethod PaymentMethod
	Configuration Configuration
	Module        Module
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Users:         NewUsers(repos),
		Student:       NewStudent(repos),
		Professor:     NewProfessor(repos),
		Subject:       NewSubject(repos),
		Note:          NewNote(repos),
		Role:          NewRole(repos),
		Permission:    NewPermission(repos),
		Careers:       NewCareers(repos),
		Cycle:         NewCycle(repos),
		Blog:          NewBlog(repos),
		Request:       NewRequest(repos),
		Notification:  NewNotification(repos),
		Activity:      NewActivity(repos),
		Payment:       NewPayment(repos),
		PaymentMethod: NewPaymentMethod(repos),
		Configuration: NewConfiguration(repos),
		Module:        NewModule(repos),
	}
}
