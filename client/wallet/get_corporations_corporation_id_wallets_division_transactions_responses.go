// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetCorporationsCorporationIDWalletsDivisionTransactionsReader is a Reader for the GetCorporationsCorporationIDWalletsDivisionTransactions structure.
type GetCorporationsCorporationIDWalletsDivisionTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsOK creates a GetCorporationsCorporationIDWalletsDivisionTransactionsOK with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsOK() *GetCorporationsCorporationIDWalletsDivisionTransactionsOK {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsOK{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsOK handles this case with default header values.

Wallet transactions
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsOK struct {
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

	Payload []*models.GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsNotModified creates a GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsNotModified() *GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified struct {
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

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest creates a GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest() *GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized creates a GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized() *GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsForbidden creates a GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsForbidden() *GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm creates a GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm() *GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError creates a GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError() *GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable creates a GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable() *GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout creates a GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout() *GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout{}
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}