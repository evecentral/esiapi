package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCharactersCharacterIDBookmarksFoldersInternalServerError Internal server error
// swagger:model get_characters_character_id_bookmarks_folders_internal_server_error
type GetCharactersCharacterIDBookmarksFoldersInternalServerError struct {

	// get_characters_character_id_bookmarks_folders_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character id bookmarks folders internal server error
func (m *GetCharactersCharacterIDBookmarksFoldersInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
