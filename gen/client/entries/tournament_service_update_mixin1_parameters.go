// Code generated by go-swagger; DO NOT EDIT.

package entries

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

// NewTournamentServiceUpdateMixin1Params creates a new TournamentServiceUpdateMixin1Params object
// with the default values initialized.
func NewTournamentServiceUpdateMixin1Params() *TournamentServiceUpdateMixin1Params {
	var ()
	return &TournamentServiceUpdateMixin1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewTournamentServiceUpdateMixin1ParamsWithTimeout creates a new TournamentServiceUpdateMixin1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewTournamentServiceUpdateMixin1ParamsWithTimeout(timeout time.Duration) *TournamentServiceUpdateMixin1Params {
	var ()
	return &TournamentServiceUpdateMixin1Params{

		timeout: timeout,
	}
}

// NewTournamentServiceUpdateMixin1ParamsWithContext creates a new TournamentServiceUpdateMixin1Params object
// with the default values initialized, and the ability to set a context for a request
func NewTournamentServiceUpdateMixin1ParamsWithContext(ctx context.Context) *TournamentServiceUpdateMixin1Params {
	var ()
	return &TournamentServiceUpdateMixin1Params{

		Context: ctx,
	}
}

// NewTournamentServiceUpdateMixin1ParamsWithHTTPClient creates a new TournamentServiceUpdateMixin1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTournamentServiceUpdateMixin1ParamsWithHTTPClient(client *http.Client) *TournamentServiceUpdateMixin1Params {
	var ()
	return &TournamentServiceUpdateMixin1Params{
		HTTPClient: client,
	}
}

/*TournamentServiceUpdateMixin1Params contains all the parameters to send to the API endpoint
for the tournament service update mixin1 operation typically these are written to a http.Request
*/
type TournamentServiceUpdateMixin1Params struct {

	/*Body*/
	Body *models.APIUpdateRequest
	/*UUID*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) WithTimeout(timeout time.Duration) *TournamentServiceUpdateMixin1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) WithContext(ctx context.Context) *TournamentServiceUpdateMixin1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) WithHTTPClient(client *http.Client) *TournamentServiceUpdateMixin1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) WithBody(body *models.APIUpdateRequest) *TournamentServiceUpdateMixin1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) SetBody(body *models.APIUpdateRequest) {
	o.Body = body
}

// WithUUID adds the uuid to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) WithUUID(uuid string) *TournamentServiceUpdateMixin1Params {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the tournament service update mixin1 params
func (o *TournamentServiceUpdateMixin1Params) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *TournamentServiceUpdateMixin1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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