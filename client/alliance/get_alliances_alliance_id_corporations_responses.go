// Code generated by go-swagger; DO NOT EDIT.

package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetAlliancesAllianceIDCorporationsReader is a Reader for the GetAlliancesAllianceIDCorporations structure.
type GetAlliancesAllianceIDCorporationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesAllianceIDCorporationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlliancesAllianceIDCorporationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetAlliancesAllianceIDCorporationsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetAlliancesAllianceIDCorporationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetAlliancesAllianceIDCorporationsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAlliancesAllianceIDCorporationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAlliancesAllianceIDCorporationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetAlliancesAllianceIDCorporationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlliancesAllianceIDCorporationsOK creates a GetAlliancesAllianceIDCorporationsOK with default headers values
func NewGetAlliancesAllianceIDCorporationsOK() *GetAlliancesAllianceIDCorporationsOK {
	return &GetAlliancesAllianceIDCorporationsOK{}
}

/*GetAlliancesAllianceIDCorporationsOK handles this case with default header values.

List of corporation IDs
*/
type GetAlliancesAllianceIDCorporationsOK struct {
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

	Payload []*int32
}

func (o *GetAlliancesAllianceIDCorporationsOK) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/corporations/][%d] getAlliancesAllianceIdCorporationsOK  %+v", 200, o.Payload)
}

func (o *GetAlliancesAllianceIDCorporationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDCorporationsNotModified creates a GetAlliancesAllianceIDCorporationsNotModified with default headers values
func NewGetAlliancesAllianceIDCorporationsNotModified() *GetAlliancesAllianceIDCorporationsNotModified {
	return &GetAlliancesAllianceIDCorporationsNotModified{}
}

/*GetAlliancesAllianceIDCorporationsNotModified handles this case with default header values.

Not modified
*/
type GetAlliancesAllianceIDCorporationsNotModified struct {
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

func (o *GetAlliancesAllianceIDCorporationsNotModified) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/corporations/][%d] getAlliancesAllianceIdCorporationsNotModified ", 304)
}

func (o *GetAlliancesAllianceIDCorporationsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDCorporationsBadRequest creates a GetAlliancesAllianceIDCorporationsBadRequest with default headers values
func NewGetAlliancesAllianceIDCorporationsBadRequest() *GetAlliancesAllianceIDCorporationsBadRequest {
	return &GetAlliancesAllianceIDCorporationsBadRequest{}
}

/*GetAlliancesAllianceIDCorporationsBadRequest handles this case with default header values.

Bad request
*/
type GetAlliancesAllianceIDCorporationsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetAlliancesAllianceIDCorporationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/corporations/][%d] getAlliancesAllianceIdCorporationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlliancesAllianceIDCorporationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDCorporationsEnhanceYourCalm creates a GetAlliancesAllianceIDCorporationsEnhanceYourCalm with default headers values
func NewGetAlliancesAllianceIDCorporationsEnhanceYourCalm() *GetAlliancesAllianceIDCorporationsEnhanceYourCalm {
	return &GetAlliancesAllianceIDCorporationsEnhanceYourCalm{}
}

/*GetAlliancesAllianceIDCorporationsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetAlliancesAllianceIDCorporationsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetAlliancesAllianceIDCorporationsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/corporations/][%d] getAlliancesAllianceIdCorporationsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetAlliancesAllianceIDCorporationsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDCorporationsInternalServerError creates a GetAlliancesAllianceIDCorporationsInternalServerError with default headers values
func NewGetAlliancesAllianceIDCorporationsInternalServerError() *GetAlliancesAllianceIDCorporationsInternalServerError {
	return &GetAlliancesAllianceIDCorporationsInternalServerError{}
}

/*GetAlliancesAllianceIDCorporationsInternalServerError handles this case with default header values.

Internal server error
*/
type GetAlliancesAllianceIDCorporationsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAlliancesAllianceIDCorporationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/corporations/][%d] getAlliancesAllianceIdCorporationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesAllianceIDCorporationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDCorporationsServiceUnavailable creates a GetAlliancesAllianceIDCorporationsServiceUnavailable with default headers values
func NewGetAlliancesAllianceIDCorporationsServiceUnavailable() *GetAlliancesAllianceIDCorporationsServiceUnavailable {
	return &GetAlliancesAllianceIDCorporationsServiceUnavailable{}
}

/*GetAlliancesAllianceIDCorporationsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetAlliancesAllianceIDCorporationsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetAlliancesAllianceIDCorporationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/corporations/][%d] getAlliancesAllianceIdCorporationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlliancesAllianceIDCorporationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDCorporationsGatewayTimeout creates a GetAlliancesAllianceIDCorporationsGatewayTimeout with default headers values
func NewGetAlliancesAllianceIDCorporationsGatewayTimeout() *GetAlliancesAllianceIDCorporationsGatewayTimeout {
	return &GetAlliancesAllianceIDCorporationsGatewayTimeout{}
}

/*GetAlliancesAllianceIDCorporationsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetAlliancesAllianceIDCorporationsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetAlliancesAllianceIDCorporationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/corporations/][%d] getAlliancesAllianceIdCorporationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlliancesAllianceIDCorporationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
