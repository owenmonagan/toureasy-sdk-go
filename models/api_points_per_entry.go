// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPointsPerEntry api points per entry
//
// swagger:model apiPointsPerEntry
type APIPointsPerEntry struct {

	// number of entries
	NumberOfEntries int32 `json:"numberOfEntries,omitempty"`

	// points
	Points int32 `json:"points,omitempty"`

	// valid for more entries
	ValidForMoreEntries bool `json:"validForMoreEntries,omitempty"`
}

// Validate validates this api points per entry
func (m *APIPointsPerEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPointsPerEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPointsPerEntry) UnmarshalBinary(b []byte) error {
	var res APIPointsPerEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
