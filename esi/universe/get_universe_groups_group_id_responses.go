// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// GetUniverseGroupsGroupIDReader is a Reader for the GetUniverseGroupsGroupID structure.
type GetUniverseGroupsGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseGroupsGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseGroupsGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseGroupsGroupIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseGroupsGroupIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseGroupsGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseGroupsGroupIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseGroupsGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseGroupsGroupIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseGroupsGroupIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseGroupsGroupIDOK creates a GetUniverseGroupsGroupIDOK with default headers values
func NewGetUniverseGroupsGroupIDOK() *GetUniverseGroupsGroupIDOK {
	return &GetUniverseGroupsGroupIDOK{}
}

/* GetUniverseGroupsGroupIDOK describes a response with status code 200, with default header values.

Information about an item group
*/
type GetUniverseGroupsGroupIDOK struct {

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

	Payload *GetUniverseGroupsGroupIDOKBody
}

func (o *GetUniverseGroupsGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdOK  %+v", 200, o.Payload)
}
func (o *GetUniverseGroupsGroupIDOK) GetPayload() *GetUniverseGroupsGroupIDOKBody {
	return o.Payload
}

func (o *GetUniverseGroupsGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseGroupsGroupIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDNotModified creates a GetUniverseGroupsGroupIDNotModified with default headers values
func NewGetUniverseGroupsGroupIDNotModified() *GetUniverseGroupsGroupIDNotModified {
	return &GetUniverseGroupsGroupIDNotModified{}
}

/* GetUniverseGroupsGroupIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseGroupsGroupIDNotModified struct {

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

func (o *GetUniverseGroupsGroupIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdNotModified ", 304)
}

func (o *GetUniverseGroupsGroupIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseGroupsGroupIDBadRequest creates a GetUniverseGroupsGroupIDBadRequest with default headers values
func NewGetUniverseGroupsGroupIDBadRequest() *GetUniverseGroupsGroupIDBadRequest {
	return &GetUniverseGroupsGroupIDBadRequest{}
}

/* GetUniverseGroupsGroupIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseGroupsGroupIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseGroupsGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetUniverseGroupsGroupIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseGroupsGroupIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDNotFound creates a GetUniverseGroupsGroupIDNotFound with default headers values
func NewGetUniverseGroupsGroupIDNotFound() *GetUniverseGroupsGroupIDNotFound {
	return &GetUniverseGroupsGroupIDNotFound{}
}

/* GetUniverseGroupsGroupIDNotFound describes a response with status code 404, with default header values.

Group not found
*/
type GetUniverseGroupsGroupIDNotFound struct {
	Payload *GetUniverseGroupsGroupIDNotFoundBody
}

func (o *GetUniverseGroupsGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdNotFound  %+v", 404, o.Payload)
}
func (o *GetUniverseGroupsGroupIDNotFound) GetPayload() *GetUniverseGroupsGroupIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseGroupsGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseGroupsGroupIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDEnhanceYourCalm creates a GetUniverseGroupsGroupIDEnhanceYourCalm with default headers values
func NewGetUniverseGroupsGroupIDEnhanceYourCalm() *GetUniverseGroupsGroupIDEnhanceYourCalm {
	return &GetUniverseGroupsGroupIDEnhanceYourCalm{}
}

/* GetUniverseGroupsGroupIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseGroupsGroupIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseGroupsGroupIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetUniverseGroupsGroupIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseGroupsGroupIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDInternalServerError creates a GetUniverseGroupsGroupIDInternalServerError with default headers values
func NewGetUniverseGroupsGroupIDInternalServerError() *GetUniverseGroupsGroupIDInternalServerError {
	return &GetUniverseGroupsGroupIDInternalServerError{}
}

/* GetUniverseGroupsGroupIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseGroupsGroupIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseGroupsGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUniverseGroupsGroupIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseGroupsGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDServiceUnavailable creates a GetUniverseGroupsGroupIDServiceUnavailable with default headers values
func NewGetUniverseGroupsGroupIDServiceUnavailable() *GetUniverseGroupsGroupIDServiceUnavailable {
	return &GetUniverseGroupsGroupIDServiceUnavailable{}
}

/* GetUniverseGroupsGroupIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseGroupsGroupIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseGroupsGroupIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetUniverseGroupsGroupIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseGroupsGroupIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseGroupsGroupIDGatewayTimeout creates a GetUniverseGroupsGroupIDGatewayTimeout with default headers values
func NewGetUniverseGroupsGroupIDGatewayTimeout() *GetUniverseGroupsGroupIDGatewayTimeout {
	return &GetUniverseGroupsGroupIDGatewayTimeout{}
}

/* GetUniverseGroupsGroupIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseGroupsGroupIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseGroupsGroupIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/groups/{group_id}/][%d] getUniverseGroupsGroupIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetUniverseGroupsGroupIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseGroupsGroupIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUniverseGroupsGroupIDNotFoundBody get_universe_groups_group_id_not_found
//
// Not found
swagger:model GetUniverseGroupsGroupIDNotFoundBody
*/
type GetUniverseGroupsGroupIDNotFoundBody struct {

	// get_universe_groups_group_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe groups group ID not found body
func (o *GetUniverseGroupsGroupIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe groups group ID not found body based on context it is used
func (o *GetUniverseGroupsGroupIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseGroupsGroupIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseGroupsGroupIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseGroupsGroupIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUniverseGroupsGroupIDOKBody get_universe_groups_group_id_ok
//
// 200 ok object
swagger:model GetUniverseGroupsGroupIDOKBody
*/
type GetUniverseGroupsGroupIDOKBody struct {

	// get_universe_groups_group_id_category_id
	//
	// category_id integer
	// Required: true
	CategoryID *int32 `json:"category_id"`

	// get_universe_groups_group_id_group_id
	//
	// group_id integer
	// Required: true
	GroupID *int32 `json:"group_id"`

	// get_universe_groups_group_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_groups_group_id_published
	//
	// published boolean
	// Required: true
	Published *bool `json:"published"`

	// get_universe_groups_group_id_types
	//
	// types array
	// Required: true
	// Max Items: 10000
	Types []int32 `json:"types"`
}

// Validate validates this get universe groups group ID o k body
func (o *GetUniverseGroupsGroupIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCategoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePublished(formats); err != nil {
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

func (o *GetUniverseGroupsGroupIDOKBody) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGroupsGroupIdOK"+"."+"category_id", "body", o.CategoryID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGroupsGroupIDOKBody) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGroupsGroupIdOK"+"."+"group_id", "body", o.GroupID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGroupsGroupIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGroupsGroupIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGroupsGroupIDOKBody) validatePublished(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGroupsGroupIdOK"+"."+"published", "body", o.Published); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseGroupsGroupIDOKBody) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseGroupsGroupIdOK"+"."+"types", "body", o.Types); err != nil {
		return err
	}

	iTypesSize := int64(len(o.Types))

	if err := validate.MaxItems("getUniverseGroupsGroupIdOK"+"."+"types", "body", iTypesSize, 10000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe groups group ID o k body based on context it is used
func (o *GetUniverseGroupsGroupIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseGroupsGroupIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseGroupsGroupIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseGroupsGroupIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
