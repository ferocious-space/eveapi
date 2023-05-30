// Code generated by go-swagger; DO NOT EDIT.

package loyalty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// GetLoyaltyStoresCorporationIDOffersReader is a Reader for the GetLoyaltyStoresCorporationIDOffers structure.
type GetLoyaltyStoresCorporationIDOffersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoyaltyStoresCorporationIDOffersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoyaltyStoresCorporationIDOffersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetLoyaltyStoresCorporationIDOffersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetLoyaltyStoresCorporationIDOffersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLoyaltyStoresCorporationIDOffersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetLoyaltyStoresCorporationIDOffersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLoyaltyStoresCorporationIDOffersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLoyaltyStoresCorporationIDOffersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLoyaltyStoresCorporationIDOffersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLoyaltyStoresCorporationIDOffersOK creates a GetLoyaltyStoresCorporationIDOffersOK with default headers values
func NewGetLoyaltyStoresCorporationIDOffersOK() *GetLoyaltyStoresCorporationIDOffersOK {
	return &GetLoyaltyStoresCorporationIDOffersOK{}
}

/*
GetLoyaltyStoresCorporationIDOffersOK describes a response with status code 200, with default header values.

A list of offers
*/
type GetLoyaltyStoresCorporationIDOffersOK struct {

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

	Payload []*GetLoyaltyStoresCorporationIDOffersOKBodyItems0
}

// IsSuccess returns true when this get loyalty stores corporation Id offers o k response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get loyalty stores corporation Id offers o k response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get loyalty stores corporation Id offers o k response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get loyalty stores corporation Id offers o k response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get loyalty stores corporation Id offers o k response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get loyalty stores corporation Id offers o k response
func (o *GetLoyaltyStoresCorporationIDOffersOK) Code() int {
	return 200
}

func (o *GetLoyaltyStoresCorporationIDOffersOK) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersOK  %+v", 200, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersOK) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersOK  %+v", 200, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersOK) GetPayload() []*GetLoyaltyStoresCorporationIDOffersOKBodyItems0 {
	return o.Payload
}

