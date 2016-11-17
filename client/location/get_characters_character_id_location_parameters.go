package location

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

// NewGetCharactersCharacterIDLocationParams creates a new GetCharactersCharacterIDLocationParams object
// with the default values initialized.
func NewGetCharactersCharacterIDLocationParams() *GetCharactersCharacterIDLocationParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDLocationParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDLocationParamsWithTimeout creates a new GetCharactersCharacterIDLocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDLocationParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDLocationParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDLocationParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDLocationParamsWithContext creates a new GetCharactersCharacterIDLocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDLocationParamsWithContext(ctx context.Context) *GetCharactersCharacterIDLocationParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDLocationParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDLocationParams contains all the parameters to send to the API endpoint
for the get characters character id location operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDLocationParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) WithContext(ctx context.Context) *GetCharactersCharacterIDLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDLocationParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) WithDatasource(datasource *string) *GetCharactersCharacterIDLocationParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id location params
func (o *GetCharactersCharacterIDLocationParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
