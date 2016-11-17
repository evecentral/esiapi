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

// PostUniverseNamesReader is a Reader for the PostUniverseNames structure.
type PostUniverseNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUniverseNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostUniverseNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostUniverseNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostUniverseNamesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUniverseNamesOK creates a PostUniverseNamesOK with default headers values
func NewPostUniverseNamesOK() *PostUniverseNamesOK {
	return &PostUniverseNamesOK{}
}

/*PostUniverseNamesOK handles this case with default header values.

List of id/name associations for a set of ID's. ID's that cannot be resolved are not returned.
*/
type PostUniverseNamesOK struct {
	Payload []*models.PostUniverseNames200Ok
}

func (o *PostUniverseNamesOK) Error() string {
	return fmt.Sprintf("[POST /universe/names/][%d] postUniverseNamesOK  %+v", 200, o.Payload)
}

func (o *PostUniverseNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUniverseNamesNotFound creates a PostUniverseNamesNotFound with default headers values
func NewPostUniverseNamesNotFound() *PostUniverseNamesNotFound {
	return &PostUniverseNamesNotFound{}
}

/*PostUniverseNamesNotFound handles this case with default header values.

no valid IDs found
*/
type PostUniverseNamesNotFound struct {
	Payload *models.PostUniverseNamesNotFound
}

func (o *PostUniverseNamesNotFound) Error() string {
	return fmt.Sprintf("[POST /universe/names/][%d] postUniverseNamesNotFound  %+v", 404, o.Payload)
}

func (o *PostUniverseNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostUniverseNamesNotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUniverseNamesInternalServerError creates a PostUniverseNamesInternalServerError with default headers values
func NewPostUniverseNamesInternalServerError() *PostUniverseNamesInternalServerError {
	return &PostUniverseNamesInternalServerError{}
}

/*PostUniverseNamesInternalServerError handles this case with default header values.

Internal server error
*/
type PostUniverseNamesInternalServerError struct {
	Payload *models.PostUniverseNamesInternalServerError
}

func (o *PostUniverseNamesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /universe/names/][%d] postUniverseNamesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUniverseNamesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostUniverseNamesInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
