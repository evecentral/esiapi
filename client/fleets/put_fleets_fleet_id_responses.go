// Code generated by go-swagger; DO NOT EDIT.

package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// PutFleetsFleetIDReader is a Reader for the PutFleetsFleetID structure.
type PutFleetsFleetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetsFleetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutFleetsFleetIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutFleetsFleetIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutFleetsFleetIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutFleetsFleetIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutFleetsFleetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewPutFleetsFleetIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutFleetsFleetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPutFleetsFleetIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewPutFleetsFleetIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutFleetsFleetIDNoContent creates a PutFleetsFleetIDNoContent with default headers values
func NewPutFleetsFleetIDNoContent() *PutFleetsFleetIDNoContent {
	return &PutFleetsFleetIDNoContent{}
}

/*PutFleetsFleetIDNoContent handles this case with default header values.

Fleet updated
*/
type PutFleetsFleetIDNoContent struct {
}

func (o *PutFleetsFleetIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdNoContent ", 204)
}

func (o *PutFleetsFleetIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetsFleetIDBadRequest creates a PutFleetsFleetIDBadRequest with default headers values
func NewPutFleetsFleetIDBadRequest() *PutFleetsFleetIDBadRequest {
	return &PutFleetsFleetIDBadRequest{}
}

/*PutFleetsFleetIDBadRequest handles this case with default header values.

Bad request
*/
type PutFleetsFleetIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *PutFleetsFleetIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetsFleetIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDUnauthorized creates a PutFleetsFleetIDUnauthorized with default headers values
func NewPutFleetsFleetIDUnauthorized() *PutFleetsFleetIDUnauthorized {
	return &PutFleetsFleetIDUnauthorized{}
}

/*PutFleetsFleetIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PutFleetsFleetIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PutFleetsFleetIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFleetsFleetIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDForbidden creates a PutFleetsFleetIDForbidden with default headers values
func NewPutFleetsFleetIDForbidden() *PutFleetsFleetIDForbidden {
	return &PutFleetsFleetIDForbidden{}
}

/*PutFleetsFleetIDForbidden handles this case with default header values.

Forbidden
*/
type PutFleetsFleetIDForbidden struct {
	Payload *models.Forbidden
}

func (o *PutFleetsFleetIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDNotFound creates a PutFleetsFleetIDNotFound with default headers values
func NewPutFleetsFleetIDNotFound() *PutFleetsFleetIDNotFound {
	return &PutFleetsFleetIDNotFound{}
}

/*PutFleetsFleetIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PutFleetsFleetIDNotFound struct {
	Payload *models.PutFleetsFleetIDNotFoundBody
}

func (o *PutFleetsFleetIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutFleetsFleetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDEnhanceYourCalm creates a PutFleetsFleetIDEnhanceYourCalm with default headers values
func NewPutFleetsFleetIDEnhanceYourCalm() *PutFleetsFleetIDEnhanceYourCalm {
	return &PutFleetsFleetIDEnhanceYourCalm{}
}

/*PutFleetsFleetIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type PutFleetsFleetIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PutFleetsFleetIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutFleetsFleetIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDInternalServerError creates a PutFleetsFleetIDInternalServerError with default headers values
func NewPutFleetsFleetIDInternalServerError() *PutFleetsFleetIDInternalServerError {
	return &PutFleetsFleetIDInternalServerError{}
}

/*PutFleetsFleetIDInternalServerError handles this case with default header values.

Internal server error
*/
type PutFleetsFleetIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PutFleetsFleetIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDServiceUnavailable creates a PutFleetsFleetIDServiceUnavailable with default headers values
func NewPutFleetsFleetIDServiceUnavailable() *PutFleetsFleetIDServiceUnavailable {
	return &PutFleetsFleetIDServiceUnavailable{}
}

/*PutFleetsFleetIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type PutFleetsFleetIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PutFleetsFleetIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFleetsFleetIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDGatewayTimeout creates a PutFleetsFleetIDGatewayTimeout with default headers values
func NewPutFleetsFleetIDGatewayTimeout() *PutFleetsFleetIDGatewayTimeout {
	return &PutFleetsFleetIDGatewayTimeout{}
}

/*PutFleetsFleetIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type PutFleetsFleetIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PutFleetsFleetIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/][%d] putFleetsFleetIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFleetsFleetIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
