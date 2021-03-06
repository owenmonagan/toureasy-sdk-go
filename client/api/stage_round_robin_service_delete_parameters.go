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

// NewStageRoundRobinServiceDeleteParams creates a new StageRoundRobinServiceDeleteParams object
// with the default values initialized.
func NewStageRoundRobinServiceDeleteParams() *StageRoundRobinServiceDeleteParams {
	var ()
	return &StageRoundRobinServiceDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStageRoundRobinServiceDeleteParamsWithTimeout creates a new StageRoundRobinServiceDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStageRoundRobinServiceDeleteParamsWithTimeout(timeout time.Duration) *StageRoundRobinServiceDeleteParams {
	var ()
	return &StageRoundRobinServiceDeleteParams{

		timeout: timeout,
	}
}

// NewStageRoundRobinServiceDeleteParamsWithContext creates a new StageRoundRobinServiceDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewStageRoundRobinServiceDeleteParamsWithContext(ctx context.Context) *StageRoundRobinServiceDeleteParams {
	var ()
	return &StageRoundRobinServiceDeleteParams{

		Context: ctx,
	}
}

// NewStageRoundRobinServiceDeleteParamsWithHTTPClient creates a new StageRoundRobinServiceDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStageRoundRobinServiceDeleteParamsWithHTTPClient(client *http.Client) *StageRoundRobinServiceDeleteParams {
	var ()
	return &StageRoundRobinServiceDeleteParams{
		HTTPClient: client,
	}
}

/*StageRoundRobinServiceDeleteParams contains all the parameters to send to the API endpoint
for the stage round robin service delete operation typically these are written to a http.Request
*/
type StageRoundRobinServiceDeleteParams struct {

	/*ID*/
	ID string
	/*TournamentID*/
	TournamentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) WithTimeout(timeout time.Duration) *StageRoundRobinServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) WithContext(ctx context.Context) *StageRoundRobinServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) WithHTTPClient(client *http.Client) *StageRoundRobinServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) WithID(id string) *StageRoundRobinServiceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) SetID(id string) {
	o.ID = id
}

// WithTournamentID adds the tournamentID to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) WithTournamentID(tournamentID *string) *StageRoundRobinServiceDeleteParams {
	o.SetTournamentID(tournamentID)
	return o
}

// SetTournamentID adds the tournamentId to the stage round robin service delete params
func (o *StageRoundRobinServiceDeleteParams) SetTournamentID(tournamentID *string) {
	o.TournamentID = tournamentID
}

// WriteToRequest writes these params to a swagger request
func (o *StageRoundRobinServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
