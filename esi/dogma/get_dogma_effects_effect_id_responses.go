// Code generated by go-swagger; DO NOT EDIT.

package dogma

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

// GetDogmaEffectsEffectIDReader is a Reader for the GetDogmaEffectsEffectID structure.
type GetDogmaEffectsEffectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDogmaEffectsEffectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDogmaEffectsEffectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetDogmaEffectsEffectIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetDogmaEffectsEffectIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDogmaEffectsEffectIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetDogmaEffectsEffectIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDogmaEffectsEffectIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDogmaEffectsEffectIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDogmaEffectsEffectIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDogmaEffectsEffectIDOK creates a GetDogmaEffectsEffectIDOK with default headers values
func NewGetDogmaEffectsEffectIDOK() *GetDogmaEffectsEffectIDOK {
	return &GetDogmaEffectsEffectIDOK{}
}

/*
GetDogmaEffectsEffectIDOK describes a response with status code 200, with default header values.

Information about a dogma effect
*/
type GetDogmaEffectsEffectIDOK struct {

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

	Payload *GetDogmaEffectsEffectIDOKBody
}

// IsSuccess returns true when this get dogma effects effect Id o k response has a 2xx status code
func (o *GetDogmaEffectsEffectIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dogma effects effect Id o k response has a 3xx status code
func (o *GetDogmaEffectsEffectIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dogma effects effect Id o k response has a 4xx status code
func (o *GetDogmaEffectsEffectIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dogma effects effect Id o k response has a 5xx status code
func (o *GetDogmaEffectsEffectIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dogma effects effect Id o k response a status code equal to that given
func (o *GetDogmaEffectsEffectIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDogmaEffectsEffectIDOK) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdOK  %+v", 200, o.Payload)
}

func (o *GetDogmaEffectsEffectIDOK) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdOK  %+v", 200, o.Payload)
}

func (o *GetDogmaEffectsEffectIDOK) GetPayload() *GetDogmaEffectsEffectIDOKBody {
	return o.Payload
}

func (o *GetDogmaEffectsEffectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetDogmaEffectsEffectIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDNotModified creates a GetDogmaEffectsEffectIDNotModified with default headers values
func NewGetDogmaEffectsEffectIDNotModified() *GetDogmaEffectsEffectIDNotModified {
	return &GetDogmaEffectsEffectIDNotModified{}
}

/*
GetDogmaEffectsEffectIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetDogmaEffectsEffectIDNotModified struct {

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

// IsSuccess returns true when this get dogma effects effect Id not modified response has a 2xx status code
func (o *GetDogmaEffectsEffectIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dogma effects effect Id not modified response has a 3xx status code
func (o *GetDogmaEffectsEffectIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get dogma effects effect Id not modified response has a 4xx status code
func (o *GetDogmaEffectsEffectIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dogma effects effect Id not modified response has a 5xx status code
func (o *GetDogmaEffectsEffectIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get dogma effects effect Id not modified response a status code equal to that given
func (o *GetDogmaEffectsEffectIDNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetDogmaEffectsEffectIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdNotModified ", 304)
}

func (o *GetDogmaEffectsEffectIDNotModified) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdNotModified ", 304)
}

func (o *GetDogmaEffectsEffectIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDogmaEffectsEffectIDBadRequest creates a GetDogmaEffectsEffectIDBadRequest with default headers values
func NewGetDogmaEffectsEffectIDBadRequest() *GetDogmaEffectsEffectIDBadRequest {
	return &GetDogmaEffectsEffectIDBadRequest{}
}

/*
GetDogmaEffectsEffectIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetDogmaEffectsEffectIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get dogma effects effect Id bad request response has a 2xx status code
func (o *GetDogmaEffectsEffectIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dogma effects effect Id bad request response has a 3xx status code
func (o *GetDogmaEffectsEffectIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dogma effects effect Id bad request response has a 4xx status code
func (o *GetDogmaEffectsEffectIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dogma effects effect Id bad request response has a 5xx status code
func (o *GetDogmaEffectsEffectIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get dogma effects effect Id bad request response a status code equal to that given
func (o *GetDogmaEffectsEffectIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDogmaEffectsEffectIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetDogmaEffectsEffectIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetDogmaEffectsEffectIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetDogmaEffectsEffectIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDNotFound creates a GetDogmaEffectsEffectIDNotFound with default headers values
func NewGetDogmaEffectsEffectIDNotFound() *GetDogmaEffectsEffectIDNotFound {
	return &GetDogmaEffectsEffectIDNotFound{}
}

/*
GetDogmaEffectsEffectIDNotFound describes a response with status code 404, with default header values.

Dogma effect not found
*/
type GetDogmaEffectsEffectIDNotFound struct {
	Payload *GetDogmaEffectsEffectIDNotFoundBody
}

// IsSuccess returns true when this get dogma effects effect Id not found response has a 2xx status code
func (o *GetDogmaEffectsEffectIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dogma effects effect Id not found response has a 3xx status code
func (o *GetDogmaEffectsEffectIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dogma effects effect Id not found response has a 4xx status code
func (o *GetDogmaEffectsEffectIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dogma effects effect Id not found response has a 5xx status code
func (o *GetDogmaEffectsEffectIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dogma effects effect Id not found response a status code equal to that given
func (o *GetDogmaEffectsEffectIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDogmaEffectsEffectIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDogmaEffectsEffectIDNotFound) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDogmaEffectsEffectIDNotFound) GetPayload() *GetDogmaEffectsEffectIDNotFoundBody {
	return o.Payload
}

func (o *GetDogmaEffectsEffectIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDogmaEffectsEffectIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDEnhanceYourCalm creates a GetDogmaEffectsEffectIDEnhanceYourCalm with default headers values
func NewGetDogmaEffectsEffectIDEnhanceYourCalm() *GetDogmaEffectsEffectIDEnhanceYourCalm {
	return &GetDogmaEffectsEffectIDEnhanceYourCalm{}
}

/*
GetDogmaEffectsEffectIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetDogmaEffectsEffectIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get dogma effects effect Id enhance your calm response has a 2xx status code
func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dogma effects effect Id enhance your calm response has a 3xx status code
func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dogma effects effect Id enhance your calm response has a 4xx status code
func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dogma effects effect Id enhance your calm response has a 5xx status code
func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get dogma effects effect Id enhance your calm response a status code equal to that given
func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDInternalServerError creates a GetDogmaEffectsEffectIDInternalServerError with default headers values
func NewGetDogmaEffectsEffectIDInternalServerError() *GetDogmaEffectsEffectIDInternalServerError {
	return &GetDogmaEffectsEffectIDInternalServerError{}
}

/*
GetDogmaEffectsEffectIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetDogmaEffectsEffectIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get dogma effects effect Id internal server error response has a 2xx status code
func (o *GetDogmaEffectsEffectIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dogma effects effect Id internal server error response has a 3xx status code
func (o *GetDogmaEffectsEffectIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dogma effects effect Id internal server error response has a 4xx status code
func (o *GetDogmaEffectsEffectIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dogma effects effect Id internal server error response has a 5xx status code
func (o *GetDogmaEffectsEffectIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dogma effects effect Id internal server error response a status code equal to that given
func (o *GetDogmaEffectsEffectIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDogmaEffectsEffectIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDogmaEffectsEffectIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDogmaEffectsEffectIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetDogmaEffectsEffectIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDServiceUnavailable creates a GetDogmaEffectsEffectIDServiceUnavailable with default headers values
func NewGetDogmaEffectsEffectIDServiceUnavailable() *GetDogmaEffectsEffectIDServiceUnavailable {
	return &GetDogmaEffectsEffectIDServiceUnavailable{}
}

/*
GetDogmaEffectsEffectIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetDogmaEffectsEffectIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get dogma effects effect Id service unavailable response has a 2xx status code
func (o *GetDogmaEffectsEffectIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dogma effects effect Id service unavailable response has a 3xx status code
func (o *GetDogmaEffectsEffectIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dogma effects effect Id service unavailable response has a 4xx status code
func (o *GetDogmaEffectsEffectIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dogma effects effect Id service unavailable response has a 5xx status code
func (o *GetDogmaEffectsEffectIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get dogma effects effect Id service unavailable response a status code equal to that given
func (o *GetDogmaEffectsEffectIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetDogmaEffectsEffectIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDogmaEffectsEffectIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDogmaEffectsEffectIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetDogmaEffectsEffectIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDGatewayTimeout creates a GetDogmaEffectsEffectIDGatewayTimeout with default headers values
func NewGetDogmaEffectsEffectIDGatewayTimeout() *GetDogmaEffectsEffectIDGatewayTimeout {
	return &GetDogmaEffectsEffectIDGatewayTimeout{}
}

/*
GetDogmaEffectsEffectIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetDogmaEffectsEffectIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get dogma effects effect Id gateway timeout response has a 2xx status code
func (o *GetDogmaEffectsEffectIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dogma effects effect Id gateway timeout response has a 3xx status code
func (o *GetDogmaEffectsEffectIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dogma effects effect Id gateway timeout response has a 4xx status code
func (o *GetDogmaEffectsEffectIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dogma effects effect Id gateway timeout response has a 5xx status code
func (o *GetDogmaEffectsEffectIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get dogma effects effect Id gateway timeout response a status code equal to that given
func (o *GetDogmaEffectsEffectIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetDogmaEffectsEffectIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDogmaEffectsEffectIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDogmaEffectsEffectIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetDogmaEffectsEffectIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetDogmaEffectsEffectIDNotFoundBody get_dogma_effects_effect_id_not_found
//
// Not found
swagger:model GetDogmaEffectsEffectIDNotFoundBody
*/
type GetDogmaEffectsEffectIDNotFoundBody struct {

	// get_dogma_effects_effect_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get dogma effects effect ID not found body
func (o *GetDogmaEffectsEffectIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get dogma effects effect ID not found body based on context it is used
func (o *GetDogmaEffectsEffectIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDogmaEffectsEffectIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDogmaEffectsEffectIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetDogmaEffectsEffectIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDogmaEffectsEffectIDOKBody get_dogma_effects_effect_id_ok
//
// 200 ok object
swagger:model GetDogmaEffectsEffectIDOKBody
*/
type GetDogmaEffectsEffectIDOKBody struct {

	// get_dogma_effects_effect_id_description
	//
	// description string
	Description string `json:"description,omitempty"`

	// get_dogma_effects_effect_id_disallow_auto_repeat
	//
	// disallow_auto_repeat boolean
	DisallowAutoRepeat bool `json:"disallow_auto_repeat,omitempty"`

	// get_dogma_effects_effect_id_discharge_attribute_id
	//
	// discharge_attribute_id integer
	DischargeAttributeID int32 `json:"discharge_attribute_id,omitempty"`

	// get_dogma_effects_effect_id_display_name
	//
	// display_name string
	DisplayName string `json:"display_name,omitempty"`

	// get_dogma_effects_effect_id_duration_attribute_id
	//
	// duration_attribute_id integer
	DurationAttributeID int32 `json:"duration_attribute_id,omitempty"`

	// get_dogma_effects_effect_id_effect_category
	//
	// effect_category integer
	EffectCategory int32 `json:"effect_category,omitempty"`

	// get_dogma_effects_effect_id_effect_id
	//
	// effect_id integer
	// Required: true
	EffectID *int32 `json:"effect_id"`

	// get_dogma_effects_effect_id_electronic_chance
	//
	// electronic_chance boolean
	ElectronicChance bool `json:"electronic_chance,omitempty"`

	// get_dogma_effects_effect_id_falloff_attribute_id
	//
	// falloff_attribute_id integer
	FalloffAttributeID int32 `json:"falloff_attribute_id,omitempty"`

	// get_dogma_effects_effect_id_icon_id
	//
	// icon_id integer
	IconID int32 `json:"icon_id,omitempty"`

	// get_dogma_effects_effect_id_is_assistance
	//
	// is_assistance boolean
	IsAssistance bool `json:"is_assistance,omitempty"`

	// get_dogma_effects_effect_id_is_offensive
	//
	// is_offensive boolean
	IsOffensive bool `json:"is_offensive,omitempty"`

	// get_dogma_effects_effect_id_is_warp_safe
	//
	// is_warp_safe boolean
	IsWarpSafe bool `json:"is_warp_safe,omitempty"`

	// get_dogma_effects_effect_id_modifiers
	//
	// modifiers array
	// Max Items: 100
	Modifiers []*GetDogmaEffectsEffectIDOKBodyModifiersItems0 `json:"modifiers"`

	// get_dogma_effects_effect_id_name
	//
	// name string
	Name string `json:"name,omitempty"`

	// get_dogma_effects_effect_id_post_expression
	//
	// post_expression integer
	PostExpression int32 `json:"post_expression,omitempty"`

	// get_dogma_effects_effect_id_pre_expression
	//
	// pre_expression integer
	PreExpression int32 `json:"pre_expression,omitempty"`

	// get_dogma_effects_effect_id_published
	//
	// published boolean
	Published bool `json:"published,omitempty"`

	// get_dogma_effects_effect_id_range_attribute_id
	//
	// range_attribute_id integer
	RangeAttributeID int32 `json:"range_attribute_id,omitempty"`

	// get_dogma_effects_effect_id_range_chance
	//
	// range_chance boolean
	RangeChance bool `json:"range_chance,omitempty"`

	// get_dogma_effects_effect_id_tracking_speed_attribute_id
	//
	// tracking_speed_attribute_id integer
	TrackingSpeedAttributeID int32 `json:"tracking_speed_attribute_id,omitempty"`
}

// Validate validates this get dogma effects effect ID o k body
func (o *GetDogmaEffectsEffectIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEffectID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateModifiers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDogmaEffectsEffectIDOKBody) validateEffectID(formats strfmt.Registry) error {

	if err := validate.Required("getDogmaEffectsEffectIdOK"+"."+"effect_id", "body", o.EffectID); err != nil {
		return err
	}

	return nil
}

func (o *GetDogmaEffectsEffectIDOKBody) validateModifiers(formats strfmt.Registry) error {
	if swag.IsZero(o.Modifiers) { // not required
		return nil
	}

	iModifiersSize := int64(len(o.Modifiers))

	if err := validate.MaxItems("getDogmaEffectsEffectIdOK"+"."+"modifiers", "body", iModifiersSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(o.Modifiers); i++ {
		if swag.IsZero(o.Modifiers[i]) { // not required
			continue
		}

		if o.Modifiers[i] != nil {
			if err := o.Modifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDogmaEffectsEffectIdOK" + "." + "modifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDogmaEffectsEffectIdOK" + "." + "modifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get dogma effects effect ID o k body based on the context it is used
func (o *GetDogmaEffectsEffectIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateModifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDogmaEffectsEffectIDOKBody) contextValidateModifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Modifiers); i++ {

		if o.Modifiers[i] != nil {
			if err := o.Modifiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDogmaEffectsEffectIdOK" + "." + "modifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDogmaEffectsEffectIdOK" + "." + "modifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDogmaEffectsEffectIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDogmaEffectsEffectIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetDogmaEffectsEffectIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDogmaEffectsEffectIDOKBodyModifiersItems0 get_dogma_effects_effect_id_modifier
//
// modifier object
swagger:model GetDogmaEffectsEffectIDOKBodyModifiersItems0
*/
type GetDogmaEffectsEffectIDOKBodyModifiersItems0 struct {

	// get_dogma_effects_effect_id_domain
	//
	// domain string
	Domain string `json:"domain,omitempty"`

	// get_dogma_effects_effect_id_modifier_effect_id
	//
	// effect_id integer
	EffectID int32 `json:"effect_id,omitempty"`

	// get_dogma_effects_effect_id_func
	//
	// func string
	// Required: true
	Func *string `json:"func"`

	// get_dogma_effects_effect_id_modified_attribute_id
	//
	// modified_attribute_id integer
	ModifiedAttributeID int32 `json:"modified_attribute_id,omitempty"`

	// get_dogma_effects_effect_id_modifying_attribute_id
	//
	// modifying_attribute_id integer
	ModifyingAttributeID int32 `json:"modifying_attribute_id,omitempty"`

	// get_dogma_effects_effect_id_operator
	//
	// operator integer
	Operator int32 `json:"operator,omitempty"`
}

// Validate validates this get dogma effects effect ID o k body modifiers items0
func (o *GetDogmaEffectsEffectIDOKBodyModifiersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFunc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDogmaEffectsEffectIDOKBodyModifiersItems0) validateFunc(formats strfmt.Registry) error {

	if err := validate.Required("func", "body", o.Func); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get dogma effects effect ID o k body modifiers items0 based on context it is used
func (o *GetDogmaEffectsEffectIDOKBodyModifiersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDogmaEffectsEffectIDOKBodyModifiersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDogmaEffectsEffectIDOKBodyModifiersItems0) UnmarshalBinary(b []byte) error {
	var res GetDogmaEffectsEffectIDOKBodyModifiersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
