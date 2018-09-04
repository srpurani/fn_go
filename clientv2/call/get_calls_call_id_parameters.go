// Code generated by go-swagger; DO NOT EDIT.

package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCallsCallIDParams creates a new GetCallsCallIDParams object
// with the default values initialized.
func NewGetCallsCallIDParams() *GetCallsCallIDParams {
	var ()
	return &GetCallsCallIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCallsCallIDParamsWithTimeout creates a new GetCallsCallIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCallsCallIDParamsWithTimeout(timeout time.Duration) *GetCallsCallIDParams {
	var ()
	return &GetCallsCallIDParams{

		timeout: timeout,
	}
}

// NewGetCallsCallIDParamsWithContext creates a new GetCallsCallIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCallsCallIDParamsWithContext(ctx context.Context) *GetCallsCallIDParams {
	var ()
	return &GetCallsCallIDParams{

		Context: ctx,
	}
}

// NewGetCallsCallIDParamsWithHTTPClient creates a new GetCallsCallIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCallsCallIDParamsWithHTTPClient(client *http.Client) *GetCallsCallIDParams {
	var ()
	return &GetCallsCallIDParams{
		HTTPClient: client,
	}
}

/*GetCallsCallIDParams contains all the parameters to send to the API endpoint
for the get calls call ID operation typically these are written to a http.Request
*/
type GetCallsCallIDParams struct {

	/*AppID
	  Application ID.

	*/
	AppID *string
	/*CallID
	  Opaque, unique Call ID.

	*/
	CallID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get calls call ID params
func (o *GetCallsCallIDParams) WithTimeout(timeout time.Duration) *GetCallsCallIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get calls call ID params
func (o *GetCallsCallIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get calls call ID params
func (o *GetCallsCallIDParams) WithContext(ctx context.Context) *GetCallsCallIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get calls call ID params
func (o *GetCallsCallIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get calls call ID params
func (o *GetCallsCallIDParams) WithHTTPClient(client *http.Client) *GetCallsCallIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get calls call ID params
func (o *GetCallsCallIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get calls call ID params
func (o *GetCallsCallIDParams) WithAppID(appID *string) *GetCallsCallIDParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get calls call ID params
func (o *GetCallsCallIDParams) SetAppID(appID *string) {
	o.AppID = appID
}

// WithCallID adds the callID to the get calls call ID params
func (o *GetCallsCallIDParams) WithCallID(callID string) *GetCallsCallIDParams {
	o.SetCallID(callID)
	return o
}

// SetCallID adds the callId to the get calls call ID params
func (o *GetCallsCallIDParams) SetCallID(callID string) {
	o.CallID = callID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCallsCallIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppID != nil {

		// query param app_id
		var qrAppID string
		if o.AppID != nil {
			qrAppID = *o.AppID
		}
		qAppID := qrAppID
		if qAppID != "" {
			if err := r.SetQueryParam("app_id", qAppID); err != nil {
				return err
			}
		}

	}

	// path param callID
	if err := r.SetPathParam("callID", o.CallID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}