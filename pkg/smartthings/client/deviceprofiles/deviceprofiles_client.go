// Code generated by go-swagger; DO NOT EDIT.

package deviceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new deviceprofiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deviceprofiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDeviceProfile creates a device profile

Create a device profile.
*/
func (a *Client) CreateDeviceProfile(params *CreateDeviceProfileParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeviceProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDeviceProfile",
		Method:             "POST",
		PathPattern:        "/deviceprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeviceProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDeviceProfileOK), nil

}

/*
DeleteDeviceProfile deletes a device profile

Delete a device profile by ID. Admin use only
*/
func (a *Client) DeleteDeviceProfile(params *DeleteDeviceProfileParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeviceProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeviceProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDeviceProfile",
		Method:             "DELETE",
		PathPattern:        "/deviceprofiles/{deviceProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeviceProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDeviceProfileOK), nil

}

/*
GetDeviceProfile gs e t a device profile s description

GET a device profile's description.
*/
func (a *Client) GetDeviceProfile(params *GetDeviceProfileParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceProfile",
		Method:             "GET",
		PathPattern:        "/deviceprofiles/{deviceProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceProfileOK), nil

}

/*
ListDeviceProfiles lists all device profiles for the authenticated user

List device profiles.
*/
func (a *Client) ListDeviceProfiles(params *ListDeviceProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*ListDeviceProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDeviceProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDeviceProfiles",
		Method:             "GET",
		PathPattern:        "/deviceprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDeviceProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDeviceProfilesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
