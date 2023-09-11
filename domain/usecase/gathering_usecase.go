package usecase

import (
	"database/sql"
	"log"
	entities "simple-invitation/domain/entities"
	repository "simple-invitation/domain/repository"
	"time"

	"github.com/google/uuid"
)

type gatheringUsecase struct {
	gatheringRepository repository.IGatheringRepository
	memberRepository    repository.IMemberRepository
}

type IGatheringUsecase interface {
	CreateNewGathering(
		idAttendees []uuid.UUID,
		idCreator uuid.UUID,
		gatheringType string,
		name string,
		location string,
		scheduled_at time.Time) (*entities.Gathering, error)
	FetchGatheringTypes() ([]*entities.GatheringType, error)
	StoreNewGatheringType(name string) (*entities.GatheringType, error)
	FetchGatheringTypeById(id uuid.UUID) (*entities.GatheringType, error)
	FetchInvitationStatuses() ([]*entities.InvitationStatus, error)
	FetchInvitationStatusById(id uuid.UUID) (*entities.InvitationStatus, error)
	StoreInvitationStatus(name string) (*entities.InvitationStatus, error)
}

func NewGatheringUsecase(dbReader, dbWriter *sql.DB) *gatheringUsecase {
	var gatheringRepository repository.IGatheringRepository = repository.NewGatheringRepository(dbReader, dbWriter)
	var memberRepository repository.IMemberRepository = repository.NewMemberRepository(dbReader, dbWriter)
	return &gatheringUsecase{
		gatheringRepository: gatheringRepository,
		memberRepository:    memberRepository,
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

func (guc *gatheringUsecase) CreateNewGathering(
	idAttendees []uuid.UUID,
	idCreator uuid.UUID,
	gatheringType string,
	name string,
	location string,
	scheduled_at time.Time) (*entities.Gathering, error) {

	attendees, err := guc.memberRepository.FetchMemberByListId(idAttendees)
	if err != nil {
		return nil, err
	}

	creator, err := guc.memberRepository.FetchMemberById(idCreator)
	log.Println(creator)
	if err != nil {
		return nil, err
	}

	// Create Invitation and Send Slack Invitation
	return guc.gatheringRepository.CreateNewGathering(attendees, creator, gatheringType, name, location, scheduled_at)
}
