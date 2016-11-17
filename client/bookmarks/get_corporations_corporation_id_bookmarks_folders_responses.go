package bookmarks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/esiswagger/models"
)

// GetCorporationsCorporationIDBookmarksFoldersReader is a Reader for the GetCorporationsCorporationIDBookmarksFolders structure.
type GetCorporationsCorporationIDBookmarksFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDBookmarksFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetCorporationsCorporationIDBookmarksFoldersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetCorporationsCorporationIDBookmarksFoldersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDBookmarksFoldersNoContent creates a GetCorporationsCorporationIDBookmarksFoldersNoContent with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersNoContent() *GetCorporationsCorporationIDBookmarksFoldersNoContent {
	return &GetCorporationsCorporationIDBookmarksFoldersNoContent{}
}

/*GetCorporationsCorporationIDBookmarksFoldersNoContent handles this case with default header values.

Nice Dummy
*/
type GetCorporationsCorporationIDBookmarksFoldersNoContent struct {
}

func (o *GetCorporationsCorporationIDBookmarksFoldersNoContent) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersNoContent ", 204)
}

func (o *GetCorporationsCorporationIDBookmarksFoldersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCorporationsCorporationIDBookmarksFoldersInternalServerError creates a GetCorporationsCorporationIDBookmarksFoldersInternalServerError with default headers values
func NewGetCorporationsCorporationIDBookmarksFoldersInternalServerError() *GetCorporationsCorporationIDBookmarksFoldersInternalServerError {
	return &GetCorporationsCorporationIDBookmarksFoldersInternalServerError{}
}

/*GetCorporationsCorporationIDBookmarksFoldersInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDBookmarksFoldersInternalServerError struct {
	Payload *models.GetCorporationsCorporationIDBookmarksFoldersInternalServerError
}

func (o *GetCorporationsCorporationIDBookmarksFoldersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/bookmarks/folders/][%d] getCorporationsCorporationIdBookmarksFoldersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDBookmarksFoldersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCorporationsCorporationIDBookmarksFoldersInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
