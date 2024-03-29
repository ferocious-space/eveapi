// Code generated by go-swagger; DO NOT EDIT.

package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dogma API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dogma API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDogmaAttributes(params *GetDogmaAttributesParams, opts ...ClientOption) (*GetDogmaAttributesOK, error)

	GetDogmaAttributesAttributeID(params *GetDogmaAttributesAttributeIDParams, opts ...ClientOption) (*GetDogmaAttributesAttributeIDOK, error)

	GetDogmaDynamicItemsTypeIDItemID(params *GetDogmaDynamicItemsTypeIDItemIDParams, opts ...ClientOption) (*GetDogmaDynamicItemsTypeIDItemIDOK, error)

	GetDogmaEffects(params *GetDogmaEffectsParams, opts ...ClientOption) (*GetDogmaEffectsOK, error)

	GetDogmaEffectsEffectID(params *GetDogmaEffectsEffectIDParams, opts ...ClientOption) (*GetDogmaEffectsEffectIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetDogmaAttributes gets attributes

	Get a list of dogma attribute ids

---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaAttributes(params *GetDogmaAttributesParams, opts ...ClientOption) (*GetDogmaAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaAttributesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_dogma_attributes",
		Method:             "GET",
		PathPattern:        "/v1/dogma/attributes/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaAttributesReader{formats: a.formats},
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
	success, ok := result.(*GetDogmaAttributesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_dogma_attributes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetDogmaAttributesAttributeID gets attribute information

	Get information on a dogma attribute

---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaAttributesAttributeID(params *GetDogmaAttributesAttributeIDParams, opts ...ClientOption) (*GetDogmaAttributesAttributeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaAttributesAttributeIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_dogma_attributes_attribute_id",
		Method:             "GET",
		PathPattern:        "/v1/dogma/attributes/{attribute_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaAttributesAttributeIDReader{formats: a.formats},
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
	success, ok := result.(*GetDogmaAttributesAttributeIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_dogma_attributes_attribute_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetDogmaDynamicItemsTypeIDItemID gets dynamic item information

	Returns info about a dynamic item resulting from mutation with a mutaplasmid.

---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaDynamicItemsTypeIDItemID(params *GetDogmaDynamicItemsTypeIDItemIDParams, opts ...ClientOption) (*GetDogmaDynamicItemsTypeIDItemIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaDynamicItemsTypeIDItemIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_dogma_dynamic_items_type_id_item_id",
		Method:             "GET",
		PathPattern:        "/v1/dogma/dynamic/items/{type_id}/{item_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaDynamicItemsTypeIDItemIDReader{formats: a.formats},
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
	success, ok := result.(*GetDogmaDynamicItemsTypeIDItemIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_dogma_dynamic_items_type_id_item_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetDogmaEffects gets effects

	Get a list of dogma effect ids

---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaEffects(params *GetDogmaEffectsParams, opts ...ClientOption) (*GetDogmaEffectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaEffectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_dogma_effects",
		Method:             "GET",
		PathPattern:        "/v1/dogma/effects/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaEffectsReader{formats: a.formats},
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
	success, ok := result.(*GetDogmaEffectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_dogma_effects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetDogmaEffectsEffectID gets effect information

	Get information on a dogma effect

---

This route expires daily at 11:05
*/
func (a *Client) GetDogmaEffectsEffectID(params *GetDogmaEffectsEffectIDParams, opts ...ClientOption) (*GetDogmaEffectsEffectIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDogmaEffectsEffectIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_dogma_effects_effect_id",
		Method:             "GET",
		PathPattern:        "/v2/dogma/effects/{effect_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDogmaEffectsEffectIDReader{formats: a.formats},
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
	success, ok := result.(*GetDogmaEffectsEffectIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_dogma_effects_effect_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
