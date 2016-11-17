package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetSearchReader is a Reader for the GetSearch structure.
type GetSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSearchOK creates a GetSearchOK with default headers values
func NewGetSearchOK() *GetSearchOK {
	return &GetSearchOK{}
}

/*GetSearchOK handles this case with default header values.

A list of search results
*/
type GetSearchOK struct {
	Payload *models.GetSearchOk
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchOK  %+v", 200, o.Payload)
}

func (o *GetSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSearchOk)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchInternalServerError creates a GetSearchInternalServerError with default headers values
func NewGetSearchInternalServerError() *GetSearchInternalServerError {
	return &GetSearchInternalServerError{}
}

/*GetSearchInternalServerError handles this case with default header values.

Internal server error
*/
type GetSearchInternalServerError struct {
	Payload *models.GetSearchInternalServerError
}

func (o *GetSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /search/][%d] getSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSearchInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
