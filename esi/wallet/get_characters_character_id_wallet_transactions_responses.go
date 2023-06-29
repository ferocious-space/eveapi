// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// GetCharactersCharacterIDWalletTransactionsReader is a Reader for the GetCharactersCharacterIDWalletTransactions structure.
type GetCharactersCharacterIDWalletTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDWalletTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDWalletTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDWalletTransactionsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDWalletTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDWalletTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDWalletTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDWalletTransactionsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDWalletTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDWalletTransactionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDWalletTransactionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/characters/{character_id}/wallet/transactions/] get_characters_character_id_wallet_transactions", response, response.Code())
	}
}

// NewGetCharactersCharacterIDWalletTransactionsOK creates a GetCharactersCharacterIDWalletTransactionsOK with default headers values
func NewGetCharactersCharacterIDWalletTransactionsOK() *GetCharactersCharacterIDWalletTransactionsOK {
	return &GetCharactersCharacterIDWalletTransactionsOK{}
}

/*
GetCharactersCharacterIDWalletTransactionsOK describes a response with status code 200, with default header values.

Wallet transactions
*/
type GetCharactersCharacterIDWalletTransactionsOK struct {

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

	Payload []*GetCharactersCharacterIDWalletTransactionsOKBodyItems0
}

// IsSuccess returns true when this get characters character Id wallet transactions o k response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get characters character Id wallet transactions o k response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions o k response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet transactions o k response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet transactions o k response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get characters character Id wallet transactions o k response
func (o *GetCharactersCharacterIDWalletTransactionsOK) Code() int {
	return 200
}

func (o *GetCharactersCharacterIDWalletTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsOK) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsOK) GetPayload() []*GetCharactersCharacterIDWalletTransactionsOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDWalletTransactionsNotModified creates a GetCharactersCharacterIDWalletTransactionsNotModified with default headers values
func NewGetCharactersCharacterIDWalletTransactionsNotModified() *GetCharactersCharacterIDWalletTransactionsNotModified {
	return &GetCharactersCharacterIDWalletTransactionsNotModified{}
}

