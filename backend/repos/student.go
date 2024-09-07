package repos

import (
	"mocku/backend/ent/users"
	"net/http"

	inertia "github.com/romsar/gonertia"
	"mocku/backend/common"
	"mocku/backend/ent"
	"mocku/backend/ent/student"
)

func (r *Repo) CreateStudent(studentRequest common.StudentRequestDto, user *ent.Users, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	if studentRequest.ID != 0 {
		return r.UpdateStudent(studentRequest, i, w, req)
	}

	_, err := r.DB.Student.
		Create().
		SetPhone(studentRequest.Phone).
		SetDistrict(studentRequest.District).
		SetCity(studentRequest.City).
		SetPostalCode(studentRequest.PostalCode).
		SetAddress(studentRequest.Address).
		SetIdentityCard(studentRequest.IdentityCard).
		SetBirthDate(studentRequest.BirthDate.Time).
		SetCreditUnitsAccumulated(studentRequest.CreditUnitsAccumulated).
		SetSemester(studentRequest.Semester).
		SetTotalAverage(studentRequest.TotalAverage).
		SetCareerID(studentRequest.CareerId).
		SetUser(user).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error creating student: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetStudentById(id int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Student, error) {
	st, err := r.DB.Student.
		Query().
		Where(student.IDEQ(id)).
		WithUser().
		WithCareer().
		First(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying student: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return st, nil
}

func (r *Repo) UpdateStudent(studentRequest common.StudentRequestDto, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Student.
		UpdateOneID(studentRequest.ID).
		SetPhone(studentRequest.Phone).
		SetDistrict(studentRequest.District).
		SetCity(studentRequest.City).
		SetPostalCode(studentRequest.PostalCode).
		SetAddress(studentRequest.Address).
		SetIdentityCard(studentRequest.IdentityCard).
		SetBirthDate(studentRequest.BirthDate.Time).
		SetCreditUnitsAccumulated(studentRequest.CreditUnitsAccumulated).
		SetSemester(studentRequest.Semester).
		SetTotalAverage(studentRequest.TotalAverage).
		SetCareerID(studentRequest.CareerId).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating student: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetStudents(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Student, error) {
	students, err := r.DB.Student.
		Query().
		WithUser().
		WithCareer().
		All(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying students: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return students, nil
}

func (r *Repo) GetStudentByUserId(userId int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Student, error) {
	st, err := r.DB.Student.
		Query().
		Where(student.HasUserWith(users.IDEQ(userId))).
		WithUser().
		WithCareer().
		First(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying student: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return st, nil
}
