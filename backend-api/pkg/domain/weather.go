package domain

import "gitlab.com/henkvanramshorst/forest/backend-api/gen/dto"

type WeatherService interface {
	GetAlmanacForGeoLocation(lat, long string) (*dto.WeatherAlmanac, error)
}
