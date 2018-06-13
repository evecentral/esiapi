// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// GetAlliancesAllianceIDContactsReader is a Reader for the GetAlliancesAllianceIDContacts structure.
type GetAlliancesAllianceIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesAllianceIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlliancesAllianceIDContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetAlliancesAllianceIDContactsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetAlliancesAllianceIDContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAlliancesAllianceIDContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAlliancesAllianceIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetAlliancesAllianceIDContactsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAlliancesAllianceIDContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAlliancesAllianceIDContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetAlliancesAllianceIDContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlliancesAllianceIDContactsOK creates a GetAlliancesAllianceIDContactsOK with default headers values
func NewGetAlliancesAllianceIDContactsOK() *GetAlliancesAllianceIDContactsOK {
	return &GetAlliancesAllianceIDContactsOK{
		XPages: 1,
	}
}

/*GetAlliancesAllianceIDContactsOK handles this case with default header values.

A list of contacts
*/
type GetAlliancesAllianceIDContactsOK struct {
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

	Payload []*models.GetAlliancesAllianceIDContactsOKBodyItems
}

func (o *GetAlliancesAllianceIDContactsOK) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsOK  %+v", 200, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDContactsNotModified creates a GetAlliancesAllianceIDContactsNotModified with default headers values
func NewGetAlliancesAllianceIDContactsNotModified() *GetAlliancesAllianceIDContactsNotModified {
	return &GetAlliancesAllianceIDContactsNotModified{}
}

/*GetAlliancesAllianceIDContactsNotModified handles this case with default header values.

Not modified
*/
type GetAlliancesAllianceIDContactsNotModified struct {
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

func (o *GetAlliancesAllianceIDContactsNotModified) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsNotModified ", 304)
}

func (o *GetAlliancesAllianceIDContactsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDContactsBadRequest creates a GetAlliancesAllianceIDContactsBadRequest with default headers values
func NewGetAlliancesAllianceIDContactsBadRequest() *GetAlliancesAllianceIDContactsBadRequest {
	return &GetAlliancesAllianceIDContactsBadRequest{}
}

/*GetAlliancesAllianceIDContactsBadRequest handles this case with default header values.

Bad request
*/
type GetAlliancesAllianceIDContactsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetAlliancesAllianceIDContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsUnauthorized creates a GetAlliancesAllianceIDContactsUnauthorized with default headers values
func NewGetAlliancesAllianceIDContactsUnauthorized() *GetAlliancesAllianceIDContactsUnauthorized {
	return &GetAlliancesAllianceIDContactsUnauthorized{}
}

/*GetAlliancesAllianceIDContactsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAlliancesAllianceIDContactsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetAlliancesAllianceIDContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsForbidden creates a GetAlliancesAllianceIDContactsForbidden with default headers values
func NewGetAlliancesAllianceIDContactsForbidden() *GetAlliancesAllianceIDContactsForbidden {
	return &GetAlliancesAllianceIDContactsForbidden{}
}

/*GetAlliancesAllianceIDContactsForbidden handles this case with default header values.

Forbidden
*/
type GetAlliancesAllianceIDContactsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetAlliancesAllianceIDContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsEnhanceYourCalm creates a GetAlliancesAllianceIDContactsEnhanceYourCalm with default headers values
func NewGetAlliancesAllianceIDContactsEnhanceYourCalm() *GetAlliancesAllianceIDContactsEnhanceYourCalm {
	return &GetAlliancesAllianceIDContactsEnhanceYourCalm{}
}

/*GetAlliancesAllianceIDContactsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetAlliancesAllianceIDContactsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsInternalServerError creates a GetAlliancesAllianceIDContactsInternalServerError with default headers values
func NewGetAlliancesAllianceIDContactsInternalServerError() *GetAlliancesAllianceIDContactsInternalServerError {
	return &GetAlliancesAllianceIDContactsInternalServerError{}
}

/*GetAlliancesAllianceIDContactsInternalServerError handles this case with default header values.

Internal server error
*/
type GetAlliancesAllianceIDContactsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAlliancesAllianceIDContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsServiceUnavailable creates a GetAlliancesAllianceIDContactsServiceUnavailable with default headers values
func NewGetAlliancesAllianceIDContactsServiceUnavailable() *GetAlliancesAllianceIDContactsServiceUnavailable {
	return &GetAlliancesAllianceIDContactsServiceUnavailable{}
}

/*GetAlliancesAllianceIDContactsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetAlliancesAllianceIDContactsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetAlliancesAllianceIDContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsGatewayTimeout creates a GetAlliancesAllianceIDContactsGatewayTimeout with default headers values
func NewGetAlliancesAllianceIDContactsGatewayTimeout() *GetAlliancesAllianceIDContactsGatewayTimeout {
	return &GetAlliancesAllianceIDContactsGatewayTimeout{}
}

/*GetAlliancesAllianceIDContactsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetAlliancesAllianceIDContactsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetAlliancesAllianceIDContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}