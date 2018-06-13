// Code generated by go-swagger; DO NOT EDIT.

package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new mail API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mail API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCharactersCharacterIDMailLabelsLabelID deletes a mail label

Delete a mail label

---
Alternate route: `/dev/characters/{character_id}/mail/labels/{label_id}/`

Alternate route: `/legacy/characters/{character_id}/mail/labels/{label_id}/`

Alternate route: `/v1/characters/{character_id}/mail/labels/{label_id}/`

*/
func (a *Client) DeleteCharactersCharacterIDMailLabelsLabelID(params *DeleteCharactersCharacterIDMailLabelsLabelIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCharactersCharacterIDMailLabelsLabelIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCharactersCharacterIDMailLabelsLabelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_characters_character_id_mail_labels_label_id",
		Method:             "DELETE",
		PathPattern:        "/characters/{character_id}/mail/labels/{label_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCharactersCharacterIDMailLabelsLabelIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCharactersCharacterIDMailLabelsLabelIDNoContent), nil

}

/*
DeleteCharactersCharacterIDMailMailID deletes a mail

Delete a mail

---
Alternate route: `/dev/characters/{character_id}/mail/{mail_id}/`

Alternate route: `/legacy/characters/{character_id}/mail/{mail_id}/`

Alternate route: `/v1/characters/{character_id}/mail/{mail_id}/`

*/
func (a *Client) DeleteCharactersCharacterIDMailMailID(params *DeleteCharactersCharacterIDMailMailIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCharactersCharacterIDMailMailIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCharactersCharacterIDMailMailIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_characters_character_id_mail_mail_id",
		Method:             "DELETE",
		PathPattern:        "/characters/{character_id}/mail/{mail_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCharactersCharacterIDMailMailIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCharactersCharacterIDMailMailIDNoContent), nil

}

/*
GetCharactersCharacterIDMail returns mail headers

Return the 50 most recent mail headers belonging to the character that match the query criteria. Queries can be filtered by label, and last_mail_id can be used to paginate backwards.

---
Alternate route: `/dev/characters/{character_id}/mail/`

Alternate route: `/legacy/characters/{character_id}/mail/`

Alternate route: `/v1/characters/{character_id}/mail/`

---
This route is cached for up to 30 seconds
*/
func (a *Client) GetCharactersCharacterIDMail(params *GetCharactersCharacterIDMailParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDMailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDMailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_mail",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/mail/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDMailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDMailOK), nil

}

/*
GetCharactersCharacterIDMailLabels gets mail labels and unread counts

Return a list of the users mail labels, unread counts for each label and a total unread count.

---
Alternate route: `/dev/characters/{character_id}/mail/labels/`

Alternate route: `/v3/characters/{character_id}/mail/labels/`

---
This route is cached for up to 30 seconds
*/
func (a *Client) GetCharactersCharacterIDMailLabels(params *GetCharactersCharacterIDMailLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDMailLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDMailLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_mail_labels",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/mail/labels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDMailLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDMailLabelsOK), nil

}

/*
GetCharactersCharacterIDMailLists returns mailing list subscriptions

Return all mailing lists that the character is subscribed to

---
Alternate route: `/dev/characters/{character_id}/mail/lists/`

Alternate route: `/legacy/characters/{character_id}/mail/lists/`

Alternate route: `/v1/characters/{character_id}/mail/lists/`

---
This route is cached for up to 120 seconds
*/
func (a *Client) GetCharactersCharacterIDMailLists(params *GetCharactersCharacterIDMailListsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDMailListsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDMailListsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_mail_lists",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/mail/lists/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDMailListsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDMailListsOK), nil

}

/*
GetCharactersCharacterIDMailMailID returns a mail

Return the contents of an EVE mail

---
Alternate route: `/dev/characters/{character_id}/mail/{mail_id}/`

Alternate route: `/legacy/characters/{character_id}/mail/{mail_id}/`

Alternate route: `/v1/characters/{character_id}/mail/{mail_id}/`

---
This route is cached for up to 30 seconds
*/
func (a *Client) GetCharactersCharacterIDMailMailID(params *GetCharactersCharacterIDMailMailIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDMailMailIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDMailMailIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_mail_mail_id",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/mail/{mail_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDMailMailIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDMailMailIDOK), nil

}

/*
PostCharactersCharacterIDMail sends a new mail

Create and send a new mail

---
Alternate route: `/dev/characters/{character_id}/mail/`

Alternate route: `/legacy/characters/{character_id}/mail/`

Alternate route: `/v1/characters/{character_id}/mail/`

*/
func (a *Client) PostCharactersCharacterIDMail(params *PostCharactersCharacterIDMailParams, authInfo runtime.ClientAuthInfoWriter) (*PostCharactersCharacterIDMailCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharactersCharacterIDMailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_characters_character_id_mail",
		Method:             "POST",
		PathPattern:        "/characters/{character_id}/mail/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCharactersCharacterIDMailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCharactersCharacterIDMailCreated), nil

}

/*
PostCharactersCharacterIDMailLabels creates a mail label

Create a mail label

---
Alternate route: `/dev/characters/{character_id}/mail/labels/`

Alternate route: `/legacy/characters/{character_id}/mail/labels/`

Alternate route: `/v2/characters/{character_id}/mail/labels/`

*/
func (a *Client) PostCharactersCharacterIDMailLabels(params *PostCharactersCharacterIDMailLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCharactersCharacterIDMailLabelsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharactersCharacterIDMailLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_characters_character_id_mail_labels",
		Method:             "POST",
		PathPattern:        "/characters/{character_id}/mail/labels/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCharactersCharacterIDMailLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCharactersCharacterIDMailLabelsCreated), nil

}

/*
PutCharactersCharacterIDMailMailID updates metadata about a mail

Update metadata about a mail

---
Alternate route: `/dev/characters/{character_id}/mail/{mail_id}/`

Alternate route: `/legacy/characters/{character_id}/mail/{mail_id}/`

Alternate route: `/v1/characters/{character_id}/mail/{mail_id}/`

*/
func (a *Client) PutCharactersCharacterIDMailMailID(params *PutCharactersCharacterIDMailMailIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutCharactersCharacterIDMailMailIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCharactersCharacterIDMailMailIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "put_characters_character_id_mail_mail_id",
		Method:             "PUT",
		PathPattern:        "/characters/{character_id}/mail/{mail_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutCharactersCharacterIDMailMailIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCharactersCharacterIDMailMailIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
