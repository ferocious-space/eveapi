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

// GetUniverseRacesReader is a Reader for the GetUniverseRaces structure.
type GetUniverseRacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseRacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseRacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseRacesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseRacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseRacesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseRacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseRacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseRacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/universe/races/] get_universe_races", response, response.Code())
	}
}

// NewGetUniverseRacesOK creates a GetUniverseRacesOK with default headers values
func NewGetUniverseRacesOK() *GetUniverseRacesOK {
	return &GetUniverseRacesOK{}
}

/*
GetUniverseRacesOK describes a response with status code 200, with default header values.

A list of character races
*/
type GetUniverseRacesOK struct {

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

	Payload []*GetUniverseRacesOKBodyItems0
}

// IsSuccess returns true when this get universe races o k response has a 2xx status code
func (o *GetUniverseRacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe races o k response has a 3xx status code
func (o *GetUniverseRacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe races o k response has a 4xx status code
func (o *GetUniverseRacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe races o k response has a 5xx status code
func (o *GetUniverseRacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe races o k response a status code equal to that given
func (o *GetUniverseRacesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe races o k response
func (o *GetUniverseRacesOK) Code() int {
	return 200
}

func (o *GetUniverseRacesOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseRacesOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseRacesOK) GetPayload() []*GetUniverseRacesOKBodyItems0 {
	return o.Payload
}

func (o *GetUniverseRacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRacesNotModified creates a GetUniverseRacesNotModified with default headers values
func NewGetUniverseRacesNotModified() *GetUniverseRacesNotModified {
	return &GetUniverseRacesNotModified{}
}

/*
GetUniverseRacesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseRacesNotModified struct {

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

// IsSuccess returns true when this get universe races not modified response has a 2xx status code
func (o *GetUniverseRacesNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe races not modified response has a 3xx status code
func (o *GetUniverseRacesNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe races not modified response has a 4xx status code
func (o *GetUniverseRacesNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe races not modified response has a 5xx status code
func (o *GetUniverseRacesNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe races not modified response a status code equal to that given
func (o *GetUniverseRacesNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe races not modified response
func (o *GetUniverseRacesNotModified) Code() int {
	return 304
}

func (o *GetUniverseRacesNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesNotModified ", 304)
}

func (o *GetUniverseRacesNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesNotModified ", 304)
}

func (o *GetUniverseRacesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseRacesBadRequest creates a GetUniverseRacesBadRequest with default headers values
func NewGetUniverseRacesBadRequest() *GetUniverseRacesBadRequest {
	return &GetUniverseRacesBadRequest{}
}

/*
GetUniverseRacesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseRacesBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe races bad request response has a 2xx status code
func (o *GetUniverseRacesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe races bad request response has a 3xx status code
func (o *GetUniverseRacesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe races bad request response has a 4xx status code
func (o *GetUniverseRacesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe races bad request response has a 5xx status code
func (o *GetUniverseRacesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe races bad request response a status code equal to that given
func (o *GetUniverseRacesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe races bad request response
func (o *GetUniverseRacesBadRequest) Code() int {
	return 400
}

func (o *GetUniverseRacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseRacesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseRacesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseRacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRacesEnhanceYourCalm creates a GetUniverseRacesEnhanceYourCalm with default headers values
func NewGetUniverseRacesEnhanceYourCalm() *GetUniverseRacesEnhanceYourCalm {
	return &GetUniverseRacesEnhanceYourCalm{}
}

/*
GetUniverseRacesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseRacesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe races enhance your calm response has a 2xx status code
func (o *GetUniverseRacesEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe races enhance your calm response has a 3xx status code
func (o *GetUniverseRacesEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe races enhance your calm response has a 4xx status code
func (o *GetUniverseRacesEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe races enhance your calm response has a 5xx status code
func (o *GetUniverseRacesEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe races enhance your calm response a status code equal to that given
func (o *GetUniverseRacesEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe races enhance your calm response
func (o *GetUniverseRacesEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseRacesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseRacesEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseRacesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseRacesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRacesInternalServerError creates a GetUniverseRacesInternalServerError with default headers values
func NewGetUniverseRacesInternalServerError() *GetUniverseRacesInternalServerError {
	return &GetUniverseRacesInternalServerError{}
}

/*
GetUniverseRacesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseRacesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe races internal server error response has a 2xx status code
func (o *GetUniverseRacesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe races internal server error response has a 3xx status code
func (o *GetUniverseRacesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe races internal server error response has a 4xx status code
func (o *GetUniverseRacesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe races internal server error response has a 5xx status code
func (o *GetUniverseRacesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe races internal server error response a status code equal to that given
func (o *GetUniverseRacesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe races internal server error response
func (o *GetUniverseRacesInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseRacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseRacesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseRacesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseRacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRacesServiceUnavailable creates a GetUniverseRacesServiceUnavailable with default headers values
func NewGetUniverseRacesServiceUnavailable() *GetUniverseRacesServiceUnavailable {
	return &GetUniverseRacesServiceUnavailable{}
}

/*
GetUniverseRacesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseRacesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe races service unavailable response has a 2xx status code
func (o *GetUniverseRacesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe races service unavailable response has a 3xx status code
func (o *GetUniverseRacesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe races service unavailable response has a 4xx status code
func (o *GetUniverseRacesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe races service unavailable response has a 5xx status code
func (o *GetUniverseRacesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe races service unavailable response a status code equal to that given
func (o *GetUniverseRacesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe races service unavailable response
func (o *GetUniverseRacesServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseRacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseRacesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseRacesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseRacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseRacesGatewayTimeout creates a GetUniverseRacesGatewayTimeout with default headers values
func NewGetUniverseRacesGatewayTimeout() *GetUniverseRacesGatewayTimeout {
	return &GetUniverseRacesGatewayTimeout{}
}

/*
GetUniverseRacesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseRacesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe races gateway timeout response has a 2xx status code
func (o *GetUniverseRacesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe races gateway timeout response has a 3xx status code
func (o *GetUniverseRacesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe races gateway timeout response has a 4xx status code
func (o *GetUniverseRacesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe races gateway timeout response has a 5xx status code
func (o *GetUniverseRacesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe races gateway timeout response a status code equal to that given
func (o *GetUniverseRacesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe races gateway timeout response
func (o *GetUniverseRacesGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseRacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseRacesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/races/][%d] getUniverseRacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseRacesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseRacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseRacesOKBodyItems0 get_universe_races_200_ok
//
// 200 ok object
swagger:model GetUniverseRacesOKBodyItems0
*/
type GetUniverseRacesOKBodyItems0 struct {

	// get_universe_races_alliance_id
	//
	// The alliance generally associated with this race
	// Required: true
	AllianceID *int32 `json:"alliance_id"`

	// get_universe_races_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_universe_races_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_races_race_id
	//
	// race_id integer
	// Required: true
	RaceID *int32 `json:"race_id"`
}

// Validate validates this get universe races o k body items0
func (o *GetUniverseRacesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllianceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRaceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateAllianceID(formats strfmt.Registry) error {

	if err := validate.Required("alliance_id", "body", o.AllianceID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseRacesOKBodyItems0) validateRaceID(formats strfmt.Registry) error {

	if err := validate.Required("race_id", "body", o.RaceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe races o k body items0 based on context it is used
func (o *GetUniverseRacesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseRacesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseRacesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseRacesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
