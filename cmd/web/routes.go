package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/api/municipio", app.GetCity)
	mux.HandleFunc("/api/municipio/search", app.SearchCities)
	mux.HandleFunc("/api/geocoder", app.geocoder)
	mux.HandleFunc("/api/area", app.saveArea)
	mux.HandleFunc("/api/areas", app.getAreas)

	return mux
}
