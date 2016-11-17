package dummy

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

// NewGetCorporationsCorporationIDBookmarksFoldersParams creates a new GetCorporationsCorporationIDBookmarksFoldersParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDBookmarksFoldersParams() *GetCorporationsCorporationIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDBookmarksFoldersParamsWithTimeout creates a new GetCorporationsCorporationIDBookmarksFoldersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDBookmarksFoldersParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDBookmarksFoldersParamsWithContext creates a new GetCorporationsCorporationIDBookmarksFoldersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDBookmarksFoldersParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDBookmarksFoldersParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDBookmarksFoldersParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCorporationsCorporationIDBookmarksFoldersParams contains all the parameters to send to the API endpoint
for the get corporations corporation id bookmarks folders operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDBookmarksFoldersParams struct {

	/*CorporationID
	  An EVE corporation ID

	*/
	CorporationID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDBookmarksFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDBookmarksFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCorporationID adds the corporationID to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDBookmarksFoldersParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDBookmarksFoldersParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id bookmarks folders params
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDBookmarksFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
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
