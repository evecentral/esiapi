package dummy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCorporationsCorporationIDWalletsWalletIDJournalReader is a Reader for the GetCorporationsCorporationIDWalletsWalletIDJournal structure.
type GetCorporationsCorporationIDWalletsWalletIDJournalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDWalletsWalletIDJournalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetCorporationsCorporationIDWalletsWalletIDJournalNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDWalletsWalletIDJournalNoContent creates a GetCorporationsCorporationIDWalletsWalletIDJournalNoContent with default headers values
func NewGetCorporationsCorporationIDWalletsWalletIDJournalNoContent() *GetCorporationsCorporationIDWalletsWalletIDJournalNoContent {
	return &GetCorporationsCorporationIDWalletsWalletIDJournalNoContent{}
}

/*GetCorporationsCorporationIDWalletsWalletIDJournalNoContent handles this case with default header values.

Nice Dummy
*/
type GetCorporationsCorporationIDWalletsWalletIDJournalNoContent struct {
}

func (o *GetCorporationsCorporationIDWalletsWalletIDJournalNoContent) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{wallet_id}/journal/][%d] getCorporationsCorporationIdWalletsWalletIdJournalNoContent ", 204)
}

func (o *GetCorporationsCorporationIDWalletsWalletIDJournalNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError creates a GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError with default headers values
func NewGetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError() *GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError {
	return &GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError{}
}

/*GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError struct {
	Payload *models.GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError
}

func (o *GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/wallets/{wallet_id}/journal/][%d] getCorporationsCorporationIdWalletsWalletIdJournalInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCorporationsCorporationIDWalletsWalletIDJournalInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
