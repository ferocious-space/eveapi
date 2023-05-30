// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/models"
)

// GetUniverseGraphicsReader is a Reader for the GetUniverseGraphics structure.
type GetUniverseGraphicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseGraphicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseGraphicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseGraphicsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseGraphicsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseGraphicsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseGraphicsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseGraphicsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseGraphicsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseGraphicsOK creates a GetUniverseGraphicsOK with default headers values
func NewGetUniverseGraphicsOK() *GetUniverseGraphicsOK {
	return &GetUniverseGraphicsOK{}
}

/*
GetUniverseGraphicsOK describes a response with status code 200, with default header values.

A list of graphic ids
*/
type GetUniverseGraphicsOK struct {

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

// IsSuccess returns true when this get universe graphics o k response has a 2xx status code
func (o *GetUniverseGraphicsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe graphics o k response has a 3xx status code
func (o *GetUniverseGraphicsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe graphics o k response has a 4xx status code
func (o *GetUniverseGraphicsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe graphics o k response has a 5xx status code
func (o *GetUniverseGraphicsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe graphics o k response a status code equal to that given
func (o *GetUniverseGraphicsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe graphics o k response
func (o *GetUniverseGraphicsOK) Code() int {
	return 200
}

func (o *GetUniverseGraphicsOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseGraphicsOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseGraphicsOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetUniverseGraphicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseGraphicsNotModified creates a GetUniverseGraphicsNotModified with default headers values
func NewGetUniverseGraphicsNotModified() *GetUniverseGraphicsNotModified {
	return &GetUniverseGraphicsNotModified{}
}

/*
GetUniverseGraphicsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseGraphicsNotModified struct {

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

// IsSuccess returns true when this get universe graphics not modified response has a 2xx status code
func (o *GetUniverseGraphicsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe graphics not modified response has a 3xx status code
func (o *GetUniverseGraphicsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe graphics not modified response has a 4xx status code
func (o *GetUniverseGraphicsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe graphics not modified response has a 5xx status code
func (o *GetUniverseGraphicsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe graphics not modified response a status code equal to that given
func (o *GetUniverseGraphicsNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe graphics not modified response
func (o *GetUniverseGraphicsNotModified) Code() int {
	return 304
}

func (o *GetUniverseGraphicsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsNotModified ", 304)
}

func (o *GetUniverseGraphicsNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsNotModified ", 304)
}

func (o *GetUniverseGraphicsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseGraphicsBadRequest creates a GetUniverseGraphicsBadRequest with default headers values
func NewGetUniverseGraphicsBadRequest() *GetUniverseGraphicsBadRequest {
	return &GetUniverseGraphicsBadRequest{}
}

/*
GetUniverseGraphicsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseGraphicsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe graphics bad request response has a 2xx status code
func (o *GetUniverseGraphicsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe graphics bad request response has a 3xx status code
func (o *GetUniverseGraphicsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe graphics bad request response has a 4xx status code
func (o *GetUniverseGraphicsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe graphics bad request response has a 5xx status code
func (o *GetUniverseGraphicsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe graphics bad request response a status code equal to that given
func (o *GetUniverseGraphicsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe graphics bad request response
func (o *GetUniverseGraphicsBadRequest) Code() int {
	return 400
}

func (o *GetUniverseGraphicsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseGraphicsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseGraphicsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseGraphicsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsEnhanceYourCalm creates a GetUniverseGraphicsEnhanceYourCalm with default headers values
func NewGetUniverseGraphicsEnhanceYourCalm() *GetUniverseGraphicsEnhanceYourCalm {
	return &GetUniverseGraphicsEnhanceYourCalm{}
}

/*
GetUniverseGraphicsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseGraphicsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe graphics enhance your calm response has a 2xx status code
func (o *GetUniverseGraphicsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe graphics enhance your calm response has a 3xx status code
func (o *GetUniverseGraphicsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe graphics enhance your calm response has a 4xx status code
func (o *GetUniverseGraphicsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe graphics enhance your calm response has a 5xx status code
func (o *GetUniverseGraphicsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe graphics enhance your calm response a status code equal to that given
func (o *GetUniverseGraphicsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe graphics enhance your calm response
func (o *GetUniverseGraphicsEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseGraphicsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseGraphicsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseGraphicsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseGraphicsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsInternalServerError creates a GetUniverseGraphicsInternalServerError with default headers values
func NewGetUniverseGraphicsInternalServerError() *GetUniverseGraphicsInternalServerError {
	return &GetUniverseGraphicsInternalServerError{}
}

/*
GetUniverseGraphicsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseGraphicsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe graphics internal server error response has a 2xx status code
func (o *GetUniverseGraphicsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe graphics internal server error response has a 3xx status code
func (o *GetUniverseGraphicsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe graphics internal server error response has a 4xx status code
func (o *GetUniverseGraphicsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe graphics internal server error response has a 5xx status code
func (o *GetUniverseGraphicsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe graphics internal server error response a status code equal to that given
func (o *GetUniverseGraphicsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe graphics internal server error response
func (o *GetUniverseGraphicsInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseGraphicsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseGraphicsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseGraphicsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseGraphicsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsServiceUnavailable creates a GetUniverseGraphicsServiceUnavailable with default headers values
func NewGetUniverseGraphicsServiceUnavailable() *GetUniverseGraphicsServiceUnavailable {
	return &GetUniverseGraphicsServiceUnavailable{}
}

/*
GetUniverseGraphicsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseGraphicsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe graphics service unavailable response has a 2xx status code
func (o *GetUniverseGraphicsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe graphics service unavailable response has a 3xx status code
func (o *GetUniverseGraphicsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe graphics service unavailable response has a 4xx status code
func (o *GetUniverseGraphicsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe graphics service unavailable response has a 5xx status code
func (o *GetUniverseGraphicsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe graphics service unavailable response a status code equal to that given
func (o *GetUniverseGraphicsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe graphics service unavailable response
func (o *GetUniverseGraphicsServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseGraphicsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseGraphicsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseGraphicsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseGraphicsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGraphicsGatewayTimeout creates a GetUniverseGraphicsGatewayTimeout with default headers values
func NewGetUniverseGraphicsGatewayTimeout() *GetUniverseGraphicsGatewayTimeout {
	return &GetUniverseGraphicsGatewayTimeout{}
}

/*
GetUniverseGraphicsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseGraphicsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe graphics gateway timeout response has a 2xx status code
func (o *GetUniverseGraphicsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe graphics gateway timeout response has a 3xx status code
func (o *GetUniverseGraphicsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe graphics gateway timeout response has a 4xx status code
func (o *GetUniverseGraphicsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe graphics gateway timeout response has a 5xx status code
func (o *GetUniverseGraphicsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe graphics gateway timeout response a status code equal to that given
func (o *GetUniverseGraphicsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe graphics gateway timeout response
func (o *GetUniverseGraphicsGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseGraphicsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseGraphicsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/graphics/][%d] getUniverseGraphicsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseGraphicsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseGraphicsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
