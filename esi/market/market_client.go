// Code generated by go-swagger; DO NOT EDIT.

package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new market API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for market API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCharactersCharacterIDOrders(params *GetCharactersCharacterIDOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDOrdersOK, error)

	GetCharactersCharacterIDOrdersHistory(params *GetCharactersCharacterIDOrdersHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDOrdersHistoryOK, error)

	GetCorporationsCorporationIDOrders(params *GetCorporationsCorporationIDOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDOrdersOK, error)

	GetCorporationsCorporationIDOrdersHistory(params *GetCorporationsCorporationIDOrdersHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDOrdersHistoryOK, error)

	GetMarketsGroups(params *GetMarketsGroupsParams, opts ...ClientOption) (*GetMarketsGroupsOK, error)

	GetMarketsGroupsMarketGroupID(params *GetMarketsGroupsMarketGroupIDParams, opts ...ClientOption) (*GetMarketsGroupsMarketGroupIDOK, error)

	GetMarketsPrices(params *GetMarketsPricesParams, opts ...ClientOption) (*GetMarketsPricesOK, error)

	GetMarketsRegionIDHistory(params *GetMarketsRegionIDHistoryParams, opts ...ClientOption) (*GetMarketsRegionIDHistoryOK, error)

	GetMarketsRegionIDOrders(params *GetMarketsRegionIDOrdersParams, opts ...ClientOption) (*GetMarketsRegionIDOrdersOK, error)

	GetMarketsRegionIDTypes(params *GetMarketsRegionIDTypesParams, opts ...ClientOption) (*GetMarketsRegionIDTypesOK, error)

	GetMarketsStructuresStructureID(params *GetMarketsStructuresStructureIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMarketsStructuresStructureIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetCharactersCharacterIDOrders lists open orders from a character

	List open market orders placed by a character

---

This route is cached for up to 1200 seconds
*/
func (a *Client) GetCharactersCharacterIDOrders(params *GetCharactersCharacterIDOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_orders",
		Method:             "GET",
		PathPattern:        "/v2/characters/{character_id}/orders/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDOrdersReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_orders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetCharactersCharacterIDOrdersHistory lists historical orders by a character

	List cancelled and expired market orders placed by a character up to 90 days in the past.

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDOrdersHistory(params *GetCharactersCharacterIDOrdersHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDOrdersHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDOrdersHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_orders_history",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/orders/history/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDOrdersHistoryReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDOrdersHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_orders_history: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetCorporationsCorporationIDOrders lists open orders from a corporation

	List open market orders placed on behalf of a corporation

---

# This route is cached for up to 1200 seconds

---
Requires one of the following EVE corporation role(s): Accountant, Trader
*/
func (a *Client) GetCorporationsCorporationIDOrders(params *GetCorporationsCorporationIDOrdersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_orders",
		Method:             "GET",
		PathPattern:        "/v3/corporations/{corporation_id}/orders/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDOrdersReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_orders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetCorporationsCorporationIDOrdersHistory lists historical orders from a corporation

	List cancelled and expired market orders placed on behalf of a corporation up to 90 days in the past.

---

# This route is cached for up to 3600 seconds

---
Requires one of the following EVE corporation role(s): Accountant, Trader
*/
func (a *Client) GetCorporationsCorporationIDOrdersHistory(params *GetCorporationsCorporationIDOrdersHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCorporationsCorporationIDOrdersHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDOrdersHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_orders_history",
		Method:             "GET",
		PathPattern:        "/v2/corporations/{corporation_id}/orders/history/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDOrdersHistoryReader{formats: a.formats},
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
	success, ok := result.(*GetCorporationsCorporationIDOrdersHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_corporations_corporation_id_orders_history: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMarketsGroups gets item groups

	Get a list of item groups

---

This route expires daily at 11:05
*/
func (a *Client) GetMarketsGroups(params *GetMarketsGroupsParams, opts ...ClientOption) (*GetMarketsGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_markets_groups",
		Method:             "GET",
		PathPattern:        "/v1/markets/groups/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsGroupsReader{formats: a.formats},
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
	success, ok := result.(*GetMarketsGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_markets_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMarketsGroupsMarketGroupID gets item group information

	Get information on an item group

---

This route expires daily at 11:05
*/
func (a *Client) GetMarketsGroupsMarketGroupID(params *GetMarketsGroupsMarketGroupIDParams, opts ...ClientOption) (*GetMarketsGroupsMarketGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsGroupsMarketGroupIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_markets_groups_market_group_id",
		Method:             "GET",
		PathPattern:        "/v1/markets/groups/{market_group_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsGroupsMarketGroupIDReader{formats: a.formats},
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
	success, ok := result.(*GetMarketsGroupsMarketGroupIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_markets_groups_market_group_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMarketsPrices lists market prices

	Return a list of prices

---

This route is cached for up to 3600 seconds
*/
func (a *Client) GetMarketsPrices(params *GetMarketsPricesParams, opts ...ClientOption) (*GetMarketsPricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsPricesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_markets_prices",
		Method:             "GET",
		PathPattern:        "/v1/markets/prices/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsPricesReader{formats: a.formats},
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
	success, ok := result.(*GetMarketsPricesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_markets_prices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMarketsRegionIDHistory lists historical market statistics in a region

	Return a list of historical market statistics for the specified type in a region

---

This route expires daily at 11:05
*/
func (a *Client) GetMarketsRegionIDHistory(params *GetMarketsRegionIDHistoryParams, opts ...ClientOption) (*GetMarketsRegionIDHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsRegionIDHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_markets_region_id_history",
		Method:             "GET",
		PathPattern:        "/v1/markets/{region_id}/history/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsRegionIDHistoryReader{formats: a.formats},
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
	success, ok := result.(*GetMarketsRegionIDHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_markets_region_id_history: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMarketsRegionIDOrders lists orders in a region

	Return a list of orders in a region

---

This route is cached for up to 300 seconds
*/
func (a *Client) GetMarketsRegionIDOrders(params *GetMarketsRegionIDOrdersParams, opts ...ClientOption) (*GetMarketsRegionIDOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsRegionIDOrdersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_markets_region_id_orders",
		Method:             "GET",
		PathPattern:        "/v1/markets/{region_id}/orders/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsRegionIDOrdersReader{formats: a.formats},
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
	success, ok := result.(*GetMarketsRegionIDOrdersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_markets_region_id_orders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMarketsRegionIDTypes lists type i ds relevant to a market

	Return a list of type IDs that have active orders in the region, for efficient market indexing.

---

This route is cached for up to 600 seconds
*/
func (a *Client) GetMarketsRegionIDTypes(params *GetMarketsRegionIDTypesParams, opts ...ClientOption) (*GetMarketsRegionIDTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsRegionIDTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_markets_region_id_types",
		Method:             "GET",
		PathPattern:        "/v1/markets/{region_id}/types/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsRegionIDTypesReader{formats: a.formats},
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
	success, ok := result.(*GetMarketsRegionIDTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_markets_region_id_types: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetMarketsStructuresStructureID lists orders in a structure

	Return all orders in a structure

---

This route is cached for up to 300 seconds
*/
func (a *Client) GetMarketsStructuresStructureID(params *GetMarketsStructuresStructureIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMarketsStructuresStructureIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsStructuresStructureIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_markets_structures_structure_id",
		Method:             "GET",
		PathPattern:        "/v1/markets/structures/{structure_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsStructuresStructureIDReader{formats: a.formats},
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
	success, ok := result.(*GetMarketsStructuresStructureIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_markets_structures_structure_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
