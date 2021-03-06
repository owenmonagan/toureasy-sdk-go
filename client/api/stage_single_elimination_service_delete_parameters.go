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

// NewStageSingleEliminationServiceDeleteParams creates a new StageSingleEliminationServiceDeleteParams object
// with the default values initialized.
func NewStageSingleEliminationServiceDeleteParams() *StageSingleEliminationServiceDeleteParams {
	var ()
	return &StageSingleEliminationServiceDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStageSingleEliminationServiceDeleteParamsWithTimeout creates a new StageSingleEliminationServiceDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStageSingleEliminationServiceDeleteParamsWithTimeout(timeout time.Duration) *StageSingleEliminationServiceDeleteParams {
	var ()
	return &StageSingleEliminationServiceDeleteParams{

		timeout: timeout,
	}
}

// NewStageSingleEliminationServiceDeleteParamsWithContext creates a new StageSingleEliminationServiceDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewStageSingleEliminationServiceDeleteParamsWithContext(ctx context.Context) *StageSingleEliminationServiceDeleteParams {
	var ()
	return &StageSingleEliminationServiceDeleteParams{

		Context: ctx,
	}
}

// NewStageSingleEliminationServiceDeleteParamsWithHTTPClient creates a new StageSingleEliminationServiceDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStageSingleEliminationServiceDeleteParamsWithHTTPClient(client *http.Client) *StageSingleEliminationServiceDeleteParams {
	var ()
	return &StageSingleEliminationServiceDeleteParams{
		HTTPClient: client,
	}
}

/*StageSingleEliminationServiceDeleteParams contains all the parameters to send to the API endpoint
for the stage single elimination service delete operation typically these are written to a http.Request
*/
type StageSingleEliminationServiceDeleteParams struct {

	/*ID*/
	ID string
	/*TournamentID*/
	TournamentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) WithTimeout(timeout time.Duration) *StageSingleEliminationServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) WithContext(ctx context.Context) *StageSingleEliminationServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) WithHTTPClient(client *http.Client) *StageSingleEliminationServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) WithID(id string) *StageSingleEliminationServiceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) SetID(id string) {
	o.ID = id
}

// WithTournamentID adds the tournamentID to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) WithTournamentID(tournamentID *string) *StageSingleEliminationServiceDeleteParams {
	o.SetTournamentID(tournamentID)
	return o
}

// SetTournamentID adds the tournamentId to the stage single elimination service delete params
func (o *StageSingleEliminationServiceDeleteParams) SetTournamentID(tournamentID *string) {
	o.TournamentID = tournamentID
}

// WriteToRequest writes these params to a swagger request
func (o *StageSingleEliminationServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
