// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// NewGetCorporationsCorporationIDWalletsDivisionJournalParams creates a new GetCorporationsCorporationIDWalletsDivisionJournalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCorporationsCorporationIDWalletsDivisionJournalParams() *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	return &GetCorporationsCorporationIDWalletsDivisionJournalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalParamsWithTimeout creates a new GetCorporationsCorporationIDWalletsDivisionJournalParams object
// with the ability to set a timeout on a request.
func NewGetCorporationsCorporationIDWalletsDivisionJournalParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	return &GetCorporationsCorporationIDWalletsDivisionJournalParams{
		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalParamsWithContext creates a new GetCorporationsCorporationIDWalletsDivisionJournalParams object
// with the ability to set a context for a request.
func NewGetCorporationsCorporationIDWalletsDivisionJournalParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	return &GetCorporationsCorporationIDWalletsDivisionJournalParams{
		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalParamsWithHTTPClient creates a new GetCorporationsCorporationIDWalletsDivisionJournalParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCorporationsCorporationIDWalletsDivisionJournalParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	return &GetCorporationsCorporationIDWalletsDivisionJournalParams{
		HTTPClient: client,
	}
}

/*
GetCorporationsCorporationIDWalletsDivisionJournalParams contains all the parameters to send to the API endpoint

	for the get corporations corporation id wallets division journal operation.

	Typically these are written to a http.Request.
*/
type GetCorporationsCorporationIDWalletsDivisionJournalParams struct {

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

	/* Division.

	   Wallet key of the division to fetch journals from

	   Format: int32
	*/
	Division int32

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

// WithDefaults hydrates default values in the get corporations corporation id wallets division journal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithDefaults() *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get corporations corporation id wallets division journal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		pageDefault = int32(1)
	)

	val := GetCorporationsCorporationIDWalletsDivisionJournalParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithDivision adds the division to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithDivision(division int32) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetDivision(division)
	return o
}

// SetDivision adds the division to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetDivision(division int32) {
	o.Division = division
}

// WithPage adds the page to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithPage(page *int32) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WithToken(token *string) *GetCorporationsCorporationIDWalletsDivisionJournalParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id wallets division journal params
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDWalletsDivisionJournalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param division
	if err := r.SetPathParam("division", swag.FormatInt32(o.Division)); err != nil {
		return err
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
