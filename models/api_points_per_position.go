// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPointsPerPosition api points per position
//
// swagger:model apiPointsPerPosition
type APIPointsPerPosition struct {

	// points per entry
	PointsPerEntry []*APIPointsPerEntry `json:"pointsPerEntry"`

	// position
	Position int32 `json:"position,omitempty"`
}

// Validate validates this api points per position
func (m *APIPointsPerPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePointsPerEntry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPointsPerPosition) validatePointsPerEntry(formats strfmt.Registry) error {

	if swag.IsZero(m.PointsPerEntry) { // not required
		return nil
	}

	for i := 0; i < len(m.PointsPerEntry); i++ {
		if swag.IsZero(m.PointsPerEntry[i]) { // not required
			continue
		}

		if m.PointsPerEntry[i] != nil {
			if err := m.PointsPerEntry[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pointsPerEntry" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIPointsPerPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPointsPerPosition) UnmarshalBinary(b []byte) error {
	var res APIPointsPerPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
