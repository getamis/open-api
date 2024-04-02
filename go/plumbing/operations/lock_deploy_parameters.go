// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewLockDeployParams creates a new LockDeployParams object
// with the default values initialized.
func NewLockDeployParams() *LockDeployParams {
	var ()
	return &LockDeployParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLockDeployParamsWithTimeout creates a new LockDeployParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLockDeployParamsWithTimeout(timeout time.Duration) *LockDeployParams {
	var ()
	return &LockDeployParams{

		timeout: timeout,
	}
}

// NewLockDeployParamsWithContext creates a new LockDeployParams object
// with the default values initialized, and the ability to set a context for a request
func NewLockDeployParamsWithContext(ctx context.Context) *LockDeployParams {
	var ()
	return &LockDeployParams{

		Context: ctx,
	}
}

// NewLockDeployParamsWithHTTPClient creates a new LockDeployParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLockDeployParamsWithHTTPClient(client *http.Client) *LockDeployParams {
	var ()
	return &LockDeployParams{
		HTTPClient: client,
	}
}

/*
LockDeployParams contains all the parameters to send to the API endpoint
for the lock deploy operation typically these are written to a http.Request
*/
type LockDeployParams struct {

	/*DeployID*/
	DeployID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lock deploy params
func (o *LockDeployParams) WithTimeout(timeout time.Duration) *LockDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lock deploy params
func (o *LockDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lock deploy params
func (o *LockDeployParams) WithContext(ctx context.Context) *LockDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lock deploy params
func (o *LockDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lock deploy params
func (o *LockDeployParams) WithHTTPClient(client *http.Client) *LockDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lock deploy params
func (o *LockDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployID adds the deployID to the lock deploy params
func (o *LockDeployParams) WithDeployID(deployID string) *LockDeployParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the lock deploy params
func (o *LockDeployParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WriteToRequest writes these params to a swagger request
func (o *LockDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
