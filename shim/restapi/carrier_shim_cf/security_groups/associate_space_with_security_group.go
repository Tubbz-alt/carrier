// Code generated by go-swagger; DO NOT EDIT.

package security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AssociateSpaceWithSecurityGroupHandlerFunc turns a function with the right signature into a associate space with security group handler
type AssociateSpaceWithSecurityGroupHandlerFunc func(AssociateSpaceWithSecurityGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssociateSpaceWithSecurityGroupHandlerFunc) Handle(params AssociateSpaceWithSecurityGroupParams) middleware.Responder {
	return fn(params)
}

// AssociateSpaceWithSecurityGroupHandler interface for that can handle valid associate space with security group params
type AssociateSpaceWithSecurityGroupHandler interface {
	Handle(AssociateSpaceWithSecurityGroupParams) middleware.Responder
}

// NewAssociateSpaceWithSecurityGroup creates a new http.Handler for the associate space with security group operation
func NewAssociateSpaceWithSecurityGroup(ctx *middleware.Context, handler AssociateSpaceWithSecurityGroupHandler) *AssociateSpaceWithSecurityGroup {
	return &AssociateSpaceWithSecurityGroup{Context: ctx, Handler: handler}
}

/*AssociateSpaceWithSecurityGroup swagger:route PUT /security_groups/{guid}/spaces/{space_guid} securityGroups associateSpaceWithSecurityGroup

Associate Space with the Security Group

curl --insecure -i %s/v2/security_groups/{guid}/spaces/{space_guid} -X PUT -H 'Authorization: %s'

*/
type AssociateSpaceWithSecurityGroup struct {
	Context *middleware.Context
	Handler AssociateSpaceWithSecurityGroupHandler
}

func (o *AssociateSpaceWithSecurityGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAssociateSpaceWithSecurityGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}