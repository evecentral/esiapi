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

// GetDogmaEffectsReader is a Reader for the GetDogmaEffects structure.
type GetDogmaEffectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDogmaEffectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDogmaEffectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetDogmaEffectsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetDogmaEffectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetDogmaEffectsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDogmaEffectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetDogmaEffectsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetDogmaEffectsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDogmaEffectsOK creates a GetDogmaEffectsOK with default headers values
func NewGetDogmaEffectsOK() *GetDogmaEffectsOK {
	return &GetDogmaEffectsOK{}
}

/*GetDogmaEffectsOK handles this case with default header values.

A list of dogma effect ids
*/
type GetDogmaEffectsOK struct {
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

	Payload []int32
}

func (o *GetDogmaEffectsOK) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/][%d] getDogmaEffectsOK  %+v", 200, o.Payload)
}

func (o *GetDogmaEffectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDogmaEffectsNotModified creates a GetDogmaEffectsNotModified with default headers values
func NewGetDogmaEffectsNotModified() *GetDogmaEffectsNotModified {
	return &GetDogmaEffectsNotModified{}
}

/*GetDogmaEffectsNotModified handles this case with default header values.

Not modified
*/
type GetDogmaEffectsNotModified struct {
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

func (o *GetDogmaEffectsNotModified) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/][%d] getDogmaEffectsNotModified ", 304)
}

func (o *GetDogmaEffectsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDogmaEffectsBadRequest creates a GetDogmaEffectsBadRequest with default headers values
func NewGetDogmaEffectsBadRequest() *GetDogmaEffectsBadRequest {
	return &GetDogmaEffectsBadRequest{}
}

/*GetDogmaEffectsBadRequest handles this case with default header values.

Bad request
*/
type GetDogmaEffectsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetDogmaEffectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/][%d] getDogmaEffectsBadRequest  %+v", 400, o.Payload)
}

func (o *GetDogmaEffectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEnhanceYourCalm creates a GetDogmaEffectsEnhanceYourCalm with default headers values
func NewGetDogmaEffectsEnhanceYourCalm() *GetDogmaEffectsEnhanceYourCalm {
	return &GetDogmaEffectsEnhanceYourCalm{}
}

/*GetDogmaEffectsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetDogmaEffectsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetDogmaEffectsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/][%d] getDogmaEffectsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetDogmaEffectsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsInternalServerError creates a GetDogmaEffectsInternalServerError with default headers values
func NewGetDogmaEffectsInternalServerError() *GetDogmaEffectsInternalServerError {
	return &GetDogmaEffectsInternalServerError{}
}

/*GetDogmaEffectsInternalServerError handles this case with default header values.

Internal server error
*/
type GetDogmaEffectsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetDogmaEffectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/][%d] getDogmaEffectsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDogmaEffectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsServiceUnavailable creates a GetDogmaEffectsServiceUnavailable with default headers values
func NewGetDogmaEffectsServiceUnavailable() *GetDogmaEffectsServiceUnavailable {
	return &GetDogmaEffectsServiceUnavailable{}
}

/*GetDogmaEffectsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetDogmaEffectsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetDogmaEffectsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/][%d] getDogmaEffectsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDogmaEffectsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsGatewayTimeout creates a GetDogmaEffectsGatewayTimeout with default headers values
func NewGetDogmaEffectsGatewayTimeout() *GetDogmaEffectsGatewayTimeout {
	return &GetDogmaEffectsGatewayTimeout{}
}

/*GetDogmaEffectsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetDogmaEffectsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetDogmaEffectsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/][%d] getDogmaEffectsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDogmaEffectsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
