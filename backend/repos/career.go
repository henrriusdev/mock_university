package repos

import (
	inertia "github.com/romsar/gonertia"
	"mocku/backend/common"
	"mocku/backend/ent"
	"net/http"
)

func (r *Repo) GetCareers(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Careers, error) {
	careersArray, err := r.DB.Careers.Query().All(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying careers: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return careersArray, nil
}

func (r *Repo) CreateCareer(careerRequest common.CareerRequestDto, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	if careerRequest.ID != 0 {
		return r.UpdateCareer(careerRequest, i, w, req)
	}

	_, err := r.DB.Careers.Create().
		SetName(careerRequest.Name).
		SetDescription(careerRequest.Description).
		SetNillableLeaderID(careerRequest.LeaderId).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error creating career: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) UpdateCareer(careerRequest common.CareerRequestDto, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Careers.UpdateOneID(careerRequest.ID).
		SetName(careerRequest.Name).
		SetDescription(careerRequest.Description).
		SetNillableLeaderID(careerRequest.LeaderId).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating career: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}

func (r *Repo) GetCareersWithLeader(i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Careers, error) {
	careersArray, err := r.DB.Careers.Query().
		WithLeader(func(q *ent.ProfessorQuery) { q.WithUser() }).
		All(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying careers: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return careersArray, nil
}
