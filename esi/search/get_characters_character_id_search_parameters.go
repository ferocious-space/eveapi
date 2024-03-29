// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewGetCharactersCharacterIDSearchParams creates a new GetCharactersCharacterIDSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDSearchParams() *GetCharactersCharacterIDSearchParams {
	return &GetCharactersCharacterIDSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDSearchParamsWithTimeout creates a new GetCharactersCharacterIDSearchParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDSearchParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDSearchParams {
	return &GetCharactersCharacterIDSearchParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDSearchParamsWithContext creates a new GetCharactersCharacterIDSearchParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDSearchParamsWithContext(ctx context.Context) *GetCharactersCharacterIDSearchParams {
	return &GetCharactersCharacterIDSearchParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDSearchParamsWithHTTPClient creates a new GetCharactersCharacterIDSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDSearchParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDSearchParams {
	return &GetCharactersCharacterIDSearchParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDSearchParams contains all the parameters to send to the API endpoint

	for the get characters character id search operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDSearchParams struct {

	/* AcceptLanguage.

	   Language to use in the response

	   Default: "en"
	*/
	AcceptLanguage *string

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Categories.

	   Type of entities to search for
	*/
	Categories []string

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

	/* Language.

	   Language to use in the response, takes precedence over Accept-Language

	   Default: "en"
	*/
	Language *string

	/* Search.

	   The string to search on
	*/
	Search string

	/* Strict.

	   Whether the search should be a strict match
	*/
	Strict *bool

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get characters character id search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDSearchParams) WithDefaults() *GetCharactersCharacterIDSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDSearchParams) SetDefaults() {
	var (
		acceptLanguageDefault = string("en")

		datasourceDefault = string("tranquility")

		languageDefault = string("en")

		strictDefault = bool(false)
	)

	val := GetCharactersCharacterIDSearchParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
		Strict:         &strictDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithContext(ctx context.Context) *GetCharactersCharacterIDSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithAcceptLanguage(acceptLanguage *string) *GetCharactersCharacterIDSearchParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDSearchParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCategories adds the categories to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithCategories(categories []string) *GetCharactersCharacterIDSearchParams {
	o.SetCategories(categories)
	return o
}

// SetCategories adds the categories to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetCategories(categories []string) {
	o.Categories = categories
}

// WithCharacterID adds the characterID to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDSearchParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithDatasource(datasource *string) *GetCharactersCharacterIDSearchParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithLanguage(language *string) *GetCharactersCharacterIDSearchParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetLanguage(language *string) {
	o.Language = language
}

// WithSearch adds the search to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithSearch(search string) *GetCharactersCharacterIDSearchParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetSearch(search string) {
	o.Search = search
}

// WithStrict adds the strict to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithStrict(strict *bool) *GetCharactersCharacterIDSearchParams {
	o.SetStrict(strict)
	return o
}

// SetStrict adds the strict to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetStrict(strict *bool) {
	o.Strict = strict
}

// WithToken adds the token to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) WithToken(token *string) *GetCharactersCharacterIDSearchParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id search params
func (o *GetCharactersCharacterIDSearchParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}
	}

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}
	}

	if o.Categories != nil {

		// binding items for categories
		joinedCategories := o.bindParamCategories(reg)

		// query array param categories
		if err := r.SetQueryParam("categories", joinedCategories...); err != nil {
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

	if o.Language != nil {

		// query param language
		var qrLanguage string

		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {

			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}
	}

	// query param search
	qrSearch := o.Search
	qSearch := qrSearch
	if qSearch != "" {

		if err := r.SetQueryParam("search", qSearch); err != nil {
			return err
		}
	}

	if o.Strict != nil {

		// query param strict
		var qrStrict bool

		if o.Strict != nil {
			qrStrict = *o.Strict
		}
		qStrict := swag.FormatBool(qrStrict)
		if qStrict != "" {

			if err := r.SetQueryParam("strict", qStrict); err != nil {
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

// bindParamGetCharactersCharacterIDSearch binds the parameter categories
func (o *GetCharactersCharacterIDSearchParams) bindParamCategories(formats strfmt.Registry) []string {
	categoriesIR := o.Categories

	var categoriesIC []string
	for _, categoriesIIR := range categoriesIR { // explode []string

		categoriesIIV := categoriesIIR // string as string
		categoriesIC = append(categoriesIC, categoriesIIV)
	}

	// items.CollectionFormat: ""
	categoriesIS := swag.JoinByFormat(categoriesIC, "")

	return categoriesIS
}
