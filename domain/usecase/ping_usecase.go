package usecase

import "simple-invitation/domain/repository"

type pingUsecase struct {
	pingRepository repository.IPingRepository
}

type IPingUsecase interface {
	PingUsecase() (bool, error)
	PingRepository() (bool, error)
}

func NewPingUsecase() *pingUsecase {
	var pingRepository repository.IPingRepository = repository.NewPingRepository()
	return &pingUsecase{
		pingRepository: pingRepository,
	}
}

func (puc *pingUsecase) PingUsecase() (bool, error) {
	return true, nil
}

func (puc *pingUsecase) PingRepository() (bool, error) {
	return puc.pingRepository.GetPingRepository()
}
