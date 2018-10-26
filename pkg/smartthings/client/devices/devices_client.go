// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new devices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for devices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDeviceEvents creates device events

Create events for a device. When a device is managed by a SmartApp then it is responsible for creating events
to update the attributes of the device in the SmartThings platform.
The token must be for a SmartApp and it must be the SmartApp that created the Device.

*/
func (a *Client) CreateDeviceEvents(params *CreateDeviceEventsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeviceEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDeviceEvents",
		Method:             "POST",
		PathPattern:        "/devices/{deviceId}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeviceEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDeviceEventsOK), nil

}

/*
DeleteDevice deletes a device

Delete a device by device id.
If the token is for a SmartApp that created the device then it implicitly has permission for this api.

*/
func (a *Client) DeleteDevice(params *DeleteDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDevice",
		Method:             "DELETE",
		PathPattern:        "/devices/{deviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDeviceOK), nil

}

/*
ExecuteDeviceCommands executes commands on device

Execute commands on a device.
*/
func (a *Client) ExecuteDeviceCommands(params *ExecuteDeviceCommandsParams, authInfo runtime.ClientAuthInfoWriter) (*ExecuteDeviceCommandsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteDeviceCommandsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "executeDeviceCommands",
		Method:             "POST",
		PathPattern:        "/devices/{deviceId}/commands",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteDeviceCommandsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ExecuteDeviceCommandsOK), nil

}

/*
GetDevice gets a device s description

Get a device's description.
*/
func (a *Client) GetDevice(params *GetDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDevice",
		Method:             "GET",
		PathPattern:        "/devices/{deviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceOK), nil

}

/*
GetDeviceComponentStatus gets a device component s status

Get the status of all attributes of a the component.
The results may be filtered if the requester only has permission to view a subset of the component's capabilities.
If the token is for a SmartApp that created the device then it implicitly has permission for this api.

*/
func (a *Client) GetDeviceComponentStatus(params *GetDeviceComponentStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceComponentStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceComponentStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceComponentStatus",
		Method:             "GET",
		PathPattern:        "/devices/{deviceId}/components/{componentId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceComponentStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceComponentStatusOK), nil

}

/*
GetDeviceStatus gets the full status of a device

Get the current status of all of a device's component's attributes.
The results may be filtered if the requester only has permission to
view a subset of the device's components or capabilities.
If the token is for a SmartApp that created the device then it implicitly has permission for this api.

*/
func (a *Client) GetDeviceStatus(params *GetDeviceStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceStatus",
		Method:             "GET",
		PathPattern:        "/devices/{deviceId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceStatusOK), nil

}

/*
GetDeviceStatusByCapability gets a capability s status

Get the current status of a device component's capability.
If the token is for a SmartApp that created the device then it implicitly has permission for this api.

*/
func (a *Client) GetDeviceStatusByCapability(params *GetDeviceStatusByCapabilityParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceStatusByCapabilityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceStatusByCapabilityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceStatusByCapability",
		Method:             "GET",
		PathPattern:        "/devices/{deviceId}/components/{componentId}/capabilities/{capabilityId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceStatusByCapabilityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceStatusByCapabilityOK), nil

}

/*
GetDevices lists devices

Get a list of devices.
*/
func (a *Client) GetDevices(params *GetDevicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDevicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDevices",
		Method:             "GET",
		PathPattern:        "/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDevicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDevicesOK), nil

}

/*
InstallDevice installs a device

Install a device. This is only available for SmartApp managed devices.
The SmartApp that creates the device is responsible for handling commands for the device and
updating the status of the device by creating events.

*/
func (a *Client) InstallDevice(params *InstallDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*InstallDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInstallDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "installDevice",
		Method:             "POST",
		PathPattern:        "/devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InstallDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InstallDeviceOK), nil

}

/*
UpdateDevice updates a device

Update the properties of a device.
If the token is for a SmartApp that created the device then it implicitly has permission for this api.

*/
func (a *Client) UpdateDevice(params *UpdateDeviceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDevice",
		Method:             "PUT",
		PathPattern:        "/devices/{deviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDeviceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}