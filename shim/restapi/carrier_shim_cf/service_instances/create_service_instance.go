// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateServiceInstanceHandlerFunc turns a function with the right signature into a create service instance handler
type CreateServiceInstanceHandlerFunc func(CreateServiceInstanceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateServiceInstanceHandlerFunc) Handle(params CreateServiceInstanceParams) middleware.Responder {
	return fn(params)
}

// CreateServiceInstanceHandler interface for that can handle valid create service instance params
type CreateServiceInstanceHandler interface {
	Handle(CreateServiceInstanceParams) middleware.Responder
}

// NewCreateServiceInstance creates a new http.Handler for the create service instance operation
func NewCreateServiceInstance(ctx *middleware.Context, handler CreateServiceInstanceHandler) *CreateServiceInstance {
	return &CreateServiceInstance{Context: ctx, Handler: handler}
}

/*CreateServiceInstance swagger:route POST /service_instances serviceInstances createServiceInstance

Creating a Service Instance

curl --insecure -i %s/v2/service_instances -X POST -H 'Authorization: %s' -d '%s'

*/
type CreateServiceInstance struct {
	Context *middleware.Context
	Handler CreateServiceInstanceHandler
}

func (o *CreateServiceInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateServiceInstanceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}