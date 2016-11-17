package alliance

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

// NewGetAlliancesAllianceIDIconsParams creates a new GetAlliancesAllianceIDIconsParams object
// with the default values initialized.
func NewGetAlliancesAllianceIDIconsParams() *GetAlliancesAllianceIDIconsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesAllianceIDIconsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlliancesAllianceIDIconsParamsWithTimeout creates a new GetAlliancesAllianceIDIconsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlliancesAllianceIDIconsParamsWithTimeout(timeout time.Duration) *GetAlliancesAllianceIDIconsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesAllianceIDIconsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetAlliancesAllianceIDIconsParamsWithContext creates a new GetAlliancesAllianceIDIconsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlliancesAllianceIDIconsParamsWithContext(ctx context.Context) *GetAlliancesAllianceIDIconsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesAllianceIDIconsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetAlliancesAllianceIDIconsParams contains all the parameters to send to the API endpoint
for the get alliances alliance id icons operation typically these are written to a http.Request
*/
type GetAlliancesAllianceIDIconsParams struct {

	/*AllianceID
	  An EVE alliance ID

	*/
	AllianceID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithTimeout(timeout time.Duration) *GetAlliancesAllianceIDIconsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithContext(ctx context.Context) *GetAlliancesAllianceIDIconsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAllianceID adds the allianceID to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithAllianceID(allianceID int32) *GetAlliancesAllianceIDIconsParams {
	o.SetAllianceID(allianceID)
	return o
}

// SetAllianceID adds the allianceId to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetAllianceID(allianceID int32) {
	o.AllianceID = allianceID
}

// WithDatasource adds the datasource to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) WithDatasource(datasource *string) *GetAlliancesAllianceIDIconsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get alliances alliance id icons params
func (o *GetAlliancesAllianceIDIconsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlliancesAllianceIDIconsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param alliance_id
	if err := r.SetPathParam("alliance_id", swag.FormatInt32(o.AllianceID)); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
