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

	"github.com/netlify/open-api/v2/go/models"
)

// NewUpdateSplitTestParams creates a new UpdateSplitTestParams object
// with the default values initialized.
func NewUpdateSplitTestParams() *UpdateSplitTestParams {
	var ()
	return &UpdateSplitTestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSplitTestParamsWithTimeout creates a new UpdateSplitTestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSplitTestParamsWithTimeout(timeout time.Duration) *UpdateSplitTestParams {
	var ()
	return &UpdateSplitTestParams{

		timeout: timeout,
	}
}

// NewUpdateSplitTestParamsWithContext creates a new UpdateSplitTestParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSplitTestParamsWithContext(ctx context.Context) *UpdateSplitTestParams {
	var ()
	return &UpdateSplitTestParams{

		Context: ctx,
	}
}

// NewUpdateSplitTestParamsWithHTTPClient creates a new UpdateSplitTestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSplitTestParamsWithHTTPClient(client *http.Client) *UpdateSplitTestParams {
	var ()
	return &UpdateSplitTestParams{
		HTTPClient: client,
	}
}

/*
UpdateSplitTestParams contains all the parameters to send to the API endpoint
for the update split test operation typically these are written to a http.Request
*/
type UpdateSplitTestParams struct {

	/*BranchTests*/
	BranchTests *models.SplitTestSetup
	/*SiteID*/
	SiteID string
	/*SplitTestID*/
	SplitTestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update split test params
func (o *UpdateSplitTestParams) WithTimeout(timeout time.Duration) *UpdateSplitTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update split test params
func (o *UpdateSplitTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update split test params
func (o *UpdateSplitTestParams) WithContext(ctx context.Context) *UpdateSplitTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update split test params
func (o *UpdateSplitTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update split test params
func (o *UpdateSplitTestParams) WithHTTPClient(client *http.Client) *UpdateSplitTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update split test params
func (o *UpdateSplitTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranchTests adds the branchTests to the update split test params
func (o *UpdateSplitTestParams) WithBranchTests(branchTests *models.SplitTestSetup) *UpdateSplitTestParams {
	o.SetBranchTests(branchTests)
	return o
}

// SetBranchTests adds the branchTests to the update split test params
func (o *UpdateSplitTestParams) SetBranchTests(branchTests *models.SplitTestSetup) {
	o.BranchTests = branchTests
}

// WithSiteID adds the siteID to the update split test params
func (o *UpdateSplitTestParams) WithSiteID(siteID string) *UpdateSplitTestParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the update split test params
func (o *UpdateSplitTestParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSplitTestID adds the splitTestID to the update split test params
func (o *UpdateSplitTestParams) WithSplitTestID(splitTestID string) *UpdateSplitTestParams {
	o.SetSplitTestID(splitTestID)
	return o
}

// SetSplitTestID adds the splitTestId to the update split test params
func (o *UpdateSplitTestParams) SetSplitTestID(splitTestID string) {
	o.SplitTestID = splitTestID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSplitTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BranchTests != nil {
		if err := r.SetBodyParam(o.BranchTests); err != nil {
			return err
		}
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param split_test_id
	if err := r.SetPathParam("split_test_id", o.SplitTestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
