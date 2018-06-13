// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDContractsContractIDItemsReader is a Reader for the GetCharactersCharacterIDContractsContractIDItems structure.
type GetCharactersCharacterIDContractsContractIDItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContractsContractIDItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDContractsContractIDItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDContractsContractIDItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContractsContractIDItemsOK creates a GetCharactersCharacterIDContractsContractIDItemsOK with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsOK() *GetCharactersCharacterIDContractsContractIDItemsOK {
	return &GetCharactersCharacterIDContractsContractIDItemsOK{}
}

/*GetCharactersCharacterIDContractsContractIDItemsOK handles this case with default header values.

A list of contracts
*/
type GetCharactersCharacterIDContractsContractIDItemsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsForbidden creates a GetCharactersCharacterIDContractsContractIDItemsForbidden with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsForbidden() *GetCharactersCharacterIDContractsContractIDItemsForbidden {
	return &GetCharactersCharacterIDContractsContractIDItemsForbidden{}
}

/*GetCharactersCharacterIDContractsContractIDItemsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDContractsContractIDItemsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError creates a GetCharactersCharacterIDContractsContractIDItemsInternalServerError with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError() *GetCharactersCharacterIDContractsContractIDItemsInternalServerError {
	return &GetCharactersCharacterIDContractsContractIDItemsInternalServerError{}
}

/*GetCharactersCharacterIDContractsContractIDItemsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDContractsContractIDItemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 get_characters_character_id_contracts_contract_id_items_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0
*/

type GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 struct {

	// get_characters_character_id_contracts_contract_id_items_is_included
	//
	// true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract.
	// Required: true
	IsIncluded *bool `json:"is_included"`

	// get_characters_character_id_contracts_contract_id_items_is_singleton
	//
	// is_singleton boolean
	// Required: true
	IsSingleton *bool `json:"is_singleton"`

	// get_characters_character_id_contracts_contract_id_items_quantity
	//
	// Number of items in the stack
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_characters_character_id_contracts_contract_id_items_raw_quantity
	//
	// -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy
	RawQuantity int32 `json:"raw_quantity,omitempty"`

	// get_characters_character_id_contracts_contract_id_items_record_id
	//
	// Unique ID for the item
	// Required: true
	RecordID *int64 `json:"record_id"`

	// get_characters_character_id_contracts_contract_id_items_type_id
	//
	// Type ID for item
	// Required: true
	TypeID *int32 `json:"type_id"`
}

/* polymorph GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 is_included false */

/* polymorph GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 is_singleton false */

/* polymorph GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 quantity false */

/* polymorph GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 raw_quantity false */

/* polymorph GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 record_id false */

/* polymorph GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 type_id false */

// Validate validates this get characters character ID contracts contract ID items o k body items0
func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsIncluded(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateIsSingleton(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRecordID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateIsIncluded(formats strfmt.Registry) error {

	if err := validate.Required("is_included", "body", o.IsIncluded); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateIsSingleton(formats strfmt.Registry) error {

	if err := validate.Required("is_singleton", "body", o.IsSingleton); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateRecordID(formats strfmt.Registry) error {

	if err := validate.Required("record_id", "body", o.RecordID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}