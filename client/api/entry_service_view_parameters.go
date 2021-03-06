// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEntryServiceViewParams creates a new EntryServiceViewParams object
// with the default values initialized.
func NewEntryServiceViewParams() *EntryServiceViewParams {
	var ()
	return &EntryServiceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEntryServiceViewParamsWithTimeout creates a new EntryServiceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEntryServiceViewParamsWithTimeout(timeout time.Duration) *EntryServiceViewParams {
	var ()
	return &EntryServiceViewParams{

		timeout: timeout,
	}
}

// NewEntryServiceViewParamsWithContext creates a new EntryServiceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewEntryServiceViewParamsWithContext(ctx context.Context) *EntryServiceViewParams {
	var ()
	return &EntryServiceViewParams{

		Context: ctx,
	}
}

// NewEntryServiceViewParamsWithHTTPClient creates a new EntryServiceViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEntryServiceViewParamsWithHTTPClient(client *http.Client) *EntryServiceViewParams {
	var ()
	return &EntryServiceViewParams{
		HTTPClient: client,
	}
}

/*EntryServiceViewParams contains all the parameters to send to the API endpoint
for the entry service view operation typically these are written to a http.Request
*/
type EntryServiceViewParams struct {

	/*ID*/
	ID string
	/*TournamentID*/
	TournamentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the entry service view params
func (o *EntryServiceViewParams) WithTimeout(timeout time.Duration) *EntryServiceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the entry service view params
func (o *EntryServiceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the entry service view params
func (o *EntryServiceViewParams) WithContext(ctx context.Context) *EntryServiceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the entry service view params
func (o *EntryServiceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the entry service view params
func (o *EntryServiceViewParams) WithHTTPClient(client *http.Client) *EntryServiceViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the entry service view params
func (o *EntryServiceViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the entry service view params
func (o *EntryServiceViewParams) WithID(id string) *EntryServiceViewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the entry service view params
func (o *EntryServiceViewParams) SetID(id string) {
	o.ID = id
}

// WithTournamentID adds the tournamentID to the entry service view params
func (o *EntryServiceViewParams) WithTournamentID(tournamentID *string) *EntryServiceViewParams {
	o.SetTournamentID(tournamentID)
	return o
}

// SetTournamentID adds the tournamentId to the entry service view params
func (o *EntryServiceViewParams) SetTournamentID(tournamentID *string) {
	o.TournamentID = tournamentID
}

// WriteToRequest writes these params to a swagger request
func (o *EntryServiceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.TournamentID != nil {

		// query param tournamentId
		var qrTournamentID string
		if o.TournamentID != nil {
			qrTournamentID = *o.TournamentID
		}
		qTournamentID := qrTournamentID
		if qTournamentID != "" {
			if err := r.SetQueryParam("tournamentId", qTournamentID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
