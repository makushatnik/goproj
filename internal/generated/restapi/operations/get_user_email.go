// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserEmailHandlerFunc turns a function with the right signature into a get user email handler
type GetUserEmailHandlerFunc func(GetUserEmailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserEmailHandlerFunc) Handle(params GetUserEmailParams) middleware.Responder {
	return fn(params)
}

// GetUserEmailHandler interface for that can handle valid get user email params
type GetUserEmailHandler interface {
	Handle(GetUserEmailParams) middleware.Responder
}

// NewGetUserEmail creates a new http.Handler for the get user email operation
func NewGetUserEmail(ctx *middleware.Context, handler GetUserEmailHandler) *GetUserEmail {
	return &GetUserEmail{Context: ctx, Handler: handler}
}

/*
	GetUserEmail swagger:route GET /user/{email} getUserEmail

Get User by email
*/
type GetUserEmail struct {
	Context *middleware.Context
	Handler GetUserEmailHandler
}

func (o *GetUserEmail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserEmailParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
