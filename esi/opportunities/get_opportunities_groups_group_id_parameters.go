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

package opportunities

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

// NewGetOpportunitiesGroupsGroupIDParams creates a new GetOpportunitiesGroupsGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOpportunitiesGroupsGroupIDParams() *GetOpportunitiesGroupsGroupIDParams {
	return &GetOpportunitiesGroupsGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOpportunitiesGroupsGroupIDParamsWithTimeout creates a new GetOpportunitiesGroupsGroupIDParams object
// with the ability to set a timeout on a request.
func NewGetOpportunitiesGroupsGroupIDParamsWithTimeout(timeout time.Duration) *GetOpportunitiesGroupsGroupIDParams {
	return &GetOpportunitiesGroupsGroupIDParams{
		timeout: timeout,
	}
}

// NewGetOpportunitiesGroupsGroupIDParamsWithContext creates a new GetOpportunitiesGroupsGroupIDParams object
// with the ability to set a context for a request.
func NewGetOpportunitiesGroupsGroupIDParamsWithContext(ctx context.Context) *GetOpportunitiesGroupsGroupIDParams {
	return &GetOpportunitiesGroupsGroupIDParams{
		Context: ctx,
	}
}

// NewGetOpportunitiesGroupsGroupIDParamsWithHTTPClient creates a new GetOpportunitiesGroupsGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOpportunitiesGroupsGroupIDParamsWithHTTPClient(client *http.Client) *GetOpportunitiesGroupsGroupIDParams {
	return &GetOpportunitiesGroupsGroupIDParams{
		HTTPClient: client,
	}
}

/* GetOpportunitiesGroupsGroupIDParams contains all the parameters to send to the API endpoint
   for the get opportunities groups group id operation.

   Typically these are written to a http.Request.
*/
type GetOpportunitiesGroupsGroupIDParams struct {

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

	/* GroupID.

	   ID of an opportunities group

	   Format: int32
	*/
	GroupID int32

	/* Language.

	   Language to use in the response, takes precedence over Accept-Language

	   Default: "en"
	*/
	Language *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get opportunities groups group id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOpportunitiesGroupsGroupIDParams) WithDefaults() *GetOpportunitiesGroupsGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get opportunities groups group id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOpportunitiesGroupsGroupIDParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en")

		datasourceDefault = string("tranquility")

		languageDefault = string("en")
	)

	val := GetOpportunitiesGroupsGroupIDParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithTimeout(timeout time.Duration) *GetOpportunitiesGroupsGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithContext(ctx context.Context) *GetOpportunitiesGroupsGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithHTTPClient(client *http.Client) *GetOpportunitiesGroupsGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithAcceptLanguage(acceptLanguage *string) *GetOpportunitiesGroupsGroupIDParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithIfNoneMatch adds the ifNoneMatch to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetOpportunitiesGroupsGroupIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithDatasource(datasource *string) *GetOpportunitiesGroupsGroupIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithGroupID adds the groupID to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithGroupID(groupID int32) *GetOpportunitiesGroupsGroupIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetGroupID(groupID int32) {
	o.GroupID = groupID
}

// WithLanguage adds the language to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) WithLanguage(language *string) *GetOpportunitiesGroupsGroupIDParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get opportunities groups group id params
func (o *GetOpportunitiesGroupsGroupIDParams) SetLanguage(language *string) {
	o.Language = language
}

// WriteToRequest writes these params to a swagger request
func (o *GetOpportunitiesGroupsGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt32(o.GroupID)); err != nil {
		return err
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
