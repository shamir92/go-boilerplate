package entities

import (
	"database/sql"
	"time"
)

type Invitation struct {
	Id                 string `gorm:"type:char(36);primary_key;" json:"uuid"`
	InvitationStatusId string `gorm:"type:char(36);"`
	MemberId           string
	GatheringId        string
	CreatedAt          time.Time    `gorm:"autoCreateTime;not null" json:"created_at"`
	UpdatedAt          time.Time    `gorm:"autoUpdateTime;not null" json:"updated_at"`
	DeletedAt          sql.NullTime `gorm:"index" json:"deleted_at"`
}
