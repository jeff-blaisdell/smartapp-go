// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AppOAuth app o auth
// swagger:model AppOAuth
type AppOAuth struct {

	// A name given to the OAuth Client.
	ClientName string `json:"clientName,omitempty"`

	// A list of SmartThings API OAuth scope identifiers that maybe required to execute your integration.
	Scope []string `json:"scope"`
}

// Validate validates this app o auth
func (m *AppOAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppOAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppOAuth) UnmarshalBinary(b []byte) error {
	var res AppOAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}