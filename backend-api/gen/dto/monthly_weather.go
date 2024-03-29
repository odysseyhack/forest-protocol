// Code generated by go-swagger; DO NOT EDIT.

package dto

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MonthlyWeather monthly weather
// swagger:model MonthlyWeather
type MonthlyWeather struct {

	// avg high
	AvgHigh string `json:"avgHigh,omitempty"`

	// avg low
	AvgLow string `json:"avgLow,omitempty"`

	// mean temp
	MeanTemp string `json:"meanTemp,omitempty"`

	// month 01=jan, 02=feb etc
	Month string `json:"month,omitempty"`

	// precip
	Precip string `json:"precip,omitempty"`
}

// Validate validates this monthly weather
func (m *MonthlyWeather) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MonthlyWeather) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonthlyWeather) UnmarshalBinary(b []byte) error {
	var res MonthlyWeather
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
