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

// GetUniverseStructuresStructureIDReader is a Reader for the GetUniverseStructuresStructureID structure.
type GetUniverseStructuresStructureIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStructuresStructureIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseStructuresStructureIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseStructuresStructureIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseStructuresStructureIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUniverseStructuresStructureIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUniverseStructuresStructureIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseStructuresStructureIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseStructuresStructureIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseStructuresStructureIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseStructuresStructureIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseStructuresStructureIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseStructuresStructureIDOK creates a GetUniverseStructuresStructureIDOK with default headers values
func NewGetUniverseStructuresStructureIDOK() *GetUniverseStructuresStructureIDOK {
	return &GetUniverseStructuresStructureIDOK{}
}

/*
GetUniverseStructuresStructureIDOK describes a response with status code 200, with default header values.

Data about a structure
*/
type GetUniverseStructuresStructureIDOK struct {

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

	Payload *GetUniverseStructuresStructureIDOKBody
}

// IsSuccess returns true when this get universe structures structure Id o k response has a 2xx status code
func (o *GetUniverseStructuresStructureIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe structures structure Id o k response has a 3xx status code
func (o *GetUniverseStructuresStructureIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id o k response has a 4xx status code
func (o *GetUniverseStructuresStructureIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures structure Id o k response has a 5xx status code
func (o *GetUniverseStructuresStructureIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures structure Id o k response a status code equal to that given
func (o *GetUniverseStructuresStructureIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUniverseStructuresStructureIDOK) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStructuresStructureIDOK) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStructuresStructureIDOK) GetPayload() *GetUniverseStructuresStructureIDOKBody {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseStructuresStructureIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDNotModified creates a GetUniverseStructuresStructureIDNotModified with default headers values
func NewGetUniverseStructuresStructureIDNotModified() *GetUniverseStructuresStructureIDNotModified {
	return &GetUniverseStructuresStructureIDNotModified{}
}

/*
GetUniverseStructuresStructureIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseStructuresStructureIDNotModified struct {

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

// IsSuccess returns true when this get universe structures structure Id not modified response has a 2xx status code
func (o *GetUniverseStructuresStructureIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id not modified response has a 3xx status code
func (o *GetUniverseStructuresStructureIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe structures structure Id not modified response has a 4xx status code
func (o *GetUniverseStructuresStructureIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures structure Id not modified response has a 5xx status code
func (o *GetUniverseStructuresStructureIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures structure Id not modified response a status code equal to that given
func (o *GetUniverseStructuresStructureIDNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetUniverseStructuresStructureIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdNotModified ", 304)
}

func (o *GetUniverseStructuresStructureIDNotModified) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdNotModified ", 304)
}

func (o *GetUniverseStructuresStructureIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStructuresStructureIDBadRequest creates a GetUniverseStructuresStructureIDBadRequest with default headers values
func NewGetUniverseStructuresStructureIDBadRequest() *GetUniverseStructuresStructureIDBadRequest {
	return &GetUniverseStructuresStructureIDBadRequest{}
}

/*
GetUniverseStructuresStructureIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseStructuresStructureIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe structures structure Id bad request response has a 2xx status code
func (o *GetUniverseStructuresStructureIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id bad request response has a 3xx status code
func (o *GetUniverseStructuresStructureIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id bad request response has a 4xx status code
func (o *GetUniverseStructuresStructureIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe structures structure Id bad request response has a 5xx status code
func (o *GetUniverseStructuresStructureIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures structure Id bad request response a status code equal to that given
func (o *GetUniverseStructuresStructureIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUniverseStructuresStructureIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStructuresStructureIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStructuresStructureIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDUnauthorized creates a GetUniverseStructuresStructureIDUnauthorized with default headers values
func NewGetUniverseStructuresStructureIDUnauthorized() *GetUniverseStructuresStructureIDUnauthorized {
	return &GetUniverseStructuresStructureIDUnauthorized{}
}

/*
GetUniverseStructuresStructureIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetUniverseStructuresStructureIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get universe structures structure Id unauthorized response has a 2xx status code
func (o *GetUniverseStructuresStructureIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id unauthorized response has a 3xx status code
func (o *GetUniverseStructuresStructureIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id unauthorized response has a 4xx status code
func (o *GetUniverseStructuresStructureIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe structures structure Id unauthorized response has a 5xx status code
func (o *GetUniverseStructuresStructureIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures structure Id unauthorized response a status code equal to that given
func (o *GetUniverseStructuresStructureIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUniverseStructuresStructureIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUniverseStructuresStructureIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUniverseStructuresStructureIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDForbidden creates a GetUniverseStructuresStructureIDForbidden with default headers values
func NewGetUniverseStructuresStructureIDForbidden() *GetUniverseStructuresStructureIDForbidden {
	return &GetUniverseStructuresStructureIDForbidden{}
}

/*
GetUniverseStructuresStructureIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUniverseStructuresStructureIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get universe structures structure Id forbidden response has a 2xx status code
func (o *GetUniverseStructuresStructureIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id forbidden response has a 3xx status code
func (o *GetUniverseStructuresStructureIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id forbidden response has a 4xx status code
func (o *GetUniverseStructuresStructureIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe structures structure Id forbidden response has a 5xx status code
func (o *GetUniverseStructuresStructureIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures structure Id forbidden response a status code equal to that given
func (o *GetUniverseStructuresStructureIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUniverseStructuresStructureIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdForbidden  %+v", 403, o.Payload)
}

func (o *GetUniverseStructuresStructureIDForbidden) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdForbidden  %+v", 403, o.Payload)
}

func (o *GetUniverseStructuresStructureIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDNotFound creates a GetUniverseStructuresStructureIDNotFound with default headers values
func NewGetUniverseStructuresStructureIDNotFound() *GetUniverseStructuresStructureIDNotFound {
	return &GetUniverseStructuresStructureIDNotFound{}
}

/*
GetUniverseStructuresStructureIDNotFound describes a response with status code 404, with default header values.

Structure not found
*/
type GetUniverseStructuresStructureIDNotFound struct {
	Payload *GetUniverseStructuresStructureIDNotFoundBody
}

// IsSuccess returns true when this get universe structures structure Id not found response has a 2xx status code
func (o *GetUniverseStructuresStructureIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id not found response has a 3xx status code
func (o *GetUniverseStructuresStructureIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id not found response has a 4xx status code
func (o *GetUniverseStructuresStructureIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe structures structure Id not found response has a 5xx status code
func (o *GetUniverseStructuresStructureIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures structure Id not found response a status code equal to that given
func (o *GetUniverseStructuresStructureIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUniverseStructuresStructureIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseStructuresStructureIDNotFound) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseStructuresStructureIDNotFound) GetPayload() *GetUniverseStructuresStructureIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseStructuresStructureIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDEnhanceYourCalm creates a GetUniverseStructuresStructureIDEnhanceYourCalm with default headers values
func NewGetUniverseStructuresStructureIDEnhanceYourCalm() *GetUniverseStructuresStructureIDEnhanceYourCalm {
	return &GetUniverseStructuresStructureIDEnhanceYourCalm{}
}

/*
GetUniverseStructuresStructureIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseStructuresStructureIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe structures structure Id enhance your calm response has a 2xx status code
func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id enhance your calm response has a 3xx status code
func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id enhance your calm response has a 4xx status code
func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe structures structure Id enhance your calm response has a 5xx status code
func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe structures structure Id enhance your calm response a status code equal to that given
func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDInternalServerError creates a GetUniverseStructuresStructureIDInternalServerError with default headers values
func NewGetUniverseStructuresStructureIDInternalServerError() *GetUniverseStructuresStructureIDInternalServerError {
	return &GetUniverseStructuresStructureIDInternalServerError{}
}

/*
GetUniverseStructuresStructureIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseStructuresStructureIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe structures structure Id internal server error response has a 2xx status code
func (o *GetUniverseStructuresStructureIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id internal server error response has a 3xx status code
func (o *GetUniverseStructuresStructureIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id internal server error response has a 4xx status code
func (o *GetUniverseStructuresStructureIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures structure Id internal server error response has a 5xx status code
func (o *GetUniverseStructuresStructureIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe structures structure Id internal server error response a status code equal to that given
func (o *GetUniverseStructuresStructureIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUniverseStructuresStructureIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStructuresStructureIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStructuresStructureIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDServiceUnavailable creates a GetUniverseStructuresStructureIDServiceUnavailable with default headers values
func NewGetUniverseStructuresStructureIDServiceUnavailable() *GetUniverseStructuresStructureIDServiceUnavailable {
	return &GetUniverseStructuresStructureIDServiceUnavailable{}
}

/*
GetUniverseStructuresStructureIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseStructuresStructureIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe structures structure Id service unavailable response has a 2xx status code
func (o *GetUniverseStructuresStructureIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id service unavailable response has a 3xx status code
func (o *GetUniverseStructuresStructureIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id service unavailable response has a 4xx status code
func (o *GetUniverseStructuresStructureIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures structure Id service unavailable response has a 5xx status code
func (o *GetUniverseStructuresStructureIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe structures structure Id service unavailable response a status code equal to that given
func (o *GetUniverseStructuresStructureIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUniverseStructuresStructureIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStructuresStructureIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStructuresStructureIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresStructureIDGatewayTimeout creates a GetUniverseStructuresStructureIDGatewayTimeout with default headers values
func NewGetUniverseStructuresStructureIDGatewayTimeout() *GetUniverseStructuresStructureIDGatewayTimeout {
	return &GetUniverseStructuresStructureIDGatewayTimeout{}
}

/*
GetUniverseStructuresStructureIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseStructuresStructureIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe structures structure Id gateway timeout response has a 2xx status code
func (o *GetUniverseStructuresStructureIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe structures structure Id gateway timeout response has a 3xx status code
func (o *GetUniverseStructuresStructureIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe structures structure Id gateway timeout response has a 4xx status code
func (o *GetUniverseStructuresStructureIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe structures structure Id gateway timeout response has a 5xx status code
func (o *GetUniverseStructuresStructureIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe structures structure Id gateway timeout response a status code equal to that given
func (o *GetUniverseStructuresStructureIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUniverseStructuresStructureIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStructuresStructureIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/universe/structures/{structure_id}/][%d] getUniverseStructuresStructureIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStructuresStructureIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseStructuresStructureIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseStructuresStructureIDNotFoundBody get_universe_structures_structure_id_not_found
//
// Not found
swagger:model GetUniverseStructuresStructureIDNotFoundBody
*/
type GetUniverseStructuresStructureIDNotFoundBody struct {

	// get_universe_structures_structure_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe structures structure ID not found body
func (o *GetUniverseStructuresStructureIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe structures structure ID not found body based on context it is used
func (o *GetUniverseStructuresStructureIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStructuresStructureIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStructuresStructureIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStructuresStructureIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseStructuresStructureIDOKBody get_universe_structures_structure_id_ok
//
// 200 ok object
swagger:model GetUniverseStructuresStructureIDOKBody
*/
type GetUniverseStructuresStructureIDOKBody struct {

	// get_universe_structures_structure_id_name
	//
	// The full name of the structure
	// Required: true
	Name *string `json:"name"`

	// get_universe_structures_structure_id_owner_id
	//
	// The ID of the corporation who owns this particular structure
	// Required: true
	OwnerID *int32 `json:"owner_id"`

	// position
	Position *GetUniverseStructuresStructureIDOKBodyPosition `json:"position,omitempty"`

	// get_universe_structures_structure_id_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_universe_structures_structure_id_type_id
	//
	// type_id integer
	TypeID int32 `json:"type_id,omitempty"`
}

// Validate validates this get universe structures structure ID o k body
func (o *GetUniverseStructuresStructureIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"owner_id", "body", o.OwnerID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validatePosition(formats strfmt.Registry) error {
	if swag.IsZero(o.Position) { // not required
		return nil
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStructuresStructureIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStructuresStructureIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe structures structure ID o k body based on the context it is used
func (o *GetUniverseStructuresStructureIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseStructuresStructureIDOKBody) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {
		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStructuresStructureIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStructuresStructureIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStructuresStructureIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStructuresStructureIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStructuresStructureIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseStructuresStructureIDOKBodyPosition get_universe_structures_structure_id_position
//
// Coordinates of the structure in Cartesian space relative to the Sun, in metres.
//
swagger:model GetUniverseStructuresStructureIDOKBodyPosition
*/
type GetUniverseStructuresStructureIDOKBodyPosition struct {

	// get_universe_structures_structure_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_structures_structure_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_structures_structure_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe structures structure ID o k body position
func (o *GetUniverseStructuresStructureIDOKBodyPosition) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseStructuresStructureIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStructuresStructureIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStructuresStructureIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe structures structure ID o k body position based on context it is used
func (o *GetUniverseStructuresStructureIDOKBodyPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStructuresStructureIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStructuresStructureIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseStructuresStructureIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
