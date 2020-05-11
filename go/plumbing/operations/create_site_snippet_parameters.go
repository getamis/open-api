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

	"github.com/netlify/open-api/go/models"
)

// NewCreateSiteSnippetParams creates a new CreateSiteSnippetParams object
// with the default values initialized.
func NewCreateSiteSnippetParams() *CreateSiteSnippetParams {
	var ()
	return &CreateSiteSnippetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSiteSnippetParamsWithTimeout creates a new CreateSiteSnippetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSiteSnippetParamsWithTimeout(timeout time.Duration) *CreateSiteSnippetParams {
	var ()
	return &CreateSiteSnippetParams{

		timeout: timeout,
	}
}

// NewCreateSiteSnippetParamsWithContext creates a new CreateSiteSnippetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSiteSnippetParamsWithContext(ctx context.Context) *CreateSiteSnippetParams {
	var ()
	return &CreateSiteSnippetParams{

		Context: ctx,
	}
}

// NewCreateSiteSnippetParamsWithHTTPClient creates a new CreateSiteSnippetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSiteSnippetParamsWithHTTPClient(client *http.Client) *CreateSiteSnippetParams {
	var ()
	return &CreateSiteSnippetParams{
		HTTPClient: client,
	}
}

/*CreateSiteSnippetParams contains all the parameters to send to the API endpoint
for the create site snippet operation typically these are written to a http.Request
*/
type CreateSiteSnippetParams struct {

	/*SiteID*/
	SiteID string
	/*Snippet*/
	Snippet *models.Snippet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create site snippet params
func (o *CreateSiteSnippetParams) WithTimeout(timeout time.Duration) *CreateSiteSnippetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create site snippet params
func (o *CreateSiteSnippetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create site snippet params
func (o *CreateSiteSnippetParams) WithContext(ctx context.Context) *CreateSiteSnippetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create site snippet params
func (o *CreateSiteSnippetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create site snippet params
func (o *CreateSiteSnippetParams) WithHTTPClient(client *http.Client) *CreateSiteSnippetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create site snippet params
func (o *CreateSiteSnippetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the create site snippet params
func (o *CreateSiteSnippetParams) WithSiteID(siteID string) *CreateSiteSnippetParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the create site snippet params
func (o *CreateSiteSnippetParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSnippet adds the snippet to the create site snippet params
func (o *CreateSiteSnippetParams) WithSnippet(snippet *models.Snippet) *CreateSiteSnippetParams {
	o.SetSnippet(snippet)
	return o
}

// SetSnippet adds the snippet to the create site snippet params
func (o *CreateSiteSnippetParams) SetSnippet(snippet *models.Snippet) {
	o.Snippet = snippet
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSiteSnippetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if o.Snippet != nil {
		if err := r.SetBodyParam(o.Snippet); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
