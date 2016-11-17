package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDMailMailIDOk 200 ok object
// swagger:model get_characters_character_id_mail_mail_id_ok
type GetCharactersCharacterIDMailMailIDOk struct {

	// get_characters_character_id_mail_mail_id_body
	//
	// Mail's body
	Body string `json:"body,omitempty"`

	// get_characters_character_id_mail_mail_id_from
	//
	// From whom the mail was sent
	From int32 `json:"from,omitempty"`

	// get_characters_character_id_mail_mail_id_labels
	//
	// Labels attached to the mail
	Labels []int64 `json:"labels"`

	// get_characters_character_id_mail_mail_id_read
	//
	// Whether the mail is flagged as read
	Read bool `json:"read,omitempty"`

	// get_characters_character_id_mail_mail_id_recipients
	//
	// Recipients of the mail
	Recipients []*GetCharactersCharacterIDMailMailIDRecipient `json:"recipients"`

	// get_characters_character_id_mail_mail_id_subject
	//
	// Mail subject
	Subject string `json:"subject,omitempty"`

	// get_characters_character_id_mail_mail_id_timestamp
	//
	// When the mail was sent
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this get characters character id mail mail id ok
func (m *GetCharactersCharacterIDMailMailIDOk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDMailMailIDOk) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {

		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("labels"+"."+strconv.Itoa(i), "body", int64(m.Labels[i]), 0, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *GetCharactersCharacterIDMailMailIDOk) validateRecipients(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipients) { // not required
		return nil
	}

	for i := 0; i < len(m.Recipients); i++ {

		if swag.IsZero(m.Recipients[i]) { // not required
			continue
		}

		if m.Recipients[i] != nil {

			if err := m.Recipients[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}