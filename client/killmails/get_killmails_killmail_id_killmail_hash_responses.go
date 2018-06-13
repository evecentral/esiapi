// Code generated by go-swagger; DO NOT EDIT.

package killmails

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetKillmailsKillmailIDKillmailHashReader is a Reader for the GetKillmailsKillmailIDKillmailHash structure.
type GetKillmailsKillmailIDKillmailHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKillmailsKillmailIDKillmailHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKillmailsKillmailIDKillmailHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetKillmailsKillmailIDKillmailHashNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetKillmailsKillmailIDKillmailHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetKillmailsKillmailIDKillmailHashEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetKillmailsKillmailIDKillmailHashUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetKillmailsKillmailIDKillmailHashInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetKillmailsKillmailIDKillmailHashServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetKillmailsKillmailIDKillmailHashGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKillmailsKillmailIDKillmailHashOK creates a GetKillmailsKillmailIDKillmailHashOK with default headers values
func NewGetKillmailsKillmailIDKillmailHashOK() *GetKillmailsKillmailIDKillmailHashOK {
	return &GetKillmailsKillmailIDKillmailHashOK{}
}

/*GetKillmailsKillmailIDKillmailHashOK handles this case with default header values.

A killmail
*/
type GetKillmailsKillmailIDKillmailHashOK struct {
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

	Payload *models.GetKillmailsKillmailIDKillmailHashOKBody
}

func (o *GetKillmailsKillmailIDKillmailHashOK) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashOK  %+v", 200, o.Payload)
}

func (o *GetKillmailsKillmailIDKillmailHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetKillmailsKillmailIDKillmailHashOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKillmailsKillmailIDKillmailHashNotModified creates a GetKillmailsKillmailIDKillmailHashNotModified with default headers values
func NewGetKillmailsKillmailIDKillmailHashNotModified() *GetKillmailsKillmailIDKillmailHashNotModified {
	return &GetKillmailsKillmailIDKillmailHashNotModified{}
}

/*GetKillmailsKillmailIDKillmailHashNotModified handles this case with default header values.

Not modified
*/
type GetKillmailsKillmailIDKillmailHashNotModified struct {
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

func (o *GetKillmailsKillmailIDKillmailHashNotModified) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashNotModified ", 304)
}

func (o *GetKillmailsKillmailIDKillmailHashNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetKillmailsKillmailIDKillmailHashBadRequest creates a GetKillmailsKillmailIDKillmailHashBadRequest with default headers values
func NewGetKillmailsKillmailIDKillmailHashBadRequest() *GetKillmailsKillmailIDKillmailHashBadRequest {
	return &GetKillmailsKillmailIDKillmailHashBadRequest{}
}

/*GetKillmailsKillmailIDKillmailHashBadRequest handles this case with default header values.

Bad request
*/
type GetKillmailsKillmailIDKillmailHashBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetKillmailsKillmailIDKillmailHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetKillmailsKillmailIDKillmailHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKillmailsKillmailIDKillmailHashEnhanceYourCalm creates a GetKillmailsKillmailIDKillmailHashEnhanceYourCalm with default headers values
func NewGetKillmailsKillmailIDKillmailHashEnhanceYourCalm() *GetKillmailsKillmailIDKillmailHashEnhanceYourCalm {
	return &GetKillmailsKillmailIDKillmailHashEnhanceYourCalm{}
}

/*GetKillmailsKillmailIDKillmailHashEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetKillmailsKillmailIDKillmailHashEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetKillmailsKillmailIDKillmailHashEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetKillmailsKillmailIDKillmailHashEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKillmailsKillmailIDKillmailHashUnprocessableEntity creates a GetKillmailsKillmailIDKillmailHashUnprocessableEntity with default headers values
func NewGetKillmailsKillmailIDKillmailHashUnprocessableEntity() *GetKillmailsKillmailIDKillmailHashUnprocessableEntity {
	return &GetKillmailsKillmailIDKillmailHashUnprocessableEntity{}
}

/*GetKillmailsKillmailIDKillmailHashUnprocessableEntity handles this case with default header values.

Invalid killmail_id and/or killmail_hash
*/
type GetKillmailsKillmailIDKillmailHashUnprocessableEntity struct {
	Payload *models.GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody
}

func (o *GetKillmailsKillmailIDKillmailHashUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetKillmailsKillmailIDKillmailHashUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetKillmailsKillmailIDKillmailHashUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKillmailsKillmailIDKillmailHashInternalServerError creates a GetKillmailsKillmailIDKillmailHashInternalServerError with default headers values
func NewGetKillmailsKillmailIDKillmailHashInternalServerError() *GetKillmailsKillmailIDKillmailHashInternalServerError {
	return &GetKillmailsKillmailIDKillmailHashInternalServerError{}
}

/*GetKillmailsKillmailIDKillmailHashInternalServerError handles this case with default header values.

Internal server error
*/
type GetKillmailsKillmailIDKillmailHashInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetKillmailsKillmailIDKillmailHashInternalServerError) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKillmailsKillmailIDKillmailHashInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKillmailsKillmailIDKillmailHashServiceUnavailable creates a GetKillmailsKillmailIDKillmailHashServiceUnavailable with default headers values
func NewGetKillmailsKillmailIDKillmailHashServiceUnavailable() *GetKillmailsKillmailIDKillmailHashServiceUnavailable {
	return &GetKillmailsKillmailIDKillmailHashServiceUnavailable{}
}

/*GetKillmailsKillmailIDKillmailHashServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetKillmailsKillmailIDKillmailHashServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetKillmailsKillmailIDKillmailHashServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKillmailsKillmailIDKillmailHashServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKillmailsKillmailIDKillmailHashGatewayTimeout creates a GetKillmailsKillmailIDKillmailHashGatewayTimeout with default headers values
func NewGetKillmailsKillmailIDKillmailHashGatewayTimeout() *GetKillmailsKillmailIDKillmailHashGatewayTimeout {
	return &GetKillmailsKillmailIDKillmailHashGatewayTimeout{}
}

/*GetKillmailsKillmailIDKillmailHashGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetKillmailsKillmailIDKillmailHashGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetKillmailsKillmailIDKillmailHashGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /killmails/{killmail_id}/{killmail_hash}/][%d] getKillmailsKillmailIdKillmailHashGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKillmailsKillmailIDKillmailHashGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
