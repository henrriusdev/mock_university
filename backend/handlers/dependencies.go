package handlers

import "mocku/backend/ent"

type Handler struct {
	DB *ent.Client
	// Add more dependencies here
}
