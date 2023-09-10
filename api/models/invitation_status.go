package entities

import (
	"database/sql"
	"time"
)

type InvitationStatus struct {
	Name      string       `json:"name"`
	Id        string       `json:"uuid"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
