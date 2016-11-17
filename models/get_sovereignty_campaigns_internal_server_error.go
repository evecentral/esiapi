package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetSovereigntyCampaignsInternalServerError Internal server error
// swagger:model get_sovereignty_campaigns_internal_server_error
type GetSovereigntyCampaignsInternalServerError struct {

	// get_sovereignty_campaigns_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get sovereignty campaigns internal server error
func (m *GetSovereigntyCampaignsInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}