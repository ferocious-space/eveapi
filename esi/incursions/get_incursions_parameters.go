// Code generated by go-swagger; DO NOT EDIT.

package incursions

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

// NewGetIncursionsParams creates a new GetIncursionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIncursionsParams() *GetIncursionsParams {
	return &GetIncursionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIncursionsParamsWithTimeout creates a new GetIncursionsParams object
// with the ability to set a timeout on a request.
func NewGetIncursionsParamsWithTimeout(timeout time.Duration) *GetIncursionsParams {
	return &GetIncursionsParams{
		timeout: timeout,
	}
}

// NewGetIncursionsParamsWithContext creates a new GetIncursionsParams object
// with the ability to set a context for a request.
func NewGetIncursionsParamsWithContext(ctx context.Context) *GetIncursionsParams {
	return &GetIncursionsParams{
		Context: ctx,
	}
}

// NewGetIncursionsParamsWithHTTPClient creates a new GetIncursionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIncursionsParamsWithHTTPClient(client *http.Client) *GetIncursionsParams {
	return &GetIncursionsParams{
		HTTPClient: client,
	}
}

/*
GetIncursionsParams contains all the parameters to send to the API endpoint

	for the get incursions operation.

	Typically these are written to a http.Request.
*/
type GetIncursionsParams struct {

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

// WithDefaults hydrates default values in the get incursions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIncursionsParams) WithDefaults() *GetIncursionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get incursions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIncursionsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetIncursionsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get incursions params
func (o *GetIncursionsParams) WithTimeout(timeout time.Duration) *GetIncursionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get incursions params
func (o *GetIncursionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get incursions params
func (o *GetIncursionsParams) WithContext(ctx context.Context) *GetIncursionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get incursions params
func (o *GetIncursionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get incursions params
func (o *GetIncursionsParams) WithHTTPClient(client *http.Client) *GetIncursionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get incursions params
func (o *GetIncursionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get incursions params
func (o *GetIncursionsParams) WithIfNoneMatch(ifNoneMatch *string) *GetIncursionsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get incursions params
func (o *GetIncursionsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get incursions params
func (o *GetIncursionsParams) WithDatasource(datasource *string) *GetIncursionsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get incursions params
func (o *GetIncursionsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetIncursionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
