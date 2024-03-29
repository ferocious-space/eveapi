// Code generated by go-swagger; DO NOT EDIT.

package dogma

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

// NewGetDogmaDynamicItemsTypeIDItemIDParams creates a new GetDogmaDynamicItemsTypeIDItemIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDogmaDynamicItemsTypeIDItemIDParams() *GetDogmaDynamicItemsTypeIDItemIDParams {
	return &GetDogmaDynamicItemsTypeIDItemIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDogmaDynamicItemsTypeIDItemIDParamsWithTimeout creates a new GetDogmaDynamicItemsTypeIDItemIDParams object
// with the ability to set a timeout on a request.
func NewGetDogmaDynamicItemsTypeIDItemIDParamsWithTimeout(timeout time.Duration) *GetDogmaDynamicItemsTypeIDItemIDParams {
	return &GetDogmaDynamicItemsTypeIDItemIDParams{
		timeout: timeout,
	}
}

// NewGetDogmaDynamicItemsTypeIDItemIDParamsWithContext creates a new GetDogmaDynamicItemsTypeIDItemIDParams object
// with the ability to set a context for a request.
func NewGetDogmaDynamicItemsTypeIDItemIDParamsWithContext(ctx context.Context) *GetDogmaDynamicItemsTypeIDItemIDParams {
	return &GetDogmaDynamicItemsTypeIDItemIDParams{
		Context: ctx,
	}
}

// NewGetDogmaDynamicItemsTypeIDItemIDParamsWithHTTPClient creates a new GetDogmaDynamicItemsTypeIDItemIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDogmaDynamicItemsTypeIDItemIDParamsWithHTTPClient(client *http.Client) *GetDogmaDynamicItemsTypeIDItemIDParams {
	return &GetDogmaDynamicItemsTypeIDItemIDParams{
		HTTPClient: client,
	}
}

/*
GetDogmaDynamicItemsTypeIDItemIDParams contains all the parameters to send to the API endpoint

	for the get dogma dynamic items type id item id operation.

	Typically these are written to a http.Request.
*/
type GetDogmaDynamicItemsTypeIDItemIDParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* ItemID.

	   item_id integer

	   Format: int64
	*/
	ItemID int64

	/* TypeID.

	   type_id integer

	   Format: int32
	*/
	TypeID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dogma dynamic items type id item id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithDefaults() *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dogma dynamic items type id item id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := GetDogmaDynamicItemsTypeIDItemIDParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithTimeout(timeout time.Duration) *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithContext(ctx context.Context) *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithHTTPClient(client *http.Client) *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithDatasource(datasource *string) *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithItemID adds the itemID to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithItemID(itemID int64) *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetItemID(itemID int64) {
	o.ItemID = itemID
}

// WithTypeID adds the typeID to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WithTypeID(typeID int32) *GetDogmaDynamicItemsTypeIDItemIDParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the get dogma dynamic items type id item id params
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) SetTypeID(typeID int32) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDogmaDynamicItemsTypeIDItemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param item_id
	if err := r.SetPathParam("item_id", swag.FormatInt64(o.ItemID)); err != nil {
		return err
	}

	// path param type_id
	if err := r.SetPathParam("type_id", swag.FormatInt32(o.TypeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
