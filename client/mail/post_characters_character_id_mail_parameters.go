package mail

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

// NewPostCharactersCharacterIDMailParams creates a new PostCharactersCharacterIDMailParams object
// with the default values initialized.
func NewPostCharactersCharacterIDMailParams() *PostCharactersCharacterIDMailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDMailParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCharactersCharacterIDMailParamsWithTimeout creates a new PostCharactersCharacterIDMailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCharactersCharacterIDMailParamsWithTimeout(timeout time.Duration) *PostCharactersCharacterIDMailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDMailParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostCharactersCharacterIDMailParamsWithContext creates a new PostCharactersCharacterIDMailParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCharactersCharacterIDMailParamsWithContext(ctx context.Context) *PostCharactersCharacterIDMailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDMailParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*PostCharactersCharacterIDMailParams contains all the parameters to send to the API endpoint
for the post characters character id mail operation typically these are written to a http.Request
*/
type PostCharactersCharacterIDMailParams struct {

	/*CharacterID
	  The sender's character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Mail
	  The mail to send

	*/
	Mail *models.PostCharactersCharacterIDMailMail

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) WithTimeout(timeout time.Duration) *PostCharactersCharacterIDMailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) WithContext(ctx context.Context) *PostCharactersCharacterIDMailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) WithCharacterID(characterID int32) *PostCharactersCharacterIDMailParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) WithDatasource(datasource *string) *PostCharactersCharacterIDMailParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMail adds the mail to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) WithMail(mail *models.PostCharactersCharacterIDMailMail) *PostCharactersCharacterIDMailParams {
	o.SetMail(mail)
	return o
}

// SetMail adds the mail to the post characters character id mail params
func (o *PostCharactersCharacterIDMailParams) SetMail(mail *models.PostCharactersCharacterIDMailMail) {
	o.Mail = mail
}

// WriteToRequest writes these params to a swagger request
func (o *PostCharactersCharacterIDMailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Mail == nil {
		o.Mail = new(models.PostCharactersCharacterIDMailMail)
	}

	if err := r.SetBodyParam(o.Mail); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
