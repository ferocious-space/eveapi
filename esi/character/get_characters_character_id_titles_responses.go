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

// GetCharactersCharacterIDTitlesReader is a Reader for the GetCharactersCharacterIDTitles structure.
type GetCharactersCharacterIDTitlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDTitlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDTitlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDTitlesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDTitlesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDTitlesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDTitlesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDTitlesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDTitlesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDTitlesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDTitlesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDTitlesOK creates a GetCharactersCharacterIDTitlesOK with default headers values
func NewGetCharactersCharacterIDTitlesOK() *GetCharactersCharacterIDTitlesOK {
	return &GetCharactersCharacterIDTitlesOK{}
}

/*
GetCharactersCharacterIDTitlesOK describes a response with status code 200, with default header values.

A list of titles
*/
type GetCharactersCharacterIDTitlesOK struct {

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

	Payload []*GetCharactersCharacterIDTitlesOKBodyItems0
}

// IsSuccess returns true when this get characters character Id titles o k response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id titles o k response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles o k response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id titles o k response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id titles o k response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCharactersCharacterIDTitlesOK) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesOK) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesOK) GetPayload() []*GetCharactersCharacterIDTitlesOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDTitlesNotModified creates a GetCharactersCharacterIDTitlesNotModified with default headers values
func NewGetCharactersCharacterIDTitlesNotModified() *GetCharactersCharacterIDTitlesNotModified {
	return &GetCharactersCharacterIDTitlesNotModified{}
}

/*
GetCharactersCharacterIDTitlesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDTitlesNotModified struct {

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

// IsSuccess returns true when this get characters character Id titles not modified response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles not modified response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id titles not modified response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id titles not modified response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id titles not modified response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetCharactersCharacterIDTitlesNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesNotModified ", 304)
}

func (o *GetCharactersCharacterIDTitlesNotModified) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesNotModified ", 304)
}

func (o *GetCharactersCharacterIDTitlesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDTitlesBadRequest creates a GetCharactersCharacterIDTitlesBadRequest with default headers values
func NewGetCharactersCharacterIDTitlesBadRequest() *GetCharactersCharacterIDTitlesBadRequest {
	return &GetCharactersCharacterIDTitlesBadRequest{}
}

/*
GetCharactersCharacterIDTitlesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDTitlesBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id titles bad request response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles bad request response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles bad request response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id titles bad request response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id titles bad request response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCharactersCharacterIDTitlesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDTitlesUnauthorized creates a GetCharactersCharacterIDTitlesUnauthorized with default headers values
func NewGetCharactersCharacterIDTitlesUnauthorized() *GetCharactersCharacterIDTitlesUnauthorized {
	return &GetCharactersCharacterIDTitlesUnauthorized{}
}

/*
GetCharactersCharacterIDTitlesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDTitlesUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id titles unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id titles unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id titles unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCharactersCharacterIDTitlesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDTitlesForbidden creates a GetCharactersCharacterIDTitlesForbidden with default headers values
func NewGetCharactersCharacterIDTitlesForbidden() *GetCharactersCharacterIDTitlesForbidden {
	return &GetCharactersCharacterIDTitlesForbidden{}
}

/*
GetCharactersCharacterIDTitlesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDTitlesForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id titles forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id titles forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id titles forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCharactersCharacterIDTitlesForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesForbidden) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDTitlesEnhanceYourCalm creates a GetCharactersCharacterIDTitlesEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDTitlesEnhanceYourCalm() *GetCharactersCharacterIDTitlesEnhanceYourCalm {
	return &GetCharactersCharacterIDTitlesEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDTitlesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDTitlesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id titles enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id titles enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id titles enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDTitlesInternalServerError creates a GetCharactersCharacterIDTitlesInternalServerError with default headers values
func NewGetCharactersCharacterIDTitlesInternalServerError() *GetCharactersCharacterIDTitlesInternalServerError {
	return &GetCharactersCharacterIDTitlesInternalServerError{}
}

/*
GetCharactersCharacterIDTitlesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDTitlesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id titles internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id titles internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id titles internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCharactersCharacterIDTitlesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDTitlesServiceUnavailable creates a GetCharactersCharacterIDTitlesServiceUnavailable with default headers values
func NewGetCharactersCharacterIDTitlesServiceUnavailable() *GetCharactersCharacterIDTitlesServiceUnavailable {
	return &GetCharactersCharacterIDTitlesServiceUnavailable{}
}

/*
GetCharactersCharacterIDTitlesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDTitlesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id titles service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id titles service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id titles service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCharactersCharacterIDTitlesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDTitlesGatewayTimeout creates a GetCharactersCharacterIDTitlesGatewayTimeout with default headers values
func NewGetCharactersCharacterIDTitlesGatewayTimeout() *GetCharactersCharacterIDTitlesGatewayTimeout {
	return &GetCharactersCharacterIDTitlesGatewayTimeout{}
}

/*
GetCharactersCharacterIDTitlesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDTitlesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id titles gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDTitlesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id titles gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDTitlesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id titles gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDTitlesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id titles gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDTitlesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id titles gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDTitlesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCharactersCharacterIDTitlesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/titles/][%d] getCharactersCharacterIdTitlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDTitlesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDTitlesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCharactersCharacterIDTitlesOKBodyItems0 get_characters_character_id_titles_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDTitlesOKBodyItems0
*/
type GetCharactersCharacterIDTitlesOKBodyItems0 struct {

	// get_characters_character_id_titles_name
	//
	// name string
	Name string `json:"name,omitempty"`

	// get_characters_character_id_titles_title_id
	//
	// title_id integer
	TitleID int32 `json:"title_id,omitempty"`
}

// Validate validates this get characters character ID titles o k body items0
func (o *GetCharactersCharacterIDTitlesOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get characters character ID titles o k body items0 based on context it is used
func (o *GetCharactersCharacterIDTitlesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDTitlesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDTitlesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDTitlesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
