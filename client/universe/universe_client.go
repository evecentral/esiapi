package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new universe API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for universe API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetUniversePlanetsPlanetID gets planet information

Information on a planet

---

Alternate route: `/v1/universe/planets/{planet_id}/`

Alternate route: `/legacy/universe/planets/{planet_id}/`

Alternate route: `/dev/universe/planets/{planet_id}/`

*/
func (a *Client) GetUniversePlanetsPlanetID(params *GetUniversePlanetsPlanetIDParams) (*GetUniversePlanetsPlanetIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUniversePlanetsPlanetIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_universe_planets_planet_id",
		Method:             "GET",
		PathPattern:        "/universe/planets/{planet_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUniversePlanetsPlanetIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUniversePlanetsPlanetIDOK), nil

}

/*
GetUniverseStationsStationID gets station information

Public information on stations

---

Alternate route: `/v1/universe/stations/{station_id}/`

Alternate route: `/legacy/universe/stations/{station_id}/`

Alternate route: `/dev/universe/stations/{station_id}/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetUniverseStationsStationID(params *GetUniverseStationsStationIDParams) (*GetUniverseStationsStationIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUniverseStationsStationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_universe_stations_station_id",
		Method:             "GET",
		PathPattern:        "/universe/stations/{station_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUniverseStationsStationIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUniverseStationsStationIDOK), nil

}

/*
GetUniverseStructures lists all public structures

List all public structures

---

Alternate route: `/v1/universe/structures/`

Alternate route: `/legacy/universe/structures/`

Alternate route: `/dev/universe/structures/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetUniverseStructures(params *GetUniverseStructuresParams) (*GetUniverseStructuresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUniverseStructuresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_universe_structures",
		Method:             "GET",
		PathPattern:        "/universe/structures/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUniverseStructuresReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUniverseStructuresOK), nil

}

/*
GetUniverseStructuresStructureID gets structure information

Returns information on requested structure, if you are on the ACL. Otherwise, returns "Forbidden" for all inputs.

---

Alternate route: `/v1/universe/structures/{structure_id}/`

Alternate route: `/legacy/universe/structures/{structure_id}/`

Alternate route: `/dev/universe/structures/{structure_id}/`

*/
func (a *Client) GetUniverseStructuresStructureID(params *GetUniverseStructuresStructureIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetUniverseStructuresStructureIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUniverseStructuresStructureIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_universe_structures_structure_id",
		Method:             "GET",
		PathPattern:        "/universe/structures/{structure_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUniverseStructuresStructureIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUniverseStructuresStructureIDOK), nil

}

/*
GetUniverseSystemsSystemID gets solar system information

Information on solar systems

---

Alternate route: `/v1/universe/systems/{system_id}/`

Alternate route: `/legacy/universe/systems/{system_id}/`

Alternate route: `/dev/universe/systems/{system_id}/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetUniverseSystemsSystemID(params *GetUniverseSystemsSystemIDParams) (*GetUniverseSystemsSystemIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUniverseSystemsSystemIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_universe_systems_system_id",
		Method:             "GET",
		PathPattern:        "/universe/systems/{system_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUniverseSystemsSystemIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUniverseSystemsSystemIDOK), nil

}

/*
GetUniverseTypesTypeID gets type information

Get information on a type

---

Alternate route: `/v1/universe/types/{type_id}/`

Alternate route: `/legacy/universe/types/{type_id}/`

Alternate route: `/dev/universe/types/{type_id}/`


---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetUniverseTypesTypeID(params *GetUniverseTypesTypeIDParams) (*GetUniverseTypesTypeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUniverseTypesTypeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_universe_types_type_id",
		Method:             "GET",
		PathPattern:        "/universe/types/{type_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUniverseTypesTypeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUniverseTypesTypeIDOK), nil

}

/*
PostUniverseNames gets names and categories for a set of ID s

Resolve a set of IDs to names and categories. Supported ID's for resolving are: Characters, Corporations, Alliances, Stations, Solar Systems, Constellations, Regions, Types.

---

Alternate route: `/v1/universe/names/`

Alternate route: `/legacy/universe/names/`

*/
func (a *Client) PostUniverseNames(params *PostUniverseNamesParams) (*PostUniverseNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUniverseNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post_universe_names",
		Method:             "POST",
		PathPattern:        "/universe/names/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUniverseNamesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostUniverseNamesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
