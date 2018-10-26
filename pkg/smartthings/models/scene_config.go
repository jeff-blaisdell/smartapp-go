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

// SceneConfig A scene configuration.
// swagger:model SceneConfig
type SceneConfig struct {

	// permissions
	// Max Items: 25
	Permissions []string `json:"permissions"`

	// The ID of the scene.
	SceneID string `json:"sceneId,omitempty"`
}

// Validate validates this scene config
func (m *SceneConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SceneConfig) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	iPermissionsSize := int64(len(m.Permissions))

	if err := validate.MaxItems("permissions", "body", iPermissionsSize, 25); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SceneConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SceneConfig) UnmarshalBinary(b []byte) error {
	var res SceneConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}