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

// NewGetCorporationsCorporationIDWalletsWalletIDTransactionsParams creates a new GetCorporationsCorporationIDWalletsWalletIDTransactionsParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDWalletsWalletIDTransactionsParams() *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDWalletsWalletIDTransactionsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDWalletsWalletIDTransactionsParamsWithTimeout creates a new GetCorporationsCorporationIDWalletsWalletIDTransactionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDWalletsWalletIDTransactionsParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDWalletsWalletIDTransactionsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDWalletsWalletIDTransactionsParamsWithContext creates a new GetCorporationsCorporationIDWalletsWalletIDTransactionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDWalletsWalletIDTransactionsParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDWalletsWalletIDTransactionsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCorporationsCorporationIDWalletsWalletIDTransactionsParams contains all the parameters to send to the API endpoint
for the get corporations corporation id wallets wallet id transactions operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDWalletsWalletIDTransactionsParams struct {

	/*CorporationID
	  An EVE corporation ID

	*/
	CorporationID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*WalletID
	  Wallet ID

	*/
	WalletID int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCorporationID adds the corporationID to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithWalletID adds the walletID to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) WithWalletID(walletID int32) *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams {
	o.SetWalletID(walletID)
	return o
}

// SetWalletID adds the walletId to the get corporations corporation id wallets wallet id transactions params
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) SetWalletID(walletID int32) {
	o.WalletID = walletID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param wallet_id
	if err := r.SetPathParam("wallet_id", swag.FormatInt32(o.WalletID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
