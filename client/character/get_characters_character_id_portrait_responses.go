package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCharactersCharacterIDPortraitReader is a Reader for the GetCharactersCharacterIDPortrait structure.
type GetCharactersCharacterIDPortraitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDPortraitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDPortraitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCharactersCharacterIDPortraitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDPortraitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDPortraitOK creates a GetCharactersCharacterIDPortraitOK with default headers values
func NewGetCharactersCharacterIDPortraitOK() *GetCharactersCharacterIDPortraitOK {
	return &GetCharactersCharacterIDPortraitOK{}
}

/*GetCharactersCharacterIDPortraitOK handles this case with default header values.

Public data for the given character
*/
type GetCharactersCharacterIDPortraitOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDPortraitOk
}

func (o *GetCharactersCharacterIDPortraitOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDPortraitOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitNotFound creates a GetCharactersCharacterIDPortraitNotFound with default headers values
func NewGetCharactersCharacterIDPortraitNotFound() *GetCharactersCharacterIDPortraitNotFound {
	return &GetCharactersCharacterIDPortraitNotFound{}
}

/*GetCharactersCharacterIDPortraitNotFound handles this case with default header values.

No image server for this datasource
*/
type GetCharactersCharacterIDPortraitNotFound struct {
	Payload *models.GetCharactersCharacterIDPortraitNotFound
}

func (o *GetCharactersCharacterIDPortraitNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDPortraitNotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitInternalServerError creates a GetCharactersCharacterIDPortraitInternalServerError with default headers values
func NewGetCharactersCharacterIDPortraitInternalServerError() *GetCharactersCharacterIDPortraitInternalServerError {
	return &GetCharactersCharacterIDPortraitInternalServerError{}
}

/*GetCharactersCharacterIDPortraitInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDPortraitInternalServerError struct {
	Payload *models.GetCharactersCharacterIDPortraitInternalServerError
}

func (o *GetCharactersCharacterIDPortraitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDPortraitInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
