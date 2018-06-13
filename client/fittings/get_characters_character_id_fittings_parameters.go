// Code generated by go-swagger; DO NOT EDIT.

package fittings

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

// NewGetCharactersCharacterIDFittingsParams creates a new GetCharactersCharacterIDFittingsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDFittingsParams() *GetCharactersCharacterIDFittingsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFittingsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDFittingsParamsWithTimeout creates a new GetCharactersCharacterIDFittingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDFittingsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDFittingsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFittingsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDFittingsParamsWithContext creates a new GetCharactersCharacterIDFittingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDFittingsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDFittingsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFittingsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDFittingsParamsWithHTTPClient creates a new GetCharactersCharacterIDFittingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDFittingsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDFittingsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFittingsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDFittingsParams contains all the parameters to send to the API endpoint
for the get characters character id fittings operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDFittingsParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDFittingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDFittingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDFittingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDFittingsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDFittingsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDFittingsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) WithToken(token *string) *GetCharactersCharacterIDFittingsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id fittings params
func (o *GetCharactersCharacterIDFittingsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDFittingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
