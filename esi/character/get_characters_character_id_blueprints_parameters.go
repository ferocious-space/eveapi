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

// NewGetCharactersCharacterIDBlueprintsParams creates a new GetCharactersCharacterIDBlueprintsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDBlueprintsParams() *GetCharactersCharacterIDBlueprintsParams {
	return &GetCharactersCharacterIDBlueprintsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDBlueprintsParamsWithTimeout creates a new GetCharactersCharacterIDBlueprintsParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDBlueprintsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDBlueprintsParams {
	return &GetCharactersCharacterIDBlueprintsParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDBlueprintsParamsWithContext creates a new GetCharactersCharacterIDBlueprintsParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDBlueprintsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDBlueprintsParams {
	return &GetCharactersCharacterIDBlueprintsParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDBlueprintsParamsWithHTTPClient creates a new GetCharactersCharacterIDBlueprintsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDBlueprintsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDBlueprintsParams {
	return &GetCharactersCharacterIDBlueprintsParams{
		HTTPClient: client,
	}
}

/* GetCharactersCharacterIDBlueprintsParams contains all the parameters to send to the API endpoint
   for the get characters character id blueprints operation.

   Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDBlueprintsParams struct {

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

// WithDefaults hydrates default values in the get characters character id blueprints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDBlueprintsParams) WithDefaults() *GetCharactersCharacterIDBlueprintsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id blueprints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDBlueprintsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		pageDefault = int32(1)
	)

	val := GetCharactersCharacterIDBlueprintsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDBlueprintsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDBlueprintsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDBlueprintsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDBlueprintsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDBlueprintsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDBlueprintsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithPage(page *int32) *GetCharactersCharacterIDBlueprintsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) WithToken(token *string) *GetCharactersCharacterIDBlueprintsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id blueprints params
func (o *GetCharactersCharacterIDBlueprintsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDBlueprintsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
