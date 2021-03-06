// Code generated by go-swagger; DO NOT EDIT.

package wars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetWarsWarIDReader is a Reader for the GetWarsWarID structure.
type GetWarsWarIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWarsWarIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWarsWarIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetWarsWarIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetWarsWarIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetWarsWarIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetWarsWarIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetWarsWarIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetWarsWarIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetWarsWarIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWarsWarIDOK creates a GetWarsWarIDOK with default headers values
func NewGetWarsWarIDOK() *GetWarsWarIDOK {
	return &GetWarsWarIDOK{}
}

/*GetWarsWarIDOK handles this case with default header values.

Details about a war
*/
type GetWarsWarIDOK struct {
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

	Payload *models.GetWarsWarIDOKBody
}

func (o *GetWarsWarIDOK) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdOK  %+v", 200, o.Payload)
}

func (o *GetWarsWarIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetWarsWarIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDNotModified creates a GetWarsWarIDNotModified with default headers values
func NewGetWarsWarIDNotModified() *GetWarsWarIDNotModified {
	return &GetWarsWarIDNotModified{}
}

/*GetWarsWarIDNotModified handles this case with default header values.

Not modified
*/
type GetWarsWarIDNotModified struct {
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

func (o *GetWarsWarIDNotModified) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdNotModified ", 304)
}

func (o *GetWarsWarIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWarsWarIDBadRequest creates a GetWarsWarIDBadRequest with default headers values
func NewGetWarsWarIDBadRequest() *GetWarsWarIDBadRequest {
	return &GetWarsWarIDBadRequest{}
}

/*GetWarsWarIDBadRequest handles this case with default header values.

Bad request
*/
type GetWarsWarIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetWarsWarIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetWarsWarIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDEnhanceYourCalm creates a GetWarsWarIDEnhanceYourCalm with default headers values
func NewGetWarsWarIDEnhanceYourCalm() *GetWarsWarIDEnhanceYourCalm {
	return &GetWarsWarIDEnhanceYourCalm{}
}

/*GetWarsWarIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetWarsWarIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetWarsWarIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetWarsWarIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDUnprocessableEntity creates a GetWarsWarIDUnprocessableEntity with default headers values
func NewGetWarsWarIDUnprocessableEntity() *GetWarsWarIDUnprocessableEntity {
	return &GetWarsWarIDUnprocessableEntity{}
}

/*GetWarsWarIDUnprocessableEntity handles this case with default header values.

War not found
*/
type GetWarsWarIDUnprocessableEntity struct {
	Payload *models.GetWarsWarIDUnprocessableEntityBody
}

func (o *GetWarsWarIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetWarsWarIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetWarsWarIDUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDInternalServerError creates a GetWarsWarIDInternalServerError with default headers values
func NewGetWarsWarIDInternalServerError() *GetWarsWarIDInternalServerError {
	return &GetWarsWarIDInternalServerError{}
}

/*GetWarsWarIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetWarsWarIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetWarsWarIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWarsWarIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDServiceUnavailable creates a GetWarsWarIDServiceUnavailable with default headers values
func NewGetWarsWarIDServiceUnavailable() *GetWarsWarIDServiceUnavailable {
	return &GetWarsWarIDServiceUnavailable{}
}

/*GetWarsWarIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetWarsWarIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetWarsWarIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWarsWarIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWarsWarIDGatewayTimeout creates a GetWarsWarIDGatewayTimeout with default headers values
func NewGetWarsWarIDGatewayTimeout() *GetWarsWarIDGatewayTimeout {
	return &GetWarsWarIDGatewayTimeout{}
}

/*GetWarsWarIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetWarsWarIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetWarsWarIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /wars/{war_id}/][%d] getWarsWarIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWarsWarIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
