// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InitializeSetting The initial setting to be returned when starting a new configuration install.
// swagger:model InitializeSetting
type InitializeSetting struct {
	Setting

	// Disable the ability for the user to customize the display name.
	DisableCustomDisplayName *bool `json:"disableCustomDisplayName,omitempty"`

	// Disable the ability to remove the app from the configuration flow.
	DisableRemoveApp *bool `json:"disableRemoveApp,omitempty"`

	// A developer defined page ID of the first page to display. Must be URL safe characters.
	FirstPageID string `json:"firstPageId,omitempty"`

	// permissions
	Permissions []string `json:"permissions"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InitializeSetting) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Setting
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Setting = aO0

	// AO1
	var dataAO1 struct {
		DisableCustomDisplayName *bool `json:"disableCustomDisplayName,omitempty"`

		DisableRemoveApp *bool `json:"disableRemoveApp,omitempty"`

		FirstPageID string `json:"firstPageId,omitempty"`

		Permissions []string `json:"permissions,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DisableCustomDisplayName = dataAO1.DisableCustomDisplayName

	m.DisableRemoveApp = dataAO1.DisableRemoveApp

	m.FirstPageID = dataAO1.FirstPageID

	m.Permissions = dataAO1.Permissions

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InitializeSetting) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Setting)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		DisableCustomDisplayName *bool `json:"disableCustomDisplayName,omitempty"`

		DisableRemoveApp *bool `json:"disableRemoveApp,omitempty"`

		FirstPageID string `json:"firstPageId,omitempty"`

		Permissions []string `json:"permissions,omitempty"`
	}

	dataAO1.DisableCustomDisplayName = m.DisableCustomDisplayName

	dataAO1.DisableRemoveApp = m.DisableRemoveApp

	dataAO1.FirstPageID = m.FirstPageID

	dataAO1.Permissions = m.Permissions

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this initialize setting
func (m *InitializeSetting) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Setting
	if err := m.Setting.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InitializeSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitializeSetting) UnmarshalBinary(b []byte) error {
	var res InitializeSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
