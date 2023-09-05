// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetUserEmailParams creates a new GetUserEmailParams object
//
// There are no default values defined in the spec.
func NewGetUserEmailParams() GetUserEmailParams {

	return GetUserEmailParams{}
}

// GetUserEmailParams contains all the bound params for the get user email operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetUserEmail
type GetUserEmailParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Email string
	/*
	  Required: true
	  In: header
	*/
	Fingerprint string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUserEmailParams() beforehand.
func (o *GetUserEmailParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEmail, rhkEmail, _ := route.Params.GetOK("email")
	if err := o.bindEmail(rEmail, rhkEmail, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindFingerprint(r.Header[http.CanonicalHeaderKey("fingerprint")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEmail binds and validates parameter Email from path.
func (o *GetUserEmailParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Email = raw

	return nil
}

// bindFingerprint binds and validates parameter Fingerprint from header.
func (o *GetUserEmailParams) bindFingerprint(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("fingerprint", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("fingerprint", "header", raw); err != nil {
		return err
	}
	o.Fingerprint = raw

	return nil
}