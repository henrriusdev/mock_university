package repos

import (
	inertia "github.com/romsar/gonertia"
	"mocku/backend/common"
	"mocku/backend/ent"
	"mocku/backend/ent/users"
	"net/http"
)

func (r *Repo) GetByEmail(email string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Users, error) {
	user, err := r.DB.Users.Query().
		Where(users.EmailEQ(email)).
		WithRole().
		First(req.Context())
	if err != nil {
		r.Logger.Printf("Error querying user: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return user, nil
}

func (r *Repo) CreateUser(request common.User, password, filePath string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Users, error) {
	if request.GetID() != 0 {
		user, err := r.GetByEmail(request.GetEmail(), i, w, req)
		if err != nil {
			return nil, err
		}

		return nil, r.UpdateUser(user.ID, request, filePath, i, w, req)
	}
	user, err := r.DB.Users.Create().
		SetEmail(request.GetEmail()).
		SetUsername(request.GetUsername()).
		SetPassword(password).
		SetName(request.GetName()).
		SetAvatar(filePath).
		SetIsActive(true).
		SetRoleID(6).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error creating user: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return user, nil
}

func (r *Repo) UpdateUser(userID int, request common.User, filePath string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Users.UpdateOneID(userID).
		SetEmail(request.GetEmail()).
		SetUsername(request.GetUsername()).
		SetName(request.GetName()).
		SetAvatar(filePath).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating user: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}
