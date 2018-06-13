// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDOnlineReader is a Reader for the GetCharactersCharacterIDOnline structure.
type GetCharactersCharacterIDOnlineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDOnlineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDOnlineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDOnlineNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDOnlineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDOnlineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDOnlineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDOnlineEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDOnlineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDOnlineServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDOnlineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOnlineOK creates a GetCharactersCharacterIDOnlineOK with default headers values
func NewGetCharactersCharacterIDOnlineOK() *GetCharactersCharacterIDOnlineOK {
	return &GetCharactersCharacterIDOnlineOK{}
}

/*GetCharactersCharacterIDOnlineOK handles this case with default header values.

Object describing the character's online status
*/
type GetCharactersCharacterIDOnlineOK struct {
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

	Payload *models.GetCharactersCharacterIDOnlineOKBody
}

func (o *GetCharactersCharacterIDOnlineOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDOnlineOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineNotModified creates a GetCharactersCharacterIDOnlineNotModified with default headers values
func NewGetCharactersCharacterIDOnlineNotModified() *GetCharactersCharacterIDOnlineNotModified {
	return &GetCharactersCharacterIDOnlineNotModified{}
}

/*GetCharactersCharacterIDOnlineNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDOnlineNotModified struct {
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

func (o *GetCharactersCharacterIDOnlineNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineNotModified ", 304)
}

func (o *GetCharactersCharacterIDOnlineNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOnlineBadRequest creates a GetCharactersCharacterIDOnlineBadRequest with default headers values
func NewGetCharactersCharacterIDOnlineBadRequest() *GetCharactersCharacterIDOnlineBadRequest {
	return &GetCharactersCharacterIDOnlineBadRequest{}
}

/*GetCharactersCharacterIDOnlineBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDOnlineBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDOnlineBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineUnauthorized creates a GetCharactersCharacterIDOnlineUnauthorized with default headers values
func NewGetCharactersCharacterIDOnlineUnauthorized() *GetCharactersCharacterIDOnlineUnauthorized {
	return &GetCharactersCharacterIDOnlineUnauthorized{}
}

/*GetCharactersCharacterIDOnlineUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDOnlineUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDOnlineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineForbidden creates a GetCharactersCharacterIDOnlineForbidden with default headers values
func NewGetCharactersCharacterIDOnlineForbidden() *GetCharactersCharacterIDOnlineForbidden {
	return &GetCharactersCharacterIDOnlineForbidden{}
}

/*GetCharactersCharacterIDOnlineForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDOnlineForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDOnlineForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineEnhanceYourCalm creates a GetCharactersCharacterIDOnlineEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDOnlineEnhanceYourCalm() *GetCharactersCharacterIDOnlineEnhanceYourCalm {
	return &GetCharactersCharacterIDOnlineEnhanceYourCalm{}
}

/*GetCharactersCharacterIDOnlineEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDOnlineEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDOnlineEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineInternalServerError creates a GetCharactersCharacterIDOnlineInternalServerError with default headers values
func NewGetCharactersCharacterIDOnlineInternalServerError() *GetCharactersCharacterIDOnlineInternalServerError {
	return &GetCharactersCharacterIDOnlineInternalServerError{}
}

/*GetCharactersCharacterIDOnlineInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDOnlineInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDOnlineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineServiceUnavailable creates a GetCharactersCharacterIDOnlineServiceUnavailable with default headers values
func NewGetCharactersCharacterIDOnlineServiceUnavailable() *GetCharactersCharacterIDOnlineServiceUnavailable {
	return &GetCharactersCharacterIDOnlineServiceUnavailable{}
}

/*GetCharactersCharacterIDOnlineServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDOnlineServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDOnlineServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOnlineGatewayTimeout creates a GetCharactersCharacterIDOnlineGatewayTimeout with default headers values
func NewGetCharactersCharacterIDOnlineGatewayTimeout() *GetCharactersCharacterIDOnlineGatewayTimeout {
	return &GetCharactersCharacterIDOnlineGatewayTimeout{}
}

/*GetCharactersCharacterIDOnlineGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDOnlineGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDOnlineGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/online/][%d] getCharactersCharacterIdOnlineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDOnlineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
