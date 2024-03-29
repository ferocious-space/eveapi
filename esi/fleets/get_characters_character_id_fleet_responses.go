// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// GetCharactersCharacterIDFleetReader is a Reader for the GetCharactersCharacterIDFleet structure.
type GetCharactersCharacterIDFleetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDFleetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDFleetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDFleetNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDFleetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDFleetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDFleetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCharactersCharacterIDFleetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDFleetEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDFleetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDFleetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDFleetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/characters/{character_id}/fleet/] get_characters_character_id_fleet", response, response.Code())
	}
}

// NewGetCharactersCharacterIDFleetOK creates a GetCharactersCharacterIDFleetOK with default headers values
func NewGetCharactersCharacterIDFleetOK() *GetCharactersCharacterIDFleetOK {
	return &GetCharactersCharacterIDFleetOK{}
}

/*
GetCharactersCharacterIDFleetOK describes a response with status code 200, with default header values.

Details about the character's fleet
*/
type GetCharactersCharacterIDFleetOK struct {

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

	Payload *GetCharactersCharacterIDFleetOKBody
}

// IsSuccess returns true when this get characters character Id fleet o k response has a 2xx status code
func (o *GetCharactersCharacterIDFleetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id fleet o k response has a 3xx status code
func (o *GetCharactersCharacterIDFleetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet o k response has a 4xx status code
func (o *GetCharactersCharacterIDFleetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id fleet o k response has a 5xx status code
func (o *GetCharactersCharacterIDFleetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id fleet o k response a status code equal to that given
func (o *GetCharactersCharacterIDFleetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get characters character Id fleet o k response
func (o *GetCharactersCharacterIDFleetOK) Code() int {
	return 200
}

func (o *GetCharactersCharacterIDFleetOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDFleetOK) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDFleetOK) GetPayload() *GetCharactersCharacterIDFleetOKBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCharactersCharacterIDFleetOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetNotModified creates a GetCharactersCharacterIDFleetNotModified with default headers values
func NewGetCharactersCharacterIDFleetNotModified() *GetCharactersCharacterIDFleetNotModified {
	return &GetCharactersCharacterIDFleetNotModified{}
}

/*
GetCharactersCharacterIDFleetNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDFleetNotModified struct {

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

// IsSuccess returns true when this get characters character Id fleet not modified response has a 2xx status code
func (o *GetCharactersCharacterIDFleetNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet not modified response has a 3xx status code
func (o *GetCharactersCharacterIDFleetNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id fleet not modified response has a 4xx status code
func (o *GetCharactersCharacterIDFleetNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id fleet not modified response has a 5xx status code
func (o *GetCharactersCharacterIDFleetNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id fleet not modified response a status code equal to that given
func (o *GetCharactersCharacterIDFleetNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get characters character Id fleet not modified response
func (o *GetCharactersCharacterIDFleetNotModified) Code() int {
	return 304
}

func (o *GetCharactersCharacterIDFleetNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetNotModified ", 304)
}

func (o *GetCharactersCharacterIDFleetNotModified) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetNotModified ", 304)
}

func (o *GetCharactersCharacterIDFleetNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDFleetBadRequest creates a GetCharactersCharacterIDFleetBadRequest with default headers values
func NewGetCharactersCharacterIDFleetBadRequest() *GetCharactersCharacterIDFleetBadRequest {
	return &GetCharactersCharacterIDFleetBadRequest{}
}

/*
GetCharactersCharacterIDFleetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDFleetBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id fleet bad request response has a 2xx status code
func (o *GetCharactersCharacterIDFleetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet bad request response has a 3xx status code
func (o *GetCharactersCharacterIDFleetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet bad request response has a 4xx status code
func (o *GetCharactersCharacterIDFleetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id fleet bad request response has a 5xx status code
func (o *GetCharactersCharacterIDFleetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id fleet bad request response a status code equal to that given
func (o *GetCharactersCharacterIDFleetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get characters character Id fleet bad request response
func (o *GetCharactersCharacterIDFleetBadRequest) Code() int {
	return 400
}

func (o *GetCharactersCharacterIDFleetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDFleetBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDFleetBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetUnauthorized creates a GetCharactersCharacterIDFleetUnauthorized with default headers values
func NewGetCharactersCharacterIDFleetUnauthorized() *GetCharactersCharacterIDFleetUnauthorized {
	return &GetCharactersCharacterIDFleetUnauthorized{}
}

/*
GetCharactersCharacterIDFleetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDFleetUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id fleet unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDFleetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDFleetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDFleetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id fleet unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDFleetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id fleet unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDFleetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get characters character Id fleet unauthorized response
func (o *GetCharactersCharacterIDFleetUnauthorized) Code() int {
	return 401
}

func (o *GetCharactersCharacterIDFleetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDFleetUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDFleetUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetForbidden creates a GetCharactersCharacterIDFleetForbidden with default headers values
func NewGetCharactersCharacterIDFleetForbidden() *GetCharactersCharacterIDFleetForbidden {
	return &GetCharactersCharacterIDFleetForbidden{}
}

/*
GetCharactersCharacterIDFleetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDFleetForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id fleet forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDFleetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDFleetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDFleetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id fleet forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDFleetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id fleet forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDFleetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get characters character Id fleet forbidden response
func (o *GetCharactersCharacterIDFleetForbidden) Code() int {
	return 403
}

func (o *GetCharactersCharacterIDFleetForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDFleetForbidden) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDFleetForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetNotFound creates a GetCharactersCharacterIDFleetNotFound with default headers values
func NewGetCharactersCharacterIDFleetNotFound() *GetCharactersCharacterIDFleetNotFound {
	return &GetCharactersCharacterIDFleetNotFound{}
}

/*
GetCharactersCharacterIDFleetNotFound describes a response with status code 404, with default header values.

The character is not in a fleet
*/
type GetCharactersCharacterIDFleetNotFound struct {
	Payload *GetCharactersCharacterIDFleetNotFoundBody
}

// IsSuccess returns true when this get characters character Id fleet not found response has a 2xx status code
func (o *GetCharactersCharacterIDFleetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet not found response has a 3xx status code
func (o *GetCharactersCharacterIDFleetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet not found response has a 4xx status code
func (o *GetCharactersCharacterIDFleetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id fleet not found response has a 5xx status code
func (o *GetCharactersCharacterIDFleetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id fleet not found response a status code equal to that given
func (o *GetCharactersCharacterIDFleetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get characters character Id fleet not found response
func (o *GetCharactersCharacterIDFleetNotFound) Code() int {
	return 404
}

func (o *GetCharactersCharacterIDFleetNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDFleetNotFound) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDFleetNotFound) GetPayload() *GetCharactersCharacterIDFleetNotFoundBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCharactersCharacterIDFleetNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetEnhanceYourCalm creates a GetCharactersCharacterIDFleetEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDFleetEnhanceYourCalm() *GetCharactersCharacterIDFleetEnhanceYourCalm {
	return &GetCharactersCharacterIDFleetEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDFleetEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDFleetEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id fleet enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id fleet enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id fleet enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get characters character Id fleet enhance your calm response
func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetInternalServerError creates a GetCharactersCharacterIDFleetInternalServerError with default headers values
func NewGetCharactersCharacterIDFleetInternalServerError() *GetCharactersCharacterIDFleetInternalServerError {
	return &GetCharactersCharacterIDFleetInternalServerError{}
}

/*
GetCharactersCharacterIDFleetInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDFleetInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id fleet internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDFleetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDFleetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDFleetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id fleet internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDFleetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id fleet internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDFleetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get characters character Id fleet internal server error response
func (o *GetCharactersCharacterIDFleetInternalServerError) Code() int {
	return 500
}

func (o *GetCharactersCharacterIDFleetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDFleetInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDFleetInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetServiceUnavailable creates a GetCharactersCharacterIDFleetServiceUnavailable with default headers values
func NewGetCharactersCharacterIDFleetServiceUnavailable() *GetCharactersCharacterIDFleetServiceUnavailable {
	return &GetCharactersCharacterIDFleetServiceUnavailable{}
}

/*
GetCharactersCharacterIDFleetServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDFleetServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id fleet service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDFleetServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDFleetServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDFleetServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id fleet service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDFleetServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id fleet service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDFleetServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get characters character Id fleet service unavailable response
func (o *GetCharactersCharacterIDFleetServiceUnavailable) Code() int {
	return 503
}

func (o *GetCharactersCharacterIDFleetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDFleetServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDFleetServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDFleetGatewayTimeout creates a GetCharactersCharacterIDFleetGatewayTimeout with default headers values
func NewGetCharactersCharacterIDFleetGatewayTimeout() *GetCharactersCharacterIDFleetGatewayTimeout {
	return &GetCharactersCharacterIDFleetGatewayTimeout{}
}

/*
GetCharactersCharacterIDFleetGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDFleetGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id fleet gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDFleetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id fleet gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDFleetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id fleet gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDFleetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id fleet gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDFleetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id fleet gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDFleetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get characters character Id fleet gateway timeout response
func (o *GetCharactersCharacterIDFleetGatewayTimeout) Code() int {
	return 504
}

func (o *GetCharactersCharacterIDFleetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDFleetGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/fleet/][%d] getCharactersCharacterIdFleetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDFleetGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDFleetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCharactersCharacterIDFleetNotFoundBody get_characters_character_id_fleet_not_found
//
// Not found
swagger:model GetCharactersCharacterIDFleetNotFoundBody
*/
type GetCharactersCharacterIDFleetNotFoundBody struct {

	// get_characters_character_id_fleet_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character ID fleet not found body
func (o *GetCharactersCharacterIDFleetNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get characters character ID fleet not found body based on context it is used
func (o *GetCharactersCharacterIDFleetNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFleetNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetCharactersCharacterIDFleetOKBody get_characters_character_id_fleet_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDFleetOKBody
*/
type GetCharactersCharacterIDFleetOKBody struct {

	// get_characters_character_id_fleet_fleet_id
	//
	// The character's current fleet ID
	// Required: true
	FleetID *int64 `json:"fleet_id"`

	// get_characters_character_id_fleet_role
	//
	// Member’s role in fleet
	// Required: true
	// Enum: [fleet_commander squad_commander squad_member wing_commander]
	Role *string `json:"role"`

	// get_characters_character_id_fleet_squad_id
	//
	// ID of the squad the member is in. If not applicable, will be set to -1
	// Required: true
	SquadID *int64 `json:"squad_id"`

	// get_characters_character_id_fleet_wing_id
	//
	// ID of the wing the member is in. If not applicable, will be set to -1
	// Required: true
	WingID *int64 `json:"wing_id"`
}

// Validate validates this get characters character ID fleet o k body
func (o *GetCharactersCharacterIDFleetOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFleetID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSquadID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateFleetID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"fleet_id", "body", o.FleetID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdFleetOKBodyTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fleet_commander","squad_commander","squad_member","wing_commander"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdFleetOKBodyTypeRolePropEnum = append(getCharactersCharacterIdFleetOKBodyTypeRolePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDFleetOKBodyRoleFleetCommander captures enum value "fleet_commander"
	GetCharactersCharacterIDFleetOKBodyRoleFleetCommander string = "fleet_commander"

	// GetCharactersCharacterIDFleetOKBodyRoleSquadCommander captures enum value "squad_commander"
	GetCharactersCharacterIDFleetOKBodyRoleSquadCommander string = "squad_commander"

	// GetCharactersCharacterIDFleetOKBodyRoleSquadMember captures enum value "squad_member"
	GetCharactersCharacterIDFleetOKBodyRoleSquadMember string = "squad_member"

	// GetCharactersCharacterIDFleetOKBodyRoleWingCommander captures enum value "wing_commander"
	GetCharactersCharacterIDFleetOKBodyRoleWingCommander string = "wing_commander"
)

// prop value enum
func (o *GetCharactersCharacterIDFleetOKBody) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdFleetOKBodyTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	// value enum
	if err := o.validateRoleEnum("getCharactersCharacterIdFleetOK"+"."+"role", "body", *o.Role); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateSquadID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"squad_id", "body", o.SquadID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDFleetOKBody) validateWingID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdFleetOK"+"."+"wing_id", "body", o.WingID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID fleet o k body based on context it is used
func (o *GetCharactersCharacterIDFleetOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDFleetOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFleetOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
