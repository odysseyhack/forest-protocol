package handler

import (
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/backend-api"
	"github.com/unchainio/interfaces/logger"
	"net/http"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/weather"
	"gitlab.com/henkvanramshorst/forest/backend-api/gen/dto"
	"github.com/go-chi/render"
	"encoding/json"
	"math/rand"
	"time"
)

type LocationHandler struct {
	logger logger.Logger
	cfg *backend_api.Config
}

func NewLocationHandler(logger logger.Logger, cfg *backend_api.Config) *LocationHandler {
	return &LocationHandler{logger: logger, cfg: cfg}
}

func (h *LocationHandler) GetLocationData(w http.ResponseWriter, r *http.Request) {
	var location dto.LatLong

	err := render.DecodeJSON(r.Body, &location)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error unmarshalling"))
		return
	}

	weatherService := weather.New(h.logger, h.cfg)

	almanac, err := weatherService.GetAlmanacForGeoLocation(location.Lat, location.Long)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("could not get almanac data"))
		return
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	silt := random.Int63n(100)
	sand := random.Int63n(100)
	clay := random.Int63n(100)

	response := &dto.LocationData{
		WeatherAlmanac: almanac,
		Soiltype: &dto.SoilType{
			Clay: clay,
			Sand: sand,
			Silt: silt,
		},
	}

	respBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("could not marshal final response"))
		return
	}

	w.Write(respBytes)
}

