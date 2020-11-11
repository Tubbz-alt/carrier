// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateSpaceQuotaDefinitionRequest create space quota definition request
//
// swagger:model createSpaceQuotaDefinitionRequest
type CreateSpaceQuotaDefinitionRequest struct {

	// The maximum amount of memory in megabytes an application instance can have. (-1 represents an unlimited amount)
	InstanceMemoryLimit string `json:"instance_memory_limit,omitempty"`

	// How much memory in megabytes a space can have.
	MemoryLimit int64 `json:"memory_limit,omitempty"`

	// The name for the Space Quota Definition.
	Name string `json:"name,omitempty"`

	// If a space can have non basic services
	NonBasicServicesAllowed bool `json:"non_basic_services_allowed,omitempty"`

	// The owning organization of the space quota
	OrganizationGUID string `json:"organization_guid,omitempty"`

	// How many routes a space can have.
	TotalRoutes int64 `json:"total_routes,omitempty"`

	// How many services a space can have.
	TotalServices int64 `json:"total_services,omitempty"`
}

// Validate validates this create space quota definition request
func (m *CreateSpaceQuotaDefinitionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSpaceQuotaDefinitionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSpaceQuotaDefinitionRequest) UnmarshalBinary(b []byte) error {
	var res CreateSpaceQuotaDefinitionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}