// Code generated by go-swagger; DO NOT EDIT.

package bookmarks

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

// NewGetCharactersCharacterIDBookmarksFoldersParams creates a new GetCharactersCharacterIDBookmarksFoldersParams object
// with the default values initialized.
func NewGetCharactersCharacterIDBookmarksFoldersParams() *GetCharactersCharacterIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDBookmarksFoldersParamsWithTimeout creates a new GetCharactersCharacterIDBookmarksFoldersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDBookmarksFoldersParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDBookmarksFoldersParamsWithContext creates a new GetCharactersCharacterIDBookmarksFoldersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDBookmarksFoldersParamsWithContext(ctx context.Context) *GetCharactersCharacterIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDBookmarksFoldersParamsWithHTTPClient creates a new GetCharactersCharacterIDBookmarksFoldersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDBookmarksFoldersParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDBookmarksFoldersParams contains all the parameters to send to the API endpoint
for the get characters character id bookmarks folders operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDBookmarksFoldersParams struct {

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
	/*Page
	  Which page of results to return

	*/
	Page *int32
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithContext(ctx context.Context) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithDatasource(datasource *string) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithPage(page *int32) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WithToken(token *string) *GetCharactersCharacterIDBookmarksFoldersParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id bookmarks folders params
func (o *GetCharactersCharacterIDBookmarksFoldersParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDBookmarksFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
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
