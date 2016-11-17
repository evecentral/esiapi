package wallet

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

// NewGetCharactersCharacterIDWalletsTransactionsParams creates a new GetCharactersCharacterIDWalletsTransactionsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDWalletsTransactionsParams() *GetCharactersCharacterIDWalletsTransactionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDWalletsTransactionsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDWalletsTransactionsParamsWithTimeout creates a new GetCharactersCharacterIDWalletsTransactionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDWalletsTransactionsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDWalletsTransactionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDWalletsTransactionsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDWalletsTransactionsParamsWithContext creates a new GetCharactersCharacterIDWalletsTransactionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDWalletsTransactionsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDWalletsTransactionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDWalletsTransactionsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDWalletsTransactionsParams contains all the parameters to send to the API endpoint
for the get characters character id wallets transactions operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDWalletsTransactionsParams struct {

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

// WithTimeout adds the timeout to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDWalletsTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDWalletsTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDWalletsTransactionsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDWalletsTransactionsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id wallets transactions params
func (o *GetCharactersCharacterIDWalletsTransactionsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDWalletsTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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