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

// GetCharactersCharacterIDFatigueReader is a Reader for the GetCharactersCharacterIDFatigue structure.
type GetCharactersCharacterIDFatigueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDFatigueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDFatigueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDFatigueNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDFatigueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDFatigueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDFatigueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDFatigueEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDFatigueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDFatigueServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDFatigueGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDFatigueOK creates a GetCharactersCharacterIDFatigueOK with default headers values
func NewGetCharactersCharacterIDFatigueOK() *GetCharactersCharacterIDFatigueOK {
	return &GetCharactersCharacterIDFatigueOK{}
}

/*GetCharactersCharacterIDFatigueOK handles this case with default header values.

Jump activation and fatigue information
*/
type GetCharactersCharacterIDFatigueOK struct {
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

	Payload *models.GetCharactersCharacterIDFatigueOKBody
}

func (o *GetCharactersCharacterIDFatigueOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDFatigueOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueNotModified creates a GetCharactersCharacterIDFatigueNotModified with default headers values
func NewGetCharactersCharacterIDFatigueNotModified() *GetCharactersCharacterIDFatigueNotModified {
	return &GetCharactersCharacterIDFatigueNotModified{}
}

/*GetCharactersCharacterIDFatigueNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDFatigueNotModified struct {
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

func (o *GetCharactersCharacterIDFatigueNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueNotModified ", 304)
}

func (o *GetCharactersCharacterIDFatigueNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDFatigueBadRequest creates a GetCharactersCharacterIDFatigueBadRequest with default headers values
func NewGetCharactersCharacterIDFatigueBadRequest() *GetCharactersCharacterIDFatigueBadRequest {
	return &GetCharactersCharacterIDFatigueBadRequest{}
}

/*GetCharactersCharacterIDFatigueBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDFatigueBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDFatigueBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueUnauthorized creates a GetCharactersCharacterIDFatigueUnauthorized with default headers values
func NewGetCharactersCharacterIDFatigueUnauthorized() *GetCharactersCharacterIDFatigueUnauthorized {
	return &GetCharactersCharacterIDFatigueUnauthorized{}
}

/*GetCharactersCharacterIDFatigueUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDFatigueUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDFatigueUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueForbidden creates a GetCharactersCharacterIDFatigueForbidden with default headers values
func NewGetCharactersCharacterIDFatigueForbidden() *GetCharactersCharacterIDFatigueForbidden {
	return &GetCharactersCharacterIDFatigueForbidden{}
}

/*GetCharactersCharacterIDFatigueForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDFatigueForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDFatigueForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueEnhanceYourCalm creates a GetCharactersCharacterIDFatigueEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDFatigueEnhanceYourCalm() *GetCharactersCharacterIDFatigueEnhanceYourCalm {
	return &GetCharactersCharacterIDFatigueEnhanceYourCalm{}
}

/*GetCharactersCharacterIDFatigueEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDFatigueEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDFatigueEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueInternalServerError creates a GetCharactersCharacterIDFatigueInternalServerError with default headers values
func NewGetCharactersCharacterIDFatigueInternalServerError() *GetCharactersCharacterIDFatigueInternalServerError {
	return &GetCharactersCharacterIDFatigueInternalServerError{}
}

/*GetCharactersCharacterIDFatigueInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDFatigueInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDFatigueInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueServiceUnavailable creates a GetCharactersCharacterIDFatigueServiceUnavailable with default headers values
func NewGetCharactersCharacterIDFatigueServiceUnavailable() *GetCharactersCharacterIDFatigueServiceUnavailable {
	return &GetCharactersCharacterIDFatigueServiceUnavailable{}
}

/*GetCharactersCharacterIDFatigueServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDFatigueServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDFatigueServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueGatewayTimeout creates a GetCharactersCharacterIDFatigueGatewayTimeout with default headers values
func NewGetCharactersCharacterIDFatigueGatewayTimeout() *GetCharactersCharacterIDFatigueGatewayTimeout {
	return &GetCharactersCharacterIDFatigueGatewayTimeout{}
}

/*GetCharactersCharacterIDFatigueGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDFatigueGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDFatigueGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDFatigueGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
