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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDogmaAttributesAttributeIDParams creates a new GetDogmaAttributesAttributeIDParams object
// with the default values initialized.
func NewGetDogmaAttributesAttributeIDParams() *GetDogmaAttributesAttributeIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesAttributeIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDogmaAttributesAttributeIDParamsWithTimeout creates a new GetDogmaAttributesAttributeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDogmaAttributesAttributeIDParamsWithTimeout(timeout time.Duration) *GetDogmaAttributesAttributeIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesAttributeIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetDogmaAttributesAttributeIDParamsWithContext creates a new GetDogmaAttributesAttributeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDogmaAttributesAttributeIDParamsWithContext(ctx context.Context) *GetDogmaAttributesAttributeIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesAttributeIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetDogmaAttributesAttributeIDParamsWithHTTPClient creates a new GetDogmaAttributesAttributeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDogmaAttributesAttributeIDParamsWithHTTPClient(client *http.Client) *GetDogmaAttributesAttributeIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetDogmaAttributesAttributeIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetDogmaAttributesAttributeIDParams contains all the parameters to send to the API endpoint
for the get dogma attributes attribute id operation typically these are written to a http.Request
*/
type GetDogmaAttributesAttributeIDParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*AttributeID
	  A dogma attribute ID

	*/
	AttributeID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) WithTimeout(timeout time.Duration) *GetDogmaAttributesAttributeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) WithContext(ctx context.Context) *GetDogmaAttributesAttributeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) WithHTTPClient(client *http.Client) *GetDogmaAttributesAttributeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetDogmaAttributesAttributeIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithAttributeID adds the attributeID to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) WithAttributeID(attributeID int32) *GetDogmaAttributesAttributeIDParams {
	o.SetAttributeID(attributeID)
	return o
}

// SetAttributeID adds the attributeId to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) SetAttributeID(attributeID int32) {
	o.AttributeID = attributeID
}

// WithDatasource adds the datasource to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) WithDatasource(datasource *string) *GetDogmaAttributesAttributeIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get dogma attributes attribute id params
func (o *GetDogmaAttributesAttributeIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetDogmaAttributesAttributeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param attribute_id
	if err := r.SetPathParam("attribute_id", swag.FormatInt32(o.AttributeID)); err != nil {
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
