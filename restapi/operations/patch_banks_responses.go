// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bankapp2/app/models"
)

// PatchBanksCreatedCode is the HTTP code returned for type PatchBanksCreated
const PatchBanksCreatedCode int = 201

/*
PatchBanksCreated Bank patched

swagger:response patchBanksCreated
*/
type PatchBanksCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Bank `json:"body,omitempty"`
}

// NewPatchBanksCreated creates PatchBanksCreated with default headers values
func NewPatchBanksCreated() *PatchBanksCreated {

	return &PatchBanksCreated{}
}

// WithPayload adds the payload to the patch banks created response
func (o *PatchBanksCreated) WithPayload(payload *models.Bank) *PatchBanksCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch banks created response
func (o *PatchBanksCreated) SetPayload(payload *models.Bank) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchBanksCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PatchBanksDefault Общая ошибка

swagger:response patchBanksDefault
*/
type PatchBanksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPatchBanksDefault creates PatchBanksDefault with default headers values
func NewPatchBanksDefault(code int) *PatchBanksDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchBanksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch banks default response
func (o *PatchBanksDefault) WithStatusCode(code int) *PatchBanksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch banks default response
func (o *PatchBanksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch banks default response
func (o *PatchBanksDefault) WithPayload(payload *models.ErrorResponse) *PatchBanksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch banks default response
func (o *PatchBanksDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchBanksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
