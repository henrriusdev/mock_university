package repos

import (
	inertia "github.com/romsar/gonertia"
	"mocku/backend/common"
	"mocku/backend/ent"
	"mocku/backend/ent/professor"
	"net/http"
)

func (r *Repo) CreateProfessor(professorRequest common.ProfessorRequestDto, user *ent.Users, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	if professorRequest.ID != 0 {
		return r.UpdateProfessor(professorRequest, i, w, req)
	}

	_, err := r.DB.Professor.Create().
		SetAddress(professorRequest.Address).
		SetBirthDate(professorRequest.BirthDate).
		SetPhone(professorRequest.Phone).
		SetIdentityCard(professorRequest.IdentityCard).
		SetNillableBossID(professorRequest.BossId).
		SetUser(user).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error creating professor: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetProfessorById(id int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Professor, error) {
	prof, err := r.DB.Professor.Query().
		Where(professor.IDEQ(id)).
		WithUser().
		First(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying professor: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return prof, nil
}

func (r *Repo) UpdateProfessor(professorRequest common.ProfessorRequestDto, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Professor.UpdateOneID(professorRequest.ID).
		SetAddress(professorRequest.Address).
		SetBirthDate(professorRequest.BirthDate).
		SetPhone(professorRequest.Phone).
		SetIdentityCard(professorRequest.IdentityCard).
		SetNillableBossID(professorRequest.BossId).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating professor: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetProfessors(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Professor, error) {
	professorsArray, err := r.DB.Professor.Query().WithUser().All(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying professors: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return professorsArray, nil
}

func (r *Repo) GetProfessorsWithoutBoss(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Professor, error) {
	professorsArray, err := r.DB.Professor.Query().
		Where(professor.Not(professor.HasBoss())).
		WithUser().
		All(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying professors: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return professorsArray, nil
}
