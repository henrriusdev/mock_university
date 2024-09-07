package repos

import (
	"net/http"

	inertia "github.com/romsar/gonertia"
	"mocku/backend/common"
	"mocku/backend/ent"
	"mocku/backend/ent/subject"
)

func (r *Repo) CreateSubject(subjectRequest common.SubjectRequestDto, carersIDs, subjectsIDs []int, classSchedule map[string][]string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	if subjectRequest.ID != 0 {
		return r.UpdateSubject(subjectRequest, carersIDs, subjectsIDs, classSchedule, i, w, req)
	}

	_, err := r.DB.Subject.
		Create().
		SetName(subjectRequest.Name).
		SetDescription(subjectRequest.Description).
		SetCreditUnits(subjectRequest.CreditUnits).
		SetSemester(subjectRequest.Semester).
		SetCode(subjectRequest.Code).
		SetPracticeHours(subjectRequest.PracticeHours).
		SetTheoryHours(subjectRequest.TheoryHours).
		SetLabHours(subjectRequest.LabHours).
		SetTotalHours(subjectRequest.TotalHours).
		SetClassSchedule(classSchedule).
		SetProfessorID(subjectRequest.ProfessorId).
		AddCareerIDs(carersIDs...).
		AddPrerequisiteIDs(subjectsIDs...).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error creating subject: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetSubjectById(id int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Subject, error) {
	sub, err := r.DB.Subject.
		Query().
		Where(subject.IDEQ(id)).
		WithProfessor(func(query *ent.ProfessorQuery) { query.WithUser() }).
		WithCareer().
		WithPrerequisites().
		First(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying subject: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return sub, nil
}

func (r *Repo) UpdateSubject(subjectRequest common.SubjectRequestDto, careerIDs, subjectsIDs []int, classSchedule map[string][]string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Subject.
		UpdateOneID(subjectRequest.ID).
		SetName(subjectRequest.Name).
		SetDescription(subjectRequest.Description).
		SetCreditUnits(subjectRequest.CreditUnits).
		SetSemester(subjectRequest.Semester).
		SetCode(subjectRequest.Code).
		SetPracticeHours(subjectRequest.PracticeHours).
		SetTheoryHours(subjectRequest.TheoryHours).
		SetLabHours(subjectRequest.LabHours).
		SetTotalHours(subjectRequest.TotalHours).
		SetClassSchedule(classSchedule).
		SetProfessorID(subjectRequest.ProfessorId).
		AddCareerIDs(careerIDs...).
		AddPrerequisiteIDs(subjectsIDs...).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating subject: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetSubjects(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Subject, error) {
	subjectsArray, err := r.DB.Subject.
		Query().
		WithProfessor(func(query *ent.ProfessorQuery) {
			query.WithUser()
		}).
		WithCareer().
		All(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying subjects: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return subjectsArray, nil
}

func (r *Repo) GetSubjectsNotID(id int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Subject, error) {
	subjectsArray, err := r.DB.Subject.
		Query().
		Where(subject.IDNEQ(id)).
		WithProfessor(func(query *ent.ProfessorQuery) {
			query.WithUser()
		}).
		WithCareer().
		All(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying subjects: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return subjectsArray, nil
}
