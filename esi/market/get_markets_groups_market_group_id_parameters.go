/*
 *    Copyright 2021 FerociousBite and Contributors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// Code generated by go-swagger; DO NOT EDIT.

package market

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

// NewGetMarketsGroupsMarketGroupIDParams creates a new GetMarketsGroupsMarketGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMarketsGroupsMarketGroupIDParams() *GetMarketsGroupsMarketGroupIDParams {
	return &GetMarketsGroupsMarketGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketsGroupsMarketGroupIDParamsWithTimeout creates a new GetMarketsGroupsMarketGroupIDParams object
// with the ability to set a timeout on a request.
func NewGetMarketsGroupsMarketGroupIDParamsWithTimeout(timeout time.Duration) *GetMarketsGroupsMarketGroupIDParams {
	return &GetMarketsGroupsMarketGroupIDParams{
		timeout: timeout,
	}
}

// NewGetMarketsGroupsMarketGroupIDParamsWithContext creates a new GetMarketsGroupsMarketGroupIDParams object
// with the ability to set a context for a request.
func NewGetMarketsGroupsMarketGroupIDParamsWithContext(ctx context.Context) *GetMarketsGroupsMarketGroupIDParams {
	return &GetMarketsGroupsMarketGroupIDParams{
		Context: ctx,
	}
}

// NewGetMarketsGroupsMarketGroupIDParamsWithHTTPClient creates a new GetMarketsGroupsMarketGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMarketsGroupsMarketGroupIDParamsWithHTTPClient(client *http.Client) *GetMarketsGroupsMarketGroupIDParams {
	return &GetMarketsGroupsMarketGroupIDParams{
		HTTPClient: client,
	}
}

/* GetMarketsGroupsMarketGroupIDParams contains all the parameters to send to the API endpoint
   for the get markets groups market group id operation.

   Typically these are written to a http.Request.
*/
type GetMarketsGroupsMarketGroupIDParams struct {

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

	/* MarketGroupID.

	   An Eve item group ID

	   Format: int32
	*/
	MarketGroupID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get markets groups market group id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketsGroupsMarketGroupIDParams) WithDefaults() *GetMarketsGroupsMarketGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get markets groups market group id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketsGroupsMarketGroupIDParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en")

		datasourceDefault = string("tranquility")

		languageDefault = string("en")
	)

	val := GetMarketsGroupsMarketGroupIDParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithTimeout(timeout time.Duration) *GetMarketsGroupsMarketGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithContext(ctx context.Context) *GetMarketsGroupsMarketGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithHTTPClient(client *http.Client) *GetMarketsGroupsMarketGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithAcceptLanguage(acceptLanguage *string) *GetMarketsGroupsMarketGroupIDParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithIfNoneMatch adds the ifNoneMatch to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetMarketsGroupsMarketGroupIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithDatasource(datasource *string) *GetMarketsGroupsMarketGroupIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithLanguage(language *string) *GetMarketsGroupsMarketGroupIDParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetLanguage(language *string) {
	o.Language = language
}

// WithMarketGroupID adds the marketGroupID to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) WithMarketGroupID(marketGroupID int32) *GetMarketsGroupsMarketGroupIDParams {
	o.SetMarketGroupID(marketGroupID)
	return o
}

// SetMarketGroupID adds the marketGroupId to the get markets groups market group id params
func (o *GetMarketsGroupsMarketGroupIDParams) SetMarketGroupID(marketGroupID int32) {
	o.MarketGroupID = marketGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketsGroupsMarketGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param market_group_id
	if err := r.SetPathParam("market_group_id", swag.FormatInt32(o.MarketGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