func (o *GetLoyaltyStoresCorporationIDOffersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLoyaltyStoresCorporationIDOffersNotModified creates a GetLoyaltyStoresCorporationIDOffersNotModified with default headers values
func NewGetLoyaltyStoresCorporationIDOffersNotModified() *GetLoyaltyStoresCorporationIDOffersNotModified {
	return &GetLoyaltyStoresCorporationIDOffersNotModified{}
}

/*
GetLoyaltyStoresCorporationIDOffersNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetLoyaltyStoresCorporationIDOffersNotModified struct {

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

// IsSuccess returns true when this get loyalty stores corporation Id offers not modified response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get loyalty stores corporation Id offers not modified response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get loyalty stores corporation Id offers not modified response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get loyalty stores corporation Id offers not modified response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get loyalty stores corporation Id offers not modified response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get loyalty stores corporation Id offers not modified response
func (o *GetLoyaltyStoresCorporationIDOffersNotModified) Code() int {
	return 304
}

func (o *GetLoyaltyStoresCorporationIDOffersNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersNotModified ", 304)
}

func (o *GetLoyaltyStoresCorporationIDOffersNotModified) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersNotModified ", 304)
}

func (o *GetLoyaltyStoresCorporationIDOffersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLoyaltyStoresCorporationIDOffersBadRequest creates a GetLoyaltyStoresCorporationIDOffersBadRequest with default headers values
func NewGetLoyaltyStoresCorporationIDOffersBadRequest() *GetLoyaltyStoresCorporationIDOffersBadRequest {
	return &GetLoyaltyStoresCorporationIDOffersBadRequest{}
}

/*
GetLoyaltyStoresCorporationIDOffersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetLoyaltyStoresCorporationIDOffersBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get loyalty stores corporation Id offers bad request response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get loyalty stores corporation Id offers bad request response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get loyalty stores corporation Id offers bad request response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get loyalty stores corporation Id offers bad request response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get loyalty stores corporation Id offers bad request response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get loyalty stores corporation Id offers bad request response
func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) Code() int {
	return 400
}

func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersBadRequest  %+v", 400, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersBadRequest  %+v", 400, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetLoyaltyStoresCorporationIDOffersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoyaltyStoresCorporationIDOffersNotFound creates a GetLoyaltyStoresCorporationIDOffersNotFound with default headers values
func NewGetLoyaltyStoresCorporationIDOffersNotFound() *GetLoyaltyStoresCorporationIDOffersNotFound {
	return &GetLoyaltyStoresCorporationIDOffersNotFound{}
}

/*
GetLoyaltyStoresCorporationIDOffersNotFound describes a response with status code 404, with default header values.

No loyalty point store found for the provided corporation
*/
type GetLoyaltyStoresCorporationIDOffersNotFound struct {
	Payload *GetLoyaltyStoresCorporationIDOffersNotFoundBody
}

// IsSuccess returns true when this get loyalty stores corporation Id offers not found response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get loyalty stores corporation Id offers not found response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get loyalty stores corporation Id offers not found response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get loyalty stores corporation Id offers not found response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get loyalty stores corporation Id offers not found response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get loyalty stores corporation Id offers not found response
func (o *GetLoyaltyStoresCorporationIDOffersNotFound) Code() int {
	return 404
}

func (o *GetLoyaltyStoresCorporationIDOffersNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersNotFound  %+v", 404, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersNotFound) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersNotFound  %+v", 404, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersNotFound) GetPayload() *GetLoyaltyStoresCorporationIDOffersNotFoundBody {
	return o.Payload
}

func (o *GetLoyaltyStoresCorporationIDOffersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetLoyaltyStoresCorporationIDOffersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoyaltyStoresCorporationIDOffersEnhanceYourCalm creates a GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm with default headers values
func NewGetLoyaltyStoresCorporationIDOffersEnhanceYourCalm() *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm {
	return &GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm{}
}

/*
GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get loyalty stores corporation Id offers enhance your calm response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get loyalty stores corporation Id offers enhance your calm response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get loyalty stores corporation Id offers enhance your calm response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get loyalty stores corporation Id offers enhance your calm response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get loyalty stores corporation Id offers enhance your calm response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get loyalty stores corporation Id offers enhance your calm response
func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetLoyaltyStoresCorporationIDOffersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoyaltyStoresCorporationIDOffersInternalServerError creates a GetLoyaltyStoresCorporationIDOffersInternalServerError with default headers values
func NewGetLoyaltyStoresCorporationIDOffersInternalServerError() *GetLoyaltyStoresCorporationIDOffersInternalServerError {
	return &GetLoyaltyStoresCorporationIDOffersInternalServerError{}
}

/*
GetLoyaltyStoresCorporationIDOffersInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLoyaltyStoresCorporationIDOffersInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get loyalty stores corporation Id offers internal server error response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get loyalty stores corporation Id offers internal server error response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get loyalty stores corporation Id offers internal server error response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get loyalty stores corporation Id offers internal server error response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get loyalty stores corporation Id offers internal server error response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get loyalty stores corporation Id offers internal server error response
func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) Code() int {
	return 500
}

func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetLoyaltyStoresCorporationIDOffersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoyaltyStoresCorporationIDOffersServiceUnavailable creates a GetLoyaltyStoresCorporationIDOffersServiceUnavailable with default headers values
func NewGetLoyaltyStoresCorporationIDOffersServiceUnavailable() *GetLoyaltyStoresCorporationIDOffersServiceUnavailable {
	return &GetLoyaltyStoresCorporationIDOffersServiceUnavailable{}
}

/*
GetLoyaltyStoresCorporationIDOffersServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetLoyaltyStoresCorporationIDOffersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get loyalty stores corporation Id offers service unavailable response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get loyalty stores corporation Id offers service unavailable response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get loyalty stores corporation Id offers service unavailable response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get loyalty stores corporation Id offers service unavailable response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get loyalty stores corporation Id offers service unavailable response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get loyalty stores corporation Id offers service unavailable response
func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) Code() int {
	return 503
}

func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetLoyaltyStoresCorporationIDOffersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoyaltyStoresCorporationIDOffersGatewayTimeout creates a GetLoyaltyStoresCorporationIDOffersGatewayTimeout with default headers values
func NewGetLoyaltyStoresCorporationIDOffersGatewayTimeout() *GetLoyaltyStoresCorporationIDOffersGatewayTimeout {
	return &GetLoyaltyStoresCorporationIDOffersGatewayTimeout{}
}

/*
GetLoyaltyStoresCorporationIDOffersGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetLoyaltyStoresCorporationIDOffersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get loyalty stores corporation Id offers gateway timeout response has a 2xx status code
func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get loyalty stores corporation Id offers gateway timeout response has a 3xx status code
func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get loyalty stores corporation Id offers gateway timeout response has a 4xx status code
func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get loyalty stores corporation Id offers gateway timeout response has a 5xx status code
func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get loyalty stores corporation Id offers gateway timeout response a status code equal to that given
func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get loyalty stores corporation Id offers gateway timeout response
func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) Code() int {
	return 504
}

func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/loyalty/stores/{corporation_id}/offers/][%d] getLoyaltyStoresCorporationIdOffersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetLoyaltyStoresCorporationIDOffersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetLoyaltyStoresCorporationIDOffersNotFoundBody get_loyalty_stores_corporation_id_offers_not_found
//
// Not found
swagger:model GetLoyaltyStoresCorporationIDOffersNotFoundBody
*/
type GetLoyaltyStoresCorporationIDOffersNotFoundBody struct {

	// get_loyalty_stores_corporation_id_offers_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get loyalty stores corporation ID offers not found body
func (o *GetLoyaltyStoresCorporationIDOffersNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get loyalty stores corporation ID offers not found body based on context it is used
func (o *GetLoyaltyStoresCorporationIDOffersNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLoyaltyStoresCorporationIDOffersNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLoyaltyStoresCorporationIDOffersNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetLoyaltyStoresCorporationIDOffersNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLoyaltyStoresCorporationIDOffersOKBodyItems0 get_loyalty_stores_corporation_id_offers_200_ok
//
// 200 ok object
swagger:model GetLoyaltyStoresCorporationIDOffersOKBodyItems0
*/
type GetLoyaltyStoresCorporationIDOffersOKBodyItems0 struct {

	// get_loyalty_stores_corporation_id_offers_ak_cost
	//
	// Analysis kredit cost
	AkCost int32 `json:"ak_cost,omitempty"`

	// get_loyalty_stores_corporation_id_offers_isk_cost
	//
	// isk_cost integer
	// Required: true
	IskCost *int64 `json:"isk_cost"`

	// get_loyalty_stores_corporation_id_offers_lp_cost
	//
	// lp_cost integer
	// Required: true
	LpCost *int32 `json:"lp_cost"`

	// get_loyalty_stores_corporation_id_offers_offer_id
	//
	// offer_id integer
	// Required: true
	OfferID *int32 `json:"offer_id"`

	// get_loyalty_stores_corporation_id_offers_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_loyalty_stores_corporation_id_offers_required_items
	//
	// required_items array
	// Required: true
	// Max Items: 100
	RequiredItems []*GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0 `json:"required_items"`

	// get_loyalty_stores_corporation_id_offers_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get loyalty stores corporation ID offers o k body items0
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIskCost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLpCost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOfferID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRequiredItems(formats); err != nil {
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

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) validateIskCost(formats strfmt.Registry) error {

	if err := validate.Required("isk_cost", "body", o.IskCost); err != nil {
		return err
	}

	return nil
}

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) validateLpCost(formats strfmt.Registry) error {

	if err := validate.Required("lp_cost", "body", o.LpCost); err != nil {
		return err
	}

	return nil
}

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) validateOfferID(formats strfmt.Registry) error {

	if err := validate.Required("offer_id", "body", o.OfferID); err != nil {
		return err
	}

	return nil
}

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) validateRequiredItems(formats strfmt.Registry) error {

	if err := validate.Required("required_items", "body", o.RequiredItems); err != nil {
		return err
	}

	iRequiredItemsSize := int64(len(o.RequiredItems))

	if err := validate.MaxItems("required_items", "body", iRequiredItemsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(o.RequiredItems); i++ {
		if swag.IsZero(o.RequiredItems[i]) { // not required
			continue
		}

		if o.RequiredItems[i] != nil {
			if err := o.RequiredItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("required_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get loyalty stores corporation ID offers o k body items0 based on the context it is used
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRequiredItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) contextValidateRequiredItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RequiredItems); i++ {

		if o.RequiredItems[i] != nil {
			if err := o.RequiredItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("required_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetLoyaltyStoresCorporationIDOffersOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0 get_loyalty_stores_corporation_id_offers_required_item
//
// required_item object
swagger:model GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0
*/
type GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0 struct {

	// get_loyalty_stores_corporation_id_offers_required_item_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_loyalty_stores_corporation_id_offers_required_item_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get loyalty stores corporation ID offers o k body items0 required items items0
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQuantity(formats); err != nil {
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

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get loyalty stores corporation ID offers o k body items0 required items items0 based on context it is used
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0) UnmarshalBinary(b []byte) error {
	var res GetLoyaltyStoresCorporationIDOffersOKBodyItems0RequiredItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
