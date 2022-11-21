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
	"github.com/go-openapi/swag"
)

// NewGetUniverseMoonsMoonIDParams creates a new GetUniverseMoonsMoonIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUniverseMoonsMoonIDParams() *GetUniverseMoonsMoonIDParams {
	return &GetUniverseMoonsMoonIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseMoonsMoonIDParamsWithTimeout creates a new GetUniverseMoonsMoonIDParams object
// with the ability to set a timeout on a request.
func NewGetUniverseMoonsMoonIDParamsWithTimeout(timeout time.Duration) *GetUniverseMoonsMoonIDParams {
	return &GetUniverseMoonsMoonIDParams{
		timeout: timeout,
	}
}

// NewGetUniverseMoonsMoonIDParamsWithContext creates a new GetUniverseMoonsMoonIDParams object
// with the ability to set a context for a request.
func NewGetUniverseMoonsMoonIDParamsWithContext(ctx context.Context) *GetUniverseMoonsMoonIDParams {
	return &GetUniverseMoonsMoonIDParams{
		Context: ctx,
	}
}

// NewGetUniverseMoonsMoonIDParamsWithHTTPClient creates a new GetUniverseMoonsMoonIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUniverseMoonsMoonIDParamsWithHTTPClient(client *http.Client) *GetUniverseMoonsMoonIDParams {
	return &GetUniverseMoonsMoonIDParams{
		HTTPClient: client,
	}
}

/*
GetUniverseMoonsMoonIDParams contains all the parameters to send to the API endpoint

	for the get universe moons moon id operation.

	Typically these are written to a http.Request.
*/
type GetUniverseMoonsMoonIDParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* MoonID.

	   moon_id integer

	   Format: int32
	*/
	MoonID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get universe moons moon id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseMoonsMoonIDParams) WithDefaults() *GetUniverseMoonsMoonIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get universe moons moon id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseMoonsMoonIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetUniverseMoonsMoonIDParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithTimeout(timeout time.Duration) *GetUniverseMoonsMoonIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithContext(ctx context.Context) *GetUniverseMoonsMoonIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithHTTPClient(client *http.Client) *GetUniverseMoonsMoonIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseMoonsMoonIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithDatasource(datasource *string) *GetUniverseMoonsMoonIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMoonID adds the moonID to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) WithMoonID(moonID int32) *GetUniverseMoonsMoonIDParams {
	o.SetMoonID(moonID)
	return o
}

// SetMoonID adds the moonId to the get universe moons moon id params
func (o *GetUniverseMoonsMoonIDParams) SetMoonID(moonID int32) {
	o.MoonID = moonID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseMoonsMoonIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param moon_id
	if err := r.SetPathParam("moon_id", swag.FormatInt32(o.MoonID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
