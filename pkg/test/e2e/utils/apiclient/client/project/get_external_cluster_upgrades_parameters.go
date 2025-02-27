// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetExternalClusterUpgradesParams creates a new GetExternalClusterUpgradesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalClusterUpgradesParams() *GetExternalClusterUpgradesParams {
	return &GetExternalClusterUpgradesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalClusterUpgradesParamsWithTimeout creates a new GetExternalClusterUpgradesParams object
// with the ability to set a timeout on a request.
func NewGetExternalClusterUpgradesParamsWithTimeout(timeout time.Duration) *GetExternalClusterUpgradesParams {
	return &GetExternalClusterUpgradesParams{
		timeout: timeout,
	}
}

// NewGetExternalClusterUpgradesParamsWithContext creates a new GetExternalClusterUpgradesParams object
// with the ability to set a context for a request.
func NewGetExternalClusterUpgradesParamsWithContext(ctx context.Context) *GetExternalClusterUpgradesParams {
	return &GetExternalClusterUpgradesParams{
		Context: ctx,
	}
}

// NewGetExternalClusterUpgradesParamsWithHTTPClient creates a new GetExternalClusterUpgradesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalClusterUpgradesParamsWithHTTPClient(client *http.Client) *GetExternalClusterUpgradesParams {
	return &GetExternalClusterUpgradesParams{
		HTTPClient: client,
	}
}

/* GetExternalClusterUpgradesParams contains all the parameters to send to the API endpoint
   for the get external cluster upgrades operation.

   Typically these are written to a http.Request.
*/
type GetExternalClusterUpgradesParams struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get external cluster upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalClusterUpgradesParams) WithDefaults() *GetExternalClusterUpgradesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get external cluster upgrades params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalClusterUpgradesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) WithTimeout(timeout time.Duration) *GetExternalClusterUpgradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) WithContext(ctx context.Context) *GetExternalClusterUpgradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) WithHTTPClient(client *http.Client) *GetExternalClusterUpgradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) WithClusterID(clusterID string) *GetExternalClusterUpgradesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) WithProjectID(projectID string) *GetExternalClusterUpgradesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get external cluster upgrades params
func (o *GetExternalClusterUpgradesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalClusterUpgradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
