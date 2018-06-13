// Code generated by go-swagger; DO NOT EDIT.

package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetDogmaAttributesAttributeIDReader is a Reader for the GetDogmaAttributesAttributeID structure.
type GetDogmaAttributesAttributeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDogmaAttributesAttributeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDogmaAttributesAttributeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetDogmaAttributesAttributeIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetDogmaAttributesAttributeIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDogmaAttributesAttributeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetDogmaAttributesAttributeIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDogmaAttributesAttributeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetDogmaAttributesAttributeIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetDogmaAttributesAttributeIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDogmaAttributesAttributeIDOK creates a GetDogmaAttributesAttributeIDOK with default headers values
func NewGetDogmaAttributesAttributeIDOK() *GetDogmaAttributesAttributeIDOK {
	return &GetDogmaAttributesAttributeIDOK{}
}

/*GetDogmaAttributesAttributeIDOK handles this case with default header values.

Information about a dogma attribute
*/
type GetDogmaAttributesAttributeIDOK struct {
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

	Payload *models.GetDogmaAttributesAttributeIDOKBody
}

func (o *GetDogmaAttributesAttributeIDOK) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdOK  %+v", 200, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetDogmaAttributesAttributeIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaAttributesAttributeIDNotModified creates a GetDogmaAttributesAttributeIDNotModified with default headers values
func NewGetDogmaAttributesAttributeIDNotModified() *GetDogmaAttributesAttributeIDNotModified {
	return &GetDogmaAttributesAttributeIDNotModified{}
}

/*GetDogmaAttributesAttributeIDNotModified handles this case with default header values.

Not modified
*/
type GetDogmaAttributesAttributeIDNotModified struct {
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

func (o *GetDogmaAttributesAttributeIDNotModified) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdNotModified ", 304)
}

func (o *GetDogmaAttributesAttributeIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDogmaAttributesAttributeIDBadRequest creates a GetDogmaAttributesAttributeIDBadRequest with default headers values
func NewGetDogmaAttributesAttributeIDBadRequest() *GetDogmaAttributesAttributeIDBadRequest {
	return &GetDogmaAttributesAttributeIDBadRequest{}
}

/*GetDogmaAttributesAttributeIDBadRequest handles this case with default header values.

Bad request
*/
type GetDogmaAttributesAttributeIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetDogmaAttributesAttributeIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaAttributesAttributeIDNotFound creates a GetDogmaAttributesAttributeIDNotFound with default headers values
func NewGetDogmaAttributesAttributeIDNotFound() *GetDogmaAttributesAttributeIDNotFound {
	return &GetDogmaAttributesAttributeIDNotFound{}
}

/*GetDogmaAttributesAttributeIDNotFound handles this case with default header values.

Dogma attribute not found
*/
type GetDogmaAttributesAttributeIDNotFound struct {
	Payload *models.GetDogmaAttributesAttributeIDNotFoundBody
}

func (o *GetDogmaAttributesAttributeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDogmaAttributesAttributeIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaAttributesAttributeIDEnhanceYourCalm creates a GetDogmaAttributesAttributeIDEnhanceYourCalm with default headers values
func NewGetDogmaAttributesAttributeIDEnhanceYourCalm() *GetDogmaAttributesAttributeIDEnhanceYourCalm {
	return &GetDogmaAttributesAttributeIDEnhanceYourCalm{}
}

/*GetDogmaAttributesAttributeIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetDogmaAttributesAttributeIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetDogmaAttributesAttributeIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaAttributesAttributeIDInternalServerError creates a GetDogmaAttributesAttributeIDInternalServerError with default headers values
func NewGetDogmaAttributesAttributeIDInternalServerError() *GetDogmaAttributesAttributeIDInternalServerError {
	return &GetDogmaAttributesAttributeIDInternalServerError{}
}

/*GetDogmaAttributesAttributeIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetDogmaAttributesAttributeIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetDogmaAttributesAttributeIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaAttributesAttributeIDServiceUnavailable creates a GetDogmaAttributesAttributeIDServiceUnavailable with default headers values
func NewGetDogmaAttributesAttributeIDServiceUnavailable() *GetDogmaAttributesAttributeIDServiceUnavailable {
	return &GetDogmaAttributesAttributeIDServiceUnavailable{}
}

/*GetDogmaAttributesAttributeIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetDogmaAttributesAttributeIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetDogmaAttributesAttributeIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaAttributesAttributeIDGatewayTimeout creates a GetDogmaAttributesAttributeIDGatewayTimeout with default headers values
func NewGetDogmaAttributesAttributeIDGatewayTimeout() *GetDogmaAttributesAttributeIDGatewayTimeout {
	return &GetDogmaAttributesAttributeIDGatewayTimeout{}
}

/*GetDogmaAttributesAttributeIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetDogmaAttributesAttributeIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetDogmaAttributesAttributeIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dogma/attributes/{attribute_id}/][%d] getDogmaAttributesAttributeIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDogmaAttributesAttributeIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
