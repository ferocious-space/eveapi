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

// GetCorporationsCorporationIDFacilitiesReader is a Reader for the GetCorporationsCorporationIDFacilities structure.
type GetCorporationsCorporationIDFacilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDFacilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDFacilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDFacilitiesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDFacilitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDFacilitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDFacilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDFacilitiesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDFacilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDFacilitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDFacilitiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/corporations/{corporation_id}/facilities/] get_corporations_corporation_id_facilities", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDFacilitiesOK creates a GetCorporationsCorporationIDFacilitiesOK with default headers values
func NewGetCorporationsCorporationIDFacilitiesOK() *GetCorporationsCorporationIDFacilitiesOK {
	return &GetCorporationsCorporationIDFacilitiesOK{}
}

/*
GetCorporationsCorporationIDFacilitiesOK describes a response with status code 200, with default header values.

List of corporation facilities
*/
type GetCorporationsCorporationIDFacilitiesOK struct {

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

	Payload []*GetCorporationsCorporationIDFacilitiesOKBodyItems0
}

// IsSuccess returns true when this get corporations corporation Id facilities o k response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get corporations corporation Id facilities o k response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities o k response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id facilities o k response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id facilities o k response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get corporations corporation Id facilities o k response
func (o *GetCorporationsCorporationIDFacilitiesOK) Code() int {
	return 200
}

func (o *GetCorporationsCorporationIDFacilitiesOK) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesOK) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesOK) GetPayload() []*GetCorporationsCorporationIDFacilitiesOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDFacilitiesNotModified creates a GetCorporationsCorporationIDFacilitiesNotModified with default headers values
func NewGetCorporationsCorporationIDFacilitiesNotModified() *GetCorporationsCorporationIDFacilitiesNotModified {
	return &GetCorporationsCorporationIDFacilitiesNotModified{}
}

