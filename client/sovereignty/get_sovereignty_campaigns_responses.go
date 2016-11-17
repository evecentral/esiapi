package sovereignty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetSovereigntyCampaignsReader is a Reader for the GetSovereigntyCampaigns structure.
type GetSovereigntyCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSovereigntyCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSovereigntyCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSovereigntyCampaignsOK creates a GetSovereigntyCampaignsOK with default headers values
func NewGetSovereigntyCampaignsOK() *GetSovereigntyCampaignsOK {
	return &GetSovereigntyCampaignsOK{}
}

/*GetSovereigntyCampaignsOK handles this case with default header values.

A list of sovereignty campaigns
*/
type GetSovereigntyCampaignsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetSovereigntyCampaigns200Ok
}

func (o *GetSovereigntyCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyCampaignsInternalServerError creates a GetSovereigntyCampaignsInternalServerError with default headers values
func NewGetSovereigntyCampaignsInternalServerError() *GetSovereigntyCampaignsInternalServerError {
	return &GetSovereigntyCampaignsInternalServerError{}
}

/*GetSovereigntyCampaignsInternalServerError handles this case with default header values.

Internal server error
*/
type GetSovereigntyCampaignsInternalServerError struct {
	Payload *models.GetSovereigntyCampaignsInternalServerError
}

func (o *GetSovereigntyCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetSovereigntyCampaignsInternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
