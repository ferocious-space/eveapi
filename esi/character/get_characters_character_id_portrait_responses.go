// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ferocious-space/eveapi/models"
)

// GetCharactersCharacterIDPortraitReader is a Reader for the GetCharactersCharacterIDPortrait structure.
type GetCharactersCharacterIDPortraitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDPortraitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDPortraitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDPortraitNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDPortraitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCharactersCharacterIDPortraitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDPortraitEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDPortraitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDPortraitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDPortraitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDPortraitOK creates a GetCharactersCharacterIDPortraitOK with default headers values
func NewGetCharactersCharacterIDPortraitOK() *GetCharactersCharacterIDPortraitOK {
	return &GetCharactersCharacterIDPortraitOK{}
}

/* GetCharactersCharacterIDPortraitOK describes a response with status code 200, with default header values.

Public data for the given character
*/
type GetCharactersCharacterIDPortraitOK struct {

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

	Payload *GetCharactersCharacterIDPortraitOKBody
}

func (o *GetCharactersCharacterIDPortraitOK) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDPortraitOK) GetPayload() *GetCharactersCharacterIDPortraitOKBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDPortraitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCharactersCharacterIDPortraitOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitNotModified creates a GetCharactersCharacterIDPortraitNotModified with default headers values
func NewGetCharactersCharacterIDPortraitNotModified() *GetCharactersCharacterIDPortraitNotModified {
	return &GetCharactersCharacterIDPortraitNotModified{}
}

/* GetCharactersCharacterIDPortraitNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDPortraitNotModified struct {

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

func (o *GetCharactersCharacterIDPortraitNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitNotModified ", 304)
}

func (o *GetCharactersCharacterIDPortraitNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDPortraitBadRequest creates a GetCharactersCharacterIDPortraitBadRequest with default headers values
func NewGetCharactersCharacterIDPortraitBadRequest() *GetCharactersCharacterIDPortraitBadRequest {
	return &GetCharactersCharacterIDPortraitBadRequest{}
}

/* GetCharactersCharacterIDPortraitBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDPortraitBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDPortraitBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDPortraitBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDPortraitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitNotFound creates a GetCharactersCharacterIDPortraitNotFound with default headers values
func NewGetCharactersCharacterIDPortraitNotFound() *GetCharactersCharacterIDPortraitNotFound {
	return &GetCharactersCharacterIDPortraitNotFound{}
}

/* GetCharactersCharacterIDPortraitNotFound describes a response with status code 404, with default header values.

No image server for this datasource
*/
type GetCharactersCharacterIDPortraitNotFound struct {
	Payload *GetCharactersCharacterIDPortraitNotFoundBody
}

func (o *GetCharactersCharacterIDPortraitNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitNotFound  %+v", 404, o.Payload)
}
func (o *GetCharactersCharacterIDPortraitNotFound) GetPayload() *GetCharactersCharacterIDPortraitNotFoundBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDPortraitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCharactersCharacterIDPortraitNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitEnhanceYourCalm creates a GetCharactersCharacterIDPortraitEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDPortraitEnhanceYourCalm() *GetCharactersCharacterIDPortraitEnhanceYourCalm {
	return &GetCharactersCharacterIDPortraitEnhanceYourCalm{}
}

/* GetCharactersCharacterIDPortraitEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDPortraitEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDPortraitEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDPortraitEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDPortraitEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitInternalServerError creates a GetCharactersCharacterIDPortraitInternalServerError with default headers values
func NewGetCharactersCharacterIDPortraitInternalServerError() *GetCharactersCharacterIDPortraitInternalServerError {
	return &GetCharactersCharacterIDPortraitInternalServerError{}
}

/* GetCharactersCharacterIDPortraitInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDPortraitInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDPortraitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDPortraitInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDPortraitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitServiceUnavailable creates a GetCharactersCharacterIDPortraitServiceUnavailable with default headers values
func NewGetCharactersCharacterIDPortraitServiceUnavailable() *GetCharactersCharacterIDPortraitServiceUnavailable {
	return &GetCharactersCharacterIDPortraitServiceUnavailable{}
}

/* GetCharactersCharacterIDPortraitServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDPortraitServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDPortraitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDPortraitServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDPortraitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDPortraitGatewayTimeout creates a GetCharactersCharacterIDPortraitGatewayTimeout with default headers values
func NewGetCharactersCharacterIDPortraitGatewayTimeout() *GetCharactersCharacterIDPortraitGatewayTimeout {
	return &GetCharactersCharacterIDPortraitGatewayTimeout{}
}

/* GetCharactersCharacterIDPortraitGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDPortraitGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDPortraitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/portrait/][%d] getCharactersCharacterIdPortraitGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDPortraitGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDPortraitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDPortraitNotFoundBody get_characters_character_id_portrait_not_found
//
// No image server for this datasource
swagger:model GetCharactersCharacterIDPortraitNotFoundBody
*/
type GetCharactersCharacterIDPortraitNotFoundBody struct {

	// get_characters_character_id_portrait_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character ID portrait not found body
func (o *GetCharactersCharacterIDPortraitNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get characters character ID portrait not found body based on context it is used
func (o *GetCharactersCharacterIDPortraitNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDPortraitNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDPortraitNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDPortraitNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCharactersCharacterIDPortraitOKBody get_characters_character_id_portrait_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDPortraitOKBody
*/
type GetCharactersCharacterIDPortraitOKBody struct {

	// get_characters_character_id_portrait_px128x128
	//
	// px128x128 string
	Px128x128 string `json:"px128x128,omitempty"`

	// get_characters_character_id_portrait_px256x256
	//
	// px256x256 string
	Px256x256 string `json:"px256x256,omitempty"`

	// get_characters_character_id_portrait_px512x512
	//
	// px512x512 string
	Px512x512 string `json:"px512x512,omitempty"`

	// get_characters_character_id_portrait_px64x64
	//
	// px64x64 string
	Px64x64 string `json:"px64x64,omitempty"`
}

// Validate validates this get characters character ID portrait o k body
func (o *GetCharactersCharacterIDPortraitOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get characters character ID portrait o k body based on context it is used
func (o *GetCharactersCharacterIDPortraitOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDPortraitOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDPortraitOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDPortraitOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
