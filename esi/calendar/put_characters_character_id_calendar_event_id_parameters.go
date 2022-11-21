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

// NewPutCharactersCharacterIDCalendarEventIDParams creates a new PutCharactersCharacterIDCalendarEventIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCharactersCharacterIDCalendarEventIDParams() *PutCharactersCharacterIDCalendarEventIDParams {
	return &PutCharactersCharacterIDCalendarEventIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCharactersCharacterIDCalendarEventIDParamsWithTimeout creates a new PutCharactersCharacterIDCalendarEventIDParams object
// with the ability to set a timeout on a request.
func NewPutCharactersCharacterIDCalendarEventIDParamsWithTimeout(timeout time.Duration) *PutCharactersCharacterIDCalendarEventIDParams {
	return &PutCharactersCharacterIDCalendarEventIDParams{
		timeout: timeout,
	}
}

// NewPutCharactersCharacterIDCalendarEventIDParamsWithContext creates a new PutCharactersCharacterIDCalendarEventIDParams object
// with the ability to set a context for a request.
func NewPutCharactersCharacterIDCalendarEventIDParamsWithContext(ctx context.Context) *PutCharactersCharacterIDCalendarEventIDParams {
	return &PutCharactersCharacterIDCalendarEventIDParams{
		Context: ctx,
	}
}

// NewPutCharactersCharacterIDCalendarEventIDParamsWithHTTPClient creates a new PutCharactersCharacterIDCalendarEventIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCharactersCharacterIDCalendarEventIDParamsWithHTTPClient(client *http.Client) *PutCharactersCharacterIDCalendarEventIDParams {
	return &PutCharactersCharacterIDCalendarEventIDParams{
		HTTPClient: client,
	}
}

/*
PutCharactersCharacterIDCalendarEventIDParams contains all the parameters to send to the API endpoint

	for the put characters character id calendar event id operation.

	Typically these are written to a http.Request.
*/
type PutCharactersCharacterIDCalendarEventIDParams struct {

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

	   The ID of the event requested

	   Format: int32
	*/
	EventID int32

	/* Response.

	   The response value to set, overriding current value
	*/
	Response PutCharactersCharacterIDCalendarEventIDBody

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put characters character id calendar event id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithDefaults() *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put characters character id calendar event id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := PutCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithTimeout(timeout time.Duration) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithContext(ctx context.Context) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithHTTPClient(client *http.Client) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithCharacterID(characterID int32) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithDatasource(datasource *string) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithEventID adds the eventID to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithEventID(eventID int32) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetEventID(eventID int32) {
	o.EventID = eventID
}

// WithResponse adds the response to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithResponse(response PutCharactersCharacterIDCalendarEventIDBody) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetResponse(response)
	return o
}

// SetResponse adds the response to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetResponse(response PutCharactersCharacterIDCalendarEventIDBody) {
	o.Response = response
}

// WithToken adds the token to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) WithToken(token *string) *PutCharactersCharacterIDCalendarEventIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the put characters character id calendar event id params
func (o *PutCharactersCharacterIDCalendarEventIDParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PutCharactersCharacterIDCalendarEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
	if err := r.SetBodyParam(o.Response); err != nil {
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
