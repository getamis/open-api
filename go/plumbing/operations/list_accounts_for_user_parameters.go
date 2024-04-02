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

// NewListAccountsForUserParams creates a new ListAccountsForUserParams object
// with the default values initialized.
func NewListAccountsForUserParams() *ListAccountsForUserParams {

	return &ListAccountsForUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAccountsForUserParamsWithTimeout creates a new ListAccountsForUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAccountsForUserParamsWithTimeout(timeout time.Duration) *ListAccountsForUserParams {

	return &ListAccountsForUserParams{

		timeout: timeout,
	}
}

// NewListAccountsForUserParamsWithContext creates a new ListAccountsForUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAccountsForUserParamsWithContext(ctx context.Context) *ListAccountsForUserParams {

	return &ListAccountsForUserParams{

		Context: ctx,
	}
}

// NewListAccountsForUserParamsWithHTTPClient creates a new ListAccountsForUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAccountsForUserParamsWithHTTPClient(client *http.Client) *ListAccountsForUserParams {

	return &ListAccountsForUserParams{
		HTTPClient: client,
	}
}

/*
ListAccountsForUserParams contains all the parameters to send to the API endpoint
for the list accounts for user operation typically these are written to a http.Request
*/
type ListAccountsForUserParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list accounts for user params
func (o *ListAccountsForUserParams) WithTimeout(timeout time.Duration) *ListAccountsForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list accounts for user params
func (o *ListAccountsForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list accounts for user params
func (o *ListAccountsForUserParams) WithContext(ctx context.Context) *ListAccountsForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list accounts for user params
func (o *ListAccountsForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list accounts for user params
func (o *ListAccountsForUserParams) WithHTTPClient(client *http.Client) *ListAccountsForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list accounts for user params
func (o *ListAccountsForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAccountsForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
