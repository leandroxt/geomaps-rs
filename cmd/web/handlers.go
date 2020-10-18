package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func (app *application) GetCity(w http.ResponseWriter, r *http.Request) {
	cityID, err := strconv.Atoi(r.URL.Query().Get("cityID"))
	if err != nil {
		app.errorLog.Println("Erro ao recuperar city ID: ", cityID)
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	app.infoLog.Println("Buscando município com o ID: ", cityID)
	gj, err := city.NewServiceImpl(app.db).GetCity(cityID)

	if err != nil {
		app.errorLog.Println("Erro ao buscar município com ID: ", cityID)
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(gj)
}

func (app *application) SearchCities(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	gj, err := city.NewServiceImpl(app.db).SearchCities(name)

	if err != nil {
		app.errorLog.Println("Erro ao buscar municípios com o nome: ", name)
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(gj)
}
