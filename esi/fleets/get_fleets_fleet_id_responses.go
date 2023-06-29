// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// GetFleetsFleetIDReader is a Reader for the GetFleetsFleetID structure.
type GetFleetsFleetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetsFleetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFleetsFleetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetFleetsFleetIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetFleetsFleetIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFleetsFleetIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFleetsFleetIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFleetsFleetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetFleetsFleetIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFleetsFleetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFleetsFleetIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFleetsFleetIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/fleets/{fleet_id}/] get_fleets_fleet_id", response, response.Code())
	}
}

// NewGetFleetsFleetIDOK creates a GetFleetsFleetIDOK with default headers values
func NewGetFleetsFleetIDOK() *GetFleetsFleetIDOK {
	return &GetFleetsFleetIDOK{}
}

/*
GetFleetsFleetIDOK describes a response with status code 200, with default header values.

Details about a fleet
*/
type GetFleetsFleetIDOK struct {

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

	Payload *GetFleetsFleetIDOKBody
}

// IsSuccess returns true when this get fleets fleet Id o k response has a 2xx status code
func (o *GetFleetsFleetIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fleets fleet Id o k response has a 3xx status code
func (o *GetFleetsFleetIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id o k response has a 4xx status code
func (o *GetFleetsFleetIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id o k response has a 5xx status code
func (o *GetFleetsFleetIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id o k response a status code equal to that given
func (o *GetFleetsFleetIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get fleets fleet Id o k response
func (o *GetFleetsFleetIDOK) Code() int {
	return 200
}

func (o *GetFleetsFleetIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdOK  %+v", 200, o.Payload)
}

func (o *GetFleetsFleetIDOK) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdOK  %+v", 200, o.Payload)
}

func (o *GetFleetsFleetIDOK) GetPayload() *GetFleetsFleetIDOKBody {
	return o.Payload
}

func (o *GetFleetsFleetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetFleetsFleetIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDNotModified creates a GetFleetsFleetIDNotModified with default headers values
func NewGetFleetsFleetIDNotModified() *GetFleetsFleetIDNotModified {
	return &GetFleetsFleetIDNotModified{}
}

/*
GetFleetsFleetIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetFleetsFleetIDNotModified struct {

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

// IsSuccess returns true when this get fleets fleet Id not modified response has a 2xx status code
func (o *GetFleetsFleetIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id not modified response has a 3xx status code
func (o *GetFleetsFleetIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get fleets fleet Id not modified response has a 4xx status code
func (o *GetFleetsFleetIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id not modified response has a 5xx status code
func (o *GetFleetsFleetIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id not modified response a status code equal to that given
func (o *GetFleetsFleetIDNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get fleets fleet Id not modified response
func (o *GetFleetsFleetIDNotModified) Code() int {
	return 304
}

func (o *GetFleetsFleetIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdNotModified ", 304)
}

func (o *GetFleetsFleetIDNotModified) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdNotModified ", 304)
}

func (o *GetFleetsFleetIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFleetsFleetIDBadRequest creates a GetFleetsFleetIDBadRequest with default headers values
func NewGetFleetsFleetIDBadRequest() *GetFleetsFleetIDBadRequest {
	return &GetFleetsFleetIDBadRequest{}
}

/*
GetFleetsFleetIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFleetsFleetIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get fleets fleet Id bad request response has a 2xx status code
func (o *GetFleetsFleetIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id bad request response has a 3xx status code
func (o *GetFleetsFleetIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id bad request response has a 4xx status code
func (o *GetFleetsFleetIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id bad request response has a 5xx status code
func (o *GetFleetsFleetIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id bad request response a status code equal to that given
func (o *GetFleetsFleetIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get fleets fleet Id bad request response
func (o *GetFleetsFleetIDBadRequest) Code() int {
	return 400
}

func (o *GetFleetsFleetIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetsFleetIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetsFleetIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetFleetsFleetIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDUnauthorized creates a GetFleetsFleetIDUnauthorized with default headers values
func NewGetFleetsFleetIDUnauthorized() *GetFleetsFleetIDUnauthorized {
	return &GetFleetsFleetIDUnauthorized{}
}

/*
GetFleetsFleetIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFleetsFleetIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get fleets fleet Id unauthorized response has a 2xx status code
func (o *GetFleetsFleetIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id unauthorized response has a 3xx status code
func (o *GetFleetsFleetIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id unauthorized response has a 4xx status code
func (o *GetFleetsFleetIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id unauthorized response has a 5xx status code
func (o *GetFleetsFleetIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id unauthorized response a status code equal to that given
func (o *GetFleetsFleetIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get fleets fleet Id unauthorized response
func (o *GetFleetsFleetIDUnauthorized) Code() int {
	return 401
}

func (o *GetFleetsFleetIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFleetsFleetIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFleetsFleetIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetFleetsFleetIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDForbidden creates a GetFleetsFleetIDForbidden with default headers values
func NewGetFleetsFleetIDForbidden() *GetFleetsFleetIDForbidden {
	return &GetFleetsFleetIDForbidden{}
}

/*
GetFleetsFleetIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFleetsFleetIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get fleets fleet Id forbidden response has a 2xx status code
func (o *GetFleetsFleetIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id forbidden response has a 3xx status code
func (o *GetFleetsFleetIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id forbidden response has a 4xx status code
func (o *GetFleetsFleetIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id forbidden response has a 5xx status code
func (o *GetFleetsFleetIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id forbidden response a status code equal to that given
func (o *GetFleetsFleetIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get fleets fleet Id forbidden response
func (o *GetFleetsFleetIDForbidden) Code() int {
	return 403
}

func (o *GetFleetsFleetIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdForbidden  %+v", 403, o.Payload)
}

func (o *GetFleetsFleetIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdForbidden  %+v", 403, o.Payload)
}

func (o *GetFleetsFleetIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetFleetsFleetIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDNotFound creates a GetFleetsFleetIDNotFound with default headers values
func NewGetFleetsFleetIDNotFound() *GetFleetsFleetIDNotFound {
	return &GetFleetsFleetIDNotFound{}
}

/*
GetFleetsFleetIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type GetFleetsFleetIDNotFound struct {
	Payload *GetFleetsFleetIDNotFoundBody
}

// IsSuccess returns true when this get fleets fleet Id not found response has a 2xx status code
func (o *GetFleetsFleetIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id not found response has a 3xx status code
func (o *GetFleetsFleetIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id not found response has a 4xx status code
func (o *GetFleetsFleetIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id not found response has a 5xx status code
func (o *GetFleetsFleetIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id not found response a status code equal to that given
func (o *GetFleetsFleetIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get fleets fleet Id not found response
func (o *GetFleetsFleetIDNotFound) Code() int {
	return 404
}

func (o *GetFleetsFleetIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdNotFound  %+v", 404, o.Payload)
}

func (o *GetFleetsFleetIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdNotFound  %+v", 404, o.Payload)
}

func (o *GetFleetsFleetIDNotFound) GetPayload() *GetFleetsFleetIDNotFoundBody {
	return o.Payload
}

func (o *GetFleetsFleetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFleetsFleetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDEnhanceYourCalm creates a GetFleetsFleetIDEnhanceYourCalm with default headers values
func NewGetFleetsFleetIDEnhanceYourCalm() *GetFleetsFleetIDEnhanceYourCalm {
	return &GetFleetsFleetIDEnhanceYourCalm{}
}

/*
GetFleetsFleetIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetFleetsFleetIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get fleets fleet Id enhance your calm response has a 2xx status code
func (o *GetFleetsFleetIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id enhance your calm response has a 3xx status code
func (o *GetFleetsFleetIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id enhance your calm response has a 4xx status code
func (o *GetFleetsFleetIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fleets fleet Id enhance your calm response has a 5xx status code
func (o *GetFleetsFleetIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get fleets fleet Id enhance your calm response a status code equal to that given
func (o *GetFleetsFleetIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get fleets fleet Id enhance your calm response
func (o *GetFleetsFleetIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetFleetsFleetIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetFleetsFleetIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetFleetsFleetIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetFleetsFleetIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDInternalServerError creates a GetFleetsFleetIDInternalServerError with default headers values
func NewGetFleetsFleetIDInternalServerError() *GetFleetsFleetIDInternalServerError {
	return &GetFleetsFleetIDInternalServerError{}
}

/*
GetFleetsFleetIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetFleetsFleetIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get fleets fleet Id internal server error response has a 2xx status code
func (o *GetFleetsFleetIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id internal server error response has a 3xx status code
func (o *GetFleetsFleetIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id internal server error response has a 4xx status code
func (o *GetFleetsFleetIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id internal server error response has a 5xx status code
func (o *GetFleetsFleetIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get fleets fleet Id internal server error response a status code equal to that given
func (o *GetFleetsFleetIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get fleets fleet Id internal server error response
func (o *GetFleetsFleetIDInternalServerError) Code() int {
	return 500
}

func (o *GetFleetsFleetIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFleetsFleetIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFleetsFleetIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetFleetsFleetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDServiceUnavailable creates a GetFleetsFleetIDServiceUnavailable with default headers values
func NewGetFleetsFleetIDServiceUnavailable() *GetFleetsFleetIDServiceUnavailable {
	return &GetFleetsFleetIDServiceUnavailable{}
}

/*
GetFleetsFleetIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetFleetsFleetIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get fleets fleet Id service unavailable response has a 2xx status code
func (o *GetFleetsFleetIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id service unavailable response has a 3xx status code
func (o *GetFleetsFleetIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id service unavailable response has a 4xx status code
func (o *GetFleetsFleetIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id service unavailable response has a 5xx status code
func (o *GetFleetsFleetIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get fleets fleet Id service unavailable response a status code equal to that given
func (o *GetFleetsFleetIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get fleets fleet Id service unavailable response
func (o *GetFleetsFleetIDServiceUnavailable) Code() int {
	return 503
}

func (o *GetFleetsFleetIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFleetsFleetIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFleetsFleetIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetFleetsFleetIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDGatewayTimeout creates a GetFleetsFleetIDGatewayTimeout with default headers values
func NewGetFleetsFleetIDGatewayTimeout() *GetFleetsFleetIDGatewayTimeout {
	return &GetFleetsFleetIDGatewayTimeout{}
}

/*
GetFleetsFleetIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetFleetsFleetIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get fleets fleet Id gateway timeout response has a 2xx status code
func (o *GetFleetsFleetIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fleets fleet Id gateway timeout response has a 3xx status code
func (o *GetFleetsFleetIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fleets fleet Id gateway timeout response has a 4xx status code
func (o *GetFleetsFleetIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fleets fleet Id gateway timeout response has a 5xx status code
func (o *GetFleetsFleetIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get fleets fleet Id gateway timeout response a status code equal to that given
func (o *GetFleetsFleetIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get fleets fleet Id gateway timeout response
func (o *GetFleetsFleetIDGatewayTimeout) Code() int {
	return 504
}

func (o *GetFleetsFleetIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFleetsFleetIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/][%d] getFleetsFleetIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFleetsFleetIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetFleetsFleetIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetFleetsFleetIDNotFoundBody get_fleets_fleet_id_not_found
//
// Not found
swagger:model GetFleetsFleetIDNotFoundBody
*/
type GetFleetsFleetIDNotFoundBody struct {

	// get_fleets_fleet_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get fleets fleet ID not found body
func (o *GetFleetsFleetIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get fleets fleet ID not found body based on context it is used
func (o *GetFleetsFleetIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetFleetsFleetIDOKBody get_fleets_fleet_id_ok
//
// 200 ok object
swagger:model GetFleetsFleetIDOKBody
*/
type GetFleetsFleetIDOKBody struct {

	// get_fleets_fleet_id_is_free_move
	//
	// Is free-move enabled
	// Required: true
	IsFreeMove *bool `json:"is_free_move"`

	// get_fleets_fleet_id_is_registered
	//
	// Does the fleet have an active fleet advertisement
	// Required: true
	IsRegistered *bool `json:"is_registered"`

	// get_fleets_fleet_id_is_voice_enabled
	//
	// Is EVE Voice enabled
	// Required: true
	IsVoiceEnabled *bool `json:"is_voice_enabled"`

	// get_fleets_fleet_id_motd
	//
	// Fleet MOTD in CCP flavoured HTML
	// Required: true
	Motd *string `json:"motd"`
}

// Validate validates this get fleets fleet ID o k body
func (o *GetFleetsFleetIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsFreeMove(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsRegistered(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsVoiceEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMotd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFleetsFleetIDOKBody) validateIsFreeMove(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"is_free_move", "body", o.IsFreeMove); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDOKBody) validateIsRegistered(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"is_registered", "body", o.IsRegistered); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDOKBody) validateIsVoiceEnabled(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"is_voice_enabled", "body", o.IsVoiceEnabled); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDOKBody) validateMotd(formats strfmt.Registry) error {

	if err := validate.Required("getFleetsFleetIdOK"+"."+"motd", "body", o.Motd); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get fleets fleet ID o k body based on context it is used
func (o *GetFleetsFleetIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
