package repos

import (
	"net/http"

	inertia "github.com/romsar/gonertia"
	"mocku/backend/common"
	"mocku/backend/ent"
	"mocku/backend/ent/note"
	"mocku/backend/ent/student"
)

func (r *Repo) GetStudentNotes(id int, i *inertia.Inertia, w http.ResponseWriter, req *http.Request) ([]*ent.Note, error) {
	st, err := r.GetStudentByUserId(id, i, w, req)
	if err != nil {
		r.Logger.Printf("Error getting student: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	notes, err := r.DB.Note.
		Query().
		Where(note.HasStudentWith(student.ID(st.ID))).
		All(req.Context())

	if err != nil {
		r.Logger.Printf("Error getting notes: %v", err)
		common.HandleServerErr(i, err).ServeHTTP(w, req)
		return nil, err
	}

	return notes, nil
}
