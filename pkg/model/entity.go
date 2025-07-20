package model

import (
	"time"
)

// Users represents the user entity
type Users struct {
	ID        uint      `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	Avatar    string    `json:"avatar" db:"avatar"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Relationships
	RoleID uint `json:"role_id" db:"role_id"`
	Role   Role `json:"role" db:"-"`

	// Reverse relationships
	RequestsMade     []Request      `json:"requests_made" db:"-"`
	RequestsReceived []Request      `json:"requests_received" db:"-"`
	Blogs            []Blog         `json:"blogs" db:"-"`
	Notifications    []Notification `json:"notifications" db:"-"`
	Activities       []Activity     `json:"activities" db:"-"`
	Student          *Student       `json:"student" db:"-"`
	Professor        *Professor     `json:"professor" db:"-"`
}

// Student represents the student entity
type Student struct {
	ID                     uint    `json:"id" db:"id"`
	IdentityCard           string  `json:"identity_card" db:"identity_card"`
	BirthDate              Date    `json:"birth_date" db:"birth_date"`
	Phone                  string  `json:"phone" db:"phone"`
	Address                string  `json:"address" db:"address"`
	District               string  `json:"district" db:"district"`
	City                   string  `json:"city" db:"city"`
	PostalCode             int     `json:"postal_code" db:"postal_code"`
	CreditUnitsAccumulated int     `json:"credit_units_accumulated" db:"credit_units_accumulated"`
	TotalAverage           float64 `json:"total_average" db:"total_average"`
	Semester               int     `json:"semester" db:"semester"`

	// Relationships
	UserID   uint    `json:"user_id" db:"user_id"`
	User     Users   `json:"user" db:"-"`
	CareerID uint    `json:"career_id" db:"career_id"`
	Career   Careers `json:"career" db:"-"`

	// Reverse relationships
	Notes    []Note    `json:"notes" db:"-"`
	Payments []Payment `json:"payments" db:"-"`
}

// Professor represents the professor entity
type Professor struct {
	ID           uint   `json:"id" db:"id"`
	IdentityCard string `json:"identity_card" db:"identity_card"`
	BirthDate    Date   `json:"birth_date" db:"birth_date"`
	Phone        string `json:"phone" db:"phone"`
	Address      string `json:"address" db:"address"`

	// Relationships
	UserID uint       `json:"user_id" db:"user_id"`
	User   Users      `json:"user" db:"-"`
	BossID *uint      `json:"boss_id" db:"boss_id"`
	Boss   *Professor `json:"boss" db:"-"`

	// Reverse relationships
	Subordinates []Professor `json:"subordinates" db:"-"`
	Subjects     []Subject   `json:"subjects" db:"-"`
	Careers      []Careers   `json:"careers" db:"-"`
}

// Subject represents the subject entity
type Subject struct {
	ID            uint                `json:"id" db:"id"`
	Name          string              `json:"name" db:"name"`
	Description   string              `json:"description" db:"description"`
	CreditUnits   int                 `json:"credit_units" db:"credit_units"`
	Semester      int                 `json:"semester" db:"semester"`
	Code          string              `json:"code" db:"code"`
	PracticeHours int                 `json:"practice_hours" db:"practice_hours"`
	TheoryHours   int                 `json:"theory_hours" db:"theory_hours"`
	LabHours      int                 `json:"lab_hours" db:"lab_hours"`
	TotalHours    int                 `json:"total_hours" db:"total_hours"`
	ClassSchedule map[string][]string `json:"class_schedule" db:"class_schedule"`

	// Relationships
	ProfessorID uint      `json:"professor_id" db:"professor_id"`
	Professor   Professor `json:"professor" db:"-"`
	CareerID    uint      `json:"career_id" db:"career_id"`
	Career      Careers   `json:"career" db:"-"`

	// Reverse relationships
	Notes         []Note         `json:"notes" db:"-"`
	Prerequisites []Prerequisite `json:"prerequisites" db:"-"`
}

// Prerequisite represents a prerequisite relationship between subjects
type Prerequisite struct {
	ID             uint `json:"id" db:"id"`
	SubjectID      uint `json:"subject_id" db:"subject_id"`
	PrerequisiteID uint `json:"prerequisite_id" db:"prerequisite_id"`

	// Non-DB fields for convenience
	Subject      Subject `json:"subject" db:"-"`
	Prerequisite Subject `json:"prerequisite" db:"-"`
}

