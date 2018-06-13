// Code generated by go-swagger; DO NOT EDIT.

package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDOrdersHistoryReader is a Reader for the GetCharactersCharacterIDOrdersHistory structure.
type GetCharactersCharacterIDOrdersHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDOrdersHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDOrdersHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDOrdersHistoryNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDOrdersHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDOrdersHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDOrdersHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDOrdersHistoryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDOrdersHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDOrdersHistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDOrdersHistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOrdersHistoryOK creates a GetCharactersCharacterIDOrdersHistoryOK with default headers values
func NewGetCharactersCharacterIDOrdersHistoryOK() *GetCharactersCharacterIDOrdersHistoryOK {
	return &GetCharactersCharacterIDOrdersHistoryOK{
		XPages: 1,
	}
}

/*GetCharactersCharacterIDOrdersHistoryOK handles this case with default header values.

Expired and cancelled market orders placed by a character
*/
type GetCharactersCharacterIDOrdersHistoryOK struct {
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
	/*Maximum page number
	 */
	XPages int32

	Payload []*models.GetCharactersCharacterIDOrdersHistoryOKBodyItems
}

func (o *GetCharactersCharacterIDOrdersHistoryOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header X-Pages
	xPages, err := swag.ConvertInt32(response.GetHeader("X-Pages"))
	if err != nil {
		return errors.InvalidType("X-Pages", "header", "int32", response.GetHeader("X-Pages"))
	}
	o.XPages = xPages

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryNotModified creates a GetCharactersCharacterIDOrdersHistoryNotModified with default headers values
func NewGetCharactersCharacterIDOrdersHistoryNotModified() *GetCharactersCharacterIDOrdersHistoryNotModified {
	return &GetCharactersCharacterIDOrdersHistoryNotModified{}
}

/*GetCharactersCharacterIDOrdersHistoryNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDOrdersHistoryNotModified struct {
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

func (o *GetCharactersCharacterIDOrdersHistoryNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryNotModified ", 304)
}

func (o *GetCharactersCharacterIDOrdersHistoryNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOrdersHistoryBadRequest creates a GetCharactersCharacterIDOrdersHistoryBadRequest with default headers values
func NewGetCharactersCharacterIDOrdersHistoryBadRequest() *GetCharactersCharacterIDOrdersHistoryBadRequest {
	return &GetCharactersCharacterIDOrdersHistoryBadRequest{}
}

/*GetCharactersCharacterIDOrdersHistoryBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDOrdersHistoryBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDOrdersHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryUnauthorized creates a GetCharactersCharacterIDOrdersHistoryUnauthorized with default headers values
func NewGetCharactersCharacterIDOrdersHistoryUnauthorized() *GetCharactersCharacterIDOrdersHistoryUnauthorized {
	return &GetCharactersCharacterIDOrdersHistoryUnauthorized{}
}

/*GetCharactersCharacterIDOrdersHistoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDOrdersHistoryUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDOrdersHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryForbidden creates a GetCharactersCharacterIDOrdersHistoryForbidden with default headers values
func NewGetCharactersCharacterIDOrdersHistoryForbidden() *GetCharactersCharacterIDOrdersHistoryForbidden {
	return &GetCharactersCharacterIDOrdersHistoryForbidden{}
}

/*GetCharactersCharacterIDOrdersHistoryForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDOrdersHistoryForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDOrdersHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryEnhanceYourCalm creates a GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDOrdersHistoryEnhanceYourCalm() *GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm {
	return &GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm{}
}

/*GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryInternalServerError creates a GetCharactersCharacterIDOrdersHistoryInternalServerError with default headers values
func NewGetCharactersCharacterIDOrdersHistoryInternalServerError() *GetCharactersCharacterIDOrdersHistoryInternalServerError {
	return &GetCharactersCharacterIDOrdersHistoryInternalServerError{}
}

/*GetCharactersCharacterIDOrdersHistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDOrdersHistoryInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDOrdersHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryServiceUnavailable creates a GetCharactersCharacterIDOrdersHistoryServiceUnavailable with default headers values
func NewGetCharactersCharacterIDOrdersHistoryServiceUnavailable() *GetCharactersCharacterIDOrdersHistoryServiceUnavailable {
	return &GetCharactersCharacterIDOrdersHistoryServiceUnavailable{}
}

/*GetCharactersCharacterIDOrdersHistoryServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDOrdersHistoryServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDOrdersHistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOrdersHistoryGatewayTimeout creates a GetCharactersCharacterIDOrdersHistoryGatewayTimeout with default headers values
func NewGetCharactersCharacterIDOrdersHistoryGatewayTimeout() *GetCharactersCharacterIDOrdersHistoryGatewayTimeout {
	return &GetCharactersCharacterIDOrdersHistoryGatewayTimeout{}
}

/*GetCharactersCharacterIDOrdersHistoryGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDOrdersHistoryGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDOrdersHistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/orders/history/][%d] getCharactersCharacterIdOrdersHistoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDOrdersHistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
