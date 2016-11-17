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

// NewGetCharactersCharacterIDCalendarEventIDParams creates a new GetCharactersCharacterIDCalendarEventIDParams object
// with the default values initialized.
func NewGetCharactersCharacterIDCalendarEventIDParams() *GetCharactersCharacterIDCalendarEventIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDParamsWithTimeout creates a new GetCharactersCharacterIDCalendarEventIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDCalendarEventIDParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDCalendarEventIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDCalendarEventIDParamsWithContext creates a new GetCharactersCharacterIDCalendarEventIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDCalendarEventIDParamsWithContext(ctx context.Context) *GetCharactersCharacterIDCalendarEventIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDCalendarEventIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDCalendarEventIDParams contains all the parameters to send to the API endpoint
for the get characters character id calendar event id operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDCalendarEventIDParams struct {

	/*CharacterID
	  The character id requesting the event

	*/
	CharacterID int64
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*EventID
	  The id of the event requested

	*/
	EventID int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithContext(ctx context.Context) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithCharacterID(characterID int64) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetCharacterID(characterID int64) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithDatasource(datasource *string) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithEventID adds the eventID to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) WithEventID(eventID int32) *GetCharactersCharacterIDCalendarEventIDParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get characters character id calendar event id params
func (o *GetCharactersCharacterIDCalendarEventIDParams) SetEventID(eventID int32) {
	o.EventID = eventID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDCalendarEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt64(o.CharacterID)); err != nil {
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

	// path param event_id
	if err := r.SetPathParam("event_id", swag.FormatInt32(o.EventID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
