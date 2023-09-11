package usecase

import (
	"database/sql"
	entities "simple-invitation/domain/entities"
	repository "simple-invitation/domain/repository"

	"github.com/google/uuid"
)

type memberUsecase struct {
	memberRepository repository.IMemberRepository
}

type IMemberUsecase interface {
	StoreNewMeber(email string, firstname string, lastname string) error
	FetchMemberByEmail(email string) (*entities.Member, error)
	FetchMemberById(id string) (*entities.Member, error)
	FetchAll() ([]*entities.Member, error)
}

func NewMemberUsecase(dbReader, dbWriter *sql.DB) *memberUsecase {
	var memberRepository repository.IMemberRepository = repository.NewMemberRepository(dbReader, dbWriter)
	return &memberUsecase{
		memberRepository: memberRepository,
	}
}
func (muc *memberUsecase) FetchAll() ([]*entities.Member, error) {
	return muc.memberRepository.FetchAll()
}

func (muc *memberUsecase) FetchMemberByEmail(email string) (*entities.Member, error) {
	return muc.memberRepository.FetchMemberByEmail(email)
}

func (muc *memberUsecase) FetchMemberById(idMember string) (*entities.Member, error) {
	id, err := uuid.Parse(idMember)
	if err != nil {
		return nil, err
	}
	return muc.memberRepository.FetchMemberById(id)
}

func (muc *memberUsecase) StoreNewMeber(email string, firstname string, lastname string) error {

	var res = muc.memberRepository.StoreNewMeber(email, firstname, lastname)
	if res != nil {
		panic(res)
	}
	return nil

}
