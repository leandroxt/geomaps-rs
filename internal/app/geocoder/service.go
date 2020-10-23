package geocoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/leandroxt/geomaps-rs/internal/entities"
)

// Service create a geocoder service interface
type Service interface {
	Geocoder(address string) (entities.GeocoderResponse, error)
}

// ServiceImpl create a geocoder service implementation
type ServiceImpl struct {
	MapsURL       string
	GoogleMapsKey string
}

// NewService return the service implementation
func NewService(mapsURL, googleMapsKey string) Service {
	return ServiceImpl{
		MapsURL:       mapsURL,
		GoogleMapsKey: googleMapsKey,
	}
}

// Geocoder receive an address and return a list of possible coordinates
func (s ServiceImpl) Geocoder(address string) (entities.GeocoderResponse, error) {
	path := "/geocode/json"
	rawURL := fmt.Sprintf("%s%s?address=%s&key=%s",
		s.MapsURL,
		path,
		address,
		s.GoogleMapsKey)
	URL := strings.ReplaceAll(rawURL, " ", "%20")

	geoResp := &entities.GeocoderResponse{}
	response, err := http.Get(URL)
	if err != nil {
		log.Println("Erro na chamada do geocoder:", err.Error())
		return *geoResp, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Erro na leitura do corpo da response: ", err.Error())
		return *geoResp, err
	}

	r := bytes.NewReader(data)
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&geoResp)
	if err != nil {
		log.Println("Não foi possível decodificar JSON: ", err.Error())
		return *geoResp, err
	}

	return *geoResp, nil
}
