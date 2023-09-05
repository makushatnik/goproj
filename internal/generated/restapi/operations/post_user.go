// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUserHandlerFunc turns a function with the right signature into a post user handler
type PostUserHandlerFunc func(PostUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUserHandlerFunc) Handle(params PostUserParams) middleware.Responder {
	return fn(params)
}

// PostUserHandler interface for that can handle valid post user params
type PostUserHandler interface {
	Handle(PostUserParams) middleware.Responder
}

// NewPostUser creates a new http.Handler for the post user operation
func NewPostUser(ctx *middleware.Context, handler PostUserHandler) *PostUser {
	return &PostUser{Context: ctx, Handler: handler}
}

/*
	PostUser swagger:route POST /user postUser

Create User
*/
type PostUser struct {
	Context *middleware.Context
	Handler PostUserHandler
}

func (o *PostUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
