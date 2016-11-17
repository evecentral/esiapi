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

// GetCharactersCharacterIDReader is a Reader for the GetCharactersCharacterID structure.
type GetCharactersCharacterIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewGetCharactersCharacterIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOK creates a GetCharactersCharacterIDOK with default headers values
func NewGetCharactersCharacterIDOK() *GetCharactersCharacterIDOK {
	return &GetCharactersCharacterIDOK{}
}

/*GetCharactersCharacterIDOK handles this case with default header values.

Public data for the given character
*/
type GetCharactersCharacterIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDOk
}

func (o *GetCharactersCharacterIDOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/][%d] getCharactersCharacterIdOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDUnprocessableEntity creates a GetCharactersCharacterIDUnprocessableEntity with default headers values
func NewGetCharactersCharacterIDUnprocessableEntity() *GetCharactersCharacterIDUnprocessableEntity {
	return &GetCharactersCharacterIDUnprocessableEntity{}
}

/*GetCharactersCharacterIDUnprocessableEntity handles this case with default header values.

Is not a character ID
*/
type GetCharactersCharacterIDUnprocessableEntity struct {
	Payload *models.GetCharactersCharacterIDUnprocessableEntity
}

func (o *GetCharactersCharacterIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/][%d] getCharactersCharacterIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetCharactersCharacterIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDUnprocessableEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDInternalServerError creates a GetCharactersCharacterIDInternalServerError with default headers values
func NewGetCharactersCharacterIDInternalServerError() *GetCharactersCharacterIDInternalServerError {
	return &GetCharactersCharacterIDInternalServerError{}
}

/*GetCharactersCharacterIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDInternalServerError struct {
	Payload *models.GetCharactersCharacterIDInternalServerError
}

func (o *GetCharactersCharacterIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/][%d] getCharactersCharacterIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
