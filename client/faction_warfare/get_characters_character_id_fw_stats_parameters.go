// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

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

// NewGetCharactersCharacterIDFwStatsParams creates a new GetCharactersCharacterIDFwStatsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDFwStatsParams() *GetCharactersCharacterIDFwStatsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFwStatsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDFwStatsParamsWithTimeout creates a new GetCharactersCharacterIDFwStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDFwStatsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDFwStatsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFwStatsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDFwStatsParamsWithContext creates a new GetCharactersCharacterIDFwStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDFwStatsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDFwStatsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFwStatsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDFwStatsParamsWithHTTPClient creates a new GetCharactersCharacterIDFwStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDFwStatsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDFwStatsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDFwStatsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDFwStatsParams contains all the parameters to send to the API endpoint
for the get characters character id fw stats operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDFwStatsParams struct {

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

// WithTimeout adds the timeout to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDFwStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDFwStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDFwStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDFwStatsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDFwStatsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDFwStatsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) WithToken(token *string) *GetCharactersCharacterIDFwStatsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id fw stats params
func (o *GetCharactersCharacterIDFwStatsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDFwStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
