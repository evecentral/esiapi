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

// PostCharactersCharacterIDAssetsNamesReader is a Reader for the PostCharactersCharacterIDAssetsNames structure.
type PostCharactersCharacterIDAssetsNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDAssetsNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostCharactersCharacterIDAssetsNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostCharactersCharacterIDAssetsNamesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostCharactersCharacterIDAssetsNamesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostCharactersCharacterIDAssetsNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewPostCharactersCharacterIDAssetsNamesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostCharactersCharacterIDAssetsNamesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostCharactersCharacterIDAssetsNamesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewPostCharactersCharacterIDAssetsNamesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCharactersCharacterIDAssetsNamesOK creates a PostCharactersCharacterIDAssetsNamesOK with default headers values
func NewPostCharactersCharacterIDAssetsNamesOK() *PostCharactersCharacterIDAssetsNamesOK {
	return &PostCharactersCharacterIDAssetsNamesOK{}
}

/*PostCharactersCharacterIDAssetsNamesOK handles this case with default header values.

List of asset names
*/
type PostCharactersCharacterIDAssetsNamesOK struct {
	Payload []*models.PostCharactersCharacterIDAssetsNamesOKBodyItems
}

func (o *PostCharactersCharacterIDAssetsNamesOK) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesOK  %+v", 200, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsNamesBadRequest creates a PostCharactersCharacterIDAssetsNamesBadRequest with default headers values
func NewPostCharactersCharacterIDAssetsNamesBadRequest() *PostCharactersCharacterIDAssetsNamesBadRequest {
	return &PostCharactersCharacterIDAssetsNamesBadRequest{}
}

/*PostCharactersCharacterIDAssetsNamesBadRequest handles this case with default header values.

Bad request
*/
type PostCharactersCharacterIDAssetsNamesBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostCharactersCharacterIDAssetsNamesBadRequest) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesBadRequest  %+v", 400, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsNamesUnauthorized creates a PostCharactersCharacterIDAssetsNamesUnauthorized with default headers values
func NewPostCharactersCharacterIDAssetsNamesUnauthorized() *PostCharactersCharacterIDAssetsNamesUnauthorized {
	return &PostCharactersCharacterIDAssetsNamesUnauthorized{}
}

/*PostCharactersCharacterIDAssetsNamesUnauthorized handles this case with default header values.

Unauthorized
*/
type PostCharactersCharacterIDAssetsNamesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostCharactersCharacterIDAssetsNamesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsNamesForbidden creates a PostCharactersCharacterIDAssetsNamesForbidden with default headers values
func NewPostCharactersCharacterIDAssetsNamesForbidden() *PostCharactersCharacterIDAssetsNamesForbidden {
	return &PostCharactersCharacterIDAssetsNamesForbidden{}
}

/*PostCharactersCharacterIDAssetsNamesForbidden handles this case with default header values.

Forbidden
*/
type PostCharactersCharacterIDAssetsNamesForbidden struct {
	Payload *models.Forbidden
}

func (o *PostCharactersCharacterIDAssetsNamesForbidden) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsNamesEnhanceYourCalm creates a PostCharactersCharacterIDAssetsNamesEnhanceYourCalm with default headers values
func NewPostCharactersCharacterIDAssetsNamesEnhanceYourCalm() *PostCharactersCharacterIDAssetsNamesEnhanceYourCalm {
	return &PostCharactersCharacterIDAssetsNamesEnhanceYourCalm{}
}

/*PostCharactersCharacterIDAssetsNamesEnhanceYourCalm handles this case with default header values.

Error limited
*/
type PostCharactersCharacterIDAssetsNamesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostCharactersCharacterIDAssetsNamesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsNamesInternalServerError creates a PostCharactersCharacterIDAssetsNamesInternalServerError with default headers values
func NewPostCharactersCharacterIDAssetsNamesInternalServerError() *PostCharactersCharacterIDAssetsNamesInternalServerError {
	return &PostCharactersCharacterIDAssetsNamesInternalServerError{}
}

/*PostCharactersCharacterIDAssetsNamesInternalServerError handles this case with default header values.

Internal server error
*/
type PostCharactersCharacterIDAssetsNamesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostCharactersCharacterIDAssetsNamesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsNamesServiceUnavailable creates a PostCharactersCharacterIDAssetsNamesServiceUnavailable with default headers values
func NewPostCharactersCharacterIDAssetsNamesServiceUnavailable() *PostCharactersCharacterIDAssetsNamesServiceUnavailable {
	return &PostCharactersCharacterIDAssetsNamesServiceUnavailable{}
}

/*PostCharactersCharacterIDAssetsNamesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type PostCharactersCharacterIDAssetsNamesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostCharactersCharacterIDAssetsNamesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsNamesGatewayTimeout creates a PostCharactersCharacterIDAssetsNamesGatewayTimeout with default headers values
func NewPostCharactersCharacterIDAssetsNamesGatewayTimeout() *PostCharactersCharacterIDAssetsNamesGatewayTimeout {
	return &PostCharactersCharacterIDAssetsNamesGatewayTimeout{}
}

/*PostCharactersCharacterIDAssetsNamesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type PostCharactersCharacterIDAssetsNamesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostCharactersCharacterIDAssetsNamesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /characters/{character_id}/assets/names/][%d] postCharactersCharacterIdAssetsNamesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsNamesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
