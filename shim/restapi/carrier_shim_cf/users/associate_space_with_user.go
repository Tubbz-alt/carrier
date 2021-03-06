// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AssociateSpaceWithUserHandlerFunc turns a function with the right signature into a associate space with user handler
type AssociateSpaceWithUserHandlerFunc func(AssociateSpaceWithUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssociateSpaceWithUserHandlerFunc) Handle(params AssociateSpaceWithUserParams) middleware.Responder {
	return fn(params)
}

// AssociateSpaceWithUserHandler interface for that can handle valid associate space with user params
type AssociateSpaceWithUserHandler interface {
	Handle(AssociateSpaceWithUserParams) middleware.Responder
}

// NewAssociateSpaceWithUser creates a new http.Handler for the associate space with user operation
func NewAssociateSpaceWithUser(ctx *middleware.Context, handler AssociateSpaceWithUserHandler) *AssociateSpaceWithUser {
	return &AssociateSpaceWithUser{Context: ctx, Handler: handler}
}

/*AssociateSpaceWithUser swagger:route PUT /users/{guid}/spaces/{space_guid} users associateSpaceWithUser

Associate Space with the User

curl --insecure -i %s/v2/users/{guid}/spaces/{space_guid} -X PUT -H 'Authorization: %s'

*/
type AssociateSpaceWithUser struct {
	Context *middleware.Context
	Handler AssociateSpaceWithUserHandler
}

func (o *AssociateSpaceWithUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAssociateSpaceWithUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
