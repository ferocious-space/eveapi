// Code generated by go-swagger; DO NOT EDIT.

package killmails

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

// NewGetCharactersCharacterIDKillmailsRecentParams creates a new GetCharactersCharacterIDKillmailsRecentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDKillmailsRecentParams() *GetCharactersCharacterIDKillmailsRecentParams {
	return &GetCharactersCharacterIDKillmailsRecentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDKillmailsRecentParamsWithTimeout creates a new GetCharactersCharacterIDKillmailsRecentParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDKillmailsRecentParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDKillmailsRecentParams {
	return &GetCharactersCharacterIDKillmailsRecentParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDKillmailsRecentParamsWithContext creates a new GetCharactersCharacterIDKillmailsRecentParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDKillmailsRecentParamsWithContext(ctx context.Context) *GetCharactersCharacterIDKillmailsRecentParams {
	return &GetCharactersCharacterIDKillmailsRecentParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDKillmailsRecentParamsWithHTTPClient creates a new GetCharactersCharacterIDKillmailsRecentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDKillmailsRecentParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDKillmailsRecentParams {
	return &GetCharactersCharacterIDKillmailsRecentParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDKillmailsRecentParams contains all the parameters to send to the API endpoint

	for the get characters character id killmails recent operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDKillmailsRecentParams struct {

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

// WithDefaults hydrates default values in the get characters character id killmails recent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithDefaults() *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id killmails recent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		pageDefault = int32(1)
	)

	val := GetCharactersCharacterIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithContext(ctx context.Context) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithDatasource(datasource *string) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithPage(page *int32) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithToken(token *string) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDKillmailsRecentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
