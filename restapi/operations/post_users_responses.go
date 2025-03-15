// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bankapp2/app/models"
)

// PostUsersCreatedCode is the HTTP code returned for type PostUsersCreated
const PostUsersCreatedCode int = 201

/*
PostUsersCreated User created

swagger:response postUsersCreated
*/
type PostUsersCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUsersCreated creates PostUsersCreated with default headers values
func NewPostUsersCreated() *PostUsersCreated {

	return &PostUsersCreated{}
}

// WithPayload adds the payload to the post users created response
func (o *PostUsersCreated) WithPayload(payload *models.User) *PostUsersCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users created response
func (o *PostUsersCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostUsersDefault Общая ошибка

swagger:response postUsersDefault
*/
type PostUsersDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPostUsersDefault creates PostUsersDefault with default headers values
func NewPostUsersDefault(code int) *PostUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post users default response
func (o *PostUsersDefault) WithStatusCode(code int) *PostUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post users default response
func (o *PostUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post users default response
func (o *PostUsersDefault) WithPayload(payload *models.ErrorResponse) *PostUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users default response
func (o *PostUsersDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
