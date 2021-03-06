// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetUniverseSystemJumpsReader is a Reader for the GetUniverseSystemJumps structure.
type GetUniverseSystemJumpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseSystemJumpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseSystemJumpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetUniverseSystemJumpsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetUniverseSystemJumpsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUniverseSystemJumpsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseSystemJumpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUniverseSystemJumpsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetUniverseSystemJumpsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseSystemJumpsOK creates a GetUniverseSystemJumpsOK with default headers values
func NewGetUniverseSystemJumpsOK() *GetUniverseSystemJumpsOK {
	return &GetUniverseSystemJumpsOK{}
}

/*GetUniverseSystemJumpsOK handles this case with default header values.

A list of systems and number of jumps
*/
type GetUniverseSystemJumpsOK struct {
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

	Payload []*models.GetUniverseSystemJumpsOKBodyItems
}

func (o *GetUniverseSystemJumpsOK) Error() string {
	return fmt.Sprintf("[GET /universe/system_jumps/][%d] getUniverseSystemJumpsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseSystemJumpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseSystemJumpsNotModified creates a GetUniverseSystemJumpsNotModified with default headers values
func NewGetUniverseSystemJumpsNotModified() *GetUniverseSystemJumpsNotModified {
	return &GetUniverseSystemJumpsNotModified{}
}

/*GetUniverseSystemJumpsNotModified handles this case with default header values.

Not modified
*/
type GetUniverseSystemJumpsNotModified struct {
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

func (o *GetUniverseSystemJumpsNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/system_jumps/][%d] getUniverseSystemJumpsNotModified ", 304)
}

func (o *GetUniverseSystemJumpsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseSystemJumpsBadRequest creates a GetUniverseSystemJumpsBadRequest with default headers values
func NewGetUniverseSystemJumpsBadRequest() *GetUniverseSystemJumpsBadRequest {
	return &GetUniverseSystemJumpsBadRequest{}
}

/*GetUniverseSystemJumpsBadRequest handles this case with default header values.

Bad request
*/
type GetUniverseSystemJumpsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseSystemJumpsBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/system_jumps/][%d] getUniverseSystemJumpsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseSystemJumpsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemJumpsEnhanceYourCalm creates a GetUniverseSystemJumpsEnhanceYourCalm with default headers values
func NewGetUniverseSystemJumpsEnhanceYourCalm() *GetUniverseSystemJumpsEnhanceYourCalm {
	return &GetUniverseSystemJumpsEnhanceYourCalm{}
}

/*GetUniverseSystemJumpsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetUniverseSystemJumpsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseSystemJumpsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/system_jumps/][%d] getUniverseSystemJumpsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseSystemJumpsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemJumpsInternalServerError creates a GetUniverseSystemJumpsInternalServerError with default headers values
func NewGetUniverseSystemJumpsInternalServerError() *GetUniverseSystemJumpsInternalServerError {
	return &GetUniverseSystemJumpsInternalServerError{}
}

/*GetUniverseSystemJumpsInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseSystemJumpsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseSystemJumpsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/system_jumps/][%d] getUniverseSystemJumpsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseSystemJumpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemJumpsServiceUnavailable creates a GetUniverseSystemJumpsServiceUnavailable with default headers values
func NewGetUniverseSystemJumpsServiceUnavailable() *GetUniverseSystemJumpsServiceUnavailable {
	return &GetUniverseSystemJumpsServiceUnavailable{}
}

/*GetUniverseSystemJumpsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetUniverseSystemJumpsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseSystemJumpsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/system_jumps/][%d] getUniverseSystemJumpsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseSystemJumpsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemJumpsGatewayTimeout creates a GetUniverseSystemJumpsGatewayTimeout with default headers values
func NewGetUniverseSystemJumpsGatewayTimeout() *GetUniverseSystemJumpsGatewayTimeout {
	return &GetUniverseSystemJumpsGatewayTimeout{}
}

/*GetUniverseSystemJumpsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetUniverseSystemJumpsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseSystemJumpsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/system_jumps/][%d] getUniverseSystemJumpsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseSystemJumpsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
