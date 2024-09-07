package common

import (
	"errors"
	"net/http"

	inertia "github.com/romsar/gonertia"
)

var ErrMethodNotAllowed = errors.New("method not allowed")

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
		err := i.Render(w, r, "Errors/NotFound", inertia.Props{})
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

func HandleBadRequest(i *inertia.Inertia, validErr error) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Errors/BadRequest", inertia.Props{
			"error": validErr.Error(),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}
