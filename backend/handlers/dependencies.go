package handlers

import (
	"log"

	"mocku/backend/ent"
)

type Handler struct {
	DB     *ent.Client
	Logger *log.Logger
}
