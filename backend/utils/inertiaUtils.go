package utils

import (
	"encoding/json"
	"fmt"
	inertia "github.com/romsar/gonertia"
	"log"
	"net/http"
	"os"
	"path"
)

func Vite(manifestPath, buildDir string) func(path string) (string, error) {
	f, err := os.Open(manifestPath)
	if err != nil {
		log.Fatalf("cannot open provided vite manifest file: %s", err)
	}
	defer f.Close()

	viteAssets := make(map[string]*struct {
		File   string `json:"file"`
		Source string `json:"src"`
	})
	err = json.NewDecoder(f).Decode(&viteAssets)
	// print content of viteAssets
	for k, v := range viteAssets {
		log.Printf("%s: %s\n", k, v.File)
	}

	if err != nil {
		log.Fatalf("cannot unmarshal vite manifest file to json: %s", err)
	}

	return func(p string) (string, error) {
		if val, ok := viteAssets[p]; ok {
			return path.Join("/", buildDir, val.File), nil
		}
		return "", fmt.Errorf("asset %q not found", p)
	}
}

func LoginRedirect(role int, w http.ResponseWriter, r *http.Request, i *inertia.Inertia) {
	var view string
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

	i.Redirect(w, r, view, 302)
}
