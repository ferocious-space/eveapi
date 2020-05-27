/*
 *    Copyright 2021 FerociousBite and Contributors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// Code generated by go-swagger; DO NOT EDIT.

package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new industry API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for industry API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDIndustryJobs(params *GetCharactersCharacterIDIndustryJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDIndustryJobsOK, error)

	GetCharactersCharacterIDMining(params *GetCharactersCharacterIDMiningParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDMiningOK, error)

	GetCorporationCorporationIDMiningExtractions(params *GetCorporationCorporationIDMiningExtractionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationCorporationIDMiningExtractionsOK, error)

	GetCorporationCorporationIDMiningObservers(params *GetCorporationCorporationIDMiningObserversParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationCorporationIDMiningObserversOK, error)

	GetCorporationCorporationIDMiningObserversObserverID(params *GetCorporationCorporationIDMiningObserversObserverIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationCorporationIDMiningObserversObserverIDOK, error)

	GetCorporationsCorporationIDIndustryJobs(params *GetCorporationsCorporationIDIndustryJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDIndustryJobsOK, error)

	GetIndustryFacilities(params *GetIndustryFacilitiesParams, opts ...ClientOption) (*GetIndustryFacilitiesOK, error)

	GetIndustrySystems(params *GetIndustrySystemsParams, opts ...ClientOption) (*GetIndustrySystemsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDIndustryJobs lists character industry jobs

  List industry jobs placed by a character

---

This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDIndustryJobs(params *GetCharactersCharacterIDIndustryJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDIndustryJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDIndustryJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_industry_jobs",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/industry/jobs/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDIndustryJobsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDIndustryJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_industry_jobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDMining characters mining ledger

  Paginated record of all mining done by a character for the past 30 days

---

This route is cached for up to 600 seconds
*/
func (a *Client) GetCharactersCharacterIDMining(params *GetCharactersCharacterIDMiningParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDMiningOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDMiningParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_mining",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/mining/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDMiningReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDMiningOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_mining: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationCorporationIDMiningExtractions moons extraction timers

  Extraction timers for all moon chunks being extracted by refineries belonging to a corporation.

---

This route is cached for up to 1800 seconds

---
Requires one of the following EVE corporation role(s): Station_Manager
*/
func (a *Client) GetCorporationCorporationIDMiningExtractions(params *GetCorporationCorporationIDMiningExtractionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationCorporationIDMiningExtractionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationCorporationIDMiningExtractionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporation_corporation_id_mining_extractions",
		Method:             "GET",
		PathPattern:        "/v1/corporation/{corporation_id}/mining/extractions/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationCorporationIDMiningExtractionsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationCorporationIDMiningExtractionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporation_corporation_id_mining_extractions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationCorporationIDMiningObservers corporations mining observers

  Paginated list of all entities capable of observing and recording mining for a corporation

---

This route is cached for up to 3600 seconds

---
Requires one of the following EVE corporation role(s): Accountant
*/
func (a *Client) GetCorporationCorporationIDMiningObservers(params *GetCorporationCorporationIDMiningObserversParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationCorporationIDMiningObserversOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationCorporationIDMiningObserversParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporation_corporation_id_mining_observers",
		Method:             "GET",
		PathPattern:        "/v1/corporation/{corporation_id}/mining/observers/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationCorporationIDMiningObserversReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationCorporationIDMiningObserversOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporation_corporation_id_mining_observers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationCorporationIDMiningObserversObserverID observeds corporation mining

  Paginated record of all mining seen by an observer

---

This route is cached for up to 3600 seconds

---
Requires one of the following EVE corporation role(s): Accountant
*/
func (a *Client) GetCorporationCorporationIDMiningObserversObserverID(params *GetCorporationCorporationIDMiningObserversObserverIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationCorporationIDMiningObserversObserverIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationCorporationIDMiningObserversObserverIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporation_corporation_id_mining_observers_observer_id",
		Method:             "GET",
		PathPattern:        "/v1/corporation/{corporation_id}/mining/observers/{observer_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationCorporationIDMiningObserversObserverIDReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationCorporationIDMiningObserversObserverIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporation_corporation_id_mining_observers_observer_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDIndustryJobs lists corporation industry jobs

  List industry jobs run by a corporation

---

This route is cached for up to 300 seconds

---
Requires one of the following EVE corporation role(s): Factory_Manager
*/
func (a *Client) GetCorporationsCorporationIDIndustryJobs(params *GetCorporationsCorporationIDIndustryJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDIndustryJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDIndustryJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_industry_jobs",
		Method:             "GET",
		PathPattern:        "/v1/corporations/{corporation_id}/industry/jobs/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDIndustryJobsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDIndustryJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_industry_jobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIndustryFacilities lists industry facilities

  Return a list of industry facilities

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetIndustryFacilities(params *GetIndustryFacilitiesParams, opts ...ClientOption) (*GetIndustryFacilitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIndustryFacilitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_industry_facilities",
		Method:             "GET",
		PathPattern:        "/v1/industry/facilities/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIndustryFacilitiesReader{formats: a.formats},
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
	success, ok := result.(*GetIndustryFacilitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_industry_facilities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIndustrySystems lists solar system cost indices

  Return cost indices for solar systems

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetIndustrySystems(params *GetIndustrySystemsParams, opts ...ClientOption) (*GetIndustrySystemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIndustrySystemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_industry_systems",
		Method:             "GET",
		PathPattern:        "/v1/industry/systems/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIndustrySystemsReader{formats: a.formats},
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
	success, ok := result.(*GetIndustrySystemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_industry_systems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
