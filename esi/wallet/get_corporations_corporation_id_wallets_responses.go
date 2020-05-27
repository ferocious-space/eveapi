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

// GetCorporationsCorporationIDWalletsReader is a Reader for the GetCorporationsCorporationIDWallets structure.
type GetCorporationsCorporationIDWalletsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDWalletsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDWalletsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDWalletsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDWalletsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDWalletsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDWalletsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDWalletsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDWalletsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDWalletsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDWalletsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDWalletsOK creates a GetCorporationsCorporationIDWalletsOK with default headers values
func NewGetCorporationsCorporationIDWalletsOK() *GetCorporationsCorporationIDWalletsOK {
	return &GetCorporationsCorporationIDWalletsOK{}
}

/* GetCorporationsCorporationIDWalletsOK describes a response with status code 200, with default header values.

List of corporation wallets
*/
type GetCorporationsCorporationIDWalletsOK struct {

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

	Payload []*GetCorporationsCorporationIDWalletsOKBodyItems0
}

func (o *GetCorporationsCorporationIDWalletsOK) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsOK) GetPayload() []*GetCorporationsCorporationIDWalletsOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDWalletsNotModified creates a GetCorporationsCorporationIDWalletsNotModified with default headers values
func NewGetCorporationsCorporationIDWalletsNotModified() *GetCorporationsCorporationIDWalletsNotModified {
	return &GetCorporationsCorporationIDWalletsNotModified{}
}

/* GetCorporationsCorporationIDWalletsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDWalletsNotModified struct {

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

func (o *GetCorporationsCorporationIDWalletsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDWalletsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDWalletsBadRequest creates a GetCorporationsCorporationIDWalletsBadRequest with default headers values
func NewGetCorporationsCorporationIDWalletsBadRequest() *GetCorporationsCorporationIDWalletsBadRequest {
	return &GetCorporationsCorporationIDWalletsBadRequest{}
}

/* GetCorporationsCorporationIDWalletsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDWalletsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDWalletsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsUnauthorized creates a GetCorporationsCorporationIDWalletsUnauthorized with default headers values
func NewGetCorporationsCorporationIDWalletsUnauthorized() *GetCorporationsCorporationIDWalletsUnauthorized {
	return &GetCorporationsCorporationIDWalletsUnauthorized{}
}

/* GetCorporationsCorporationIDWalletsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDWalletsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDWalletsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsForbidden creates a GetCorporationsCorporationIDWalletsForbidden with default headers values
func NewGetCorporationsCorporationIDWalletsForbidden() *GetCorporationsCorporationIDWalletsForbidden {
	return &GetCorporationsCorporationIDWalletsForbidden{}
}

/* GetCorporationsCorporationIDWalletsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDWalletsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDWalletsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsEnhanceYourCalm creates a GetCorporationsCorporationIDWalletsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDWalletsEnhanceYourCalm() *GetCorporationsCorporationIDWalletsEnhanceYourCalm {
	return &GetCorporationsCorporationIDWalletsEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDWalletsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDWalletsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDWalletsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsInternalServerError creates a GetCorporationsCorporationIDWalletsInternalServerError with default headers values
func NewGetCorporationsCorporationIDWalletsInternalServerError() *GetCorporationsCorporationIDWalletsInternalServerError {
	return &GetCorporationsCorporationIDWalletsInternalServerError{}
}

/* GetCorporationsCorporationIDWalletsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDWalletsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDWalletsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsServiceUnavailable creates a GetCorporationsCorporationIDWalletsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDWalletsServiceUnavailable() *GetCorporationsCorporationIDWalletsServiceUnavailable {
	return &GetCorporationsCorporationIDWalletsServiceUnavailable{}
}

/* GetCorporationsCorporationIDWalletsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDWalletsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDWalletsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsGatewayTimeout creates a GetCorporationsCorporationIDWalletsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDWalletsGatewayTimeout() *GetCorporationsCorporationIDWalletsGatewayTimeout {
	return &GetCorporationsCorporationIDWalletsGatewayTimeout{}
}

/* GetCorporationsCorporationIDWalletsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDWalletsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDWalletsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/][%d] getCorporationsCorporationIdWalletsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDWalletsOKBodyItems0 get_corporations_corporation_id_wallets_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDWalletsOKBodyItems0
*/
type GetCorporationsCorporationIDWalletsOKBodyItems0 struct {

	// get_corporations_corporation_id_wallets_balance
	//
	// balance number
	// Required: true
	Balance *float64 `json:"balance"`

	// get_corporations_corporation_id_wallets_division
	//
	// division integer
	// Required: true
	// Maximum: 7
	// Minimum: 1
	Division *int32 `json:"division"`
}

// Validate validates this get corporations corporation ID wallets o k body items0
func (o *GetCorporationsCorporationIDWalletsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDWalletsOKBodyItems0) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", o.Balance); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsOKBodyItems0) validateDivision(formats strfmt.Registry) error {

	if err := validate.Required("division", "body", o.Division); err != nil {
		return err
	}

	if err := validate.MinimumInt("division", "body", int64(*o.Division), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("division", "body", int64(*o.Division), 7, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID wallets o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDWalletsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDWalletsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDWalletsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDWalletsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
