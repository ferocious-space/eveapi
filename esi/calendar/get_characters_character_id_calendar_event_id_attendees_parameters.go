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

// NewGetCharactersCharacterIDCalendarEventIDAttendeesParams creates a new GetCharactersCharacterIDCalendarEventIDAttendeesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCharactersCharacterIDCalendarEventIDAttendeesParams() *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	return &GetCharactersCharacterIDCalendarEventIDAttendeesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDAttendeesParamsWithTimeout creates a new GetCharactersCharacterIDCalendarEventIDAttendeesParams object
// with the ability to set a timeout on a request.
func NewGetCharactersCharacterIDCalendarEventIDAttendeesParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	return &GetCharactersCharacterIDCalendarEventIDAttendeesParams{
		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDAttendeesParamsWithContext creates a new GetCharactersCharacterIDCalendarEventIDAttendeesParams object
// with the ability to set a context for a request.
func NewGetCharactersCharacterIDCalendarEventIDAttendeesParamsWithContext(ctx context.Context) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	return &GetCharactersCharacterIDCalendarEventIDAttendeesParams{
		Context: ctx,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDAttendeesParamsWithHTTPClient creates a new GetCharactersCharacterIDCalendarEventIDAttendeesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCharactersCharacterIDCalendarEventIDAttendeesParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	return &GetCharactersCharacterIDCalendarEventIDAttendeesParams{
		HTTPClient: client,
	}
}

/*
GetCharactersCharacterIDCalendarEventIDAttendeesParams contains all the parameters to send to the API endpoint

	for the get characters character id calendar event id attendees operation.

	Typically these are written to a http.Request.
*/
type GetCharactersCharacterIDCalendarEventIDAttendeesParams struct {

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

// WithDefaults hydrates default values in the get characters character id calendar event id attendees params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithDefaults() *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get characters character id calendar event id attendees params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetCharactersCharacterIDCalendarEventIDAttendeesParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithContext(ctx context.Context) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithDatasource(datasource *string) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithEventID adds the eventID to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithEventID(eventID int32) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetEventID(eventID int32) {
	o.EventID = eventID
}

// WithToken adds the token to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WithToken(token *string) *GetCharactersCharacterIDCalendarEventIDAttendeesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id calendar event id attendees params
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDCalendarEventIDAttendeesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
