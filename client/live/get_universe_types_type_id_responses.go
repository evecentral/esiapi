package live

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetUniverseTypesTypeIDReader is a Reader for the GetUniverseTypesTypeID structure.
type GetUniverseTypesTypeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseTypesTypeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseTypesTypeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUniverseTypesTypeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseTypesTypeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseTypesTypeIDOK creates a GetUniverseTypesTypeIDOK with default headers values
func NewGetUniverseTypesTypeIDOK() *GetUniverseTypesTypeIDOK {
	return &GetUniverseTypesTypeIDOK{}
}

/*GetUniverseTypesTypeIDOK handles this case with default header values.

Public data about a type
*/
type GetUniverseTypesTypeIDOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *models.GetUniverseTypesTypeIDOk
}

func (o *GetUniverseTypesTypeIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseTypesTypeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniverseTypesTypeIDOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDNotFound creates a GetUniverseTypesTypeIDNotFound with default headers values
func NewGetUniverseTypesTypeIDNotFound() *GetUniverseTypesTypeIDNotFound {
	return &GetUniverseTypesTypeIDNotFound{}
}

/*GetUniverseTypesTypeIDNotFound handles this case with default header values.

Type not found
*/
type GetUniverseTypesTypeIDNotFound struct {
	Payload *models.GetUniverseTypesTypeIDNotFound
}

func (o *GetUniverseTypesTypeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseTypesTypeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseTypesTypeIDNotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesTypeIDInternalServerError creates a GetUniverseTypesTypeIDInternalServerError with default headers values
func NewGetUniverseTypesTypeIDInternalServerError() *GetUniverseTypesTypeIDInternalServerError {
	return &GetUniverseTypesTypeIDInternalServerError{}
}

/*GetUniverseTypesTypeIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseTypesTypeIDInternalServerError struct {
	Payload *models.GetUniverseTypesTypeIDInternalServerError
}

func (o *GetUniverseTypesTypeIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/types/{type_id}/][%d] getUniverseTypesTypeIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseTypesTypeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseTypesTypeIDInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}