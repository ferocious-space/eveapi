// Code generated by go-swagger; DO NOT EDIT.

package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/models"
)

// GetCorporationsCorporationIDMembersLimitReader is a Reader for the GetCorporationsCorporationIDMembersLimit structure.
type GetCorporationsCorporationIDMembersLimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMembersLimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDMembersLimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDMembersLimitNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDMembersLimitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDMembersLimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDMembersLimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDMembersLimitEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDMembersLimitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDMembersLimitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDMembersLimitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/corporations/{corporation_id}/members/limit/] get_corporations_corporation_id_members_limit", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMembersLimitOK creates a GetCorporationsCorporationIDMembersLimitOK with default headers values
func NewGetCorporationsCorporationIDMembersLimitOK() *GetCorporationsCorporationIDMembersLimitOK {
	return &GetCorporationsCorporationIDMembersLimitOK{}
}

/*
GetCorporationsCorporationIDMembersLimitOK describes a response with status code 200, with default header values.

The corporation's member limit
*/
type GetCorporationsCorporationIDMembersLimitOK struct {

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

	Payload int32
}

// IsSuccess returns true when this get corporations corporation Id members limit o k response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get corporations corporation Id members limit o k response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit o k response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members limit o k response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members limit o k response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get corporations corporation Id members limit o k response
func (o *GetCorporationsCorporationIDMembersLimitOK) Code() int {
	return 200
}

func (o *GetCorporationsCorporationIDMembersLimitOK) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitOK) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitOK) GetPayload() int32 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersLimitNotModified creates a GetCorporationsCorporationIDMembersLimitNotModified with default headers values
func NewGetCorporationsCorporationIDMembersLimitNotModified() *GetCorporationsCorporationIDMembersLimitNotModified {
	return &GetCorporationsCorporationIDMembersLimitNotModified{}
}

/*
GetCorporationsCorporationIDMembersLimitNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDMembersLimitNotModified struct {

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

// IsSuccess returns true when this get corporations corporation Id members limit not modified response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit not modified response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get corporations corporation Id members limit not modified response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members limit not modified response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members limit not modified response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get corporations corporation Id members limit not modified response
func (o *GetCorporationsCorporationIDMembersLimitNotModified) Code() int {
	return 304
}

func (o *GetCorporationsCorporationIDMembersLimitNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembersLimitNotModified) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembersLimitNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersLimitBadRequest creates a GetCorporationsCorporationIDMembersLimitBadRequest with default headers values
func NewGetCorporationsCorporationIDMembersLimitBadRequest() *GetCorporationsCorporationIDMembersLimitBadRequest {
	return &GetCorporationsCorporationIDMembersLimitBadRequest{}
}

/*
GetCorporationsCorporationIDMembersLimitBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDMembersLimitBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get corporations corporation Id members limit bad request response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit bad request response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit bad request response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members limit bad request response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members limit bad request response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get corporations corporation Id members limit bad request response
func (o *GetCorporationsCorporationIDMembersLimitBadRequest) Code() int {
	return 400
}

func (o *GetCorporationsCorporationIDMembersLimitBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersLimitUnauthorized creates a GetCorporationsCorporationIDMembersLimitUnauthorized with default headers values
func NewGetCorporationsCorporationIDMembersLimitUnauthorized() *GetCorporationsCorporationIDMembersLimitUnauthorized {
	return &GetCorporationsCorporationIDMembersLimitUnauthorized{}
}

/*
GetCorporationsCorporationIDMembersLimitUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMembersLimitUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get corporations corporation Id members limit unauthorized response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit unauthorized response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit unauthorized response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members limit unauthorized response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members limit unauthorized response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get corporations corporation Id members limit unauthorized response
func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) Code() int {
	return 401
}

func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersLimitForbidden creates a GetCorporationsCorporationIDMembersLimitForbidden with default headers values
func NewGetCorporationsCorporationIDMembersLimitForbidden() *GetCorporationsCorporationIDMembersLimitForbidden {
	return &GetCorporationsCorporationIDMembersLimitForbidden{}
}

/*
GetCorporationsCorporationIDMembersLimitForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMembersLimitForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get corporations corporation Id members limit forbidden response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit forbidden response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit forbidden response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members limit forbidden response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members limit forbidden response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get corporations corporation Id members limit forbidden response
func (o *GetCorporationsCorporationIDMembersLimitForbidden) Code() int {
	return 403
}

func (o *GetCorporationsCorporationIDMembersLimitForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitForbidden) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersLimitEnhanceYourCalm creates a GetCorporationsCorporationIDMembersLimitEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMembersLimitEnhanceYourCalm() *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm {
	return &GetCorporationsCorporationIDMembersLimitEnhanceYourCalm{}
}

/*
GetCorporationsCorporationIDMembersLimitEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDMembersLimitEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get corporations corporation Id members limit enhance your calm response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit enhance your calm response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit enhance your calm response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members limit enhance your calm response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members limit enhance your calm response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get corporations corporation Id members limit enhance your calm response
func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersLimitInternalServerError creates a GetCorporationsCorporationIDMembersLimitInternalServerError with default headers values
func NewGetCorporationsCorporationIDMembersLimitInternalServerError() *GetCorporationsCorporationIDMembersLimitInternalServerError {
	return &GetCorporationsCorporationIDMembersLimitInternalServerError{}
}

/*
GetCorporationsCorporationIDMembersLimitInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMembersLimitInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get corporations corporation Id members limit internal server error response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit internal server error response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit internal server error response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members limit internal server error response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members limit internal server error response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get corporations corporation Id members limit internal server error response
func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) Code() int {
	return 500
}

func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersLimitServiceUnavailable creates a GetCorporationsCorporationIDMembersLimitServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMembersLimitServiceUnavailable() *GetCorporationsCorporationIDMembersLimitServiceUnavailable {
	return &GetCorporationsCorporationIDMembersLimitServiceUnavailable{}
}

/*
GetCorporationsCorporationIDMembersLimitServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMembersLimitServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get corporations corporation Id members limit service unavailable response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit service unavailable response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit service unavailable response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members limit service unavailable response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members limit service unavailable response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get corporations corporation Id members limit service unavailable response
func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) Code() int {
	return 503
}

func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersLimitGatewayTimeout creates a GetCorporationsCorporationIDMembersLimitGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMembersLimitGatewayTimeout() *GetCorporationsCorporationIDMembersLimitGatewayTimeout {
	return &GetCorporationsCorporationIDMembersLimitGatewayTimeout{}
}

/*
GetCorporationsCorporationIDMembersLimitGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMembersLimitGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get corporations corporation Id members limit gateway timeout response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members limit gateway timeout response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members limit gateway timeout response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members limit gateway timeout response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members limit gateway timeout response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get corporations corporation Id members limit gateway timeout response
func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) Code() int {
	return 504
}

func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/limit/][%d] getCorporationsCorporationIdMembersLimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersLimitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
