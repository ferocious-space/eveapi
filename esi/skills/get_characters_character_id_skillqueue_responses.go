// Code generated by go-swagger; DO NOT EDIT.

package skills

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

// GetCharactersCharacterIDSkillqueueReader is a Reader for the GetCharactersCharacterIDSkillqueue structure.
type GetCharactersCharacterIDSkillqueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDSkillqueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDSkillqueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDSkillqueueNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDSkillqueueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDSkillqueueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDSkillqueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDSkillqueueEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDSkillqueueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDSkillqueueServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDSkillqueueGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/characters/{character_id}/skillqueue/] get_characters_character_id_skillqueue", response, response.Code())
	}
}

// NewGetCharactersCharacterIDSkillqueueOK creates a GetCharactersCharacterIDSkillqueueOK with default headers values
func NewGetCharactersCharacterIDSkillqueueOK() *GetCharactersCharacterIDSkillqueueOK {
	return &GetCharactersCharacterIDSkillqueueOK{}
}

/*
GetCharactersCharacterIDSkillqueueOK describes a response with status code 200, with default header values.

The current skill queue, sorted ascending by finishing time
*/
type GetCharactersCharacterIDSkillqueueOK struct {

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

	Payload []*GetCharactersCharacterIDSkillqueueOKBodyItems0
}

// IsSuccess returns true when this get characters character Id skillqueue o k response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id skillqueue o k response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue o k response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id skillqueue o k response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id skillqueue o k response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get characters character Id skillqueue o k response
func (o *GetCharactersCharacterIDSkillqueueOK) Code() int {
	return 200
}

func (o *GetCharactersCharacterIDSkillqueueOK) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueOK) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueOK) GetPayload() []*GetCharactersCharacterIDSkillqueueOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDSkillqueueNotModified creates a GetCharactersCharacterIDSkillqueueNotModified with default headers values
func NewGetCharactersCharacterIDSkillqueueNotModified() *GetCharactersCharacterIDSkillqueueNotModified {
	return &GetCharactersCharacterIDSkillqueueNotModified{}
}

