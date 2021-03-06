// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPosition api position
//
// swagger:model apiPosition
type APIPosition struct {

	// id
	ID string `json:"id,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`

	// timestamps
	Timestamps *APITimestamps `json:"timestamps,omitempty"`
}

// Validate validates this api position
func (m *APIPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPosition) validateTimestamps(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamps) { // not required
		return nil
	}

	if m.Timestamps != nil {
		if err := m.Timestamps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timestamps")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPosition) UnmarshalBinary(b []byte) error {
	var res APIPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
