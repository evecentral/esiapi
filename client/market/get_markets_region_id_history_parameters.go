// Code generated by go-swagger; DO NOT EDIT.

package market

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

// NewGetMarketsRegionIDHistoryParams creates a new GetMarketsRegionIDHistoryParams object
// with the default values initialized.
func NewGetMarketsRegionIDHistoryParams() *GetMarketsRegionIDHistoryParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsRegionIDHistoryParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketsRegionIDHistoryParamsWithTimeout creates a new GetMarketsRegionIDHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMarketsRegionIDHistoryParamsWithTimeout(timeout time.Duration) *GetMarketsRegionIDHistoryParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsRegionIDHistoryParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetMarketsRegionIDHistoryParamsWithContext creates a new GetMarketsRegionIDHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMarketsRegionIDHistoryParamsWithContext(ctx context.Context) *GetMarketsRegionIDHistoryParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsRegionIDHistoryParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetMarketsRegionIDHistoryParamsWithHTTPClient creates a new GetMarketsRegionIDHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMarketsRegionIDHistoryParamsWithHTTPClient(client *http.Client) *GetMarketsRegionIDHistoryParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetMarketsRegionIDHistoryParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetMarketsRegionIDHistoryParams contains all the parameters to send to the API endpoint
for the get markets region id history operation typically these are written to a http.Request
*/
type GetMarketsRegionIDHistoryParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*RegionID
	  Return statistics in this region

	*/
	RegionID int32
	/*TypeID
	  Return statistics for this type

	*/
	TypeID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) WithTimeout(timeout time.Duration) *GetMarketsRegionIDHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) WithContext(ctx context.Context) *GetMarketsRegionIDHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) WithHTTPClient(client *http.Client) *GetMarketsRegionIDHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) WithIfNoneMatch(ifNoneMatch *string) *GetMarketsRegionIDHistoryParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) WithDatasource(datasource *string) *GetMarketsRegionIDHistoryParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithRegionID adds the regionID to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) WithRegionID(regionID int32) *GetMarketsRegionIDHistoryParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) SetRegionID(regionID int32) {
	o.RegionID = regionID
}

// WithTypeID adds the typeID to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) WithTypeID(typeID int32) *GetMarketsRegionIDHistoryParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the get markets region id history params
func (o *GetMarketsRegionIDHistoryParams) SetTypeID(typeID int32) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketsRegionIDHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param region_id
	if err := r.SetPathParam("region_id", swag.FormatInt32(o.RegionID)); err != nil {
		return err
	}

	// query param type_id
	qrTypeID := o.TypeID
	qTypeID := swag.FormatInt32(qrTypeID)
	if qTypeID != "" {
		if err := r.SetQueryParam("type_id", qTypeID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
