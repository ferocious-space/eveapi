// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// NewPutFleetsFleetIDWingsWingIDParams creates a new PutFleetsFleetIDWingsWingIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutFleetsFleetIDWingsWingIDParams() *PutFleetsFleetIDWingsWingIDParams {
	return &PutFleetsFleetIDWingsWingIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutFleetsFleetIDWingsWingIDParamsWithTimeout creates a new PutFleetsFleetIDWingsWingIDParams object
// with the ability to set a timeout on a request.
func NewPutFleetsFleetIDWingsWingIDParamsWithTimeout(timeout time.Duration) *PutFleetsFleetIDWingsWingIDParams {
	return &PutFleetsFleetIDWingsWingIDParams{
		timeout: timeout,
	}
}

// NewPutFleetsFleetIDWingsWingIDParamsWithContext creates a new PutFleetsFleetIDWingsWingIDParams object
// with the ability to set a context for a request.
func NewPutFleetsFleetIDWingsWingIDParamsWithContext(ctx context.Context) *PutFleetsFleetIDWingsWingIDParams {
	return &PutFleetsFleetIDWingsWingIDParams{
		Context: ctx,
	}
}

// NewPutFleetsFleetIDWingsWingIDParamsWithHTTPClient creates a new PutFleetsFleetIDWingsWingIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutFleetsFleetIDWingsWingIDParamsWithHTTPClient(client *http.Client) *PutFleetsFleetIDWingsWingIDParams {
	return &PutFleetsFleetIDWingsWingIDParams{
		HTTPClient: client,
	}
}

/*
PutFleetsFleetIDWingsWingIDParams contains all the parameters to send to the API endpoint

	for the put fleets fleet id wings wing id operation.

	Typically these are written to a http.Request.
*/
type PutFleetsFleetIDWingsWingIDParams struct {

	/* Datasource.

	   The server name you would like data from

	   Default: "tranquility"
	*/
	Datasource *string

	/* FleetID.

	   ID for a fleet

	   Format: int64
	*/
	FleetID int64

	/* Naming.

	   New name of the wing
	*/
	Naming PutFleetsFleetIDWingsWingIDBody

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	/* WingID.

	   The wing to rename

	   Format: int64
	*/
	WingID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put fleets fleet id wings wing id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFleetsFleetIDWingsWingIDParams) WithDefaults() *PutFleetsFleetIDWingsWingIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put fleets fleet id wings wing id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFleetsFleetIDWingsWingIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := PutFleetsFleetIDWingsWingIDParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithTimeout(timeout time.Duration) *PutFleetsFleetIDWingsWingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithContext(ctx context.Context) *PutFleetsFleetIDWingsWingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithHTTPClient(client *http.Client) *PutFleetsFleetIDWingsWingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasource adds the datasource to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithDatasource(datasource *string) *PutFleetsFleetIDWingsWingIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFleetID adds the fleetID to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithFleetID(fleetID int64) *PutFleetsFleetIDWingsWingIDParams {
	o.SetFleetID(fleetID)
	return o
}

// SetFleetID adds the fleetId to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetFleetID(fleetID int64) {
	o.FleetID = fleetID
}

// WithNaming adds the naming to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithNaming(naming PutFleetsFleetIDWingsWingIDBody) *PutFleetsFleetIDWingsWingIDParams {
	o.SetNaming(naming)
	return o
}

// SetNaming adds the naming to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetNaming(naming PutFleetsFleetIDWingsWingIDBody) {
	o.Naming = naming
}

// WithToken adds the token to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithToken(token *string) *PutFleetsFleetIDWingsWingIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetToken(token *string) {
	o.Token = token
}

// WithWingID adds the wingID to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) WithWingID(wingID int64) *PutFleetsFleetIDWingsWingIDParams {
	o.SetWingID(wingID)
	return o
}

// SetWingID adds the wingId to the put fleets fleet id wings wing id params
func (o *PutFleetsFleetIDWingsWingIDParams) SetWingID(wingID int64) {
	o.WingID = wingID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFleetsFleetIDWingsWingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param fleet_id
	if err := r.SetPathParam("fleet_id", swag.FormatInt64(o.FleetID)); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Naming); err != nil {
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

	// path param wing_id
	if err := r.SetPathParam("wing_id", swag.FormatInt64(o.WingID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
