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

// NewGameServiceViewParams creates a new GameServiceViewParams object
// with the default values initialized.
func NewGameServiceViewParams() *GameServiceViewParams {
	var ()
	return &GameServiceViewParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGameServiceViewParamsWithTimeout creates a new GameServiceViewParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGameServiceViewParamsWithTimeout(timeout time.Duration) *GameServiceViewParams {
	var ()
	return &GameServiceViewParams{

		timeout: timeout,
	}
}

// NewGameServiceViewParamsWithContext creates a new GameServiceViewParams object
// with the default values initialized, and the ability to set a context for a request
func NewGameServiceViewParamsWithContext(ctx context.Context) *GameServiceViewParams {
	var ()
	return &GameServiceViewParams{

		Context: ctx,
	}
}

// NewGameServiceViewParamsWithHTTPClient creates a new GameServiceViewParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGameServiceViewParamsWithHTTPClient(client *http.Client) *GameServiceViewParams {
	var ()
	return &GameServiceViewParams{
		HTTPClient: client,
	}
}

/*GameServiceViewParams contains all the parameters to send to the API endpoint
for the game service view operation typically these are written to a http.Request
*/
type GameServiceViewParams struct {

	/*ID*/
	ID string
	/*TournamentID*/
	TournamentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the game service view params
func (o *GameServiceViewParams) WithTimeout(timeout time.Duration) *GameServiceViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the game service view params
func (o *GameServiceViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the game service view params
func (o *GameServiceViewParams) WithContext(ctx context.Context) *GameServiceViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the game service view params
func (o *GameServiceViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the game service view params
func (o *GameServiceViewParams) WithHTTPClient(client *http.Client) *GameServiceViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the game service view params
func (o *GameServiceViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the game service view params
func (o *GameServiceViewParams) WithID(id string) *GameServiceViewParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the game service view params
func (o *GameServiceViewParams) SetID(id string) {
	o.ID = id
}

// WithTournamentID adds the tournamentID to the game service view params
func (o *GameServiceViewParams) WithTournamentID(tournamentID *string) *GameServiceViewParams {
	o.SetTournamentID(tournamentID)
	return o
}

// SetTournamentID adds the tournamentId to the game service view params
func (o *GameServiceViewParams) SetTournamentID(tournamentID *string) {
	o.TournamentID = tournamentID
}

// WriteToRequest writes these params to a swagger request
func (o *GameServiceViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
