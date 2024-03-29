// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// GetAlliancesAllianceIDContactsReader is a Reader for the GetAlliancesAllianceIDContacts structure.
type GetAlliancesAllianceIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesAllianceIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlliancesAllianceIDContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetAlliancesAllianceIDContactsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetAlliancesAllianceIDContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAlliancesAllianceIDContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlliancesAllianceIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetAlliancesAllianceIDContactsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlliancesAllianceIDContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlliancesAllianceIDContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlliancesAllianceIDContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/alliances/{alliance_id}/contacts/] get_alliances_alliance_id_contacts", response, response.Code())
	}
}

// NewGetAlliancesAllianceIDContactsOK creates a GetAlliancesAllianceIDContactsOK with default headers values
func NewGetAlliancesAllianceIDContactsOK() *GetAlliancesAllianceIDContactsOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetAlliancesAllianceIDContactsOK{

		XPages: xPagesDefault,
	}
}

/*
GetAlliancesAllianceIDContactsOK describes a response with status code 200, with default header values.

A list of contacts
*/
type GetAlliancesAllianceIDContactsOK struct {

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

	Payload []*GetAlliancesAllianceIDContactsOKBodyItems0
}

// IsSuccess returns true when this get alliances alliance Id contacts o k response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alliances alliance Id contacts o k response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts o k response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alliances alliance Id contacts o k response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alliances alliance Id contacts o k response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alliances alliance Id contacts o k response
func (o *GetAlliancesAllianceIDContactsOK) Code() int {
	return 200
}

