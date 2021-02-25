// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// APIStatusOptions api status options
//
// swagger:model apiStatusOptions
type APIStatusOptions string

const (

	// APIStatusOptionsWAITING captures enum value "WAITING"
	APIStatusOptionsWAITING APIStatusOptions = "WAITING"

	// APIStatusOptionsREADY captures enum value "READY"
	APIStatusOptionsREADY APIStatusOptions = "READY"

	// APIStatusOptionsACTIVE captures enum value "ACTIVE"
	APIStatusOptionsACTIVE APIStatusOptions = "ACTIVE"

	// APIStatusOptionsCOMPLETE captures enum value "COMPLETE"
	APIStatusOptionsCOMPLETE APIStatusOptions = "COMPLETE"
)

// for schema
var apiStatusOptionsEnum []interface{}

func init() {
	var res []APIStatusOptions
	if err := json.Unmarshal([]byte(`["WAITING","READY","ACTIVE","COMPLETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiStatusOptionsEnum = append(apiStatusOptionsEnum, v)
	}
}

func (m APIStatusOptions) validateAPIStatusOptionsEnum(path, location string, value APIStatusOptions) error {
	if err := validate.EnumCase(path, location, value, apiStatusOptionsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this api status options
func (m APIStatusOptions) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPIStatusOptionsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}