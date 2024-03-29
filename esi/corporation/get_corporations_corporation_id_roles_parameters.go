// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// NewGetCorporationsCorporationIDRolesParams creates a new GetCorporationsCorporationIDRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCorporationsCorporationIDRolesParams() *GetCorporationsCorporationIDRolesParams {
	return &GetCorporationsCorporationIDRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDRolesParamsWithTimeout creates a new GetCorporationsCorporationIDRolesParams object
// with the ability to set a timeout on a request.
func NewGetCorporationsCorporationIDRolesParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDRolesParams {
	return &GetCorporationsCorporationIDRolesParams{
		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDRolesParamsWithContext creates a new GetCorporationsCorporationIDRolesParams object
// with the ability to set a context for a request.
func NewGetCorporationsCorporationIDRolesParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDRolesParams {
	return &GetCorporationsCorporationIDRolesParams{
		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDRolesParamsWithHTTPClient creates a new GetCorporationsCorporationIDRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCorporationsCorporationIDRolesParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDRolesParams {
	return &GetCorporationsCorporationIDRolesParams{
		HTTPClient: client,
	}
}

/*
GetCorporationsCorporationIDRolesParams contains all the parameters to send to the API endpoint

	for the get corporations corporation id roles operation.

	Typically these are written to a http.Request.
*/
type GetCorporationsCorporationIDRolesParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* CorporationID.

	   An EVE corporation ID

	   Format: int32
	*/
	CorporationID int32

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get corporations corporation id roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDRolesParams) WithDefaults() *GetCorporationsCorporationIDRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get corporations corporation id roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDRolesParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCorporationsCorporationIDRolesParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDRolesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDRolesParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDRolesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) WithToken(token *string) *GetCorporationsCorporationIDRolesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id roles params
func (o *GetCorporationsCorporationIDRolesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
		return err
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

	if o.Token != nil {

		// query param token
		var qrToken string

		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {

			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
