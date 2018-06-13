// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCharactersCharacterIDMailParamsBody post_characters_character_id_mail_mail
//
// mail object
// swagger:model postCharactersCharacterIdMailParamsBody
type PostCharactersCharacterIDMailParamsBody struct {

	// post_characters_character_id_mail_approved_cost
	//
	// approved_cost integer
	ApprovedCost int64 `json:"approved_cost,omitempty"`

	// post_characters_character_id_mail_body
	//
	// body string
	// Required: true
	// Max Length: 10000
	Body *string `json:"body"`

	// post_characters_character_id_mail_recipients
	//
	// recipients array
	// Required: true
	// Max Items: 50
	// Min Items: 1
	Recipients []*PostCharactersCharacterIDMailParamsBodyRecipientsItems `json:"recipients"`

	// post_characters_character_id_mail_subject
	//
	// subject string
	// Required: true
	// Max Length: 1000
	Subject *string `json:"subject"`
}

// Validate validates this post characters character Id mail params body
func (m *PostCharactersCharacterIDMailParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCharactersCharacterIDMailParamsBody) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	if err := validate.MaxLength("body", "body", string(*m.Body), 10000); err != nil {
		return err
	}

	return nil
}

func (m *PostCharactersCharacterIDMailParamsBody) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("recipients", "body", m.Recipients); err != nil {
		return err
	}

	iRecipientsSize := int64(len(m.Recipients))

	if err := validate.MinItems("recipients", "body", iRecipientsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("recipients", "body", iRecipientsSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(m.Recipients); i++ {
		if swag.IsZero(m.Recipients[i]) { // not required
			continue
		}

		if m.Recipients[i] != nil {
			if err := m.Recipients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostCharactersCharacterIDMailParamsBody) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	if err := validate.MaxLength("subject", "body", string(*m.Subject), 1000); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDMailParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDMailParamsBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDMailParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
