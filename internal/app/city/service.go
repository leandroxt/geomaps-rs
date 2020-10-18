package city

import (
	"database/sql"

	"github.com/leandroxt/geomaps-rs/internal/entities"
)

// Service create service interface
type Service interface {
	GetCity(name string) (*entities.GeoJSON, error)
	SearchCities(name string) ([]*entities.GeoJSON, error)
}

// ServiceImpl create a service implementation
type ServiceImpl struct {
	Repo Repo
}

// NewServiceImpl return the service implementation
func NewServiceImpl(db *sql.DB) Service {
	return ServiceImpl{
		Repo: NewRepoImpl(db),
	}
}

// GetCity return a city and its geometry by city name
func (s ServiceImpl) GetCity(name string) (*entities.GeoJSON, error) {
	gj, err := s.Repo.GetCity(name)
	if err != nil {
		return nil, err
	}
	return gj, nil
}

// SearchCities return a list of city based name search
func (s ServiceImpl) SearchCities(name string) ([]*entities.GeoJSON, error) {
	gj, err := s.Repo.SearchCities(name)
	if err != nil {
		return nil, err
	}
	return gj, nil
}