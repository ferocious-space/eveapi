// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new character API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for character API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterID(params *GetCharactersCharacterIDParams, opts ...ClientOption) (*GetCharactersCharacterIDOK, error)

	GetCharactersCharacterIDAgentsResearch(params *GetCharactersCharacterIDAgentsResearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDAgentsResearchOK, error)

	GetCharactersCharacterIDBlueprints(params *GetCharactersCharacterIDBlueprintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDBlueprintsOK, error)

	GetCharactersCharacterIDCorporationhistory(params *GetCharactersCharacterIDCorporationhistoryParams, opts ...ClientOption) (*GetCharactersCharacterIDCorporationhistoryOK, error)

	GetCharactersCharacterIDFatigue(params *GetCharactersCharacterIDFatigueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDFatigueOK, error)

	GetCharactersCharacterIDMedals(params *GetCharactersCharacterIDMedalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDMedalsOK, error)

	GetCharactersCharacterIDNotifications(params *GetCharactersCharacterIDNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDNotificationsOK, error)

	GetCharactersCharacterIDNotificationsContacts(params *GetCharactersCharacterIDNotificationsContactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDNotificationsContactsOK, error)

	GetCharactersCharacterIDPortrait(params *GetCharactersCharacterIDPortraitParams, opts ...ClientOption) (*GetCharactersCharacterIDPortraitOK, error)

	GetCharactersCharacterIDRoles(params *GetCharactersCharacterIDRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDRolesOK, error)

	GetCharactersCharacterIDStandings(params *GetCharactersCharacterIDStandingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDStandingsOK, error)

	GetCharactersCharacterIDTitles(params *GetCharactersCharacterIDTitlesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDTitlesOK, error)

	PostCharactersAffiliation(params *PostCharactersAffiliationParams, opts ...ClientOption) (*PostCharactersAffiliationOK, error)

	PostCharactersCharacterIDCspa(params *PostCharactersCharacterIDCspaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCharactersCharacterIDCspaCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterID gets character s public information

  Public information about a character

---

This route is cached for up to 86400 seconds
*/
func (a *Client) GetCharactersCharacterID(params *GetCharactersCharacterIDParams, opts ...ClientOption) (*GetCharactersCharacterIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id",
		Method:             "GET",
		PathPattern:        "/v4/characters/{character_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDAgentsResearch gets agents research

  Return a list of agents research information for a character. The formula for finding the current research points with an agent is: currentPoints = remainderPoints + pointsPerDay * days(currentTime - researchStartDate)

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDAgentsResearch(params *GetCharactersCharacterIDAgentsResearchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDAgentsResearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDAgentsResearchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_agents_research",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/agents_research/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDAgentsResearchReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDAgentsResearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_agents_research: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDBlueprints gets blueprints

  Return a list of blueprints the character owns

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDBlueprints(params *GetCharactersCharacterIDBlueprintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDBlueprintsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDBlueprintsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_blueprints",
		Method:             "GET",
		PathPattern:        "/v2/characters/{character_id}/blueprints/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDBlueprintsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDBlueprintsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_blueprints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDCorporationhistory gets corporation history

  Get a list of all the corporations a character has been a member of

---

This route is cached for up to 86400 seconds
*/
func (a *Client) GetCharactersCharacterIDCorporationhistory(params *GetCharactersCharacterIDCorporationhistoryParams, opts ...ClientOption) (*GetCharactersCharacterIDCorporationhistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDCorporationhistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_corporationhistory",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/corporationhistory/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDCorporationhistoryReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDCorporationhistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_corporationhistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDFatigue gets jump fatigue

  Return a character's jump activation and fatigue information

---

This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDFatigue(params *GetCharactersCharacterIDFatigueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDFatigueOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDFatigueParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_fatigue",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/fatigue/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDFatigueReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDFatigueOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_fatigue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDMedals gets medals

  Return a list of medals the character has

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDMedals(params *GetCharactersCharacterIDMedalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDMedalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDMedalsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_medals",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/medals/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDMedalsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDMedalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_medals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDNotifications gets character notifications

  Return character notifications

---

This route is cached for up to 600 seconds
*/
func (a *Client) GetCharactersCharacterIDNotifications(params *GetCharactersCharacterIDNotificationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDNotificationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_notifications",
		Method:             "GET",
		PathPattern:        "/v5/characters/{character_id}/notifications/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDNotificationsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_notifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDNotificationsContacts gets new contact notifications

  Return notifications about having been added to someone's contact list

---

This route is cached for up to 600 seconds
*/
func (a *Client) GetCharactersCharacterIDNotificationsContacts(params *GetCharactersCharacterIDNotificationsContactsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDNotificationsContactsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDNotificationsContactsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_notifications_contacts",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/notifications/contacts/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDNotificationsContactsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDNotificationsContactsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_notifications_contacts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDPortrait gets character portraits

  Get portrait urls for a character

---

This route expires daily at 11:05
*/
func (a *Client) GetCharactersCharacterIDPortrait(params *GetCharactersCharacterIDPortraitParams, opts ...ClientOption) (*GetCharactersCharacterIDPortraitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDPortraitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_portrait",
		Method:             "GET",
		PathPattern:        "/v2/characters/{character_id}/portrait/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDPortraitReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDPortraitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_portrait: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDRoles gets character corporation roles

  Returns a character's corporation roles

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDRoles(params *GetCharactersCharacterIDRolesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_roles",
		Method:             "GET",
		PathPattern:        "/v2/characters/{character_id}/roles/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDRolesReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_roles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDStandings gets standings

  Return character standings from agents, NPC corporations, and factions

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDStandings(params *GetCharactersCharacterIDStandingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDStandingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDStandingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_standings",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/standings/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDStandingsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDStandingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_standings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDTitles gets character corporation titles

  Returns a character's titles

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDTitles(params *GetCharactersCharacterIDTitlesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDTitlesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDTitlesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_titles",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/titles/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDTitlesReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDTitlesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_titles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCharactersAffiliation characters affiliation

  Bulk lookup of character IDs to corporation, alliance and faction

---

This route is cached for up to 3600 seconds

---
[Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#POST-/characters/affiliation/)
*/
func (a *Client) PostCharactersAffiliation(params *PostCharactersAffiliationParams, opts ...ClientOption) (*PostCharactersAffiliationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharactersAffiliationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_characters_affiliation",
		Method:             "POST",
		PathPattern:        "/v1/characters/affiliation/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCharactersAffiliationReader{formats: a.formats},
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
	success, ok := result.(*PostCharactersAffiliationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_characters_affiliation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCharactersCharacterIDCspa calculates a c s p a charge cost

  Takes a source character ID in the url and a set of target character ID's in the body, returns a CSPA charge cost

---
*/
func (a *Client) PostCharactersCharacterIDCspa(params *PostCharactersCharacterIDCspaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCharactersCharacterIDCspaCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharactersCharacterIDCspaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_characters_character_id_cspa",
		Method:             "POST",
		PathPattern:        "/v4/characters/{character_id}/cspa/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCharactersCharacterIDCspaReader{formats: a.formats},
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
	success, ok := result.(*PostCharactersCharacterIDCspaCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_characters_character_id_cspa: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
