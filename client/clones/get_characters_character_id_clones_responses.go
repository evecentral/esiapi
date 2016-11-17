package clones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCharactersCharacterIDClonesReader is a Reader for the GetCharactersCharacterIDClones structure.
type GetCharactersCharacterIDClonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDClonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDClonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDClonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDClonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDClonesOK creates a GetCharactersCharacterIDClonesOK with default headers values
func NewGetCharactersCharacterIDClonesOK() *GetCharactersCharacterIDClonesOK {
	return &GetCharactersCharacterIDClonesOK{}
}

/*GetCharactersCharacterIDClonesOK handles this case with default header values.

Clone information for the given character
*/
type GetCharactersCharacterIDClonesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDClonesOk
}

func (o *GetCharactersCharacterIDClonesOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/clones/][%d] getCharactersCharacterIdClonesOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDClonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDClonesOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDClonesForbidden creates a GetCharactersCharacterIDClonesForbidden with default headers values
func NewGetCharactersCharacterIDClonesForbidden() *GetCharactersCharacterIDClonesForbidden {
	return &GetCharactersCharacterIDClonesForbidden{}
}

/*GetCharactersCharacterIDClonesForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDClonesForbidden struct {
	Payload *models.GetCharactersCharacterIDClonesForbidden
}

func (o *GetCharactersCharacterIDClonesForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/clones/][%d] getCharactersCharacterIdClonesForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDClonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDClonesForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDClonesInternalServerError creates a GetCharactersCharacterIDClonesInternalServerError with default headers values
func NewGetCharactersCharacterIDClonesInternalServerError() *GetCharactersCharacterIDClonesInternalServerError {
	return &GetCharactersCharacterIDClonesInternalServerError{}
}

/*GetCharactersCharacterIDClonesInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDClonesInternalServerError struct {
	Payload *models.GetCharactersCharacterIDClonesInternalServerError
}

func (o *GetCharactersCharacterIDClonesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/clones/][%d] getCharactersCharacterIdClonesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDClonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDClonesInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
