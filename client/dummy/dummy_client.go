package dummy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new dummy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dummy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharactersCharacterIDWalletsJournal gets character wallet journal

Returns the most recent 50 entries for the characters wallet journal. Optionally, takes an argument with a reference ID, and returns the prior 50 entries from the journal.

---

Alternate route: `/v1/characters/{character_id}/wallets/journal/`

Alternate route: `/legacy/characters/{character_id}/wallets/journal/`

Alternate route: `/dev/characters/{character_id}/wallets/journal/`

*/
func (a *Client) GetCharactersCharacterIDWalletsJournal(params *GetCharactersCharacterIDWalletsJournalParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDWalletsJournalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDWalletsJournalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_wallets_journal",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/wallets/journal/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDWalletsJournalReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDWalletsJournalOK), nil

}

/*
GetCharactersCharacterIDWalletsTransactions gets wallet transactions

Gets the 50 most recent transactions in a characters wallet. Optionally, takes an argument with a transaction ID, and returns the prior 50 transactions

---

Alternate route: `/v1/characters/{character_id}/wallets/transactions/`

Alternate route: `/legacy/characters/{character_id}/wallets/transactions/`

Alternate route: `/dev/characters/{character_id}/wallets/transactions/`

*/
func (a *Client) GetCharactersCharacterIDWalletsTransactions(params *GetCharactersCharacterIDWalletsTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDWalletsTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDWalletsTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_wallets_transactions",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/wallets/transactions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDWalletsTransactionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDWalletsTransactionsOK), nil

}

/*
GetCorporationsCorporationIDBookmarks dummies endpoint please ignore

Dummy

---

Alternate route: `/v1/corporations/{corporation_id}/bookmarks/`

Alternate route: `/legacy/corporations/{corporation_id}/bookmarks/`

Alternate route: `/dev/corporations/{corporation_id}/bookmarks/`

*/
func (a *Client) GetCorporationsCorporationIDBookmarks(params *GetCorporationsCorporationIDBookmarksParams) (*GetCorporationsCorporationIDBookmarksNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDBookmarksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_bookmarks",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/bookmarks/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDBookmarksReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDBookmarksNoContent), nil

}

/*
GetCorporationsCorporationIDBookmarksFolders dummies endpoint please ignore

Dummy

---

Alternate route: `/v1/corporations/{corporation_id}/bookmarks/folders/`

Alternate route: `/legacy/corporations/{corporation_id}/bookmarks/folders/`

Alternate route: `/dev/corporations/{corporation_id}/bookmarks/folders/`

*/
func (a *Client) GetCorporationsCorporationIDBookmarksFolders(params *GetCorporationsCorporationIDBookmarksFoldersParams) (*GetCorporationsCorporationIDBookmarksFoldersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDBookmarksFoldersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_bookmarks_folders",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/bookmarks/folders/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDBookmarksFoldersReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDBookmarksFoldersNoContent), nil

}

/*
GetCorporationsCorporationIDWallets dummies endpoint please ignore

Dummy

---

Alternate route: `/v1/corporations/{corporation_id}/wallets/`

Alternate route: `/legacy/corporations/{corporation_id}/wallets/`

Alternate route: `/dev/corporations/{corporation_id}/wallets/`

*/
func (a *Client) GetCorporationsCorporationIDWallets(params *GetCorporationsCorporationIDWalletsParams) (*GetCorporationsCorporationIDWalletsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDWalletsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_wallets",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/wallets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDWalletsReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDWalletsNoContent), nil

}

/*
GetCorporationsCorporationIDWalletsWalletIDJournal dummies endpoint please ignore

Dummy

---

Alternate route: `/v1/corporations/{corporation_id}/wallets/{wallet_id}/journal/`

Alternate route: `/legacy/corporations/{corporation_id}/wallets/{wallet_id}/journal/`

Alternate route: `/dev/corporations/{corporation_id}/wallets/{wallet_id}/journal/`

*/
func (a *Client) GetCorporationsCorporationIDWalletsWalletIDJournal(params *GetCorporationsCorporationIDWalletsWalletIDJournalParams) (*GetCorporationsCorporationIDWalletsWalletIDJournalNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDWalletsWalletIDJournalParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_wallets_wallet_id_journal",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/wallets/{wallet_id}/journal/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDWalletsWalletIDJournalReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDWalletsWalletIDJournalNoContent), nil

}

/*
GetCorporationsCorporationIDWalletsWalletIDTransactions dummies endpoint please ignore

Dummy

---

Alternate route: `/v1/corporations/{corporation_id}/wallets/{wallet_id}/transactions/`

Alternate route: `/legacy/corporations/{corporation_id}/wallets/{wallet_id}/transactions/`

Alternate route: `/dev/corporations/{corporation_id}/wallets/{wallet_id}/transactions/`

*/
func (a *Client) GetCorporationsCorporationIDWalletsWalletIDTransactions(params *GetCorporationsCorporationIDWalletsWalletIDTransactionsParams) (*GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDWalletsWalletIDTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_wallets_wallet_id_transactions",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/wallets/{wallet_id}/transactions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDWalletsWalletIDTransactionsReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDWalletsWalletIDTransactionsNoContent), nil

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
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUniversePlanetsPlanetIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
