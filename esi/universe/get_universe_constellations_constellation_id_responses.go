// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// GetUniverseConstellationsConstellationIDReader is a Reader for the GetUniverseConstellationsConstellationID structure.
type GetUniverseConstellationsConstellationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseConstellationsConstellationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseConstellationsConstellationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseConstellationsConstellationIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseConstellationsConstellationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseConstellationsConstellationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseConstellationsConstellationIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseConstellationsConstellationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseConstellationsConstellationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseConstellationsConstellationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/universe/constellations/{constellation_id}/] get_universe_constellations_constellation_id", response, response.Code())
	}
}

// NewGetUniverseConstellationsConstellationIDOK creates a GetUniverseConstellationsConstellationIDOK with default headers values
func NewGetUniverseConstellationsConstellationIDOK() *GetUniverseConstellationsConstellationIDOK {
	return &GetUniverseConstellationsConstellationIDOK{}
}

/*
GetUniverseConstellationsConstellationIDOK describes a response with status code 200, with default header values.

Information about a constellation
*/
type GetUniverseConstellationsConstellationIDOK struct {

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

	Payload *GetUniverseConstellationsConstellationIDOKBody
}

// IsSuccess returns true when this get universe constellations constellation Id o k response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe constellations constellation Id o k response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe constellations constellation Id o k response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe constellations constellation Id o k response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe constellations constellation Id o k response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe constellations constellation Id o k response
func (o *GetUniverseConstellationsConstellationIDOK) Code() int {
	return 200
}

func (o *GetUniverseConstellationsConstellationIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDOK) GetPayload() *GetUniverseConstellationsConstellationIDOKBody {
	return o.Payload
}

func (o *GetUniverseConstellationsConstellationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseConstellationsConstellationIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsConstellationIDNotModified creates a GetUniverseConstellationsConstellationIDNotModified with default headers values
func NewGetUniverseConstellationsConstellationIDNotModified() *GetUniverseConstellationsConstellationIDNotModified {
	return &GetUniverseConstellationsConstellationIDNotModified{}
}

/*
GetUniverseConstellationsConstellationIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseConstellationsConstellationIDNotModified struct {

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

// IsSuccess returns true when this get universe constellations constellation Id not modified response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe constellations constellation Id not modified response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe constellations constellation Id not modified response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe constellations constellation Id not modified response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe constellations constellation Id not modified response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe constellations constellation Id not modified response
func (o *GetUniverseConstellationsConstellationIDNotModified) Code() int {
	return 304
}

func (o *GetUniverseConstellationsConstellationIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdNotModified ", 304)
}

func (o *GetUniverseConstellationsConstellationIDNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdNotModified ", 304)
}

func (o *GetUniverseConstellationsConstellationIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseConstellationsConstellationIDBadRequest creates a GetUniverseConstellationsConstellationIDBadRequest with default headers values
func NewGetUniverseConstellationsConstellationIDBadRequest() *GetUniverseConstellationsConstellationIDBadRequest {
	return &GetUniverseConstellationsConstellationIDBadRequest{}
}

/*
GetUniverseConstellationsConstellationIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseConstellationsConstellationIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe constellations constellation Id bad request response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe constellations constellation Id bad request response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe constellations constellation Id bad request response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe constellations constellation Id bad request response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe constellations constellation Id bad request response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe constellations constellation Id bad request response
func (o *GetUniverseConstellationsConstellationIDBadRequest) Code() int {
	return 400
}

func (o *GetUniverseConstellationsConstellationIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseConstellationsConstellationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsConstellationIDNotFound creates a GetUniverseConstellationsConstellationIDNotFound with default headers values
func NewGetUniverseConstellationsConstellationIDNotFound() *GetUniverseConstellationsConstellationIDNotFound {
	return &GetUniverseConstellationsConstellationIDNotFound{}
}

/*
GetUniverseConstellationsConstellationIDNotFound describes a response with status code 404, with default header values.

Constellation not found
*/
type GetUniverseConstellationsConstellationIDNotFound struct {
	Payload *GetUniverseConstellationsConstellationIDNotFoundBody
}

// IsSuccess returns true when this get universe constellations constellation Id not found response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe constellations constellation Id not found response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe constellations constellation Id not found response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe constellations constellation Id not found response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe constellations constellation Id not found response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get universe constellations constellation Id not found response
func (o *GetUniverseConstellationsConstellationIDNotFound) Code() int {
	return 404
}

func (o *GetUniverseConstellationsConstellationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDNotFound) GetPayload() *GetUniverseConstellationsConstellationIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseConstellationsConstellationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseConstellationsConstellationIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsConstellationIDEnhanceYourCalm creates a GetUniverseConstellationsConstellationIDEnhanceYourCalm with default headers values
func NewGetUniverseConstellationsConstellationIDEnhanceYourCalm() *GetUniverseConstellationsConstellationIDEnhanceYourCalm {
	return &GetUniverseConstellationsConstellationIDEnhanceYourCalm{}
}

/*
GetUniverseConstellationsConstellationIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseConstellationsConstellationIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe constellations constellation Id enhance your calm response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe constellations constellation Id enhance your calm response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe constellations constellation Id enhance your calm response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe constellations constellation Id enhance your calm response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe constellations constellation Id enhance your calm response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe constellations constellation Id enhance your calm response
func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseConstellationsConstellationIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsConstellationIDInternalServerError creates a GetUniverseConstellationsConstellationIDInternalServerError with default headers values
func NewGetUniverseConstellationsConstellationIDInternalServerError() *GetUniverseConstellationsConstellationIDInternalServerError {
	return &GetUniverseConstellationsConstellationIDInternalServerError{}
}

/*
GetUniverseConstellationsConstellationIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseConstellationsConstellationIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe constellations constellation Id internal server error response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe constellations constellation Id internal server error response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe constellations constellation Id internal server error response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe constellations constellation Id internal server error response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe constellations constellation Id internal server error response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe constellations constellation Id internal server error response
func (o *GetUniverseConstellationsConstellationIDInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseConstellationsConstellationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseConstellationsConstellationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsConstellationIDServiceUnavailable creates a GetUniverseConstellationsConstellationIDServiceUnavailable with default headers values
func NewGetUniverseConstellationsConstellationIDServiceUnavailable() *GetUniverseConstellationsConstellationIDServiceUnavailable {
	return &GetUniverseConstellationsConstellationIDServiceUnavailable{}
}

/*
GetUniverseConstellationsConstellationIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseConstellationsConstellationIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe constellations constellation Id service unavailable response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe constellations constellation Id service unavailable response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe constellations constellation Id service unavailable response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe constellations constellation Id service unavailable response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe constellations constellation Id service unavailable response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe constellations constellation Id service unavailable response
func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseConstellationsConstellationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseConstellationsConstellationIDGatewayTimeout creates a GetUniverseConstellationsConstellationIDGatewayTimeout with default headers values
func NewGetUniverseConstellationsConstellationIDGatewayTimeout() *GetUniverseConstellationsConstellationIDGatewayTimeout {
	return &GetUniverseConstellationsConstellationIDGatewayTimeout{}
}

/*
GetUniverseConstellationsConstellationIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseConstellationsConstellationIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe constellations constellation Id gateway timeout response has a 2xx status code
func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe constellations constellation Id gateway timeout response has a 3xx status code
func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe constellations constellation Id gateway timeout response has a 4xx status code
func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe constellations constellation Id gateway timeout response has a 5xx status code
func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe constellations constellation Id gateway timeout response a status code equal to that given
func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe constellations constellation Id gateway timeout response
func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/constellations/{constellation_id}/][%d] getUniverseConstellationsConstellationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseConstellationsConstellationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseConstellationsConstellationIDNotFoundBody get_universe_constellations_constellation_id_not_found
//
// Not found
swagger:model GetUniverseConstellationsConstellationIDNotFoundBody
*/
type GetUniverseConstellationsConstellationIDNotFoundBody struct {

	// get_universe_constellations_constellation_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe constellations constellation ID not found body
func (o *GetUniverseConstellationsConstellationIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe constellations constellation ID not found body based on context it is used
func (o *GetUniverseConstellationsConstellationIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseConstellationsConstellationIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseConstellationsConstellationIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseConstellationsConstellationIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseConstellationsConstellationIDOKBody get_universe_constellations_constellation_id_ok
//
// 200 ok object
swagger:model GetUniverseConstellationsConstellationIDOKBody
*/
type GetUniverseConstellationsConstellationIDOKBody struct {

	// get_universe_constellations_constellation_id_constellation_id
	//
	// constellation_id integer
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_universe_constellations_constellation_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// position
	// Required: true
	Position *GetUniverseConstellationsConstellationIDOKBodyPosition `json:"position"`

	// get_universe_constellations_constellation_id_region_id
	//
	// The region this constellation is in
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_universe_constellations_constellation_id_systems
	//
	// systems array
	// Required: true
	// Max Items: 10000
	Systems []int32 `json:"systems"`
}

// Validate validates this get universe constellations constellation ID o k body
func (o *GetUniverseConstellationsConstellationIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConstellationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSystems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBody) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"constellation_id", "body", o.ConstellationID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseConstellationsConstellationIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseConstellationsConstellationIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBody) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"region_id", "body", o.RegionID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBody) validateSystems(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"systems", "body", o.Systems); err != nil {
		return err
	}

	iSystemsSize := int64(len(o.Systems))

	if err := validate.MaxItems("getUniverseConstellationsConstellationIdOK"+"."+"systems", "body", iSystemsSize, 10000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe constellations constellation ID o k body based on the context it is used
func (o *GetUniverseConstellationsConstellationIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBody) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {

		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseConstellationsConstellationIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseConstellationsConstellationIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseConstellationsConstellationIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseConstellationsConstellationIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseConstellationsConstellationIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseConstellationsConstellationIDOKBodyPosition get_universe_constellations_constellation_id_position
//
// position object
swagger:model GetUniverseConstellationsConstellationIDOKBodyPosition
*/
type GetUniverseConstellationsConstellationIDOKBodyPosition struct {

	// get_universe_constellations_constellation_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_constellations_constellation_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_constellations_constellation_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe constellations constellation ID o k body position
func (o *GetUniverseConstellationsConstellationIDOKBodyPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateX(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateY(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateZ(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseConstellationsConstellationIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseConstellationsConstellationIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe constellations constellation ID o k body position based on context it is used
func (o *GetUniverseConstellationsConstellationIDOKBodyPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseConstellationsConstellationIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseConstellationsConstellationIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseConstellationsConstellationIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
