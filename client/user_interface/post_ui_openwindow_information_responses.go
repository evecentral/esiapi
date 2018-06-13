// Code generated by go-swagger; DO NOT EDIT.

package user_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// PostUIOpenwindowInformationReader is a Reader for the PostUIOpenwindowInformation structure.
type PostUIOpenwindowInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIOpenwindowInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostUIOpenwindowInformationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewPostUIOpenwindowInformationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostUIOpenwindowInformationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUIOpenwindowInformationNoContent creates a PostUIOpenwindowInformationNoContent with default headers values
func NewPostUIOpenwindowInformationNoContent() *PostUIOpenwindowInformationNoContent {
	return &PostUIOpenwindowInformationNoContent{}
}

/*PostUIOpenwindowInformationNoContent handles this case with default header values.

Open window request received
*/
type PostUIOpenwindowInformationNoContent struct {
}

func (o *PostUIOpenwindowInformationNoContent) Error() string {
	return fmt.Sprintf("[POST /ui/openwindow/information/][%d] postUiOpenwindowInformationNoContent ", 204)
}

func (o *PostUIOpenwindowInformationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIOpenwindowInformationForbidden creates a PostUIOpenwindowInformationForbidden with default headers values
func NewPostUIOpenwindowInformationForbidden() *PostUIOpenwindowInformationForbidden {
	return &PostUIOpenwindowInformationForbidden{}
}

/*PostUIOpenwindowInformationForbidden handles this case with default header values.

Forbidden
*/
type PostUIOpenwindowInformationForbidden struct {
	Payload *models.Forbidden
}

func (o *PostUIOpenwindowInformationForbidden) Error() string {
	return fmt.Sprintf("[POST /ui/openwindow/information/][%d] postUiOpenwindowInformationForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowInformationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowInformationInternalServerError creates a PostUIOpenwindowInformationInternalServerError with default headers values
func NewPostUIOpenwindowInformationInternalServerError() *PostUIOpenwindowInformationInternalServerError {
	return &PostUIOpenwindowInformationInternalServerError{}
}

/*PostUIOpenwindowInformationInternalServerError handles this case with default header values.

Internal server error
*/
type PostUIOpenwindowInformationInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostUIOpenwindowInformationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ui/openwindow/information/][%d] postUiOpenwindowInformationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowInformationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}