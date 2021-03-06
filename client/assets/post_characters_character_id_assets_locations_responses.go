// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// PostCharactersCharacterIDAssetsLocationsReader is a Reader for the PostCharactersCharacterIDAssetsLocations structure.
type PostCharactersCharacterIDAssetsLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDAssetsLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCharactersCharacterIDAssetsLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCharactersCharacterIDAssetsLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostCharactersCharacterIDAssetsLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostCharactersCharacterIDAssetsLocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewPostCharactersCharacterIDAssetsLocationsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersCharacterIDAssetsLocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostCharactersCharacterIDAssetsLocationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewPostCharactersCharacterIDAssetsLocationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersCharacterIDAssetsLocationsOK creates a PostCharactersCharacterIDAssetsLocationsOK with default headers values
func NewPostCharactersCharacterIDAssetsLocationsOK() *PostCharactersCharacterIDAssetsLocationsOK {
	return &PostCharactersCharacterIDAssetsLocationsOK{}
}

/*PostCharactersCharacterIDAssetsLocationsOK handles this case with default header values.

List of asset locations
*/
type PostCharactersCharacterIDAssetsLocationsOK struct {
	Payload []*models.PostCharactersCharacterIDAssetsLocationsOKBodyItems
}

func (o *PostCharactersCharacterIDAssetsLocationsOK) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsOK  %+v", 200, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsBadRequest creates a PostCharactersCharacterIDAssetsLocationsBadRequest with default headers values
func NewPostCharactersCharacterIDAssetsLocationsBadRequest() *PostCharactersCharacterIDAssetsLocationsBadRequest {
	return &PostCharactersCharacterIDAssetsLocationsBadRequest{}
}

/*PostCharactersCharacterIDAssetsLocationsBadRequest handles this case with default header values.

Bad request
*/
type PostCharactersCharacterIDAssetsLocationsBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsUnauthorized creates a PostCharactersCharacterIDAssetsLocationsUnauthorized with default headers values
func NewPostCharactersCharacterIDAssetsLocationsUnauthorized() *PostCharactersCharacterIDAssetsLocationsUnauthorized {
	return &PostCharactersCharacterIDAssetsLocationsUnauthorized{}
}

/*PostCharactersCharacterIDAssetsLocationsUnauthorized handles this case with default header values.

Unauthorized
*/
type PostCharactersCharacterIDAssetsLocationsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsForbidden creates a PostCharactersCharacterIDAssetsLocationsForbidden with default headers values
func NewPostCharactersCharacterIDAssetsLocationsForbidden() *PostCharactersCharacterIDAssetsLocationsForbidden {
	return &PostCharactersCharacterIDAssetsLocationsForbidden{}
}

/*PostCharactersCharacterIDAssetsLocationsForbidden handles this case with default header values.

Forbidden
*/
type PostCharactersCharacterIDAssetsLocationsForbidden struct {
	Payload *models.Forbidden
}

func (o *PostCharactersCharacterIDAssetsLocationsForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsEnhanceYourCalm creates a PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm with default headers values
func NewPostCharactersCharacterIDAssetsLocationsEnhanceYourCalm() *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm {
	return &PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm{}
}

/*PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsInternalServerError creates a PostCharactersCharacterIDAssetsLocationsInternalServerError with default headers values
func NewPostCharactersCharacterIDAssetsLocationsInternalServerError() *PostCharactersCharacterIDAssetsLocationsInternalServerError {
	return &PostCharactersCharacterIDAssetsLocationsInternalServerError{}
}

/*PostCharactersCharacterIDAssetsLocationsInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersCharacterIDAssetsLocationsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsServiceUnavailable creates a PostCharactersCharacterIDAssetsLocationsServiceUnavailable with default headers values
func NewPostCharactersCharacterIDAssetsLocationsServiceUnavailable() *PostCharactersCharacterIDAssetsLocationsServiceUnavailable {
	return &PostCharactersCharacterIDAssetsLocationsServiceUnavailable{}
}

/*PostCharactersCharacterIDAssetsLocationsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type PostCharactersCharacterIDAssetsLocationsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsGatewayTimeout creates a PostCharactersCharacterIDAssetsLocationsGatewayTimeout with default headers values
func NewPostCharactersCharacterIDAssetsLocationsGatewayTimeout() *PostCharactersCharacterIDAssetsLocationsGatewayTimeout {
	return &PostCharactersCharacterIDAssetsLocationsGatewayTimeout{}
}

/*PostCharactersCharacterIDAssetsLocationsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type PostCharactersCharacterIDAssetsLocationsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
