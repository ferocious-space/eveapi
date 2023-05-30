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

// GetUniverseStructuresReader is a Reader for the GetUniverseStructures structure.
type GetUniverseStructuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStructuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseStructuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseStructuresNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseStructuresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseStructuresEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseStructuresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseStructuresServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseStructuresGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseStructuresOK creates a GetUniverseStructuresOK with default headers values
func NewGetUniverseStructuresOK() *GetUniverseStructuresOK {
	return &GetUniverseStructuresOK{}
}

/*
GetUniverseStructuresOK describes a response with status code 200, with default header values.

List of public structure IDs
*/
type GetUniverseStructuresOK struct {

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

	Payload []*int64
}

// IsSuccess returns true when this get universe structures o k response has a 2xx status code
func (o *GetUniverseStructuresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe structures o k response has a 3xx status code
func (o *GetUniverseStructuresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures o k response has a 4xx status code
func (o *GetUniverseStructuresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures o k response has a 5xx status code
func (o *GetUniverseStructuresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures o k response a status code equal to that given
func (o *GetUniverseStructuresOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe structures o k response
func (o *GetUniverseStructuresOK) Code() int {
	return 200
}

func (o *GetUniverseStructuresOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStructuresOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStructuresOK) GetPayload() []*int64 {
	return o.Payload
}

func (o *GetUniverseStructuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStructuresNotModified creates a GetUniverseStructuresNotModified with default headers values
func NewGetUniverseStructuresNotModified() *GetUniverseStructuresNotModified {
	return &GetUniverseStructuresNotModified{}
}

/*
GetUniverseStructuresNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseStructuresNotModified struct {

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

// IsSuccess returns true when this get universe structures not modified response has a 2xx status code
func (o *GetUniverseStructuresNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures not modified response has a 3xx status code
func (o *GetUniverseStructuresNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe structures not modified response has a 4xx status code
func (o *GetUniverseStructuresNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures not modified response has a 5xx status code
func (o *GetUniverseStructuresNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures not modified response a status code equal to that given
func (o *GetUniverseStructuresNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe structures not modified response
func (o *GetUniverseStructuresNotModified) Code() int {
	return 304
}

func (o *GetUniverseStructuresNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresNotModified ", 304)
}

func (o *GetUniverseStructuresNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresNotModified ", 304)
}

func (o *GetUniverseStructuresNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStructuresBadRequest creates a GetUniverseStructuresBadRequest with default headers values
func NewGetUniverseStructuresBadRequest() *GetUniverseStructuresBadRequest {
	return &GetUniverseStructuresBadRequest{}
}

/*
GetUniverseStructuresBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseStructuresBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe structures bad request response has a 2xx status code
func (o *GetUniverseStructuresBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures bad request response has a 3xx status code
func (o *GetUniverseStructuresBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures bad request response has a 4xx status code
func (o *GetUniverseStructuresBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe structures bad request response has a 5xx status code
func (o *GetUniverseStructuresBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures bad request response a status code equal to that given
func (o *GetUniverseStructuresBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe structures bad request response
func (o *GetUniverseStructuresBadRequest) Code() int {
	return 400
}

func (o *GetUniverseStructuresBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStructuresBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStructuresBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseStructuresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresEnhanceYourCalm creates a GetUniverseStructuresEnhanceYourCalm with default headers values
func NewGetUniverseStructuresEnhanceYourCalm() *GetUniverseStructuresEnhanceYourCalm {
	return &GetUniverseStructuresEnhanceYourCalm{}
}

/*
GetUniverseStructuresEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseStructuresEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe structures enhance your calm response has a 2xx status code
func (o *GetUniverseStructuresEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures enhance your calm response has a 3xx status code
func (o *GetUniverseStructuresEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures enhance your calm response has a 4xx status code
func (o *GetUniverseStructuresEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe structures enhance your calm response has a 5xx status code
func (o *GetUniverseStructuresEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures enhance your calm response a status code equal to that given
func (o *GetUniverseStructuresEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe structures enhance your calm response
func (o *GetUniverseStructuresEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseStructuresEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStructuresEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStructuresEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseStructuresEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresInternalServerError creates a GetUniverseStructuresInternalServerError with default headers values
func NewGetUniverseStructuresInternalServerError() *GetUniverseStructuresInternalServerError {
	return &GetUniverseStructuresInternalServerError{}
}

/*
GetUniverseStructuresInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseStructuresInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe structures internal server error response has a 2xx status code
func (o *GetUniverseStructuresInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures internal server error response has a 3xx status code
func (o *GetUniverseStructuresInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures internal server error response has a 4xx status code
func (o *GetUniverseStructuresInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures internal server error response has a 5xx status code
func (o *GetUniverseStructuresInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe structures internal server error response a status code equal to that given
func (o *GetUniverseStructuresInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe structures internal server error response
func (o *GetUniverseStructuresInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseStructuresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStructuresInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStructuresInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseStructuresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresServiceUnavailable creates a GetUniverseStructuresServiceUnavailable with default headers values
func NewGetUniverseStructuresServiceUnavailable() *GetUniverseStructuresServiceUnavailable {
	return &GetUniverseStructuresServiceUnavailable{}
}

/*
GetUniverseStructuresServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseStructuresServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe structures service unavailable response has a 2xx status code
func (o *GetUniverseStructuresServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures service unavailable response has a 3xx status code
func (o *GetUniverseStructuresServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures service unavailable response has a 4xx status code
func (o *GetUniverseStructuresServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures service unavailable response has a 5xx status code
func (o *GetUniverseStructuresServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe structures service unavailable response a status code equal to that given
func (o *GetUniverseStructuresServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe structures service unavailable response
func (o *GetUniverseStructuresServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseStructuresServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStructuresServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStructuresServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseStructuresServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresGatewayTimeout creates a GetUniverseStructuresGatewayTimeout with default headers values
func NewGetUniverseStructuresGatewayTimeout() *GetUniverseStructuresGatewayTimeout {
	return &GetUniverseStructuresGatewayTimeout{}
}

/*
GetUniverseStructuresGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseStructuresGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe structures gateway timeout response has a 2xx status code
func (o *GetUniverseStructuresGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures gateway timeout response has a 3xx status code
func (o *GetUniverseStructuresGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures gateway timeout response has a 4xx status code
func (o *GetUniverseStructuresGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures gateway timeout response has a 5xx status code
func (o *GetUniverseStructuresGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe structures gateway timeout response a status code equal to that given
func (o *GetUniverseStructuresGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe structures gateway timeout response
func (o *GetUniverseStructuresGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseStructuresGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStructuresGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/structures/][%d] getUniverseStructuresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStructuresGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseStructuresGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
