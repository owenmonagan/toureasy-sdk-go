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

// APIStageFromReferences api stage from references
//
// swagger:model apiStageFromReferences
type APIStageFromReferences struct {

	// stage rr ids
	StageRrIds []*APIEdge `json:"stageRrIds"`

	// stage se ids
	StageSeIds []*APIEdge `json:"stageSeIds"`
}

// Validate validates this api stage from references
func (m *APIStageFromReferences) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStageRrIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageSeIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIStageFromReferences) validateStageRrIds(formats strfmt.Registry) error {

	if swag.IsZero(m.StageRrIds) { // not required
		return nil
	}

	for i := 0; i < len(m.StageRrIds); i++ {
		if swag.IsZero(m.StageRrIds[i]) { // not required
			continue
		}

		if m.StageRrIds[i] != nil {
			if err := m.StageRrIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stageRrIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIStageFromReferences) validateStageSeIds(formats strfmt.Registry) error {

	if swag.IsZero(m.StageSeIds) { // not required
		return nil
	}

	for i := 0; i < len(m.StageSeIds); i++ {
		if swag.IsZero(m.StageSeIds[i]) { // not required
			continue
		}

		if m.StageSeIds[i] != nil {
			if err := m.StageSeIds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stageSeIds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIStageFromReferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIStageFromReferences) UnmarshalBinary(b []byte) error {
	var res APIStageFromReferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
