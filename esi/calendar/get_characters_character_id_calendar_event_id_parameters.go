// Code generated by go-swagger; DO NOT EDIT.

package calendar

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

// NewGetCharactersCharacterIDCalendarEventIDParams creates a new GetCharactersCharacterIDCalendarEventIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDCalendarEventIDParams() *GetCharactersCharacterIDCalendarEventIDParams {
	return &GetCharactersCharacterIDCalendarEventIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDParamsWithTimeout creates a new GetCharactersCharacterIDCalendarEventIDParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDCalendarEventIDParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDCalendarEventIDParams {
	return &GetCharactersCharacterIDCalendarEventIDParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDParamsWithContext creates a new GetCharactersCharacterIDCalendarEventIDParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDCalendarEventIDParamsWithContext(ctx context.Context) *GetCharactersCharacterIDCalendarEventIDParams {
	return &GetCharactersCharacterIDCalendarEventIDParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDParamsWithHTTPClient creates a new GetCharactersCharacterIDCalendarEventIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDCalendarEventIDParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDCalendarEventIDParams {
	return &GetCharactersCharacterIDCalendarEventIDParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDCalendarEventIDParams contains all the parameters to send to the API endpoint

	for the get characters character id calendar event id operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDCalendarEventIDParams struct {

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

	/* EventID.

	   The id of the event requested

	   Format: int32
	*/
	EventID int32

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get characters character id calendar event id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithDefaults() *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id calendar event id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithContext(ctx context.Context) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithDatasource(datasource *string) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithEventID adds the eventID to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithEventID(eventID int32) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetEventID(eventID int32) {
	o.EventID = eventID
}

// WithToken adds the token to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithToken(token *string) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDCalendarEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param event_id
	if err := r.SetPathParam("event_id", swag.FormatInt32(o.EventID)); err != nil {
		return err
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
