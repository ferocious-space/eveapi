// Code generated by go-swagger; DO NOT EDIT.

package character

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

// NewGetCharactersCharacterIDNotificationsParams creates a new GetCharactersCharacterIDNotificationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDNotificationsParams() *GetCharactersCharacterIDNotificationsParams {
	return &GetCharactersCharacterIDNotificationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDNotificationsParamsWithTimeout creates a new GetCharactersCharacterIDNotificationsParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDNotificationsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDNotificationsParams {
	return &GetCharactersCharacterIDNotificationsParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDNotificationsParamsWithContext creates a new GetCharactersCharacterIDNotificationsParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDNotificationsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDNotificationsParams {
	return &GetCharactersCharacterIDNotificationsParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDNotificationsParamsWithHTTPClient creates a new GetCharactersCharacterIDNotificationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDNotificationsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDNotificationsParams {
	return &GetCharactersCharacterIDNotificationsParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDNotificationsParams contains all the parameters to send to the API endpoint

	for the get characters character id notifications operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDNotificationsParams struct {

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

// WithDefaults hydrates default values in the get characters character id notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDNotificationsParams) WithDefaults() *GetCharactersCharacterIDNotificationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDNotificationsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCharactersCharacterIDNotificationsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDNotificationsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDNotificationsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDNotificationsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) WithToken(token *string) *GetCharactersCharacterIDNotificationsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id notifications params
func (o *GetCharactersCharacterIDNotificationsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
