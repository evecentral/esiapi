package live

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// PostCharactersCharacterIDMailReader is a Reader for the PostCharactersCharacterIDMail structure.
type PostCharactersCharacterIDMailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDMailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCharactersCharacterIDMailCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCharactersCharacterIDMailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostCharactersCharacterIDMailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersCharacterIDMailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersCharacterIDMailCreated creates a PostCharactersCharacterIDMailCreated with default headers values
func NewPostCharactersCharacterIDMailCreated() *PostCharactersCharacterIDMailCreated {
	return &PostCharactersCharacterIDMailCreated{}
}

/*PostCharactersCharacterIDMailCreated handles this case with default header values.

Mail created
*/
type PostCharactersCharacterIDMailCreated struct {
	Payload int32
}

func (o *PostCharactersCharacterIDMailCreated) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailCreated  %+v", 201, o.Payload)
}

func (o *PostCharactersCharacterIDMailCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailBadRequest creates a PostCharactersCharacterIDMailBadRequest with default headers values
func NewPostCharactersCharacterIDMailBadRequest() *PostCharactersCharacterIDMailBadRequest {
	return &PostCharactersCharacterIDMailBadRequest{}
}

/*PostCharactersCharacterIDMailBadRequest handles this case with default header values.

Only one corporation, alliance, or mailing list can be the
recipient of a mail

*/
type PostCharactersCharacterIDMailBadRequest struct {
	Payload *models.PostCharactersCharacterIDMailBadRequest
}

func (o *PostCharactersCharacterIDMailBadRequest) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailBadRequest  %+v", 400, o.Payload)
}

func (o *PostCharactersCharacterIDMailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersCharacterIDMailBadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailForbidden creates a PostCharactersCharacterIDMailForbidden with default headers values
func NewPostCharactersCharacterIDMailForbidden() *PostCharactersCharacterIDMailForbidden {
	return &PostCharactersCharacterIDMailForbidden{}
}

/*PostCharactersCharacterIDMailForbidden handles this case with default header values.

Forbidden
*/
type PostCharactersCharacterIDMailForbidden struct {
	Payload *models.PostCharactersCharacterIDMailForbidden
}

func (o *PostCharactersCharacterIDMailForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDMailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersCharacterIDMailForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailInternalServerError creates a PostCharactersCharacterIDMailInternalServerError with default headers values
func NewPostCharactersCharacterIDMailInternalServerError() *PostCharactersCharacterIDMailInternalServerError {
	return &PostCharactersCharacterIDMailInternalServerError{}
}

/*PostCharactersCharacterIDMailInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersCharacterIDMailInternalServerError struct {
	Payload *models.PostCharactersCharacterIDMailInternalServerError
}

func (o *PostCharactersCharacterIDMailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDMailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersCharacterIDMailInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
