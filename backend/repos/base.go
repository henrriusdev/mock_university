package repos

import (
	"log"
	"mocku/backend/ent"
)

type Repo struct {
	DB     *ent.Client
	Logger *log.Logger
}
