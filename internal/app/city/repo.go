package city

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/leandroxt/geomaps-rs/internal/entities"
)

// Repo creates city respository interface
type Repo interface {
	GetCity(cityID int) (*entities.GeoJSON, error)
	SearchCities(name string) ([]*entities.GeoJSON, error)
}

// RepoImpl create city repository implementation
type RepoImpl struct {
	db *sql.DB
}

// NewRepoImpl creates new repo implementation
func NewRepoImpl(db *sql.DB) Repo {
	return RepoImpl{
		db: db,
	}
}

// GetCity retrieve geometry in database
func (r RepoImpl) GetCity(cityID int) (*entities.GeoJSON, error) {
	stmt := "select id, name, geojson, ST_AsGeoJSON(geom) from municipio where id = $1"

	gj := &entities.GeoJSON{}
	err := r.db.QueryRow(stmt, cityID).Scan(&gj.ID, &gj.Name, &gj.JSON, &gj.Geom)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no_city_found")
		}
		return nil, err
	}

	return gj, nil
}

// SearchCities return a list of city based name search
func (r RepoImpl) SearchCities(name string) ([]*entities.GeoJSON, error) {
	stmt := "select distinct id, name from municipio where name ILIKE $1 order by 1"
	gj := []*entities.GeoJSON{}

	rows, err := r.db.Query(stmt, fmt.Sprintf("%%%s%%", name))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		g := &entities.GeoJSON{}
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			return nil, err
		}

		gj = append(gj, g)
	}

	return gj, nil
}