/*
GetCharactersCharacterIDWalletTransactionsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDWalletTransactionsNotModified struct {

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

// IsSuccess returns true when this get characters character Id wallet transactions not modified response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions not modified response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get characters character Id wallet transactions not modified response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet transactions not modified response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet transactions not modified response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get characters character Id wallet transactions not modified response
func (o *GetCharactersCharacterIDWalletTransactionsNotModified) Code() int {
	return 304
}

func (o *GetCharactersCharacterIDWalletTransactionsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsNotModified ", 304)
}

func (o *GetCharactersCharacterIDWalletTransactionsNotModified) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsNotModified ", 304)
}

func (o *GetCharactersCharacterIDWalletTransactionsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDWalletTransactionsBadRequest creates a GetCharactersCharacterIDWalletTransactionsBadRequest with default headers values
func NewGetCharactersCharacterIDWalletTransactionsBadRequest() *GetCharactersCharacterIDWalletTransactionsBadRequest {
	return &GetCharactersCharacterIDWalletTransactionsBadRequest{}
}

/*
GetCharactersCharacterIDWalletTransactionsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDWalletTransactionsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get characters character Id wallet transactions bad request response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions bad request response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions bad request response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet transactions bad request response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet transactions bad request response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get characters character Id wallet transactions bad request response
func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) Code() int {
	return 400
}

func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsUnauthorized creates a GetCharactersCharacterIDWalletTransactionsUnauthorized with default headers values
func NewGetCharactersCharacterIDWalletTransactionsUnauthorized() *GetCharactersCharacterIDWalletTransactionsUnauthorized {
	return &GetCharactersCharacterIDWalletTransactionsUnauthorized{}
}

/*
GetCharactersCharacterIDWalletTransactionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDWalletTransactionsUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this get characters character Id wallet transactions unauthorized response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions unauthorized response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions unauthorized response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet transactions unauthorized response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet transactions unauthorized response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get characters character Id wallet transactions unauthorized response
func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) Code() int {
	return 401
}

func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsForbidden creates a GetCharactersCharacterIDWalletTransactionsForbidden with default headers values
func NewGetCharactersCharacterIDWalletTransactionsForbidden() *GetCharactersCharacterIDWalletTransactionsForbidden {
	return &GetCharactersCharacterIDWalletTransactionsForbidden{}
}

/*
GetCharactersCharacterIDWalletTransactionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDWalletTransactionsForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this get characters character Id wallet transactions forbidden response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions forbidden response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions forbidden response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet transactions forbidden response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet transactions forbidden response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get characters character Id wallet transactions forbidden response
func (o *GetCharactersCharacterIDWalletTransactionsForbidden) Code() int {
	return 403
}

func (o *GetCharactersCharacterIDWalletTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsEnhanceYourCalm creates a GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDWalletTransactionsEnhanceYourCalm() *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm {
	return &GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm{}
}

/*
GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get characters character Id wallet transactions enhance your calm response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions enhance your calm response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions enhance your calm response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get characters character Id wallet transactions enhance your calm response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get characters character Id wallet transactions enhance your calm response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get characters character Id wallet transactions enhance your calm response
func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsInternalServerError creates a GetCharactersCharacterIDWalletTransactionsInternalServerError with default headers values
func NewGetCharactersCharacterIDWalletTransactionsInternalServerError() *GetCharactersCharacterIDWalletTransactionsInternalServerError {
	return &GetCharactersCharacterIDWalletTransactionsInternalServerError{}
}

/*
GetCharactersCharacterIDWalletTransactionsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDWalletTransactionsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get characters character Id wallet transactions internal server error response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions internal server error response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions internal server error response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet transactions internal server error response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id wallet transactions internal server error response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get characters character Id wallet transactions internal server error response
func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) Code() int {
	return 500
}

func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsServiceUnavailable creates a GetCharactersCharacterIDWalletTransactionsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDWalletTransactionsServiceUnavailable() *GetCharactersCharacterIDWalletTransactionsServiceUnavailable {
	return &GetCharactersCharacterIDWalletTransactionsServiceUnavailable{}
}

/*
GetCharactersCharacterIDWalletTransactionsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDWalletTransactionsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get characters character Id wallet transactions service unavailable response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions service unavailable response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions service unavailable response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet transactions service unavailable response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id wallet transactions service unavailable response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get characters character Id wallet transactions service unavailable response
func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) Code() int {
	return 503
}

func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsGatewayTimeout creates a GetCharactersCharacterIDWalletTransactionsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDWalletTransactionsGatewayTimeout() *GetCharactersCharacterIDWalletTransactionsGatewayTimeout {
	return &GetCharactersCharacterIDWalletTransactionsGatewayTimeout{}
}

/*
GetCharactersCharacterIDWalletTransactionsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDWalletTransactionsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get characters character Id wallet transactions gateway timeout response has a 2xx status code
func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get characters character Id wallet transactions gateway timeout response has a 3xx status code
func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get characters character Id wallet transactions gateway timeout response has a 4xx status code
func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get characters character Id wallet transactions gateway timeout response has a 5xx status code
func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get characters character Id wallet transactions gateway timeout response a status code equal to that given
func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get characters character Id wallet transactions gateway timeout response
func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) Code() int {
	return 504
}

func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCharactersCharacterIDWalletTransactionsOKBodyItems0 get_characters_character_id_wallet_transactions_200_ok
//
// wallet transaction
swagger:model GetCharactersCharacterIDWalletTransactionsOKBodyItems0
*/
type GetCharactersCharacterIDWalletTransactionsOKBodyItems0 struct {

	// get_characters_character_id_wallet_transactions_client_id
	//
	// client_id integer
	// Required: true
	ClientID *int32 `json:"client_id"`

	// get_characters_character_id_wallet_transactions_date
	//
	// Date and time of transaction
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// get_characters_character_id_wallet_transactions_is_buy
	//
	// is_buy boolean
	// Required: true
	IsBuy *bool `json:"is_buy"`

	// get_characters_character_id_wallet_transactions_is_personal
	//
	// is_personal boolean
	// Required: true
	IsPersonal *bool `json:"is_personal"`

	// get_characters_character_id_wallet_transactions_journal_ref_id
	//
	// journal_ref_id integer
	// Required: true
	JournalRefID *int64 `json:"journal_ref_id"`

	// get_characters_character_id_wallet_transactions_location_id
	//
	// location_id integer
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_characters_character_id_wallet_transactions_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_characters_character_id_wallet_transactions_transaction_id
	//
	// Unique transaction ID
	// Required: true
	TransactionID *int64 `json:"transaction_id"`

	// get_characters_character_id_wallet_transactions_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_characters_character_id_wallet_transactions_unit_price
	//
	// Amount paid per unit
	// Required: true
	UnitPrice *float64 `json:"unit_price"`
}

// Validate validates this get characters character ID wallet transactions o k body items0
func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsBuy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsPersonal(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateJournalRefID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnitPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", o.ClientID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", o.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", o.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateIsBuy(formats strfmt.Registry) error {

	if err := validate.Required("is_buy", "body", o.IsBuy); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateIsPersonal(formats strfmt.Registry) error {

	if err := validate.Required("is_personal", "body", o.IsPersonal); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateJournalRefID(formats strfmt.Registry) error {

	if err := validate.Required("journal_ref_id", "body", o.JournalRefID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transaction_id", "body", o.TransactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) validateUnitPrice(formats strfmt.Registry) error {

	if err := validate.Required("unit_price", "body", o.UnitPrice); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID wallet transactions o k body items0 based on context it is used
func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDWalletTransactionsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDWalletTransactionsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
