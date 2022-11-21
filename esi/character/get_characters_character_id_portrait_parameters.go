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

// NewGetCharactersCharacterIDPortraitParams creates a new GetCharactersCharacterIDPortraitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDPortraitParams() *GetCharactersCharacterIDPortraitParams {
	return &GetCharactersCharacterIDPortraitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDPortraitParamsWithTimeout creates a new GetCharactersCharacterIDPortraitParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDPortraitParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDPortraitParams {
	return &GetCharactersCharacterIDPortraitParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDPortraitParamsWithContext creates a new GetCharactersCharacterIDPortraitParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDPortraitParamsWithContext(ctx context.Context) *GetCharactersCharacterIDPortraitParams {
	return &GetCharactersCharacterIDPortraitParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDPortraitParamsWithHTTPClient creates a new GetCharactersCharacterIDPortraitParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDPortraitParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDPortraitParams {
	return &GetCharactersCharacterIDPortraitParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDPortraitParams contains all the parameters to send to the API endpoint

	for the get characters character id portrait operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDPortraitParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get characters character id portrait params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDPortraitParams) WithDefaults() *GetCharactersCharacterIDPortraitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id portrait params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDPortraitParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCharactersCharacterIDPortraitParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDPortraitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) WithContext(ctx context.Context) *GetCharactersCharacterIDPortraitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDPortraitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDPortraitParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDPortraitParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) WithDatasource(datasource *string) *GetCharactersCharacterIDPortraitParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id portrait params
func (o *GetCharactersCharacterIDPortraitParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDPortraitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
