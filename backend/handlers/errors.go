package handlers

import (
	inertia "github.com/romsar/gonertia"
	"net/http"
)

func HandleServerErr(i *inertia.Inertia, err error) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Errors/Server", inertia.Props{
			"error": err.Error(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}

func HandleNotFound(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Errors/NotFound", nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}

func HandleUnauthorized(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Errors/Unauthorized", nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}
