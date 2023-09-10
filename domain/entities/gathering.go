package entities

type Gathering struct {
	Id         string
	Creator    string
	ScheduleAt string
	Name       string
	Location   string
	Type       *GatheringType
	Attendee   []*Member
}
