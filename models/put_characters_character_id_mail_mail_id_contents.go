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

// PutCharactersCharacterIDMailMailIDContents put_characters_character_id_mail_mail_id_contents
//
// contents object
// swagger:model put_characters_character_id_mail_mail_id_contents
type PutCharactersCharacterIDMailMailIDContents struct {

	// put_characters_character_id_mail_mail_id_labels
	//
	// Labels to assign to the mail. Pre-existing labels are unassigned.
	Labels []int64 `json:"labels"`

	// put_characters_character_id_mail_mail_id_read
	//
	// Whether the mail is flagged as read
	Read bool `json:"read,omitempty"`
}

// Validate validates this put characters character id mail mail id contents
func (m *PutCharactersCharacterIDMailMailIDContents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutCharactersCharacterIDMailMailIDContents) validateLabels(formats strfmt.Registry) error {

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