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

// GetCorporationsCorporationIDMembersReader is a Reader for the GetCorporationsCorporationIDMembers structure.
type GetCorporationsCorporationIDMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDMembersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDMembersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMembersOK creates a GetCorporationsCorporationIDMembersOK with default headers values
func NewGetCorporationsCorporationIDMembersOK() *GetCorporationsCorporationIDMembersOK {
	return &GetCorporationsCorporationIDMembersOK{}
}

/*
GetCorporationsCorporationIDMembersOK describes a response with status code 200, with default header values.

List of member character IDs
*/
type GetCorporationsCorporationIDMembersOK struct {

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

	Payload []int32
}

// IsSuccess returns true when this get corporations corporation Id members o k response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get corporations corporation Id members o k response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members o k response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members o k response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members o k response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get corporations corporation Id members o k response
func (o *GetCorporationsCorporationIDMembersOK) Code() int {
	return 200
}

func (o *GetCorporationsCorporationIDMembersOK) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersOK) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersNotModified creates a GetCorporationsCorporationIDMembersNotModified with default headers values
func NewGetCorporationsCorporationIDMembersNotModified() *GetCorporationsCorporationIDMembersNotModified {
	return &GetCorporationsCorporationIDMembersNotModified{}
}

/*
GetCorporationsCorporationIDMembersNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDMembersNotModified struct {

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

// IsSuccess returns true when this get corporations corporation Id members not modified response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members not modified response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get corporations corporation Id members not modified response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members not modified response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members not modified response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get corporations corporation Id members not modified response
func (o *GetCorporationsCorporationIDMembersNotModified) Code() int {
	return 304
}

func (o *GetCorporationsCorporationIDMembersNotModified) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembersNotModified) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersBadRequest creates a GetCorporationsCorporationIDMembersBadRequest with default headers values
func NewGetCorporationsCorporationIDMembersBadRequest() *GetCorporationsCorporationIDMembersBadRequest {
	return &GetCorporationsCorporationIDMembersBadRequest{}
}

/*
GetCorporationsCorporationIDMembersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDMembersBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get corporations corporation Id members bad request response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members bad request response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members bad request response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members bad request response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members bad request response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get corporations corporation Id members bad request response
func (o *GetCorporationsCorporationIDMembersBadRequest) Code() int {
	return 400
}

func (o *GetCorporationsCorporationIDMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersBadRequest) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersUnauthorized creates a GetCorporationsCorporationIDMembersUnauthorized with default headers values
func NewGetCorporationsCorporationIDMembersUnauthorized() *GetCorporationsCorporationIDMembersUnauthorized {
	return &GetCorporationsCorporationIDMembersUnauthorized{}
}

/*
GetCorporationsCorporationIDMembersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMembersUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get corporations corporation Id members unauthorized response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members unauthorized response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members unauthorized response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members unauthorized response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members unauthorized response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get corporations corporation Id members unauthorized response
func (o *GetCorporationsCorporationIDMembersUnauthorized) Code() int {
	return 401
}

func (o *GetCorporationsCorporationIDMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersUnauthorized) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersForbidden creates a GetCorporationsCorporationIDMembersForbidden with default headers values
func NewGetCorporationsCorporationIDMembersForbidden() *GetCorporationsCorporationIDMembersForbidden {
	return &GetCorporationsCorporationIDMembersForbidden{}
}

/*
GetCorporationsCorporationIDMembersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMembersForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get corporations corporation Id members forbidden response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members forbidden response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members forbidden response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members forbidden response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members forbidden response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get corporations corporation Id members forbidden response
func (o *GetCorporationsCorporationIDMembersForbidden) Code() int {
	return 403
}

func (o *GetCorporationsCorporationIDMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersForbidden) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersEnhanceYourCalm creates a GetCorporationsCorporationIDMembersEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMembersEnhanceYourCalm() *GetCorporationsCorporationIDMembersEnhanceYourCalm {
	return &GetCorporationsCorporationIDMembersEnhanceYourCalm{}
}

/*
GetCorporationsCorporationIDMembersEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDMembersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get corporations corporation Id members enhance your calm response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members enhance your calm response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members enhance your calm response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members enhance your calm response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members enhance your calm response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get corporations corporation Id members enhance your calm response
func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersInternalServerError creates a GetCorporationsCorporationIDMembersInternalServerError with default headers values
func NewGetCorporationsCorporationIDMembersInternalServerError() *GetCorporationsCorporationIDMembersInternalServerError {
	return &GetCorporationsCorporationIDMembersInternalServerError{}
}

/*
GetCorporationsCorporationIDMembersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMembersInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get corporations corporation Id members internal server error response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members internal server error response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members internal server error response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members internal server error response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members internal server error response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get corporations corporation Id members internal server error response
func (o *GetCorporationsCorporationIDMembersInternalServerError) Code() int {
	return 500
}

func (o *GetCorporationsCorporationIDMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersServiceUnavailable creates a GetCorporationsCorporationIDMembersServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMembersServiceUnavailable() *GetCorporationsCorporationIDMembersServiceUnavailable {
	return &GetCorporationsCorporationIDMembersServiceUnavailable{}
}

/*
GetCorporationsCorporationIDMembersServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMembersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get corporations corporation Id members service unavailable response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members service unavailable response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members service unavailable response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members service unavailable response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members service unavailable response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get corporations corporation Id members service unavailable response
func (o *GetCorporationsCorporationIDMembersServiceUnavailable) Code() int {
	return 503
}

func (o *GetCorporationsCorporationIDMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersGatewayTimeout creates a GetCorporationsCorporationIDMembersGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMembersGatewayTimeout() *GetCorporationsCorporationIDMembersGatewayTimeout {
	return &GetCorporationsCorporationIDMembersGatewayTimeout{}
}

/*
GetCorporationsCorporationIDMembersGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMembersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get corporations corporation Id members gateway timeout response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members gateway timeout response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members gateway timeout response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members gateway timeout response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members gateway timeout response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get corporations corporation Id members gateway timeout response
func (o *GetCorporationsCorporationIDMembersGatewayTimeout) Code() int {
	return 504
}

func (o *GetCorporationsCorporationIDMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/members/][%d] getCorporationsCorporationIdMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
