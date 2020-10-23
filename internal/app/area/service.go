package area

import (
	"database/sql"

	"github.com/leandroxt/geomaps-rs/internal/entities"
)

// Service create service interface
type Service interface {
	SaveArea(area entities.Area) (bool, error)
}

// ServiceImpl create a service implementation
type ServiceImpl struct {
	Repo Repo
}

// NewService return the service implementation
func NewService(db *sql.DB) Service {
	return ServiceImpl{
		Repo: NewRepo(db),
	}
}

// SaveArea return a city and its geometry by city name
func (s ServiceImpl) SaveArea(area entities.Area) (bool, error) {
	i, err := s.Repo.SaveArea(area)
	if err != nil {
		return false, err
	}
	return i, nil
}
