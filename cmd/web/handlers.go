package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/leandroxt/geomaps-rs/internal/entities"

	"github.com/leandroxt/geomaps-rs/internal/app/area"
	"github.com/leandroxt/geomaps-rs/internal/app/city"
	"github.com/leandroxt/geomaps-rs/internal/app/geocoder"
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
	gj, err := city.NewService(app.db).GetCity(cityID)

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

	gj, err := city.NewService(app.db).SearchCities(name)

	if err != nil {
		app.errorLog.Println("Erro ao buscar municípios com o nome: ", name)
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(gj)
}

func (app *application) geocoder(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")

	app.infoLog.Println("Init geocoder. Address:", address)
	geocoder, err := geocoder.NewService(app.mapsURL, app.mapsKey).Geocoder(address)
	if err != nil {
		app.errorLog.Println("Erro ao geocodificar endereço: ", address)
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(geocoder)
}

func (app *application) saveArea(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var a entities.Area
	if err := decoder.Decode(&a); err != nil {
		http.Error(w, http.StatusText(500)+": check json format and types", http.StatusInternalServerError)
		app.errorLog.Println("Erro ao decodificar area: ", err.Error())
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	success, err := area.NewService(app.db).SaveArea(a)
	if err != nil || !success {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		app.errorLog.Println("Erro ao salvar area: ", err.Error())
		json.NewEncoder(w).Encode(&Err{
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
