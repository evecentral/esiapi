// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewGetCharactersCharacterIDContactsLabelsParams creates a new GetCharactersCharacterIDContactsLabelsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDContactsLabelsParams() *GetCharactersCharacterIDContactsLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContactsLabelsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDContactsLabelsParamsWithTimeout creates a new GetCharactersCharacterIDContactsLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDContactsLabelsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDContactsLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContactsLabelsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDContactsLabelsParamsWithContext creates a new GetCharactersCharacterIDContactsLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDContactsLabelsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDContactsLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContactsLabelsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDContactsLabelsParamsWithHTTPClient creates a new GetCharactersCharacterIDContactsLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDContactsLabelsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDContactsLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDContactsLabelsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDContactsLabelsParams contains all the parameters to send to the API endpoint
for the get characters character id contacts labels operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDContactsLabelsParams struct {

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

// WithTimeout adds the timeout to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDContactsLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDContactsLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDContactsLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDContactsLabelsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDContactsLabelsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDContactsLabelsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) WithToken(token *string) *GetCharactersCharacterIDContactsLabelsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id contacts labels params
func (o *GetCharactersCharacterIDContactsLabelsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDContactsLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
