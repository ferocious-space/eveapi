// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// NewPostCharactersCharacterIDAssetsLocationsParams creates a new PostCharactersCharacterIDAssetsLocationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostCharactersCharacterIDAssetsLocationsParams() *PostCharactersCharacterIDAssetsLocationsParams {
	return &PostCharactersCharacterIDAssetsLocationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostCharactersCharacterIDAssetsLocationsParamsWithTimeout creates a new PostCharactersCharacterIDAssetsLocationsParams object
// with the ability to set a timeout on a request.
func NewPostCharactersCharacterIDAssetsLocationsParamsWithTimeout(timeout time.Duration) *PostCharactersCharacterIDAssetsLocationsParams {
	return &PostCharactersCharacterIDAssetsLocationsParams{
		timeout: timeout,
	}
}

// NewPostCharactersCharacterIDAssetsLocationsParamsWithContext creates a new PostCharactersCharacterIDAssetsLocationsParams object
// with the ability to set a context for a request.
func NewPostCharactersCharacterIDAssetsLocationsParamsWithContext(ctx context.Context) *PostCharactersCharacterIDAssetsLocationsParams {
	return &PostCharactersCharacterIDAssetsLocationsParams{
		Context: ctx,
	}
}

// NewPostCharactersCharacterIDAssetsLocationsParamsWithHTTPClient creates a new PostCharactersCharacterIDAssetsLocationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostCharactersCharacterIDAssetsLocationsParamsWithHTTPClient(client *http.Client) *PostCharactersCharacterIDAssetsLocationsParams {
	return &PostCharactersCharacterIDAssetsLocationsParams{
		HTTPClient: client,
	}
}

/*
PostCharactersCharacterIDAssetsLocationsParams contains all the parameters to send to the API endpoint

	for the post characters character id assets locations operation.

	Typically these are written to a http.Request.
*/
type PostCharactersCharacterIDAssetsLocationsParams struct {

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

	/* ItemIds.

	   A list of item ids
	*/
	ItemIds []int64

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post characters character id assets locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithDefaults() *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post characters character id assets locations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := PostCharactersCharacterIDAssetsLocationsParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithTimeout(timeout time.Duration) *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithContext(ctx context.Context) *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithHTTPClient(client *http.Client) *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithCharacterID(characterID int32) *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithDatasource(datasource *string) *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithItemIds adds the itemIds to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithItemIds(itemIds []int64) *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetItemIds(itemIds)
	return o
}

// SetItemIds adds the itemIds to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetItemIds(itemIds []int64) {
	o.ItemIds = itemIds
}

// WithToken adds the token to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) WithToken(token *string) *PostCharactersCharacterIDAssetsLocationsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post characters character id assets locations params
func (o *PostCharactersCharacterIDAssetsLocationsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PostCharactersCharacterIDAssetsLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.ItemIds != nil {
		if err := r.SetBodyParam(o.ItemIds); err != nil {
			return err
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
