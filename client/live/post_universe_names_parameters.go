package live

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// NewPostUniverseNamesParams creates a new PostUniverseNamesParams object
// with the default values initialized.
func NewPostUniverseNamesParams() *PostUniverseNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUniverseNamesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUniverseNamesParamsWithTimeout creates a new PostUniverseNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUniverseNamesParamsWithTimeout(timeout time.Duration) *PostUniverseNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUniverseNamesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostUniverseNamesParamsWithContext creates a new PostUniverseNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUniverseNamesParamsWithContext(ctx context.Context) *PostUniverseNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUniverseNamesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*PostUniverseNamesParams contains all the parameters to send to the API endpoint
for the post universe names operation typically these are written to a http.Request
*/
type PostUniverseNamesParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Ids
	  The ids to resolve

	*/
	Ids *models.PostUniverseNamesIds

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post universe names params
func (o *PostUniverseNamesParams) WithTimeout(timeout time.Duration) *PostUniverseNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post universe names params
func (o *PostUniverseNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post universe names params
func (o *PostUniverseNamesParams) WithContext(ctx context.Context) *PostUniverseNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post universe names params
func (o *PostUniverseNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDatasource adds the datasource to the post universe names params
func (o *PostUniverseNamesParams) WithDatasource(datasource *string) *PostUniverseNamesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post universe names params
func (o *PostUniverseNamesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithIds adds the ids to the post universe names params
func (o *PostUniverseNamesParams) WithIds(ids *models.PostUniverseNamesIds) *PostUniverseNamesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the post universe names params
func (o *PostUniverseNamesParams) SetIds(ids *models.PostUniverseNamesIds) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *PostUniverseNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Ids == nil {
		o.Ids = new(models.PostUniverseNamesIds)
	}

	if err := r.SetBodyParam(o.Ids); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
