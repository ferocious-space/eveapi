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

package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fleets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fleets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteFleetsFleetIDMembersMemberID(params *DeleteFleetsFleetIDMembersMemberIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFleetsFleetIDMembersMemberIDNoContent, error)

	DeleteFleetsFleetIDSquadsSquadID(params *DeleteFleetsFleetIDSquadsSquadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFleetsFleetIDSquadsSquadIDNoContent, error)

	DeleteFleetsFleetIDWingsWingID(params *DeleteFleetsFleetIDWingsWingIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFleetsFleetIDWingsWingIDNoContent, error)

	GetCharactersCharacterIDFleet(params *GetCharactersCharacterIDFleetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDFleetOK, error)

	GetFleetsFleetID(params *GetFleetsFleetIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFleetsFleetIDOK, error)

	GetFleetsFleetIDMembers(params *GetFleetsFleetIDMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFleetsFleetIDMembersOK, error)

	GetFleetsFleetIDWings(params *GetFleetsFleetIDWingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFleetsFleetIDWingsOK, error)

	PostFleetsFleetIDMembers(params *PostFleetsFleetIDMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFleetsFleetIDMembersNoContent, error)

	PostFleetsFleetIDWings(params *PostFleetsFleetIDWingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFleetsFleetIDWingsCreated, error)

	PostFleetsFleetIDWingsWingIDSquads(params *PostFleetsFleetIDWingsWingIDSquadsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFleetsFleetIDWingsWingIDSquadsCreated, error)

	PutFleetsFleetID(params *PutFleetsFleetIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDNoContent, error)

	PutFleetsFleetIDMembersMemberID(params *PutFleetsFleetIDMembersMemberIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDMembersMemberIDNoContent, error)

	PutFleetsFleetIDSquadsSquadID(params *PutFleetsFleetIDSquadsSquadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDSquadsSquadIDNoContent, error)

	PutFleetsFleetIDWingsWingID(params *PutFleetsFleetIDWingsWingIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDWingsWingIDNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteFleetsFleetIDMembersMemberID kicks fleet member

  Kick a fleet member

---

*/
func (a *Client) DeleteFleetsFleetIDMembersMemberID(params *DeleteFleetsFleetIDMembersMemberIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFleetsFleetIDMembersMemberIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFleetsFleetIDMembersMemberIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_fleets_fleet_id_members_member_id",
		Method:             "DELETE",
		PathPattern:        "/v1/fleets/{fleet_id}/members/{member_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFleetsFleetIDMembersMemberIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteFleetsFleetIDMembersMemberIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_fleets_fleet_id_members_member_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteFleetsFleetIDSquadsSquadID deletes fleet squad

  Delete a fleet squad, only empty squads can be deleted

---

*/
func (a *Client) DeleteFleetsFleetIDSquadsSquadID(params *DeleteFleetsFleetIDSquadsSquadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFleetsFleetIDSquadsSquadIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFleetsFleetIDSquadsSquadIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_fleets_fleet_id_squads_squad_id",
		Method:             "DELETE",
		PathPattern:        "/v1/fleets/{fleet_id}/squads/{squad_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFleetsFleetIDSquadsSquadIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteFleetsFleetIDSquadsSquadIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_fleets_fleet_id_squads_squad_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteFleetsFleetIDWingsWingID deletes fleet wing

  Delete a fleet wing, only empty wings can be deleted. The wing may contain squads, but the squads must be empty

---

*/
func (a *Client) DeleteFleetsFleetIDWingsWingID(params *DeleteFleetsFleetIDWingsWingIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteFleetsFleetIDWingsWingIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFleetsFleetIDWingsWingIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_fleets_fleet_id_wings_wing_id",
		Method:             "DELETE",
		PathPattern:        "/v1/fleets/{fleet_id}/wings/{wing_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFleetsFleetIDWingsWingIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteFleetsFleetIDWingsWingIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_fleets_fleet_id_wings_wing_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCharactersCharacterIDFleet gets character fleet info

  Return the fleet ID the character is in, if any.

---

This route is cached for up to 60 seconds

---
Warning: This route has an upgrade available

---
[Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#GET-/characters/{character_id}/fleet/)
*/
func (a *Client) GetCharactersCharacterIDFleet(params *GetCharactersCharacterIDFleetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCharactersCharacterIDFleetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDFleetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_characters_character_id_fleet",
		Method:             "GET",
		PathPattern:        "/v1/characters/{character_id}/fleet/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDFleetReader{formats: a.formats},
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
	success, ok := result.(*GetCharactersCharacterIDFleetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_characters_character_id_fleet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFleetsFleetID gets fleet information

  Return details about a fleet

---

This route is cached for up to 5 seconds
*/
func (a *Client) GetFleetsFleetID(params *GetFleetsFleetIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFleetsFleetIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetsFleetIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fleets_fleet_id",
		Method:             "GET",
		PathPattern:        "/v1/fleets/{fleet_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetsFleetIDReader{formats: a.formats},
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
	success, ok := result.(*GetFleetsFleetIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fleets_fleet_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFleetsFleetIDMembers gets fleet members

  Return information about fleet members

---

This route is cached for up to 5 seconds
*/
func (a *Client) GetFleetsFleetIDMembers(params *GetFleetsFleetIDMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFleetsFleetIDMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetsFleetIDMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fleets_fleet_id_members",
		Method:             "GET",
		PathPattern:        "/v1/fleets/{fleet_id}/members/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetsFleetIDMembersReader{formats: a.formats},
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
	success, ok := result.(*GetFleetsFleetIDMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fleets_fleet_id_members: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFleetsFleetIDWings gets fleet wings

  Return information about wings in a fleet

---

This route is cached for up to 5 seconds
*/
func (a *Client) GetFleetsFleetIDWings(params *GetFleetsFleetIDWingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFleetsFleetIDWingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetsFleetIDWingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_fleets_fleet_id_wings",
		Method:             "GET",
		PathPattern:        "/v1/fleets/{fleet_id}/wings/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetsFleetIDWingsReader{formats: a.formats},
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
	success, ok := result.(*GetFleetsFleetIDWingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_fleets_fleet_id_wings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostFleetsFleetIDMembers creates fleet invitation

  Invite a character into the fleet. If a character has a CSPA charge set it is not possible to invite them to the fleet using ESI

---

*/
func (a *Client) PostFleetsFleetIDMembers(params *PostFleetsFleetIDMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFleetsFleetIDMembersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFleetsFleetIDMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_fleets_fleet_id_members",
		Method:             "POST",
		PathPattern:        "/v1/fleets/{fleet_id}/members/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFleetsFleetIDMembersReader{formats: a.formats},
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
	success, ok := result.(*PostFleetsFleetIDMembersNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_fleets_fleet_id_members: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostFleetsFleetIDWings creates fleet wing

  Create a new wing in a fleet

---

*/
func (a *Client) PostFleetsFleetIDWings(params *PostFleetsFleetIDWingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFleetsFleetIDWingsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFleetsFleetIDWingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_fleets_fleet_id_wings",
		Method:             "POST",
		PathPattern:        "/v1/fleets/{fleet_id}/wings/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFleetsFleetIDWingsReader{formats: a.formats},
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
	success, ok := result.(*PostFleetsFleetIDWingsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_fleets_fleet_id_wings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostFleetsFleetIDWingsWingIDSquads creates fleet squad

  Create a new squad in a fleet

---

*/
func (a *Client) PostFleetsFleetIDWingsWingIDSquads(params *PostFleetsFleetIDWingsWingIDSquadsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostFleetsFleetIDWingsWingIDSquadsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFleetsFleetIDWingsWingIDSquadsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_fleets_fleet_id_wings_wing_id_squads",
		Method:             "POST",
		PathPattern:        "/v1/fleets/{fleet_id}/wings/{wing_id}/squads/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFleetsFleetIDWingsWingIDSquadsReader{formats: a.formats},
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
	success, ok := result.(*PostFleetsFleetIDWingsWingIDSquadsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_fleets_fleet_id_wings_wing_id_squads: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutFleetsFleetID updates fleet

  Update settings about a fleet

---

*/
func (a *Client) PutFleetsFleetID(params *PutFleetsFleetIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFleetsFleetIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "put_fleets_fleet_id",
		Method:             "PUT",
		PathPattern:        "/v1/fleets/{fleet_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutFleetsFleetIDReader{formats: a.formats},
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
	success, ok := result.(*PutFleetsFleetIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for put_fleets_fleet_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutFleetsFleetIDMembersMemberID moves fleet member

  Move a fleet member around

---

*/
func (a *Client) PutFleetsFleetIDMembersMemberID(params *PutFleetsFleetIDMembersMemberIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDMembersMemberIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFleetsFleetIDMembersMemberIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "put_fleets_fleet_id_members_member_id",
		Method:             "PUT",
		PathPattern:        "/v1/fleets/{fleet_id}/members/{member_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutFleetsFleetIDMembersMemberIDReader{formats: a.formats},
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
	success, ok := result.(*PutFleetsFleetIDMembersMemberIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for put_fleets_fleet_id_members_member_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutFleetsFleetIDSquadsSquadID renames fleet squad

  Rename a fleet squad

---

*/
func (a *Client) PutFleetsFleetIDSquadsSquadID(params *PutFleetsFleetIDSquadsSquadIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDSquadsSquadIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFleetsFleetIDSquadsSquadIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "put_fleets_fleet_id_squads_squad_id",
		Method:             "PUT",
		PathPattern:        "/v1/fleets/{fleet_id}/squads/{squad_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutFleetsFleetIDSquadsSquadIDReader{formats: a.formats},
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
	success, ok := result.(*PutFleetsFleetIDSquadsSquadIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for put_fleets_fleet_id_squads_squad_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutFleetsFleetIDWingsWingID renames fleet wing

  Rename a fleet wing

---

*/
func (a *Client) PutFleetsFleetIDWingsWingID(params *PutFleetsFleetIDWingsWingIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutFleetsFleetIDWingsWingIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFleetsFleetIDWingsWingIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "put_fleets_fleet_id_wings_wing_id",
		Method:             "PUT",
		PathPattern:        "/v1/fleets/{fleet_id}/wings/{wing_id}/",
		ProducesMediaTypes: []string{"application/json", "text/html; charset=utf-8", "text/plain; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutFleetsFleetIDWingsWingIDReader{formats: a.formats},
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
	success, ok := result.(*PutFleetsFleetIDWingsWingIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for put_fleets_fleet_id_wings_wing_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
