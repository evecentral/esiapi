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

	case 400:
		result := NewDeleteCharactersCharacterIDMailMailIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteCharactersCharacterIDMailMailIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteCharactersCharacterIDMailMailIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewDeleteCharactersCharacterIDMailMailIDEnhanceYourCalm()
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

	case 503:
		result := NewDeleteCharactersCharacterIDMailMailIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewDeleteCharactersCharacterIDMailMailIDGatewayTimeout()
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

// NewDeleteCharactersCharacterIDMailMailIDBadRequest creates a DeleteCharactersCharacterIDMailMailIDBadRequest with default headers values
func NewDeleteCharactersCharacterIDMailMailIDBadRequest() *DeleteCharactersCharacterIDMailMailIDBadRequest {
	return &DeleteCharactersCharacterIDMailMailIDBadRequest{}
}

/*DeleteCharactersCharacterIDMailMailIDBadRequest handles this case with default header values.

Bad request
*/
type DeleteCharactersCharacterIDMailMailIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDUnauthorized creates a DeleteCharactersCharacterIDMailMailIDUnauthorized with default headers values
func NewDeleteCharactersCharacterIDMailMailIDUnauthorized() *DeleteCharactersCharacterIDMailMailIDUnauthorized {
	return &DeleteCharactersCharacterIDMailMailIDUnauthorized{}
}

/*DeleteCharactersCharacterIDMailMailIDUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCharactersCharacterIDMailMailIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

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
	Payload *models.Forbidden
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDEnhanceYourCalm creates a DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm with default headers values
func NewDeleteCharactersCharacterIDMailMailIDEnhanceYourCalm() *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm {
	return &DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm{}
}

/*DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

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
	Payload *models.InternalServerError
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDServiceUnavailable creates a DeleteCharactersCharacterIDMailMailIDServiceUnavailable with default headers values
func NewDeleteCharactersCharacterIDMailMailIDServiceUnavailable() *DeleteCharactersCharacterIDMailMailIDServiceUnavailable {
	return &DeleteCharactersCharacterIDMailMailIDServiceUnavailable{}
}

/*DeleteCharactersCharacterIDMailMailIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type DeleteCharactersCharacterIDMailMailIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDGatewayTimeout creates a DeleteCharactersCharacterIDMailMailIDGatewayTimeout with default headers values
func NewDeleteCharactersCharacterIDMailMailIDGatewayTimeout() *DeleteCharactersCharacterIDMailMailIDGatewayTimeout {
	return &DeleteCharactersCharacterIDMailMailIDGatewayTimeout{}
}

/*DeleteCharactersCharacterIDMailMailIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type DeleteCharactersCharacterIDMailMailIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
