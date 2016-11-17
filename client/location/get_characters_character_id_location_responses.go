package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDLocationReader is a Reader for the GetCharactersCharacterIDLocation structure.
type GetCharactersCharacterIDLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDLocationOK creates a GetCharactersCharacterIDLocationOK with default headers values
func NewGetCharactersCharacterIDLocationOK() *GetCharactersCharacterIDLocationOK {
	return &GetCharactersCharacterIDLocationOK{}
}

/*GetCharactersCharacterIDLocationOK handles this case with default header values.

Information about the characters current location. Returns the current solar system id, and also the current station or structure ID if applicable.
*/
type GetCharactersCharacterIDLocationOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDLocationOk
}

func (o *GetCharactersCharacterIDLocationOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDLocationOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationForbidden creates a GetCharactersCharacterIDLocationForbidden with default headers values
func NewGetCharactersCharacterIDLocationForbidden() *GetCharactersCharacterIDLocationForbidden {
	return &GetCharactersCharacterIDLocationForbidden{}
}

/*GetCharactersCharacterIDLocationForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDLocationForbidden struct {
	Payload *models.GetCharactersCharacterIDLocationForbidden
}

func (o *GetCharactersCharacterIDLocationForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDLocationForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDLocationInternalServerError creates a GetCharactersCharacterIDLocationInternalServerError with default headers values
func NewGetCharactersCharacterIDLocationInternalServerError() *GetCharactersCharacterIDLocationInternalServerError {
	return &GetCharactersCharacterIDLocationInternalServerError{}
}

/*GetCharactersCharacterIDLocationInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDLocationInternalServerError struct {
	Payload *models.GetCharactersCharacterIDLocationInternalServerError
}

func (o *GetCharactersCharacterIDLocationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/location/][%d] getCharactersCharacterIdLocationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDLocationInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
