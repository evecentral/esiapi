// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDPortraitReader is a Reader for the GetCharactersCharacterIDPortrait structure.
type GetCharactersCharacterIDPortraitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDPortraitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDPortraitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDPortraitNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDPortraitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCharactersCharacterIDPortraitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDPortraitEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDPortraitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDPortraitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDPortraitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDPortraitOK creates a GetCharactersCharacterIDPortraitOK with default headers values
func NewGetCharactersCharacterIDPortraitOK() *GetCharactersCharacterIDPortraitOK {
	return &GetCharactersCharacterIDPortraitOK{}
}

/*GetCharactersCharacterIDPortraitOK handles this case with default header values.

Public data for the given character
*/
type GetCharactersCharacterIDPortraitOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7232 compliant entity tag
	 */
	ETag string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCharactersCharacterIDPortraitOKBody
}

func (o *GetCharactersCharacterIDPortraitOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDPortraitOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitNotModified creates a GetCharactersCharacterIDPortraitNotModified with default headers values
func NewGetCharactersCharacterIDPortraitNotModified() *GetCharactersCharacterIDPortraitNotModified {
	return &GetCharactersCharacterIDPortraitNotModified{}
}

/*GetCharactersCharacterIDPortraitNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDPortraitNotModified struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7232 compliant entity tag
	 */
	ETag string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string
}

func (o *GetCharactersCharacterIDPortraitNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitNotModified ", 304)
}

func (o *GetCharactersCharacterIDPortraitNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	return nil
}

// NewGetCharactersCharacterIDPortraitBadRequest creates a GetCharactersCharacterIDPortraitBadRequest with default headers values
func NewGetCharactersCharacterIDPortraitBadRequest() *GetCharactersCharacterIDPortraitBadRequest {
	return &GetCharactersCharacterIDPortraitBadRequest{}
}

/*GetCharactersCharacterIDPortraitBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDPortraitBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDPortraitBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitNotFound creates a GetCharactersCharacterIDPortraitNotFound with default headers values
func NewGetCharactersCharacterIDPortraitNotFound() *GetCharactersCharacterIDPortraitNotFound {
	return &GetCharactersCharacterIDPortraitNotFound{}
}

/*GetCharactersCharacterIDPortraitNotFound handles this case with default header values.

No image server for this datasource
*/
type GetCharactersCharacterIDPortraitNotFound struct {
	Payload *models.GetCharactersCharacterIDPortraitNotFoundBody
}

func (o *GetCharactersCharacterIDPortraitNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDPortraitNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitEnhanceYourCalm creates a GetCharactersCharacterIDPortraitEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDPortraitEnhanceYourCalm() *GetCharactersCharacterIDPortraitEnhanceYourCalm {
	return &GetCharactersCharacterIDPortraitEnhanceYourCalm{}
}

/*GetCharactersCharacterIDPortraitEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDPortraitEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDPortraitEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitInternalServerError creates a GetCharactersCharacterIDPortraitInternalServerError with default headers values
func NewGetCharactersCharacterIDPortraitInternalServerError() *GetCharactersCharacterIDPortraitInternalServerError {
	return &GetCharactersCharacterIDPortraitInternalServerError{}
}

/*GetCharactersCharacterIDPortraitInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDPortraitInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDPortraitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitServiceUnavailable creates a GetCharactersCharacterIDPortraitServiceUnavailable with default headers values
func NewGetCharactersCharacterIDPortraitServiceUnavailable() *GetCharactersCharacterIDPortraitServiceUnavailable {
	return &GetCharactersCharacterIDPortraitServiceUnavailable{}
}

/*GetCharactersCharacterIDPortraitServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDPortraitServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDPortraitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitGatewayTimeout creates a GetCharactersCharacterIDPortraitGatewayTimeout with default headers values
func NewGetCharactersCharacterIDPortraitGatewayTimeout() *GetCharactersCharacterIDPortraitGatewayTimeout {
	return &GetCharactersCharacterIDPortraitGatewayTimeout{}
}

/*GetCharactersCharacterIDPortraitGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDPortraitGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDPortraitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDPortraitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
