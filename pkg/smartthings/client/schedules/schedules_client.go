// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new schedules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schedules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSchedule saves an installed app schedule

Create a schedule for an installed app. The installed app must be in the location specified and the installed app must have permission to create schedules.

*/
func (a *Client) CreateSchedule(params *CreateScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateScheduleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSchedule",
		Method:             "POST",
		PathPattern:        "/installedapps/{installedAppId}/schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateScheduleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateScheduleOK), nil

}

/*
DeleteSchedule deletes a schedule

Delete a specific schedule for the installed app.

*/
func (a *Client) DeleteSchedule(params *DeleteScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteScheduleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSchedule",
		Method:             "DELETE",
		PathPattern:        "/installedapps/{installedAppId}/schedules/{scheduleName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteScheduleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteScheduleOK), nil

}

/*
DeleteSchedules deletes all of an installed app s schedules

Delete all schedules for the installed app.

*/
func (a *Client) DeleteSchedules(params *DeleteSchedulesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSchedulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSchedules",
		Method:             "DELETE",
		PathPattern:        "/installedapps/{installedAppId}/schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSchedulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSchedulesOK), nil

}

/*
GetSchedule gets an installed app s schedule

Get a specific schedule for the installed app.

*/
func (a *Client) GetSchedule(params *GetScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*GetScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSchedule",
		Method:             "GET",
		PathPattern:        "/installedapps/{installedAppId}/schedules/{scheduleName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScheduleOK), nil

}

/*
GetSchedules lists installed app schedules

List the schedules for the installed app.

*/
func (a *Client) GetSchedules(params *GetSchedulesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchedulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSchedules",
		Method:             "GET",
		PathPattern:        "/installedapps/{installedAppId}/schedules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchedulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSchedulesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
