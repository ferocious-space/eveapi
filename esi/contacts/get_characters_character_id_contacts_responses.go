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

// GetCharactersCharacterIDContactsReader is a Reader for the GetCharactersCharacterIDContacts structure.
type GetCharactersCharacterIDContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDContactsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDContactsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContactsOK creates a GetCharactersCharacterIDContactsOK with default headers values
func NewGetCharactersCharacterIDContactsOK() *GetCharactersCharacterIDContactsOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCharactersCharacterIDContactsOK{

		XPages: xPagesDefault,
	}
}

/*
GetCharactersCharacterIDContactsOK describes a response with status code 200, with default header values.

A list of contacts
*/
type GetCharactersCharacterIDContactsOK struct {

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

	Payload []*GetCharactersCharacterIDContactsOKBodyItems0
}

// IsSuccess returns true when this get characters character Id contacts o k response has a 2xx status code
func (o *GetCharactersCharacterIDContactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id contacts o k response has a 3xx status code
func (o *GetCharactersCharacterIDContactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts o k response has a 4xx status code
func (o *GetCharactersCharacterIDContactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contacts o k response has a 5xx status code
func (o *GetCharactersCharacterIDContactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contacts o k response a status code equal to that given
func (o *GetCharactersCharacterIDContactsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get characters character Id contacts o k response
func (o *GetCharactersCharacterIDContactsOK) Code() int {
	return 200
}

func (o *GetCharactersCharacterIDContactsOK) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContactsOK) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContactsOK) GetPayload() []*GetCharactersCharacterIDContactsOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContactsNotModified creates a GetCharactersCharacterIDContactsNotModified with default headers values
func NewGetCharactersCharacterIDContactsNotModified() *GetCharactersCharacterIDContactsNotModified {
	return &GetCharactersCharacterIDContactsNotModified{}
}

/*
GetCharactersCharacterIDContactsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDContactsNotModified struct {

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

// IsSuccess returns true when this get characters character Id contacts not modified response has a 2xx status code
func (o *GetCharactersCharacterIDContactsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts not modified response has a 3xx status code
func (o *GetCharactersCharacterIDContactsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id contacts not modified response has a 4xx status code
func (o *GetCharactersCharacterIDContactsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contacts not modified response has a 5xx status code
func (o *GetCharactersCharacterIDContactsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contacts not modified response a status code equal to that given
func (o *GetCharactersCharacterIDContactsNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get characters character Id contacts not modified response
func (o *GetCharactersCharacterIDContactsNotModified) Code() int {
	return 304
}

func (o *GetCharactersCharacterIDContactsNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsNotModified ", 304)
}

func (o *GetCharactersCharacterIDContactsNotModified) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsNotModified ", 304)
}

func (o *GetCharactersCharacterIDContactsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContactsBadRequest creates a GetCharactersCharacterIDContactsBadRequest with default headers values
func NewGetCharactersCharacterIDContactsBadRequest() *GetCharactersCharacterIDContactsBadRequest {
	return &GetCharactersCharacterIDContactsBadRequest{}
}

/*
GetCharactersCharacterIDContactsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDContactsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id contacts bad request response has a 2xx status code
func (o *GetCharactersCharacterIDContactsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts bad request response has a 3xx status code
func (o *GetCharactersCharacterIDContactsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts bad request response has a 4xx status code
func (o *GetCharactersCharacterIDContactsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contacts bad request response has a 5xx status code
func (o *GetCharactersCharacterIDContactsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contacts bad request response a status code equal to that given
func (o *GetCharactersCharacterIDContactsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get characters character Id contacts bad request response
func (o *GetCharactersCharacterIDContactsBadRequest) Code() int {
	return 400
}

func (o *GetCharactersCharacterIDContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDContactsBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDContactsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsUnauthorized creates a GetCharactersCharacterIDContactsUnauthorized with default headers values
func NewGetCharactersCharacterIDContactsUnauthorized() *GetCharactersCharacterIDContactsUnauthorized {
	return &GetCharactersCharacterIDContactsUnauthorized{}
}

/*
GetCharactersCharacterIDContactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDContactsUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id contacts unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDContactsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDContactsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDContactsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contacts unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDContactsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contacts unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDContactsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get characters character Id contacts unauthorized response
func (o *GetCharactersCharacterIDContactsUnauthorized) Code() int {
	return 401
}

func (o *GetCharactersCharacterIDContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDContactsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDContactsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsForbidden creates a GetCharactersCharacterIDContactsForbidden with default headers values
func NewGetCharactersCharacterIDContactsForbidden() *GetCharactersCharacterIDContactsForbidden {
	return &GetCharactersCharacterIDContactsForbidden{}
}

/*
GetCharactersCharacterIDContactsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDContactsForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id contacts forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDContactsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDContactsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDContactsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contacts forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDContactsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contacts forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDContactsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get characters character Id contacts forbidden response
func (o *GetCharactersCharacterIDContactsForbidden) Code() int {
	return 403
}

func (o *GetCharactersCharacterIDContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContactsForbidden) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContactsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsEnhanceYourCalm creates a GetCharactersCharacterIDContactsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDContactsEnhanceYourCalm() *GetCharactersCharacterIDContactsEnhanceYourCalm {
	return &GetCharactersCharacterIDContactsEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDContactsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDContactsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id contacts enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contacts enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contacts enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get characters character Id contacts enhance your calm response
func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsInternalServerError creates a GetCharactersCharacterIDContactsInternalServerError with default headers values
func NewGetCharactersCharacterIDContactsInternalServerError() *GetCharactersCharacterIDContactsInternalServerError {
	return &GetCharactersCharacterIDContactsInternalServerError{}
}

/*
GetCharactersCharacterIDContactsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDContactsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id contacts internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDContactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDContactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDContactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contacts internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDContactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id contacts internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDContactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get characters character Id contacts internal server error response
func (o *GetCharactersCharacterIDContactsInternalServerError) Code() int {
	return 500
}

func (o *GetCharactersCharacterIDContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContactsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContactsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsServiceUnavailable creates a GetCharactersCharacterIDContactsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDContactsServiceUnavailable() *GetCharactersCharacterIDContactsServiceUnavailable {
	return &GetCharactersCharacterIDContactsServiceUnavailable{}
}

/*
GetCharactersCharacterIDContactsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDContactsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id contacts service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDContactsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDContactsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDContactsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contacts service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDContactsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id contacts service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDContactsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get characters character Id contacts service unavailable response
func (o *GetCharactersCharacterIDContactsServiceUnavailable) Code() int {
	return 503
}

func (o *GetCharactersCharacterIDContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDContactsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDContactsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsGatewayTimeout creates a GetCharactersCharacterIDContactsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDContactsGatewayTimeout() *GetCharactersCharacterIDContactsGatewayTimeout {
	return &GetCharactersCharacterIDContactsGatewayTimeout{}
}

/*
GetCharactersCharacterIDContactsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDContactsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id contacts gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDContactsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contacts gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDContactsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contacts gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDContactsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contacts gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDContactsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id contacts gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDContactsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get characters character Id contacts gateway timeout response
func (o *GetCharactersCharacterIDContactsGatewayTimeout) Code() int {
	return 504
}

func (o *GetCharactersCharacterIDContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDContactsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/contacts/][%d] getCharactersCharacterIdContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDContactsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCharactersCharacterIDContactsOKBodyItems0 get_characters_character_id_contacts_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDContactsOKBodyItems0
*/
type GetCharactersCharacterIDContactsOKBodyItems0 struct {

	// get_characters_character_id_contacts_contact_id
	//
	// contact_id integer
	// Required: true
	ContactID *int32 `json:"contact_id"`

	// get_characters_character_id_contacts_contact_type
	//
	// contact_type string
	// Required: true
	// Enum: [character corporation alliance faction]
	ContactType *string `json:"contact_type"`

	// get_characters_character_id_contacts_is_blocked
	//
	// Whether this contact is in the blocked list. Note a missing value denotes unknown, not true or false
	IsBlocked bool `json:"is_blocked,omitempty"`

	// get_characters_character_id_contacts_is_watched
	//
	// Whether this contact is being watched
	IsWatched bool `json:"is_watched,omitempty"`

	// get_characters_character_id_contacts_label_ids
	//
	// label_ids array
	// Max Items: 63
	LabelIds []int64 `json:"label_ids"`

	// get_characters_character_id_contacts_standing
	//
	// Standing of the contact
	// Required: true
	Standing *float32 `json:"standing"`
}

// Validate validates this get characters character ID contacts o k body items0
func (o *GetCharactersCharacterIDContactsOKBodyItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetCharactersCharacterIDContactsOKBodyItems0) validateContactID(formats strfmt.Registry) error {

	if err := validate.Required("contact_id", "body", o.ContactID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdContactsOKBodyItems0TypeContactTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["character","corporation","alliance","faction"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdContactsOKBodyItems0TypeContactTypePropEnum = append(getCharactersCharacterIdContactsOKBodyItems0TypeContactTypePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDContactsOKBodyItems0ContactTypeCharacter captures enum value "character"
	GetCharactersCharacterIDContactsOKBodyItems0ContactTypeCharacter string = "character"

	// GetCharactersCharacterIDContactsOKBodyItems0ContactTypeCorporation captures enum value "corporation"
	GetCharactersCharacterIDContactsOKBodyItems0ContactTypeCorporation string = "corporation"

	// GetCharactersCharacterIDContactsOKBodyItems0ContactTypeAlliance captures enum value "alliance"
	GetCharactersCharacterIDContactsOKBodyItems0ContactTypeAlliance string = "alliance"

	// GetCharactersCharacterIDContactsOKBodyItems0ContactTypeFaction captures enum value "faction"
	GetCharactersCharacterIDContactsOKBodyItems0ContactTypeFaction string = "faction"
)

// prop value enum
func (o *GetCharactersCharacterIDContactsOKBodyItems0) validateContactTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdContactsOKBodyItems0TypeContactTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDContactsOKBodyItems0) validateContactType(formats strfmt.Registry) error {

	if err := validate.Required("contact_type", "body", o.ContactType); err != nil {
		return err
	}

	// value enum
	if err := o.validateContactTypeEnum("contact_type", "body", *o.ContactType); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContactsOKBodyItems0) validateLabelIds(formats strfmt.Registry) error {
	if swag.IsZero(o.LabelIds) { // not required
		return nil
	}

	iLabelIdsSize := int64(len(o.LabelIds))

	if err := validate.MaxItems("label_ids", "body", iLabelIdsSize, 63); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContactsOKBodyItems0) validateStanding(formats strfmt.Registry) error {

	if err := validate.Required("standing", "body", o.Standing); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID contacts o k body items0 based on context it is used
func (o *GetCharactersCharacterIDContactsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDContactsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDContactsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDContactsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
