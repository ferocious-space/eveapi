// Code generated by go-swagger; DO NOT EDIT.

package location

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

// GetCharactersCharacterIDShipReader is a Reader for the GetCharactersCharacterIDShip structure.
type GetCharactersCharacterIDShipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDShipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDShipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDShipNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDShipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDShipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDShipForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDShipEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDShipInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDShipServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDShipGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/characters/{character_id}/ship/] get_characters_character_id_ship", response, response.Code())
	}
}

// NewGetCharactersCharacterIDShipOK creates a GetCharactersCharacterIDShipOK with default headers values
func NewGetCharactersCharacterIDShipOK() *GetCharactersCharacterIDShipOK {
	return &GetCharactersCharacterIDShipOK{}
}

/*
GetCharactersCharacterIDShipOK describes a response with status code 200, with default header values.

Get the current ship type, name and id
*/
type GetCharactersCharacterIDShipOK struct {

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

	Payload *GetCharactersCharacterIDShipOKBody
}

// IsSuccess returns true when this get characters character Id ship o k response has a 2xx status code
func (o *GetCharactersCharacterIDShipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id ship o k response has a 3xx status code
func (o *GetCharactersCharacterIDShipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship o k response has a 4xx status code
func (o *GetCharactersCharacterIDShipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id ship o k response has a 5xx status code
func (o *GetCharactersCharacterIDShipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id ship o k response a status code equal to that given
func (o *GetCharactersCharacterIDShipOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get characters character Id ship o k response
func (o *GetCharactersCharacterIDShipOK) Code() int {
	return 200
}

func (o *GetCharactersCharacterIDShipOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDShipOK) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDShipOK) GetPayload() *GetCharactersCharacterIDShipOKBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCharactersCharacterIDShipOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipNotModified creates a GetCharactersCharacterIDShipNotModified with default headers values
func NewGetCharactersCharacterIDShipNotModified() *GetCharactersCharacterIDShipNotModified {
	return &GetCharactersCharacterIDShipNotModified{}
}

/*
GetCharactersCharacterIDShipNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDShipNotModified struct {

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

// IsSuccess returns true when this get characters character Id ship not modified response has a 2xx status code
func (o *GetCharactersCharacterIDShipNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship not modified response has a 3xx status code
func (o *GetCharactersCharacterIDShipNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id ship not modified response has a 4xx status code
func (o *GetCharactersCharacterIDShipNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id ship not modified response has a 5xx status code
func (o *GetCharactersCharacterIDShipNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id ship not modified response a status code equal to that given
func (o *GetCharactersCharacterIDShipNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get characters character Id ship not modified response
func (o *GetCharactersCharacterIDShipNotModified) Code() int {
	return 304
}

func (o *GetCharactersCharacterIDShipNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipNotModified ", 304)
}

func (o *GetCharactersCharacterIDShipNotModified) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipNotModified ", 304)
}

func (o *GetCharactersCharacterIDShipNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDShipBadRequest creates a GetCharactersCharacterIDShipBadRequest with default headers values
func NewGetCharactersCharacterIDShipBadRequest() *GetCharactersCharacterIDShipBadRequest {
	return &GetCharactersCharacterIDShipBadRequest{}
}

/*
GetCharactersCharacterIDShipBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDShipBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id ship bad request response has a 2xx status code
func (o *GetCharactersCharacterIDShipBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship bad request response has a 3xx status code
func (o *GetCharactersCharacterIDShipBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship bad request response has a 4xx status code
func (o *GetCharactersCharacterIDShipBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id ship bad request response has a 5xx status code
func (o *GetCharactersCharacterIDShipBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id ship bad request response a status code equal to that given
func (o *GetCharactersCharacterIDShipBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get characters character Id ship bad request response
func (o *GetCharactersCharacterIDShipBadRequest) Code() int {
	return 400
}

func (o *GetCharactersCharacterIDShipBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDShipBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDShipBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipUnauthorized creates a GetCharactersCharacterIDShipUnauthorized with default headers values
func NewGetCharactersCharacterIDShipUnauthorized() *GetCharactersCharacterIDShipUnauthorized {
	return &GetCharactersCharacterIDShipUnauthorized{}
}

/*
GetCharactersCharacterIDShipUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDShipUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id ship unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDShipUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDShipUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDShipUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id ship unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDShipUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id ship unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDShipUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get characters character Id ship unauthorized response
func (o *GetCharactersCharacterIDShipUnauthorized) Code() int {
	return 401
}

func (o *GetCharactersCharacterIDShipUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDShipUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDShipUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipForbidden creates a GetCharactersCharacterIDShipForbidden with default headers values
func NewGetCharactersCharacterIDShipForbidden() *GetCharactersCharacterIDShipForbidden {
	return &GetCharactersCharacterIDShipForbidden{}
}

/*
GetCharactersCharacterIDShipForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDShipForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id ship forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDShipForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDShipForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDShipForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id ship forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDShipForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id ship forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDShipForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get characters character Id ship forbidden response
func (o *GetCharactersCharacterIDShipForbidden) Code() int {
	return 403
}

func (o *GetCharactersCharacterIDShipForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDShipForbidden) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDShipForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipEnhanceYourCalm creates a GetCharactersCharacterIDShipEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDShipEnhanceYourCalm() *GetCharactersCharacterIDShipEnhanceYourCalm {
	return &GetCharactersCharacterIDShipEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDShipEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDShipEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id ship enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDShipEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDShipEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDShipEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id ship enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDShipEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id ship enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDShipEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get characters character Id ship enhance your calm response
func (o *GetCharactersCharacterIDShipEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCharactersCharacterIDShipEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDShipEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDShipEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipInternalServerError creates a GetCharactersCharacterIDShipInternalServerError with default headers values
func NewGetCharactersCharacterIDShipInternalServerError() *GetCharactersCharacterIDShipInternalServerError {
	return &GetCharactersCharacterIDShipInternalServerError{}
}

/*
GetCharactersCharacterIDShipInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDShipInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id ship internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDShipInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDShipInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDShipInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id ship internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDShipInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id ship internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDShipInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get characters character Id ship internal server error response
func (o *GetCharactersCharacterIDShipInternalServerError) Code() int {
	return 500
}

func (o *GetCharactersCharacterIDShipInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDShipInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDShipInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipServiceUnavailable creates a GetCharactersCharacterIDShipServiceUnavailable with default headers values
func NewGetCharactersCharacterIDShipServiceUnavailable() *GetCharactersCharacterIDShipServiceUnavailable {
	return &GetCharactersCharacterIDShipServiceUnavailable{}
}

/*
GetCharactersCharacterIDShipServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDShipServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id ship service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDShipServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDShipServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDShipServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id ship service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDShipServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id ship service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDShipServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get characters character Id ship service unavailable response
func (o *GetCharactersCharacterIDShipServiceUnavailable) Code() int {
	return 503
}

func (o *GetCharactersCharacterIDShipServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDShipServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDShipServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDShipGatewayTimeout creates a GetCharactersCharacterIDShipGatewayTimeout with default headers values
func NewGetCharactersCharacterIDShipGatewayTimeout() *GetCharactersCharacterIDShipGatewayTimeout {
	return &GetCharactersCharacterIDShipGatewayTimeout{}
}

/*
GetCharactersCharacterIDShipGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDShipGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id ship gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDShipGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id ship gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDShipGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id ship gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDShipGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id ship gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDShipGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id ship gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDShipGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get characters character Id ship gateway timeout response
func (o *GetCharactersCharacterIDShipGatewayTimeout) Code() int {
	return 504
}

func (o *GetCharactersCharacterIDShipGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDShipGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/ship/][%d] getCharactersCharacterIdShipGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDShipGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDShipGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCharactersCharacterIDShipOKBody get_characters_character_id_ship_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDShipOKBody
*/
type GetCharactersCharacterIDShipOKBody struct {

	// get_characters_character_id_ship_ship_item_id
	//
	// Item id's are unique to a ship and persist until it is repackaged. This value can be used to track repeated uses of a ship, or detect when a pilot changes into a different instance of the same ship type.
	// Required: true
	ShipItemID *int64 `json:"ship_item_id"`

	// get_characters_character_id_ship_ship_name
	//
	// ship_name string
	// Required: true
	ShipName *string `json:"ship_name"`

	// get_characters_character_id_ship_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`
}

// Validate validates this get characters character ID ship o k body
func (o *GetCharactersCharacterIDShipOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateShipItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDShipOKBody) validateShipItemID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdShipOK"+"."+"ship_item_id", "body", o.ShipItemID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDShipOKBody) validateShipName(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdShipOK"+"."+"ship_name", "body", o.ShipName); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDShipOKBody) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getCharactersCharacterIdShipOK"+"."+"ship_type_id", "body", o.ShipTypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID ship o k body based on context it is used
func (o *GetCharactersCharacterIDShipOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDShipOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDShipOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDShipOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
