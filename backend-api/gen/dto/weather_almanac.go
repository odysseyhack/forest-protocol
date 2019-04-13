// Code generated by go-swagger; DO NOT EDIT.

package dto

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WeatherAlmanac weather almanac
// swagger:model WeatherAlmanac
type WeatherAlmanac struct {

	// lat
	Lat string `json:"lat,omitempty"`

	// long
	Long string `json:"long,omitempty"`

	// months
	Months WeatherAlmanacMonths `json:"months"`
}

// Validate validates this weather almanac
func (m *WeatherAlmanac) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WeatherAlmanac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WeatherAlmanac) UnmarshalBinary(b []byte) error {
	var res WeatherAlmanac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
