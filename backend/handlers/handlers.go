package handlers

import (
	inertia "github.com/romsar/gonertia"
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
