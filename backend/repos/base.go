package repos

import (
	"log"
	"mocku/backend/ent"
)

type Repo struct {
	DB     *ent.Client
	Logger *log.Logger
}

func NewRepo(client *ent.Client, logger *log.Logger) *Repo {
	return &Repo{
		DB:     client,
		Logger: logger,
	}
}
