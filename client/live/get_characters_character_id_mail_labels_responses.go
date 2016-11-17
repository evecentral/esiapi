package live

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDMailLabelsReader is a Reader for the GetCharactersCharacterIDMailLabels structure.
type GetCharactersCharacterIDMailLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMailLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMailLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDMailLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMailLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMailLabelsOK creates a GetCharactersCharacterIDMailLabelsOK with default headers values
func NewGetCharactersCharacterIDMailLabelsOK() *GetCharactersCharacterIDMailLabelsOK {
	return &GetCharactersCharacterIDMailLabelsOK{}
}

/*GetCharactersCharacterIDMailLabelsOK handles this case with default header values.

A list of mail labels and unread counts
*/
type GetCharactersCharacterIDMailLabelsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDMailLabelsOk
}

func (o *GetCharactersCharacterIDMailLabelsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDMailLabelsOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsForbidden creates a GetCharactersCharacterIDMailLabelsForbidden with default headers values
func NewGetCharactersCharacterIDMailLabelsForbidden() *GetCharactersCharacterIDMailLabelsForbidden {
	return &GetCharactersCharacterIDMailLabelsForbidden{}
}

/*GetCharactersCharacterIDMailLabelsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMailLabelsForbidden struct {
	Payload *models.GetCharactersCharacterIDMailLabelsForbidden
}

func (o *GetCharactersCharacterIDMailLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDMailLabelsForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsInternalServerError creates a GetCharactersCharacterIDMailLabelsInternalServerError with default headers values
func NewGetCharactersCharacterIDMailLabelsInternalServerError() *GetCharactersCharacterIDMailLabelsInternalServerError {
	return &GetCharactersCharacterIDMailLabelsInternalServerError{}
}

/*GetCharactersCharacterIDMailLabelsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMailLabelsInternalServerError struct {
	Payload *models.GetCharactersCharacterIDMailLabelsInternalServerError
}

func (o *GetCharactersCharacterIDMailLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDMailLabelsInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
