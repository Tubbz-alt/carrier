// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateServiceBrokerResponse create service broker response
//
// swagger:model createServiceBrokerResponse
type CreateServiceBrokerResponse struct {

	// The auth Username
	AuthUsername string `json:"auth_username,omitempty"`

	// The broker Url
	BrokerURL string `json:"broker_url,omitempty"`

	// The name
	Name string `json:"name,omitempty"`
}

// Validate validates this create service broker response
func (m *CreateServiceBrokerResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateServiceBrokerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateServiceBrokerResponse) UnmarshalBinary(b []byte) error {
	var res CreateServiceBrokerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
