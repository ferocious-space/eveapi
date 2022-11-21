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

// NewPostUniverseIdsParams creates a new PostUniverseIdsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUniverseIdsParams() *PostUniverseIdsParams {
	return &PostUniverseIdsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUniverseIdsParamsWithTimeout creates a new PostUniverseIdsParams object
// with the ability to set a timeout on a request.
func NewPostUniverseIdsParamsWithTimeout(timeout time.Duration) *PostUniverseIdsParams {
	return &PostUniverseIdsParams{
		timeout: timeout,
	}
}

// NewPostUniverseIdsParamsWithContext creates a new PostUniverseIdsParams object
// with the ability to set a context for a request.
func NewPostUniverseIdsParamsWithContext(ctx context.Context) *PostUniverseIdsParams {
	return &PostUniverseIdsParams{
		Context: ctx,
	}
}

// NewPostUniverseIdsParamsWithHTTPClient creates a new PostUniverseIdsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUniverseIdsParamsWithHTTPClient(client *http.Client) *PostUniverseIdsParams {
	return &PostUniverseIdsParams{
		HTTPClient: client,
	}
}

/*
PostUniverseIdsParams contains all the parameters to send to the API endpoint

	for the post universe ids operation.

	Typically these are written to a http.Request.
*/
type PostUniverseIdsParams struct {

	/* AcceptLanguage.

	   Language to use in the response

	   Default: "en"
	*/
	AcceptLanguage *string

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

	/* Names.

	   The names to resolve
	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post universe ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUniverseIdsParams) WithDefaults() *PostUniverseIdsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post universe ids params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUniverseIdsParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en")

		datasourceDefault = string("tranquility")

		languageDefault = string("en")
	)

	val := PostUniverseIdsParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post universe ids params
func (o *PostUniverseIdsParams) WithTimeout(timeout time.Duration) *PostUniverseIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post universe ids params
func (o *PostUniverseIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post universe ids params
func (o *PostUniverseIdsParams) WithContext(ctx context.Context) *PostUniverseIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post universe ids params
func (o *PostUniverseIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post universe ids params
func (o *PostUniverseIdsParams) WithHTTPClient(client *http.Client) *PostUniverseIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post universe ids params
func (o *PostUniverseIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the post universe ids params
func (o *PostUniverseIdsParams) WithAcceptLanguage(acceptLanguage *string) *PostUniverseIdsParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the post universe ids params
func (o *PostUniverseIdsParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithDatasource adds the datasource to the post universe ids params
func (o *PostUniverseIdsParams) WithDatasource(datasource *string) *PostUniverseIdsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post universe ids params
func (o *PostUniverseIdsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the post universe ids params
func (o *PostUniverseIdsParams) WithLanguage(language *string) *PostUniverseIdsParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the post universe ids params
func (o *PostUniverseIdsParams) SetLanguage(language *string) {
	o.Language = language
}

// WithNames adds the names to the post universe ids params
func (o *PostUniverseIdsParams) WithNames(names []string) *PostUniverseIdsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the post universe ids params
func (o *PostUniverseIdsParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *PostUniverseIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Names != nil {
		if err := r.SetBodyParam(o.Names); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