func (o *GetAlliancesAllianceIDContactsOK) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsOK  %+v", 200, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsOK) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsOK  %+v", 200, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsOK) GetPayload() []*GetAlliancesAllianceIDContactsOKBodyItems0 {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDContactsNotModified creates a GetAlliancesAllianceIDContactsNotModified with default headers values
func NewGetAlliancesAllianceIDContactsNotModified() *GetAlliancesAllianceIDContactsNotModified {
	return &GetAlliancesAllianceIDContactsNotModified{}
}

/*
GetAlliancesAllianceIDContactsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetAlliancesAllianceIDContactsNotModified struct {

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

// IsSuccess returns true when this get alliances alliance Id contacts not modified response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts not modified response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get alliances alliance Id contacts not modified response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alliances alliance Id contacts not modified response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get alliances alliance Id contacts not modified response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get alliances alliance Id contacts not modified response
func (o *GetAlliancesAllianceIDContactsNotModified) Code() int {
	return 304
}

func (o *GetAlliancesAllianceIDContactsNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsNotModified ", 304)
}

func (o *GetAlliancesAllianceIDContactsNotModified) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsNotModified ", 304)
}

func (o *GetAlliancesAllianceIDContactsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesAllianceIDContactsBadRequest creates a GetAlliancesAllianceIDContactsBadRequest with default headers values
func NewGetAlliancesAllianceIDContactsBadRequest() *GetAlliancesAllianceIDContactsBadRequest {
	return &GetAlliancesAllianceIDContactsBadRequest{}
}

/*
GetAlliancesAllianceIDContactsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlliancesAllianceIDContactsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get alliances alliance Id contacts bad request response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts bad request response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts bad request response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alliances alliance Id contacts bad request response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get alliances alliance Id contacts bad request response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get alliances alliance Id contacts bad request response
func (o *GetAlliancesAllianceIDContactsBadRequest) Code() int {
	return 400
}

func (o *GetAlliancesAllianceIDContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsUnauthorized creates a GetAlliancesAllianceIDContactsUnauthorized with default headers values
func NewGetAlliancesAllianceIDContactsUnauthorized() *GetAlliancesAllianceIDContactsUnauthorized {
	return &GetAlliancesAllianceIDContactsUnauthorized{}
}

/*
GetAlliancesAllianceIDContactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAlliancesAllianceIDContactsUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get alliances alliance Id contacts unauthorized response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts unauthorized response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts unauthorized response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alliances alliance Id contacts unauthorized response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alliances alliance Id contacts unauthorized response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get alliances alliance Id contacts unauthorized response
func (o *GetAlliancesAllianceIDContactsUnauthorized) Code() int {
	return 401
}

func (o *GetAlliancesAllianceIDContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsForbidden creates a GetAlliancesAllianceIDContactsForbidden with default headers values
func NewGetAlliancesAllianceIDContactsForbidden() *GetAlliancesAllianceIDContactsForbidden {
	return &GetAlliancesAllianceIDContactsForbidden{}
}

/*
GetAlliancesAllianceIDContactsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAlliancesAllianceIDContactsForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get alliances alliance Id contacts forbidden response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts forbidden response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts forbidden response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alliances alliance Id contacts forbidden response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get alliances alliance Id contacts forbidden response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get alliances alliance Id contacts forbidden response
func (o *GetAlliancesAllianceIDContactsForbidden) Code() int {
	return 403
}

func (o *GetAlliancesAllianceIDContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsForbidden) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsEnhanceYourCalm creates a GetAlliancesAllianceIDContactsEnhanceYourCalm with default headers values
func NewGetAlliancesAllianceIDContactsEnhanceYourCalm() *GetAlliancesAllianceIDContactsEnhanceYourCalm {
	return &GetAlliancesAllianceIDContactsEnhanceYourCalm{}
}

/*
GetAlliancesAllianceIDContactsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetAlliancesAllianceIDContactsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get alliances alliance Id contacts enhance your calm response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts enhance your calm response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts enhance your calm response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alliances alliance Id contacts enhance your calm response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get alliances alliance Id contacts enhance your calm response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get alliances alliance Id contacts enhance your calm response
func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsInternalServerError creates a GetAlliancesAllianceIDContactsInternalServerError with default headers values
func NewGetAlliancesAllianceIDContactsInternalServerError() *GetAlliancesAllianceIDContactsInternalServerError {
	return &GetAlliancesAllianceIDContactsInternalServerError{}
}

/*
GetAlliancesAllianceIDContactsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAlliancesAllianceIDContactsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get alliances alliance Id contacts internal server error response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts internal server error response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts internal server error response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alliances alliance Id contacts internal server error response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alliances alliance Id contacts internal server error response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get alliances alliance Id contacts internal server error response
func (o *GetAlliancesAllianceIDContactsInternalServerError) Code() int {
	return 500
}

func (o *GetAlliancesAllianceIDContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsServiceUnavailable creates a GetAlliancesAllianceIDContactsServiceUnavailable with default headers values
func NewGetAlliancesAllianceIDContactsServiceUnavailable() *GetAlliancesAllianceIDContactsServiceUnavailable {
	return &GetAlliancesAllianceIDContactsServiceUnavailable{}
}

/*
GetAlliancesAllianceIDContactsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetAlliancesAllianceIDContactsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get alliances alliance Id contacts service unavailable response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts service unavailable response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts service unavailable response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alliances alliance Id contacts service unavailable response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get alliances alliance Id contacts service unavailable response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get alliances alliance Id contacts service unavailable response
func (o *GetAlliancesAllianceIDContactsServiceUnavailable) Code() int {
	return 503
}

func (o *GetAlliancesAllianceIDContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesAllianceIDContactsGatewayTimeout creates a GetAlliancesAllianceIDContactsGatewayTimeout with default headers values
func NewGetAlliancesAllianceIDContactsGatewayTimeout() *GetAlliancesAllianceIDContactsGatewayTimeout {
	return &GetAlliancesAllianceIDContactsGatewayTimeout{}
}

/*
GetAlliancesAllianceIDContactsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetAlliancesAllianceIDContactsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get alliances alliance Id contacts gateway timeout response has a 2xx status code
func (o *GetAlliancesAllianceIDContactsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alliances alliance Id contacts gateway timeout response has a 3xx status code
func (o *GetAlliancesAllianceIDContactsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alliances alliance Id contacts gateway timeout response has a 4xx status code
func (o *GetAlliancesAllianceIDContactsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alliances alliance Id contacts gateway timeout response has a 5xx status code
func (o *GetAlliancesAllianceIDContactsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get alliances alliance Id contacts gateway timeout response a status code equal to that given
func (o *GetAlliancesAllianceIDContactsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get alliances alliance Id contacts gateway timeout response
func (o *GetAlliancesAllianceIDContactsGatewayTimeout) Code() int {
	return 504
}

func (o *GetAlliancesAllianceIDContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/alliances/{alliance_id}/contacts/][%d] getAlliancesAllianceIdContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlliancesAllianceIDContactsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetAlliancesAllianceIDContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetAlliancesAllianceIDContactsOKBodyItems0 get_alliances_alliance_id_contacts_200_ok
//
// 200 ok object
swagger:model GetAlliancesAllianceIDContactsOKBodyItems0
*/
type GetAlliancesAllianceIDContactsOKBodyItems0 struct {

	// get_alliances_alliance_id_contacts_contact_id
	//
	// contact_id integer
	// Required: true
	ContactID *int32 `json:"contact_id"`

	// get_alliances_alliance_id_contacts_contact_type
	//
	// contact_type string
	// Required: true
	// Enum: [character corporation alliance faction]
	ContactType *string `json:"contact_type"`

	// get_alliances_alliance_id_contacts_label_ids
	//
	// label_ids array
	// Max Items: 63
	LabelIds []int64 `json:"label_ids"`

	// get_alliances_alliance_id_contacts_standing
	//
	// Standing of the contact
	// Required: true
	Standing *float32 `json:"standing"`
}

// Validate validates this get alliances alliance ID contacts o k body items0
func (o *GetAlliancesAllianceIDContactsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContactID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateContactType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabelIds(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStanding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAlliancesAllianceIDContactsOKBodyItems0) validateContactID(formats strfmt.Registry) error {

	if err := validate.Required("contact_id", "body", o.ContactID); err != nil {
		return err
	}

	return nil
}

var getAlliancesAllianceIdContactsOKBodyItems0TypeContactTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["character","corporation","alliance","faction"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAlliancesAllianceIdContactsOKBodyItems0TypeContactTypePropEnum = append(getAlliancesAllianceIdContactsOKBodyItems0TypeContactTypePropEnum, v)
	}
}

const (

	// GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeCharacter captures enum value "character"
	GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeCharacter string = "character"

	// GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeCorporation captures enum value "corporation"
	GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeCorporation string = "corporation"

	// GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeAlliance captures enum value "alliance"
	GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeAlliance string = "alliance"

	// GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeFaction captures enum value "faction"
	GetAlliancesAllianceIDContactsOKBodyItems0ContactTypeFaction string = "faction"
)

// prop value enum
func (o *GetAlliancesAllianceIDContactsOKBodyItems0) validateContactTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getAlliancesAllianceIdContactsOKBodyItems0TypeContactTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetAlliancesAllianceIDContactsOKBodyItems0) validateContactType(formats strfmt.Registry) error {

	if err := validate.Required("contact_type", "body", o.ContactType); err != nil {
		return err
	}

	// value enum
	if err := o.validateContactTypeEnum("contact_type", "body", *o.ContactType); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDContactsOKBodyItems0) validateLabelIds(formats strfmt.Registry) error {
	if swag.IsZero(o.LabelIds) { // not required
		return nil
	}

	iLabelIdsSize := int64(len(o.LabelIds))

	if err := validate.MaxItems("label_ids", "body", iLabelIdsSize, 63); err != nil {
		return err
	}

	return nil
}

func (o *GetAlliancesAllianceIDContactsOKBodyItems0) validateStanding(formats strfmt.Registry) error {

	if err := validate.Required("standing", "body", o.Standing); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get alliances alliance ID contacts o k body items0 based on context it is used
func (o *GetAlliancesAllianceIDContactsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAlliancesAllianceIDContactsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAlliancesAllianceIDContactsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDContactsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
