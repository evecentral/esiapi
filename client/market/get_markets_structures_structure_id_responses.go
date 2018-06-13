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

// GetMarketsStructuresStructureIDReader is a Reader for the GetMarketsStructuresStructureID structure.
type GetMarketsStructuresStructureIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsStructuresStructureIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMarketsStructuresStructureIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetMarketsStructuresStructureIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetMarketsStructuresStructureIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetMarketsStructuresStructureIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetMarketsStructuresStructureIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetMarketsStructuresStructureIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMarketsStructuresStructureIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetMarketsStructuresStructureIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetMarketsStructuresStructureIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMarketsStructuresStructureIDOK creates a GetMarketsStructuresStructureIDOK with default headers values
func NewGetMarketsStructuresStructureIDOK() *GetMarketsStructuresStructureIDOK {
	return &GetMarketsStructuresStructureIDOK{
		XPages: 1,
	}
}

/*GetMarketsStructuresStructureIDOK handles this case with default header values.

A list of orders
*/
type GetMarketsStructuresStructureIDOK struct {
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

	Payload []*models.GetMarketsStructuresStructureIDOKBodyItems
}

func (o *GetMarketsStructuresStructureIDOK) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdOK  %+v", 200, o.Payload)
}

func (o *GetMarketsStructuresStructureIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsStructuresStructureIDNotModified creates a GetMarketsStructuresStructureIDNotModified with default headers values
func NewGetMarketsStructuresStructureIDNotModified() *GetMarketsStructuresStructureIDNotModified {
	return &GetMarketsStructuresStructureIDNotModified{}
}

/*GetMarketsStructuresStructureIDNotModified handles this case with default header values.

Not modified
*/
type GetMarketsStructuresStructureIDNotModified struct {
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

func (o *GetMarketsStructuresStructureIDNotModified) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdNotModified ", 304)
}

func (o *GetMarketsStructuresStructureIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsStructuresStructureIDBadRequest creates a GetMarketsStructuresStructureIDBadRequest with default headers values
func NewGetMarketsStructuresStructureIDBadRequest() *GetMarketsStructuresStructureIDBadRequest {
	return &GetMarketsStructuresStructureIDBadRequest{}
}

/*GetMarketsStructuresStructureIDBadRequest handles this case with default header values.

Bad request
*/
type GetMarketsStructuresStructureIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetMarketsStructuresStructureIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetMarketsStructuresStructureIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDUnauthorized creates a GetMarketsStructuresStructureIDUnauthorized with default headers values
func NewGetMarketsStructuresStructureIDUnauthorized() *GetMarketsStructuresStructureIDUnauthorized {
	return &GetMarketsStructuresStructureIDUnauthorized{}
}

/*GetMarketsStructuresStructureIDUnauthorized handles this case with default header values.

Unauthorized
*/
type GetMarketsStructuresStructureIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetMarketsStructuresStructureIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMarketsStructuresStructureIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDForbidden creates a GetMarketsStructuresStructureIDForbidden with default headers values
func NewGetMarketsStructuresStructureIDForbidden() *GetMarketsStructuresStructureIDForbidden {
	return &GetMarketsStructuresStructureIDForbidden{}
}

/*GetMarketsStructuresStructureIDForbidden handles this case with default header values.

Forbidden
*/
type GetMarketsStructuresStructureIDForbidden struct {
	Payload *models.Forbidden
}

func (o *GetMarketsStructuresStructureIDForbidden) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdForbidden  %+v", 403, o.Payload)
}

func (o *GetMarketsStructuresStructureIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDEnhanceYourCalm creates a GetMarketsStructuresStructureIDEnhanceYourCalm with default headers values
func NewGetMarketsStructuresStructureIDEnhanceYourCalm() *GetMarketsStructuresStructureIDEnhanceYourCalm {
	return &GetMarketsStructuresStructureIDEnhanceYourCalm{}
}

/*GetMarketsStructuresStructureIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetMarketsStructuresStructureIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetMarketsStructuresStructureIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetMarketsStructuresStructureIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDInternalServerError creates a GetMarketsStructuresStructureIDInternalServerError with default headers values
func NewGetMarketsStructuresStructureIDInternalServerError() *GetMarketsStructuresStructureIDInternalServerError {
	return &GetMarketsStructuresStructureIDInternalServerError{}
}

/*GetMarketsStructuresStructureIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetMarketsStructuresStructureIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetMarketsStructuresStructureIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMarketsStructuresStructureIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDServiceUnavailable creates a GetMarketsStructuresStructureIDServiceUnavailable with default headers values
func NewGetMarketsStructuresStructureIDServiceUnavailable() *GetMarketsStructuresStructureIDServiceUnavailable {
	return &GetMarketsStructuresStructureIDServiceUnavailable{}
}

/*GetMarketsStructuresStructureIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetMarketsStructuresStructureIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetMarketsStructuresStructureIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetMarketsStructuresStructureIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsStructuresStructureIDGatewayTimeout creates a GetMarketsStructuresStructureIDGatewayTimeout with default headers values
func NewGetMarketsStructuresStructureIDGatewayTimeout() *GetMarketsStructuresStructureIDGatewayTimeout {
	return &GetMarketsStructuresStructureIDGatewayTimeout{}
}

/*GetMarketsStructuresStructureIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetMarketsStructuresStructureIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetMarketsStructuresStructureIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /markets/structures/{structure_id}/][%d] getMarketsStructuresStructureIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetMarketsStructuresStructureIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
