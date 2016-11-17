package character

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

// NewGetCharactersCharacterIDCorporationhistoryParams creates a new GetCharactersCharacterIDCorporationhistoryParams object
// with the default values initialized.
func NewGetCharactersCharacterIDCorporationhistoryParams() *GetCharactersCharacterIDCorporationhistoryParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDCorporationhistoryParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDCorporationhistoryParamsWithTimeout creates a new GetCharactersCharacterIDCorporationhistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDCorporationhistoryParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDCorporationhistoryParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDCorporationhistoryParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDCorporationhistoryParamsWithContext creates a new GetCharactersCharacterIDCorporationhistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDCorporationhistoryParamsWithContext(ctx context.Context) *GetCharactersCharacterIDCorporationhistoryParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDCorporationhistoryParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDCorporationhistoryParams contains all the parameters to send to the API endpoint
for the get characters character id corporationhistory operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDCorporationhistoryParams struct {

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

// WithTimeout adds the timeout to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithContext(ctx context.Context) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) WithDatasource(datasource *string) *GetCharactersCharacterIDCorporationhistoryParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id corporationhistory params
func (o *GetCharactersCharacterIDCorporationhistoryParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDCorporationhistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
