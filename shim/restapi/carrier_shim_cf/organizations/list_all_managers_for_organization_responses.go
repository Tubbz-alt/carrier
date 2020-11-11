// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllManagersForOrganizationOKCode is the HTTP code returned for type ListAllManagersForOrganizationOK
const ListAllManagersForOrganizationOKCode int = 200

/*ListAllManagersForOrganizationOK successful response

swagger:response listAllManagersForOrganizationOK
*/
type ListAllManagersForOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllManagersForOrganizationResponsePaged `json:"body,omitempty"`
}

// NewListAllManagersForOrganizationOK creates ListAllManagersForOrganizationOK with default headers values
func NewListAllManagersForOrganizationOK() *ListAllManagersForOrganizationOK {

	return &ListAllManagersForOrganizationOK{}
}

// WithPayload adds the payload to the list all managers for organization o k response
func (o *ListAllManagersForOrganizationOK) WithPayload(payload *models.ListAllManagersForOrganizationResponsePaged) *ListAllManagersForOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all managers for organization o k response
func (o *ListAllManagersForOrganizationOK) SetPayload(payload *models.ListAllManagersForOrganizationResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllManagersForOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}