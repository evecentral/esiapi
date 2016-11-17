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
)

// NewGetCharactersCharacterIDMailMailIDParams creates a new GetCharactersCharacterIDMailMailIDParams object
// with the default values initialized.
func NewGetCharactersCharacterIDMailMailIDParams() *GetCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDMailMailIDParamsWithTimeout creates a new GetCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDMailMailIDParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDMailMailIDParamsWithContext creates a new GetCharactersCharacterIDMailMailIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDMailMailIDParamsWithContext(ctx context.Context) *GetCharactersCharacterIDMailMailIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailMailIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDMailMailIDParams contains all the parameters to send to the API endpoint
for the get characters character id mail mail id operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDMailMailIDParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
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

// WithTimeout adds the timeout to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailMailIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithContext(ctx context.Context) *GetCharactersCharacterIDMailMailIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDMailMailIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithDatasource(datasource *string) *GetCharactersCharacterIDMailMailIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMailID adds the mailID to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) WithMailID(mailID int32) *GetCharactersCharacterIDMailMailIDParams {
	o.SetMailID(mailID)
	return o
}

// SetMailID adds the mailId to the get characters character id mail mail id params
func (o *GetCharactersCharacterIDMailMailIDParams) SetMailID(mailID int32) {
	o.MailID = mailID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDMailMailIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
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
