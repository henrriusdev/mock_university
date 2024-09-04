package handlers

import (
	"log"
	"mocku/backend/repos"
)

type Handler struct {
	Repo   repos.Repo
	Logger *log.Logger
}
