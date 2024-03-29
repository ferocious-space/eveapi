// Code generated by go-swagger; DO NOT EDIT.

package market

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

// NewGetMarketsRegionIDOrdersParams creates a new GetMarketsRegionIDOrdersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMarketsRegionIDOrdersParams() *GetMarketsRegionIDOrdersParams {
	return &GetMarketsRegionIDOrdersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketsRegionIDOrdersParamsWithTimeout creates a new GetMarketsRegionIDOrdersParams object
// with the ability to set a timeout on a request.
func NewGetMarketsRegionIDOrdersParamsWithTimeout(timeout time.Duration) *GetMarketsRegionIDOrdersParams {
	return &GetMarketsRegionIDOrdersParams{
		timeout: timeout,
	}
}

// NewGetMarketsRegionIDOrdersParamsWithContext creates a new GetMarketsRegionIDOrdersParams object
// with the ability to set a context for a request.
func NewGetMarketsRegionIDOrdersParamsWithContext(ctx context.Context) *GetMarketsRegionIDOrdersParams {
	return &GetMarketsRegionIDOrdersParams{
		Context: ctx,
	}
}

// NewGetMarketsRegionIDOrdersParamsWithHTTPClient creates a new GetMarketsRegionIDOrdersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMarketsRegionIDOrdersParamsWithHTTPClient(client *http.Client) *GetMarketsRegionIDOrdersParams {
	return &GetMarketsRegionIDOrdersParams{
		HTTPClient: client,
	}
}

/*
GetMarketsRegionIDOrdersParams contains all the parameters to send to the API endpoint

	for the get markets region id orders operation.

	Typically these are written to a http.Request.
*/
type GetMarketsRegionIDOrdersParams struct {

	/* IfNoneMatch.

	   ETag from a previous request. A 304 will be returned if this matches the current ETag
	*/
	IfNoneMatch *string

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* OrderType.

	   Filter buy/sell orders, return all orders by default. If you query without type_id, we always return both buy and sell orders

	   Default: "all"
	*/
	OrderType string

	/* Page.

	   Which page of results to return

	   Format: int32
	   Default: 1
	*/
	Page *int32

	/* RegionID.

	   Return orders in this region

	   Format: int32
	*/
	RegionID int32

	/* TypeID.

	   Return orders only for this type

	   Format: int32
	*/
	TypeID *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get markets region id orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketsRegionIDOrdersParams) WithDefaults() *GetMarketsRegionIDOrdersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get markets region id orders params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketsRegionIDOrdersParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")

		orderTypeDefault = string("all")

		pageDefault = int32(1)
	)

	val := GetMarketsRegionIDOrdersParams{
		Datasource: &datasourceDefault,
		OrderType:  orderTypeDefault,
		Page:       &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithTimeout(timeout time.Duration) *GetMarketsRegionIDOrdersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithContext(ctx context.Context) *GetMarketsRegionIDOrdersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithHTTPClient(client *http.Client) *GetMarketsRegionIDOrdersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithIfNoneMatch(ifNoneMatch *string) *GetMarketsRegionIDOrdersParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithDatasource(datasource *string) *GetMarketsRegionIDOrdersParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithOrderType adds the orderType to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithOrderType(orderType string) *GetMarketsRegionIDOrdersParams {
	o.SetOrderType(orderType)
	return o
}

// SetOrderType adds the orderType to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetOrderType(orderType string) {
	o.OrderType = orderType
}

// WithPage adds the page to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithPage(page *int32) *GetMarketsRegionIDOrdersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetPage(page *int32) {
	o.Page = page
}

// WithRegionID adds the regionID to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithRegionID(regionID int32) *GetMarketsRegionIDOrdersParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetRegionID(regionID int32) {
	o.RegionID = regionID
}

// WithTypeID adds the typeID to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) WithTypeID(typeID *int32) *GetMarketsRegionIDOrdersParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the get markets region id orders params
func (o *GetMarketsRegionIDOrdersParams) SetTypeID(typeID *int32) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketsRegionIDOrdersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param order_type
	qrOrderType := o.OrderType
	qOrderType := qrOrderType
	if qOrderType != "" {

		if err := r.SetQueryParam("order_type", qOrderType); err != nil {
			return err
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

	// path param region_id
	if err := r.SetPathParam("region_id", swag.FormatInt32(o.RegionID)); err != nil {
		return err
	}

	if o.TypeID != nil {

		// query param type_id
		var qrTypeID int32

		if o.TypeID != nil {
			qrTypeID = *o.TypeID
		}
		qTypeID := swag.FormatInt32(qrTypeID)
		if qTypeID != "" {

			if err := r.SetQueryParam("type_id", qTypeID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
