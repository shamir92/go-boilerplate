package repository

type pingRepository struct {
}

type IPingRepository interface {
	GetPingRepository() (bool, error)
}

func NewPingRepository() *pingRepository {
	return &pingRepository{}
}

func (repo *pingRepository) GetPingRepository() (bool, error) {
	return true, nil
}
