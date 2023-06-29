// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
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

// GetUniverseStationsStationIDReader is a Reader for the GetUniverseStationsStationID structure.
type GetUniverseStationsStationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStationsStationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseStationsStationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseStationsStationIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseStationsStationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUniverseStationsStationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseStationsStationIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseStationsStationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseStationsStationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseStationsStationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v2/universe/stations/{station_id}/] get_universe_stations_station_id", response, response.Code())
	}
}

// NewGetUniverseStationsStationIDOK creates a GetUniverseStationsStationIDOK with default headers values
func NewGetUniverseStationsStationIDOK() *GetUniverseStationsStationIDOK {
	return &GetUniverseStationsStationIDOK{}
}

/*
GetUniverseStationsStationIDOK describes a response with status code 200, with default header values.

Information about a station
*/
type GetUniverseStationsStationIDOK struct {

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

	Payload *GetUniverseStationsStationIDOKBody
}

// IsSuccess returns true when this get universe stations station Id o k response has a 2xx status code
func (o *GetUniverseStationsStationIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe stations station Id o k response has a 3xx status code
func (o *GetUniverseStationsStationIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stations station Id o k response has a 4xx status code
func (o *GetUniverseStationsStationIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stations station Id o k response has a 5xx status code
func (o *GetUniverseStationsStationIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stations station Id o k response a status code equal to that given
func (o *GetUniverseStationsStationIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe stations station Id o k response
func (o *GetUniverseStationsStationIDOK) Code() int {
	return 200
}

func (o *GetUniverseStationsStationIDOK) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStationsStationIDOK) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStationsStationIDOK) GetPayload() *GetUniverseStationsStationIDOKBody {
	return o.Payload
}

func (o *GetUniverseStationsStationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetUniverseStationsStationIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStationsStationIDNotModified creates a GetUniverseStationsStationIDNotModified with default headers values
func NewGetUniverseStationsStationIDNotModified() *GetUniverseStationsStationIDNotModified {
	return &GetUniverseStationsStationIDNotModified{}
}

/*
GetUniverseStationsStationIDNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseStationsStationIDNotModified struct {

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

// IsSuccess returns true when this get universe stations station Id not modified response has a 2xx status code
func (o *GetUniverseStationsStationIDNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stations station Id not modified response has a 3xx status code
func (o *GetUniverseStationsStationIDNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe stations station Id not modified response has a 4xx status code
func (o *GetUniverseStationsStationIDNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stations station Id not modified response has a 5xx status code
func (o *GetUniverseStationsStationIDNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stations station Id not modified response a status code equal to that given
func (o *GetUniverseStationsStationIDNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe stations station Id not modified response
func (o *GetUniverseStationsStationIDNotModified) Code() int {
	return 304
}

func (o *GetUniverseStationsStationIDNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdNotModified ", 304)
}

func (o *GetUniverseStationsStationIDNotModified) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdNotModified ", 304)
}

func (o *GetUniverseStationsStationIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStationsStationIDBadRequest creates a GetUniverseStationsStationIDBadRequest with default headers values
func NewGetUniverseStationsStationIDBadRequest() *GetUniverseStationsStationIDBadRequest {
	return &GetUniverseStationsStationIDBadRequest{}
}

/*
GetUniverseStationsStationIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseStationsStationIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe stations station Id bad request response has a 2xx status code
func (o *GetUniverseStationsStationIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stations station Id bad request response has a 3xx status code
func (o *GetUniverseStationsStationIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stations station Id bad request response has a 4xx status code
func (o *GetUniverseStationsStationIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe stations station Id bad request response has a 5xx status code
func (o *GetUniverseStationsStationIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stations station Id bad request response a status code equal to that given
func (o *GetUniverseStationsStationIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe stations station Id bad request response
func (o *GetUniverseStationsStationIDBadRequest) Code() int {
	return 400
}

func (o *GetUniverseStationsStationIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStationsStationIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStationsStationIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseStationsStationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStationsStationIDNotFound creates a GetUniverseStationsStationIDNotFound with default headers values
func NewGetUniverseStationsStationIDNotFound() *GetUniverseStationsStationIDNotFound {
	return &GetUniverseStationsStationIDNotFound{}
}

/*
GetUniverseStationsStationIDNotFound describes a response with status code 404, with default header values.

Station not found
*/
type GetUniverseStationsStationIDNotFound struct {
	Payload *GetUniverseStationsStationIDNotFoundBody
}

// IsSuccess returns true when this get universe stations station Id not found response has a 2xx status code
func (o *GetUniverseStationsStationIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stations station Id not found response has a 3xx status code
func (o *GetUniverseStationsStationIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stations station Id not found response has a 4xx status code
func (o *GetUniverseStationsStationIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe stations station Id not found response has a 5xx status code
func (o *GetUniverseStationsStationIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stations station Id not found response a status code equal to that given
func (o *GetUniverseStationsStationIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get universe stations station Id not found response
func (o *GetUniverseStationsStationIDNotFound) Code() int {
	return 404
}

func (o *GetUniverseStationsStationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseStationsStationIDNotFound) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseStationsStationIDNotFound) GetPayload() *GetUniverseStationsStationIDNotFoundBody {
	return o.Payload
}

func (o *GetUniverseStationsStationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUniverseStationsStationIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStationsStationIDEnhanceYourCalm creates a GetUniverseStationsStationIDEnhanceYourCalm with default headers values
func NewGetUniverseStationsStationIDEnhanceYourCalm() *GetUniverseStationsStationIDEnhanceYourCalm {
	return &GetUniverseStationsStationIDEnhanceYourCalm{}
}

/*
GetUniverseStationsStationIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseStationsStationIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe stations station Id enhance your calm response has a 2xx status code
func (o *GetUniverseStationsStationIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stations station Id enhance your calm response has a 3xx status code
func (o *GetUniverseStationsStationIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stations station Id enhance your calm response has a 4xx status code
func (o *GetUniverseStationsStationIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe stations station Id enhance your calm response has a 5xx status code
func (o *GetUniverseStationsStationIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe stations station Id enhance your calm response a status code equal to that given
func (o *GetUniverseStationsStationIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe stations station Id enhance your calm response
func (o *GetUniverseStationsStationIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseStationsStationIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStationsStationIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStationsStationIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseStationsStationIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStationsStationIDInternalServerError creates a GetUniverseStationsStationIDInternalServerError with default headers values
func NewGetUniverseStationsStationIDInternalServerError() *GetUniverseStationsStationIDInternalServerError {
	return &GetUniverseStationsStationIDInternalServerError{}
}

/*
GetUniverseStationsStationIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseStationsStationIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe stations station Id internal server error response has a 2xx status code
func (o *GetUniverseStationsStationIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stations station Id internal server error response has a 3xx status code
func (o *GetUniverseStationsStationIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stations station Id internal server error response has a 4xx status code
func (o *GetUniverseStationsStationIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stations station Id internal server error response has a 5xx status code
func (o *GetUniverseStationsStationIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe stations station Id internal server error response a status code equal to that given
func (o *GetUniverseStationsStationIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe stations station Id internal server error response
func (o *GetUniverseStationsStationIDInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseStationsStationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStationsStationIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStationsStationIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseStationsStationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStationsStationIDServiceUnavailable creates a GetUniverseStationsStationIDServiceUnavailable with default headers values
func NewGetUniverseStationsStationIDServiceUnavailable() *GetUniverseStationsStationIDServiceUnavailable {
	return &GetUniverseStationsStationIDServiceUnavailable{}
}

/*
GetUniverseStationsStationIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseStationsStationIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe stations station Id service unavailable response has a 2xx status code
func (o *GetUniverseStationsStationIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stations station Id service unavailable response has a 3xx status code
func (o *GetUniverseStationsStationIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stations station Id service unavailable response has a 4xx status code
func (o *GetUniverseStationsStationIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stations station Id service unavailable response has a 5xx status code
func (o *GetUniverseStationsStationIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe stations station Id service unavailable response a status code equal to that given
func (o *GetUniverseStationsStationIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe stations station Id service unavailable response
func (o *GetUniverseStationsStationIDServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseStationsStationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStationsStationIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStationsStationIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseStationsStationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStationsStationIDGatewayTimeout creates a GetUniverseStationsStationIDGatewayTimeout with default headers values
func NewGetUniverseStationsStationIDGatewayTimeout() *GetUniverseStationsStationIDGatewayTimeout {
	return &GetUniverseStationsStationIDGatewayTimeout{}
}

/*
GetUniverseStationsStationIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseStationsStationIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe stations station Id gateway timeout response has a 2xx status code
func (o *GetUniverseStationsStationIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe stations station Id gateway timeout response has a 3xx status code
func (o *GetUniverseStationsStationIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe stations station Id gateway timeout response has a 4xx status code
func (o *GetUniverseStationsStationIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe stations station Id gateway timeout response has a 5xx status code
func (o *GetUniverseStationsStationIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe stations station Id gateway timeout response a status code equal to that given
func (o *GetUniverseStationsStationIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe stations station Id gateway timeout response
func (o *GetUniverseStationsStationIDGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseStationsStationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStationsStationIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/universe/stations/{station_id}/][%d] getUniverseStationsStationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStationsStationIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseStationsStationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseStationsStationIDNotFoundBody get_universe_stations_station_id_not_found
//
// Not found
swagger:model GetUniverseStationsStationIDNotFoundBody
*/
type GetUniverseStationsStationIDNotFoundBody struct {

	// get_universe_stations_station_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe stations station ID not found body
func (o *GetUniverseStationsStationIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get universe stations station ID not found body based on context it is used
func (o *GetUniverseStationsStationIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStationsStationIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStationsStationIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStationsStationIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseStationsStationIDOKBody get_universe_stations_station_id_ok
//
// 200 ok object
swagger:model GetUniverseStationsStationIDOKBody
*/
type GetUniverseStationsStationIDOKBody struct {

	// get_universe_stations_station_id_max_dockable_ship_volume
	//
	// max_dockable_ship_volume number
	// Required: true
	MaxDockableShipVolume *float32 `json:"max_dockable_ship_volume"`

	// get_universe_stations_station_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_stations_station_id_office_rental_cost
	//
	// office_rental_cost number
	// Required: true
	OfficeRentalCost *float32 `json:"office_rental_cost"`

	// get_universe_stations_station_id_owner
	//
	// ID of the corporation that controls this station
	Owner int32 `json:"owner,omitempty"`

	// position
	// Required: true
	Position *GetUniverseStationsStationIDOKBodyPosition `json:"position"`

	// get_universe_stations_station_id_race_id
	//
	// race_id integer
	RaceID int32 `json:"race_id,omitempty"`

	// get_universe_stations_station_id_reprocessing_efficiency
	//
	// reprocessing_efficiency number
	// Required: true
	ReprocessingEfficiency *float32 `json:"reprocessing_efficiency"`

	// get_universe_stations_station_id_reprocessing_stations_take
	//
	// reprocessing_stations_take number
	// Required: true
	ReprocessingStationsTake *float32 `json:"reprocessing_stations_take"`

	// get_universe_stations_station_id_services
	//
	// services array
	// Required: true
	// Max Items: 30
	Services []string `json:"services"`

	// get_universe_stations_station_id_station_id
	//
	// station_id integer
	// Required: true
	StationID *int32 `json:"station_id"`

	// get_universe_stations_station_id_system_id
	//
	// The solar system this station is in
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_universe_stations_station_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get universe stations station ID o k body
func (o *GetUniverseStationsStationIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMaxDockableShipVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOfficeRentalCost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReprocessingEfficiency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReprocessingStationsTake(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStationID(formats); err != nil {
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

func (o *GetUniverseStationsStationIDOKBody) validateMaxDockableShipVolume(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"max_dockable_ship_volume", "body", o.MaxDockableShipVolume); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateOfficeRentalCost(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"office_rental_cost", "body", o.OfficeRentalCost); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStationsStationIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStationsStationIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateReprocessingEfficiency(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"reprocessing_efficiency", "body", o.ReprocessingEfficiency); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateReprocessingStationsTake(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"reprocessing_stations_take", "body", o.ReprocessingStationsTake); err != nil {
		return err
	}

	return nil
}

var getUniverseStationsStationIdOKBodyServicesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bounty-missions","assasination-missions","courier-missions","interbus","reprocessing-plant","refinery","market","black-market","stock-exchange","cloning","surgery","dna-therapy","repair-facilities","factory","labratory","gambling","fitting","paintshop","news","storage","insurance","docking","office-rental","jump-clone-facility","loyalty-point-store","navy-offices","security-offices"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getUniverseStationsStationIdOKBodyServicesItemsEnum = append(getUniverseStationsStationIdOKBodyServicesItemsEnum, v)
	}
}

func (o *GetUniverseStationsStationIDOKBody) validateServicesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getUniverseStationsStationIdOKBodyServicesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"services", "body", o.Services); err != nil {
		return err
	}

	iServicesSize := int64(len(o.Services))

	if err := validate.MaxItems("getUniverseStationsStationIdOK"+"."+"services", "body", iServicesSize, 30); err != nil {
		return err
	}

	for i := 0; i < len(o.Services); i++ {

		// value enum
		if err := o.validateServicesItemsEnum("getUniverseStationsStationIdOK"+"."+"services"+"."+strconv.Itoa(i), "body", o.Services[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"station_id", "body", o.StationID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBody) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get universe stations station ID o k body based on the context it is used
func (o *GetUniverseStationsStationIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseStationsStationIDOKBody) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {

		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUniverseStationsStationIdOK" + "." + "position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getUniverseStationsStationIdOK" + "." + "position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStationsStationIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStationsStationIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseStationsStationIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetUniverseStationsStationIDOKBodyPosition get_universe_stations_station_id_position
//
// position object
swagger:model GetUniverseStationsStationIDOKBodyPosition
*/
type GetUniverseStationsStationIDOKBodyPosition struct {

	// get_universe_stations_station_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_stations_station_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_stations_station_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe stations station ID o k body position
func (o *GetUniverseStationsStationIDOKBodyPosition) Validate(formats strfmt.Registry) error {
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

func (o *GetUniverseStationsStationIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseStationsStationIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("getUniverseStationsStationIdOK"+"."+"position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe stations station ID o k body position based on context it is used
func (o *GetUniverseStationsStationIDOKBodyPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseStationsStationIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseStationsStationIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseStationsStationIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
