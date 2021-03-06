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

// GetCorporationsCorporationIDOrdersHistoryReader is a Reader for the GetCorporationsCorporationIDOrdersHistory structure.
type GetCorporationsCorporationIDOrdersHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDOrdersHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDOrdersHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDOrdersHistoryNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDOrdersHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDOrdersHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDOrdersHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDOrdersHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDOrdersHistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDOrdersHistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDOrdersHistoryOK creates a GetCorporationsCorporationIDOrdersHistoryOK with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryOK() *GetCorporationsCorporationIDOrdersHistoryOK {
	return &GetCorporationsCorporationIDOrdersHistoryOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDOrdersHistoryOK handles this case with default header values.

Expired and cancelled market orders placed on behalf of a corporation
*/
type GetCorporationsCorporationIDOrdersHistoryOK struct {
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

	Payload []*models.GetCorporationsCorporationIDOrdersHistoryOKBodyItems
}

func (o *GetCorporationsCorporationIDOrdersHistoryOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDOrdersHistoryNotModified creates a GetCorporationsCorporationIDOrdersHistoryNotModified with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryNotModified() *GetCorporationsCorporationIDOrdersHistoryNotModified {
	return &GetCorporationsCorporationIDOrdersHistoryNotModified{}
}

/*GetCorporationsCorporationIDOrdersHistoryNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDOrdersHistoryNotModified struct {
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

func (o *GetCorporationsCorporationIDOrdersHistoryNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryNotModified ", 304)
}

func (o *GetCorporationsCorporationIDOrdersHistoryNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDOrdersHistoryBadRequest creates a GetCorporationsCorporationIDOrdersHistoryBadRequest with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryBadRequest() *GetCorporationsCorporationIDOrdersHistoryBadRequest {
	return &GetCorporationsCorporationIDOrdersHistoryBadRequest{}
}

/*GetCorporationsCorporationIDOrdersHistoryBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDOrdersHistoryBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDOrdersHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryUnauthorized creates a GetCorporationsCorporationIDOrdersHistoryUnauthorized with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryUnauthorized() *GetCorporationsCorporationIDOrdersHistoryUnauthorized {
	return &GetCorporationsCorporationIDOrdersHistoryUnauthorized{}
}

/*GetCorporationsCorporationIDOrdersHistoryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDOrdersHistoryUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDOrdersHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryForbidden creates a GetCorporationsCorporationIDOrdersHistoryForbidden with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryForbidden() *GetCorporationsCorporationIDOrdersHistoryForbidden {
	return &GetCorporationsCorporationIDOrdersHistoryForbidden{}
}

/*GetCorporationsCorporationIDOrdersHistoryForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDOrdersHistoryForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDOrdersHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm creates a GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm() *GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm {
	return &GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryInternalServerError creates a GetCorporationsCorporationIDOrdersHistoryInternalServerError with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryInternalServerError() *GetCorporationsCorporationIDOrdersHistoryInternalServerError {
	return &GetCorporationsCorporationIDOrdersHistoryInternalServerError{}
}

/*GetCorporationsCorporationIDOrdersHistoryInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDOrdersHistoryInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDOrdersHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryServiceUnavailable creates a GetCorporationsCorporationIDOrdersHistoryServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryServiceUnavailable() *GetCorporationsCorporationIDOrdersHistoryServiceUnavailable {
	return &GetCorporationsCorporationIDOrdersHistoryServiceUnavailable{}
}

/*GetCorporationsCorporationIDOrdersHistoryServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDOrdersHistoryServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDOrdersHistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersHistoryGatewayTimeout creates a GetCorporationsCorporationIDOrdersHistoryGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDOrdersHistoryGatewayTimeout() *GetCorporationsCorporationIDOrdersHistoryGatewayTimeout {
	return &GetCorporationsCorporationIDOrdersHistoryGatewayTimeout{}
}

/*GetCorporationsCorporationIDOrdersHistoryGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDOrdersHistoryGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDOrdersHistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/history/][%d] getCorporationsCorporationIdOrdersHistoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersHistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
