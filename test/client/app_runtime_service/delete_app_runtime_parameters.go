// Code generated by go-swagger; DO NOT EDIT.

package app_runtime_service

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

// NewDeleteAppRuntimeParams creates a new DeleteAppRuntimeParams object
// with the default values initialized.
func NewDeleteAppRuntimeParams() *DeleteAppRuntimeParams {
	var ()
	return &DeleteAppRuntimeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppRuntimeParamsWithTimeout creates a new DeleteAppRuntimeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAppRuntimeParamsWithTimeout(timeout time.Duration) *DeleteAppRuntimeParams {
	var ()
	return &DeleteAppRuntimeParams{

		timeout: timeout,
	}
}

// NewDeleteAppRuntimeParamsWithContext creates a new DeleteAppRuntimeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAppRuntimeParamsWithContext(ctx context.Context) *DeleteAppRuntimeParams {
	var ()
	return &DeleteAppRuntimeParams{

		Context: ctx,
	}
}

// NewDeleteAppRuntimeParamsWithHTTPClient creates a new DeleteAppRuntimeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAppRuntimeParamsWithHTTPClient(client *http.Client) *DeleteAppRuntimeParams {
	var ()
	return &DeleteAppRuntimeParams{
		HTTPClient: client,
	}
}

/*DeleteAppRuntimeParams contains all the parameters to send to the API endpoint
for the delete app runtime operation typically these are written to a http.Request
*/
type DeleteAppRuntimeParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete app runtime params
func (o *DeleteAppRuntimeParams) WithTimeout(timeout time.Duration) *DeleteAppRuntimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete app runtime params
func (o *DeleteAppRuntimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete app runtime params
func (o *DeleteAppRuntimeParams) WithContext(ctx context.Context) *DeleteAppRuntimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete app runtime params
func (o *DeleteAppRuntimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete app runtime params
func (o *DeleteAppRuntimeParams) WithHTTPClient(client *http.Client) *DeleteAppRuntimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete app runtime params
func (o *DeleteAppRuntimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete app runtime params
func (o *DeleteAppRuntimeParams) WithID(id string) *DeleteAppRuntimeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete app runtime params
func (o *DeleteAppRuntimeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppRuntimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
