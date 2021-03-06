// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// GetCorporationsCorporationIDMedalsReader is a Reader for the GetCorporationsCorporationIDMedals structure.
type GetCorporationsCorporationIDMedalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMedalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDMedalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDMedalsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDMedalsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDMedalsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDMedalsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDMedalsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDMedalsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDMedalsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDMedalsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMedalsOK creates a GetCorporationsCorporationIDMedalsOK with default headers values
func NewGetCorporationsCorporationIDMedalsOK() *GetCorporationsCorporationIDMedalsOK {
	return &GetCorporationsCorporationIDMedalsOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDMedalsOK handles this case with default header values.

A list of medals
*/
type GetCorporationsCorporationIDMedalsOK struct {
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

	Payload []*models.GetCorporationsCorporationIDMedalsOKBodyItems
}

func (o *GetCorporationsCorporationIDMedalsOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsNotModified creates a GetCorporationsCorporationIDMedalsNotModified with default headers values
func NewGetCorporationsCorporationIDMedalsNotModified() *GetCorporationsCorporationIDMedalsNotModified {
	return &GetCorporationsCorporationIDMedalsNotModified{}
}

/*GetCorporationsCorporationIDMedalsNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDMedalsNotModified struct {
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

func (o *GetCorporationsCorporationIDMedalsNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMedalsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsBadRequest creates a GetCorporationsCorporationIDMedalsBadRequest with default headers values
func NewGetCorporationsCorporationIDMedalsBadRequest() *GetCorporationsCorporationIDMedalsBadRequest {
	return &GetCorporationsCorporationIDMedalsBadRequest{}
}

/*GetCorporationsCorporationIDMedalsBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDMedalsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDMedalsBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsUnauthorized creates a GetCorporationsCorporationIDMedalsUnauthorized with default headers values
func NewGetCorporationsCorporationIDMedalsUnauthorized() *GetCorporationsCorporationIDMedalsUnauthorized {
	return &GetCorporationsCorporationIDMedalsUnauthorized{}
}

/*GetCorporationsCorporationIDMedalsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMedalsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDMedalsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsForbidden creates a GetCorporationsCorporationIDMedalsForbidden with default headers values
func NewGetCorporationsCorporationIDMedalsForbidden() *GetCorporationsCorporationIDMedalsForbidden {
	return &GetCorporationsCorporationIDMedalsForbidden{}
}

/*GetCorporationsCorporationIDMedalsForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMedalsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDMedalsForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsEnhanceYourCalm creates a GetCorporationsCorporationIDMedalsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMedalsEnhanceYourCalm() *GetCorporationsCorporationIDMedalsEnhanceYourCalm {
	return &GetCorporationsCorporationIDMedalsEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDMedalsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDMedalsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDMedalsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsInternalServerError creates a GetCorporationsCorporationIDMedalsInternalServerError with default headers values
func NewGetCorporationsCorporationIDMedalsInternalServerError() *GetCorporationsCorporationIDMedalsInternalServerError {
	return &GetCorporationsCorporationIDMedalsInternalServerError{}
}

/*GetCorporationsCorporationIDMedalsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMedalsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDMedalsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsServiceUnavailable creates a GetCorporationsCorporationIDMedalsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMedalsServiceUnavailable() *GetCorporationsCorporationIDMedalsServiceUnavailable {
	return &GetCorporationsCorporationIDMedalsServiceUnavailable{}
}

/*GetCorporationsCorporationIDMedalsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMedalsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDMedalsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsGatewayTimeout creates a GetCorporationsCorporationIDMedalsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMedalsGatewayTimeout() *GetCorporationsCorporationIDMedalsGatewayTimeout {
	return &GetCorporationsCorporationIDMedalsGatewayTimeout{}
}

/*GetCorporationsCorporationIDMedalsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMedalsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDMedalsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
