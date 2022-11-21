// Code generated by go-swagger; DO NOT EDIT.

package user_interface

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

// NewPostUIOpenwindowContractParams creates a new PostUIOpenwindowContractParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostUIOpenwindowContractParams() *PostUIOpenwindowContractParams {
	return &PostUIOpenwindowContractParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostUIOpenwindowContractParamsWithTimeout creates a new PostUIOpenwindowContractParams object
// with the ability to set a timeout on a request.
func NewPostUIOpenwindowContractParamsWithTimeout(timeout time.Duration) *PostUIOpenwindowContractParams {
	return &PostUIOpenwindowContractParams{
		timeout: timeout,
	}
}

// NewPostUIOpenwindowContractParamsWithContext creates a new PostUIOpenwindowContractParams object
// with the ability to set a context for a request.
func NewPostUIOpenwindowContractParamsWithContext(ctx context.Context) *PostUIOpenwindowContractParams {
	return &PostUIOpenwindowContractParams{
		Context: ctx,
	}
}

// NewPostUIOpenwindowContractParamsWithHTTPClient creates a new PostUIOpenwindowContractParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostUIOpenwindowContractParamsWithHTTPClient(client *http.Client) *PostUIOpenwindowContractParams {
	return &PostUIOpenwindowContractParams{
		HTTPClient: client,
	}
}

/*
PostUIOpenwindowContractParams contains all the parameters to send to the API endpoint

	for the post ui openwindow contract operation.

	Typically these are written to a http.Request.
*/
type PostUIOpenwindowContractParams struct {

	/* ContractID.

	   The contract to open

	   Format: int32
	*/
	ContractID int32

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

// WithDefaults hydrates default values in the post ui openwindow contract params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUIOpenwindowContractParams) WithDefaults() *PostUIOpenwindowContractParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post ui openwindow contract params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostUIOpenwindowContractParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := PostUIOpenwindowContractParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) WithTimeout(timeout time.Duration) *PostUIOpenwindowContractParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) WithContext(ctx context.Context) *PostUIOpenwindowContractParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) WithHTTPClient(client *http.Client) *PostUIOpenwindowContractParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractID adds the contractID to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) WithContractID(contractID int32) *PostUIOpenwindowContractParams {
	o.SetContractID(contractID)
	return o
}

// SetContractID adds the contractId to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) SetContractID(contractID int32) {
	o.ContractID = contractID
}

// WithDatasource adds the datasource to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) WithDatasource(datasource *string) *PostUIOpenwindowContractParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) WithToken(token *string) *PostUIOpenwindowContractParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post ui openwindow contract params
func (o *PostUIOpenwindowContractParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PostUIOpenwindowContractParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param contract_id
	qrContractID := o.ContractID
	qContractID := swag.FormatInt32(qrContractID)
	if qContractID != "" {

		if err := r.SetQueryParam("contract_id", qContractID); err != nil {
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
