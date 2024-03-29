// Code generated by go-swagger; DO NOT EDIT.

package mail

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

// NewGetCharactersCharacterIDMailListsParams creates a new GetCharactersCharacterIDMailListsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDMailListsParams() *GetCharactersCharacterIDMailListsParams {
	return &GetCharactersCharacterIDMailListsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDMailListsParamsWithTimeout creates a new GetCharactersCharacterIDMailListsParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDMailListsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailListsParams {
	return &GetCharactersCharacterIDMailListsParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDMailListsParamsWithContext creates a new GetCharactersCharacterIDMailListsParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDMailListsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDMailListsParams {
	return &GetCharactersCharacterIDMailListsParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDMailListsParamsWithHTTPClient creates a new GetCharactersCharacterIDMailListsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDMailListsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDMailListsParams {
	return &GetCharactersCharacterIDMailListsParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDMailListsParams contains all the parameters to send to the API endpoint

	for the get characters character id mail lists operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDMailListsParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* CharacterID.

	   An EVE character ID

	   Format: int32
	*/
	CharacterID int32

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

// WithDefaults hydrates default values in the get characters character id mail lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDMailListsParams) WithDefaults() *GetCharactersCharacterIDMailListsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id mail lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDMailListsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCharactersCharacterIDMailListsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDMailListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDMailListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDMailListsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDMailListsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDMailListsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) WithToken(token *string) *GetCharactersCharacterIDMailListsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id mail lists params
func (o *GetCharactersCharacterIDMailListsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDMailListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
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
