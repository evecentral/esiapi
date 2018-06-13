// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/models"
)

// GetFwSystemsReader is a Reader for the GetFwSystems structure.
type GetFwSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFwSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFwSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetFwSystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFwSystemsOK creates a GetFwSystemsOK with default headers values
func NewGetFwSystemsOK() *GetFwSystemsOK {
	return &GetFwSystemsOK{}
}

/*GetFwSystemsOK handles this case with default header values.

All faction war solar systems
*/
type GetFwSystemsOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetFwSystemsOKBodyItems0
}

func (o *GetFwSystemsOK) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsOK  %+v", 200, o.Payload)
}

func (o *GetFwSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFwSystemsInternalServerError creates a GetFwSystemsInternalServerError with default headers values
func NewGetFwSystemsInternalServerError() *GetFwSystemsInternalServerError {
	return &GetFwSystemsInternalServerError{}
}

/*GetFwSystemsInternalServerError handles this case with default header values.

Internal server error
*/
type GetFwSystemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetFwSystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /fw/systems/][%d] getFwSystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFwSystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFwSystemsOKBodyItems0 get_fw_systems_200_ok
//
// 200 ok object
swagger:model GetFwSystemsOKBodyItems0
*/

type GetFwSystemsOKBodyItems0 struct {

	// get_fw_systems_contested
	//
	// contested boolean
	// Required: true
	Contested *bool `json:"contested"`

	// get_fw_systems_occupier_faction_id
	//
	// occupier_faction_id integer
	// Required: true
	OccupierFactionID *int32 `json:"occupier_faction_id"`

	// get_fw_systems_owner_faction_id
	//
	// owner_faction_id integer
	// Required: true
	OwnerFactionID *int32 `json:"owner_faction_id"`

	// get_fw_systems_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_fw_systems_victory_points
	//
	// victory_points integer
	// Required: true
	VictoryPoints *int32 `json:"victory_points"`

	// get_fw_systems_victory_points_threshold
	//
	// victory_points_threshold integer
	// Required: true
	VictoryPointsThreshold *int32 `json:"victory_points_threshold"`
}

/* polymorph GetFwSystemsOKBodyItems0 contested false */

/* polymorph GetFwSystemsOKBodyItems0 occupier_faction_id false */

/* polymorph GetFwSystemsOKBodyItems0 owner_faction_id false */

/* polymorph GetFwSystemsOKBodyItems0 solar_system_id false */

/* polymorph GetFwSystemsOKBodyItems0 victory_points false */

/* polymorph GetFwSystemsOKBodyItems0 victory_points_threshold false */

// Validate validates this get fw systems o k body items0
func (o *GetFwSystemsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContested(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOccupierFactionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateOwnerFactionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateVictoryPoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateVictoryPointsThreshold(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateContested(formats strfmt.Registry) error {

	if err := validate.Required("contested", "body", o.Contested); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateOccupierFactionID(formats strfmt.Registry) error {

	if err := validate.Required("occupier_faction_id", "body", o.OccupierFactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateOwnerFactionID(formats strfmt.Registry) error {

	if err := validate.Required("owner_faction_id", "body", o.OwnerFactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateVictoryPoints(formats strfmt.Registry) error {

	if err := validate.Required("victory_points", "body", o.VictoryPoints); err != nil {
		return err
	}

	return nil
}

func (o *GetFwSystemsOKBodyItems0) validateVictoryPointsThreshold(formats strfmt.Registry) error {

	if err := validate.Required("victory_points_threshold", "body", o.VictoryPointsThreshold); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFwSystemsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFwSystemsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetFwSystemsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}