package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUniverseStructuresStructureIDParams creates a new GetUniverseStructuresStructureIDParams object
// with the default values initialized.
func NewGetUniverseStructuresStructureIDParams() *GetUniverseStructuresStructureIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresStructureIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseStructuresStructureIDParamsWithTimeout creates a new GetUniverseStructuresStructureIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseStructuresStructureIDParamsWithTimeout(timeout time.Duration) *GetUniverseStructuresStructureIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresStructureIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseStructuresStructureIDParamsWithContext creates a new GetUniverseStructuresStructureIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseStructuresStructureIDParamsWithContext(ctx context.Context) *GetUniverseStructuresStructureIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresStructureIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetUniverseStructuresStructureIDParams contains all the parameters to send to the API endpoint
for the get universe structures structure id operation typically these are written to a http.Request
*/
type GetUniverseStructuresStructureIDParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*StructureID
	  An Eve structure ID

	*/
	StructureID int64

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithTimeout(timeout time.Duration) *GetUniverseStructuresStructureIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithContext(ctx context.Context) *GetUniverseStructuresStructureIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithDatasource(datasource *string) *GetUniverseStructuresStructureIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithStructureID adds the structureID to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) WithStructureID(structureID int64) *GetUniverseStructuresStructureIDParams {
	o.SetStructureID(structureID)
	return o
}

// SetStructureID adds the structureId to the get universe structures structure id params
func (o *GetUniverseStructuresStructureIDParams) SetStructureID(structureID int64) {
	o.StructureID = structureID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseStructuresStructureIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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

	// path param structure_id
	if err := r.SetPathParam("structure_id", swag.FormatInt64(o.StructureID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
