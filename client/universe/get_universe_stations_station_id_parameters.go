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

// NewGetUniverseStationsStationIDParams creates a new GetUniverseStationsStationIDParams object
// with the default values initialized.
func NewGetUniverseStationsStationIDParams() *GetUniverseStationsStationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStationsStationIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseStationsStationIDParamsWithTimeout creates a new GetUniverseStationsStationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseStationsStationIDParamsWithTimeout(timeout time.Duration) *GetUniverseStationsStationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStationsStationIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseStationsStationIDParamsWithContext creates a new GetUniverseStationsStationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseStationsStationIDParamsWithContext(ctx context.Context) *GetUniverseStationsStationIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStationsStationIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetUniverseStationsStationIDParams contains all the parameters to send to the API endpoint
for the get universe stations station id operation typically these are written to a http.Request
*/
type GetUniverseStationsStationIDParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*StationID
	  An Eve station ID

	*/
	StationID int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithTimeout(timeout time.Duration) *GetUniverseStationsStationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithContext(ctx context.Context) *GetUniverseStationsStationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithDatasource(datasource *string) *GetUniverseStationsStationIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithStationID adds the stationID to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) WithStationID(stationID int32) *GetUniverseStationsStationIDParams {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the get universe stations station id params
func (o *GetUniverseStationsStationIDParams) SetStationID(stationID int32) {
	o.StationID = stationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseStationsStationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param station_id
	if err := r.SetPathParam("station_id", swag.FormatInt32(o.StationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
