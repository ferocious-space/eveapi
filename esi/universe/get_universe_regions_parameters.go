// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// NewGetUniverseRegionsParams creates a new GetUniverseRegionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUniverseRegionsParams() *GetUniverseRegionsParams {
	return &GetUniverseRegionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseRegionsParamsWithTimeout creates a new GetUniverseRegionsParams object
// with the ability to set a timeout on a request.
func NewGetUniverseRegionsParamsWithTimeout(timeout time.Duration) *GetUniverseRegionsParams {
	return &GetUniverseRegionsParams{
		timeout: timeout,
	}
}

// NewGetUniverseRegionsParamsWithContext creates a new GetUniverseRegionsParams object
// with the ability to set a context for a request.
func NewGetUniverseRegionsParamsWithContext(ctx context.Context) *GetUniverseRegionsParams {
	return &GetUniverseRegionsParams{
		Context: ctx,
	}
}

// NewGetUniverseRegionsParamsWithHTTPClient creates a new GetUniverseRegionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUniverseRegionsParamsWithHTTPClient(client *http.Client) *GetUniverseRegionsParams {
	return &GetUniverseRegionsParams{
		HTTPClient: client,
	}
}

/*
GetUniverseRegionsParams contains all the parameters to send to the API endpoint

	for the get universe regions operation.

	Typically these are written to a http.Request.
*/
type GetUniverseRegionsParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get universe regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseRegionsParams) WithDefaults() *GetUniverseRegionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get universe regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseRegionsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetUniverseRegionsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get universe regions params
func (o *GetUniverseRegionsParams) WithTimeout(timeout time.Duration) *GetUniverseRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe regions params
func (o *GetUniverseRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe regions params
func (o *GetUniverseRegionsParams) WithContext(ctx context.Context) *GetUniverseRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe regions params
func (o *GetUniverseRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe regions params
func (o *GetUniverseRegionsParams) WithHTTPClient(client *http.Client) *GetUniverseRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe regions params
func (o *GetUniverseRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe regions params
func (o *GetUniverseRegionsParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseRegionsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe regions params
func (o *GetUniverseRegionsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get universe regions params
func (o *GetUniverseRegionsParams) WithDatasource(datasource *string) *GetUniverseRegionsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe regions params
func (o *GetUniverseRegionsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}
	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string

		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {

			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
