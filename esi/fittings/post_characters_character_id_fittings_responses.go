// Code generated by go-swagger; DO NOT EDIT.

package fittings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// PostCharactersCharacterIDFittingsReader is a Reader for the PostCharactersCharacterIDFittings structure.
type PostCharactersCharacterIDFittingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDFittingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostCharactersCharacterIDFittingsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCharactersCharacterIDFittingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCharactersCharacterIDFittingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCharactersCharacterIDFittingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostCharactersCharacterIDFittingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCharactersCharacterIDFittingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCharactersCharacterIDFittingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCharactersCharacterIDFittingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCharactersCharacterIDFittingsCreated creates a PostCharactersCharacterIDFittingsCreated with default headers values
func NewPostCharactersCharacterIDFittingsCreated() *PostCharactersCharacterIDFittingsCreated {
	return &PostCharactersCharacterIDFittingsCreated{}
}

/* PostCharactersCharacterIDFittingsCreated describes a response with status code 201, with default header values.

A list of fittings
*/
type PostCharactersCharacterIDFittingsCreated struct {
	Payload *PostCharactersCharacterIDFittingsCreatedBody
}

func (o *PostCharactersCharacterIDFittingsCreated) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsCreated  %+v", 201, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsCreated) GetPayload() *PostCharactersCharacterIDFittingsCreatedBody {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostCharactersCharacterIDFittingsCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsBadRequest creates a PostCharactersCharacterIDFittingsBadRequest with default headers values
func NewPostCharactersCharacterIDFittingsBadRequest() *PostCharactersCharacterIDFittingsBadRequest {
	return &PostCharactersCharacterIDFittingsBadRequest{}
}

/* PostCharactersCharacterIDFittingsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostCharactersCharacterIDFittingsBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostCharactersCharacterIDFittingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsBadRequest  %+v", 400, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsUnauthorized creates a PostCharactersCharacterIDFittingsUnauthorized with default headers values
func NewPostCharactersCharacterIDFittingsUnauthorized() *PostCharactersCharacterIDFittingsUnauthorized {
	return &PostCharactersCharacterIDFittingsUnauthorized{}
}

/* PostCharactersCharacterIDFittingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostCharactersCharacterIDFittingsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostCharactersCharacterIDFittingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsUnauthorized  %+v", 401, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsForbidden creates a PostCharactersCharacterIDFittingsForbidden with default headers values
func NewPostCharactersCharacterIDFittingsForbidden() *PostCharactersCharacterIDFittingsForbidden {
	return &PostCharactersCharacterIDFittingsForbidden{}
}

/* PostCharactersCharacterIDFittingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostCharactersCharacterIDFittingsForbidden struct {
	Payload *models.Forbidden
}

func (o *PostCharactersCharacterIDFittingsForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsForbidden  %+v", 403, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsEnhanceYourCalm creates a PostCharactersCharacterIDFittingsEnhanceYourCalm with default headers values
func NewPostCharactersCharacterIDFittingsEnhanceYourCalm() *PostCharactersCharacterIDFittingsEnhanceYourCalm {
	return &PostCharactersCharacterIDFittingsEnhanceYourCalm{}
}

/* PostCharactersCharacterIDFittingsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostCharactersCharacterIDFittingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostCharactersCharacterIDFittingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsInternalServerError creates a PostCharactersCharacterIDFittingsInternalServerError with default headers values
func NewPostCharactersCharacterIDFittingsInternalServerError() *PostCharactersCharacterIDFittingsInternalServerError {
	return &PostCharactersCharacterIDFittingsInternalServerError{}
}

/* PostCharactersCharacterIDFittingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostCharactersCharacterIDFittingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostCharactersCharacterIDFittingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsServiceUnavailable creates a PostCharactersCharacterIDFittingsServiceUnavailable with default headers values
func NewPostCharactersCharacterIDFittingsServiceUnavailable() *PostCharactersCharacterIDFittingsServiceUnavailable {
	return &PostCharactersCharacterIDFittingsServiceUnavailable{}
}

/* PostCharactersCharacterIDFittingsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostCharactersCharacterIDFittingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostCharactersCharacterIDFittingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDFittingsGatewayTimeout creates a PostCharactersCharacterIDFittingsGatewayTimeout with default headers values
func NewPostCharactersCharacterIDFittingsGatewayTimeout() *PostCharactersCharacterIDFittingsGatewayTimeout {
	return &PostCharactersCharacterIDFittingsGatewayTimeout{}
}

/* PostCharactersCharacterIDFittingsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostCharactersCharacterIDFittingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostCharactersCharacterIDFittingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/fittings/][%d] postCharactersCharacterIdFittingsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *PostCharactersCharacterIDFittingsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostCharactersCharacterIDFittingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostCharactersCharacterIDFittingsBody post_characters_character_id_fittings_fitting
//
// fitting object
swagger:model PostCharactersCharacterIDFittingsBody
*/
type PostCharactersCharacterIDFittingsBody struct {

	// post_characters_character_id_fittings_description
	//
	// description string
	// Required: true
	// Max Length: 500
	// Min Length: 0
	Description *string `json:"description"`

	// post_characters_character_id_fittings_items
	//
	// items array
	// Required: true
	// Max Items: 255
	// Min Items: 1
	Items []*PostCharactersCharacterIDFittingsParamsBodyItemsItems0 `json:"items"`

	// post_characters_character_id_fittings_name
	//
	// name string
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// post_characters_character_id_fittings_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`
}

