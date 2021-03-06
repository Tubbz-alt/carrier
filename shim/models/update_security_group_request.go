// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateSecurityGroupRequest update security group request
//
// swagger:model updateSecurityGroupRequest
type UpdateSecurityGroupRequest struct {

	// The name of the security group.
	Name string `json:"name,omitempty"`

	// The egress rules for apps that belong to this security group. A rule consists of a protocol (tcp icmp udp all) destination CIDR or destination range  port or port range (tcp udp all) type (control signal for icmp) code (control signal for icmp)  log (enables logging for the egress rule)
	Rules []GenericObject `json:"rules"`

	// The list of associated spaces.
	SpaceGuids GenericObject `json:"space_guids,omitempty"`
}

// Validate validates this update security group request
func (m *UpdateSecurityGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpaceGuids(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSecurityGroupRequest) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {

		if err := m.Rules[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rules" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UpdateSecurityGroupRequest) validateSpaceGuids(formats strfmt.Registry) error {

	if swag.IsZero(m.SpaceGuids) { // not required
		return nil
	}

	if err := m.SpaceGuids.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("space_guids")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSecurityGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSecurityGroupRequest) UnmarshalBinary(b []byte) error {
	var res UpdateSecurityGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
