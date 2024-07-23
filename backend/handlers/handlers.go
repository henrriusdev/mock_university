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

		var view string
		if utils.CheckPassword(user.Password, formData.Password) {
			role := user.Edges.Role.ID
			switch role {
			case 1:
				view = "Admin/Index"
			case 2:
				view = "Payments/Index"
			case 3:
				view = "Control/Index"
			case 4, 5:
				view = "Professor/Index"
			case 6:
				view = "Student/Index"
			}
		}

		println(view)

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
