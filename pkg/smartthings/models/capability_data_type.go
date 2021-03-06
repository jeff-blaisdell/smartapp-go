// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CapabilityDataType __DEPRECATED__ only available while migrating from `type` to `schema`
// swagger:model CapabilityDataType
type CapabilityDataType string

const (

	// CapabilityDataTypeSTRING captures enum value "STRING"
	CapabilityDataTypeSTRING CapabilityDataType = "STRING"

	// CapabilityDataTypeNUMBER captures enum value "NUMBER"
	CapabilityDataTypeNUMBER CapabilityDataType = "NUMBER"

	// CapabilityDataTypeENUM captures enum value "ENUM"
	CapabilityDataTypeENUM CapabilityDataType = "ENUM"

	// CapabilityDataTypeJSONOBJECT captures enum value "JSON_OBJECT"
	CapabilityDataTypeJSONOBJECT CapabilityDataType = "JSON_OBJECT"

	// CapabilityDataTypeDATE captures enum value "DATE"
	CapabilityDataTypeDATE CapabilityDataType = "DATE"

	// CapabilityDataTypeVECTOR3 captures enum value "VECTOR3"
	CapabilityDataTypeVECTOR3 CapabilityDataType = "VECTOR3"

	// CapabilityDataTypeDYNAMICENUM captures enum value "DYNAMIC_ENUM"
	CapabilityDataTypeDYNAMICENUM CapabilityDataType = "DYNAMIC_ENUM"

	// CapabilityDataTypeCOLORMAP captures enum value "COLOR_MAP"
	CapabilityDataTypeCOLORMAP CapabilityDataType = "COLOR_MAP"

	// CapabilityDataTypeBOOLEAN captures enum value "BOOLEAN"
	CapabilityDataTypeBOOLEAN CapabilityDataType = "BOOLEAN"
)

// for schema
var capabilityDataTypeEnum []interface{}

func init() {
	var res []CapabilityDataType
	if err := json.Unmarshal([]byte(`["STRING","NUMBER","ENUM","JSON_OBJECT","DATE","VECTOR3","DYNAMIC_ENUM","COLOR_MAP","BOOLEAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		capabilityDataTypeEnum = append(capabilityDataTypeEnum, v)
	}
}

func (m CapabilityDataType) validateCapabilityDataTypeEnum(path, location string, value CapabilityDataType) error {
	if err := validate.Enum(path, location, value, capabilityDataTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this capability data type
func (m CapabilityDataType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCapabilityDataTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
