package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetMarketsPricesReader is a Reader for the GetMarketsPrices structure.
type GetMarketsPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMarketsPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetMarketsPricesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMarketsPricesOK creates a GetMarketsPricesOK with default headers values
func NewGetMarketsPricesOK() *GetMarketsPricesOK {
	return &GetMarketsPricesOK{}
}

/*GetMarketsPricesOK handles this case with default header values.

A list of prices
*/
type GetMarketsPricesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetMarketsPrices200Ok
}

func (o *GetMarketsPricesOK) Error() string {
	return fmt.Sprintf("[GET /markets/prices/][%d] getMarketsPricesOK  %+v", 200, o.Payload)
}

func (o *GetMarketsPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsPricesInternalServerError creates a GetMarketsPricesInternalServerError with default headers values
func NewGetMarketsPricesInternalServerError() *GetMarketsPricesInternalServerError {
	return &GetMarketsPricesInternalServerError{}
}

/*GetMarketsPricesInternalServerError handles this case with default header values.

Internal server error
*/
type GetMarketsPricesInternalServerError struct {
	Payload *models.GetMarketsPricesInternalServerError
}

func (o *GetMarketsPricesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /markets/prices/][%d] getMarketsPricesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMarketsPricesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMarketsPricesInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
