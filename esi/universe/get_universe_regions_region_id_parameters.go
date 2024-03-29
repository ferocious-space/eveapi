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

// NewGetUniverseRegionsRegionIDParams creates a new GetUniverseRegionsRegionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUniverseRegionsRegionIDParams() *GetUniverseRegionsRegionIDParams {
	return &GetUniverseRegionsRegionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseRegionsRegionIDParamsWithTimeout creates a new GetUniverseRegionsRegionIDParams object
// with the ability to set a timeout on a request.
func NewGetUniverseRegionsRegionIDParamsWithTimeout(timeout time.Duration) *GetUniverseRegionsRegionIDParams {
	return &GetUniverseRegionsRegionIDParams{
		timeout: timeout,
	}
}

// NewGetUniverseRegionsRegionIDParamsWithContext creates a new GetUniverseRegionsRegionIDParams object
// with the ability to set a context for a request.
func NewGetUniverseRegionsRegionIDParamsWithContext(ctx context.Context) *GetUniverseRegionsRegionIDParams {
	return &GetUniverseRegionsRegionIDParams{
		Context: ctx,
	}
}

// NewGetUniverseRegionsRegionIDParamsWithHTTPClient creates a new GetUniverseRegionsRegionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUniverseRegionsRegionIDParamsWithHTTPClient(client *http.Client) *GetUniverseRegionsRegionIDParams {
	return &GetUniverseRegionsRegionIDParams{
		HTTPClient: client,
	}
}

/*
GetUniverseRegionsRegionIDParams contains all the parameters to send to the API endpoint

	for the get universe regions region id operation.

	Typically these are written to a http.Request.
*/
type GetUniverseRegionsRegionIDParams struct {

	/* AcceptLanguage.

	   Language to use in the response

	   Default: "en"
	*/
	AcceptLanguage *string

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* Language.

	   Language to use in the response, takes precedence over Accept-Language

	   Default: "en"
	*/
	Language *string

	/* RegionID.

	   region_id integer

	   Format: int32
	*/
	RegionID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get universe regions region id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseRegionsRegionIDParams) WithDefaults() *GetUniverseRegionsRegionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get universe regions region id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseRegionsRegionIDParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en")

		datasourceDefault = string("tranquility")

		languageDefault = string("en")
	)

	val := GetUniverseRegionsRegionIDParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithTimeout(timeout time.Duration) *GetUniverseRegionsRegionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithContext(ctx context.Context) *GetUniverseRegionsRegionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithHTTPClient(client *http.Client) *GetUniverseRegionsRegionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithAcceptLanguage(acceptLanguage *string) *GetUniverseRegionsRegionIDParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseRegionsRegionIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithDatasource(datasource *string) *GetUniverseRegionsRegionIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithLanguage(language *string) *GetUniverseRegionsRegionIDParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetLanguage(language *string) {
	o.Language = language
}

// WithRegionID adds the regionID to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) WithRegionID(regionID int32) *GetUniverseRegionsRegionIDParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the get universe regions region id params
func (o *GetUniverseRegionsRegionIDParams) SetRegionID(regionID int32) {
	o.RegionID = regionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseRegionsRegionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}
	}

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

	if o.Language != nil {

		// query param language
		var qrLanguage string

		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {

			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}
	}

	// path param region_id
	if err := r.SetPathParam("region_id", swag.FormatInt32(o.RegionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
