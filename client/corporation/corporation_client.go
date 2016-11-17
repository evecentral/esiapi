package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new corporation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for corporation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCorporationsCorporationID gets corporation information

Public information about a corporation

---

Alternate route: `/v2/corporations/{corporation_id}/`

Alternate route: `/dev/corporations/{corporation_id}/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationID(params *GetCorporationsCorporationIDParams) (*GetCorporationsCorporationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDOK), nil

}

/*
GetCorporationsCorporationIDAlliancehistory gets alliance history

Get a list of all the alliances a corporation has been a member of

---

Alternate route: `/v1/corporations/{corporation_id}/alliancehistory/`

Alternate route: `/legacy/corporations/{corporation_id}/alliancehistory/`

Alternate route: `/dev/corporations/{corporation_id}/alliancehistory/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDAlliancehistory(params *GetCorporationsCorporationIDAlliancehistoryParams) (*GetCorporationsCorporationIDAlliancehistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDAlliancehistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_alliancehistory",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/alliancehistory/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDAlliancehistoryReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDAlliancehistoryOK), nil

}

/*
GetCorporationsCorporationIDIcons gets corporation icon

Get the icon urls for a corporation

---

Alternate route: `/v1/corporations/{corporation_id}/icons/`

Alternate route: `/legacy/corporations/{corporation_id}/icons/`

Alternate route: `/dev/corporations/{corporation_id}/icons/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDIcons(params *GetCorporationsCorporationIDIconsParams) (*GetCorporationsCorporationIDIconsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDIconsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_icons",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/icons/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDIconsReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDIconsOK), nil

}

/*
GetCorporationsCorporationIDMembers gets corporation members

Read the current list of members if the calling character is a member.

---

Alternate route: `/v2/corporations/{corporation_id}/members/`

Alternate route: `/legacy/corporations/{corporation_id}/members/`

Alternate route: `/dev/corporations/{corporation_id}/members/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDMembers(params *GetCorporationsCorporationIDMembersParams, authInfo runtime.ClientAuthInfoWriter) (*GetCorporationsCorporationIDMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_members",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/members/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDMembersOK), nil

}

/*
GetCorporationsCorporationIDRoles gets corporation members

Return the roles of all members if the character has the personnel manager role or any grantable role.

---

Alternate route: `/v1/corporations/{corporation_id}/roles/`

Alternate route: `/legacy/corporations/{corporation_id}/roles/`

Alternate route: `/dev/corporations/{corporation_id}/roles/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDRoles(params *GetCorporationsCorporationIDRolesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCorporationsCorporationIDRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_roles",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDRolesOK), nil

}

/*
GetCorporationsNames gets corporation names

Resolve a set of corporation IDs to corporation names

---

Alternate route: `/v1/corporations/names/`

Alternate route: `/legacy/corporations/names/`

Alternate route: `/dev/corporations/names/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsNames(params *GetCorporationsNamesParams) (*GetCorporationsNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_names",
		Method:             "GET",
		PathPattern:        "/corporations/names/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsNamesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsNamesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
