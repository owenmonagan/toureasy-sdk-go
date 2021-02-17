// Code generated by go-swagger; DO NOT EDIT.

package round_robin

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

	"github.com/owenmonagan/toureasy-sdk-go/gen/models"
)

// NewStageRoundRobinServiceUpdateParams creates a new StageRoundRobinServiceUpdateParams object
// with the default values initialized.
func NewStageRoundRobinServiceUpdateParams() *StageRoundRobinServiceUpdateParams {
	var ()
	return &StageRoundRobinServiceUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStageRoundRobinServiceUpdateParamsWithTimeout creates a new StageRoundRobinServiceUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStageRoundRobinServiceUpdateParamsWithTimeout(timeout time.Duration) *StageRoundRobinServiceUpdateParams {
	var ()
	return &StageRoundRobinServiceUpdateParams{

		timeout: timeout,
	}
}

// NewStageRoundRobinServiceUpdateParamsWithContext creates a new StageRoundRobinServiceUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewStageRoundRobinServiceUpdateParamsWithContext(ctx context.Context) *StageRoundRobinServiceUpdateParams {
	var ()
	return &StageRoundRobinServiceUpdateParams{

		Context: ctx,
	}
}

// NewStageRoundRobinServiceUpdateParamsWithHTTPClient creates a new StageRoundRobinServiceUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStageRoundRobinServiceUpdateParamsWithHTTPClient(client *http.Client) *StageRoundRobinServiceUpdateParams {
	var ()
	return &StageRoundRobinServiceUpdateParams{
		HTTPClient: client,
	}
}

/*StageRoundRobinServiceUpdateParams contains all the parameters to send to the API endpoint
for the stage round robin service update operation typically these are written to a http.Request
*/
type StageRoundRobinServiceUpdateParams struct {

	/*Body*/
	Body *models.APIStageRoundRobinUpdateRequest
	/*UUID*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) WithTimeout(timeout time.Duration) *StageRoundRobinServiceUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) WithContext(ctx context.Context) *StageRoundRobinServiceUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) WithHTTPClient(client *http.Client) *StageRoundRobinServiceUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) WithBody(body *models.APIStageRoundRobinUpdateRequest) *StageRoundRobinServiceUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) SetBody(body *models.APIStageRoundRobinUpdateRequest) {
	o.Body = body
}

// WithUUID adds the uuid to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) WithUUID(uuid string) *StageRoundRobinServiceUpdateParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the stage round robin service update params
func (o *StageRoundRobinServiceUpdateParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StageRoundRobinServiceUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}