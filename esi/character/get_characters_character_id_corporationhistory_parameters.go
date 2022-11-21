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

// NewGetCharactersCharacterIDCorporationhistoryParams creates a new GetCharactersCharacterIDCorporationhistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDCorporationhistoryParams() *GetCharactersCharacterIDCorporationhistoryParams {
	return &GetCharactersCharacterIDCorporationhistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDCorporationhistoryParamsWithTimeout creates a new GetCharactersCharacterIDCorporationhistoryParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDCorporationhistoryParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDCorporationhistoryParams {
	return &GetCharactersCharacterIDCorporationhistoryParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDCorporationhistoryParamsWithContext creates a new GetCharactersCharacterIDCorporationhistoryParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDCorporationhistoryParamsWithContext(ctx context.Context) *GetCharactersCharacterIDCorporationhistoryParams {
	return &GetCharactersCharacterIDCorporationhistoryParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDCorporationhistoryParamsWithHTTPClient creates a new GetCharactersCharacterIDCorporationhistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDCorporationhistoryParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDCorporationhistoryParams {
	return &GetCharactersCharacterIDCorporationhistoryParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDCorporationhistoryParams contains all the parameters to send to the API endpoint

	for the get characters character id corporationhistory operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDCorporationhistoryParams struct {

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

// WithDefaults hydrates default values in the get characters character id corporationhistory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithDefaults() *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id corporationhistory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCharactersCharacterIDCorporationhistoryParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithContext(ctx context.Context) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithDatasource(datasource *string) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDCorporationhistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
