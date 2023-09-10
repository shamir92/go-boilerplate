package entities

import (
	"database/sql"

	"github.com/google/uuid"
)

type GatheringType struct {
	Name      string       `json:"name"`
	Id        uuid.UUID    `json:"uuid"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
