// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/henkvanramshorst/forest/backend-api/gen/dto"
)

// GetLocationBasedDataOKCode is the HTTP code returned for type GetLocationBasedDataOK
const GetLocationBasedDataOKCode int = 200

/*GetLocationBasedDataOK all the data!

swagger:response getLocationBasedDataOK
*/
type GetLocationBasedDataOK struct {

	/*
	  In: Body
	*/
	Payload *dto.LocationData `json:"body,omitempty"`
}

// NewGetLocationBasedDataOK creates GetLocationBasedDataOK with default headers values
func NewGetLocationBasedDataOK() *GetLocationBasedDataOK {
	return &GetLocationBasedDataOK{}
}

// WithPayload adds the payload to the get location based data o k response
func (o *GetLocationBasedDataOK) WithPayload(payload *dto.LocationData) *GetLocationBasedDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get location based data o k response
func (o *GetLocationBasedDataOK) SetPayload(payload *dto.LocationData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLocationBasedDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
