package handlers

import (
	inertia "github.com/romsar/gonertia"
	"net/http"
)

func (h *Handler) HomeHandler(i *inertia.Inertia) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := i.Render(w, r, "Home/Index", inertia.Props{
			"text": "Inertia.js with Svelte and Go! 💙",
		})

		if err != nil {
			HandleServerErr(w, err)
			return
		}
	}

	return http.HandlerFunc(fn)
}
