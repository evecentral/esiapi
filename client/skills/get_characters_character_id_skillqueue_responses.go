package skills

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCharactersCharacterIDSkillqueueReader is a Reader for the GetCharactersCharacterIDSkillqueue structure.
type GetCharactersCharacterIDSkillqueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDSkillqueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDSkillqueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDSkillqueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDSkillqueueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDSkillqueueOK creates a GetCharactersCharacterIDSkillqueueOK with default headers values
func NewGetCharactersCharacterIDSkillqueueOK() *GetCharactersCharacterIDSkillqueueOK {
	return &GetCharactersCharacterIDSkillqueueOK{}
}

/*GetCharactersCharacterIDSkillqueueOK handles this case with default header values.

The current skill queue, sorted ascending by finishing time
*/
type GetCharactersCharacterIDSkillqueueOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetCharactersCharacterIDSkillqueue200Ok
}

func (o *GetCharactersCharacterIDSkillqueueOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDSkillqueueForbidden creates a GetCharactersCharacterIDSkillqueueForbidden with default headers values
func NewGetCharactersCharacterIDSkillqueueForbidden() *GetCharactersCharacterIDSkillqueueForbidden {
	return &GetCharactersCharacterIDSkillqueueForbidden{}
}

/*GetCharactersCharacterIDSkillqueueForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDSkillqueueForbidden struct {
	Payload *models.GetCharactersCharacterIDSkillqueueForbidden
}

func (o *GetCharactersCharacterIDSkillqueueForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDSkillqueueForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillqueueInternalServerError creates a GetCharactersCharacterIDSkillqueueInternalServerError with default headers values
func NewGetCharactersCharacterIDSkillqueueInternalServerError() *GetCharactersCharacterIDSkillqueueInternalServerError {
	return &GetCharactersCharacterIDSkillqueueInternalServerError{}
}

/*GetCharactersCharacterIDSkillqueueInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDSkillqueueInternalServerError struct {
	Payload *models.GetCharactersCharacterIDSkillqueueInternalServerError
}

func (o *GetCharactersCharacterIDSkillqueueInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDSkillqueueInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
