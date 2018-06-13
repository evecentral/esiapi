// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/evecentral/esiapi/client/alliance"
	"github.com/evecentral/esiapi/client/assets"
	"github.com/evecentral/esiapi/client/bookmarks"
	"github.com/evecentral/esiapi/client/calendar"
	"github.com/evecentral/esiapi/client/character"
	"github.com/evecentral/esiapi/client/clones"
	"github.com/evecentral/esiapi/client/contacts"
	"github.com/evecentral/esiapi/client/contracts"
	"github.com/evecentral/esiapi/client/corporation"
	"github.com/evecentral/esiapi/client/dogma"
	"github.com/evecentral/esiapi/client/faction_warfare"
	"github.com/evecentral/esiapi/client/fittings"
	"github.com/evecentral/esiapi/client/fleets"
	"github.com/evecentral/esiapi/client/incursions"
	"github.com/evecentral/esiapi/client/industry"
	"github.com/evecentral/esiapi/client/insurance"
	"github.com/evecentral/esiapi/client/killmails"
	"github.com/evecentral/esiapi/client/location"
	"github.com/evecentral/esiapi/client/loyalty"
	"github.com/evecentral/esiapi/client/mail"
	"github.com/evecentral/esiapi/client/market"
	"github.com/evecentral/esiapi/client/opportunities"
	"github.com/evecentral/esiapi/client/planetary_interaction"
	"github.com/evecentral/esiapi/client/routes"
	"github.com/evecentral/esiapi/client/search"
	"github.com/evecentral/esiapi/client/skills"
	"github.com/evecentral/esiapi/client/sovereignty"
	"github.com/evecentral/esiapi/client/status"
	"github.com/evecentral/esiapi/client/universe"
	"github.com/evecentral/esiapi/client/user_interface"
	"github.com/evecentral/esiapi/client/wallet"
	"github.com/evecentral/esiapi/client/wars"
)

// Default e v e swagger interface HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "esi.evetech.net"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/latest"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new e v e swagger interface HTTP client.
func NewHTTPClient(formats strfmt.Registry) *EVESwaggerInterface {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new e v e swagger interface HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *EVESwaggerInterface {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new e v e swagger interface client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *EVESwaggerInterface {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(EVESwaggerInterface)
	cli.Transport = transport

	cli.Alliance = alliance.New(transport, formats)

	cli.Assets = assets.New(transport, formats)

	cli.Bookmarks = bookmarks.New(transport, formats)

	cli.Calendar = calendar.New(transport, formats)

	cli.Character = character.New(transport, formats)

	cli.Clones = clones.New(transport, formats)

	cli.Contacts = contacts.New(transport, formats)

	cli.Contracts = contracts.New(transport, formats)

	cli.Corporation = corporation.New(transport, formats)

	cli.Dogma = dogma.New(transport, formats)

	cli.FactionWarfare = faction_warfare.New(transport, formats)

	cli.Fittings = fittings.New(transport, formats)

	cli.Fleets = fleets.New(transport, formats)

	cli.Incursions = incursions.New(transport, formats)

	cli.Industry = industry.New(transport, formats)

	cli.Insurance = insurance.New(transport, formats)

	cli.Killmails = killmails.New(transport, formats)

	cli.Location = location.New(transport, formats)

	cli.Loyalty = loyalty.New(transport, formats)

	cli.Mail = mail.New(transport, formats)

	cli.Market = market.New(transport, formats)

	cli.Opportunities = opportunities.New(transport, formats)

	cli.PlanetaryInteraction = planetary_interaction.New(transport, formats)

	cli.Routes = routes.New(transport, formats)

	cli.Search = search.New(transport, formats)

	cli.Skills = skills.New(transport, formats)

	cli.Sovereignty = sovereignty.New(transport, formats)

	cli.Status = status.New(transport, formats)

	cli.Universe = universe.New(transport, formats)

	cli.UserInterface = user_interface.New(transport, formats)

	cli.Wallet = wallet.New(transport, formats)

	cli.Wars = wars.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// EVESwaggerInterface is a client for e v e swagger interface
type EVESwaggerInterface struct {
	Alliance *alliance.Client

	Assets *assets.Client

	Bookmarks *bookmarks.Client

	Calendar *calendar.Client

	Character *character.Client

	Clones *clones.Client

	Contacts *contacts.Client

	Contracts *contracts.Client

	Corporation *corporation.Client

	Dogma *dogma.Client

	FactionWarfare *faction_warfare.Client

	Fittings *fittings.Client

	Fleets *fleets.Client

	Incursions *incursions.Client

	Industry *industry.Client

	Insurance *insurance.Client

	Killmails *killmails.Client

	Location *location.Client

	Loyalty *loyalty.Client

	Mail *mail.Client

	Market *market.Client

	Opportunities *opportunities.Client

	PlanetaryInteraction *planetary_interaction.Client

	Routes *routes.Client

	Search *search.Client

	Skills *skills.Client

	Sovereignty *sovereignty.Client

	Status *status.Client

	Universe *universe.Client

	UserInterface *user_interface.Client

	Wallet *wallet.Client

	Wars *wars.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *EVESwaggerInterface) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Alliance.SetTransport(transport)

	c.Assets.SetTransport(transport)

	c.Bookmarks.SetTransport(transport)

	c.Calendar.SetTransport(transport)

	c.Character.SetTransport(transport)

	c.Clones.SetTransport(transport)

	c.Contacts.SetTransport(transport)

	c.Contracts.SetTransport(transport)

	c.Corporation.SetTransport(transport)

	c.Dogma.SetTransport(transport)

	c.FactionWarfare.SetTransport(transport)

	c.Fittings.SetTransport(transport)

	c.Fleets.SetTransport(transport)

	c.Incursions.SetTransport(transport)

	c.Industry.SetTransport(transport)

	c.Insurance.SetTransport(transport)

	c.Killmails.SetTransport(transport)

	c.Location.SetTransport(transport)

	c.Loyalty.SetTransport(transport)

	c.Mail.SetTransport(transport)

	c.Market.SetTransport(transport)

	c.Opportunities.SetTransport(transport)

	c.PlanetaryInteraction.SetTransport(transport)

	c.Routes.SetTransport(transport)

	c.Search.SetTransport(transport)

	c.Skills.SetTransport(transport)

	c.Sovereignty.SetTransport(transport)

	c.Status.SetTransport(transport)

	c.Universe.SetTransport(transport)

	c.UserInterface.SetTransport(transport)

	c.Wallet.SetTransport(transport)

	c.Wars.SetTransport(transport)

}
