// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// GetCorporationsCorporationIDMembersTitlesReader is a Reader for the GetCorporationsCorporationIDMembersTitles structure.
type GetCorporationsCorporationIDMembersTitlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMembersTitlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDMembersTitlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDMembersTitlesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDMembersTitlesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDMembersTitlesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDMembersTitlesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDMembersTitlesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDMembersTitlesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDMembersTitlesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDMembersTitlesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMembersTitlesOK creates a GetCorporationsCorporationIDMembersTitlesOK with default headers values
func NewGetCorporationsCorporationIDMembersTitlesOK() *GetCorporationsCorporationIDMembersTitlesOK {
	return &GetCorporationsCorporationIDMembersTitlesOK{}
}

/*
GetCorporationsCorporationIDMembersTitlesOK describes a response with status code 200, with default header values.

A list of members and theirs titles
*/
type GetCorporationsCorporationIDMembersTitlesOK struct {

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

	Payload []*GetCorporationsCorporationIDMembersTitlesOKBodyItems0
}

// IsSuccess returns true when this get corporations corporation Id members titles o k response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get corporations corporation Id members titles o k response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles o k response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members titles o k response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members titles o k response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCorporationsCorporationIDMembersTitlesOK) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesOK) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesOK) GetPayload() []*GetCorporationsCorporationIDMembersTitlesOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersTitlesNotModified creates a GetCorporationsCorporationIDMembersTitlesNotModified with default headers values
func NewGetCorporationsCorporationIDMembersTitlesNotModified() *GetCorporationsCorporationIDMembersTitlesNotModified {
	return &GetCorporationsCorporationIDMembersTitlesNotModified{}
}

/*
GetCorporationsCorporationIDMembersTitlesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDMembersTitlesNotModified struct {

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

// IsSuccess returns true when this get corporations corporation Id members titles not modified response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles not modified response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get corporations corporation Id members titles not modified response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members titles not modified response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members titles not modified response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetCorporationsCorporationIDMembersTitlesNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembersTitlesNotModified) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembersTitlesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembersTitlesBadRequest creates a GetCorporationsCorporationIDMembersTitlesBadRequest with default headers values
func NewGetCorporationsCorporationIDMembersTitlesBadRequest() *GetCorporationsCorporationIDMembersTitlesBadRequest {
	return &GetCorporationsCorporationIDMembersTitlesBadRequest{}
}

/*
GetCorporationsCorporationIDMembersTitlesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDMembersTitlesBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get corporations corporation Id members titles bad request response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles bad request response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles bad request response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members titles bad request response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members titles bad request response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersTitlesUnauthorized creates a GetCorporationsCorporationIDMembersTitlesUnauthorized with default headers values
func NewGetCorporationsCorporationIDMembersTitlesUnauthorized() *GetCorporationsCorporationIDMembersTitlesUnauthorized {
	return &GetCorporationsCorporationIDMembersTitlesUnauthorized{}
}

/*
GetCorporationsCorporationIDMembersTitlesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMembersTitlesUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get corporations corporation Id members titles unauthorized response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles unauthorized response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles unauthorized response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members titles unauthorized response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members titles unauthorized response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersTitlesForbidden creates a GetCorporationsCorporationIDMembersTitlesForbidden with default headers values
func NewGetCorporationsCorporationIDMembersTitlesForbidden() *GetCorporationsCorporationIDMembersTitlesForbidden {
	return &GetCorporationsCorporationIDMembersTitlesForbidden{}
}

/*
GetCorporationsCorporationIDMembersTitlesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMembersTitlesForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get corporations corporation Id members titles forbidden response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles forbidden response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles forbidden response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members titles forbidden response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members titles forbidden response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCorporationsCorporationIDMembersTitlesForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesForbidden) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersTitlesEnhanceYourCalm creates a GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMembersTitlesEnhanceYourCalm() *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm {
	return &GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm{}
}

/*
GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get corporations corporation Id members titles enhance your calm response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles enhance your calm response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles enhance your calm response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id members titles enhance your calm response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id members titles enhance your calm response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersTitlesInternalServerError creates a GetCorporationsCorporationIDMembersTitlesInternalServerError with default headers values
func NewGetCorporationsCorporationIDMembersTitlesInternalServerError() *GetCorporationsCorporationIDMembersTitlesInternalServerError {
	return &GetCorporationsCorporationIDMembersTitlesInternalServerError{}
}

/*
GetCorporationsCorporationIDMembersTitlesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMembersTitlesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get corporations corporation Id members titles internal server error response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles internal server error response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles internal server error response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members titles internal server error response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members titles internal server error response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersTitlesServiceUnavailable creates a GetCorporationsCorporationIDMembersTitlesServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMembersTitlesServiceUnavailable() *GetCorporationsCorporationIDMembersTitlesServiceUnavailable {
	return &GetCorporationsCorporationIDMembersTitlesServiceUnavailable{}
}

/*
GetCorporationsCorporationIDMembersTitlesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMembersTitlesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get corporations corporation Id members titles service unavailable response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles service unavailable response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles service unavailable response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members titles service unavailable response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members titles service unavailable response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembersTitlesGatewayTimeout creates a GetCorporationsCorporationIDMembersTitlesGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMembersTitlesGatewayTimeout() *GetCorporationsCorporationIDMembersTitlesGatewayTimeout {
	return &GetCorporationsCorporationIDMembersTitlesGatewayTimeout{}
}

/*
GetCorporationsCorporationIDMembersTitlesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMembersTitlesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get corporations corporation Id members titles gateway timeout response has a 2xx status code
func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id members titles gateway timeout response has a 3xx status code
func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id members titles gateway timeout response has a 4xx status code
func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id members titles gateway timeout response has a 5xx status code
func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id members titles gateway timeout response a status code equal to that given
func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/members/titles/][%d] getCorporationsCorporationIdMembersTitlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMembersTitlesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCorporationsCorporationIDMembersTitlesOKBodyItems0 get_corporations_corporation_id_members_titles_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDMembersTitlesOKBodyItems0
*/
type GetCorporationsCorporationIDMembersTitlesOKBodyItems0 struct {

	// get_corporations_corporation_id_members_titles_character_id
	//
	// character_id integer
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporations_corporation_id_members_titles_titles
	//
	// A list of title_id
	// Required: true
	// Max Items: 16
	Titles []int32 `json:"titles"`
}

// Validate validates this get corporations corporation ID members titles o k body items0
func (o *GetCorporationsCorporationIDMembersTitlesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDMembersTitlesOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMembersTitlesOKBodyItems0) validateTitles(formats strfmt.Registry) error {

	if err := validate.Required("titles", "body", o.Titles); err != nil {
		return err
	}

	iTitlesSize := int64(len(o.Titles))

	if err := validate.MaxItems("titles", "body", iTitlesSize, 16); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID members titles o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDMembersTitlesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMembersTitlesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMembersTitlesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDMembersTitlesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
