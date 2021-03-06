// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetAppBitsUploadFeatureFlagResponse get app bits upload feature flag response
//
// swagger:model getAppBitsUploadFeatureFlagResponse
type GetAppBitsUploadFeatureFlagResponse struct {

	// The enabled
	Enabled bool `json:"enabled,omitempty"`

	// The error Message
	ErrorMessage GenericObject `json:"error_message,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// The url
	URL string `json:"url,omitempty"`
}

// Validate validates this get app bits upload feature flag response
func (m *GetAppBitsUploadFeatureFlagResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAppBitsUploadFeatureFlagResponse) validateErrorMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorMessage) { // not required
		return nil
	}

	if err := m.ErrorMessage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("error_message")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAppBitsUploadFeatureFlagResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAppBitsUploadFeatureFlagResponse) UnmarshalBinary(b []byte) error {
	var res GetAppBitsUploadFeatureFlagResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
