// Code generated by go-swagger; DO NOT EDIT.

package faction_warfare

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new faction warfare API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for faction warfare API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDFwStats(params *GetCharactersCharacterIDFwStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDFwStatsOK, error)

	GetCorporationsCorporationIDFwStats(params *GetCorporationsCorporationIDFwStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDFwStatsOK, error)

	GetFwLeaderboards(params *GetFwLeaderboardsParams, opts ...ClientOption) (*GetFwLeaderboardsOK, error)

	GetFwLeaderboardsCharacters(params *GetFwLeaderboardsCharactersParams, opts ...ClientOption) (*GetFwLeaderboardsCharactersOK, error)

	GetFwLeaderboardsCorporations(params *GetFwLeaderboardsCorporationsParams, opts ...ClientOption) (*GetFwLeaderboardsCorporationsOK, error)

	GetFwStats(params *GetFwStatsParams, opts ...ClientOption) (*GetFwStatsOK, error)

	GetFwSystems(params *GetFwSystemsParams, opts ...ClientOption) (*GetFwSystemsOK, error)

	GetFwWars(params *GetFwWarsParams, opts ...ClientOption) (*GetFwWarsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDFwStats overviews of a character involved in faction warfare

  Statistical overview of a character involved in faction warfare

---

This route expires daily at 11:05
*/
func (a *Client) GetCharactersCharacterIDFwStats(params *GetCharactersCharacterIDFwStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDFwStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDFwStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_fw_stats",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/fw/stats/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDFwStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCharactersCharacterIDFwStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_fw_stats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDFwStats overviews of a corporation involved in faction warfare

  Statistics about a corporation involved in faction warfare

---

This route expires daily at 11:05
*/
func (a *Client) GetCorporationsCorporationIDFwStats(params *GetCorporationsCorporationIDFwStatsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDFwStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDFwStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_fw_stats",
		Method:             "GET",
		PathPattern:        "/v1/corporations/{corporation_id}/fw/stats/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDFwStatsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCorporationsCorporationIDFwStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_fw_stats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFwLeaderboards lists of the top factions in faction warfare

  Top 4 leaderboard of factions for kills and victory points separated by total, last week and yesterday

---

This route expires daily at 11:05
*/
func (a *Client) GetFwLeaderboards(params *GetFwLeaderboardsParams, opts ...ClientOption) (*GetFwLeaderboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwLeaderboardsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fw_leaderboards",
		Method:             "GET",
		PathPattern:        "/v1/fw/leaderboards/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwLeaderboardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFwLeaderboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fw_leaderboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFwLeaderboardsCharacters lists of the top pilots in faction warfare

  Top 100 leaderboard of pilots for kills and victory points separated by total, last week and yesterday

---

This route expires daily at 11:05
*/
func (a *Client) GetFwLeaderboardsCharacters(params *GetFwLeaderboardsCharactersParams, opts ...ClientOption) (*GetFwLeaderboardsCharactersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwLeaderboardsCharactersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fw_leaderboards_characters",
		Method:             "GET",
		PathPattern:        "/v1/fw/leaderboards/characters/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwLeaderboardsCharactersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFwLeaderboardsCharactersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fw_leaderboards_characters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFwLeaderboardsCorporations lists of the top corporations in faction warfare

  Top 10 leaderboard of corporations for kills and victory points separated by total, last week and yesterday

---

This route expires daily at 11:05
*/
func (a *Client) GetFwLeaderboardsCorporations(params *GetFwLeaderboardsCorporationsParams, opts ...ClientOption) (*GetFwLeaderboardsCorporationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwLeaderboardsCorporationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fw_leaderboards_corporations",
		Method:             "GET",
		PathPattern:        "/v1/fw/leaderboards/corporations/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwLeaderboardsCorporationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFwLeaderboardsCorporationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fw_leaderboards_corporations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFwStats ans overview of statistics about factions involved in faction warfare

  Statistical overviews of factions involved in faction warfare

---

This route expires daily at 11:05
*/
func (a *Client) GetFwStats(params *GetFwStatsParams, opts ...ClientOption) (*GetFwStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwStatsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fw_stats",
		Method:             "GET",
		PathPattern:        "/v1/fw/stats/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFwStatsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fw_stats: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFwSystems ownerships of faction warfare systems

  An overview of the current ownership of faction warfare solar systems

---

This route is cached for up to 1800 seconds
*/
func (a *Client) GetFwSystems(params *GetFwSystemsParams, opts ...ClientOption) (*GetFwSystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwSystemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fw_systems",
		Method:             "GET",
		PathPattern:        "/v2/fw/systems/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwSystemsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFwSystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fw_systems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFwWars data about which n p c factions are at war

  Data about which NPC factions are at war

---

This route expires daily at 11:05
*/
func (a *Client) GetFwWars(params *GetFwWarsParams, opts ...ClientOption) (*GetFwWarsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFwWarsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fw_wars",
		Method:             "GET",
		PathPattern:        "/v1/fw/wars/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFwWarsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFwWarsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fw_wars: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
