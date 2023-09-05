// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"goproj/internal/generated/models"
)

// PostUserCreatedCode is the HTTP code returned for type PostUserCreated
const PostUserCreatedCode int = 201

/*
PostUserCreated User was successfully created

swagger:response postUserCreated
*/
type PostUserCreated struct {
}

// NewPostUserCreated creates PostUserCreated with default headers values
func NewPostUserCreated() *PostUserCreated {

	return &PostUserCreated{}
}

// WriteResponse to the client
func (o *PostUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostUserBadRequestCode is the HTTP code returned for type PostUserBadRequest
const PostUserBadRequestCode int = 400

/*
PostUserBadRequest Bad request

swagger:response postUserBadRequest
*/
type PostUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.FailResponse `json:"body,omitempty"`
}

// NewPostUserBadRequest creates PostUserBadRequest with default headers values
func NewPostUserBadRequest() *PostUserBadRequest {

	return &PostUserBadRequest{}
}

// WithPayload adds the payload to the post user bad request response
func (o *PostUserBadRequest) WithPayload(payload *models.FailResponse) *PostUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user bad request response
func (o *PostUserBadRequest) SetPayload(payload *models.FailResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostUserInternalServerErrorCode is the HTTP code returned for type PostUserInternalServerError
const PostUserInternalServerErrorCode int = 500

/*
PostUserInternalServerError Internal server error

swagger:response postUserInternalServerError
*/
type PostUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.FailResponse `json:"body,omitempty"`
}

// NewPostUserInternalServerError creates PostUserInternalServerError with default headers values
func NewPostUserInternalServerError() *PostUserInternalServerError {

	return &PostUserInternalServerError{}
}

// WithPayload adds the payload to the post user internal server error response
func (o *PostUserInternalServerError) WithPayload(payload *models.FailResponse) *PostUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user internal server error response
func (o *PostUserInternalServerError) SetPayload(payload *models.FailResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