// Note represents the note entity
type Note struct {
	ID      uint      `json:"id" db:"id"`
	Notes   []float64 `json:"notes" db:"notes"` // JSON array of float64
	Average float32   `json:"average" db:"average"`

	// Relationships
	StudentID uint    `json:"student_id" db:"student_id"`
	Student   Student `json:"student" db:"-"`
	SubjectID uint    `json:"subject_id" db:"subject_id"`
	Subject   Subject `json:"subject" db:"-"`
	CycleID   uint    `json:"cycle_id" db:"cycle_id"`
	Cycle     Cycle   `json:"cycle" db:"-"`
}

// Role represents the role entity
type Role struct {
	ID          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	// Relationships
	Permissions []Permission `json:"permissions" db:"-"`

	// Reverse relationships
	Users []Users `json:"users" db:"-"`
}

// Permission represents the permission entity
type Permission struct {
	ID          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Module      string `json:"module" db:"module"`

	// Relationships
	Roles []Role `json:"roles" db:"-"`
}

// Careers represents the careers entity
type Careers struct {
	ID          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	// Relationships
	LeaderID uint       `json:"leader_id" db:"leader_id"`
	Leader   *Professor `json:"leader" db:"-"`

	// Reverse relationships
	Students []Student `json:"students" db:"-"`
	Subjects []Subject `json:"subjects" db:"-"`
}

// Cycle represents the cycle entity
type Cycle struct {
	ID        uint      `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	Active    bool      `json:"active" db:"active"`

	// Reverse relationships
	Notes []Note `json:"notes" db:"-"`
}

// Blog represents the blog entity
type Blog struct {
	ID        uint      `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Image     string    `json:"image" db:"image"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Relationships
	OwnerID uint  `json:"owner_id" db:"owner_id"`
	Owner   Users `json:"owner" db:"-"`
}

// Request represents the request entity
type Request struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	// Relationships
	RequesterID uint  `json:"requester_id" db:"requester_id"`
	Requester   Users `json:"requester" db:"-"`
	ReceiverID  uint  `json:"receiver_id" db:"receiver_id"`
	Receiver    Users `json:"receiver" db:"-"`
}

// Notification represents the notification entity
type Notification struct {
	ID        uint      `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	IsRead    bool      `json:"is_read" db:"is_read"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Relationships
	RecipientID uint  `json:"recipient_id" db:"recipient_id"`
	Recipient   Users `json:"recipient" db:"-"`
}

// Activity represents the activity entity
type Activity struct {
	ID        uint      `json:"id" db:"id"`
	Action    string    `json:"action" db:"action"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Relationships
	UserID uint  `json:"user_id" db:"user_id"`
	User   Users `json:"user" db:"-"`
}

// Payment represents the payment entity
type Payment struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Amount      float64   `json:"amount" db:"amount"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	// Relationships
	StudentID       uint          `json:"student_id" db:"student_id"`
	Student         Student       `json:"student" db:"-"`
	PaymentMethodID uint          `json:"payment_method_id" db:"payment_method_id"`
	PaymentMethod   PaymentMethod `json:"payment_method" db:"-"`
}

// PaymentMethod represents the payment method entity
type PaymentMethod struct {
	ID          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`

	// Reverse relationships
	Payments []Payment `json:"payments" db:"-"`
}

// Configuration represents the configuration entity
type Configuration struct {
	ID                        uint        `json:"id" db:"id"`
	StartRegistrationSubjects time.Time   `json:"start_registration_subjects" db:"start_registration_subjects"`
	EndRegistrationSubjects   time.Time   `json:"end_registration_subjects" db:"end_registration_subjects"`
	BlockNotPayInscription    bool        `json:"block_not_pay_inscription" db:"block_not_pay_inscription"`
	FeeDates                  []time.Time `json:"fee_dates" db:"fee_dates"`
	NumberFees                int         `json:"number_fees" db:"number_fees"`
	NumberNotes               int         `json:"number_notes" db:"number_notes"`
	NotesPercentages          []float64   `json:"notes_percentages" db:"notes_percentages"`

	// Relationships
	CycleID uint  `json:"cycle_id" db:"cycle_id"`
	Cycle   Cycle `json:"cycle" db:"-"`
}

// Module represents the module entity
type Module struct {
	ID          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
