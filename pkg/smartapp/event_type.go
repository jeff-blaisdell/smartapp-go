// Code generated by go-swagger; DO NOT EDIT.

package smartapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// EventType The type of event passed to the app being executed. The type will be one of:
//   * DEVICE_EVENT - A device event as a result of a subscription the app created.
//   * TIMER_EVENT - An event as a result of a scheduled app execution.
//   * DEVICE_COMMANDS_EVENT - _Only applicable for cloud-to-cloud device integration apps._ An event as a result of a device command execution request.
//   * MODE_EVENT - A mode event is triggered when the location's mode is changed.
//
// swagger:model EventType
type EventType string

const (

	// EventTypeDEVICEEVENT captures enum value "DEVICE_EVENT"
	EventTypeDEVICEEVENT EventType = "DEVICE_EVENT"

	// EventTypeMODEEVENT captures enum value "MODE_EVENT"
	EventTypeMODEEVENT EventType = "MODE_EVENT"

	// EventTypeTIMEREVENT captures enum value "TIMER_EVENT"
	EventTypeTIMEREVENT EventType = "TIMER_EVENT"

	// EventTypeDEVICECOMMANDSEVENT captures enum value "DEVICE_COMMANDS_EVENT"
	EventTypeDEVICECOMMANDSEVENT EventType = "DEVICE_COMMANDS_EVENT"
)

// for schema
var eventTypeEnum []interface{}

func init() {
	var res []EventType
	if err := json.Unmarshal([]byte(`["DEVICE_EVENT","MODE_EVENT","TIMER_EVENT","DEVICE_COMMANDS_EVENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventTypeEnum = append(eventTypeEnum, v)
	}
}

func (m EventType) validateEventTypeEnum(path, location string, value EventType) error {
	if err := validate.Enum(path, location, value, eventTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this event type
func (m EventType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEventTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
