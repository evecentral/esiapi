// Code generated by go-swagger; DO NOT EDIT.

package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDogmaAttributesParams creates a new GetDogmaAttributesParams object
// with the default values initialized.
func NewGetDogmaAttributesParams() *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDogmaAttributesParamsWithTimeout creates a new GetDogmaAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDogmaAttributesParamsWithTimeout(timeout time.Duration) *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetDogmaAttributesParamsWithContext creates a new GetDogmaAttributesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDogmaAttributesParamsWithContext(ctx context.Context) *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetDogmaAttributesParamsWithHTTPClient creates a new GetDogmaAttributesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDogmaAttributesParamsWithHTTPClient(client *http.Client) *GetDogmaAttributesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetDogmaAttributesParams contains all the parameters to send to the API endpoint
for the get dogma attributes operation typically these are written to a http.Request
*/
type GetDogmaAttributesParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithTimeout(timeout time.Duration) *GetDogmaAttributesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithContext(ctx context.Context) *GetDogmaAttributesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithHTTPClient(client *http.Client) *GetDogmaAttributesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithIfNoneMatch(ifNoneMatch *string) *GetDogmaAttributesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get dogma attributes params
func (o *GetDogmaAttributesParams) WithDatasource(datasource *string) *GetDogmaAttributesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get dogma attributes params
func (o *GetDogmaAttributesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetDogmaAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
