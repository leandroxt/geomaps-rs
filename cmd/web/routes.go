package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/api/municipio", app.GetCity)
	mux.HandleFunc("/api/municipio/search", app.SearchCities)

	return mux
}
