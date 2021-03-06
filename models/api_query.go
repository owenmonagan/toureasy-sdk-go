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

// APIQuery api query
//
// swagger:model apiQuery
type APIQuery struct {

	// batch number
	BatchNumber int32 `json:"batchNumber,omitempty"`

	// entry ids
	EntryIds []string `json:"entryIds"`

	// inbound of
	InboundOf *APIAllReferences `json:"inboundOf,omitempty"`

	// outbound of
	OutboundOf *APIAllReferences `json:"outboundOf,omitempty"`

	// statuses
	Statuses []APIStatusOptions `json:"statuses"`

	// tournament Id
	TournamentID string `json:"tournamentId,omitempty"`
}

// Validate validates this api query
func (m *APIQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInboundOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatuses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIQuery) validateInboundOf(formats strfmt.Registry) error {

	if swag.IsZero(m.InboundOf) { // not required
		return nil
	}

	if m.InboundOf != nil {
		if err := m.InboundOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inboundOf")
			}
			return err
		}
	}

	return nil
}

func (m *APIQuery) validateOutboundOf(formats strfmt.Registry) error {

	if swag.IsZero(m.OutboundOf) { // not required
		return nil
	}

	if m.OutboundOf != nil {
		if err := m.OutboundOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundOf")
			}
			return err
		}
	}

	return nil
}

func (m *APIQuery) validateStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.Statuses) { // not required
		return nil
	}

	for i := 0; i < len(m.Statuses); i++ {

		if err := m.Statuses[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statuses" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIQuery) UnmarshalBinary(b []byte) error {
	var res APIQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
