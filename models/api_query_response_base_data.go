// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIQueryResponseBaseData api query response base data
//
// swagger:model apiQueryResponseBaseData
type APIQueryResponseBaseData struct {

	// inbound has
	InboundHas *APIAllReferences `json:"inboundHas,omitempty"`
}

// Validate validates this api query response base data
func (m *APIQueryResponseBaseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInboundHas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIQueryResponseBaseData) validateInboundHas(formats strfmt.Registry) error {

	if swag.IsZero(m.InboundHas) { // not required
		return nil
	}

	if m.InboundHas != nil {
		if err := m.InboundHas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inboundHas")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIQueryResponseBaseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIQueryResponseBaseData) UnmarshalBinary(b []byte) error {
	var res APIQueryResponseBaseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
