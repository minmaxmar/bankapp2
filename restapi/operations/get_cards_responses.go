// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bankapp2/app/models"
)

// GetCardsOKCode is the HTTP code returned for type GetCardsOK
const GetCardsOKCode int = 200

/*
GetCardsOK Ok

swagger:response getCardsOK
*/
type GetCardsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Card `json:"body,omitempty"`
}

// NewGetCardsOK creates GetCardsOK with default headers values
func NewGetCardsOK() *GetCardsOK {

	return &GetCardsOK{}
}

// WithPayload adds the payload to the get cards o k response
func (o *GetCardsOK) WithPayload(payload []*models.Card) *GetCardsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cards o k response
func (o *GetCardsOK) SetPayload(payload []*models.Card) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCardsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Card, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetCardsDefault Общая ошибка

swagger:response getCardsDefault
*/
type GetCardsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetCardsDefault creates GetCardsDefault with default headers values
func NewGetCardsDefault(code int) *GetCardsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCardsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get cards default response
func (o *GetCardsDefault) WithStatusCode(code int) *GetCardsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get cards default response
func (o *GetCardsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get cards default response
func (o *GetCardsDefault) WithPayload(payload *models.ErrorResponse) *GetCardsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cards default response
func (o *GetCardsDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCardsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
