package usecase

import (
	"database/sql"
	entities "simple-invitation/domain/entities"
	repository "simple-invitation/domain/repository"

	"github.com/google/uuid"
)

type gatheringUsecase struct {
	gatheringRepository repository.IGatheringRepository
}

type IGatheringUsecase interface {
	// CreateNewGathering(
	// 	attendees []uuid.UUID,
	// 	creator uuid.UUID,
	// 	gatheringType string,
	// 	name string,
	// 	location string,
	// 	scheduled_at string) (*entities.Gathering, error)
	FetchGatheringTypes() ([]*entities.GatheringType, error)
	StoreNewGatheringType(name string) (*entities.GatheringType, error)
	FetchGatheringTypeById(id uuid.UUID) (*entities.GatheringType, error)
	FetchInvitationStatuses() ([]*entities.InvitationStatus, error)
	FetchInvitationStatusById(id uuid.UUID) (*entities.InvitationStatus, error)
	StoreInvitationStatus(name string) (*entities.InvitationStatus, error)
}

func NewGatheringUsecase(dbReader, dbWriter *sql.DB) *gatheringUsecase {
	var gatheringRepository repository.IGatheringRepository = repository.NewGatheringRepository(dbReader, dbWriter)
	return &gatheringUsecase{
		gatheringRepository: gatheringRepository,
	}
}

func (guc *gatheringUsecase) FetchGatheringTypes() ([]*entities.GatheringType, error) {
	return guc.gatheringRepository.FetchGatheringTypes()
}

func (guc *gatheringUsecase) StoreNewGatheringType(name string) (*entities.GatheringType, error) {
	return guc.gatheringRepository.StoreNewGatheringType(name)
}

func (guc *gatheringUsecase) FetchGatheringTypeById(id uuid.UUID) (*entities.GatheringType, error) {
	return guc.gatheringRepository.FetchGatheringTypeById(id)
}

func (guc *gatheringUsecase) FetchInvitationStatuses() ([]*entities.InvitationStatus, error) {
	return guc.gatheringRepository.FetchInvitationStatuses()
}

func (guc *gatheringUsecase) FetchInvitationStatusById(id uuid.UUID) (*entities.InvitationStatus, error) {
	return guc.gatheringRepository.FetchInvitationStatusById(id)
}

func (guc *gatheringUsecase) StoreInvitationStatus(name string) (*entities.InvitationStatus, error) {
	return guc.gatheringRepository.StoreInvitationStatus(name)
}

// func (guc *gatheringUsecase) CreateNewGathering(
// 	attendees []uuid.UUID,
// 	creator uuid.UUID,
// 	gatheringType string,
// 	name string,
// 	location string,
// 	scheduled_at string) (*entities.Gathering, error) {
// 	return guc.gatheringRepository.CreateNewGathering(attendees, creator, gatheringType, name, location, scheduled_at)
// }
