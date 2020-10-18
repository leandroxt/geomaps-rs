package main

import (
	"encoding/json"
	"net/http"

	"github.com/leandroxt/geomaps-rs/internal/app/city"
)

// Err error json
type Err struct {
	Error string `json:"error"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errorLog.Println("url path not exists")
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from GeoMaps"))
}

func (app *application) getMunicipio(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	app.infoLog.Println("Buscando município com o nome: ", name)

	gj, err := city.NewServiceImpl(app.db).GetCity(name)

	if err != nil {
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(gj)
}

func (app *application) SearchCities(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	app.infoLog.Println("Buscando municípios com o nome: ", name)

	gj, err := city.NewServiceImpl(app.db).SearchCities(name)

	if err != nil {
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(gj)
}
