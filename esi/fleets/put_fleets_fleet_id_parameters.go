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

// NewPutFleetsFleetIDParams creates a new PutFleetsFleetIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutFleetsFleetIDParams() *PutFleetsFleetIDParams {
	return &PutFleetsFleetIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutFleetsFleetIDParamsWithTimeout creates a new PutFleetsFleetIDParams object
// with the ability to set a timeout on a request.
func NewPutFleetsFleetIDParamsWithTimeout(timeout time.Duration) *PutFleetsFleetIDParams {
	return &PutFleetsFleetIDParams{
		timeout: timeout,
	}
}

// NewPutFleetsFleetIDParamsWithContext creates a new PutFleetsFleetIDParams object
// with the ability to set a context for a request.
func NewPutFleetsFleetIDParamsWithContext(ctx context.Context) *PutFleetsFleetIDParams {
	return &PutFleetsFleetIDParams{
		Context: ctx,
	}
}

// NewPutFleetsFleetIDParamsWithHTTPClient creates a new PutFleetsFleetIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutFleetsFleetIDParamsWithHTTPClient(client *http.Client) *PutFleetsFleetIDParams {
	return &PutFleetsFleetIDParams{
		HTTPClient: client,
	}
}

/*
PutFleetsFleetIDParams contains all the parameters to send to the API endpoint

	for the put fleets fleet id operation.

	Typically these are written to a http.Request.
*/
type PutFleetsFleetIDParams struct {

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

	/* NewSettings.

	   What to update for this fleet
	*/
	NewSettings PutFleetsFleetIDBody

	/* Token.

	   Access token to use if unable to set a header
	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put fleets fleet id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFleetsFleetIDParams) WithDefaults() *PutFleetsFleetIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put fleets fleet id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutFleetsFleetIDParams) SetDefaults() {
	var (
		datasourceDefault = string("tranquility")
	)

	val := PutFleetsFleetIDParams{
		Datasource: &datasourceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) WithTimeout(timeout time.Duration) *PutFleetsFleetIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) WithContext(ctx context.Context) *PutFleetsFleetIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) WithHTTPClient(client *http.Client) *PutFleetsFleetIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasource adds the datasource to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) WithDatasource(datasource *string) *PutFleetsFleetIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFleetID adds the fleetID to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) WithFleetID(fleetID int64) *PutFleetsFleetIDParams {
	o.SetFleetID(fleetID)
	return o
}

// SetFleetID adds the fleetId to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) SetFleetID(fleetID int64) {
	o.FleetID = fleetID
}

// WithNewSettings adds the newSettings to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) WithNewSettings(newSettings PutFleetsFleetIDBody) *PutFleetsFleetIDParams {
	o.SetNewSettings(newSettings)
	return o
}

// SetNewSettings adds the newSettings to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) SetNewSettings(newSettings PutFleetsFleetIDBody) {
	o.NewSettings = newSettings
}

// WithToken adds the token to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) WithToken(token *string) *PutFleetsFleetIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the put fleets fleet id params
func (o *PutFleetsFleetIDParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PutFleetsFleetIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.NewSettings); err != nil {
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
