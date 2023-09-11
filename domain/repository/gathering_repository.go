package repository

import (
	"database/sql"
	"log"
	entities "simple-invitation/domain/entities"
	"time"

	"github.com/google/uuid"
)

type gatheringRepository struct {
	dbWriter  *sql.DB
	dbReader  *sql.DB
	gathering entities.Gathering
}

type IGatheringRepository interface {
	// Create a new member.
	FetchDBReader() *sql.DB
	FetchDBWriter() *sql.DB
	// For Gathering
	CreateNewGathering(
		attendees []*entities.Member,
		creator *entities.Member,
		gatheringTypeName string,
		gatheringName string,
		gatheringLocation string,
		scheduled_at time.Time,
	) (*entities.Gathering, error)

	// For Gathering Types
	FetchGatheringTypes() ([]*entities.GatheringType, error)
	FetchGatheringTypeById(idGatheringType uuid.UUID) (*entities.GatheringType, error)
	FetchGatheringTypeByName(gatheringTypeName string) (*entities.GatheringType, error)
	StoreNewGatheringType(name string) (*entities.GatheringType, error)

	// For Invitation
	FetchInvitationStatuses() ([]*entities.InvitationStatus, error)
	FetchInvitationStatusById(invitationStatusId uuid.UUID) (*entities.InvitationStatus, error)
	StoreInvitationStatus(name string) (*entities.InvitationStatus, error)
}

func NewGatheringRepository(dbReader *sql.DB, dbWriter *sql.DB) *gatheringRepository {
	return &gatheringRepository{
		dbWriter: dbWriter,
		dbReader: dbReader,
	}
}

func (gr *gatheringRepository) FetchDBWriter() *sql.DB {
	return gr.dbWriter
}

func (gr *gatheringRepository) FetchDBReader() *sql.DB {
	return gr.dbReader
}

