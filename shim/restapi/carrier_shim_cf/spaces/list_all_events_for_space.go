// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAllEventsForSpaceHandlerFunc turns a function with the right signature into a list all events for space handler
type ListAllEventsForSpaceHandlerFunc func(ListAllEventsForSpaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAllEventsForSpaceHandlerFunc) Handle(params ListAllEventsForSpaceParams) middleware.Responder {
	return fn(params)
}

// ListAllEventsForSpaceHandler interface for that can handle valid list all events for space params
type ListAllEventsForSpaceHandler interface {
	Handle(ListAllEventsForSpaceParams) middleware.Responder
}

// NewListAllEventsForSpace creates a new http.Handler for the list all events for space operation
func NewListAllEventsForSpace(ctx *middleware.Context, handler ListAllEventsForSpaceHandler) *ListAllEventsForSpace {
	return &ListAllEventsForSpace{Context: ctx, Handler: handler}
}

/*ListAllEventsForSpace swagger:route GET /spaces/{guid}/events spaces listAllEventsForSpace

List all Events for the Space

curl --insecure -i %s/v2/spaces/{guid}/events -X GET -H 'Authorization: %s'

*/
type ListAllEventsForSpace struct {
	Context *middleware.Context
	Handler ListAllEventsForSpaceHandler
}

func (o *ListAllEventsForSpace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAllEventsForSpaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}