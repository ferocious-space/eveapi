// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// GetUniverseSystemsSystemIDReader is a Reader for the GetUniverseSystemsSystemID structure.
type GetUniverseSystemsSystemIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseSystemsSystemIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseSystemsSystemIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseSystemsSystemIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseSystemsSystemIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseSystemsSystemIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseSystemsSystemIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseSystemsSystemIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseSystemsSystemIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseSystemsSystemIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v4/universe/systems/{system_id}/] get_universe_systems_system_id", response, response.Code())
	}
}

// NewGetUniverseSystemsSystemIDOK creates a GetUniverseSystemsSystemIDOK with default headers values
func NewGetUniverseSystemsSystemIDOK() *GetUniverseSystemsSystemIDOK {
	return &GetUniverseSystemsSystemIDOK{}
}

/*
GetUniverseSystemsSystemIDOK describes a response with status code 200, with default header values.

Information about a solar system
*/
type GetUniverseSystemsSystemIDOK struct {

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

	Payload *GetUniverseSystemsSystemIDOKBody
}

// IsSuccess returns true when this get universe systems system Id o k response has a 2xx status code
func (o *GetUniverseSystemsSystemIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe systems system Id o k response has a 3xx status code
func (o *GetUniverseSystemsSystemIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe systems system Id o k response has a 4xx status code
func (o *GetUniverseSystemsSystemIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe systems system Id o k response has a 5xx status code
func (o *GetUniverseSystemsSystemIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe systems system Id o k response a status code equal to that given
func (o *GetUniverseSystemsSystemIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe systems system Id o k response
func (o *GetUniverseSystemsSystemIDOK) Code() int {
	return 200
}

func (o *GetUniverseSystemsSystemIDOK) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseSystemsSystemIDOK) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseSystemsSystemIDOK) GetPayload() *GetUniverseSystemsSystemIDOKBody {
	return o.Payload
}

func (o *GetUniverseSystemsSystemIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseSystemsSystemIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemsSystemIDNotModified creates a GetUniverseSystemsSystemIDNotModified with default headers values
func NewGetUniverseSystemsSystemIDNotModified() *GetUniverseSystemsSystemIDNotModified {
	return &GetUniverseSystemsSystemIDNotModified{}
}

/*
GetUniverseSystemsSystemIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseSystemsSystemIDNotModified struct {

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

// IsSuccess returns true when this get universe systems system Id not modified response has a 2xx status code
func (o *GetUniverseSystemsSystemIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe systems system Id not modified response has a 3xx status code
func (o *GetUniverseSystemsSystemIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe systems system Id not modified response has a 4xx status code
func (o *GetUniverseSystemsSystemIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe systems system Id not modified response has a 5xx status code
func (o *GetUniverseSystemsSystemIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe systems system Id not modified response a status code equal to that given
func (o *GetUniverseSystemsSystemIDNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe systems system Id not modified response
func (o *GetUniverseSystemsSystemIDNotModified) Code() int {
	return 304
}

func (o *GetUniverseSystemsSystemIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdNotModified ", 304)
}

func (o *GetUniverseSystemsSystemIDNotModified) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdNotModified ", 304)
}

func (o *GetUniverseSystemsSystemIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseSystemsSystemIDBadRequest creates a GetUniverseSystemsSystemIDBadRequest with default headers values
func NewGetUniverseSystemsSystemIDBadRequest() *GetUniverseSystemsSystemIDBadRequest {
	return &GetUniverseSystemsSystemIDBadRequest{}
}

/*
GetUniverseSystemsSystemIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseSystemsSystemIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe systems system Id bad request response has a 2xx status code
func (o *GetUniverseSystemsSystemIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe systems system Id bad request response has a 3xx status code
func (o *GetUniverseSystemsSystemIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe systems system Id bad request response has a 4xx status code
func (o *GetUniverseSystemsSystemIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe systems system Id bad request response has a 5xx status code
func (o *GetUniverseSystemsSystemIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe systems system Id bad request response a status code equal to that given
func (o *GetUniverseSystemsSystemIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe systems system Id bad request response
func (o *GetUniverseSystemsSystemIDBadRequest) Code() int {
	return 400
}

func (o *GetUniverseSystemsSystemIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseSystemsSystemIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseSystemsSystemIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseSystemsSystemIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemsSystemIDNotFound creates a GetUniverseSystemsSystemIDNotFound with default headers values
func NewGetUniverseSystemsSystemIDNotFound() *GetUniverseSystemsSystemIDNotFound {
	return &GetUniverseSystemsSystemIDNotFound{}
}

/*
GetUniverseSystemsSystemIDNotFound describes a response with status code 404, with default header values.

Solar system not found
*/
type GetUniverseSystemsSystemIDNotFound struct {
	Payload *GetUniverseSystemsSystemIDNotFoundBody
}

// IsSuccess returns true when this get universe systems system Id not found response has a 2xx status code
func (o *GetUniverseSystemsSystemIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe systems system Id not found response has a 3xx status code
func (o *GetUniverseSystemsSystemIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe systems system Id not found response has a 4xx status code
func (o *GetUniverseSystemsSystemIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe systems system Id not found response has a 5xx status code
func (o *GetUniverseSystemsSystemIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe systems system Id not found response a status code equal to that given
func (o *GetUniverseSystemsSystemIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get universe systems system Id not found response
func (o *GetUniverseSystemsSystemIDNotFound) Code() int {
	return 404
}

func (o *GetUniverseSystemsSystemIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseSystemsSystemIDNotFound) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseSystemsSystemIDNotFound) GetPayload() *GetUniverseSystemsSystemIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseSystemsSystemIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseSystemsSystemIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemsSystemIDEnhanceYourCalm creates a GetUniverseSystemsSystemIDEnhanceYourCalm with default headers values
func NewGetUniverseSystemsSystemIDEnhanceYourCalm() *GetUniverseSystemsSystemIDEnhanceYourCalm {
	return &GetUniverseSystemsSystemIDEnhanceYourCalm{}
}

/*
GetUniverseSystemsSystemIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseSystemsSystemIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe systems system Id enhance your calm response has a 2xx status code
func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe systems system Id enhance your calm response has a 3xx status code
func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe systems system Id enhance your calm response has a 4xx status code
func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe systems system Id enhance your calm response has a 5xx status code
func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe systems system Id enhance your calm response a status code equal to that given
func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe systems system Id enhance your calm response
func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseSystemsSystemIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemsSystemIDInternalServerError creates a GetUniverseSystemsSystemIDInternalServerError with default headers values
func NewGetUniverseSystemsSystemIDInternalServerError() *GetUniverseSystemsSystemIDInternalServerError {
	return &GetUniverseSystemsSystemIDInternalServerError{}
}

/*
GetUniverseSystemsSystemIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseSystemsSystemIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe systems system Id internal server error response has a 2xx status code
func (o *GetUniverseSystemsSystemIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe systems system Id internal server error response has a 3xx status code
func (o *GetUniverseSystemsSystemIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe systems system Id internal server error response has a 4xx status code
func (o *GetUniverseSystemsSystemIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe systems system Id internal server error response has a 5xx status code
func (o *GetUniverseSystemsSystemIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe systems system Id internal server error response a status code equal to that given
func (o *GetUniverseSystemsSystemIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe systems system Id internal server error response
func (o *GetUniverseSystemsSystemIDInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseSystemsSystemIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseSystemsSystemIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseSystemsSystemIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseSystemsSystemIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemsSystemIDServiceUnavailable creates a GetUniverseSystemsSystemIDServiceUnavailable with default headers values
func NewGetUniverseSystemsSystemIDServiceUnavailable() *GetUniverseSystemsSystemIDServiceUnavailable {
	return &GetUniverseSystemsSystemIDServiceUnavailable{}
}

/*
GetUniverseSystemsSystemIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseSystemsSystemIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe systems system Id service unavailable response has a 2xx status code
func (o *GetUniverseSystemsSystemIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe systems system Id service unavailable response has a 3xx status code
func (o *GetUniverseSystemsSystemIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe systems system Id service unavailable response has a 4xx status code
func (o *GetUniverseSystemsSystemIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe systems system Id service unavailable response has a 5xx status code
func (o *GetUniverseSystemsSystemIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe systems system Id service unavailable response a status code equal to that given
func (o *GetUniverseSystemsSystemIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe systems system Id service unavailable response
func (o *GetUniverseSystemsSystemIDServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseSystemsSystemIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseSystemsSystemIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseSystemsSystemIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseSystemsSystemIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemsSystemIDGatewayTimeout creates a GetUniverseSystemsSystemIDGatewayTimeout with default headers values
func NewGetUniverseSystemsSystemIDGatewayTimeout() *GetUniverseSystemsSystemIDGatewayTimeout {
	return &GetUniverseSystemsSystemIDGatewayTimeout{}
}

/*
GetUniverseSystemsSystemIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseSystemsSystemIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe systems system Id gateway timeout response has a 2xx status code
func (o *GetUniverseSystemsSystemIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe systems system Id gateway timeout response has a 3xx status code
func (o *GetUniverseSystemsSystemIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe systems system Id gateway timeout response has a 4xx status code
func (o *GetUniverseSystemsSystemIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe systems system Id gateway timeout response has a 5xx status code
func (o *GetUniverseSystemsSystemIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe systems system Id gateway timeout response a status code equal to that given
func (o *GetUniverseSystemsSystemIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe systems system Id gateway timeout response
func (o *GetUniverseSystemsSystemIDGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseSystemsSystemIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseSystemsSystemIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v4/universe/systems/{system_id}/][%d] getUniverseSystemsSystemIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseSystemsSystemIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseSystemsSystemIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseSystemsSystemIDNotFoundBody get_universe_systems_system_id_not_found
//
// Not found
swagger:model GetUniverseSystemsSystemIDNotFoundBody
*/
type GetUniverseSystemsSystemIDNotFoundBody struct {

	// get_universe_systems_system_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe systems system ID not found body
func (o *GetUniverseSystemsSystemIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe systems system ID not found body based on context it is used
func (o *GetUniverseSystemsSystemIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseSystemsSystemIDOKBody get_universe_systems_system_id_ok
//
// 200 ok object
swagger:model GetUniverseSystemsSystemIDOKBody
*/
type GetUniverseSystemsSystemIDOKBody struct {

	// get_universe_systems_system_id_constellation_id
	//
	// The constellation this solar system is in
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_universe_systems_system_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_systems_system_id_planets
	//
	// planets array
	// Max Items: 1000
	Planets []*GetUniverseSystemsSystemIDOKBodyPlanetsItems0 `json:"planets"`

	// position
	// Required: true
	Position *GetUniverseSystemsSystemIDOKBodyPosition `json:"position"`

	// get_universe_systems_system_id_security_class
	//
	// security_class string
	SecurityClass string `json:"security_class,omitempty"`

	// get_universe_systems_system_id_security_status
	//
	// security_status number
	// Required: true
	SecurityStatus *float32 `json:"security_status"`

	// get_universe_systems_system_id_star_id
	//
	// star_id integer
	StarID int32 `json:"star_id,omitempty"`

	// get_universe_systems_system_id_stargates
	//
	// stargates array
	// Max Items: 25
	Stargates []int32 `json:"stargates"`

	// get_universe_systems_system_id_stations
	//
	// stations array
	// Max Items: 25
	Stations []int32 `json:"stations"`

	// get_universe_systems_system_id_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe systems system ID o k body
func (o *GetUniverseSystemsSystemIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConstellationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlanets(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecurityStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStargates(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"constellation_id", "body", o.ConstellationID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validatePlanets(formats strfmt.Registry) error {
	if swag.IsZero(o.Planets) { // not required
		return nil
	}

	iPlanetsSize := int64(len(o.Planets))

	if err := validate.MaxItems("getUniverseSystemsSystemIdOK"+"."+"planets", "body", iPlanetsSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(o.Planets); i++ {
		if swag.IsZero(o.Planets[i]) { // not required
			continue
		}

		if o.Planets[i] != nil {
			if err := o.Planets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseSystemsSystemIdOK" + "." + "planets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getUniverseSystemsSystemIdOK" + "." + "planets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseSystemsSystemIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseSystemsSystemIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateSecurityStatus(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"security_status", "body", o.SecurityStatus); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateStargates(formats strfmt.Registry) error {
	if swag.IsZero(o.Stargates) { // not required
		return nil
	}

	iStargatesSize := int64(len(o.Stargates))

	if err := validate.MaxItems("getUniverseSystemsSystemIdOK"+"."+"stargates", "body", iStargatesSize, 25); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateStations(formats strfmt.Registry) error {
	if swag.IsZero(o.Stations) { // not required
		return nil
	}

	iStationsSize := int64(len(o.Stations))

	if err := validate.MaxItems("getUniverseSystemsSystemIdOK"+"."+"stations", "body", iStationsSize, 25); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe systems system ID o k body based on the context it is used
func (o *GetUniverseSystemsSystemIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePlanets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) contextValidatePlanets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Planets); i++ {

		if o.Planets[i] != nil {

			if swag.IsZero(o.Planets[i]) { // not required
				return nil
			}

			if err := o.Planets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getUniverseSystemsSystemIdOK" + "." + "planets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getUniverseSystemsSystemIdOK" + "." + "planets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBody) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {

		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseSystemsSystemIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseSystemsSystemIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseSystemsSystemIDOKBodyPlanetsItems0 get_universe_systems_system_id_planet
//
// planet object
swagger:model GetUniverseSystemsSystemIDOKBodyPlanetsItems0
*/
type GetUniverseSystemsSystemIDOKBodyPlanetsItems0 struct {

	// get_universe_systems_system_id_asteroid_belts
	//
	// asteroid_belts array
	// Max Items: 100
	AsteroidBelts []int32 `json:"asteroid_belts"`

	// get_universe_systems_system_id_moons
	//
	// moons array
	// Max Items: 1000
	Moons []int32 `json:"moons"`

	// get_universe_systems_system_id_planet_id
	//
	// planet_id integer
	// Required: true
	PlanetID *int32 `json:"planet_id"`
}

// Validate validates this get universe systems system ID o k body planets items0
func (o *GetUniverseSystemsSystemIDOKBodyPlanetsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAsteroidBelts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMoons(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlanetID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPlanetsItems0) validateAsteroidBelts(formats strfmt.Registry) error {
	if swag.IsZero(o.AsteroidBelts) { // not required
		return nil
	}

	iAsteroidBeltsSize := int64(len(o.AsteroidBelts))

	if err := validate.MaxItems("asteroid_belts", "body", iAsteroidBeltsSize, 100); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPlanetsItems0) validateMoons(formats strfmt.Registry) error {
	if swag.IsZero(o.Moons) { // not required
		return nil
	}

	iMoonsSize := int64(len(o.Moons))

	if err := validate.MaxItems("moons", "body", iMoonsSize, 1000); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPlanetsItems0) validatePlanetID(formats strfmt.Registry) error {

	if err := validate.Required("planet_id", "body", o.PlanetID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe systems system ID o k body planets items0 based on context it is used
func (o *GetUniverseSystemsSystemIDOKBodyPlanetsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBodyPlanetsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBodyPlanetsItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDOKBodyPlanetsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseSystemsSystemIDOKBodyPosition get_universe_systems_system_id_position
//
// position object
swagger:model GetUniverseSystemsSystemIDOKBodyPosition
*/
type GetUniverseSystemsSystemIDOKBodyPosition struct {

	// get_universe_systems_system_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_systems_system_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_systems_system_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe systems system ID o k body position
func (o *GetUniverseSystemsSystemIDOKBodyPosition) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseSystemsSystemIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemsSystemIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseSystemsSystemIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe systems system ID o k body position based on context it is used
func (o *GetUniverseSystemsSystemIDOKBodyPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemsSystemIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
