// Code generated by go-swagger; DO NOT EDIT.

package skills

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

// NewGetCharactersCharacterIDAttributesParams creates a new GetCharactersCharacterIDAttributesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDAttributesParams() *GetCharactersCharacterIDAttributesParams {
	return &GetCharactersCharacterIDAttributesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDAttributesParamsWithTimeout creates a new GetCharactersCharacterIDAttributesParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDAttributesParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDAttributesParams {
	return &GetCharactersCharacterIDAttributesParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDAttributesParamsWithContext creates a new GetCharactersCharacterIDAttributesParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDAttributesParamsWithContext(ctx context.Context) *GetCharactersCharacterIDAttributesParams {
	return &GetCharactersCharacterIDAttributesParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDAttributesParamsWithHTTPClient creates a new GetCharactersCharacterIDAttributesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDAttributesParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDAttributesParams {
	return &GetCharactersCharacterIDAttributesParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDAttributesParams contains all the parameters to send to the API endpoint

	for the get characters character id attributes operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDAttributesParams struct {

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

// WithDefaults hydrates default values in the get characters character id attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDAttributesParams) WithDefaults() *GetCharactersCharacterIDAttributesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id attributes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDAttributesParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCharactersCharacterIDAttributesParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) WithContext(ctx context.Context) *GetCharactersCharacterIDAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDAttributesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDAttributesParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) WithDatasource(datasource *string) *GetCharactersCharacterIDAttributesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) WithToken(token *string) *GetCharactersCharacterIDAttributesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id attributes params
func (o *GetCharactersCharacterIDAttributesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
