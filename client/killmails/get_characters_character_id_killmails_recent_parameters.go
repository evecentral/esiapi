package killmails

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

// NewGetCharactersCharacterIDKillmailsRecentParams creates a new GetCharactersCharacterIDKillmailsRecentParams object
// with the default values initialized.
func NewGetCharactersCharacterIDKillmailsRecentParams() *GetCharactersCharacterIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		maxCountDefault   = int32(50)
	)
	return &GetCharactersCharacterIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		MaxCount:   &maxCountDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDKillmailsRecentParamsWithTimeout creates a new GetCharactersCharacterIDKillmailsRecentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDKillmailsRecentParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		maxCountDefault   = int32(50)
	)
	return &GetCharactersCharacterIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		MaxCount:   &maxCountDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDKillmailsRecentParamsWithContext creates a new GetCharactersCharacterIDKillmailsRecentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDKillmailsRecentParamsWithContext(ctx context.Context) *GetCharactersCharacterIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		maxCountDefault   = int32(50)
	)
	return &GetCharactersCharacterIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		MaxCount:   &maxCountDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDKillmailsRecentParams contains all the parameters to send to the API endpoint
for the get characters character id killmails recent operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDKillmailsRecentParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*MaxCount
	  How many killmails to return at maximum

	*/
	MaxCount *int32
	/*MaxKillID
	  Only return killmails with ID smaller than this.


	*/
	MaxKillID *int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithContext(ctx context.Context) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithDatasource(datasource *string) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithMaxCount adds the maxCount to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithMaxCount(maxCount *int32) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetMaxCount(maxCount)
	return o
}

// SetMaxCount adds the maxCount to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetMaxCount(maxCount *int32) {
	o.MaxCount = maxCount
}

// WithMaxKillID adds the maxKillID to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) WithMaxKillID(maxKillID *int32) *GetCharactersCharacterIDKillmailsRecentParams {
	o.SetMaxKillID(maxKillID)
	return o
}

// SetMaxKillID adds the maxKillId to the get characters character id killmails recent params
func (o *GetCharactersCharacterIDKillmailsRecentParams) SetMaxKillID(maxKillID *int32) {
	o.MaxKillID = maxKillID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDKillmailsRecentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MaxCount != nil {

		// query param max_count
		var qrMaxCount int32
		if o.MaxCount != nil {
			qrMaxCount = *o.MaxCount
		}
		qMaxCount := swag.FormatInt32(qrMaxCount)
		if qMaxCount != "" {
			if err := r.SetQueryParam("max_count", qMaxCount); err != nil {
				return err
			}
		}

	}

	if o.MaxKillID != nil {

		// query param max_kill_id
		var qrMaxKillID int32
		if o.MaxKillID != nil {
			qrMaxKillID = *o.MaxKillID
		}
		qMaxKillID := swag.FormatInt32(qrMaxKillID)
		if qMaxKillID != "" {
			if err := r.SetQueryParam("max_kill_id", qMaxKillID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
