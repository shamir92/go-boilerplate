package repository

import (
	"database/sql"
	entities "simple-invitation/domain/entities"

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
	// CreateNewGathering(
	// 	attendees []uuid.UUID,
	// 	creator uuid.UUID,
	// 	gatheringType string,
	// 	name string,
	// 	location string,
	// 	scheduled_at string,
	// ) (*entities.Gathering, error)
	FetchGatheringTypes() ([]*entities.GatheringType, error)
	FetchGatheringTypeById(idGatheringType uuid.UUID) (*entities.GatheringType, error)
	StoreNewGatheringType(name string) (*entities.GatheringType, error)
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

// func (gr *gatheringRepository) CreateNewGathering(
// 	attendees []uuid.UUID,
// 	creator uuid.UUID,
// 	gatheringType string,
// 	name string,
// 	location string,
// 	scheduled_at string,
// ) (*entities.Gathering, error) {

// 	// stmt, err := gr.dbWriter.Prepare("insert into ")
// 	// stmt, err := gr.dbWriter.Prepare("INSERT INTO gathering (id, creator, gathering_type, schedule_at, name, location) VALUES (?,?, ?, ?)")
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// idGathering := uuid.New()

// 	// // Execute the prepared statement.
// 	// _, err = stmt.Exec(idGathering.String(), email, firstname, lastname)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// // Close the prepared statement.
// 	// err = stmt.Close()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	return nil, nil
// }

// func (gr *gatheringRepository) createGathering() (*entities.Gathering, error) {

// 	// stmt, err := gr.dbWriter.Prepare("insert into gathering (id, creato)")
// 	return nil, nil
// }

// func (gr *gatheringRepository) getGatheringTypeByName(gatheringTypeName string) (*entities.GatheringType, error) {

// 	var id, name string
// 	var created_at, updated_at, deleted_at sql.NullTime
// 	err := gr.dbReader.QueryRow("SELECT id, name, created_at, updated_at, deleted_at FROM gathering_type WHERE name=?", gatheringTypeName).Scan(
// 		&id, &name, &created_at, &updated_at, &deleted_at)
// 	if err != nil {
// 		return nil, err
// 	}

// 	gatheringTypeId, err := uuid.Parse(id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &entities.GatheringType{
// 		Id:        gatheringTypeId,
// 		Name:      name,
// 		CreatedAt: created_at,
// 		UpdatedAt: updated_at,
// 		DeletedAt: deleted_at,
// 	}, nil
// }

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
