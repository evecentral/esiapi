// Code generated by go-swagger; DO NOT EDIT.

package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetCorporationsCorporationIDMembersReader is a Reader for the GetCorporationsCorporationIDMembers structure.
type GetCorporationsCorporationIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDMembersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDMembersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMembersOK creates a GetCorporationsCorporationIDMembersOK with default headers values
func NewGetCorporationsCorporationIDMembersOK() *GetCorporationsCorporationIDMembersOK {
	return &GetCorporationsCorporationIDMembersOK{}
}

/*GetCorporationsCorporationIDMembersOK handles this case with default header values.

List of member character IDs
*/
type GetCorporationsCorporationIDMembersOK struct {
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

func (o *GetCorporationsCorporationIDMembersOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersNotModified creates a GetCorporationsCorporationIDMembersNotModified with default headers values
func NewGetCorporationsCorporationIDMembersNotModified() *GetCorporationsCorporationIDMembersNotModified {
	return &GetCorporationsCorporationIDMembersNotModified{}
}

/*GetCorporationsCorporationIDMembersNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDMembersNotModified struct {
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

func (o *GetCorporationsCorporationIDMembersNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersBadRequest creates a GetCorporationsCorporationIDMembersBadRequest with default headers values
func NewGetCorporationsCorporationIDMembersBadRequest() *GetCorporationsCorporationIDMembersBadRequest {
	return &GetCorporationsCorporationIDMembersBadRequest{}
}

/*GetCorporationsCorporationIDMembersBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDMembersBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersUnauthorized creates a GetCorporationsCorporationIDMembersUnauthorized with default headers values
func NewGetCorporationsCorporationIDMembersUnauthorized() *GetCorporationsCorporationIDMembersUnauthorized {
	return &GetCorporationsCorporationIDMembersUnauthorized{}
}

/*GetCorporationsCorporationIDMembersUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMembersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersForbidden creates a GetCorporationsCorporationIDMembersForbidden with default headers values
func NewGetCorporationsCorporationIDMembersForbidden() *GetCorporationsCorporationIDMembersForbidden {
	return &GetCorporationsCorporationIDMembersForbidden{}
}

/*GetCorporationsCorporationIDMembersForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMembersForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersEnhanceYourCalm creates a GetCorporationsCorporationIDMembersEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMembersEnhanceYourCalm() *GetCorporationsCorporationIDMembersEnhanceYourCalm {
	return &GetCorporationsCorporationIDMembersEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDMembersEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDMembersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersInternalServerError creates a GetCorporationsCorporationIDMembersInternalServerError with default headers values
func NewGetCorporationsCorporationIDMembersInternalServerError() *GetCorporationsCorporationIDMembersInternalServerError {
	return &GetCorporationsCorporationIDMembersInternalServerError{}
}

/*GetCorporationsCorporationIDMembersInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMembersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersServiceUnavailable creates a GetCorporationsCorporationIDMembersServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMembersServiceUnavailable() *GetCorporationsCorporationIDMembersServiceUnavailable {
	return &GetCorporationsCorporationIDMembersServiceUnavailable{}
}

/*GetCorporationsCorporationIDMembersServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMembersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersGatewayTimeout creates a GetCorporationsCorporationIDMembersGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMembersGatewayTimeout() *GetCorporationsCorporationIDMembersGatewayTimeout {
	return &GetCorporationsCorporationIDMembersGatewayTimeout{}
}

/*GetCorporationsCorporationIDMembersGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMembersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