/*
GetCharactersCharacterIDSkillqueueNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDSkillqueueNotModified struct {

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

// IsSuccess returns true when this get characters character Id skillqueue not modified response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue not modified response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id skillqueue not modified response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id skillqueue not modified response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id skillqueue not modified response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get characters character Id skillqueue not modified response
func (o *GetCharactersCharacterIDSkillqueueNotModified) Code() int {
	return 304
}

func (o *GetCharactersCharacterIDSkillqueueNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueNotModified ", 304)
}

func (o *GetCharactersCharacterIDSkillqueueNotModified) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueNotModified ", 304)
}

func (o *GetCharactersCharacterIDSkillqueueNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDSkillqueueBadRequest creates a GetCharactersCharacterIDSkillqueueBadRequest with default headers values
func NewGetCharactersCharacterIDSkillqueueBadRequest() *GetCharactersCharacterIDSkillqueueBadRequest {
	return &GetCharactersCharacterIDSkillqueueBadRequest{}
}

/*
GetCharactersCharacterIDSkillqueueBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDSkillqueueBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id skillqueue bad request response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue bad request response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue bad request response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id skillqueue bad request response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id skillqueue bad request response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get characters character Id skillqueue bad request response
func (o *GetCharactersCharacterIDSkillqueueBadRequest) Code() int {
	return 400
}

func (o *GetCharactersCharacterIDSkillqueueBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillqueueUnauthorized creates a GetCharactersCharacterIDSkillqueueUnauthorized with default headers values
func NewGetCharactersCharacterIDSkillqueueUnauthorized() *GetCharactersCharacterIDSkillqueueUnauthorized {
	return &GetCharactersCharacterIDSkillqueueUnauthorized{}
}

/*
GetCharactersCharacterIDSkillqueueUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDSkillqueueUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id skillqueue unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id skillqueue unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id skillqueue unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get characters character Id skillqueue unauthorized response
func (o *GetCharactersCharacterIDSkillqueueUnauthorized) Code() int {
	return 401
}

func (o *GetCharactersCharacterIDSkillqueueUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillqueueForbidden creates a GetCharactersCharacterIDSkillqueueForbidden with default headers values
func NewGetCharactersCharacterIDSkillqueueForbidden() *GetCharactersCharacterIDSkillqueueForbidden {
	return &GetCharactersCharacterIDSkillqueueForbidden{}
}

/*
GetCharactersCharacterIDSkillqueueForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDSkillqueueForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id skillqueue forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id skillqueue forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id skillqueue forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get characters character Id skillqueue forbidden response
func (o *GetCharactersCharacterIDSkillqueueForbidden) Code() int {
	return 403
}

func (o *GetCharactersCharacterIDSkillqueueForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueForbidden) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillqueueEnhanceYourCalm creates a GetCharactersCharacterIDSkillqueueEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDSkillqueueEnhanceYourCalm() *GetCharactersCharacterIDSkillqueueEnhanceYourCalm {
	return &GetCharactersCharacterIDSkillqueueEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDSkillqueueEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDSkillqueueEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id skillqueue enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id skillqueue enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id skillqueue enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get characters character Id skillqueue enhance your calm response
func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillqueueInternalServerError creates a GetCharactersCharacterIDSkillqueueInternalServerError with default headers values
func NewGetCharactersCharacterIDSkillqueueInternalServerError() *GetCharactersCharacterIDSkillqueueInternalServerError {
	return &GetCharactersCharacterIDSkillqueueInternalServerError{}
}

/*
GetCharactersCharacterIDSkillqueueInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDSkillqueueInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id skillqueue internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id skillqueue internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id skillqueue internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get characters character Id skillqueue internal server error response
func (o *GetCharactersCharacterIDSkillqueueInternalServerError) Code() int {
	return 500
}

func (o *GetCharactersCharacterIDSkillqueueInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillqueueServiceUnavailable creates a GetCharactersCharacterIDSkillqueueServiceUnavailable with default headers values
func NewGetCharactersCharacterIDSkillqueueServiceUnavailable() *GetCharactersCharacterIDSkillqueueServiceUnavailable {
	return &GetCharactersCharacterIDSkillqueueServiceUnavailable{}
}

/*
GetCharactersCharacterIDSkillqueueServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDSkillqueueServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id skillqueue service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id skillqueue service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id skillqueue service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get characters character Id skillqueue service unavailable response
func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) Code() int {
	return 503
}

func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDSkillqueueGatewayTimeout creates a GetCharactersCharacterIDSkillqueueGatewayTimeout with default headers values
func NewGetCharactersCharacterIDSkillqueueGatewayTimeout() *GetCharactersCharacterIDSkillqueueGatewayTimeout {
	return &GetCharactersCharacterIDSkillqueueGatewayTimeout{}
}

/*
GetCharactersCharacterIDSkillqueueGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDSkillqueueGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id skillqueue gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id skillqueue gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id skillqueue gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id skillqueue gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id skillqueue gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get characters character Id skillqueue gateway timeout response
func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) Code() int {
	return 504
}

func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/skillqueue/][%d] getCharactersCharacterIdSkillqueueGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDSkillqueueGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCharactersCharacterIDSkillqueueOKBodyItems0 get_characters_character_id_skillqueue_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDSkillqueueOKBodyItems0
*/
type GetCharactersCharacterIDSkillqueueOKBodyItems0 struct {

	// get_characters_character_id_skillqueue_finish_date
	//
	// Date on which training of the skill will complete. Omitted if the skill queue is paused.
	// Format: date-time
	FinishDate strfmt.DateTime `json:"finish_date,omitempty"`

	// get_characters_character_id_skillqueue_finished_level
	//
	// finished_level integer
	// Required: true
	// Maximum: 5
	// Minimum: 0
	FinishedLevel *int32 `json:"finished_level"`

	// get_characters_character_id_skillqueue_level_end_sp
	//
	// level_end_sp integer
	LevelEndSp int32 `json:"level_end_sp,omitempty"`

	// get_characters_character_id_skillqueue_level_start_sp
	//
	// Amount of SP that was in the skill when it started training it's current level. Used to calculate % of current level complete.
	LevelStartSp int32 `json:"level_start_sp,omitempty"`

	// get_characters_character_id_skillqueue_queue_position
	//
	// queue_position integer
	// Required: true
	QueuePosition *int32 `json:"queue_position"`

	// get_characters_character_id_skillqueue_skill_id
	//
	// skill_id integer
	// Required: true
	SkillID *int32 `json:"skill_id"`

	// get_characters_character_id_skillqueue_start_date
	//
	// start_date string
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`

	// get_characters_character_id_skillqueue_training_start_sp
	//
	// training_start_sp integer
	TrainingStartSp int32 `json:"training_start_sp,omitempty"`
}

// Validate validates this get characters character ID skillqueue o k body items0
func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFinishDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFinishedLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQueuePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSkillID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) validateFinishDate(formats strfmt.Registry) error {
	if swag.IsZero(o.FinishDate) { // not required
		return nil
	}

	if err := validate.FormatOf("finish_date", "body", "date-time", o.FinishDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) validateFinishedLevel(formats strfmt.Registry) error {

	if err := validate.Required("finished_level", "body", o.FinishedLevel); err != nil {
		return err
	}

	if err := validate.MinimumInt("finished_level", "body", int64(*o.FinishedLevel), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("finished_level", "body", int64(*o.FinishedLevel), 5, false); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) validateQueuePosition(formats strfmt.Registry) error {

	if err := validate.Required("queue_position", "body", o.QueuePosition); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) validateSkillID(formats strfmt.Registry) error {

	if err := validate.Required("skill_id", "body", o.SkillID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(o.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", o.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID skillqueue o k body items0 based on context it is used
func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDSkillqueueOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDSkillqueueOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
