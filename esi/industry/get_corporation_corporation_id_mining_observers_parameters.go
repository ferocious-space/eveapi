// Code generated by go-swagger; DO NOT EDIT.

package industry

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

// NewGetCorporationCorporationIDMiningObserversParams creates a new GetCorporationCorporationIDMiningObserversParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCorporationCorporationIDMiningObserversParams() *GetCorporationCorporationIDMiningObserversParams {
	return &GetCorporationCorporationIDMiningObserversParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationCorporationIDMiningObserversParamsWithTimeout creates a new GetCorporationCorporationIDMiningObserversParams object
// with the ability to set a timeout on a request.
func NewGetCorporationCorporationIDMiningObserversParamsWithTimeout(timeout time.Duration) *GetCorporationCorporationIDMiningObserversParams {
	return &GetCorporationCorporationIDMiningObserversParams{
		timeout: timeout,
	}
}

// NewGetCorporationCorporationIDMiningObserversParamsWithContext creates a new GetCorporationCorporationIDMiningObserversParams object
// with the ability to set a context for a request.
func NewGetCorporationCorporationIDMiningObserversParamsWithContext(ctx context.Context) *GetCorporationCorporationIDMiningObserversParams {
	return &GetCorporationCorporationIDMiningObserversParams{
		Context: ctx,
	}
}

// NewGetCorporationCorporationIDMiningObserversParamsWithHTTPClient creates a new GetCorporationCorporationIDMiningObserversParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCorporationCorporationIDMiningObserversParamsWithHTTPClient(client *http.Client) *GetCorporationCorporationIDMiningObserversParams {
	return &GetCorporationCorporationIDMiningObserversParams{
		HTTPClient: client,
	}
}

/* GetCorporationCorporationIDMiningObserversParams contains all the parameters to send to the API endpoint
   for the get corporation corporation id mining observers operation.

   Typically these are written to a http.Request.
*/
type GetCorporationCorporationIDMiningObserversParams struct {

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

	/* Page.

	   Which page of results to return

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get corporation corporation id mining observers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationCorporationIDMiningObserversParams) WithDefaults() *GetCorporationCorporationIDMiningObserversParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get corporation corporation id mining observers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationCorporationIDMiningObserversParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		pageDefault = int32(1)
	)

	val := GetCorporationCorporationIDMiningObserversParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithTimeout(timeout time.Duration) *GetCorporationCorporationIDMiningObserversParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithContext(ctx context.Context) *GetCorporationCorporationIDMiningObserversParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithHTTPClient(client *http.Client) *GetCorporationCorporationIDMiningObserversParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationCorporationIDMiningObserversParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithCorporationID(corporationID int32) *GetCorporationCorporationIDMiningObserversParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithDatasource(datasource *string) *GetCorporationCorporationIDMiningObserversParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithPage(page *int32) *GetCorporationCorporationIDMiningObserversParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) WithToken(token *string) *GetCorporationCorporationIDMiningObserversParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporation corporation id mining observers params
func (o *GetCorporationCorporationIDMiningObserversParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationCorporationIDMiningObserversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
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
