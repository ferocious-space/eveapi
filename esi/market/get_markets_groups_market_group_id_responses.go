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

// GetMarketsGroupsMarketGroupIDReader is a Reader for the GetMarketsGroupsMarketGroupID structure.
type GetMarketsGroupsMarketGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketsGroupsMarketGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMarketsGroupsMarketGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetMarketsGroupsMarketGroupIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetMarketsGroupsMarketGroupIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMarketsGroupsMarketGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetMarketsGroupsMarketGroupIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMarketsGroupsMarketGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMarketsGroupsMarketGroupIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetMarketsGroupsMarketGroupIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMarketsGroupsMarketGroupIDOK creates a GetMarketsGroupsMarketGroupIDOK with default headers values
func NewGetMarketsGroupsMarketGroupIDOK() *GetMarketsGroupsMarketGroupIDOK {
	return &GetMarketsGroupsMarketGroupIDOK{}
}

/* GetMarketsGroupsMarketGroupIDOK describes a response with status code 200, with default header values.

Information about an item group
*/
type GetMarketsGroupsMarketGroupIDOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* The language used in the response
	 */
	ContentLanguage string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *GetMarketsGroupsMarketGroupIDOKBody
}

func (o *GetMarketsGroupsMarketGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdOK  %+v", 200, o.Payload)
}
func (o *GetMarketsGroupsMarketGroupIDOK) GetPayload() *GetMarketsGroupsMarketGroupIDOKBody {
	return o.Payload
}

func (o *GetMarketsGroupsMarketGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Content-Language
	hdrContentLanguage := response.GetHeader("Content-Language")

	if hdrContentLanguage != "" {
		o.ContentLanguage = hdrContentLanguage
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

	o.Payload = new(GetMarketsGroupsMarketGroupIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsGroupsMarketGroupIDNotModified creates a GetMarketsGroupsMarketGroupIDNotModified with default headers values
func NewGetMarketsGroupsMarketGroupIDNotModified() *GetMarketsGroupsMarketGroupIDNotModified {
	return &GetMarketsGroupsMarketGroupIDNotModified{}
}

/* GetMarketsGroupsMarketGroupIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetMarketsGroupsMarketGroupIDNotModified struct {

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

func (o *GetMarketsGroupsMarketGroupIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdNotModified ", 304)
}

func (o *GetMarketsGroupsMarketGroupIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMarketsGroupsMarketGroupIDBadRequest creates a GetMarketsGroupsMarketGroupIDBadRequest with default headers values
func NewGetMarketsGroupsMarketGroupIDBadRequest() *GetMarketsGroupsMarketGroupIDBadRequest {
	return &GetMarketsGroupsMarketGroupIDBadRequest{}
}

/* GetMarketsGroupsMarketGroupIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetMarketsGroupsMarketGroupIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetMarketsGroupsMarketGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetMarketsGroupsMarketGroupIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetMarketsGroupsMarketGroupIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsGroupsMarketGroupIDNotFound creates a GetMarketsGroupsMarketGroupIDNotFound with default headers values
func NewGetMarketsGroupsMarketGroupIDNotFound() *GetMarketsGroupsMarketGroupIDNotFound {
	return &GetMarketsGroupsMarketGroupIDNotFound{}
}

/* GetMarketsGroupsMarketGroupIDNotFound describes a response with status code 404, with default header values.

Market group not found
*/
type GetMarketsGroupsMarketGroupIDNotFound struct {
	Payload *GetMarketsGroupsMarketGroupIDNotFoundBody
}

func (o *GetMarketsGroupsMarketGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdNotFound  %+v", 404, o.Payload)
}
func (o *GetMarketsGroupsMarketGroupIDNotFound) GetPayload() *GetMarketsGroupsMarketGroupIDNotFoundBody {
	return o.Payload
}

func (o *GetMarketsGroupsMarketGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMarketsGroupsMarketGroupIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsGroupsMarketGroupIDEnhanceYourCalm creates a GetMarketsGroupsMarketGroupIDEnhanceYourCalm with default headers values
func NewGetMarketsGroupsMarketGroupIDEnhanceYourCalm() *GetMarketsGroupsMarketGroupIDEnhanceYourCalm {
	return &GetMarketsGroupsMarketGroupIDEnhanceYourCalm{}
}

/* GetMarketsGroupsMarketGroupIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetMarketsGroupsMarketGroupIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetMarketsGroupsMarketGroupIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetMarketsGroupsMarketGroupIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetMarketsGroupsMarketGroupIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsGroupsMarketGroupIDInternalServerError creates a GetMarketsGroupsMarketGroupIDInternalServerError with default headers values
func NewGetMarketsGroupsMarketGroupIDInternalServerError() *GetMarketsGroupsMarketGroupIDInternalServerError {
	return &GetMarketsGroupsMarketGroupIDInternalServerError{}
}

/* GetMarketsGroupsMarketGroupIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetMarketsGroupsMarketGroupIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetMarketsGroupsMarketGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetMarketsGroupsMarketGroupIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetMarketsGroupsMarketGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsGroupsMarketGroupIDServiceUnavailable creates a GetMarketsGroupsMarketGroupIDServiceUnavailable with default headers values
func NewGetMarketsGroupsMarketGroupIDServiceUnavailable() *GetMarketsGroupsMarketGroupIDServiceUnavailable {
	return &GetMarketsGroupsMarketGroupIDServiceUnavailable{}
}

/* GetMarketsGroupsMarketGroupIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetMarketsGroupsMarketGroupIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetMarketsGroupsMarketGroupIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetMarketsGroupsMarketGroupIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetMarketsGroupsMarketGroupIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketsGroupsMarketGroupIDGatewayTimeout creates a GetMarketsGroupsMarketGroupIDGatewayTimeout with default headers values
func NewGetMarketsGroupsMarketGroupIDGatewayTimeout() *GetMarketsGroupsMarketGroupIDGatewayTimeout {
	return &GetMarketsGroupsMarketGroupIDGatewayTimeout{}
}

/* GetMarketsGroupsMarketGroupIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetMarketsGroupsMarketGroupIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetMarketsGroupsMarketGroupIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/markets/groups/{market_group_id}/][%d] getMarketsGroupsMarketGroupIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetMarketsGroupsMarketGroupIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetMarketsGroupsMarketGroupIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMarketsGroupsMarketGroupIDNotFoundBody get_markets_groups_market_group_id_not_found
//
// Not found
swagger:model GetMarketsGroupsMarketGroupIDNotFoundBody
*/
type GetMarketsGroupsMarketGroupIDNotFoundBody struct {

	// get_markets_groups_market_group_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get markets groups market group ID not found body
func (o *GetMarketsGroupsMarketGroupIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get markets groups market group ID not found body based on context it is used
func (o *GetMarketsGroupsMarketGroupIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetMarketsGroupsMarketGroupIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMarketsGroupsMarketGroupIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetMarketsGroupsMarketGroupIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetMarketsGroupsMarketGroupIDOKBody get_markets_groups_market_group_id_ok
//
// 200 ok object
swagger:model GetMarketsGroupsMarketGroupIDOKBody
*/
type GetMarketsGroupsMarketGroupIDOKBody struct {

	// get_markets_groups_market_group_id_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_markets_groups_market_group_id_market_group_id
	//
	// market_group_id integer
	// Required: true
	MarketGroupID *int32 `json:"market_group_id"`

	// get_markets_groups_market_group_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_markets_groups_market_group_id_parent_group_id
	//
	// parent_group_id integer
	ParentGroupID int32 `json:"parent_group_id,omitempty"`

	// get_markets_groups_market_group_id_types
	//
	// types array
	// Required: true
	// Max Items: 5000
	Types []int32 `json:"types"`
}

// Validate validates this get markets groups market group ID o k body
func (o *GetMarketsGroupsMarketGroupIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMarketGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateMarketGroupID(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"market_group_id", "body", o.MarketGroupID); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetMarketsGroupsMarketGroupIDOKBody) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("getMarketsGroupsMarketGroupIdOK"+"."+"types", "body", o.Types); err != nil {
		return err
	}

	iTypesSize := int64(len(o.Types))

	if err := validate.MaxItems("getMarketsGroupsMarketGroupIdOK"+"."+"types", "body", iTypesSize, 5000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get markets groups market group ID o k body based on context it is used
func (o *GetMarketsGroupsMarketGroupIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetMarketsGroupsMarketGroupIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMarketsGroupsMarketGroupIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetMarketsGroupsMarketGroupIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
