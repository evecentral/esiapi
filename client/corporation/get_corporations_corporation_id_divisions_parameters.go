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

// NewGetCorporationsCorporationIDDivisionsParams creates a new GetCorporationsCorporationIDDivisionsParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDDivisionsParams() *GetCorporationsCorporationIDDivisionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDDivisionsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDDivisionsParamsWithTimeout creates a new GetCorporationsCorporationIDDivisionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDDivisionsParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDDivisionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDDivisionsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDDivisionsParamsWithContext creates a new GetCorporationsCorporationIDDivisionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDDivisionsParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDDivisionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDDivisionsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDDivisionsParamsWithHTTPClient creates a new GetCorporationsCorporationIDDivisionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDDivisionsParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDDivisionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDDivisionsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDDivisionsParams contains all the parameters to send to the API endpoint
for the get corporations corporation id divisions operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDDivisionsParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*CorporationID
	  An EVE corporation ID

	*/
	CorporationID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDDivisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDDivisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDDivisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDDivisionsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDDivisionsParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDDivisionsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) WithToken(token *string) *GetCorporationsCorporationIDDivisionsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id divisions params
func (o *GetCorporationsCorporationIDDivisionsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDDivisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
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