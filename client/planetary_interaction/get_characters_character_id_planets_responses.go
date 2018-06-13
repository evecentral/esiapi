// Code generated by go-swagger; DO NOT EDIT.

package planetary_interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/evecentral/esiapi/models"
)

// GetCharactersCharacterIDPlanetsReader is a Reader for the GetCharactersCharacterIDPlanets structure.
type GetCharactersCharacterIDPlanetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDPlanetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDPlanetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDPlanetsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDPlanetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDPlanetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDPlanetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDPlanetsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDPlanetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDPlanetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDPlanetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDPlanetsOK creates a GetCharactersCharacterIDPlanetsOK with default headers values
func NewGetCharactersCharacterIDPlanetsOK() *GetCharactersCharacterIDPlanetsOK {
	return &GetCharactersCharacterIDPlanetsOK{}
}

/*GetCharactersCharacterIDPlanetsOK handles this case with default header values.

List of colonies
*/
type GetCharactersCharacterIDPlanetsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7232 compliant entity tag
	 */
	ETag string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetCharactersCharacterIDPlanetsOKBodyItems
}

func (o *GetCharactersCharacterIDPlanetsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

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

// NewGetCharactersCharacterIDPlanetsNotModified creates a GetCharactersCharacterIDPlanetsNotModified with default headers values
func NewGetCharactersCharacterIDPlanetsNotModified() *GetCharactersCharacterIDPlanetsNotModified {
	return &GetCharactersCharacterIDPlanetsNotModified{}
}

/*GetCharactersCharacterIDPlanetsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDPlanetsNotModified struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7232 compliant entity tag
	 */
	ETag string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string
}

func (o *GetCharactersCharacterIDPlanetsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsNotModified ", 304)
}

func (o *GetCharactersCharacterIDPlanetsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	return nil
}

// NewGetCharactersCharacterIDPlanetsBadRequest creates a GetCharactersCharacterIDPlanetsBadRequest with default headers values
func NewGetCharactersCharacterIDPlanetsBadRequest() *GetCharactersCharacterIDPlanetsBadRequest {
	return &GetCharactersCharacterIDPlanetsBadRequest{}
}

/*GetCharactersCharacterIDPlanetsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDPlanetsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDPlanetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsUnauthorized creates a GetCharactersCharacterIDPlanetsUnauthorized with default headers values
func NewGetCharactersCharacterIDPlanetsUnauthorized() *GetCharactersCharacterIDPlanetsUnauthorized {
	return &GetCharactersCharacterIDPlanetsUnauthorized{}
}

/*GetCharactersCharacterIDPlanetsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDPlanetsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDPlanetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsForbidden creates a GetCharactersCharacterIDPlanetsForbidden with default headers values
func NewGetCharactersCharacterIDPlanetsForbidden() *GetCharactersCharacterIDPlanetsForbidden {
	return &GetCharactersCharacterIDPlanetsForbidden{}
}

/*GetCharactersCharacterIDPlanetsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDPlanetsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDPlanetsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsEnhanceYourCalm creates a GetCharactersCharacterIDPlanetsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDPlanetsEnhanceYourCalm() *GetCharactersCharacterIDPlanetsEnhanceYourCalm {
	return &GetCharactersCharacterIDPlanetsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDPlanetsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDPlanetsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDPlanetsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsInternalServerError creates a GetCharactersCharacterIDPlanetsInternalServerError with default headers values
func NewGetCharactersCharacterIDPlanetsInternalServerError() *GetCharactersCharacterIDPlanetsInternalServerError {
	return &GetCharactersCharacterIDPlanetsInternalServerError{}
}

/*GetCharactersCharacterIDPlanetsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDPlanetsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDPlanetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsServiceUnavailable creates a GetCharactersCharacterIDPlanetsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDPlanetsServiceUnavailable() *GetCharactersCharacterIDPlanetsServiceUnavailable {
	return &GetCharactersCharacterIDPlanetsServiceUnavailable{}
}

/*GetCharactersCharacterIDPlanetsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDPlanetsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDPlanetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPlanetsGatewayTimeout creates a GetCharactersCharacterIDPlanetsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDPlanetsGatewayTimeout() *GetCharactersCharacterIDPlanetsGatewayTimeout {
	return &GetCharactersCharacterIDPlanetsGatewayTimeout{}
}

/*GetCharactersCharacterIDPlanetsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDPlanetsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDPlanetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/planets/][%d] getCharactersCharacterIdPlanetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDPlanetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
