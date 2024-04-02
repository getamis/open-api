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

// NewListMembersForAccountParams creates a new ListMembersForAccountParams object
// with the default values initialized.
func NewListMembersForAccountParams() *ListMembersForAccountParams {
	var ()
	return &ListMembersForAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListMembersForAccountParamsWithTimeout creates a new ListMembersForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListMembersForAccountParamsWithTimeout(timeout time.Duration) *ListMembersForAccountParams {
	var ()
	return &ListMembersForAccountParams{

		timeout: timeout,
	}
}

// NewListMembersForAccountParamsWithContext creates a new ListMembersForAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewListMembersForAccountParamsWithContext(ctx context.Context) *ListMembersForAccountParams {
	var ()
	return &ListMembersForAccountParams{

		Context: ctx,
	}
}

// NewListMembersForAccountParamsWithHTTPClient creates a new ListMembersForAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListMembersForAccountParamsWithHTTPClient(client *http.Client) *ListMembersForAccountParams {
	var ()
	return &ListMembersForAccountParams{
		HTTPClient: client,
	}
}

/*
ListMembersForAccountParams contains all the parameters to send to the API endpoint
for the list members for account operation typically these are written to a http.Request
*/
type ListMembersForAccountParams struct {

	/*AccountSlug*/
	AccountSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list members for account params
func (o *ListMembersForAccountParams) WithTimeout(timeout time.Duration) *ListMembersForAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list members for account params
func (o *ListMembersForAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list members for account params
func (o *ListMembersForAccountParams) WithContext(ctx context.Context) *ListMembersForAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list members for account params
func (o *ListMembersForAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list members for account params
func (o *ListMembersForAccountParams) WithHTTPClient(client *http.Client) *ListMembersForAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list members for account params
func (o *ListMembersForAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountSlug adds the accountSlug to the list members for account params
func (o *ListMembersForAccountParams) WithAccountSlug(accountSlug string) *ListMembersForAccountParams {
	o.SetAccountSlug(accountSlug)
	return o
}

// SetAccountSlug adds the accountSlug to the list members for account params
func (o *ListMembersForAccountParams) SetAccountSlug(accountSlug string) {
	o.AccountSlug = accountSlug
}

// WriteToRequest writes these params to a swagger request
func (o *ListMembersForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_slug
	if err := r.SetPathParam("account_slug", o.AccountSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
