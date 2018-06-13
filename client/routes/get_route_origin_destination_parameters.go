// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRouteOriginDestinationParams creates a new GetRouteOriginDestinationParams object
// with the default values initialized.
func NewGetRouteOriginDestinationParams() *GetRouteOriginDestinationParams {
	var (
		datasourceDefault = string("tranquility")
		flagDefault       = string("shortest")
	)
	return &GetRouteOriginDestinationParams{
		Datasource: &datasourceDefault,
		Flag:       &flagDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRouteOriginDestinationParamsWithTimeout creates a new GetRouteOriginDestinationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRouteOriginDestinationParamsWithTimeout(timeout time.Duration) *GetRouteOriginDestinationParams {
	var (
		datasourceDefault = string("tranquility")
		flagDefault       = string("shortest")
	)
	return &GetRouteOriginDestinationParams{
		Datasource: &datasourceDefault,
		Flag:       &flagDefault,

		timeout: timeout,
	}
}

// NewGetRouteOriginDestinationParamsWithContext creates a new GetRouteOriginDestinationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRouteOriginDestinationParamsWithContext(ctx context.Context) *GetRouteOriginDestinationParams {
	var (
		datasourceDefault = string("tranquility")
		flagDefault       = string("shortest")
	)
	return &GetRouteOriginDestinationParams{
		Datasource: &datasourceDefault,
		Flag:       &flagDefault,

		Context: ctx,
	}
}

// NewGetRouteOriginDestinationParamsWithHTTPClient creates a new GetRouteOriginDestinationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRouteOriginDestinationParamsWithHTTPClient(client *http.Client) *GetRouteOriginDestinationParams {
	var (
		datasourceDefault = string("tranquility")
		flagDefault       = string("shortest")
	)
	return &GetRouteOriginDestinationParams{
		Datasource: &datasourceDefault,
		Flag:       &flagDefault,
		HTTPClient: client,
	}
}

/*GetRouteOriginDestinationParams contains all the parameters to send to the API endpoint
for the get route origin destination operation typically these are written to a http.Request
*/
type GetRouteOriginDestinationParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*Avoid
	  avoid solar system ID(s)

	*/
	Avoid []int32
	/*Connections
	  connected solar system pairs

	*/
	Connections [][]int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Destination
	  destination solar system ID

	*/
	Destination int32
	/*Flag
	  route security preference

	*/
	Flag *string
	/*Origin
	  origin solar system ID

	*/
	Origin int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithTimeout(timeout time.Duration) *GetRouteOriginDestinationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithContext(ctx context.Context) *GetRouteOriginDestinationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithHTTPClient(client *http.Client) *GetRouteOriginDestinationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithIfNoneMatch(ifNoneMatch *string) *GetRouteOriginDestinationParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithAvoid adds the avoid to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithAvoid(avoid []int32) *GetRouteOriginDestinationParams {
	o.SetAvoid(avoid)
	return o
}

// SetAvoid adds the avoid to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetAvoid(avoid []int32) {
	o.Avoid = avoid
}

// WithConnections adds the connections to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithConnections(connections [][]int32) *GetRouteOriginDestinationParams {
	o.SetConnections(connections)
	return o
}

// SetConnections adds the connections to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetConnections(connections [][]int32) {
	o.Connections = connections
}

// WithDatasource adds the datasource to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithDatasource(datasource *string) *GetRouteOriginDestinationParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithDestination adds the destination to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithDestination(destination int32) *GetRouteOriginDestinationParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetDestination(destination int32) {
	o.Destination = destination
}

// WithFlag adds the flag to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithFlag(flag *string) *GetRouteOriginDestinationParams {
	o.SetFlag(flag)
	return o
}

// SetFlag adds the flag to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetFlag(flag *string) {
	o.Flag = flag
}

// WithOrigin adds the origin to the get route origin destination params
func (o *GetRouteOriginDestinationParams) WithOrigin(origin int32) *GetRouteOriginDestinationParams {
	o.SetOrigin(origin)
	return o
}

// SetOrigin adds the origin to the get route origin destination params
func (o *GetRouteOriginDestinationParams) SetOrigin(origin int32) {
	o.Origin = origin
}

// WriteToRequest writes these params to a swagger request
func (o *GetRouteOriginDestinationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesAvoid []string
	for _, v := range o.Avoid {
		valuesAvoid = append(valuesAvoid, swag.FormatInt32(v))
	}

	joinedAvoid := swag.JoinByFormat(valuesAvoid, "")
	// query array param avoid
	if err := r.SetQueryParam("avoid", joinedAvoid...); err != nil {
		return err
	}

	valuesConnections := o.Connections

	joinedConnections := swag.JoinByFormat(valuesConnections, "")
	// query array param connections
	if err := r.SetQueryParam("connections", joinedConnections...); err != nil {
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

	// path param destination
	if err := r.SetPathParam("destination", swag.FormatInt32(o.Destination)); err != nil {
		return err
	}

	if o.Flag != nil {

		// query param flag
		var qrFlag string
		if o.Flag != nil {
			qrFlag = *o.Flag
		}
		qFlag := qrFlag
		if qFlag != "" {
			if err := r.SetQueryParam("flag", qFlag); err != nil {
				return err
			}
		}

	}

	// path param origin
	if err := r.SetPathParam("origin", swag.FormatInt32(o.Origin)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
