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

// GetUniverseConstellationsReader is a Reader for the GetUniverseConstellations structure.
type GetUniverseConstellationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseConstellationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseConstellationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseConstellationsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseConstellationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseConstellationsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseConstellationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseConstellationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseConstellationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseConstellationsOK creates a GetUniverseConstellationsOK with default headers values
func NewGetUniverseConstellationsOK() *GetUniverseConstellationsOK {
	return &GetUniverseConstellationsOK{}
}

/* GetUniverseConstellationsOK describes a response with status code 200, with default header values.

A list of constellation ids
*/
type GetUniverseConstellationsOK struct {

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

func (o *GetUniverseConstellationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/][%d] getUniverseConstellationsOK  %+v", 200, o.Payload)
}
func (o *GetUniverseConstellationsOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetUniverseConstellationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseConstellationsNotModified creates a GetUniverseConstellationsNotModified with default headers values
func NewGetUniverseConstellationsNotModified() *GetUniverseConstellationsNotModified {
	return &GetUniverseConstellationsNotModified{}
}

/* GetUniverseConstellationsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseConstellationsNotModified struct {

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

func (o *GetUniverseConstellationsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/][%d] getUniverseConstellationsNotModified ", 304)
}

func (o *GetUniverseConstellationsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseConstellationsBadRequest creates a GetUniverseConstellationsBadRequest with default headers values
func NewGetUniverseConstellationsBadRequest() *GetUniverseConstellationsBadRequest {
	return &GetUniverseConstellationsBadRequest{}
}

/* GetUniverseConstellationsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseConstellationsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseConstellationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/][%d] getUniverseConstellationsBadRequest  %+v", 400, o.Payload)
}
func (o *GetUniverseConstellationsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseConstellationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsEnhanceYourCalm creates a GetUniverseConstellationsEnhanceYourCalm with default headers values
func NewGetUniverseConstellationsEnhanceYourCalm() *GetUniverseConstellationsEnhanceYourCalm {
	return &GetUniverseConstellationsEnhanceYourCalm{}
}

/* GetUniverseConstellationsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseConstellationsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseConstellationsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/][%d] getUniverseConstellationsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetUniverseConstellationsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseConstellationsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsInternalServerError creates a GetUniverseConstellationsInternalServerError with default headers values
func NewGetUniverseConstellationsInternalServerError() *GetUniverseConstellationsInternalServerError {
	return &GetUniverseConstellationsInternalServerError{}
}

/* GetUniverseConstellationsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseConstellationsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseConstellationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/][%d] getUniverseConstellationsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUniverseConstellationsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseConstellationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsServiceUnavailable creates a GetUniverseConstellationsServiceUnavailable with default headers values
func NewGetUniverseConstellationsServiceUnavailable() *GetUniverseConstellationsServiceUnavailable {
	return &GetUniverseConstellationsServiceUnavailable{}
}

/* GetUniverseConstellationsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseConstellationsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseConstellationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/][%d] getUniverseConstellationsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUniverseConstellationsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseConstellationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsGatewayTimeout creates a GetUniverseConstellationsGatewayTimeout with default headers values
func NewGetUniverseConstellationsGatewayTimeout() *GetUniverseConstellationsGatewayTimeout {
	return &GetUniverseConstellationsGatewayTimeout{}
}

/* GetUniverseConstellationsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseConstellationsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseConstellationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/][%d] getUniverseConstellationsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUniverseConstellationsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseConstellationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
