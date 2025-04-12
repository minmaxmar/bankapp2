// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bankapp2/app/models"
)

// GetCardsIDOKCode is the HTTP code returned for type GetCardsIDOK
const GetCardsIDOKCode int = 200

/*
GetCardsIDOK Ok

swagger:response getCardsIdOK
*/
type GetCardsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Card `json:"body,omitempty"`
}

// NewGetCardsIDOK creates GetCardsIDOK with default headers values
func NewGetCardsIDOK() *GetCardsIDOK {

	return &GetCardsIDOK{}
}

// WithPayload adds the payload to the get cards Id o k response
func (o *GetCardsIDOK) WithPayload(payload *models.Card) *GetCardsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cards Id o k response
func (o *GetCardsIDOK) SetPayload(payload *models.Card) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCardsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCardsIDNotFoundCode is the HTTP code returned for type GetCardsIDNotFound
const GetCardsIDNotFoundCode int = 404

/*
GetCardsIDNotFound Card not found

swagger:response getCardsIdNotFound
*/
type GetCardsIDNotFound struct {
}

// NewGetCardsIDNotFound creates GetCardsIDNotFound with default headers values
func NewGetCardsIDNotFound() *GetCardsIDNotFound {

	return &GetCardsIDNotFound{}
}

// WriteResponse to the client
func (o *GetCardsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*
GetCardsIDDefault Общая ошибка

swagger:response getCardsIdDefault
*/
type GetCardsIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetCardsIDDefault creates GetCardsIDDefault with default headers values
func NewGetCardsIDDefault(code int) *GetCardsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCardsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get cards ID default response
func (o *GetCardsIDDefault) WithStatusCode(code int) *GetCardsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get cards ID default response
func (o *GetCardsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get cards ID default response
func (o *GetCardsIDDefault) WithPayload(payload *models.ErrorResponse) *GetCardsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cards ID default response
func (o *GetCardsIDDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCardsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
