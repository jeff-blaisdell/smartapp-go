// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateOrUpdateAppRequest create or update app request
// swagger:model CreateOrUpdateAppRequest
type CreateOrUpdateAppRequest struct {

	// A globally unique, developer-defined identifier for an app. It is alpha-numeric, may contain dashes,
	// underscores, periods, and must be less then 250 characters long.
	//
	// Required: true
	AppName *string `json:"appName"`

	// app type
	// Required: true
	AppType AppType `json:"appType"`

	// A default description for an app.
	//
	// Required: true
	// Max Length: 250
	Description *string `json:"description"`

	// A default display name for an app.
	//
	// Required: true
	// Max Length: 75
	DisplayName *string `json:"displayName"`

	// lambda smart app
	LambdaSmartApp *CreateOrUpdateLambdaSmartAppRequest `json:"lambdaSmartApp,omitempty"`

	// Inform the installation systems that a particular app can only be installed once within a user's account.
	//
	SingleInstance *bool `json:"singleInstance,omitempty"`

	// webhook smart app
	WebhookSmartApp *CreateOrUpdateWebhookSmartAppRequest `json:"webhookSmartApp,omitempty"`
}

// Validate validates this create or update app request
func (m *CreateOrUpdateAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLambdaSmartApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookSmartApp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrUpdateAppRequest) validateAppName(formats strfmt.Registry) error {

	if err := validate.Required("appName", "body", m.AppName); err != nil {
		return err
	}

	return nil
}

func (m *CreateOrUpdateAppRequest) validateAppType(formats strfmt.Registry) error {

	if err := m.AppType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("appType")
		}
		return err
	}

	return nil
}

func (m *CreateOrUpdateAppRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 250); err != nil {
		return err
	}

	return nil
}

func (m *CreateOrUpdateAppRequest) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(*m.DisplayName), 75); err != nil {
		return err
	}

	return nil
}

func (m *CreateOrUpdateAppRequest) validateLambdaSmartApp(formats strfmt.Registry) error {

	if swag.IsZero(m.LambdaSmartApp) { // not required
		return nil
	}

	if m.LambdaSmartApp != nil {
		if err := m.LambdaSmartApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lambdaSmartApp")
			}
			return err
		}
	}

	return nil
}

func (m *CreateOrUpdateAppRequest) validateWebhookSmartApp(formats strfmt.Registry) error {

	if swag.IsZero(m.WebhookSmartApp) { // not required
		return nil
	}

	if m.WebhookSmartApp != nil {
		if err := m.WebhookSmartApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhookSmartApp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrUpdateAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrUpdateAppRequest) UnmarshalBinary(b []byte) error {
	var res CreateOrUpdateAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
