// Code generated by go-swagger; DO NOT EDIT.

package opportunities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/models"
)

// GetOpportunitiesTasksReader is a Reader for the GetOpportunitiesTasks structure.
type GetOpportunitiesTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOpportunitiesTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOpportunitiesTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetOpportunitiesTasksNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetOpportunitiesTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetOpportunitiesTasksEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOpportunitiesTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOpportunitiesTasksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOpportunitiesTasksGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOpportunitiesTasksOK creates a GetOpportunitiesTasksOK with default headers values
func NewGetOpportunitiesTasksOK() *GetOpportunitiesTasksOK {
	return &GetOpportunitiesTasksOK{}
}

/* GetOpportunitiesTasksOK describes a response with status code 200, with default header values.

A list of opportunities task ids
*/
type GetOpportunitiesTasksOK struct {

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

func (o *GetOpportunitiesTasksOK) Error() string {
	return fmt.Sprintf("[GET /v1/opportunities/tasks/][%d] getOpportunitiesTasksOK  %+v", 200, o.Payload)
}
func (o *GetOpportunitiesTasksOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetOpportunitiesTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOpportunitiesTasksNotModified creates a GetOpportunitiesTasksNotModified with default headers values
func NewGetOpportunitiesTasksNotModified() *GetOpportunitiesTasksNotModified {
	return &GetOpportunitiesTasksNotModified{}
}

/* GetOpportunitiesTasksNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetOpportunitiesTasksNotModified struct {

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

func (o *GetOpportunitiesTasksNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/opportunities/tasks/][%d] getOpportunitiesTasksNotModified ", 304)
}

func (o *GetOpportunitiesTasksNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetOpportunitiesTasksBadRequest creates a GetOpportunitiesTasksBadRequest with default headers values
func NewGetOpportunitiesTasksBadRequest() *GetOpportunitiesTasksBadRequest {
	return &GetOpportunitiesTasksBadRequest{}
}

/* GetOpportunitiesTasksBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetOpportunitiesTasksBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetOpportunitiesTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/opportunities/tasks/][%d] getOpportunitiesTasksBadRequest  %+v", 400, o.Payload)
}
func (o *GetOpportunitiesTasksBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetOpportunitiesTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpportunitiesTasksEnhanceYourCalm creates a GetOpportunitiesTasksEnhanceYourCalm with default headers values
func NewGetOpportunitiesTasksEnhanceYourCalm() *GetOpportunitiesTasksEnhanceYourCalm {
	return &GetOpportunitiesTasksEnhanceYourCalm{}
}

/* GetOpportunitiesTasksEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetOpportunitiesTasksEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetOpportunitiesTasksEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/opportunities/tasks/][%d] getOpportunitiesTasksEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetOpportunitiesTasksEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetOpportunitiesTasksEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpportunitiesTasksInternalServerError creates a GetOpportunitiesTasksInternalServerError with default headers values
func NewGetOpportunitiesTasksInternalServerError() *GetOpportunitiesTasksInternalServerError {
	return &GetOpportunitiesTasksInternalServerError{}
}

/* GetOpportunitiesTasksInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetOpportunitiesTasksInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetOpportunitiesTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/opportunities/tasks/][%d] getOpportunitiesTasksInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOpportunitiesTasksInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetOpportunitiesTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpportunitiesTasksServiceUnavailable creates a GetOpportunitiesTasksServiceUnavailable with default headers values
func NewGetOpportunitiesTasksServiceUnavailable() *GetOpportunitiesTasksServiceUnavailable {
	return &GetOpportunitiesTasksServiceUnavailable{}
}

/* GetOpportunitiesTasksServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetOpportunitiesTasksServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetOpportunitiesTasksServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/opportunities/tasks/][%d] getOpportunitiesTasksServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetOpportunitiesTasksServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetOpportunitiesTasksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpportunitiesTasksGatewayTimeout creates a GetOpportunitiesTasksGatewayTimeout with default headers values
func NewGetOpportunitiesTasksGatewayTimeout() *GetOpportunitiesTasksGatewayTimeout {
	return &GetOpportunitiesTasksGatewayTimeout{}
}

/* GetOpportunitiesTasksGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetOpportunitiesTasksGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetOpportunitiesTasksGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/opportunities/tasks/][%d] getOpportunitiesTasksGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetOpportunitiesTasksGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetOpportunitiesTasksGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
