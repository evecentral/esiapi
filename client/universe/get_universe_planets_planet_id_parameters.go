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

// NewGetUniversePlanetsPlanetIDParams creates a new GetUniversePlanetsPlanetIDParams object
// with the default values initialized.
func NewGetUniversePlanetsPlanetIDParams() *GetUniversePlanetsPlanetIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniversePlanetsPlanetIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniversePlanetsPlanetIDParamsWithTimeout creates a new GetUniversePlanetsPlanetIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniversePlanetsPlanetIDParamsWithTimeout(timeout time.Duration) *GetUniversePlanetsPlanetIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniversePlanetsPlanetIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniversePlanetsPlanetIDParamsWithContext creates a new GetUniversePlanetsPlanetIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniversePlanetsPlanetIDParamsWithContext(ctx context.Context) *GetUniversePlanetsPlanetIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniversePlanetsPlanetIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetUniversePlanetsPlanetIDParams contains all the parameters to send to the API endpoint
for the get universe planets planet id operation typically these are written to a http.Request
*/
type GetUniversePlanetsPlanetIDParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*PlanetID
	  An Eve planet ID

	*/
	PlanetID int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) WithTimeout(timeout time.Duration) *GetUniversePlanetsPlanetIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) WithContext(ctx context.Context) *GetUniversePlanetsPlanetIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) WithDatasource(datasource *string) *GetUniversePlanetsPlanetIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPlanetID adds the planetID to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) WithPlanetID(planetID int32) *GetUniversePlanetsPlanetIDParams {
	o.SetPlanetID(planetID)
	return o
}

// SetPlanetID adds the planetId to the get universe planets planet id params
func (o *GetUniversePlanetsPlanetIDParams) SetPlanetID(planetID int32) {
	o.PlanetID = planetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniversePlanetsPlanetIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param planet_id
	if err := r.SetPathParam("planet_id", swag.FormatInt32(o.PlanetID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
