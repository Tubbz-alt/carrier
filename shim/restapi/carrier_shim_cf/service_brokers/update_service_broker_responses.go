// Code generated by go-swagger; DO NOT EDIT.

package service_brokers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// UpdateServiceBrokerOKCode is the HTTP code returned for type UpdateServiceBrokerOK
const UpdateServiceBrokerOKCode int = 200

/*UpdateServiceBrokerOK successful response

swagger:response updateServiceBrokerOK
*/
type UpdateServiceBrokerOK struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateServiceBrokerResponse `json:"body,omitempty"`
}

// NewUpdateServiceBrokerOK creates UpdateServiceBrokerOK with default headers values
func NewUpdateServiceBrokerOK() *UpdateServiceBrokerOK {

	return &UpdateServiceBrokerOK{}
}

// WithPayload adds the payload to the update service broker o k response
func (o *UpdateServiceBrokerOK) WithPayload(payload *models.UpdateServiceBrokerResponse) *UpdateServiceBrokerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service broker o k response
func (o *UpdateServiceBrokerOK) SetPayload(payload *models.UpdateServiceBrokerResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceBrokerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
