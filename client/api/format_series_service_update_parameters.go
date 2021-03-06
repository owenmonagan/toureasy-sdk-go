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

// NewFormatSeriesServiceUpdateParams creates a new FormatSeriesServiceUpdateParams object
// with the default values initialized.
func NewFormatSeriesServiceUpdateParams() *FormatSeriesServiceUpdateParams {
	var ()
	return &FormatSeriesServiceUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFormatSeriesServiceUpdateParamsWithTimeout creates a new FormatSeriesServiceUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFormatSeriesServiceUpdateParamsWithTimeout(timeout time.Duration) *FormatSeriesServiceUpdateParams {
	var ()
	return &FormatSeriesServiceUpdateParams{

		timeout: timeout,
	}
}

// NewFormatSeriesServiceUpdateParamsWithContext creates a new FormatSeriesServiceUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewFormatSeriesServiceUpdateParamsWithContext(ctx context.Context) *FormatSeriesServiceUpdateParams {
	var ()
	return &FormatSeriesServiceUpdateParams{

		Context: ctx,
	}
}

// NewFormatSeriesServiceUpdateParamsWithHTTPClient creates a new FormatSeriesServiceUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFormatSeriesServiceUpdateParamsWithHTTPClient(client *http.Client) *FormatSeriesServiceUpdateParams {
	var ()
	return &FormatSeriesServiceUpdateParams{
		HTTPClient: client,
	}
}

/*FormatSeriesServiceUpdateParams contains all the parameters to send to the API endpoint
for the format series service update operation typically these are written to a http.Request
*/
type FormatSeriesServiceUpdateParams struct {

	/*Body*/
	Body *models.APIFormatSeriesRequestBody
	/*ObjectID*/
	ObjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the format series service update params
func (o *FormatSeriesServiceUpdateParams) WithTimeout(timeout time.Duration) *FormatSeriesServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the format series service update params
func (o *FormatSeriesServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the format series service update params
func (o *FormatSeriesServiceUpdateParams) WithContext(ctx context.Context) *FormatSeriesServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the format series service update params
func (o *FormatSeriesServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the format series service update params
func (o *FormatSeriesServiceUpdateParams) WithHTTPClient(client *http.Client) *FormatSeriesServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the format series service update params
func (o *FormatSeriesServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the format series service update params
func (o *FormatSeriesServiceUpdateParams) WithBody(body *models.APIFormatSeriesRequestBody) *FormatSeriesServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the format series service update params
func (o *FormatSeriesServiceUpdateParams) SetBody(body *models.APIFormatSeriesRequestBody) {
	o.Body = body
}

// WithObjectID adds the objectID to the format series service update params
func (o *FormatSeriesServiceUpdateParams) WithObjectID(objectID string) *FormatSeriesServiceUpdateParams {
	o.SetObjectID(objectID)
	return o
}

// SetObjectID adds the objectId to the format series service update params
func (o *FormatSeriesServiceUpdateParams) SetObjectID(objectID string) {
	o.ObjectID = objectID
}

// WriteToRequest writes these params to a swagger request
func (o *FormatSeriesServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param object.id
	if err := r.SetPathParam("object.id", o.ObjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
