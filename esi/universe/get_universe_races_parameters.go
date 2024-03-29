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

// NewGetUniverseRacesParams creates a new GetUniverseRacesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUniverseRacesParams() *GetUniverseRacesParams {
	return &GetUniverseRacesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseRacesParamsWithTimeout creates a new GetUniverseRacesParams object
// with the ability to set a timeout on a request.
func NewGetUniverseRacesParamsWithTimeout(timeout time.Duration) *GetUniverseRacesParams {
	return &GetUniverseRacesParams{
		timeout: timeout,
	}
}

// NewGetUniverseRacesParamsWithContext creates a new GetUniverseRacesParams object
// with the ability to set a context for a request.
func NewGetUniverseRacesParamsWithContext(ctx context.Context) *GetUniverseRacesParams {
	return &GetUniverseRacesParams{
		Context: ctx,
	}
}

// NewGetUniverseRacesParamsWithHTTPClient creates a new GetUniverseRacesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUniverseRacesParamsWithHTTPClient(client *http.Client) *GetUniverseRacesParams {
	return &GetUniverseRacesParams{
		HTTPClient: client,
	}
}

/*
GetUniverseRacesParams contains all the parameters to send to the API endpoint

	for the get universe races operation.

	Typically these are written to a http.Request.
*/
type GetUniverseRacesParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get universe races params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseRacesParams) WithDefaults() *GetUniverseRacesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get universe races params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUniverseRacesParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en")

		datasourceDefault = string("tranquility")

		languageDefault = string("en")
	)

	val := GetUniverseRacesParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get universe races params
func (o *GetUniverseRacesParams) WithTimeout(timeout time.Duration) *GetUniverseRacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe races params
func (o *GetUniverseRacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe races params
func (o *GetUniverseRacesParams) WithContext(ctx context.Context) *GetUniverseRacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe races params
func (o *GetUniverseRacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe races params
func (o *GetUniverseRacesParams) WithHTTPClient(client *http.Client) *GetUniverseRacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe races params
func (o *GetUniverseRacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get universe races params
func (o *GetUniverseRacesParams) WithAcceptLanguage(acceptLanguage *string) *GetUniverseRacesParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get universe races params
func (o *GetUniverseRacesParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe races params
func (o *GetUniverseRacesParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseRacesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe races params
func (o *GetUniverseRacesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get universe races params
func (o *GetUniverseRacesParams) WithDatasource(datasource *string) *GetUniverseRacesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe races params
func (o *GetUniverseRacesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get universe races params
func (o *GetUniverseRacesParams) WithLanguage(language *string) *GetUniverseRacesParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get universe races params
func (o *GetUniverseRacesParams) SetLanguage(language *string) {
	o.Language = language
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseRacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
