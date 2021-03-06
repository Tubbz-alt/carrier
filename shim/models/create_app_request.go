// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateAppRequest create app request
//
// swagger:model createAppRequest
type CreateAppRequest struct {

	// Buildpack to build the app. 3 options: a) Blank means autodetection; b) A Git Url pointing to a buildpack; c) Name of an installed buildpack.
	Buildpack GenericObject `json:"buildpack,omitempty"`

	// The command to start an app after it is staged (e.g. 'rails s -p $PORT' or 'java com.org.Server $PORT').
	Command GenericObject `json:"command,omitempty"`

	// Open the console port for the app (at $CONSOLE_PORT).
	Console bool `json:"console,omitempty"`

	// Open the debug port for the app (at $DEBUG_PORT).
	Debug bool `json:"debug,omitempty"`

	// The command detected by the buildpack during staging.
	DetectedStartCommand GenericObject `json:"detected_start_command,omitempty"`

	// The maximum amount of disk available to an instance of an app. In megabytes.
	DiskQuota string `json:"disk_quota,omitempty"`

	// Name of the Docker image containing the app
	DockerImage string `json:"docker_image,omitempty"`

	// Key/value pairs of all the environment variables to run in your app. Does not include any system or service variables.
	EnvironmentJSON GenericObject `json:"environment_json,omitempty"`

	// Timeout for health checking of an staged app when starting up
	HealthCheckTimeout GenericObject `json:"health_check_timeout,omitempty"`

	// Type of health check to perform.
	HealthCheckType string `json:"health_check_type,omitempty"`

	// The number of instances of the app to run. To ensure optimal availability ensure there are at least 2 instances.
	Instances string `json:"instances,omitempty"`

	// The amount of memory each instance should have. In megabytes.
	Memory string `json:"memory,omitempty"`

	// The name of the app.
	Name string `json:"name,omitempty"`

	// Deprecated.
	Production bool `json:"production,omitempty"`

	// The guid of the associated space.
	SpaceGUID string `json:"space_guid,omitempty"`

	// The guid of the associated stack.
	StackGUID string `json:"stack_guid,omitempty"`

	// The current desired state of the app. One of STOPPED or STARTED.
	State string `json:"state,omitempty"`
}

// Validate validates this create app request
func (m *CreateAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildpack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetectedStartCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthCheckTimeout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAppRequest) validateBuildpack(formats strfmt.Registry) error {

	if swag.IsZero(m.Buildpack) { // not required
		return nil
	}

	if err := m.Buildpack.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("buildpack")
		}
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.Command) { // not required
		return nil
	}

	if err := m.Command.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("command")
		}
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateDetectedStartCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.DetectedStartCommand) { // not required
		return nil
	}

	if err := m.DetectedStartCommand.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("detected_start_command")
		}
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateEnvironmentJSON(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvironmentJSON) { // not required
		return nil
	}

	if err := m.EnvironmentJSON.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("environment_json")
		}
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateHealthCheckTimeout(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthCheckTimeout) { // not required
		return nil
	}

	if err := m.HealthCheckTimeout.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("health_check_timeout")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAppRequest) UnmarshalBinary(b []byte) error {
	var res CreateAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