func (gr *gatheringRepository) CreateNewGathering(
	attendees []*entities.Member,
	creator *entities.Member,
	gatheringTypeName string,
	gatheringName string,
	gatheringLocation string,
	scheduled_at time.Time,
) (*entities.Gathering, error) {
	log.Println("nyampai repository")

	log.Println(attendees[0].Id)
	log.Println(creator.Id)
	err := gr.gathering.CheckAttendeesIsNotNil(attendees)
	if err != nil {
		return nil, err
	}

	err = gr.gathering.CheckCreatorIsNotNil(creator)
	if err != nil {
		return nil, err
	}

	gatheringType, err := gr.FetchGatheringTypeByName(gatheringTypeName)
	if err != nil {
		return nil, nil
	}

	stmt, err := gr.dbWriter.Prepare("insert into gathering (id, creator,gathering_type_id, schedule_at, name, location) values (?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}

	idGathering := uuid.New()
	// Execute the prepared statement.
	_, err = stmt.Exec(idGathering, creator.Id, gatheringType.Id, scheduled_at, gatheringName, gatheringLocation)
	if err != nil {
		return nil, err
	}

	return &entities.Gathering{
		Id:         idGathering,
		Creator:    creator,
		Attendees:  attendees,
		Type:       gatheringType,
		ScheduleAt: scheduled_at,
		Name:       gatheringName,
		Location:   gatheringLocation,
	}, nil
}

func (gr *gatheringRepository) FetchGatheringTypeByName(gatheringTypeName string) (*entities.GatheringType, error) {
	var id uuid.UUID
	var name string
	var created_at, deleted_at, updated_at sql.NullTime
	err := gr.dbReader.QueryRow("select id, name, created_at, updated_at, deleted_at from gathering_type where name = ?", gatheringTypeName).Scan(
		&id, &name, &created_at, &updated_at, &deleted_at)

	if err != nil {
		return nil, err
	}

	return &entities.GatheringType{
		Id:        id,
		Name:      name,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		DeletedAt: deleted_at,
	}, nil
}

// / For gathering type
func (gr *gatheringRepository) FetchGatheringTypes() ([]*entities.GatheringType, error) {
	var gatheringTypeList []*entities.GatheringType
	result, err := gr.dbReader.Query("select id, name, created_at, updated_at, deleted_at from gathering_type")

	for result.Next() {
		var id uuid.UUID
		var name string
		var created_at, deleted_at, updated_at sql.NullTime
		result.Scan(&id, &name, &created_at, &updated_at, &deleted_at)
		// gatheringTypeId, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}

		gatheringTypeList = append(gatheringTypeList, &entities.GatheringType{
			Id:        id,
			Name:      name,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			DeletedAt: deleted_at,
		})
	}

	if err != nil {
		return nil, err
	}
	return gatheringTypeList, nil
}

func (gr *gatheringRepository) FetchGatheringTypeById(idGatheringType uuid.UUID) (*entities.GatheringType, error) {
	var id uuid.UUID
	var name string
	var created_at, deleted_at, updated_at sql.NullTime
	err := gr.dbReader.QueryRow("select id, name, created_at, updated_at, deleted_at from gathering_type where id = ?", idGatheringType).Scan(
		&id, &name, &created_at, &updated_at, &deleted_at)

	if err != nil {
		return nil, err
	}

	return &entities.GatheringType{
		Id:        id,
		Name:      name,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		DeletedAt: deleted_at,
	}, nil
}

func (gr *gatheringRepository) StoreNewGatheringType(name string) (*entities.GatheringType, error) {
	stmt, err := gr.dbWriter.Prepare("INSERT INTO gathering_type (id, name) VALUES (?,?)")
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	// Execute the prepared statement.
	_, err = stmt.Exec(id.String(), name)
	if err != nil {
		return nil, err
	}

	// Close the prepared statement.
	defer stmt.Close()
	return &entities.GatheringType{
		Id:   id,
		Name: name,
	}, nil
}

func (gr *gatheringRepository) FetchInvitationStatuses() ([]*entities.InvitationStatus, error) {
	var invitationStatusList []*entities.InvitationStatus
	result, err := gr.dbReader.Query("select id, name, created_at, updated_at, deleted_at from invitation_status")
	for result.Next() {
		var id uuid.UUID
		var name string
		var created_at, deleted_at, updated_at sql.NullTime
		result.Scan(&id, &name, &created_at, &updated_at, &deleted_at)
		if err != nil {
			return nil, err
		}
		invitationStatusList = append(invitationStatusList, &entities.InvitationStatus{
			Id:        id,
			Name:      name,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
			DeletedAt: deleted_at,
		})
	}

	if err != nil {
		return nil, err
	}
	return invitationStatusList, nil
}

func (gr *gatheringRepository) FetchInvitationStatusById(invitationStatusId uuid.UUID) (*entities.InvitationStatus, error) {
	var id uuid.UUID
	var name string
	var created_at, deleted_at, updated_at sql.NullTime

	err := gr.dbReader.QueryRow("select id, name, created_at, updated_at, deleted_at from invitation_status where id = ?",
		invitationStatusId).Scan(
		&id, &name, &created_at, &updated_at, &deleted_at)

	if err != nil {
		return nil, err
	}
	return &entities.InvitationStatus{
		Id:        id,
		Name:      name,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
		DeletedAt: deleted_at,
	}, nil
}

func (gr *gatheringRepository) StoreInvitationStatus(name string) (*entities.InvitationStatus, error) {
	stmt, err := gr.dbWriter.Prepare("INSERT INTO invitation_status (id, name) VALUES (?,?)")
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	// Execute the prepared statement.
	_, err = stmt.Exec(id.String(), name)
	if err != nil {
		return nil, err
	}

	// Close the prepared statement.
	defer stmt.Close()
	return &entities.InvitationStatus{
		Id:   id,
		Name: name,
	}, nil
}
