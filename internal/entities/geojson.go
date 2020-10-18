package entities

// GeoJSON represents a GeoJson geometry
type GeoJSON struct {
	ID   int         `json:"id"`
	Name string      `json:"name"`
	JSON string      `json:"json"`
	Geom interface{} `json:"geom"`
}
