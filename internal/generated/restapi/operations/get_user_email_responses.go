// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"goproj/internal/generated/models"
)

// GetUserEmailOKCode is the HTTP code returned for type GetUserEmailOK
const GetUserEmailOKCode int = 200

/*
GetUserEmailOK Ok

swagger:response getUserEmailOK
*/
type GetUserEmailOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserEmailOK creates GetUserEmailOK with default headers values
func NewGetUserEmailOK() *GetUserEmailOK {

	return &GetUserEmailOK{}
}

// WithPayload adds the payload to the get user email o k response
func (o *GetUserEmailOK) WithPayload(payload *models.User) *GetUserEmailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user email o k response
func (o *GetUserEmailOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserEmailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserEmailBadRequestCode is the HTTP code returned for type GetUserEmailBadRequest
const GetUserEmailBadRequestCode int = 400

/*
GetUserEmailBadRequest Bad request

swagger:response getUserEmailBadRequest
*/
type GetUserEmailBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.FailResponse `json:"body,omitempty"`
}

// NewGetUserEmailBadRequest creates GetUserEmailBadRequest with default headers values
func NewGetUserEmailBadRequest() *GetUserEmailBadRequest {

	return &GetUserEmailBadRequest{}
}

// WithPayload adds the payload to the get user email bad request response
func (o *GetUserEmailBadRequest) WithPayload(payload *models.FailResponse) *GetUserEmailBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user email bad request response
func (o *GetUserEmailBadRequest) SetPayload(payload *models.FailResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserEmailBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserEmailInternalServerErrorCode is the HTTP code returned for type GetUserEmailInternalServerError
const GetUserEmailInternalServerErrorCode int = 500

/*
GetUserEmailInternalServerError Internal server error

swagger:response getUserEmailInternalServerError
*/
type GetUserEmailInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.FailResponse `json:"body,omitempty"`
}

// NewGetUserEmailInternalServerError creates GetUserEmailInternalServerError with default headers values
func NewGetUserEmailInternalServerError() *GetUserEmailInternalServerError {

	return &GetUserEmailInternalServerError{}
}

// WithPayload adds the payload to the get user email internal server error response
func (o *GetUserEmailInternalServerError) WithPayload(payload *models.FailResponse) *GetUserEmailInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user email internal server error response
func (o *GetUserEmailInternalServerError) SetPayload(payload *models.FailResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserEmailInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
