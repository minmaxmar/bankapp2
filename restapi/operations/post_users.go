// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostUsersHandlerFunc turns a function with the right signature into a post users handler
type PostUsersHandlerFunc func(PostUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostUsersHandlerFunc) Handle(params PostUsersParams) middleware.Responder {
	return fn(params)
}

// PostUsersHandler interface for that can handle valid post users params
type PostUsersHandler interface {
	Handle(PostUsersParams) middleware.Responder
}

// NewPostUsers creates a new http.Handler for the post users operation
func NewPostUsers(ctx *middleware.Context, handler PostUsersHandler) *PostUsers {
	return &PostUsers{Context: ctx, Handler: handler}
}

/*
	PostUsers swagger:route POST /users postUsers

Create a new user
*/
type PostUsers struct {
	Context *middleware.Context
	Handler PostUsersHandler
}

func (o *PostUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostUsersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
