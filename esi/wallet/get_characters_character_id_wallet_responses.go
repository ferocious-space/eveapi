// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/models"
)

// GetCharactersCharacterIDWalletReader is a Reader for the GetCharactersCharacterIDWallet structure.
type GetCharactersCharacterIDWalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDWalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDWalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDWalletNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDWalletBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDWalletUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDWalletForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDWalletEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDWalletInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDWalletServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDWalletGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDWalletOK creates a GetCharactersCharacterIDWalletOK with default headers values
func NewGetCharactersCharacterIDWalletOK() *GetCharactersCharacterIDWalletOK {
	return &GetCharactersCharacterIDWalletOK{}
}

/*
GetCharactersCharacterIDWalletOK describes a response with status code 200, with default header values.

Wallet balance
*/
type GetCharactersCharacterIDWalletOK struct {

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

	Payload float64
}

// IsSuccess returns true when this get characters character Id wallet o k response has a 2xx status code
func (o *GetCharactersCharacterIDWalletOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id wallet o k response has a 3xx status code
func (o *GetCharactersCharacterIDWalletOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet o k response has a 4xx status code
func (o *GetCharactersCharacterIDWalletOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet o k response has a 5xx status code
func (o *GetCharactersCharacterIDWalletOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet o k response a status code equal to that given
func (o *GetCharactersCharacterIDWalletOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get characters character Id wallet o k response
func (o *GetCharactersCharacterIDWalletOK) Code() int {
	return 200
}

func (o *GetCharactersCharacterIDWalletOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDWalletOK) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDWalletOK) GetPayload() float64 {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDWalletNotModified creates a GetCharactersCharacterIDWalletNotModified with default headers values
func NewGetCharactersCharacterIDWalletNotModified() *GetCharactersCharacterIDWalletNotModified {
	return &GetCharactersCharacterIDWalletNotModified{}
}

/*
GetCharactersCharacterIDWalletNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDWalletNotModified struct {

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

// IsSuccess returns true when this get characters character Id wallet not modified response has a 2xx status code
func (o *GetCharactersCharacterIDWalletNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet not modified response has a 3xx status code
func (o *GetCharactersCharacterIDWalletNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id wallet not modified response has a 4xx status code
func (o *GetCharactersCharacterIDWalletNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet not modified response has a 5xx status code
func (o *GetCharactersCharacterIDWalletNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet not modified response a status code equal to that given
func (o *GetCharactersCharacterIDWalletNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get characters character Id wallet not modified response
func (o *GetCharactersCharacterIDWalletNotModified) Code() int {
	return 304
}

func (o *GetCharactersCharacterIDWalletNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletNotModified ", 304)
}

func (o *GetCharactersCharacterIDWalletNotModified) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletNotModified ", 304)
}

func (o *GetCharactersCharacterIDWalletNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDWalletBadRequest creates a GetCharactersCharacterIDWalletBadRequest with default headers values
func NewGetCharactersCharacterIDWalletBadRequest() *GetCharactersCharacterIDWalletBadRequest {
	return &GetCharactersCharacterIDWalletBadRequest{}
}

/*
GetCharactersCharacterIDWalletBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDWalletBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id wallet bad request response has a 2xx status code
func (o *GetCharactersCharacterIDWalletBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet bad request response has a 3xx status code
func (o *GetCharactersCharacterIDWalletBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet bad request response has a 4xx status code
func (o *GetCharactersCharacterIDWalletBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet bad request response has a 5xx status code
func (o *GetCharactersCharacterIDWalletBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet bad request response a status code equal to that given
func (o *GetCharactersCharacterIDWalletBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get characters character Id wallet bad request response
func (o *GetCharactersCharacterIDWalletBadRequest) Code() int {
	return 400
}

func (o *GetCharactersCharacterIDWalletBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDWalletBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDWalletBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletUnauthorized creates a GetCharactersCharacterIDWalletUnauthorized with default headers values
func NewGetCharactersCharacterIDWalletUnauthorized() *GetCharactersCharacterIDWalletUnauthorized {
	return &GetCharactersCharacterIDWalletUnauthorized{}
}

/*
GetCharactersCharacterIDWalletUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDWalletUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id wallet unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDWalletUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDWalletUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDWalletUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDWalletUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDWalletUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get characters character Id wallet unauthorized response
func (o *GetCharactersCharacterIDWalletUnauthorized) Code() int {
	return 401
}

func (o *GetCharactersCharacterIDWalletUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDWalletUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDWalletUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletForbidden creates a GetCharactersCharacterIDWalletForbidden with default headers values
func NewGetCharactersCharacterIDWalletForbidden() *GetCharactersCharacterIDWalletForbidden {
	return &GetCharactersCharacterIDWalletForbidden{}
}

/*
GetCharactersCharacterIDWalletForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDWalletForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id wallet forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDWalletForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDWalletForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDWalletForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDWalletForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDWalletForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get characters character Id wallet forbidden response
func (o *GetCharactersCharacterIDWalletForbidden) Code() int {
	return 403
}

func (o *GetCharactersCharacterIDWalletForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDWalletForbidden) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDWalletForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletEnhanceYourCalm creates a GetCharactersCharacterIDWalletEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDWalletEnhanceYourCalm() *GetCharactersCharacterIDWalletEnhanceYourCalm {
	return &GetCharactersCharacterIDWalletEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDWalletEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDWalletEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id wallet enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get characters character Id wallet enhance your calm response
func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletInternalServerError creates a GetCharactersCharacterIDWalletInternalServerError with default headers values
func NewGetCharactersCharacterIDWalletInternalServerError() *GetCharactersCharacterIDWalletInternalServerError {
	return &GetCharactersCharacterIDWalletInternalServerError{}
}

/*
GetCharactersCharacterIDWalletInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDWalletInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id wallet internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDWalletInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDWalletInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDWalletInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDWalletInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id wallet internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDWalletInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get characters character Id wallet internal server error response
func (o *GetCharactersCharacterIDWalletInternalServerError) Code() int {
	return 500
}

func (o *GetCharactersCharacterIDWalletInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDWalletInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDWalletInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletServiceUnavailable creates a GetCharactersCharacterIDWalletServiceUnavailable with default headers values
func NewGetCharactersCharacterIDWalletServiceUnavailable() *GetCharactersCharacterIDWalletServiceUnavailable {
	return &GetCharactersCharacterIDWalletServiceUnavailable{}
}

/*
GetCharactersCharacterIDWalletServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDWalletServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id wallet service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDWalletServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDWalletServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDWalletServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDWalletServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id wallet service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDWalletServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get characters character Id wallet service unavailable response
func (o *GetCharactersCharacterIDWalletServiceUnavailable) Code() int {
	return 503
}

func (o *GetCharactersCharacterIDWalletServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDWalletServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDWalletServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletGatewayTimeout creates a GetCharactersCharacterIDWalletGatewayTimeout with default headers values
func NewGetCharactersCharacterIDWalletGatewayTimeout() *GetCharactersCharacterIDWalletGatewayTimeout {
	return &GetCharactersCharacterIDWalletGatewayTimeout{}
}

/*
GetCharactersCharacterIDWalletGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDWalletGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id wallet gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDWalletGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDWalletGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDWalletGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDWalletGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id wallet gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDWalletGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get characters character Id wallet gateway timeout response
func (o *GetCharactersCharacterIDWalletGatewayTimeout) Code() int {
	return 504
}

func (o *GetCharactersCharacterIDWalletGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDWalletGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/][%d] getCharactersCharacterIdWalletGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDWalletGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
