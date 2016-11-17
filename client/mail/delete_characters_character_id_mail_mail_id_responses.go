package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// DeleteCharactersCharacterIDMailMailIDReader is a Reader for the DeleteCharactersCharacterIDMailMailID structure.
type DeleteCharactersCharacterIDMailMailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCharactersCharacterIDMailMailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCharactersCharacterIDMailMailIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteCharactersCharacterIDMailMailIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteCharactersCharacterIDMailMailIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCharactersCharacterIDMailMailIDNoContent creates a DeleteCharactersCharacterIDMailMailIDNoContent with default headers values
func NewDeleteCharactersCharacterIDMailMailIDNoContent() *DeleteCharactersCharacterIDMailMailIDNoContent {
	return &DeleteCharactersCharacterIDMailMailIDNoContent{}
}

/*DeleteCharactersCharacterIDMailMailIDNoContent handles this case with default header values.

Mail deleted
*/
type DeleteCharactersCharacterIDMailMailIDNoContent struct {
}

func (o *DeleteCharactersCharacterIDMailMailIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdNoContent ", 204)
}

func (o *DeleteCharactersCharacterIDMailMailIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDForbidden creates a DeleteCharactersCharacterIDMailMailIDForbidden with default headers values
func NewDeleteCharactersCharacterIDMailMailIDForbidden() *DeleteCharactersCharacterIDMailMailIDForbidden {
	return &DeleteCharactersCharacterIDMailMailIDForbidden{}
}

/*DeleteCharactersCharacterIDMailMailIDForbidden handles this case with default header values.

Forbidden
*/
type DeleteCharactersCharacterIDMailMailIDForbidden struct {
	Payload *models.DeleteCharactersCharacterIDMailMailIDForbidden
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCharactersCharacterIDMailMailIDForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDInternalServerError creates a DeleteCharactersCharacterIDMailMailIDInternalServerError with default headers values
func NewDeleteCharactersCharacterIDMailMailIDInternalServerError() *DeleteCharactersCharacterIDMailMailIDInternalServerError {
	return &DeleteCharactersCharacterIDMailMailIDInternalServerError{}
}

/*DeleteCharactersCharacterIDMailMailIDInternalServerError handles this case with default header values.

Internal server error
*/
type DeleteCharactersCharacterIDMailMailIDInternalServerError struct {
	Payload *models.DeleteCharactersCharacterIDMailMailIDInternalServerError
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCharactersCharacterIDMailMailIDInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
