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

func (r *Repo) CreateUserStudent(studentRequest common.StudentRequestDto, password, filePath string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) (*ent.Users, error) {
	if studentRequest.ID != 0 {
		user, err := r.GetByEmail(studentRequest.Email, i, w, req)
		if err != nil {
			return nil, err
		}

		return nil, r.UpdateUserStudent(user.ID, studentRequest, filePath, i, w, req)
	}
	user, err := r.DB.Users.Create().
		SetEmail(studentRequest.Email).
		SetUsername(studentRequest.Username).
		SetPassword(password).
		SetName(studentRequest.Name).
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

func (r *Repo) UpdateUserStudent(userID int, studentRequest common.StudentRequestDto, filePath string, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) error {
	_, err := r.DB.Users.UpdateOneID(userID).
		SetEmail(studentRequest.Email).
		SetUsername(studentRequest.Username).
		SetName(studentRequest.Name).
		SetAvatar(filePath).
		Save(req.Context())
	if err != nil {
		r.Logger.Printf("Error updating user: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return err
	}

	return nil
}
