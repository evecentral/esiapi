// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// GetCorporationsCorporationIDWalletsDivisionJournalReader is a Reader for the GetCorporationsCorporationIDWalletsDivisionJournal structure.
type GetCorporationsCorporationIDWalletsDivisionJournalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDWalletsDivisionJournalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalOK creates a GetCorporationsCorporationIDWalletsDivisionJournalOK with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalOK() *GetCorporationsCorporationIDWalletsDivisionJournalOK {
	return &GetCorporationsCorporationIDWalletsDivisionJournalOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalOK handles this case with default header values.

Journal entries
*/
type GetCorporationsCorporationIDWalletsDivisionJournalOK struct {
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

	Payload []*models.GetCorporationsCorporationIDWalletsDivisionJournalOKBodyItems
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDWalletsDivisionJournalNotModified creates a GetCorporationsCorporationIDWalletsDivisionJournalNotModified with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalNotModified() *GetCorporationsCorporationIDWalletsDivisionJournalNotModified {
	return &GetCorporationsCorporationIDWalletsDivisionJournalNotModified{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDWalletsDivisionJournalNotModified struct {
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

func (o *GetCorporationsCorporationIDWalletsDivisionJournalNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalNotModified ", 304)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDWalletsDivisionJournalBadRequest creates a GetCorporationsCorporationIDWalletsDivisionJournalBadRequest with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalBadRequest() *GetCorporationsCorporationIDWalletsDivisionJournalBadRequest {
	return &GetCorporationsCorporationIDWalletsDivisionJournalBadRequest{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDWalletsDivisionJournalBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalUnauthorized creates a GetCorporationsCorporationIDWalletsDivisionJournalUnauthorized with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalUnauthorized() *GetCorporationsCorporationIDWalletsDivisionJournalUnauthorized {
	return &GetCorporationsCorporationIDWalletsDivisionJournalUnauthorized{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDWalletsDivisionJournalUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalForbidden creates a GetCorporationsCorporationIDWalletsDivisionJournalForbidden with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalForbidden() *GetCorporationsCorporationIDWalletsDivisionJournalForbidden {
	return &GetCorporationsCorporationIDWalletsDivisionJournalForbidden{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDWalletsDivisionJournalForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm creates a GetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm() *GetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm {
	return &GetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalInternalServerError creates a GetCorporationsCorporationIDWalletsDivisionJournalInternalServerError with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalInternalServerError() *GetCorporationsCorporationIDWalletsDivisionJournalInternalServerError {
	return &GetCorporationsCorporationIDWalletsDivisionJournalInternalServerError{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDWalletsDivisionJournalInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable creates a GetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable() *GetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable {
	return &GetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout creates a GetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout() *GetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout {
	return &GetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout{}
}

/*GetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/journal/][%d] getCorporationsCorporationIdWalletsDivisionJournalGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionJournalGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
