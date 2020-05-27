// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// GetCharactersCharacterIDFatigueReader is a Reader for the GetCharactersCharacterIDFatigue structure.
type GetCharactersCharacterIDFatigueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDFatigueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDFatigueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDFatigueNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDFatigueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDFatigueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDFatigueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDFatigueEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDFatigueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDFatigueServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDFatigueGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDFatigueOK creates a GetCharactersCharacterIDFatigueOK with default headers values
func NewGetCharactersCharacterIDFatigueOK() *GetCharactersCharacterIDFatigueOK {
	return &GetCharactersCharacterIDFatigueOK{}
}

/* GetCharactersCharacterIDFatigueOK describes a response with status code 200, with default header values.

Jump activation and fatigue information
*/
type GetCharactersCharacterIDFatigueOK struct {

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

	Payload *GetCharactersCharacterIDFatigueOKBody
}

func (o *GetCharactersCharacterIDFatigueOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueOK) GetPayload() *GetCharactersCharacterIDFatigueOKBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCharactersCharacterIDFatigueOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueNotModified creates a GetCharactersCharacterIDFatigueNotModified with default headers values
func NewGetCharactersCharacterIDFatigueNotModified() *GetCharactersCharacterIDFatigueNotModified {
	return &GetCharactersCharacterIDFatigueNotModified{}
}

/* GetCharactersCharacterIDFatigueNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDFatigueNotModified struct {

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

func (o *GetCharactersCharacterIDFatigueNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueNotModified ", 304)
}

func (o *GetCharactersCharacterIDFatigueNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDFatigueBadRequest creates a GetCharactersCharacterIDFatigueBadRequest with default headers values
func NewGetCharactersCharacterIDFatigueBadRequest() *GetCharactersCharacterIDFatigueBadRequest {
	return &GetCharactersCharacterIDFatigueBadRequest{}
}

/* GetCharactersCharacterIDFatigueBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDFatigueBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDFatigueBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueUnauthorized creates a GetCharactersCharacterIDFatigueUnauthorized with default headers values
func NewGetCharactersCharacterIDFatigueUnauthorized() *GetCharactersCharacterIDFatigueUnauthorized {
	return &GetCharactersCharacterIDFatigueUnauthorized{}
}

/* GetCharactersCharacterIDFatigueUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDFatigueUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDFatigueUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueForbidden creates a GetCharactersCharacterIDFatigueForbidden with default headers values
func NewGetCharactersCharacterIDFatigueForbidden() *GetCharactersCharacterIDFatigueForbidden {
	return &GetCharactersCharacterIDFatigueForbidden{}
}

/* GetCharactersCharacterIDFatigueForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDFatigueForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDFatigueForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueEnhanceYourCalm creates a GetCharactersCharacterIDFatigueEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDFatigueEnhanceYourCalm() *GetCharactersCharacterIDFatigueEnhanceYourCalm {
	return &GetCharactersCharacterIDFatigueEnhanceYourCalm{}
}

/* GetCharactersCharacterIDFatigueEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDFatigueEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDFatigueEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueInternalServerError creates a GetCharactersCharacterIDFatigueInternalServerError with default headers values
func NewGetCharactersCharacterIDFatigueInternalServerError() *GetCharactersCharacterIDFatigueInternalServerError {
	return &GetCharactersCharacterIDFatigueInternalServerError{}
}

/* GetCharactersCharacterIDFatigueInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDFatigueInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDFatigueInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueServiceUnavailable creates a GetCharactersCharacterIDFatigueServiceUnavailable with default headers values
func NewGetCharactersCharacterIDFatigueServiceUnavailable() *GetCharactersCharacterIDFatigueServiceUnavailable {
	return &GetCharactersCharacterIDFatigueServiceUnavailable{}
}

/* GetCharactersCharacterIDFatigueServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDFatigueServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDFatigueServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFatigueGatewayTimeout creates a GetCharactersCharacterIDFatigueGatewayTimeout with default headers values
func NewGetCharactersCharacterIDFatigueGatewayTimeout() *GetCharactersCharacterIDFatigueGatewayTimeout {
	return &GetCharactersCharacterIDFatigueGatewayTimeout{}
}

/* GetCharactersCharacterIDFatigueGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDFatigueGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDFatigueGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fatigue/][%d] getCharactersCharacterIdFatigueGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDFatigueGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDFatigueGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDFatigueOKBody get_characters_character_id_fatigue_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDFatigueOKBody
*/
type GetCharactersCharacterIDFatigueOKBody struct {

	// get_characters_character_id_fatigue_jump_fatigue_expire_date
	//
	// Character's jump fatigue expiry
	// Format: date-time
	JumpFatigueExpireDate strfmt.DateTime `json:"jump_fatigue_expire_date,omitempty"`

	// get_characters_character_id_fatigue_last_jump_date
	//
	// Character's last jump activation
	// Format: date-time
	LastJumpDate strfmt.DateTime `json:"last_jump_date,omitempty"`

	// get_characters_character_id_fatigue_last_update_date
	//
	// Character's last jump update
	// Format: date-time
	LastUpdateDate strfmt.DateTime `json:"last_update_date,omitempty"`
}

// Validate validates this get characters character ID fatigue o k body
func (o *GetCharactersCharacterIDFatigueOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJumpFatigueExpireDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastJumpDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDFatigueOKBody) validateJumpFatigueExpireDate(formats strfmt.Registry) error {
	if swag.IsZero(o.JumpFatigueExpireDate) { // not required
		return nil
	}

	if err := validate.FormatOf("getCharactersCharacterIdFatigueOK"+"."+"jump_fatigue_expire_date", "body", "date-time", o.JumpFatigueExpireDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFatigueOKBody) validateLastJumpDate(formats strfmt.Registry) error {
	if swag.IsZero(o.LastJumpDate) { // not required
		return nil
	}

	if err := validate.FormatOf("getCharactersCharacterIdFatigueOK"+"."+"last_jump_date", "body", "date-time", o.LastJumpDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFatigueOKBody) validateLastUpdateDate(formats strfmt.Registry) error {
	if swag.IsZero(o.LastUpdateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("getCharactersCharacterIdFatigueOK"+"."+"last_update_date", "body", "date-time", o.LastUpdateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID fatigue o k body based on context it is used
func (o *GetCharactersCharacterIDFatigueOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDFatigueOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDFatigueOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFatigueOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
