// Code generated by go-swagger; DO NOT EDIT.

package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
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

	case 401:
		result := NewPostCharactersCharacterIDMailUnauthorized()
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

	case 420:
		result := NewPostCharactersCharacterIDMailEnhanceYourCalm()
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

	case 503:
		result := NewPostCharactersCharacterIDMailServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewPostCharactersCharacterIDMailGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 520:
		result := NewPostCharactersCharacterIDMail()
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

Bad request
*/
type PostCharactersCharacterIDMailBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostCharactersCharacterIDMailBadRequest) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailBadRequest  %+v", 400, o.Payload)
}

func (o *PostCharactersCharacterIDMailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailUnauthorized creates a PostCharactersCharacterIDMailUnauthorized with default headers values
func NewPostCharactersCharacterIDMailUnauthorized() *PostCharactersCharacterIDMailUnauthorized {
	return &PostCharactersCharacterIDMailUnauthorized{}
}

/*PostCharactersCharacterIDMailUnauthorized handles this case with default header values.

Unauthorized
*/
type PostCharactersCharacterIDMailUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostCharactersCharacterIDMailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCharactersCharacterIDMailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

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
	Payload *models.Forbidden
}

func (o *PostCharactersCharacterIDMailForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDMailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailEnhanceYourCalm creates a PostCharactersCharacterIDMailEnhanceYourCalm with default headers values
func NewPostCharactersCharacterIDMailEnhanceYourCalm() *PostCharactersCharacterIDMailEnhanceYourCalm {
	return &PostCharactersCharacterIDMailEnhanceYourCalm{}
}

/*PostCharactersCharacterIDMailEnhanceYourCalm handles this case with default header values.

Error limited
*/
type PostCharactersCharacterIDMailEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostCharactersCharacterIDMailEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostCharactersCharacterIDMailEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

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
	Payload *models.InternalServerError
}

func (o *PostCharactersCharacterIDMailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDMailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailServiceUnavailable creates a PostCharactersCharacterIDMailServiceUnavailable with default headers values
func NewPostCharactersCharacterIDMailServiceUnavailable() *PostCharactersCharacterIDMailServiceUnavailable {
	return &PostCharactersCharacterIDMailServiceUnavailable{}
}

/*PostCharactersCharacterIDMailServiceUnavailable handles this case with default header values.

Service unavailable
*/
type PostCharactersCharacterIDMailServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostCharactersCharacterIDMailServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCharactersCharacterIDMailServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMailGatewayTimeout creates a PostCharactersCharacterIDMailGatewayTimeout with default headers values
func NewPostCharactersCharacterIDMailGatewayTimeout() *PostCharactersCharacterIDMailGatewayTimeout {
	return &PostCharactersCharacterIDMailGatewayTimeout{}
}

/*PostCharactersCharacterIDMailGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type PostCharactersCharacterIDMailGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostCharactersCharacterIDMailGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMailGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCharactersCharacterIDMailGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDMail creates a PostCharactersCharacterIDMail with default headers values
func NewPostCharactersCharacterIDMail() *PostCharactersCharacterIDMail {
	return &PostCharactersCharacterIDMail{}
}

/*PostCharactersCharacterIDMail handles this case with default header values.

Internal error thrown from the EVE server. Most of the time this means you have hit an EVE server rate limit.
*/
type PostCharactersCharacterIDMail struct {
	Payload *models.PostCharactersCharacterIDMailBody
}

func (o *PostCharactersCharacterIDMail) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/mail/][%d] postCharactersCharacterIdMail  %+v", 520, o.Payload)
}

func (o *PostCharactersCharacterIDMail) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostCharactersCharacterIDMailBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
