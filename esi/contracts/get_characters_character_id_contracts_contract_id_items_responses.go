// Code generated by go-swagger; DO NOT EDIT.

package contracts

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

// GetCharactersCharacterIDContractsContractIDItemsReader is a Reader for the GetCharactersCharacterIDContractsContractIDItems structure.
type GetCharactersCharacterIDContractsContractIDItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContractsContractIDItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDContractsContractIDItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDContractsContractIDItemsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDContractsContractIDItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDContractsContractIDItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDContractsContractIDItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCharactersCharacterIDContractsContractIDItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDContractsContractIDItemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDContractsContractIDItemsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContractsContractIDItemsOK creates a GetCharactersCharacterIDContractsContractIDItemsOK with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsOK() *GetCharactersCharacterIDContractsContractIDItemsOK {
	return &GetCharactersCharacterIDContractsContractIDItemsOK{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsOK describes a response with status code 200, with default header values.

A list of items in this contract
*/
type GetCharactersCharacterIDContractsContractIDItemsOK struct {

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

	Payload []*GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0
}

// IsSuccess returns true when this get characters character Id contracts contract Id items o k response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id contracts contract Id items o k response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items o k response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contracts contract Id items o k response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contracts contract Id items o k response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) GetPayload() []*GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContractsContractIDItemsNotModified creates a GetCharactersCharacterIDContractsContractIDItemsNotModified with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsNotModified() *GetCharactersCharacterIDContractsContractIDItemsNotModified {
	return &GetCharactersCharacterIDContractsContractIDItemsNotModified{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDContractsContractIDItemsNotModified struct {

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

// IsSuccess returns true when this get characters character Id contracts contract Id items not modified response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items not modified response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id contracts contract Id items not modified response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contracts contract Id items not modified response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contracts contract Id items not modified response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsNotModified ", 304)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsNotModified ", 304)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContractsContractIDItemsBadRequest creates a GetCharactersCharacterIDContractsContractIDItemsBadRequest with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsBadRequest() *GetCharactersCharacterIDContractsContractIDItemsBadRequest {
	return &GetCharactersCharacterIDContractsContractIDItemsBadRequest{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDContractsContractIDItemsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id contracts contract Id items bad request response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items bad request response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items bad request response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contracts contract Id items bad request response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contracts contract Id items bad request response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsUnauthorized creates a GetCharactersCharacterIDContractsContractIDItemsUnauthorized with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsUnauthorized() *GetCharactersCharacterIDContractsContractIDItemsUnauthorized {
	return &GetCharactersCharacterIDContractsContractIDItemsUnauthorized{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDContractsContractIDItemsUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id contracts contract Id items unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contracts contract Id items unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contracts contract Id items unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsForbidden creates a GetCharactersCharacterIDContractsContractIDItemsForbidden with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsForbidden() *GetCharactersCharacterIDContractsContractIDItemsForbidden {
	return &GetCharactersCharacterIDContractsContractIDItemsForbidden{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDContractsContractIDItemsForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id contracts contract Id items forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contracts contract Id items forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contracts contract Id items forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsNotFound creates a GetCharactersCharacterIDContractsContractIDItemsNotFound with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsNotFound() *GetCharactersCharacterIDContractsContractIDItemsNotFound {
	return &GetCharactersCharacterIDContractsContractIDItemsNotFound{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetCharactersCharacterIDContractsContractIDItemsNotFound struct {
	Payload *GetCharactersCharacterIDContractsContractIDItemsNotFoundBody
}

// IsSuccess returns true when this get characters character Id contracts contract Id items not found response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items not found response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items not found response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contracts contract Id items not found response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contracts contract Id items not found response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) GetPayload() *GetCharactersCharacterIDContractsContractIDItemsNotFoundBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCharactersCharacterIDContractsContractIDItemsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm creates a GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm() *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm {
	return &GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id contracts contract Id items enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id contracts contract Id items enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id contracts contract Id items enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError creates a GetCharactersCharacterIDContractsContractIDItemsInternalServerError with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError() *GetCharactersCharacterIDContractsContractIDItemsInternalServerError {
	return &GetCharactersCharacterIDContractsContractIDItemsInternalServerError{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDContractsContractIDItemsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id contracts contract Id items internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contracts contract Id items internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id contracts contract Id items internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsServiceUnavailable creates a GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsServiceUnavailable() *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable {
	return &GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id contracts contract Id items service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contracts contract Id items service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id contracts contract Id items service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsGatewayTimeout creates a GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsGatewayTimeout() *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout {
	return &GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout{}
}

/*
GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id contracts contract Id items gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id contracts contract Id items gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id contracts contract Id items gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id contracts contract Id items gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id contracts contract Id items gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCharactersCharacterIDContractsContractIDItemsNotFoundBody get_characters_character_id_contracts_contract_id_items_not_found
//
// Not found
swagger:model GetCharactersCharacterIDContractsContractIDItemsNotFoundBody
*/
type GetCharactersCharacterIDContractsContractIDItemsNotFoundBody struct {

	// get_characters_character_id_contracts_contract_id_items_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character ID contracts contract ID items not found body
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get characters character ID contracts contract ID items not found body based on context it is used
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsContractIDItemsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDContractsContractIDItemsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 get_characters_character_id_contracts_contract_id_items_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0
*/
type GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0 struct {

	// get_characters_character_id_contracts_contract_id_items_is_included
	//
	// true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract
	// Required: true
	IsIncluded *bool `json:"is_included"`

	// get_characters_character_id_contracts_contract_id_items_is_singleton
	//
	// is_singleton boolean
	// Required: true
	IsSingleton *bool `json:"is_singleton"`

	// get_characters_character_id_contracts_contract_id_items_quantity
	//
	// Number of items in the stack
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_characters_character_id_contracts_contract_id_items_raw_quantity
	//
	// -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy
	RawQuantity int32 `json:"raw_quantity,omitempty"`

	// get_characters_character_id_contracts_contract_id_items_record_id
	//
	// Unique ID for the item
	// Required: true
	RecordID *int64 `json:"record_id"`

	// get_characters_character_id_contracts_contract_id_items_type_id
	//
	// Type ID for item
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get characters character ID contracts contract ID items o k body items0
func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIsIncluded(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsSingleton(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecordID(formats); err != nil {
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

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateIsIncluded(formats strfmt.Registry) error {

	if err := validate.Required("is_included", "body", o.IsIncluded); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateIsSingleton(formats strfmt.Registry) error {

	if err := validate.Required("is_singleton", "body", o.IsSingleton); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateRecordID(formats strfmt.Registry) error {

	if err := validate.Required("record_id", "body", o.RecordID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID contracts contract ID items o k body items0 based on context it is used
func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDContractsContractIDItemsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
