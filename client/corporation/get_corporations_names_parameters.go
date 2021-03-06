// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// NewGetCorporationsNamesParams creates a new GetCorporationsNamesParams object
// with the default values initialized.
func NewGetCorporationsNamesParams() *GetCorporationsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsNamesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsNamesParamsWithTimeout creates a new GetCorporationsNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsNamesParamsWithTimeout(timeout time.Duration) *GetCorporationsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsNamesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsNamesParamsWithContext creates a new GetCorporationsNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsNamesParamsWithContext(ctx context.Context) *GetCorporationsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsNamesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsNamesParamsWithHTTPClient creates a new GetCorporationsNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsNamesParamsWithHTTPClient(client *http.Client) *GetCorporationsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsNamesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsNamesParams contains all the parameters to send to the API endpoint
for the get corporations names operation typically these are written to a http.Request
*/
type GetCorporationsNamesParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*CorporationIds
	  A comma separated list of corporation IDs

	*/
	CorporationIds []int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get corporations names params
func (o *GetCorporationsNamesParams) WithTimeout(timeout time.Duration) *GetCorporationsNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations names params
func (o *GetCorporationsNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations names params
func (o *GetCorporationsNamesParams) WithContext(ctx context.Context) *GetCorporationsNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations names params
func (o *GetCorporationsNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations names params
func (o *GetCorporationsNamesParams) WithHTTPClient(client *http.Client) *GetCorporationsNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations names params
func (o *GetCorporationsNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations names params
func (o *GetCorporationsNamesParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsNamesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations names params
func (o *GetCorporationsNamesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationIds adds the corporationIds to the get corporations names params
func (o *GetCorporationsNamesParams) WithCorporationIds(corporationIds []int32) *GetCorporationsNamesParams {
	o.SetCorporationIds(corporationIds)
	return o
}

// SetCorporationIds adds the corporationIds to the get corporations names params
func (o *GetCorporationsNamesParams) SetCorporationIds(corporationIds []int32) {
	o.CorporationIds = corporationIds
}

// WithDatasource adds the datasource to the get corporations names params
func (o *GetCorporationsNamesParams) WithDatasource(datasource *string) *GetCorporationsNamesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations names params
func (o *GetCorporationsNamesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesCorporationIds []string
	for _, v := range o.CorporationIds {
		valuesCorporationIds = append(valuesCorporationIds, swag.FormatInt32(v))
	}

	joinedCorporationIds := swag.JoinByFormat(valuesCorporationIds, "")
	// query array param corporation_ids
	if err := r.SetQueryParam("corporation_ids", joinedCorporationIds...); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
