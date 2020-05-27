// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// GetCharactersCharacterIDStandingsReader is a Reader for the GetCharactersCharacterIDStandings structure.
type GetCharactersCharacterIDStandingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDStandingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDStandingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDStandingsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDStandingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDStandingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDStandingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDStandingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDStandingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDStandingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDStandingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDStandingsOK creates a GetCharactersCharacterIDStandingsOK with default headers values
func NewGetCharactersCharacterIDStandingsOK() *GetCharactersCharacterIDStandingsOK {
	return &GetCharactersCharacterIDStandingsOK{}
}

/* GetCharactersCharacterIDStandingsOK describes a response with status code 200, with default header values.

A list of standings
*/
type GetCharactersCharacterIDStandingsOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCharactersCharacterIDStandingsOKBodyItems0
}

func (o *GetCharactersCharacterIDStandingsOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsOK) GetPayload() []*GetCharactersCharacterIDStandingsOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
	}

	// hydrates response header Expires
	hdrExpires := response.GetHeader("Expires")

	if hdrExpires != "" {
		o.Expires = hdrExpires
	}

	// hydrates response header Last-Modified
	hdrLastModified := response.GetHeader("Last-Modified")

	if hdrLastModified != "" {
		o.LastModified = hdrLastModified
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsNotModified creates a GetCharactersCharacterIDStandingsNotModified with default headers values
func NewGetCharactersCharacterIDStandingsNotModified() *GetCharactersCharacterIDStandingsNotModified {
	return &GetCharactersCharacterIDStandingsNotModified{}
}

/* GetCharactersCharacterIDStandingsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDStandingsNotModified struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string
}

func (o *GetCharactersCharacterIDStandingsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsNotModified ", 304)
}

func (o *GetCharactersCharacterIDStandingsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
	}

	// hydrates response header Expires
	hdrExpires := response.GetHeader("Expires")

	if hdrExpires != "" {
		o.Expires = hdrExpires
	}

	// hydrates response header Last-Modified
	hdrLastModified := response.GetHeader("Last-Modified")

	if hdrLastModified != "" {
		o.LastModified = hdrLastModified
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsBadRequest creates a GetCharactersCharacterIDStandingsBadRequest with default headers values
func NewGetCharactersCharacterIDStandingsBadRequest() *GetCharactersCharacterIDStandingsBadRequest {
	return &GetCharactersCharacterIDStandingsBadRequest{}
}

/* GetCharactersCharacterIDStandingsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDStandingsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDStandingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsUnauthorized creates a GetCharactersCharacterIDStandingsUnauthorized with default headers values
func NewGetCharactersCharacterIDStandingsUnauthorized() *GetCharactersCharacterIDStandingsUnauthorized {
	return &GetCharactersCharacterIDStandingsUnauthorized{}
}

/* GetCharactersCharacterIDStandingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDStandingsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDStandingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsForbidden creates a GetCharactersCharacterIDStandingsForbidden with default headers values
func NewGetCharactersCharacterIDStandingsForbidden() *GetCharactersCharacterIDStandingsForbidden {
	return &GetCharactersCharacterIDStandingsForbidden{}
}

/* GetCharactersCharacterIDStandingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDStandingsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDStandingsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsEnhanceYourCalm creates a GetCharactersCharacterIDStandingsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDStandingsEnhanceYourCalm() *GetCharactersCharacterIDStandingsEnhanceYourCalm {
	return &GetCharactersCharacterIDStandingsEnhanceYourCalm{}
}

/* GetCharactersCharacterIDStandingsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDStandingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDStandingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsInternalServerError creates a GetCharactersCharacterIDStandingsInternalServerError with default headers values
func NewGetCharactersCharacterIDStandingsInternalServerError() *GetCharactersCharacterIDStandingsInternalServerError {
	return &GetCharactersCharacterIDStandingsInternalServerError{}
}

/* GetCharactersCharacterIDStandingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDStandingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDStandingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsServiceUnavailable creates a GetCharactersCharacterIDStandingsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDStandingsServiceUnavailable() *GetCharactersCharacterIDStandingsServiceUnavailable {
	return &GetCharactersCharacterIDStandingsServiceUnavailable{}
}

/* GetCharactersCharacterIDStandingsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDStandingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDStandingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsGatewayTimeout creates a GetCharactersCharacterIDStandingsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDStandingsGatewayTimeout() *GetCharactersCharacterIDStandingsGatewayTimeout {
	return &GetCharactersCharacterIDStandingsGatewayTimeout{}
}

/* GetCharactersCharacterIDStandingsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDStandingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDStandingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDStandingsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDStandingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDStandingsOKBodyItems0 get_characters_character_id_standings_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDStandingsOKBodyItems0
*/
type GetCharactersCharacterIDStandingsOKBodyItems0 struct {

	// get_characters_character_id_standings_from_id
	//
	// from_id integer
	// Required: true
	FromID *int32 `json:"from_id"`

	// get_characters_character_id_standings_from_type
	//
	// from_type string
	// Required: true
	// Enum: [agent npc_corp faction]
	FromType *string `json:"from_type"`

	// get_characters_character_id_standings_standing
	//
	// standing number
	// Required: true
	Standing *float32 `json:"standing"`
}

// Validate validates this get characters character ID standings o k body items0
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFromID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFromType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStanding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", o.FromID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["agent","npc_corp","faction"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum = append(getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDStandingsOKBodyItems0FromTypeAgent captures enum value "agent"
	GetCharactersCharacterIDStandingsOKBodyItems0FromTypeAgent string = "agent"

	// GetCharactersCharacterIDStandingsOKBodyItems0FromTypeNpcCorp captures enum value "npc_corp"
	GetCharactersCharacterIDStandingsOKBodyItems0FromTypeNpcCorp string = "npc_corp"

	// GetCharactersCharacterIDStandingsOKBodyItems0FromTypeFaction captures enum value "faction"
	GetCharactersCharacterIDStandingsOKBodyItems0FromTypeFaction string = "faction"
)

// prop value enum
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateFromTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdStandingsOKBodyItems0TypeFromTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateFromType(formats strfmt.Registry) error {

	if err := validate.Required("from_type", "body", o.FromType); err != nil {
		return err
	}

	// value enum
	if err := o.validateFromTypeEnum("from_type", "body", *o.FromType); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDStandingsOKBodyItems0) validateStanding(formats strfmt.Registry) error {

	if err := validate.Required("standing", "body", o.Standing); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID standings o k body items0 based on context it is used
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDStandingsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDStandingsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
