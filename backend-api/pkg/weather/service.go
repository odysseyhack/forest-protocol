package weather

import "gitlab.com/henkvanramshorst/forest/backend-api/gen/dto"

type WeatherService struct {

}

func New() (*WeatherService) {
	return nil
}

func (s *WeatherService) GetAlmanacForGeoLocation(lat, long string) (*dto.WeatherAlmanac, error) {
	return nil, nil
}
