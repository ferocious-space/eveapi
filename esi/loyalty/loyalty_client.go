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

package loyalty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new loyalty API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for loyalty API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDLoyaltyPoints(params *GetCharactersCharacterIDLoyaltyPointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDLoyaltyPointsOK, error)

	GetLoyaltyStoresCorporationIDOffers(params *GetLoyaltyStoresCorporationIDOffersParams, opts ...ClientOption) (*GetLoyaltyStoresCorporationIDOffersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCharactersCharacterIDLoyaltyPoints gets loyalty points

  Return a list of loyalty points for all corporations the character has worked for

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDLoyaltyPoints(params *GetCharactersCharacterIDLoyaltyPointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDLoyaltyPointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDLoyaltyPointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_loyalty_points",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/loyalty/points/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDLoyaltyPointsReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDLoyaltyPointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_loyalty_points: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLoyaltyStoresCorporationIDOffers lists loyalty store offers

  Return a list of offers from a specific corporation's loyalty store

---

This route expires daily at 11:05
*/
func (a *Client) GetLoyaltyStoresCorporationIDOffers(params *GetLoyaltyStoresCorporationIDOffersParams, opts ...ClientOption) (*GetLoyaltyStoresCorporationIDOffersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoyaltyStoresCorporationIDOffersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_loyalty_stores_corporation_id_offers",
		Method:             "GET",
		PathPattern:        "/v1/loyalty/stores/{corporation_id}/offers/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoyaltyStoresCorporationIDOffersReader{formats: a.formats},
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
	success, ok := result.(*GetLoyaltyStoresCorporationIDOffersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_loyalty_stores_corporation_id_offers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
