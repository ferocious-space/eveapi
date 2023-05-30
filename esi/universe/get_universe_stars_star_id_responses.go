// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// GetUniverseStarsStarIDReader is a Reader for the GetUniverseStarsStarID structure.
type GetUniverseStarsStarIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStarsStarIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseStarsStarIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseStarsStarIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseStarsStarIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseStarsStarIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseStarsStarIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseStarsStarIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseStarsStarIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseStarsStarIDOK creates a GetUniverseStarsStarIDOK with default headers values
func NewGetUniverseStarsStarIDOK() *GetUniverseStarsStarIDOK {
	return &GetUniverseStarsStarIDOK{}
}

/*
GetUniverseStarsStarIDOK describes a response with status code 200, with default header values.

Information about a star
*/
type GetUniverseStarsStarIDOK struct {

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

	Payload *GetUniverseStarsStarIDOKBody
}

// IsSuccess returns true when this get universe stars star Id o k response has a 2xx status code
func (o *GetUniverseStarsStarIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe stars star Id o k response has a 3xx status code
func (o *GetUniverseStarsStarIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stars star Id o k response has a 4xx status code
func (o *GetUniverseStarsStarIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stars star Id o k response has a 5xx status code
func (o *GetUniverseStarsStarIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stars star Id o k response a status code equal to that given
func (o *GetUniverseStarsStarIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe stars star Id o k response
func (o *GetUniverseStarsStarIDOK) Code() int {
	return 200
}

func (o *GetUniverseStarsStarIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStarsStarIDOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStarsStarIDOK) GetPayload() *GetUniverseStarsStarIDOKBody {
	return o.Payload
}

func (o *GetUniverseStarsStarIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseStarsStarIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStarsStarIDNotModified creates a GetUniverseStarsStarIDNotModified with default headers values
func NewGetUniverseStarsStarIDNotModified() *GetUniverseStarsStarIDNotModified {
	return &GetUniverseStarsStarIDNotModified{}
}

/*
GetUniverseStarsStarIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseStarsStarIDNotModified struct {

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

// IsSuccess returns true when this get universe stars star Id not modified response has a 2xx status code
func (o *GetUniverseStarsStarIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stars star Id not modified response has a 3xx status code
func (o *GetUniverseStarsStarIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe stars star Id not modified response has a 4xx status code
func (o *GetUniverseStarsStarIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stars star Id not modified response has a 5xx status code
func (o *GetUniverseStarsStarIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stars star Id not modified response a status code equal to that given
func (o *GetUniverseStarsStarIDNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe stars star Id not modified response
func (o *GetUniverseStarsStarIDNotModified) Code() int {
	return 304
}

func (o *GetUniverseStarsStarIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdNotModified ", 304)
}

func (o *GetUniverseStarsStarIDNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdNotModified ", 304)
}

func (o *GetUniverseStarsStarIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStarsStarIDBadRequest creates a GetUniverseStarsStarIDBadRequest with default headers values
func NewGetUniverseStarsStarIDBadRequest() *GetUniverseStarsStarIDBadRequest {
	return &GetUniverseStarsStarIDBadRequest{}
}

/*
GetUniverseStarsStarIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseStarsStarIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe stars star Id bad request response has a 2xx status code
func (o *GetUniverseStarsStarIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stars star Id bad request response has a 3xx status code
func (o *GetUniverseStarsStarIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stars star Id bad request response has a 4xx status code
func (o *GetUniverseStarsStarIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe stars star Id bad request response has a 5xx status code
func (o *GetUniverseStarsStarIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stars star Id bad request response a status code equal to that given
func (o *GetUniverseStarsStarIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe stars star Id bad request response
func (o *GetUniverseStarsStarIDBadRequest) Code() int {
	return 400
}

func (o *GetUniverseStarsStarIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStarsStarIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStarsStarIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseStarsStarIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStarsStarIDEnhanceYourCalm creates a GetUniverseStarsStarIDEnhanceYourCalm with default headers values
func NewGetUniverseStarsStarIDEnhanceYourCalm() *GetUniverseStarsStarIDEnhanceYourCalm {
	return &GetUniverseStarsStarIDEnhanceYourCalm{}
}

/*
GetUniverseStarsStarIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseStarsStarIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe stars star Id enhance your calm response has a 2xx status code
func (o *GetUniverseStarsStarIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stars star Id enhance your calm response has a 3xx status code
func (o *GetUniverseStarsStarIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stars star Id enhance your calm response has a 4xx status code
func (o *GetUniverseStarsStarIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe stars star Id enhance your calm response has a 5xx status code
func (o *GetUniverseStarsStarIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stars star Id enhance your calm response a status code equal to that given
func (o *GetUniverseStarsStarIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe stars star Id enhance your calm response
func (o *GetUniverseStarsStarIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseStarsStarIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStarsStarIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStarsStarIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseStarsStarIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStarsStarIDInternalServerError creates a GetUniverseStarsStarIDInternalServerError with default headers values
func NewGetUniverseStarsStarIDInternalServerError() *GetUniverseStarsStarIDInternalServerError {
	return &GetUniverseStarsStarIDInternalServerError{}
}

/*
GetUniverseStarsStarIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseStarsStarIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe stars star Id internal server error response has a 2xx status code
func (o *GetUniverseStarsStarIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stars star Id internal server error response has a 3xx status code
func (o *GetUniverseStarsStarIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stars star Id internal server error response has a 4xx status code
func (o *GetUniverseStarsStarIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stars star Id internal server error response has a 5xx status code
func (o *GetUniverseStarsStarIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe stars star Id internal server error response a status code equal to that given
func (o *GetUniverseStarsStarIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe stars star Id internal server error response
func (o *GetUniverseStarsStarIDInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseStarsStarIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStarsStarIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStarsStarIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseStarsStarIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStarsStarIDServiceUnavailable creates a GetUniverseStarsStarIDServiceUnavailable with default headers values
func NewGetUniverseStarsStarIDServiceUnavailable() *GetUniverseStarsStarIDServiceUnavailable {
	return &GetUniverseStarsStarIDServiceUnavailable{}
}

/*
GetUniverseStarsStarIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseStarsStarIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe stars star Id service unavailable response has a 2xx status code
func (o *GetUniverseStarsStarIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stars star Id service unavailable response has a 3xx status code
func (o *GetUniverseStarsStarIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stars star Id service unavailable response has a 4xx status code
func (o *GetUniverseStarsStarIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stars star Id service unavailable response has a 5xx status code
func (o *GetUniverseStarsStarIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe stars star Id service unavailable response a status code equal to that given
func (o *GetUniverseStarsStarIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe stars star Id service unavailable response
func (o *GetUniverseStarsStarIDServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseStarsStarIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStarsStarIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStarsStarIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseStarsStarIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStarsStarIDGatewayTimeout creates a GetUniverseStarsStarIDGatewayTimeout with default headers values
func NewGetUniverseStarsStarIDGatewayTimeout() *GetUniverseStarsStarIDGatewayTimeout {
	return &GetUniverseStarsStarIDGatewayTimeout{}
}

/*
GetUniverseStarsStarIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseStarsStarIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe stars star Id gateway timeout response has a 2xx status code
func (o *GetUniverseStarsStarIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stars star Id gateway timeout response has a 3xx status code
func (o *GetUniverseStarsStarIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stars star Id gateway timeout response has a 4xx status code
func (o *GetUniverseStarsStarIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stars star Id gateway timeout response has a 5xx status code
func (o *GetUniverseStarsStarIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe stars star Id gateway timeout response a status code equal to that given
func (o *GetUniverseStarsStarIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe stars star Id gateway timeout response
func (o *GetUniverseStarsStarIDGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseStarsStarIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStarsStarIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/stars/{star_id}/][%d] getUniverseStarsStarIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStarsStarIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseStarsStarIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseStarsStarIDOKBody get_universe_stars_star_id_ok
//
// 200 ok object
swagger:model GetUniverseStarsStarIDOKBody
*/
type GetUniverseStarsStarIDOKBody struct {

	// get_universe_stars_star_id_age
	//
	// Age of star in years
	// Required: true
	Age *int64 `json:"age"`

	// get_universe_stars_star_id_luminosity
	//
	// luminosity number
	// Required: true
	Luminosity *float32 `json:"luminosity"`

	// get_universe_stars_star_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_stars_star_id_radius
	//
	// radius integer
	// Required: true
	Radius *int64 `json:"radius"`

	// get_universe_stars_star_id_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_universe_stars_star_id_spectral_class
	//
	// spectral_class string
	// Required: true
	// Enum: [K2 V K4 V G2 V G8 V M7 V K7 V M2 V K5 V M3 V G0 V G7 V G3 V F9 V G5 V F6 V K8 V K9 V K6 V G9 V G6 V G4 VI G4 V F8 V F2 V F1 V K3 V F0 VI G1 VI G0 VI K1 V M4 V M1 V M6 V M0 V K2 IV G2 VI K0 V K5 IV F5 VI G6 VI F6 VI F2 IV G3 VI M8 V F1 VI K1 IV F7 V G5 VI M5 V G7 VI F5 V F4 VI F8 VI K3 IV F4 IV F0 V G7 IV G8 VI F2 VI F4 V F7 VI F3 V G1 V G9 VI F3 IV F9 VI M9 V K0 IV F1 IV G4 IV F3 VI K4 IV G5 IV G3 IV G1 IV K7 IV G0 IV K6 IV K9 IV G2 IV F9 IV F0 IV K8 IV G8 IV F6 IV F5 IV A0 A0IV A0IV2]
	SpectralClass *string `json:"spectral_class"`

	// get_universe_stars_star_id_temperature
	//
	// temperature integer
	// Required: true
	Temperature *int32 `json:"temperature"`

	// get_universe_stars_star_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get universe stars star ID o k body
func (o *GetUniverseStarsStarIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAge(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLuminosity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadius(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSpectralClass(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTemperature(formats); err != nil {
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

func (o *GetUniverseStarsStarIDOKBody) validateAge(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"age", "body", o.Age); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStarsStarIDOKBody) validateLuminosity(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"luminosity", "body", o.Luminosity); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStarsStarIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStarsStarIDOKBody) validateRadius(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"radius", "body", o.Radius); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStarsStarIDOKBody) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

var getUniverseStarsStarIdOKBodyTypeSpectralClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["K2 V","K4 V","G2 V","G8 V","M7 V","K7 V","M2 V","K5 V","M3 V","G0 V","G7 V","G3 V","F9 V","G5 V","F6 V","K8 V","K9 V","K6 V","G9 V","G6 V","G4 VI","G4 V","F8 V","F2 V","F1 V","K3 V","F0 VI","G1 VI","G0 VI","K1 V","M4 V","M1 V","M6 V","M0 V","K2 IV","G2 VI","K0 V","K5 IV","F5 VI","G6 VI","F6 VI","F2 IV","G3 VI","M8 V","F1 VI","K1 IV","F7 V","G5 VI","M5 V","G7 VI","F5 V","F4 VI","F8 VI","K3 IV","F4 IV","F0 V","G7 IV","G8 VI","F2 VI","F4 V","F7 VI","F3 V","G1 V","G9 VI","F3 IV","F9 VI","M9 V","K0 IV","F1 IV","G4 IV","F3 VI","K4 IV","G5 IV","G3 IV","G1 IV","K7 IV","G0 IV","K6 IV","K9 IV","G2 IV","F9 IV","F0 IV","K8 IV","G8 IV","F6 IV","F5 IV","A0","A0IV","A0IV2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getUniverseStarsStarIdOKBodyTypeSpectralClassPropEnum = append(getUniverseStarsStarIdOKBodyTypeSpectralClassPropEnum, v)
	}
}

const (

	// GetUniverseStarsStarIDOKBodySpectralClassK2V captures enum value "K2 V"
	GetUniverseStarsStarIDOKBodySpectralClassK2V string = "K2 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK4V captures enum value "K4 V"
	GetUniverseStarsStarIDOKBodySpectralClassK4V string = "K4 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG2V captures enum value "G2 V"
	GetUniverseStarsStarIDOKBodySpectralClassG2V string = "G2 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG8V captures enum value "G8 V"
	GetUniverseStarsStarIDOKBodySpectralClassG8V string = "G8 V"

	// GetUniverseStarsStarIDOKBodySpectralClassM7V captures enum value "M7 V"
	GetUniverseStarsStarIDOKBodySpectralClassM7V string = "M7 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK7V captures enum value "K7 V"
	GetUniverseStarsStarIDOKBodySpectralClassK7V string = "K7 V"

	// GetUniverseStarsStarIDOKBodySpectralClassM2V captures enum value "M2 V"
	GetUniverseStarsStarIDOKBodySpectralClassM2V string = "M2 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK5V captures enum value "K5 V"
	GetUniverseStarsStarIDOKBodySpectralClassK5V string = "K5 V"

	// GetUniverseStarsStarIDOKBodySpectralClassM3V captures enum value "M3 V"
	GetUniverseStarsStarIDOKBodySpectralClassM3V string = "M3 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG0V captures enum value "G0 V"
	GetUniverseStarsStarIDOKBodySpectralClassG0V string = "G0 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG7V captures enum value "G7 V"
	GetUniverseStarsStarIDOKBodySpectralClassG7V string = "G7 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG3V captures enum value "G3 V"
	GetUniverseStarsStarIDOKBodySpectralClassG3V string = "G3 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF9V captures enum value "F9 V"
	GetUniverseStarsStarIDOKBodySpectralClassF9V string = "F9 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG5V captures enum value "G5 V"
	GetUniverseStarsStarIDOKBodySpectralClassG5V string = "G5 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF6V captures enum value "F6 V"
	GetUniverseStarsStarIDOKBodySpectralClassF6V string = "F6 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK8V captures enum value "K8 V"
	GetUniverseStarsStarIDOKBodySpectralClassK8V string = "K8 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK9V captures enum value "K9 V"
	GetUniverseStarsStarIDOKBodySpectralClassK9V string = "K9 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK6V captures enum value "K6 V"
	GetUniverseStarsStarIDOKBodySpectralClassK6V string = "K6 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG9V captures enum value "G9 V"
	GetUniverseStarsStarIDOKBodySpectralClassG9V string = "G9 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG6V captures enum value "G6 V"
	GetUniverseStarsStarIDOKBodySpectralClassG6V string = "G6 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG4VI captures enum value "G4 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG4VI string = "G4 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassG4V captures enum value "G4 V"
	GetUniverseStarsStarIDOKBodySpectralClassG4V string = "G4 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF8V captures enum value "F8 V"
	GetUniverseStarsStarIDOKBodySpectralClassF8V string = "F8 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF2V captures enum value "F2 V"
	GetUniverseStarsStarIDOKBodySpectralClassF2V string = "F2 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF1V captures enum value "F1 V"
	GetUniverseStarsStarIDOKBodySpectralClassF1V string = "F1 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK3V captures enum value "K3 V"
	GetUniverseStarsStarIDOKBodySpectralClassK3V string = "K3 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF0VI captures enum value "F0 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF0VI string = "F0 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassG1VI captures enum value "G1 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG1VI string = "G1 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassG0VI captures enum value "G0 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG0VI string = "G0 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassK1V captures enum value "K1 V"
	GetUniverseStarsStarIDOKBodySpectralClassK1V string = "K1 V"

	// GetUniverseStarsStarIDOKBodySpectralClassM4V captures enum value "M4 V"
	GetUniverseStarsStarIDOKBodySpectralClassM4V string = "M4 V"

	// GetUniverseStarsStarIDOKBodySpectralClassM1V captures enum value "M1 V"
	GetUniverseStarsStarIDOKBodySpectralClassM1V string = "M1 V"

	// GetUniverseStarsStarIDOKBodySpectralClassM6V captures enum value "M6 V"
	GetUniverseStarsStarIDOKBodySpectralClassM6V string = "M6 V"

	// GetUniverseStarsStarIDOKBodySpectralClassM0V captures enum value "M0 V"
	GetUniverseStarsStarIDOKBodySpectralClassM0V string = "M0 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK2IV captures enum value "K2 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK2IV string = "K2 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG2VI captures enum value "G2 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG2VI string = "G2 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassK0V captures enum value "K0 V"
	GetUniverseStarsStarIDOKBodySpectralClassK0V string = "K0 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK5IV captures enum value "K5 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK5IV string = "K5 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF5VI captures enum value "F5 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF5VI string = "F5 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassG6VI captures enum value "G6 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG6VI string = "G6 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF6VI captures enum value "F6 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF6VI string = "F6 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF2IV captures enum value "F2 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF2IV string = "F2 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG3VI captures enum value "G3 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG3VI string = "G3 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassM8V captures enum value "M8 V"
	GetUniverseStarsStarIDOKBodySpectralClassM8V string = "M8 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF1VI captures enum value "F1 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF1VI string = "F1 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassK1IV captures enum value "K1 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK1IV string = "K1 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF7V captures enum value "F7 V"
	GetUniverseStarsStarIDOKBodySpectralClassF7V string = "F7 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG5VI captures enum value "G5 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG5VI string = "G5 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassM5V captures enum value "M5 V"
	GetUniverseStarsStarIDOKBodySpectralClassM5V string = "M5 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG7VI captures enum value "G7 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG7VI string = "G7 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF5V captures enum value "F5 V"
	GetUniverseStarsStarIDOKBodySpectralClassF5V string = "F5 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF4VI captures enum value "F4 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF4VI string = "F4 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF8VI captures enum value "F8 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF8VI string = "F8 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassK3IV captures enum value "K3 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK3IV string = "K3 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF4IV captures enum value "F4 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF4IV string = "F4 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF0V captures enum value "F0 V"
	GetUniverseStarsStarIDOKBodySpectralClassF0V string = "F0 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG7IV captures enum value "G7 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG7IV string = "G7 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG8VI captures enum value "G8 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG8VI string = "G8 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF2VI captures enum value "F2 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF2VI string = "F2 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF4V captures enum value "F4 V"
	GetUniverseStarsStarIDOKBodySpectralClassF4V string = "F4 V"

	// GetUniverseStarsStarIDOKBodySpectralClassF7VI captures enum value "F7 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF7VI string = "F7 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF3V captures enum value "F3 V"
	GetUniverseStarsStarIDOKBodySpectralClassF3V string = "F3 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG1V captures enum value "G1 V"
	GetUniverseStarsStarIDOKBodySpectralClassG1V string = "G1 V"

	// GetUniverseStarsStarIDOKBodySpectralClassG9VI captures enum value "G9 VI"
	GetUniverseStarsStarIDOKBodySpectralClassG9VI string = "G9 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassF3IV captures enum value "F3 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF3IV string = "F3 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF9VI captures enum value "F9 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF9VI string = "F9 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassM9V captures enum value "M9 V"
	GetUniverseStarsStarIDOKBodySpectralClassM9V string = "M9 V"

	// GetUniverseStarsStarIDOKBodySpectralClassK0IV captures enum value "K0 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK0IV string = "K0 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF1IV captures enum value "F1 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF1IV string = "F1 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG4IV captures enum value "G4 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG4IV string = "G4 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF3VI captures enum value "F3 VI"
	GetUniverseStarsStarIDOKBodySpectralClassF3VI string = "F3 VI"

	// GetUniverseStarsStarIDOKBodySpectralClassK4IV captures enum value "K4 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK4IV string = "K4 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG5IV captures enum value "G5 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG5IV string = "G5 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG3IV captures enum value "G3 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG3IV string = "G3 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG1IV captures enum value "G1 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG1IV string = "G1 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassK7IV captures enum value "K7 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK7IV string = "K7 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG0IV captures enum value "G0 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG0IV string = "G0 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassK6IV captures enum value "K6 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK6IV string = "K6 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassK9IV captures enum value "K9 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK9IV string = "K9 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG2IV captures enum value "G2 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG2IV string = "G2 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF9IV captures enum value "F9 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF9IV string = "F9 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF0IV captures enum value "F0 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF0IV string = "F0 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassK8IV captures enum value "K8 IV"
	GetUniverseStarsStarIDOKBodySpectralClassK8IV string = "K8 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassG8IV captures enum value "G8 IV"
	GetUniverseStarsStarIDOKBodySpectralClassG8IV string = "G8 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF6IV captures enum value "F6 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF6IV string = "F6 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassF5IV captures enum value "F5 IV"
	GetUniverseStarsStarIDOKBodySpectralClassF5IV string = "F5 IV"

	// GetUniverseStarsStarIDOKBodySpectralClassA0 captures enum value "A0"
	GetUniverseStarsStarIDOKBodySpectralClassA0 string = "A0"

	// GetUniverseStarsStarIDOKBodySpectralClassA0IV captures enum value "A0IV"
	GetUniverseStarsStarIDOKBodySpectralClassA0IV string = "A0IV"

	// GetUniverseStarsStarIDOKBodySpectralClassA0IV2 captures enum value "A0IV2"
	GetUniverseStarsStarIDOKBodySpectralClassA0IV2 string = "A0IV2"
)

// prop value enum
func (o *GetUniverseStarsStarIDOKBody) validateSpectralClassEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getUniverseStarsStarIdOKBodyTypeSpectralClassPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetUniverseStarsStarIDOKBody) validateSpectralClass(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"spectral_class", "body", o.SpectralClass); err != nil {
		return err
	}

	// value enum
	if err := o.validateSpectralClassEnum("getUniverseStarsStarIdOK"+"."+"spectral_class", "body", *o.SpectralClass); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStarsStarIDOKBody) validateTemperature(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"temperature", "body", o.Temperature); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStarsStarIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStarsStarIdOK"+"."+"type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe stars star ID o k body based on context it is used
func (o *GetUniverseStarsStarIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStarsStarIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStarsStarIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStarsStarIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
