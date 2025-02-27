// Code generated by go-swagger; DO NOT EDIT.

package aks

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

// NewListAKSNodePoolModesParams creates a new ListAKSNodePoolModesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAKSNodePoolModesParams() *ListAKSNodePoolModesParams {
	return &ListAKSNodePoolModesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAKSNodePoolModesParamsWithTimeout creates a new ListAKSNodePoolModesParams object
// with the ability to set a timeout on a request.
func NewListAKSNodePoolModesParamsWithTimeout(timeout time.Duration) *ListAKSNodePoolModesParams {
	return &ListAKSNodePoolModesParams{
		timeout: timeout,
	}
}

// NewListAKSNodePoolModesParamsWithContext creates a new ListAKSNodePoolModesParams object
// with the ability to set a context for a request.
func NewListAKSNodePoolModesParamsWithContext(ctx context.Context) *ListAKSNodePoolModesParams {
	return &ListAKSNodePoolModesParams{
		Context: ctx,
	}
}

// NewListAKSNodePoolModesParamsWithHTTPClient creates a new ListAKSNodePoolModesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAKSNodePoolModesParamsWithHTTPClient(client *http.Client) *ListAKSNodePoolModesParams {
	return &ListAKSNodePoolModesParams{
		HTTPClient: client,
	}
}

/* ListAKSNodePoolModesParams contains all the parameters to send to the API endpoint
   for the list a k s node pool modes operation.

   Typically these are written to a http.Request.
*/
type ListAKSNodePoolModesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list a k s node pool modes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAKSNodePoolModesParams) WithDefaults() *ListAKSNodePoolModesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list a k s node pool modes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAKSNodePoolModesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list a k s node pool modes params
func (o *ListAKSNodePoolModesParams) WithTimeout(timeout time.Duration) *ListAKSNodePoolModesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list a k s node pool modes params
func (o *ListAKSNodePoolModesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list a k s node pool modes params
func (o *ListAKSNodePoolModesParams) WithContext(ctx context.Context) *ListAKSNodePoolModesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list a k s node pool modes params
func (o *ListAKSNodePoolModesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list a k s node pool modes params
func (o *ListAKSNodePoolModesParams) WithHTTPClient(client *http.Client) *ListAKSNodePoolModesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list a k s node pool modes params
func (o *ListAKSNodePoolModesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAKSNodePoolModesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
