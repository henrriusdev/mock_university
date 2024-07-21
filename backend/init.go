package backend

import (
	inertia "github.com/romsar/gonertia"
	"log"
	"mocku/backend/handlers"
	"mocku/backend/utils"
	"net/http"
	"os"
)

func MountHandlers() {
	i := initInertia()
	mux := http.NewServeMux()

	mux.Handle("/", i.Middleware(handlers.HomeHandler(i)))
	mux.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("./public/build"))))

	http.ListenAndServe(":3000", mux)
}

func initInertia() *inertia.Inertia {
	manifestPath := "./public/build/manifest.json"

	// check if the manifest file exists, if not, rename it
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		// move the manifest from ./public/build/.vite/manifest.json to ./public/build/manifest.json
		// so that the vite function can find it
		err := os.Rename("./public/build/.vite/manifest.json", "./public/build/manifest.json")
		if err != nil {
			return nil
		}
	}

	i, err := inertia.NewFromFile(
		"resources/views/root.html",
		inertia.WithVersionFromFile(manifestPath),
		inertia.WithSSR(),
	)
	if err != nil {
		log.Fatal(err)
	}

	i.ShareTemplateFunc("vite", utils.Vite(manifestPath, "/build/"))

	return i
}