// Validate validates this post characters character ID fittings body
func (o *PostCharactersCharacterIDFittingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCharactersCharacterIDFittingsBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("fitting"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	if err := validate.MinLength("fitting"+"."+"description", "body", *o.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("fitting"+"."+"description", "body", *o.Description, 500); err != nil {
		return err
	}

	return nil
}

func (o *PostCharactersCharacterIDFittingsBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("fitting"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	iItemsSize := int64(len(o.Items))

	if err := validate.MinItems("fitting"+"."+"items", "body", iItemsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("fitting"+"."+"items", "body", iItemsSize, 255); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fitting" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostCharactersCharacterIDFittingsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("fitting"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("fitting"+"."+"name", "body", *o.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("fitting"+"."+"name", "body", *o.Name, 50); err != nil {
		return err
	}

	return nil
}

func (o *PostCharactersCharacterIDFittingsBody) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("fitting"+"."+"ship_type_id", "body", o.ShipTypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post characters character ID fittings body based on the context it is used
func (o *PostCharactersCharacterIDFittingsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCharactersCharacterIDFittingsBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fitting" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostCharactersCharacterIDFittingsCreatedBody post_characters_character_id_fittings_created
//
// 201 created object
swagger:model PostCharactersCharacterIDFittingsCreatedBody
*/
type PostCharactersCharacterIDFittingsCreatedBody struct {

	// post_characters_character_id_fittings_fitting_id
	//
	// fitting_id integer
	// Required: true
	FittingID *int32 `json:"fitting_id"`
}

// Validate validates this post characters character ID fittings created body
func (o *PostCharactersCharacterIDFittingsCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFittingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCharactersCharacterIDFittingsCreatedBody) validateFittingID(formats strfmt.Registry) error {

	if err := validate.Required("postCharactersCharacterIdFittingsCreated"+"."+"fitting_id", "body", o.FittingID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post characters character ID fittings created body based on context it is used
func (o *PostCharactersCharacterIDFittingsCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostCharactersCharacterIDFittingsParamsBodyItemsItems0 post_characters_character_id_fittings_item
//
// item object
swagger:model PostCharactersCharacterIDFittingsParamsBodyItemsItems0
*/
type PostCharactersCharacterIDFittingsParamsBodyItemsItems0 struct {

	// post_characters_character_id_fittings_flag
	//
	// Fitting location for the item. Entries placed in 'Invalid' will be discarded. If this leaves the fitting with nothing, it will cause an error.
	// Required: true
	// Enum: [Cargo DroneBay FighterBay HiSlot0 HiSlot1 HiSlot2 HiSlot3 HiSlot4 HiSlot5 HiSlot6 HiSlot7 Invalid LoSlot0 LoSlot1 LoSlot2 LoSlot3 LoSlot4 LoSlot5 LoSlot6 LoSlot7 MedSlot0 MedSlot1 MedSlot2 MedSlot3 MedSlot4 MedSlot5 MedSlot6 MedSlot7 RigSlot0 RigSlot1 RigSlot2 ServiceSlot0 ServiceSlot1 ServiceSlot2 ServiceSlot3 ServiceSlot4 ServiceSlot5 ServiceSlot6 ServiceSlot7 SubSystemSlot0 SubSystemSlot1 SubSystemSlot2 SubSystemSlot3]
	Flag *string `json:"flag"`

	// post_characters_character_id_fittings_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// post_characters_character_id_fittings_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this post characters character ID fittings params body items items0
func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postCharactersCharacterIdFittingsParamsBodyItemsItems0TypeFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cargo","DroneBay","FighterBay","HiSlot0","HiSlot1","HiSlot2","HiSlot3","HiSlot4","HiSlot5","HiSlot6","HiSlot7","Invalid","LoSlot0","LoSlot1","LoSlot2","LoSlot3","LoSlot4","LoSlot5","LoSlot6","LoSlot7","MedSlot0","MedSlot1","MedSlot2","MedSlot3","MedSlot4","MedSlot5","MedSlot6","MedSlot7","RigSlot0","RigSlot1","RigSlot2","ServiceSlot0","ServiceSlot1","ServiceSlot2","ServiceSlot3","ServiceSlot4","ServiceSlot5","ServiceSlot6","ServiceSlot7","SubSystemSlot0","SubSystemSlot1","SubSystemSlot2","SubSystemSlot3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postCharactersCharacterIdFittingsParamsBodyItemsItems0TypeFlagPropEnum = append(postCharactersCharacterIdFittingsParamsBodyItemsItems0TypeFlagPropEnum, v)
	}
}

const (

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagCargo captures enum value "Cargo"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagCargo string = "Cargo"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagDroneBay captures enum value "DroneBay"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagDroneBay string = "DroneBay"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagFighterBay captures enum value "FighterBay"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagFighterBay string = "FighterBay"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot0 captures enum value "HiSlot0"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot0 string = "HiSlot0"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot1 captures enum value "HiSlot1"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot1 string = "HiSlot1"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot2 captures enum value "HiSlot2"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot2 string = "HiSlot2"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot3 captures enum value "HiSlot3"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot3 string = "HiSlot3"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot4 captures enum value "HiSlot4"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot4 string = "HiSlot4"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot5 captures enum value "HiSlot5"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot5 string = "HiSlot5"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot6 captures enum value "HiSlot6"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot6 string = "HiSlot6"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot7 captures enum value "HiSlot7"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagHiSlot7 string = "HiSlot7"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagInvalid captures enum value "Invalid"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagInvalid string = "Invalid"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot0 captures enum value "LoSlot0"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot0 string = "LoSlot0"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot1 captures enum value "LoSlot1"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot1 string = "LoSlot1"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot2 captures enum value "LoSlot2"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot2 string = "LoSlot2"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot3 captures enum value "LoSlot3"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot3 string = "LoSlot3"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot4 captures enum value "LoSlot4"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot4 string = "LoSlot4"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot5 captures enum value "LoSlot5"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot5 string = "LoSlot5"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot6 captures enum value "LoSlot6"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot6 string = "LoSlot6"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot7 captures enum value "LoSlot7"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagLoSlot7 string = "LoSlot7"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot0 captures enum value "MedSlot0"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot0 string = "MedSlot0"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot1 captures enum value "MedSlot1"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot1 string = "MedSlot1"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot2 captures enum value "MedSlot2"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot2 string = "MedSlot2"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot3 captures enum value "MedSlot3"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot3 string = "MedSlot3"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot4 captures enum value "MedSlot4"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot4 string = "MedSlot4"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot5 captures enum value "MedSlot5"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot5 string = "MedSlot5"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot6 captures enum value "MedSlot6"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot6 string = "MedSlot6"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot7 captures enum value "MedSlot7"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagMedSlot7 string = "MedSlot7"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagRigSlot0 captures enum value "RigSlot0"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagRigSlot0 string = "RigSlot0"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagRigSlot1 captures enum value "RigSlot1"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagRigSlot1 string = "RigSlot1"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagRigSlot2 captures enum value "RigSlot2"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagRigSlot2 string = "RigSlot2"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot0 captures enum value "ServiceSlot0"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot0 string = "ServiceSlot0"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot1 captures enum value "ServiceSlot1"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot1 string = "ServiceSlot1"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot2 captures enum value "ServiceSlot2"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot2 string = "ServiceSlot2"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot3 captures enum value "ServiceSlot3"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot3 string = "ServiceSlot3"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot4 captures enum value "ServiceSlot4"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot4 string = "ServiceSlot4"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot5 captures enum value "ServiceSlot5"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot5 string = "ServiceSlot5"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot6 captures enum value "ServiceSlot6"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot6 string = "ServiceSlot6"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot7 captures enum value "ServiceSlot7"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagServiceSlot7 string = "ServiceSlot7"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot0 captures enum value "SubSystemSlot0"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot0 string = "SubSystemSlot0"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot1 captures enum value "SubSystemSlot1"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot1 string = "SubSystemSlot1"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot2 captures enum value "SubSystemSlot2"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot2 string = "SubSystemSlot2"

	// PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot3 captures enum value "SubSystemSlot3"
	PostCharactersCharacterIDFittingsParamsBodyItemsItems0FlagSubSystemSlot3 string = "SubSystemSlot3"
)

// prop value enum
func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) validateFlagEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postCharactersCharacterIdFittingsParamsBodyItemsItems0TypeFlagPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", o.Flag); err != nil {
		return err
	}

	// value enum
	if err := o.validateFlagEnum("flag", "body", *o.Flag); err != nil {
		return err
	}

	return nil
}

func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post characters character ID fittings params body items items0 based on context it is used
func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDFittingsParamsBodyItemsItems0) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsParamsBodyItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
