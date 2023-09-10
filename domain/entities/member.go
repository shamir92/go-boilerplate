package entities

import (
	"database/sql"
)

type Member struct {
	Firstname string       `db:"firstname" json:"firstname"`
	Lastname  string       `db:"lastname" json:"lastname"`
	Email     string       `db:"email" json:"email"`
	Id        string       `db:"id" json:"id"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}
