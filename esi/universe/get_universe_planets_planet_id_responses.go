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

// GetUniversePlanetsPlanetIDReader is a Reader for the GetUniversePlanetsPlanetID structure.
type GetUniversePlanetsPlanetIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniversePlanetsPlanetIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniversePlanetsPlanetIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniversePlanetsPlanetIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniversePlanetsPlanetIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniversePlanetsPlanetIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniversePlanetsPlanetIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniversePlanetsPlanetIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniversePlanetsPlanetIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniversePlanetsPlanetIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/universe/planets/{planet_id}/] get_universe_planets_planet_id", response, response.Code())
	}
}

// NewGetUniversePlanetsPlanetIDOK creates a GetUniversePlanetsPlanetIDOK with default headers values
func NewGetUniversePlanetsPlanetIDOK() *GetUniversePlanetsPlanetIDOK {
	return &GetUniversePlanetsPlanetIDOK{}
}

/*
GetUniversePlanetsPlanetIDOK describes a response with status code 200, with default header values.

Information about a planet
*/
type GetUniversePlanetsPlanetIDOK struct {

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

	Payload *GetUniversePlanetsPlanetIDOKBody
}

// IsSuccess returns true when this get universe planets planet Id o k response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe planets planet Id o k response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe planets planet Id o k response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe planets planet Id o k response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe planets planet Id o k response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe planets planet Id o k response
func (o *GetUniversePlanetsPlanetIDOK) Code() int {
	return 200
}

func (o *GetUniversePlanetsPlanetIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdOK  %+v", 200, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdOK  %+v", 200, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDOK) GetPayload() *GetUniversePlanetsPlanetIDOKBody {
	return o.Payload
}

func (o *GetUniversePlanetsPlanetIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniversePlanetsPlanetIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDNotModified creates a GetUniversePlanetsPlanetIDNotModified with default headers values
func NewGetUniversePlanetsPlanetIDNotModified() *GetUniversePlanetsPlanetIDNotModified {
	return &GetUniversePlanetsPlanetIDNotModified{}
}

/*
GetUniversePlanetsPlanetIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniversePlanetsPlanetIDNotModified struct {

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

// IsSuccess returns true when this get universe planets planet Id not modified response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe planets planet Id not modified response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe planets planet Id not modified response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe planets planet Id not modified response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe planets planet Id not modified response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe planets planet Id not modified response
func (o *GetUniversePlanetsPlanetIDNotModified) Code() int {
	return 304
}

func (o *GetUniversePlanetsPlanetIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdNotModified ", 304)
}

func (o *GetUniversePlanetsPlanetIDNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdNotModified ", 304)
}

func (o *GetUniversePlanetsPlanetIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniversePlanetsPlanetIDBadRequest creates a GetUniversePlanetsPlanetIDBadRequest with default headers values
func NewGetUniversePlanetsPlanetIDBadRequest() *GetUniversePlanetsPlanetIDBadRequest {
	return &GetUniversePlanetsPlanetIDBadRequest{}
}

/*
GetUniversePlanetsPlanetIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniversePlanetsPlanetIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe planets planet Id bad request response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe planets planet Id bad request response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe planets planet Id bad request response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe planets planet Id bad request response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe planets planet Id bad request response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe planets planet Id bad request response
func (o *GetUniversePlanetsPlanetIDBadRequest) Code() int {
	return 400
}

func (o *GetUniversePlanetsPlanetIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniversePlanetsPlanetIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDNotFound creates a GetUniversePlanetsPlanetIDNotFound with default headers values
func NewGetUniversePlanetsPlanetIDNotFound() *GetUniversePlanetsPlanetIDNotFound {
	return &GetUniversePlanetsPlanetIDNotFound{}
}

/*
GetUniversePlanetsPlanetIDNotFound describes a response with status code 404, with default header values.

Planet not found
*/
type GetUniversePlanetsPlanetIDNotFound struct {
	Payload *GetUniversePlanetsPlanetIDNotFoundBody
}

// IsSuccess returns true when this get universe planets planet Id not found response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe planets planet Id not found response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe planets planet Id not found response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe planets planet Id not found response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe planets planet Id not found response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get universe planets planet Id not found response
func (o *GetUniversePlanetsPlanetIDNotFound) Code() int {
	return 404
}

func (o *GetUniversePlanetsPlanetIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDNotFound) GetPayload() *GetUniversePlanetsPlanetIDNotFoundBody {
	return o.Payload
}

func (o *GetUniversePlanetsPlanetIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniversePlanetsPlanetIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDEnhanceYourCalm creates a GetUniversePlanetsPlanetIDEnhanceYourCalm with default headers values
func NewGetUniversePlanetsPlanetIDEnhanceYourCalm() *GetUniversePlanetsPlanetIDEnhanceYourCalm {
	return &GetUniversePlanetsPlanetIDEnhanceYourCalm{}
}

/*
GetUniversePlanetsPlanetIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniversePlanetsPlanetIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe planets planet Id enhance your calm response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe planets planet Id enhance your calm response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe planets planet Id enhance your calm response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe planets planet Id enhance your calm response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe planets planet Id enhance your calm response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe planets planet Id enhance your calm response
func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniversePlanetsPlanetIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDInternalServerError creates a GetUniversePlanetsPlanetIDInternalServerError with default headers values
func NewGetUniversePlanetsPlanetIDInternalServerError() *GetUniversePlanetsPlanetIDInternalServerError {
	return &GetUniversePlanetsPlanetIDInternalServerError{}
}

/*
GetUniversePlanetsPlanetIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniversePlanetsPlanetIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe planets planet Id internal server error response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe planets planet Id internal server error response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe planets planet Id internal server error response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe planets planet Id internal server error response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe planets planet Id internal server error response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe planets planet Id internal server error response
func (o *GetUniversePlanetsPlanetIDInternalServerError) Code() int {
	return 500
}

func (o *GetUniversePlanetsPlanetIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniversePlanetsPlanetIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDServiceUnavailable creates a GetUniversePlanetsPlanetIDServiceUnavailable with default headers values
func NewGetUniversePlanetsPlanetIDServiceUnavailable() *GetUniversePlanetsPlanetIDServiceUnavailable {
	return &GetUniversePlanetsPlanetIDServiceUnavailable{}
}

/*
GetUniversePlanetsPlanetIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniversePlanetsPlanetIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe planets planet Id service unavailable response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe planets planet Id service unavailable response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe planets planet Id service unavailable response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe planets planet Id service unavailable response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe planets planet Id service unavailable response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe planets planet Id service unavailable response
func (o *GetUniversePlanetsPlanetIDServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniversePlanetsPlanetIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniversePlanetsPlanetIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniversePlanetsPlanetIDGatewayTimeout creates a GetUniversePlanetsPlanetIDGatewayTimeout with default headers values
func NewGetUniversePlanetsPlanetIDGatewayTimeout() *GetUniversePlanetsPlanetIDGatewayTimeout {
	return &GetUniversePlanetsPlanetIDGatewayTimeout{}
}

/*
GetUniversePlanetsPlanetIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniversePlanetsPlanetIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe planets planet Id gateway timeout response has a 2xx status code
func (o *GetUniversePlanetsPlanetIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe planets planet Id gateway timeout response has a 3xx status code
func (o *GetUniversePlanetsPlanetIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe planets planet Id gateway timeout response has a 4xx status code
func (o *GetUniversePlanetsPlanetIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe planets planet Id gateway timeout response has a 5xx status code
func (o *GetUniversePlanetsPlanetIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe planets planet Id gateway timeout response a status code equal to that given
func (o *GetUniversePlanetsPlanetIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe planets planet Id gateway timeout response
func (o *GetUniversePlanetsPlanetIDGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniversePlanetsPlanetIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/planets/{planet_id}/][%d] getUniversePlanetsPlanetIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniversePlanetsPlanetIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniversePlanetsPlanetIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniversePlanetsPlanetIDNotFoundBody get_universe_planets_planet_id_not_found
//
// Not found
swagger:model GetUniversePlanetsPlanetIDNotFoundBody
*/
type GetUniversePlanetsPlanetIDNotFoundBody struct {

	// get_universe_planets_planet_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe planets planet ID not found body
func (o *GetUniversePlanetsPlanetIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe planets planet ID not found body based on context it is used
func (o *GetUniversePlanetsPlanetIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniversePlanetsPlanetIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniversePlanetsPlanetIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniversePlanetsPlanetIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniversePlanetsPlanetIDOKBody get_universe_planets_planet_id_ok
//
// 200 ok object
swagger:model GetUniversePlanetsPlanetIDOKBody
*/
type GetUniversePlanetsPlanetIDOKBody struct {

	// get_universe_planets_planet_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_planets_planet_id_planet_id
	//
	// planet_id integer
	// Required: true
	PlanetID *int32 `json:"planet_id"`

	// position
	// Required: true
	Position *GetUniversePlanetsPlanetIDOKBodyPosition `json:"position"`

	// get_universe_planets_planet_id_system_id
	//
	// The solar system this planet is in
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_universe_planets_planet_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get universe planets planet ID o k body
func (o *GetUniversePlanetsPlanetIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlanetID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSystemID(formats); err != nil {
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

func (o *GetUniversePlanetsPlanetIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBody) validatePlanetID(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"planet_id", "body", o.PlanetID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniversePlanetsPlanetIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniversePlanetsPlanetIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe planets planet ID o k body based on the context it is used
func (o *GetUniversePlanetsPlanetIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBody) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {

		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniversePlanetsPlanetIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniversePlanetsPlanetIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniversePlanetsPlanetIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniversePlanetsPlanetIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniversePlanetsPlanetIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniversePlanetsPlanetIDOKBodyPosition get_universe_planets_planet_id_position
//
// position object
swagger:model GetUniversePlanetsPlanetIDOKBodyPosition
*/
type GetUniversePlanetsPlanetIDOKBodyPosition struct {

	// get_universe_planets_planet_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_planets_planet_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_planets_planet_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe planets planet ID o k body position
func (o *GetUniversePlanetsPlanetIDOKBodyPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateX(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateY(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateZ(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniversePlanetsPlanetIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniversePlanetsPlanetIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe planets planet ID o k body position based on context it is used
func (o *GetUniversePlanetsPlanetIDOKBodyPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniversePlanetsPlanetIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniversePlanetsPlanetIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniversePlanetsPlanetIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
