// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// NewGetUniverseTypesTypeIDParams creates a new GetUniverseTypesTypeIDParams object
// with the default values initialized.
func NewGetUniverseTypesTypeIDParams() *GetUniverseTypesTypeIDParams {
	var (
		acceptLanguageDefault = string("en-us")
		datasourceDefault     = string("tranquility")
		languageDefault       = string("en-us")
	)
	return &GetUniverseTypesTypeIDParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseTypesTypeIDParamsWithTimeout creates a new GetUniverseTypesTypeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseTypesTypeIDParamsWithTimeout(timeout time.Duration) *GetUniverseTypesTypeIDParams {
	var (
		acceptLanguageDefault = string("en-us")
		datasourceDefault     = string("tranquility")
		languageDefault       = string("en-us")
	)
	return &GetUniverseTypesTypeIDParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,

		timeout: timeout,
	}
}

// NewGetUniverseTypesTypeIDParamsWithContext creates a new GetUniverseTypesTypeIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseTypesTypeIDParamsWithContext(ctx context.Context) *GetUniverseTypesTypeIDParams {
	var (
		acceptLanguageDefault = string("en-us")
		datasourceDefault     = string("tranquility")
		languageDefault       = string("en-us")
	)
	return &GetUniverseTypesTypeIDParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,

		Context: ctx,
	}
}

// NewGetUniverseTypesTypeIDParamsWithHTTPClient creates a new GetUniverseTypesTypeIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseTypesTypeIDParamsWithHTTPClient(client *http.Client) *GetUniverseTypesTypeIDParams {
	var (
		acceptLanguageDefault = string("en-us")
		datasourceDefault     = string("tranquility")
		languageDefault       = string("en-us")
	)
	return &GetUniverseTypesTypeIDParams{
		AcceptLanguage: &acceptLanguageDefault,
		Datasource:     &datasourceDefault,
		Language:       &languageDefault,
		HTTPClient:     client,
	}
}

/*GetUniverseTypesTypeIDParams contains all the parameters to send to the API endpoint
for the get universe types type id operation typically these are written to a http.Request
*/
type GetUniverseTypesTypeIDParams struct {

	/*AcceptLanguage
	  Language to use in the response

	*/
	AcceptLanguage *string
	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Language
	  Language to use in the response, takes precedence over Accept-Language

	*/
	Language *string
	/*TypeID
	  An Eve item type ID

	*/
	TypeID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithTimeout(timeout time.Duration) *GetUniverseTypesTypeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithContext(ctx context.Context) *GetUniverseTypesTypeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithHTTPClient(client *http.Client) *GetUniverseTypesTypeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithAcceptLanguage(acceptLanguage *string) *GetUniverseTypesTypeIDParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseTypesTypeIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithDatasource(datasource *string) *GetUniverseTypesTypeIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLanguage adds the language to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithLanguage(language *string) *GetUniverseTypesTypeIDParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetLanguage(language *string) {
	o.Language = language
}

// WithTypeID adds the typeID to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) WithTypeID(typeID int32) *GetUniverseTypesTypeIDParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the get universe types type id params
func (o *GetUniverseTypesTypeIDParams) SetTypeID(typeID int32) {
	o.TypeID = typeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseTypesTypeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

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

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	// path param type_id
	if err := r.SetPathParam("type_id", swag.FormatInt32(o.TypeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
