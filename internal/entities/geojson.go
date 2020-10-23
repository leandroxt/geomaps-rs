package entities

// GeoJSON represents a GeoJson geometry
type GeoJSON struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	JSON     string      `json:"json"`
	Geom     interface{} `json:"geom"`
	CentroID string      `json:"centroid"`
}

// Coordinate represents LatLng coordinates
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Geometry store information about the address
type Geometry struct {
	Location Coordinate `json:"location"`
}

// GeocoderResult treated info about address
type GeocoderResult struct {
	FormattedAddress string   `json:"formatted_address"`
	Geometry         Geometry `json:"geometry"`
}

// GeocoderResponse represents the geocoder result
type GeocoderResponse struct {
	Status  string           `json:"status"`
	Results []GeocoderResult `json:"results"`
}

// Area is a circle of interest
type Area struct {
	Name   string     `json:"name"`
	Center Coordinate `json:"center"`
	Radius float64    `json:"radius"`
}
