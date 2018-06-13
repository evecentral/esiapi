// Code generated by go-swagger; DO NOT EDIT.

package planetary_interaction

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

// NewGetUniverseSchematicsSchematicIDParams creates a new GetUniverseSchematicsSchematicIDParams object
// with the default values initialized.
func NewGetUniverseSchematicsSchematicIDParams() *GetUniverseSchematicsSchematicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseSchematicsSchematicIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseSchematicsSchematicIDParamsWithTimeout creates a new GetUniverseSchematicsSchematicIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseSchematicsSchematicIDParamsWithTimeout(timeout time.Duration) *GetUniverseSchematicsSchematicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseSchematicsSchematicIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseSchematicsSchematicIDParamsWithContext creates a new GetUniverseSchematicsSchematicIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseSchematicsSchematicIDParamsWithContext(ctx context.Context) *GetUniverseSchematicsSchematicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseSchematicsSchematicIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseSchematicsSchematicIDParamsWithHTTPClient creates a new GetUniverseSchematicsSchematicIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseSchematicsSchematicIDParamsWithHTTPClient(client *http.Client) *GetUniverseSchematicsSchematicIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseSchematicsSchematicIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseSchematicsSchematicIDParams contains all the parameters to send to the API endpoint
for the get universe schematics schematic id operation typically these are written to a http.Request
*/
type GetUniverseSchematicsSchematicIDParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*SchematicID
	  A PI schematic ID

	*/
	SchematicID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) WithTimeout(timeout time.Duration) *GetUniverseSchematicsSchematicIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) WithContext(ctx context.Context) *GetUniverseSchematicsSchematicIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) WithHTTPClient(client *http.Client) *GetUniverseSchematicsSchematicIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseSchematicsSchematicIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) WithDatasource(datasource *string) *GetUniverseSchematicsSchematicIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithSchematicID adds the schematicID to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) WithSchematicID(schematicID int32) *GetUniverseSchematicsSchematicIDParams {
	o.SetSchematicID(schematicID)
	return o
}

// SetSchematicID adds the schematicId to the get universe schematics schematic id params
func (o *GetUniverseSchematicsSchematicIDParams) SetSchematicID(schematicID int32) {
	o.SchematicID = schematicID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseSchematicsSchematicIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param schematic_id
	if err := r.SetPathParam("schematic_id", swag.FormatInt32(o.SchematicID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
