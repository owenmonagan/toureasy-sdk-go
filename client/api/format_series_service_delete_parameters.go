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

// NewFormatSeriesServiceDeleteParams creates a new FormatSeriesServiceDeleteParams object
// with the default values initialized.
func NewFormatSeriesServiceDeleteParams() *FormatSeriesServiceDeleteParams {
	var ()
	return &FormatSeriesServiceDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFormatSeriesServiceDeleteParamsWithTimeout creates a new FormatSeriesServiceDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFormatSeriesServiceDeleteParamsWithTimeout(timeout time.Duration) *FormatSeriesServiceDeleteParams {
	var ()
	return &FormatSeriesServiceDeleteParams{

		timeout: timeout,
	}
}

// NewFormatSeriesServiceDeleteParamsWithContext creates a new FormatSeriesServiceDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewFormatSeriesServiceDeleteParamsWithContext(ctx context.Context) *FormatSeriesServiceDeleteParams {
	var ()
	return &FormatSeriesServiceDeleteParams{

		Context: ctx,
	}
}

// NewFormatSeriesServiceDeleteParamsWithHTTPClient creates a new FormatSeriesServiceDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFormatSeriesServiceDeleteParamsWithHTTPClient(client *http.Client) *FormatSeriesServiceDeleteParams {
	var ()
	return &FormatSeriesServiceDeleteParams{
		HTTPClient: client,
	}
}

/*FormatSeriesServiceDeleteParams contains all the parameters to send to the API endpoint
for the format series service delete operation typically these are written to a http.Request
*/
type FormatSeriesServiceDeleteParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) WithTimeout(timeout time.Duration) *FormatSeriesServiceDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) WithContext(ctx context.Context) *FormatSeriesServiceDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) WithHTTPClient(client *http.Client) *FormatSeriesServiceDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) WithID(id string) *FormatSeriesServiceDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the format series service delete params
func (o *FormatSeriesServiceDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FormatSeriesServiceDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}