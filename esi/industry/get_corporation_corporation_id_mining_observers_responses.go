// Code generated by go-swagger; DO NOT EDIT.

package industry

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

// GetCorporationCorporationIDMiningObserversReader is a Reader for the GetCorporationCorporationIDMiningObservers structure.
type GetCorporationCorporationIDMiningObserversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationCorporationIDMiningObserversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationCorporationIDMiningObserversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationCorporationIDMiningObserversNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationCorporationIDMiningObserversBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationCorporationIDMiningObserversUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationCorporationIDMiningObserversForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationCorporationIDMiningObserversEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationCorporationIDMiningObserversInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationCorporationIDMiningObserversServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationCorporationIDMiningObserversGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationCorporationIDMiningObserversOK creates a GetCorporationCorporationIDMiningObserversOK with default headers values
func NewGetCorporationCorporationIDMiningObserversOK() *GetCorporationCorporationIDMiningObserversOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationCorporationIDMiningObserversOK{

		XPages: xPagesDefault,
	}
}

/*
GetCorporationCorporationIDMiningObserversOK describes a response with status code 200, with default header values.

Observer list of a corporation
*/
type GetCorporationCorporationIDMiningObserversOK struct {

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

	/* Maximum page number

	   Format: int32
	   Default: 1
	*/
	XPages int32

	Payload []*GetCorporationCorporationIDMiningObserversOKBodyItems0
}

