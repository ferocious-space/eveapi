// Code generated by go-swagger; DO NOT EDIT.

package esi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/esi/alliance"
	"github.com/ferocious-space/eveapi/esi/assets"
	"github.com/ferocious-space/eveapi/esi/bookmarks"
	"github.com/ferocious-space/eveapi/esi/calendar"
	"github.com/ferocious-space/eveapi/esi/character"
	"github.com/ferocious-space/eveapi/esi/clones"
	"github.com/ferocious-space/eveapi/esi/contacts"
	"github.com/ferocious-space/eveapi/esi/contracts"
	"github.com/ferocious-space/eveapi/esi/corporation"
	"github.com/ferocious-space/eveapi/esi/dogma"
	"github.com/ferocious-space/eveapi/esi/faction_warfare"
	"github.com/ferocious-space/eveapi/esi/fittings"
	"github.com/ferocious-space/eveapi/esi/fleets"
	"github.com/ferocious-space/eveapi/esi/incursions"
	"github.com/ferocious-space/eveapi/esi/industry"
	"github.com/ferocious-space/eveapi/esi/insurance"
	"github.com/ferocious-space/eveapi/esi/killmails"
	"github.com/ferocious-space/eveapi/esi/location"
	"github.com/ferocious-space/eveapi/esi/loyalty"
	"github.com/ferocious-space/eveapi/esi/mail"
	"github.com/ferocious-space/eveapi/esi/market"
	"github.com/ferocious-space/eveapi/esi/meta"
	"github.com/ferocious-space/eveapi/esi/opportunities"
	"github.com/ferocious-space/eveapi/esi/planetary_interaction"
	"github.com/ferocious-space/eveapi/esi/routes"
	"github.com/ferocious-space/eveapi/esi/search"
	"github.com/ferocious-space/eveapi/esi/skills"
	"github.com/ferocious-space/eveapi/esi/sovereignty"
	"github.com/ferocious-space/eveapi/esi/status"
	"github.com/ferocious-space/eveapi/esi/universe"
	"github.com/ferocious-space/eveapi/esi/user_interface"
	"github.com/ferocious-space/eveapi/esi/wallet"
	"github.com/ferocious-space/eveapi/esi/wars"
)

// Default e v e swagger interface HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "esi.evetech.net"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
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
	cli.Meta = meta.New(transport, formats)
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
	Alliance alliance.ClientService

	Assets assets.ClientService

	Bookmarks bookmarks.ClientService

	Calendar calendar.ClientService

	Character character.ClientService

	Clones clones.ClientService

	Contacts contacts.ClientService

	Contracts contracts.ClientService

	Corporation corporation.ClientService

	Dogma dogma.ClientService

	FactionWarfare faction_warfare.ClientService

	Fittings fittings.ClientService

	Fleets fleets.ClientService

	Incursions incursions.ClientService

	Industry industry.ClientService

	Insurance insurance.ClientService

	Killmails killmails.ClientService

	Location location.ClientService

	Loyalty loyalty.ClientService

	Mail mail.ClientService

	Market market.ClientService

	Meta meta.ClientService

	Opportunities opportunities.ClientService

	PlanetaryInteraction planetary_interaction.ClientService

	Routes routes.ClientService

	Search search.ClientService

	Skills skills.ClientService

	Sovereignty sovereignty.ClientService

	Status status.ClientService

	Universe universe.ClientService

	UserInterface user_interface.ClientService

	Wallet wallet.ClientService

	Wars wars.ClientService

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
	c.Meta.SetTransport(transport)
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
