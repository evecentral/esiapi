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

// GetIncursionsReader is a Reader for the GetIncursions structure.
type GetIncursionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIncursionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIncursionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetIncursionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIncursionsOK creates a GetIncursionsOK with default headers values
func NewGetIncursionsOK() *GetIncursionsOK {
	return &GetIncursionsOK{}
}

/*GetIncursionsOK handles this case with default header values.

A list of incursions
*/
type GetIncursionsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetIncursions200Ok
}

func (o *GetIncursionsOK) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsOK  %+v", 200, o.Payload)
}

func (o *GetIncursionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIncursionsInternalServerError creates a GetIncursionsInternalServerError with default headers values
func NewGetIncursionsInternalServerError() *GetIncursionsInternalServerError {
	return &GetIncursionsInternalServerError{}
}

/*GetIncursionsInternalServerError handles this case with default header values.

Internal server error
*/
type GetIncursionsInternalServerError struct {
	Payload *models.GetIncursionsInternalServerError
}

func (o *GetIncursionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /incursions/][%d] getIncursionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIncursionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIncursionsInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
