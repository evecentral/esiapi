package live

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCharactersCharacterIDWalletsReader is a Reader for the GetCharactersCharacterIDWallets structure.
type GetCharactersCharacterIDWalletsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDWalletsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDWalletsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCharactersCharacterIDWalletsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDWalletsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDWalletsOK creates a GetCharactersCharacterIDWalletsOK with default headers values
func NewGetCharactersCharacterIDWalletsOK() *GetCharactersCharacterIDWalletsOK {
	return &GetCharactersCharacterIDWalletsOK{}
}

/*GetCharactersCharacterIDWalletsOK handles this case with default header values.

Wallet data for selected user
*/
type GetCharactersCharacterIDWalletsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetCharactersCharacterIDWallets200Ok
}

func (o *GetCharactersCharacterIDWalletsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallets/][%d] getCharactersCharacterIdWalletsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDWalletsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

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

// NewGetCharactersCharacterIDWalletsForbidden creates a GetCharactersCharacterIDWalletsForbidden with default headers values
func NewGetCharactersCharacterIDWalletsForbidden() *GetCharactersCharacterIDWalletsForbidden {
	return &GetCharactersCharacterIDWalletsForbidden{}
}

/*GetCharactersCharacterIDWalletsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDWalletsForbidden struct {
	Payload *models.GetCharactersCharacterIDWalletsForbidden
}

func (o *GetCharactersCharacterIDWalletsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallets/][%d] getCharactersCharacterIdWalletsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDWalletsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDWalletsForbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletsInternalServerError creates a GetCharactersCharacterIDWalletsInternalServerError with default headers values
func NewGetCharactersCharacterIDWalletsInternalServerError() *GetCharactersCharacterIDWalletsInternalServerError {
	return &GetCharactersCharacterIDWalletsInternalServerError{}
}

/*GetCharactersCharacterIDWalletsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDWalletsInternalServerError struct {
	Payload *models.GetCharactersCharacterIDWalletsInternalServerError
}

func (o *GetCharactersCharacterIDWalletsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallets/][%d] getCharactersCharacterIdWalletsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDWalletsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDWalletsInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
