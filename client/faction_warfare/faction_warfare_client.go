// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new faction warfare API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for faction warfare API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFwLeaderboards lists of the top factions in faction warfare

Top 4 leaderboard of factions for kills and victory points separated by total, last week and yesterday.

---
Alternate route: `/v1/fw/leaderboards/`

Alternate route: `/legacy/fw/leaderboards/`

Alternate route: `/dev/fw/leaderboards/`

---
This route expires daily at 11:05
*/
func (a *Client) GetFwLeaderboards(params *GetFwLeaderboardsParams) (*GetFwLeaderboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwLeaderboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_fw_leaderboards",
		Method:             "GET",
		PathPattern:        "/fw/leaderboards/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwLeaderboardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFwLeaderboardsOK), nil

}

/*
GetFwLeaderboardsCharacters lists of the top pilots in faction warfare

Top 100 leaderboard of pilots for kills and victory points separated by total, last week and yesterday.

---
Alternate route: `/v1/fw/leaderboards/characters/`

Alternate route: `/legacy/fw/leaderboards/characters/`

Alternate route: `/dev/fw/leaderboards/characters/`

---
This route expires daily at 11:05
*/
func (a *Client) GetFwLeaderboardsCharacters(params *GetFwLeaderboardsCharactersParams) (*GetFwLeaderboardsCharactersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwLeaderboardsCharactersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_fw_leaderboards_characters",
		Method:             "GET",
		PathPattern:        "/fw/leaderboards/characters/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwLeaderboardsCharactersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFwLeaderboardsCharactersOK), nil

}

/*
GetFwLeaderboardsCorporations lists of the top corporations in faction warfare

Top 10 leaderboard of corporations for kills and victory points separated by total, last week and yesterday.

---
Alternate route: `/v1/fw/leaderboards/corporations/`

Alternate route: `/legacy/fw/leaderboards/corporations/`

Alternate route: `/dev/fw/leaderboards/corporations/`

---
This route expires daily at 11:05
*/
func (a *Client) GetFwLeaderboardsCorporations(params *GetFwLeaderboardsCorporationsParams) (*GetFwLeaderboardsCorporationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwLeaderboardsCorporationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_fw_leaderboards_corporations",
		Method:             "GET",
		PathPattern:        "/fw/leaderboards/corporations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwLeaderboardsCorporationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFwLeaderboardsCorporationsOK), nil

}

/*
GetFwStats ans overview of statistics about factions involved in faction warfare

Statistical overviews of factions involved in faction warfare

---
Alternate route: `/v1/fw/stats/`

Alternate route: `/legacy/fw/stats/`

Alternate route: `/dev/fw/stats/`

---
This route expires daily at 11:05
*/
func (a *Client) GetFwStats(params *GetFwStatsParams) (*GetFwStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwStatsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_fw_stats",
		Method:             "GET",
		PathPattern:        "/fw/stats/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFwStatsOK), nil

}

/*
GetFwSystems ownerships of faction warfare systems

An overview of the current ownership of faction warfare solar systems

---
Alternate route: `/v1/fw/systems/`

Alternate route: `/legacy/fw/systems/`

Alternate route: `/dev/fw/systems/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetFwSystems(params *GetFwSystemsParams) (*GetFwSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwSystemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_fw_systems",
		Method:             "GET",
		PathPattern:        "/fw/systems/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwSystemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFwSystemsOK), nil

}

/*
GetFwWars data about which n p c factions are at war

Data about which NPC factions are at war

---
Alternate route: `/v1/fw/wars/`

Alternate route: `/legacy/fw/wars/`

Alternate route: `/dev/fw/wars/`

---
This route expires daily at 11:05
*/
func (a *Client) GetFwWars(params *GetFwWarsParams) (*GetFwWarsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwWarsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_fw_wars",
		Method:             "GET",
		PathPattern:        "/fw/wars/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwWarsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFwWarsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}