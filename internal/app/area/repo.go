package area

import (
	"database/sql"

	"github.com/leandroxt/geomaps-rs/internal/entities"
)

// Repo creates area respository interface
type Repo interface {
	SaveArea(area entities.Area) (bool, error)
}

// RepoImpl create are repository implementation
type RepoImpl struct {
	db *sql.DB
}

// NewRepo creates new repo implementation
func NewRepo(db *sql.DB) Repo {
	return RepoImpl{
		db: db,
	}
}

// SaveArea saves interest users area
func (r RepoImpl) SaveArea(area entities.Area) (bool, error) {
	stmt := "insert into area (name, radius, center) values ($1, $2, ST_MakePoint($3, $4)) RETURNING id"

	var lastInsertID int

	err := r.db.QueryRow(stmt, area.Name, area.Radius, area.Center.Lng, area.Center.Lat).Scan(&lastInsertID)
	if err != nil {
		return false, err
	}

	return lastInsertID > 0, nil
}
