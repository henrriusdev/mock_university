package handlers

import (
	"encoding/json"
	inertia "github.com/romsar/gonertia"
	"mocku/backend/ent/users"
	"mocku/backend/utils"
	"net/http"
)

func (h *Handler) HomeHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(w, err)
			return
		}

		err = i.Render(w, r, "Home/Index", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(w, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(w, err)
			return
		}

		err = i.Render(w, r, "Auth/Login", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(w, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var formData LoginDto
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			HandleServerErr(w, err)
			return
		}

		user, err := h.DB.Users.Query().Where(users.EmailEQ(formData.Email)).First(r.Context())
		if err != nil {
			HandleServerErr(w, err)
			return
		}

		if utils.CheckPassword(user.Password, formData.Password) {
			role := user.Edges.Role.ID
			if role == 1 {
				http.Redirect(w, r, "/admin", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/home", http.StatusSeeOther)
			}
		}

		careers, err := h.DB.Careers.Query().All(r.Context())
		err = i.Render(w, r, "Auth/Login", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(w, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}
