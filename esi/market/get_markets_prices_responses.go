// Code generated by go-swagger; DO NOT EDIT.

package market

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

// GetMarketsPricesReader is a Reader for the GetMarketsPrices structure.
type GetMarketsPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMarketsPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetMarketsPricesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetMarketsPricesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetMarketsPricesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMarketsPricesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMarketsPricesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetMarketsPricesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMarketsPricesOK creates a GetMarketsPricesOK with default headers values
func NewGetMarketsPricesOK() *GetMarketsPricesOK {
	return &GetMarketsPricesOK{}
}

/* GetMarketsPricesOK describes a response with status code 200, with default header values.

A list of prices
*/
type GetMarketsPricesOK struct {

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

	Payload []*GetMarketsPricesOKBodyItems0
}

func (o *GetMarketsPricesOK) Error() string {
	return fmt.Sprintf("[GET /v1/markets/prices/][%d] getMarketsPricesOK  %+v", 200, o.Payload)
}
func (o *GetMarketsPricesOK) GetPayload() []*GetMarketsPricesOKBodyItems0 {
	return o.Payload
}

func (o *GetMarketsPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsPricesNotModified creates a GetMarketsPricesNotModified with default headers values
func NewGetMarketsPricesNotModified() *GetMarketsPricesNotModified {
	return &GetMarketsPricesNotModified{}
}

/* GetMarketsPricesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetMarketsPricesNotModified struct {

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

func (o *GetMarketsPricesNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/markets/prices/][%d] getMarketsPricesNotModified ", 304)
}

func (o *GetMarketsPricesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsPricesBadRequest creates a GetMarketsPricesBadRequest with default headers values
func NewGetMarketsPricesBadRequest() *GetMarketsPricesBadRequest {
	return &GetMarketsPricesBadRequest{}
}

/* GetMarketsPricesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetMarketsPricesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetMarketsPricesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/markets/prices/][%d] getMarketsPricesBadRequest  %+v", 400, o.Payload)
}
func (o *GetMarketsPricesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetMarketsPricesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsPricesEnhanceYourCalm creates a GetMarketsPricesEnhanceYourCalm with default headers values
func NewGetMarketsPricesEnhanceYourCalm() *GetMarketsPricesEnhanceYourCalm {
	return &GetMarketsPricesEnhanceYourCalm{}
}

/* GetMarketsPricesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetMarketsPricesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetMarketsPricesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/markets/prices/][%d] getMarketsPricesEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetMarketsPricesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetMarketsPricesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsPricesInternalServerError creates a GetMarketsPricesInternalServerError with default headers values
func NewGetMarketsPricesInternalServerError() *GetMarketsPricesInternalServerError {
	return &GetMarketsPricesInternalServerError{}
}

/* GetMarketsPricesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetMarketsPricesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetMarketsPricesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/markets/prices/][%d] getMarketsPricesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetMarketsPricesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetMarketsPricesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsPricesServiceUnavailable creates a GetMarketsPricesServiceUnavailable with default headers values
func NewGetMarketsPricesServiceUnavailable() *GetMarketsPricesServiceUnavailable {
	return &GetMarketsPricesServiceUnavailable{}
}

/* GetMarketsPricesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetMarketsPricesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetMarketsPricesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/markets/prices/][%d] getMarketsPricesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetMarketsPricesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetMarketsPricesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsPricesGatewayTimeout creates a GetMarketsPricesGatewayTimeout with default headers values
func NewGetMarketsPricesGatewayTimeout() *GetMarketsPricesGatewayTimeout {
	return &GetMarketsPricesGatewayTimeout{}
}

/* GetMarketsPricesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetMarketsPricesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetMarketsPricesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/markets/prices/][%d] getMarketsPricesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetMarketsPricesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetMarketsPricesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMarketsPricesOKBodyItems0 get_markets_prices_200_ok
//
// 200 ok object
swagger:model GetMarketsPricesOKBodyItems0
*/
type GetMarketsPricesOKBodyItems0 struct {

	// get_markets_prices_adjusted_price
	//
	// adjusted_price number
	AdjustedPrice float64 `json:"adjusted_price,omitempty"`

	// get_markets_prices_average_price
	//
	// average_price number
	AveragePrice float64 `json:"average_price,omitempty"`

	// get_markets_prices_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get markets prices o k body items0
func (o *GetMarketsPricesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMarketsPricesOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get markets prices o k body items0 based on context it is used
func (o *GetMarketsPricesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetMarketsPricesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMarketsPricesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetMarketsPricesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
