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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFwLeaderboardsCharactersParams creates a new GetFwLeaderboardsCharactersParams object
// with the default values initialized.
func NewGetFwLeaderboardsCharactersParams() *GetFwLeaderboardsCharactersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetFwLeaderboardsCharactersParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFwLeaderboardsCharactersParamsWithTimeout creates a new GetFwLeaderboardsCharactersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFwLeaderboardsCharactersParamsWithTimeout(timeout time.Duration) *GetFwLeaderboardsCharactersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetFwLeaderboardsCharactersParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetFwLeaderboardsCharactersParamsWithContext creates a new GetFwLeaderboardsCharactersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFwLeaderboardsCharactersParamsWithContext(ctx context.Context) *GetFwLeaderboardsCharactersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetFwLeaderboardsCharactersParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetFwLeaderboardsCharactersParamsWithHTTPClient creates a new GetFwLeaderboardsCharactersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFwLeaderboardsCharactersParamsWithHTTPClient(client *http.Client) *GetFwLeaderboardsCharactersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetFwLeaderboardsCharactersParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetFwLeaderboardsCharactersParams contains all the parameters to send to the API endpoint
for the get fw leaderboards characters operation typically these are written to a http.Request
*/
type GetFwLeaderboardsCharactersParams struct {

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

// WithTimeout adds the timeout to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) WithTimeout(timeout time.Duration) *GetFwLeaderboardsCharactersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) WithContext(ctx context.Context) *GetFwLeaderboardsCharactersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) WithHTTPClient(client *http.Client) *GetFwLeaderboardsCharactersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) WithIfNoneMatch(ifNoneMatch *string) *GetFwLeaderboardsCharactersParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) WithDatasource(datasource *string) *GetFwLeaderboardsCharactersParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get fw leaderboards characters params
func (o *GetFwLeaderboardsCharactersParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetFwLeaderboardsCharactersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
