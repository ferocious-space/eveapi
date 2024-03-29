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

// GetFleetsFleetIDMembersReader is a Reader for the GetFleetsFleetIDMembers structure.
type GetFleetsFleetIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetsFleetIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFleetsFleetIDMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetFleetsFleetIDMembersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetFleetsFleetIDMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFleetsFleetIDMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFleetsFleetIDMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFleetsFleetIDMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetFleetsFleetIDMembersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFleetsFleetIDMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFleetsFleetIDMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFleetsFleetIDMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/fleets/{fleet_id}/members/] get_fleets_fleet_id_members", response, response.Code())
	}
}

// NewGetFleetsFleetIDMembersOK creates a GetFleetsFleetIDMembersOK with default headers values
func NewGetFleetsFleetIDMembersOK() *GetFleetsFleetIDMembersOK {
	return &GetFleetsFleetIDMembersOK{}
}

/*
GetFleetsFleetIDMembersOK describes a response with status code 200, with default header values.

A list of fleet members
*/
type GetFleetsFleetIDMembersOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* The language used in the response
	 */
	ContentLanguage string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetFleetsFleetIDMembersOKBodyItems0
}

// IsSuccess returns true when this get fleets fleet Id members o k response has a 2xx status code
func (o *GetFleetsFleetIDMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fleets fleet Id members o k response has a 3xx status code
func (o *GetFleetsFleetIDMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members o k response has a 4xx status code
func (o *GetFleetsFleetIDMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id members o k response has a 5xx status code
func (o *GetFleetsFleetIDMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id members o k response a status code equal to that given
func (o *GetFleetsFleetIDMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get fleets fleet Id members o k response
func (o *GetFleetsFleetIDMembersOK) Code() int {
	return 200
}

func (o *GetFleetsFleetIDMembersOK) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersOK  %+v", 200, o.Payload)
}

func (o *GetFleetsFleetIDMembersOK) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersOK  %+v", 200, o.Payload)
}

func (o *GetFleetsFleetIDMembersOK) GetPayload() []*GetFleetsFleetIDMembersOKBodyItems0 {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Content-Language
	hdrContentLanguage := response.GetHeader("Content-Language")

	if hdrContentLanguage != "" {
		o.ContentLanguage = hdrContentLanguage
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

// NewGetFleetsFleetIDMembersNotModified creates a GetFleetsFleetIDMembersNotModified with default headers values
func NewGetFleetsFleetIDMembersNotModified() *GetFleetsFleetIDMembersNotModified {
	return &GetFleetsFleetIDMembersNotModified{}
}

/*
GetFleetsFleetIDMembersNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetFleetsFleetIDMembersNotModified struct {

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

// IsSuccess returns true when this get fleets fleet Id members not modified response has a 2xx status code
func (o *GetFleetsFleetIDMembersNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members not modified response has a 3xx status code
func (o *GetFleetsFleetIDMembersNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get fleets fleet Id members not modified response has a 4xx status code
func (o *GetFleetsFleetIDMembersNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id members not modified response has a 5xx status code
func (o *GetFleetsFleetIDMembersNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id members not modified response a status code equal to that given
func (o *GetFleetsFleetIDMembersNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get fleets fleet Id members not modified response
func (o *GetFleetsFleetIDMembersNotModified) Code() int {
	return 304
}

func (o *GetFleetsFleetIDMembersNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersNotModified ", 304)
}

func (o *GetFleetsFleetIDMembersNotModified) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersNotModified ", 304)
}

func (o *GetFleetsFleetIDMembersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFleetsFleetIDMembersBadRequest creates a GetFleetsFleetIDMembersBadRequest with default headers values
func NewGetFleetsFleetIDMembersBadRequest() *GetFleetsFleetIDMembersBadRequest {
	return &GetFleetsFleetIDMembersBadRequest{}
}

/*
GetFleetsFleetIDMembersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFleetsFleetIDMembersBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get fleets fleet Id members bad request response has a 2xx status code
func (o *GetFleetsFleetIDMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members bad request response has a 3xx status code
func (o *GetFleetsFleetIDMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members bad request response has a 4xx status code
func (o *GetFleetsFleetIDMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id members bad request response has a 5xx status code
func (o *GetFleetsFleetIDMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id members bad request response a status code equal to that given
func (o *GetFleetsFleetIDMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get fleets fleet Id members bad request response
func (o *GetFleetsFleetIDMembersBadRequest) Code() int {
	return 400
}

func (o *GetFleetsFleetIDMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetsFleetIDMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetsFleetIDMembersBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDMembersUnauthorized creates a GetFleetsFleetIDMembersUnauthorized with default headers values
func NewGetFleetsFleetIDMembersUnauthorized() *GetFleetsFleetIDMembersUnauthorized {
	return &GetFleetsFleetIDMembersUnauthorized{}
}

/*
GetFleetsFleetIDMembersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFleetsFleetIDMembersUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get fleets fleet Id members unauthorized response has a 2xx status code
func (o *GetFleetsFleetIDMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members unauthorized response has a 3xx status code
func (o *GetFleetsFleetIDMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members unauthorized response has a 4xx status code
func (o *GetFleetsFleetIDMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id members unauthorized response has a 5xx status code
func (o *GetFleetsFleetIDMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id members unauthorized response a status code equal to that given
func (o *GetFleetsFleetIDMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get fleets fleet Id members unauthorized response
func (o *GetFleetsFleetIDMembersUnauthorized) Code() int {
	return 401
}

func (o *GetFleetsFleetIDMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFleetsFleetIDMembersUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFleetsFleetIDMembersUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDMembersForbidden creates a GetFleetsFleetIDMembersForbidden with default headers values
func NewGetFleetsFleetIDMembersForbidden() *GetFleetsFleetIDMembersForbidden {
	return &GetFleetsFleetIDMembersForbidden{}
}

/*
GetFleetsFleetIDMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFleetsFleetIDMembersForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get fleets fleet Id members forbidden response has a 2xx status code
func (o *GetFleetsFleetIDMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members forbidden response has a 3xx status code
func (o *GetFleetsFleetIDMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members forbidden response has a 4xx status code
func (o *GetFleetsFleetIDMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id members forbidden response has a 5xx status code
func (o *GetFleetsFleetIDMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id members forbidden response a status code equal to that given
func (o *GetFleetsFleetIDMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get fleets fleet Id members forbidden response
func (o *GetFleetsFleetIDMembersForbidden) Code() int {
	return 403
}

func (o *GetFleetsFleetIDMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetFleetsFleetIDMembersForbidden) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetFleetsFleetIDMembersForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDMembersNotFound creates a GetFleetsFleetIDMembersNotFound with default headers values
func NewGetFleetsFleetIDMembersNotFound() *GetFleetsFleetIDMembersNotFound {
	return &GetFleetsFleetIDMembersNotFound{}
}

/*
GetFleetsFleetIDMembersNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type GetFleetsFleetIDMembersNotFound struct {
	Payload *GetFleetsFleetIDMembersNotFoundBody
}

// IsSuccess returns true when this get fleets fleet Id members not found response has a 2xx status code
func (o *GetFleetsFleetIDMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members not found response has a 3xx status code
func (o *GetFleetsFleetIDMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members not found response has a 4xx status code
func (o *GetFleetsFleetIDMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id members not found response has a 5xx status code
func (o *GetFleetsFleetIDMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id members not found response a status code equal to that given
func (o *GetFleetsFleetIDMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get fleets fleet Id members not found response
func (o *GetFleetsFleetIDMembersNotFound) Code() int {
	return 404
}

func (o *GetFleetsFleetIDMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetFleetsFleetIDMembersNotFound) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetFleetsFleetIDMembersNotFound) GetPayload() *GetFleetsFleetIDMembersNotFoundBody {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFleetsFleetIDMembersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDMembersEnhanceYourCalm creates a GetFleetsFleetIDMembersEnhanceYourCalm with default headers values
func NewGetFleetsFleetIDMembersEnhanceYourCalm() *GetFleetsFleetIDMembersEnhanceYourCalm {
	return &GetFleetsFleetIDMembersEnhanceYourCalm{}
}

/*
GetFleetsFleetIDMembersEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetFleetsFleetIDMembersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get fleets fleet Id members enhance your calm response has a 2xx status code
func (o *GetFleetsFleetIDMembersEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members enhance your calm response has a 3xx status code
func (o *GetFleetsFleetIDMembersEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members enhance your calm response has a 4xx status code
func (o *GetFleetsFleetIDMembersEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id members enhance your calm response has a 5xx status code
func (o *GetFleetsFleetIDMembersEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id members enhance your calm response a status code equal to that given
func (o *GetFleetsFleetIDMembersEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get fleets fleet Id members enhance your calm response
func (o *GetFleetsFleetIDMembersEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetFleetsFleetIDMembersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetFleetsFleetIDMembersEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetFleetsFleetIDMembersEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDMembersInternalServerError creates a GetFleetsFleetIDMembersInternalServerError with default headers values
func NewGetFleetsFleetIDMembersInternalServerError() *GetFleetsFleetIDMembersInternalServerError {
	return &GetFleetsFleetIDMembersInternalServerError{}
}

/*
GetFleetsFleetIDMembersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetFleetsFleetIDMembersInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get fleets fleet Id members internal server error response has a 2xx status code
func (o *GetFleetsFleetIDMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members internal server error response has a 3xx status code
func (o *GetFleetsFleetIDMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members internal server error response has a 4xx status code
func (o *GetFleetsFleetIDMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id members internal server error response has a 5xx status code
func (o *GetFleetsFleetIDMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fleets fleet Id members internal server error response a status code equal to that given
func (o *GetFleetsFleetIDMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get fleets fleet Id members internal server error response
func (o *GetFleetsFleetIDMembersInternalServerError) Code() int {
	return 500
}

func (o *GetFleetsFleetIDMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFleetsFleetIDMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFleetsFleetIDMembersInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDMembersServiceUnavailable creates a GetFleetsFleetIDMembersServiceUnavailable with default headers values
func NewGetFleetsFleetIDMembersServiceUnavailable() *GetFleetsFleetIDMembersServiceUnavailable {
	return &GetFleetsFleetIDMembersServiceUnavailable{}
}

/*
GetFleetsFleetIDMembersServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetFleetsFleetIDMembersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get fleets fleet Id members service unavailable response has a 2xx status code
func (o *GetFleetsFleetIDMembersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members service unavailable response has a 3xx status code
func (o *GetFleetsFleetIDMembersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members service unavailable response has a 4xx status code
func (o *GetFleetsFleetIDMembersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id members service unavailable response has a 5xx status code
func (o *GetFleetsFleetIDMembersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get fleets fleet Id members service unavailable response a status code equal to that given
func (o *GetFleetsFleetIDMembersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get fleets fleet Id members service unavailable response
func (o *GetFleetsFleetIDMembersServiceUnavailable) Code() int {
	return 503
}

func (o *GetFleetsFleetIDMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFleetsFleetIDMembersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFleetsFleetIDMembersServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDMembersGatewayTimeout creates a GetFleetsFleetIDMembersGatewayTimeout with default headers values
func NewGetFleetsFleetIDMembersGatewayTimeout() *GetFleetsFleetIDMembersGatewayTimeout {
	return &GetFleetsFleetIDMembersGatewayTimeout{}
}

/*
GetFleetsFleetIDMembersGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetFleetsFleetIDMembersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get fleets fleet Id members gateway timeout response has a 2xx status code
func (o *GetFleetsFleetIDMembersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id members gateway timeout response has a 3xx status code
func (o *GetFleetsFleetIDMembersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id members gateway timeout response has a 4xx status code
func (o *GetFleetsFleetIDMembersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id members gateway timeout response has a 5xx status code
func (o *GetFleetsFleetIDMembersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get fleets fleet Id members gateway timeout response a status code equal to that given
func (o *GetFleetsFleetIDMembersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get fleets fleet Id members gateway timeout response
func (o *GetFleetsFleetIDMembersGatewayTimeout) Code() int {
	return 504
}

func (o *GetFleetsFleetIDMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFleetsFleetIDMembersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/members/][%d] getFleetsFleetIdMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFleetsFleetIDMembersGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetFleetsFleetIDMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetFleetsFleetIDMembersNotFoundBody get_fleets_fleet_id_members_not_found
//
// Not found
swagger:model GetFleetsFleetIDMembersNotFoundBody
*/
type GetFleetsFleetIDMembersNotFoundBody struct {

	// get_fleets_fleet_id_members_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get fleets fleet ID members not found body
func (o *GetFleetsFleetIDMembersNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get fleets fleet ID members not found body based on context it is used
func (o *GetFleetsFleetIDMembersNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDMembersNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDMembersNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDMembersNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetFleetsFleetIDMembersOKBodyItems0 get_fleets_fleet_id_members_200_ok
//
// 200 ok object
swagger:model GetFleetsFleetIDMembersOKBodyItems0
*/
type GetFleetsFleetIDMembersOKBodyItems0 struct {

	// get_fleets_fleet_id_members_character_id
	//
	// character_id integer
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_fleets_fleet_id_members_join_time
	//
	// join_time string
	// Required: true
	// Format: date-time
	JoinTime *strfmt.DateTime `json:"join_time"`

	// get_fleets_fleet_id_members_role
	//
	// Member’s role in fleet
	// Required: true
	// Enum: [fleet_commander wing_commander squad_commander squad_member]
	Role *string `json:"role"`

	// get_fleets_fleet_id_members_role_name
	//
	// Localized role names
	// Required: true
	RoleName *string `json:"role_name"`

	// get_fleets_fleet_id_members_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`

	// get_fleets_fleet_id_members_solar_system_id
	//
	// Solar system the member is located in
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_fleets_fleet_id_members_squad_id
	//
	// ID of the squad the member is in. If not applicable, will be set to -1
	// Required: true
	SquadID *int64 `json:"squad_id"`

	// get_fleets_fleet_id_members_station_id
	//
	// Station in which the member is docked in, if applicable
	StationID int64 `json:"station_id,omitempty"`

	// get_fleets_fleet_id_members_takes_fleet_warp
	//
	// Whether the member take fleet warps
	// Required: true
	TakesFleetWarp *bool `json:"takes_fleet_warp"`

	// get_fleets_fleet_id_members_wing_id
	//
	// ID of the wing the member is in. If not applicable, will be set to -1
	// Required: true
	WingID *int64 `json:"wing_id"`
}

// Validate validates this get fleets fleet ID members o k body items0
func (o *GetFleetsFleetIDMembersOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateJoinTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoleName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSquadID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTakesFleetWarp(formats); err != nil {
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

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateJoinTime(formats strfmt.Registry) error {

	if err := validate.Required("join_time", "body", o.JoinTime); err != nil {
		return err
	}

	if err := validate.FormatOf("join_time", "body", "date-time", o.JoinTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var getFleetsFleetIdMembersOKBodyItems0TypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fleet_commander","wing_commander","squad_commander","squad_member"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getFleetsFleetIdMembersOKBodyItems0TypeRolePropEnum = append(getFleetsFleetIdMembersOKBodyItems0TypeRolePropEnum, v)
	}
}

const (

	// GetFleetsFleetIDMembersOKBodyItems0RoleFleetCommander captures enum value "fleet_commander"
	GetFleetsFleetIDMembersOKBodyItems0RoleFleetCommander string = "fleet_commander"

	// GetFleetsFleetIDMembersOKBodyItems0RoleWingCommander captures enum value "wing_commander"
	GetFleetsFleetIDMembersOKBodyItems0RoleWingCommander string = "wing_commander"

	// GetFleetsFleetIDMembersOKBodyItems0RoleSquadCommander captures enum value "squad_commander"
	GetFleetsFleetIDMembersOKBodyItems0RoleSquadCommander string = "squad_commander"

	// GetFleetsFleetIDMembersOKBodyItems0RoleSquadMember captures enum value "squad_member"
	GetFleetsFleetIDMembersOKBodyItems0RoleSquadMember string = "squad_member"
)

// prop value enum
func (o *GetFleetsFleetIDMembersOKBodyItems0) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getFleetsFleetIdMembersOKBodyItems0TypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", o.Role); err != nil {
		return err
	}

	// value enum
	if err := o.validateRoleEnum("role", "body", *o.Role); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateRoleName(formats strfmt.Registry) error {

	if err := validate.Required("role_name", "body", o.RoleName); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("ship_type_id", "body", o.ShipTypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateSquadID(formats strfmt.Registry) error {

	if err := validate.Required("squad_id", "body", o.SquadID); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateTakesFleetWarp(formats strfmt.Registry) error {

	if err := validate.Required("takes_fleet_warp", "body", o.TakesFleetWarp); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDMembersOKBodyItems0) validateWingID(formats strfmt.Registry) error {

	if err := validate.Required("wing_id", "body", o.WingID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get fleets fleet ID members o k body items0 based on context it is used
func (o *GetFleetsFleetIDMembersOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDMembersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDMembersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDMembersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
