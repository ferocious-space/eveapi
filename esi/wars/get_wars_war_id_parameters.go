// Code generated by go-swagger; DO NOT EDIT.

package wars

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

// NewGetWarsWarIDParams creates a new GetWarsWarIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWarsWarIDParams() *GetWarsWarIDParams {
	return &GetWarsWarIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWarsWarIDParamsWithTimeout creates a new GetWarsWarIDParams object
// with the ability to set a timeout on a request.
func NewGetWarsWarIDParamsWithTimeout(timeout time.Duration) *GetWarsWarIDParams {
	return &GetWarsWarIDParams{
		timeout: timeout,
	}
}

// NewGetWarsWarIDParamsWithContext creates a new GetWarsWarIDParams object
// with the ability to set a context for a request.
func NewGetWarsWarIDParamsWithContext(ctx context.Context) *GetWarsWarIDParams {
	return &GetWarsWarIDParams{
		Context: ctx,
	}
}

// NewGetWarsWarIDParamsWithHTTPClient creates a new GetWarsWarIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWarsWarIDParamsWithHTTPClient(client *http.Client) *GetWarsWarIDParams {
	return &GetWarsWarIDParams{
		HTTPClient: client,
	}
}

/*
GetWarsWarIDParams contains all the parameters to send to the API endpoint

	for the get wars war id operation.

	Typically these are written to a http.Request.
*/
type GetWarsWarIDParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* WarID.

	   ID for a war

	   Format: int32
	*/
	WarID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wars war id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWarsWarIDParams) WithDefaults() *GetWarsWarIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wars war id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWarsWarIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetWarsWarIDParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get wars war id params
func (o *GetWarsWarIDParams) WithTimeout(timeout time.Duration) *GetWarsWarIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wars war id params
func (o *GetWarsWarIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wars war id params
func (o *GetWarsWarIDParams) WithContext(ctx context.Context) *GetWarsWarIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wars war id params
func (o *GetWarsWarIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wars war id params
func (o *GetWarsWarIDParams) WithHTTPClient(client *http.Client) *GetWarsWarIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wars war id params
func (o *GetWarsWarIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get wars war id params
func (o *GetWarsWarIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetWarsWarIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get wars war id params
func (o *GetWarsWarIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get wars war id params
func (o *GetWarsWarIDParams) WithDatasource(datasource *string) *GetWarsWarIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get wars war id params
func (o *GetWarsWarIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithWarID adds the warID to the get wars war id params
func (o *GetWarsWarIDParams) WithWarID(warID int32) *GetWarsWarIDParams {
	o.SetWarID(warID)
	return o
}

// SetWarID adds the warId to the get wars war id params
func (o *GetWarsWarIDParams) SetWarID(warID int32) {
	o.WarID = warID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWarsWarIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param war_id
	if err := r.SetPathParam("war_id", swag.FormatInt32(o.WarID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
