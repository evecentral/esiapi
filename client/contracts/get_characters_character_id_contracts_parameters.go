// Code generated by go-swagger; DO NOT EDIT.

package contracts

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

// NewGetCharactersCharacterIDContractsParams creates a new GetCharactersCharacterIDContractsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDContractsParams() *GetCharactersCharacterIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContractsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDContractsParamsWithTimeout creates a new GetCharactersCharacterIDContractsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDContractsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContractsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDContractsParamsWithContext creates a new GetCharactersCharacterIDContractsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDContractsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContractsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDContractsParamsWithHTTPClient creates a new GetCharactersCharacterIDContractsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDContractsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContractsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDContractsParams contains all the parameters to send to the API endpoint
for the get characters character id contracts operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDContractsParams struct {

	/*XUserAgent
	  Client identifier, takes precedence over User-Agent

	*/
	XUserAgent *string
	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string
	/*UserAgent
	  Client identifier, takes precedence over headers

	*/
	UserAgent *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDContractsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDContractsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDContractsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXUserAgent adds the xUserAgent to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithXUserAgent(xUserAgent *string) *GetCharactersCharacterIDContractsParams {
	o.SetXUserAgent(xUserAgent)
	return o
}

// SetXUserAgent adds the xUserAgent to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetXUserAgent(xUserAgent *string) {
	o.XUserAgent = xUserAgent
}

// WithCharacterID adds the characterID to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDContractsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDContractsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithToken(token *string) *GetCharactersCharacterIDContractsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetToken(token *string) {
	o.Token = token
}

// WithUserAgent adds the userAgent to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) WithUserAgent(userAgent *string) *GetCharactersCharacterIDContractsParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get characters character id contracts params
func (o *GetCharactersCharacterIDContractsParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDContractsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XUserAgent != nil {

		// header param X-User-Agent
		if err := r.SetHeaderParam("X-User-Agent", *o.XUserAgent); err != nil {
			return err
		}

	}

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
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

	if o.UserAgent != nil {

		// query param user_agent
		var qrUserAgent string
		if o.UserAgent != nil {
			qrUserAgent = *o.UserAgent
		}
		qUserAgent := qrUserAgent
		if qUserAgent != "" {
			if err := r.SetQueryParam("user_agent", qUserAgent); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}