// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCharactersCharacterIDMailParamsBodyRecipientsItems post_characters_character_id_mail_recipient
//
// recipient object
// swagger:model postCharactersCharacterIdMailParamsBodyRecipientsItems
type PostCharactersCharacterIDMailParamsBodyRecipientsItems struct {

	// post_characters_character_id_mail_recipient_id
	//
	// recipient_id integer
	// Required: true
	RecipientID *int32 `json:"recipient_id"`

	// post_characters_character_id_mail_recipient_type
	//
	// recipient_type string
	// Required: true
	// Enum: [alliance character corporation mailing_list]
	RecipientType *string `json:"recipient_type"`
}

// Validate validates this post characters character Id mail params body recipients items
func (m *PostCharactersCharacterIDMailParamsBodyRecipientsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipientType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCharactersCharacterIDMailParamsBodyRecipientsItems) validateRecipientID(formats strfmt.Registry) error {

	if err := validate.Required("recipient_id", "body", m.RecipientID); err != nil {
		return err
	}

	return nil
}

var postCharactersCharacterIdMailParamsBodyRecipientsItemsTypeRecipientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance","character","corporation","mailing_list"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCharactersCharacterIdMailParamsBodyRecipientsItemsTypeRecipientTypePropEnum = append(postCharactersCharacterIdMailParamsBodyRecipientsItemsTypeRecipientTypePropEnum, v)
	}
}

const (

	// PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeAlliance captures enum value "alliance"
	PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeAlliance string = "alliance"

	// PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeCharacter captures enum value "character"
	PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeCharacter string = "character"

	// PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeCorporation captures enum value "corporation"
	PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeCorporation string = "corporation"

	// PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeMailingList captures enum value "mailing_list"
	PostCharactersCharacterIDMailParamsBodyRecipientsItemsRecipientTypeMailingList string = "mailing_list"
)

// prop value enum
func (m *PostCharactersCharacterIDMailParamsBodyRecipientsItems) validateRecipientTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postCharactersCharacterIdMailParamsBodyRecipientsItemsTypeRecipientTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostCharactersCharacterIDMailParamsBodyRecipientsItems) validateRecipientType(formats strfmt.Registry) error {

	if err := validate.Required("recipient_type", "body", m.RecipientType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecipientTypeEnum("recipient_type", "body", *m.RecipientType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDMailParamsBodyRecipientsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDMailParamsBodyRecipientsItems) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDMailParamsBodyRecipientsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
