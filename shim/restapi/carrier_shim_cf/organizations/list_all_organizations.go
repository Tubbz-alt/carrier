// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllOrganizationsHandlerFunc turns a function with the right signature into a list all organizations handler
type ListAllOrganizationsHandlerFunc func(ListAllOrganizationsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllOrganizationsHandlerFunc) Handle(params ListAllOrganizationsParams) middleware.Responder {
	return fn(params)
}

// ListAllOrganizationsHandler interface for that can handle valid list all organizations params
type ListAllOrganizationsHandler interface {
	Handle(ListAllOrganizationsParams) middleware.Responder
}

// NewListAllOrganizations creates a new http.Handler for the list all organizations operation
func NewListAllOrganizations(ctx *middleware.Context, handler ListAllOrganizationsHandler) *ListAllOrganizations {
	return &ListAllOrganizations{Context: ctx, Handler: handler}
}

/*ListAllOrganizations swagger:route GET /organizations organizations listAllOrganizations

List all Organizations

curl --insecure -i %s/v2/organizations -X GET -H 'Authorization: %s'

*/
type ListAllOrganizations struct {
	Context *middleware.Context
	Handler ListAllOrganizationsHandler
}

func (o *ListAllOrganizations) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllOrganizationsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
