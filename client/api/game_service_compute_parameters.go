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

	"github.com/owenmonagan/toureasy-sdk-go/models"
)

// NewGameServiceComputeParams creates a new GameServiceComputeParams object
// with the default values initialized.
func NewGameServiceComputeParams() *GameServiceComputeParams {
	var ()
	return &GameServiceComputeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGameServiceComputeParamsWithTimeout creates a new GameServiceComputeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGameServiceComputeParamsWithTimeout(timeout time.Duration) *GameServiceComputeParams {
	var ()
	return &GameServiceComputeParams{

		timeout: timeout,
	}
}

// NewGameServiceComputeParamsWithContext creates a new GameServiceComputeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGameServiceComputeParamsWithContext(ctx context.Context) *GameServiceComputeParams {
	var ()
	return &GameServiceComputeParams{

		Context: ctx,
	}
}

// NewGameServiceComputeParamsWithHTTPClient creates a new GameServiceComputeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGameServiceComputeParamsWithHTTPClient(client *http.Client) *GameServiceComputeParams {
	var ()
	return &GameServiceComputeParams{
		HTTPClient: client,
	}
}

/*GameServiceComputeParams contains all the parameters to send to the API endpoint
for the game service compute operation typically these are written to a http.Request
*/
type GameServiceComputeParams struct {

	/*Body*/
	Body *models.APIID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the game service compute params
func (o *GameServiceComputeParams) WithTimeout(timeout time.Duration) *GameServiceComputeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the game service compute params
func (o *GameServiceComputeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the game service compute params
func (o *GameServiceComputeParams) WithContext(ctx context.Context) *GameServiceComputeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the game service compute params
func (o *GameServiceComputeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the game service compute params
func (o *GameServiceComputeParams) WithHTTPClient(client *http.Client) *GameServiceComputeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the game service compute params
func (o *GameServiceComputeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the game service compute params
func (o *GameServiceComputeParams) WithBody(body *models.APIID) *GameServiceComputeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the game service compute params
func (o *GameServiceComputeParams) SetBody(body *models.APIID) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GameServiceComputeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}