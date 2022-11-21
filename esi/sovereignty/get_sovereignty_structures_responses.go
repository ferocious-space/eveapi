// Code generated by go-swagger; DO NOT EDIT.

package sovereignty

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

// GetSovereigntyStructuresReader is a Reader for the GetSovereigntyStructures structure.
type GetSovereigntyStructuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyStructuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSovereigntyStructuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetSovereigntyStructuresNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetSovereigntyStructuresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetSovereigntyStructuresEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSovereigntyStructuresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSovereigntyStructuresServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSovereigntyStructuresGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSovereigntyStructuresOK creates a GetSovereigntyStructuresOK with default headers values
func NewGetSovereigntyStructuresOK() *GetSovereigntyStructuresOK {
	return &GetSovereigntyStructuresOK{}
}

/*
GetSovereigntyStructuresOK describes a response with status code 200, with default header values.

A list of sovereignty structures
*/
type GetSovereigntyStructuresOK struct {

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

	Payload []*GetSovereigntyStructuresOKBodyItems0
}

// IsSuccess returns true when this get sovereignty structures o k response has a 2xx status code
func (o *GetSovereigntyStructuresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sovereignty structures o k response has a 3xx status code
func (o *GetSovereigntyStructuresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty structures o k response has a 4xx status code
func (o *GetSovereigntyStructuresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty structures o k response has a 5xx status code
func (o *GetSovereigntyStructuresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty structures o k response a status code equal to that given
func (o *GetSovereigntyStructuresOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSovereigntyStructuresOK) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyStructuresOK) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyStructuresOK) GetPayload() []*GetSovereigntyStructuresOKBodyItems0 {
	return o.Payload
}

func (o *GetSovereigntyStructuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyStructuresNotModified creates a GetSovereigntyStructuresNotModified with default headers values
func NewGetSovereigntyStructuresNotModified() *GetSovereigntyStructuresNotModified {
	return &GetSovereigntyStructuresNotModified{}
}

/*
GetSovereigntyStructuresNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetSovereigntyStructuresNotModified struct {

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

// IsSuccess returns true when this get sovereignty structures not modified response has a 2xx status code
func (o *GetSovereigntyStructuresNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty structures not modified response has a 3xx status code
func (o *GetSovereigntyStructuresNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get sovereignty structures not modified response has a 4xx status code
func (o *GetSovereigntyStructuresNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty structures not modified response has a 5xx status code
func (o *GetSovereigntyStructuresNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty structures not modified response a status code equal to that given
func (o *GetSovereigntyStructuresNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetSovereigntyStructuresNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresNotModified ", 304)
}

func (o *GetSovereigntyStructuresNotModified) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresNotModified ", 304)
}

func (o *GetSovereigntyStructuresNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyStructuresBadRequest creates a GetSovereigntyStructuresBadRequest with default headers values
func NewGetSovereigntyStructuresBadRequest() *GetSovereigntyStructuresBadRequest {
	return &GetSovereigntyStructuresBadRequest{}
}

/*
GetSovereigntyStructuresBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSovereigntyStructuresBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get sovereignty structures bad request response has a 2xx status code
func (o *GetSovereigntyStructuresBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty structures bad request response has a 3xx status code
func (o *GetSovereigntyStructuresBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty structures bad request response has a 4xx status code
func (o *GetSovereigntyStructuresBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sovereignty structures bad request response has a 5xx status code
func (o *GetSovereigntyStructuresBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty structures bad request response a status code equal to that given
func (o *GetSovereigntyStructuresBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSovereigntyStructuresBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyStructuresBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyStructuresBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetSovereigntyStructuresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyStructuresEnhanceYourCalm creates a GetSovereigntyStructuresEnhanceYourCalm with default headers values
func NewGetSovereigntyStructuresEnhanceYourCalm() *GetSovereigntyStructuresEnhanceYourCalm {
	return &GetSovereigntyStructuresEnhanceYourCalm{}
}

/*
GetSovereigntyStructuresEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetSovereigntyStructuresEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get sovereignty structures enhance your calm response has a 2xx status code
func (o *GetSovereigntyStructuresEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty structures enhance your calm response has a 3xx status code
func (o *GetSovereigntyStructuresEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty structures enhance your calm response has a 4xx status code
func (o *GetSovereigntyStructuresEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sovereignty structures enhance your calm response has a 5xx status code
func (o *GetSovereigntyStructuresEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty structures enhance your calm response a status code equal to that given
func (o *GetSovereigntyStructuresEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetSovereigntyStructuresEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyStructuresEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyStructuresEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetSovereigntyStructuresEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyStructuresInternalServerError creates a GetSovereigntyStructuresInternalServerError with default headers values
func NewGetSovereigntyStructuresInternalServerError() *GetSovereigntyStructuresInternalServerError {
	return &GetSovereigntyStructuresInternalServerError{}
}

/*
GetSovereigntyStructuresInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSovereigntyStructuresInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get sovereignty structures internal server error response has a 2xx status code
func (o *GetSovereigntyStructuresInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty structures internal server error response has a 3xx status code
func (o *GetSovereigntyStructuresInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty structures internal server error response has a 4xx status code
func (o *GetSovereigntyStructuresInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty structures internal server error response has a 5xx status code
func (o *GetSovereigntyStructuresInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty structures internal server error response a status code equal to that given
func (o *GetSovereigntyStructuresInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSovereigntyStructuresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyStructuresInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyStructuresInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetSovereigntyStructuresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyStructuresServiceUnavailable creates a GetSovereigntyStructuresServiceUnavailable with default headers values
func NewGetSovereigntyStructuresServiceUnavailable() *GetSovereigntyStructuresServiceUnavailable {
	return &GetSovereigntyStructuresServiceUnavailable{}
}

/*
GetSovereigntyStructuresServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetSovereigntyStructuresServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get sovereignty structures service unavailable response has a 2xx status code
func (o *GetSovereigntyStructuresServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty structures service unavailable response has a 3xx status code
func (o *GetSovereigntyStructuresServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty structures service unavailable response has a 4xx status code
func (o *GetSovereigntyStructuresServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty structures service unavailable response has a 5xx status code
func (o *GetSovereigntyStructuresServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty structures service unavailable response a status code equal to that given
func (o *GetSovereigntyStructuresServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSovereigntyStructuresServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyStructuresServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyStructuresServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetSovereigntyStructuresServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyStructuresGatewayTimeout creates a GetSovereigntyStructuresGatewayTimeout with default headers values
func NewGetSovereigntyStructuresGatewayTimeout() *GetSovereigntyStructuresGatewayTimeout {
	return &GetSovereigntyStructuresGatewayTimeout{}
}

/*
GetSovereigntyStructuresGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetSovereigntyStructuresGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get sovereignty structures gateway timeout response has a 2xx status code
func (o *GetSovereigntyStructuresGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty structures gateway timeout response has a 3xx status code
func (o *GetSovereigntyStructuresGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty structures gateway timeout response has a 4xx status code
func (o *GetSovereigntyStructuresGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty structures gateway timeout response has a 5xx status code
func (o *GetSovereigntyStructuresGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty structures gateway timeout response a status code equal to that given
func (o *GetSovereigntyStructuresGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetSovereigntyStructuresGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyStructuresGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/structures/][%d] getSovereigntyStructuresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyStructuresGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetSovereigntyStructuresGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetSovereigntyStructuresOKBodyItems0 get_sovereignty_structures_200_ok
//
// 200 ok object
swagger:model GetSovereigntyStructuresOKBodyItems0
*/
type GetSovereigntyStructuresOKBodyItems0 struct {

	// get_sovereignty_structures_alliance_id
	//
	// The alliance that owns the structure.
	//
	// Required: true
	AllianceID *int32 `json:"alliance_id"`

	// get_sovereignty_structures_solar_system_id
	//
	// Solar system in which the structure is located.
	//
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_sovereignty_structures_structure_id
	//
	// Unique item ID for this structure.
	// Required: true
	StructureID *int64 `json:"structure_id"`

	// get_sovereignty_structures_structure_type_id
	//
	// A reference to the type of structure this is.
	//
	// Required: true
	StructureTypeID *int32 `json:"structure_type_id"`

	// get_sovereignty_structures_vulnerability_occupancy_level
	//
	// The occupancy level for the next or current vulnerability window. This takes into account all development indexes and capital system bonuses. Also known as Activity Defense Multiplier from in the client. It increases the time that attackers must spend using their entosis links on the structure.
	//
	VulnerabilityOccupancyLevel float32 `json:"vulnerability_occupancy_level,omitempty"`

	// get_sovereignty_structures_vulnerable_end_time
	//
	// The time at which the next or current vulnerability window ends. At the end of a vulnerability window the next window is recalculated and locked in along with the vulnerabilityOccupancyLevel. If the structure is not in 100% entosis control of the defender, it will go in to 'overtime' and stay vulnerable for as long as that situation persists. Only once the defenders have 100% entosis control and has the vulnerableEndTime passed does the vulnerability interval expire and a new one is calculated.
	//
	// Format: date-time
	VulnerableEndTime strfmt.DateTime `json:"vulnerable_end_time,omitempty"`

	// get_sovereignty_structures_vulnerable_start_time
	//
	// The next time at which the structure will become vulnerable. Or the start time of the current window if current time is between this and vulnerableEndTime.
	//
	// Format: date-time
	VulnerableStartTime strfmt.DateTime `json:"vulnerable_start_time,omitempty"`
}

// Validate validates this get sovereignty structures o k body items0
func (o *GetSovereigntyStructuresOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllianceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStructureID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStructureTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVulnerableEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVulnerableStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyStructuresOKBodyItems0) validateAllianceID(formats strfmt.Registry) error {

	if err := validate.Required("alliance_id", "body", o.AllianceID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyStructuresOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyStructuresOKBodyItems0) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", o.StructureID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyStructuresOKBodyItems0) validateStructureTypeID(formats strfmt.Registry) error {

	if err := validate.Required("structure_type_id", "body", o.StructureTypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyStructuresOKBodyItems0) validateVulnerableEndTime(formats strfmt.Registry) error {
	if swag.IsZero(o.VulnerableEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("vulnerable_end_time", "body", "date-time", o.VulnerableEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyStructuresOKBodyItems0) validateVulnerableStartTime(formats strfmt.Registry) error {
	if swag.IsZero(o.VulnerableStartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("vulnerable_start_time", "body", "date-time", o.VulnerableStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get sovereignty structures o k body items0 based on context it is used
func (o *GetSovereigntyStructuresOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSovereigntyStructuresOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSovereigntyStructuresOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyStructuresOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
