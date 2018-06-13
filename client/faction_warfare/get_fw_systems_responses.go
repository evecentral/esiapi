// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetFwSystemsReader is a Reader for the GetFwSystems structure.
type GetFwSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFwSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFwSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetFwSystemsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetFwSystemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetFwSystemsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetFwSystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetFwSystemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetFwSystemsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFwSystemsOK creates a GetFwSystemsOK with default headers values
func NewGetFwSystemsOK() *GetFwSystemsOK {
	return &GetFwSystemsOK{}
}

/*GetFwSystemsOK handles this case with default header values.

All faction warfare solar systems
*/
type GetFwSystemsOK struct {
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

	Payload []*models.GetFwSystemsOKBodyItems
}

func (o *GetFwSystemsOK) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsOK  %+v", 200, o.Payload)
}

func (o *GetFwSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsNotModified creates a GetFwSystemsNotModified with default headers values
func NewGetFwSystemsNotModified() *GetFwSystemsNotModified {
	return &GetFwSystemsNotModified{}
}

/*GetFwSystemsNotModified handles this case with default header values.

Not modified
*/
type GetFwSystemsNotModified struct {
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

func (o *GetFwSystemsNotModified) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsNotModified ", 304)
}

func (o *GetFwSystemsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFwSystemsBadRequest creates a GetFwSystemsBadRequest with default headers values
func NewGetFwSystemsBadRequest() *GetFwSystemsBadRequest {
	return &GetFwSystemsBadRequest{}
}

/*GetFwSystemsBadRequest handles this case with default header values.

Bad request
*/
type GetFwSystemsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetFwSystemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetFwSystemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsEnhanceYourCalm creates a GetFwSystemsEnhanceYourCalm with default headers values
func NewGetFwSystemsEnhanceYourCalm() *GetFwSystemsEnhanceYourCalm {
	return &GetFwSystemsEnhanceYourCalm{}
}

/*GetFwSystemsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetFwSystemsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetFwSystemsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetFwSystemsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsInternalServerError creates a GetFwSystemsInternalServerError with default headers values
func NewGetFwSystemsInternalServerError() *GetFwSystemsInternalServerError {
	return &GetFwSystemsInternalServerError{}
}

/*GetFwSystemsInternalServerError handles this case with default header values.

Internal server error
*/
type GetFwSystemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetFwSystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFwSystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsServiceUnavailable creates a GetFwSystemsServiceUnavailable with default headers values
func NewGetFwSystemsServiceUnavailable() *GetFwSystemsServiceUnavailable {
	return &GetFwSystemsServiceUnavailable{}
}

/*GetFwSystemsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetFwSystemsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetFwSystemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFwSystemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFwSystemsGatewayTimeout creates a GetFwSystemsGatewayTimeout with default headers values
func NewGetFwSystemsGatewayTimeout() *GetFwSystemsGatewayTimeout {
	return &GetFwSystemsGatewayTimeout{}
}

/*GetFwSystemsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetFwSystemsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetFwSystemsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFwSystemsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