// IsSuccess returns true when this get corporation corporation Id mining observers o k response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get corporation corporation Id mining observers o k response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers o k response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporation corporation Id mining observers o k response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporation corporation Id mining observers o k response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCorporationCorporationIDMiningObserversOK) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversOK  %+v", 200, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversOK) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversOK  %+v", 200, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversOK) GetPayload() []*GetCorporationCorporationIDMiningObserversOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// hydrates response header X-Pages
	hdrXPages := response.GetHeader("X-Pages")

	if hdrXPages != "" {
		valxPages, err := swag.ConvertInt32(hdrXPages)
		if err != nil {
			return errors.InvalidType("X-Pages", "header", "int32", hdrXPages)
		}
		o.XPages = valxPages
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversNotModified creates a GetCorporationCorporationIDMiningObserversNotModified with default headers values
func NewGetCorporationCorporationIDMiningObserversNotModified() *GetCorporationCorporationIDMiningObserversNotModified {
	return &GetCorporationCorporationIDMiningObserversNotModified{}
}

/*
GetCorporationCorporationIDMiningObserversNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationCorporationIDMiningObserversNotModified struct {

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

// IsSuccess returns true when this get corporation corporation Id mining observers not modified response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers not modified response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get corporation corporation Id mining observers not modified response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporation corporation Id mining observers not modified response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporation corporation Id mining observers not modified response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetCorporationCorporationIDMiningObserversNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversNotModified ", 304)
}

func (o *GetCorporationCorporationIDMiningObserversNotModified) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversNotModified ", 304)
}

func (o *GetCorporationCorporationIDMiningObserversNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationCorporationIDMiningObserversBadRequest creates a GetCorporationCorporationIDMiningObserversBadRequest with default headers values
func NewGetCorporationCorporationIDMiningObserversBadRequest() *GetCorporationCorporationIDMiningObserversBadRequest {
	return &GetCorporationCorporationIDMiningObserversBadRequest{}
}

/*
GetCorporationCorporationIDMiningObserversBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationCorporationIDMiningObserversBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get corporation corporation Id mining observers bad request response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers bad request response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers bad request response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporation corporation Id mining observers bad request response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporation corporation Id mining observers bad request response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCorporationCorporationIDMiningObserversBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversUnauthorized creates a GetCorporationCorporationIDMiningObserversUnauthorized with default headers values
func NewGetCorporationCorporationIDMiningObserversUnauthorized() *GetCorporationCorporationIDMiningObserversUnauthorized {
	return &GetCorporationCorporationIDMiningObserversUnauthorized{}
}

/*
GetCorporationCorporationIDMiningObserversUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationCorporationIDMiningObserversUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get corporation corporation Id mining observers unauthorized response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers unauthorized response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers unauthorized response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporation corporation Id mining observers unauthorized response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporation corporation Id mining observers unauthorized response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCorporationCorporationIDMiningObserversUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversForbidden creates a GetCorporationCorporationIDMiningObserversForbidden with default headers values
func NewGetCorporationCorporationIDMiningObserversForbidden() *GetCorporationCorporationIDMiningObserversForbidden {
	return &GetCorporationCorporationIDMiningObserversForbidden{}
}

/*
GetCorporationCorporationIDMiningObserversForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationCorporationIDMiningObserversForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get corporation corporation Id mining observers forbidden response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers forbidden response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers forbidden response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporation corporation Id mining observers forbidden response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporation corporation Id mining observers forbidden response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCorporationCorporationIDMiningObserversForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversForbidden) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversEnhanceYourCalm creates a GetCorporationCorporationIDMiningObserversEnhanceYourCalm with default headers values
func NewGetCorporationCorporationIDMiningObserversEnhanceYourCalm() *GetCorporationCorporationIDMiningObserversEnhanceYourCalm {
	return &GetCorporationCorporationIDMiningObserversEnhanceYourCalm{}
}

/*
GetCorporationCorporationIDMiningObserversEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationCorporationIDMiningObserversEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get corporation corporation Id mining observers enhance your calm response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers enhance your calm response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers enhance your calm response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporation corporation Id mining observers enhance your calm response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporation corporation Id mining observers enhance your calm response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversInternalServerError creates a GetCorporationCorporationIDMiningObserversInternalServerError with default headers values
func NewGetCorporationCorporationIDMiningObserversInternalServerError() *GetCorporationCorporationIDMiningObserversInternalServerError {
	return &GetCorporationCorporationIDMiningObserversInternalServerError{}
}

/*
GetCorporationCorporationIDMiningObserversInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationCorporationIDMiningObserversInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get corporation corporation Id mining observers internal server error response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers internal server error response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers internal server error response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporation corporation Id mining observers internal server error response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporation corporation Id mining observers internal server error response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCorporationCorporationIDMiningObserversInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversServiceUnavailable creates a GetCorporationCorporationIDMiningObserversServiceUnavailable with default headers values
func NewGetCorporationCorporationIDMiningObserversServiceUnavailable() *GetCorporationCorporationIDMiningObserversServiceUnavailable {
	return &GetCorporationCorporationIDMiningObserversServiceUnavailable{}
}

/*
GetCorporationCorporationIDMiningObserversServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationCorporationIDMiningObserversServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get corporation corporation Id mining observers service unavailable response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers service unavailable response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers service unavailable response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporation corporation Id mining observers service unavailable response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporation corporation Id mining observers service unavailable response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationCorporationIDMiningObserversGatewayTimeout creates a GetCorporationCorporationIDMiningObserversGatewayTimeout with default headers values
func NewGetCorporationCorporationIDMiningObserversGatewayTimeout() *GetCorporationCorporationIDMiningObserversGatewayTimeout {
	return &GetCorporationCorporationIDMiningObserversGatewayTimeout{}
}

/*
GetCorporationCorporationIDMiningObserversGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationCorporationIDMiningObserversGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get corporation corporation Id mining observers gateway timeout response has a 2xx status code
func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporation corporation Id mining observers gateway timeout response has a 3xx status code
func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporation corporation Id mining observers gateway timeout response has a 4xx status code
func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporation corporation Id mining observers gateway timeout response has a 5xx status code
func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporation corporation Id mining observers gateway timeout response a status code equal to that given
func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/corporation/{corporation_id}/mining/observers/][%d] getCorporationCorporationIdMiningObserversGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationCorporationIDMiningObserversGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCorporationCorporationIDMiningObserversOKBodyItems0 get_corporation_corporation_id_mining_observers_200_ok
//
// 200 ok object
swagger:model GetCorporationCorporationIDMiningObserversOKBodyItems0
*/
type GetCorporationCorporationIDMiningObserversOKBodyItems0 struct {

	// get_corporation_corporation_id_mining_observers_last_updated
	//
	// last_updated string
	// Required: true
	// Format: date
	LastUpdated *strfmt.Date `json:"last_updated"`

	// get_corporation_corporation_id_mining_observers_observer_id
	//
	// The entity that was observing the asteroid field when it was mined.
	//
	// Required: true
	ObserverID *int64 `json:"observer_id"`

	// get_corporation_corporation_id_mining_observers_observer_type
	//
	// The category of the observing entity
	// Required: true
	// Enum: [structure]
	ObserverType *string `json:"observer_type"`
}

// Validate validates this get corporation corporation ID mining observers o k body items0
func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateObserverID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateObserverType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("last_updated", "body", o.LastUpdated); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated", "body", "date", o.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) validateObserverID(formats strfmt.Registry) error {

	if err := validate.Required("observer_id", "body", o.ObserverID); err != nil {
		return err
	}

	return nil
}

var getCorporationCorporationIdMiningObserversOKBodyItems0TypeObserverTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["structure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationCorporationIdMiningObserversOKBodyItems0TypeObserverTypePropEnum = append(getCorporationCorporationIdMiningObserversOKBodyItems0TypeObserverTypePropEnum, v)
	}
}

const (

	// GetCorporationCorporationIDMiningObserversOKBodyItems0ObserverTypeStructure captures enum value "structure"
	GetCorporationCorporationIDMiningObserversOKBodyItems0ObserverTypeStructure string = "structure"
)

// prop value enum
func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) validateObserverTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationCorporationIdMiningObserversOKBodyItems0TypeObserverTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) validateObserverType(formats strfmt.Registry) error {

	if err := validate.Required("observer_type", "body", o.ObserverType); err != nil {
		return err
	}

	// value enum
	if err := o.validateObserverTypeEnum("observer_type", "body", *o.ObserverType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporation corporation ID mining observers o k body items0 based on context it is used
func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationCorporationIDMiningObserversOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationCorporationIDMiningObserversOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
