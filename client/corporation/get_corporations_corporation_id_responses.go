package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCorporationsCorporationIDReader is a Reader for the GetCorporationsCorporationID structure.
type GetCorporationsCorporationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCorporationsCorporationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDOK creates a GetCorporationsCorporationIDOK with default headers values
func NewGetCorporationsCorporationIDOK() *GetCorporationsCorporationIDOK {
	return &GetCorporationsCorporationIDOK{}
}

/*GetCorporationsCorporationIDOK handles this case with default header values.

Public data about a corporation
*/
type GetCorporationsCorporationIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetCorporationsCorporationIDOk
}

func (o *GetCorporationsCorporationIDOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/][%d] getCorporationsCorporationIdOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCorporationsCorporationIDOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDNotFound creates a GetCorporationsCorporationIDNotFound with default headers values
func NewGetCorporationsCorporationIDNotFound() *GetCorporationsCorporationIDNotFound {
	return &GetCorporationsCorporationIDNotFound{}
}

/*GetCorporationsCorporationIDNotFound handles this case with default header values.

Corporation not found
*/
type GetCorporationsCorporationIDNotFound struct {
	Payload *models.GetCorporationsCorporationIDNotFound
}

func (o *GetCorporationsCorporationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/][%d] getCorporationsCorporationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCorporationsCorporationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCorporationsCorporationIDNotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDInternalServerError creates a GetCorporationsCorporationIDInternalServerError with default headers values
func NewGetCorporationsCorporationIDInternalServerError() *GetCorporationsCorporationIDInternalServerError {
	return &GetCorporationsCorporationIDInternalServerError{}
}

/*GetCorporationsCorporationIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDInternalServerError struct {
	Payload *models.GetCorporationsCorporationIDInternalServerError
}

func (o *GetCorporationsCorporationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/][%d] getCorporationsCorporationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCorporationsCorporationIDInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
