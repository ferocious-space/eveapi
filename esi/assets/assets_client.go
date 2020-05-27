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

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new assets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for assets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDAssets(params *GetCharactersCharacterIDAssetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDAssetsOK, error)

	GetCorporationsCorporationIDAssets(params *GetCorporationsCorporationIDAssetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDAssetsOK, error)

	PostCharactersCharacterIDAssetsLocations(params *PostCharactersCharacterIDAssetsLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCharactersCharacterIDAssetsLocationsOK, error)

	PostCharactersCharacterIDAssetsNames(params *PostCharactersCharacterIDAssetsNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCharactersCharacterIDAssetsNamesOK, error)

	PostCorporationsCorporationIDAssetsLocations(params *PostCorporationsCorporationIDAssetsLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCorporationsCorporationIDAssetsLocationsOK, error)

	PostCorporationsCorporationIDAssetsNames(params *PostCorporationsCorporationIDAssetsNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCorporationsCorporationIDAssetsNamesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDAssets gets character assets

  Return a list of the characters assets

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDAssets(params *GetCharactersCharacterIDAssetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDAssetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDAssetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_assets",
		Method:             "GET",
		PathPattern:        "/v5/characters/{character_id}/assets/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDAssetsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDAssetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_assets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCorporationsCorporationIDAssets gets corporation assets

  Return a list of the corporation assets

---

This route is cached for up to 3600 seconds

---
Requires one of the following EVE corporation role(s): Director
*/
func (a *Client) GetCorporationsCorporationIDAssets(params *GetCorporationsCorporationIDAssetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDAssetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDAssetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_assets",
		Method:             "GET",
		PathPattern:        "/v5/corporations/{corporation_id}/assets/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDAssetsReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDAssetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_assets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCharactersCharacterIDAssetsLocations gets character asset locations

  Return locations for a set of item ids, which you can get from character assets endpoint. Coordinates for items in hangars or stations are set to (0,0,0)

---

*/
func (a *Client) PostCharactersCharacterIDAssetsLocations(params *PostCharactersCharacterIDAssetsLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCharactersCharacterIDAssetsLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharactersCharacterIDAssetsLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_characters_character_id_assets_locations",
		Method:             "POST",
		PathPattern:        "/v2/characters/{character_id}/assets/locations/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCharactersCharacterIDAssetsLocationsReader{formats: a.formats},
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
	success, ok := result.(*PostCharactersCharacterIDAssetsLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_characters_character_id_assets_locations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCharactersCharacterIDAssetsNames gets character asset names

  Return names for a set of item ids, which you can get from character assets endpoint. Typically used for items that can customize names, like containers or ships.

---

*/
func (a *Client) PostCharactersCharacterIDAssetsNames(params *PostCharactersCharacterIDAssetsNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCharactersCharacterIDAssetsNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCharactersCharacterIDAssetsNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_characters_character_id_assets_names",
		Method:             "POST",
		PathPattern:        "/v1/characters/{character_id}/assets/names/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCharactersCharacterIDAssetsNamesReader{formats: a.formats},
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
	success, ok := result.(*PostCharactersCharacterIDAssetsNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_characters_character_id_assets_names: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCorporationsCorporationIDAssetsLocations gets corporation asset locations

  Return locations for a set of item ids, which you can get from corporation assets endpoint. Coordinates for items in hangars or stations are set to (0,0,0)

---

Requires one of the following EVE corporation role(s): Director
*/
func (a *Client) PostCorporationsCorporationIDAssetsLocations(params *PostCorporationsCorporationIDAssetsLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCorporationsCorporationIDAssetsLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCorporationsCorporationIDAssetsLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_corporations_corporation_id_assets_locations",
		Method:             "POST",
		PathPattern:        "/v2/corporations/{corporation_id}/assets/locations/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCorporationsCorporationIDAssetsLocationsReader{formats: a.formats},
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
	success, ok := result.(*PostCorporationsCorporationIDAssetsLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_corporations_corporation_id_assets_locations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostCorporationsCorporationIDAssetsNames gets corporation asset names

  Return names for a set of item ids, which you can get from corporation assets endpoint. Only valid for items that can customize names, like containers or ships

---

Requires one of the following EVE corporation role(s): Director
*/
func (a *Client) PostCorporationsCorporationIDAssetsNames(params *PostCorporationsCorporationIDAssetsNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostCorporationsCorporationIDAssetsNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCorporationsCorporationIDAssetsNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_corporations_corporation_id_assets_names",
		Method:             "POST",
		PathPattern:        "/v1/corporations/{corporation_id}/assets/names/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCorporationsCorporationIDAssetsNamesReader{formats: a.formats},
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
	success, ok := result.(*PostCorporationsCorporationIDAssetsNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_corporations_corporation_id_assets_names: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
