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
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Home/Index", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Auth/Login", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) LoginPostHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			HandleNotFound(i).ServeHTTP(w, r)
			return
		}

		var formData LoginDto
		err := json.NewDecoder(r.Body).Decode(&formData)
		println(formData.Email, formData.Password, err, "l63")
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		user, err := h.DB.Users.Query().Where(users.EmailEQ(formData.Email)).First(r.Context())
		println(user, err, "l69")
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		careers, err := h.DB.Careers.Query().All(r.Context())
		if !utils.CheckPassword(user.Password, formData.Password) {
			err = i.Render(w, r, "Auth/Login", inertia.Props{
				"careers": careers,
				"error":   "Credenciales incorrectas",
			})

			if err != nil {
				HandleServerErr(i, err)
				return
			}

			return
		}

		var view string
		role := user.Edges.Role.ID
		switch role {
		case 1:
			view = "/directive"
		case 2:
			view = "/payments"
		case 3:
			view = "/control"
		case 4, 5:
			view = "/professor"
		case 6:
			view = "/student"
		}

		println(view)

		i.Redirect(w, r, view, 302)
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) DirectiveDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Directive/Dash", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) PaymentsDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Payment/Dash", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) ControlDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Control/Dash", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) ProfessorDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Professor/Dash", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}

func (h *Handler) StudentDashHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		careers, err := h.DB.Careers.Query().All(r.Context())
		if err != nil {
			HandleServerErr(i, err)
			return
		}

		err = i.Render(w, r, "Student/Dash", inertia.Props{
			"careers": careers,
		})

		if err != nil {
			HandleServerErr(i, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}
