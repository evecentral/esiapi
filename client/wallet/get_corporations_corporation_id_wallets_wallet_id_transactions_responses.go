package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCorporationsCorporationIDWalletsWalletIDTransactionsReader is a Reader for the GetCorporationsCorporationIDWalletsWalletIDTransactions structure.
type GetCorporationsCorporationIDWalletsWalletIDTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent creates a GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent with default headers values
func NewGetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent() *GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent {
	return &GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent{}
}

/*GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent handles this case with default header values.

Nice Dummy
*/
type GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent struct {
}

func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{wallet_id}/transactions/][%d] getCorporationsCorporationIdWalletsWalletIdTransactionsNoContent ", 204)
}

func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError creates a GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError with default headers values
func NewGetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError() *GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError {
	return &GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError{}
}

/*GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError struct {
	Payload *models.GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError
}

func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{wallet_id}/transactions/][%d] getCorporationsCorporationIdWalletsWalletIdTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCorporationsCorporationIDWalletsWalletIDTransactionsInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
