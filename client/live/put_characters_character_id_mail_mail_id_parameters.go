package live

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

	"github.com/evecentral/esiapi/esiswagger/models"
)

// NewPutCharactersCharacterIDMailMailIDParams creates a new PutCharactersCharacterIDMailMailIDParams object
// with the default values initialized.
func NewPutCharactersCharacterIDMailMailIDParams() *PutCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPutCharactersCharacterIDMailMailIDParamsWithTimeout creates a new PutCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutCharactersCharacterIDMailMailIDParamsWithTimeout(timeout time.Duration) *PutCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPutCharactersCharacterIDMailMailIDParamsWithContext creates a new PutCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutCharactersCharacterIDMailMailIDParamsWithContext(ctx context.Context) *PutCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PutCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*PutCharactersCharacterIDMailMailIDParams contains all the parameters to send to the API endpoint
for the put characters character id mail mail id operation typically these are written to a http.Request
*/
type PutCharactersCharacterIDMailMailIDParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Contents
	  Data used to update the mail

	*/
	Contents *models.PutCharactersCharacterIDMailMailIDContents
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*MailID
	  An EVE mail ID

	*/
	MailID int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithTimeout(timeout time.Duration) *PutCharactersCharacterIDMailMailIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithContext(ctx context.Context) *PutCharactersCharacterIDMailMailIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithCharacterID(characterID int32) *PutCharactersCharacterIDMailMailIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithContents adds the contents to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithContents(contents *models.PutCharactersCharacterIDMailMailIDContents) *PutCharactersCharacterIDMailMailIDParams {
	o.SetContents(contents)
	return o
}

// SetContents adds the contents to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetContents(contents *models.PutCharactersCharacterIDMailMailIDContents) {
	o.Contents = contents
}

// WithDatasource adds the datasource to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithDatasource(datasource *string) *PutCharactersCharacterIDMailMailIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMailID adds the mailID to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) WithMailID(mailID int32) *PutCharactersCharacterIDMailMailIDParams {
	o.SetMailID(mailID)
	return o
}

// SetMailID adds the mailId to the put characters character id mail mail id params
func (o *PutCharactersCharacterIDMailMailIDParams) SetMailID(mailID int32) {
	o.MailID = mailID
}

// WriteToRequest writes these params to a swagger request
func (o *PutCharactersCharacterIDMailMailIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
		return err
	}

	if o.Contents == nil {
		o.Contents = new(models.PutCharactersCharacterIDMailMailIDContents)
	}

	if err := r.SetBodyParam(o.Contents); err != nil {
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

	// path param mail_id
	if err := r.SetPathParam("mail_id", swag.FormatInt32(o.MailID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
