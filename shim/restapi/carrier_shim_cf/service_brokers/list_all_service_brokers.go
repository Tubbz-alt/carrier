// Code generated by go-swagger; DO NOT EDIT.

package service_brokers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllServiceBrokersHandlerFunc turns a function with the right signature into a list all service brokers handler
type ListAllServiceBrokersHandlerFunc func(ListAllServiceBrokersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllServiceBrokersHandlerFunc) Handle(params ListAllServiceBrokersParams) middleware.Responder {
	return fn(params)
}

// ListAllServiceBrokersHandler interface for that can handle valid list all service brokers params
type ListAllServiceBrokersHandler interface {
	Handle(ListAllServiceBrokersParams) middleware.Responder
}

// NewListAllServiceBrokers creates a new http.Handler for the list all service brokers operation
func NewListAllServiceBrokers(ctx *middleware.Context, handler ListAllServiceBrokersHandler) *ListAllServiceBrokers {
	return &ListAllServiceBrokers{Context: ctx, Handler: handler}
}

/*ListAllServiceBrokers swagger:route GET /service_brokers serviceBrokers listAllServiceBrokers

List all Service Brokers

curl --insecure -i %s/v2/service_brokers -X GET -H 'Authorization: %s'

*/
type ListAllServiceBrokers struct {
	Context *middleware.Context
	Handler ListAllServiceBrokersHandler
}

func (o *ListAllServiceBrokers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllServiceBrokersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