/*
GetCorporationsCorporationIDFacilitiesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDFacilitiesNotModified struct {

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

// IsSuccess returns true when this get corporations corporation Id facilities not modified response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities not modified response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get corporations corporation Id facilities not modified response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id facilities not modified response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id facilities not modified response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get corporations corporation Id facilities not modified response
func (o *GetCorporationsCorporationIDFacilitiesNotModified) Code() int {
	return 304
}

func (o *GetCorporationsCorporationIDFacilitiesNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesNotModified ", 304)
}

func (o *GetCorporationsCorporationIDFacilitiesNotModified) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesNotModified ", 304)
}

func (o *GetCorporationsCorporationIDFacilitiesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDFacilitiesBadRequest creates a GetCorporationsCorporationIDFacilitiesBadRequest with default headers values
func NewGetCorporationsCorporationIDFacilitiesBadRequest() *GetCorporationsCorporationIDFacilitiesBadRequest {
	return &GetCorporationsCorporationIDFacilitiesBadRequest{}
}

/*
GetCorporationsCorporationIDFacilitiesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDFacilitiesBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get corporations corporation Id facilities bad request response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities bad request response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities bad request response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id facilities bad request response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id facilities bad request response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get corporations corporation Id facilities bad request response
func (o *GetCorporationsCorporationIDFacilitiesBadRequest) Code() int {
	return 400
}

func (o *GetCorporationsCorporationIDFacilitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesUnauthorized creates a GetCorporationsCorporationIDFacilitiesUnauthorized with default headers values
func NewGetCorporationsCorporationIDFacilitiesUnauthorized() *GetCorporationsCorporationIDFacilitiesUnauthorized {
	return &GetCorporationsCorporationIDFacilitiesUnauthorized{}
}

/*
GetCorporationsCorporationIDFacilitiesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDFacilitiesUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get corporations corporation Id facilities unauthorized response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities unauthorized response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities unauthorized response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id facilities unauthorized response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id facilities unauthorized response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get corporations corporation Id facilities unauthorized response
func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) Code() int {
	return 401
}

func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesForbidden creates a GetCorporationsCorporationIDFacilitiesForbidden with default headers values
func NewGetCorporationsCorporationIDFacilitiesForbidden() *GetCorporationsCorporationIDFacilitiesForbidden {
	return &GetCorporationsCorporationIDFacilitiesForbidden{}
}

/*
GetCorporationsCorporationIDFacilitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDFacilitiesForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get corporations corporation Id facilities forbidden response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities forbidden response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities forbidden response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id facilities forbidden response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id facilities forbidden response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get corporations corporation Id facilities forbidden response
func (o *GetCorporationsCorporationIDFacilitiesForbidden) Code() int {
	return 403
}

func (o *GetCorporationsCorporationIDFacilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesForbidden) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesEnhanceYourCalm creates a GetCorporationsCorporationIDFacilitiesEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDFacilitiesEnhanceYourCalm() *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm {
	return &GetCorporationsCorporationIDFacilitiesEnhanceYourCalm{}
}

/*
GetCorporationsCorporationIDFacilitiesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDFacilitiesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get corporations corporation Id facilities enhance your calm response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities enhance your calm response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities enhance your calm response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get corporations corporation Id facilities enhance your calm response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get corporations corporation Id facilities enhance your calm response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get corporations corporation Id facilities enhance your calm response
func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesInternalServerError creates a GetCorporationsCorporationIDFacilitiesInternalServerError with default headers values
func NewGetCorporationsCorporationIDFacilitiesInternalServerError() *GetCorporationsCorporationIDFacilitiesInternalServerError {
	return &GetCorporationsCorporationIDFacilitiesInternalServerError{}
}

/*
GetCorporationsCorporationIDFacilitiesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDFacilitiesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get corporations corporation Id facilities internal server error response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities internal server error response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities internal server error response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id facilities internal server error response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id facilities internal server error response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get corporations corporation Id facilities internal server error response
func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) Code() int {
	return 500
}

func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesServiceUnavailable creates a GetCorporationsCorporationIDFacilitiesServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDFacilitiesServiceUnavailable() *GetCorporationsCorporationIDFacilitiesServiceUnavailable {
	return &GetCorporationsCorporationIDFacilitiesServiceUnavailable{}
}

/*
GetCorporationsCorporationIDFacilitiesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDFacilitiesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get corporations corporation Id facilities service unavailable response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities service unavailable response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities service unavailable response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id facilities service unavailable response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id facilities service unavailable response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get corporations corporation Id facilities service unavailable response
func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) Code() int {
	return 503
}

func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesGatewayTimeout creates a GetCorporationsCorporationIDFacilitiesGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDFacilitiesGatewayTimeout() *GetCorporationsCorporationIDFacilitiesGatewayTimeout {
	return &GetCorporationsCorporationIDFacilitiesGatewayTimeout{}
}

/*
GetCorporationsCorporationIDFacilitiesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDFacilitiesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get corporations corporation Id facilities gateway timeout response has a 2xx status code
func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get corporations corporation Id facilities gateway timeout response has a 3xx status code
func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get corporations corporation Id facilities gateway timeout response has a 4xx status code
func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get corporations corporation Id facilities gateway timeout response has a 5xx status code
func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get corporations corporation Id facilities gateway timeout response a status code equal to that given
func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get corporations corporation Id facilities gateway timeout response
func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) Code() int {
	return 504
}

func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCorporationsCorporationIDFacilitiesOKBodyItems0 get_corporations_corporation_id_facilities_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDFacilitiesOKBodyItems0
*/
type GetCorporationsCorporationIDFacilitiesOKBodyItems0 struct {

	// get_corporations_corporation_id_facilities_facility_id
	//
	// facility_id integer
	// Required: true
	FacilityID *int64 `json:"facility_id"`

	// get_corporations_corporation_id_facilities_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_corporations_corporation_id_facilities_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporations corporation ID facilities o k body items0
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) validateFacilityID(formats strfmt.Registry) error {

	if err := validate.Required("facility_id", "body", o.FacilityID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID facilities o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDFacilitiesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
