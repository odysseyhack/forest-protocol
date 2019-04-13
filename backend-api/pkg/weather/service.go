package weather

import (
	"gitlab.com/henkvanramshorst/forest/backend-api/gen/dto"
	"github.com/unchainio/interfaces/logger"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/backend-api"
	"net/http"
	"github.com/go-chi/render"
	"fmt"
)

type WeatherService struct {
	logger logger.Logger
	cfg *backend_api.Config
}

func New(logger logger.Logger, cfg *backend_api.Config) (*WeatherService) {
	return &WeatherService{logger: logger, cfg: cfg}
}

func (s *WeatherService) GetAlmanacForGeoLocation(lat, long string) (*dto.WeatherAlmanac, error) {
	client := http.Client{}
	res, err := client.Get(s.cfg.WeatherBaseUrl + "/api/weather/v1/geocode/" + lat + "/" + long + "/almanac/monthly.json?start=01&end=12")
	if err != nil {
		fmt.Printf("error here: %s\n", err.Error())
		return nil, err
	}

	var parsedResponse WeatherServiceAlmanacResponse
	err = render.DecodeJSON(res.Body, &parsedResponse)
	if err != nil {
		fmt.Printf("error here: %s\n", err.Error())
		return nil, err
	}

	var monthData []*dto.MonthlyWeather
	for _, month := range parsedResponse.MonthData {
		m := &dto.MonthlyWeather{
			AvgHigh: fmt.Sprintf("%f", month.AvgHigh),
			AvgLow: fmt.Sprintf("%f", month.AvgLow),
			MeanTemp: fmt.Sprintf("%f", month.MeanTemp),
			Month: month.Month,
			Precip: fmt.Sprintf("%f", month.AvgPrecip),
		}
		monthData = append(monthData, m)
	}

	almanac := &dto.WeatherAlmanac{
		Long: fmt.Sprintf("%f", parsedResponse.MetaData.Longitude),
		Lat: fmt.Sprintf("%f", parsedResponse.MetaData.Latitude),
		Months: monthData,
	}

	return almanac, nil
}

