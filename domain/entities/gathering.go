package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Gathering struct {
	Id         uuid.UUID
	Creator    *Member
	ScheduleAt time.Time
	Name       string
	Location   string
	Type       *GatheringType
	Attendees  []*Member
}

func (g *Gathering) CheckAttendeesIsNotNil(attendees []*Member) error {
	if attendees != nil {
		return nil
	}
	return fmt.Errorf("attendees is nil")
}

func (g *Gathering) CheckCreatorIsNotNil(creator *Member) error {
	if creator != nil {
		return nil
	}
	return fmt.Errorf("creator is nil")
}
