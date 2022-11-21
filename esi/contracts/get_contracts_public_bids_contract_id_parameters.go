// Code generated by go-swagger; DO NOT EDIT.

package contracts

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

// NewGetContractsPublicBidsContractIDParams creates a new GetContractsPublicBidsContractIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContractsPublicBidsContractIDParams() *GetContractsPublicBidsContractIDParams {
	return &GetContractsPublicBidsContractIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContractsPublicBidsContractIDParamsWithTimeout creates a new GetContractsPublicBidsContractIDParams object
// with the ability to set a timeout on a request.
func NewGetContractsPublicBidsContractIDParamsWithTimeout(timeout time.Duration) *GetContractsPublicBidsContractIDParams {
	return &GetContractsPublicBidsContractIDParams{
		timeout: timeout,
	}
}

// NewGetContractsPublicBidsContractIDParamsWithContext creates a new GetContractsPublicBidsContractIDParams object
// with the ability to set a context for a request.
func NewGetContractsPublicBidsContractIDParamsWithContext(ctx context.Context) *GetContractsPublicBidsContractIDParams {
	return &GetContractsPublicBidsContractIDParams{
		Context: ctx,
	}
}

// NewGetContractsPublicBidsContractIDParamsWithHTTPClient creates a new GetContractsPublicBidsContractIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContractsPublicBidsContractIDParamsWithHTTPClient(client *http.Client) *GetContractsPublicBidsContractIDParams {
	return &GetContractsPublicBidsContractIDParams{
		HTTPClient: client,
	}
}

/*
GetContractsPublicBidsContractIDParams contains all the parameters to send to the API endpoint

	for the get contracts public bids contract id operation.

	Typically these are written to a http.Request.
*/
type GetContractsPublicBidsContractIDParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* ContractID.

	   ID of a contract

	   Format: int32
	*/
	ContractID int32

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contracts public bids contract id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContractsPublicBidsContractIDParams) WithDefaults() *GetContractsPublicBidsContractIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contracts public bids contract id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContractsPublicBidsContractIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		pageDefault = int32(1)
	)

	val := GetContractsPublicBidsContractIDParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) WithTimeout(timeout time.Duration) *GetContractsPublicBidsContractIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) WithContext(ctx context.Context) *GetContractsPublicBidsContractIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) WithHTTPClient(client *http.Client) *GetContractsPublicBidsContractIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetContractsPublicBidsContractIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithContractID adds the contractID to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) WithContractID(contractID int32) *GetContractsPublicBidsContractIDParams {
	o.SetContractID(contractID)
	return o
}

// SetContractID adds the contractId to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) SetContractID(contractID int32) {
	o.ContractID = contractID
}

// WithDatasource adds the datasource to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) WithDatasource(datasource *string) *GetContractsPublicBidsContractIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) WithPage(page *int32) *GetContractsPublicBidsContractIDParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get contracts public bids contract id params
func (o *GetContractsPublicBidsContractIDParams) SetPage(page *int32) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetContractsPublicBidsContractIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param contract_id
	if err := r.SetPathParam("contract_id", swag.FormatInt32(o.ContractID)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
