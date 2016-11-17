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
)

// NewGetCharactersCharacterIDMailParams creates a new GetCharactersCharacterIDMailParams object
// with the default values initialized.
func NewGetCharactersCharacterIDMailParams() *GetCharactersCharacterIDMailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDMailParamsWithTimeout creates a new GetCharactersCharacterIDMailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDMailParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDMailParamsWithContext creates a new GetCharactersCharacterIDMailParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDMailParamsWithContext(ctx context.Context) *GetCharactersCharacterIDMailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersCharacterIDMailParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

/*GetCharactersCharacterIDMailParams contains all the parameters to send to the API endpoint
for the get characters character id mail operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDMailParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Labels
	  Fetch only mails that match one or more of the given labels

	*/
	Labels []int64
	/*LastMailID
	  List only mail with an ID lower than the given ID, if present

	*/
	LastMailID *int32

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDMailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) WithContext(ctx context.Context) *GetCharactersCharacterIDMailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCharacterID adds the characterID to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDMailParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) WithDatasource(datasource *string) *GetCharactersCharacterIDMailParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLabels adds the labels to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) WithLabels(labels []int64) *GetCharactersCharacterIDMailParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) SetLabels(labels []int64) {
	o.Labels = labels
}

// WithLastMailID adds the lastMailID to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) WithLastMailID(lastMailID *int32) *GetCharactersCharacterIDMailParams {
	o.SetLastMailID(lastMailID)
	return o
}

// SetLastMailID adds the lastMailId to the get characters character id mail params
func (o *GetCharactersCharacterIDMailParams) SetLastMailID(lastMailID *int32) {
	o.LastMailID = lastMailID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDMailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesLabels []string
	for _, v := range o.Labels {
		valuesLabels = append(valuesLabels, swag.FormatInt64(v))
	}

	joinedLabels := swag.JoinByFormat(valuesLabels, "multi")
	// query array param labels
	if err := r.SetQueryParam("labels", joinedLabels...); err != nil {
		return err
	}

	if o.LastMailID != nil {

		// query param last_mail_id
		var qrLastMailID int32
		if o.LastMailID != nil {
			qrLastMailID = *o.LastMailID
		}
		qLastMailID := swag.FormatInt32(qrLastMailID)
		if qLastMailID != "" {
			if err := r.SetQueryParam("last_mail_id", qLastMailID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
