// Code generated by go-swagger; DO NOT EDIT.

package industry

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

// GetIndustryFacilitiesReader is a Reader for the GetIndustryFacilities structure.
type GetIndustryFacilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIndustryFacilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIndustryFacilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetIndustryFacilitiesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetIndustryFacilitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetIndustryFacilitiesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIndustryFacilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIndustryFacilitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIndustryFacilitiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIndustryFacilitiesOK creates a GetIndustryFacilitiesOK with default headers values
func NewGetIndustryFacilitiesOK() *GetIndustryFacilitiesOK {
	return &GetIndustryFacilitiesOK{}
}

/*
GetIndustryFacilitiesOK describes a response with status code 200, with default header values.

A list of facilities
*/
type GetIndustryFacilitiesOK struct {

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

	Payload []*GetIndustryFacilitiesOKBodyItems0
}

// IsSuccess returns true when this get industry facilities o k response has a 2xx status code
func (o *GetIndustryFacilitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get industry facilities o k response has a 3xx status code
func (o *GetIndustryFacilitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry facilities o k response has a 4xx status code
func (o *GetIndustryFacilitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry facilities o k response has a 5xx status code
func (o *GetIndustryFacilitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry facilities o k response a status code equal to that given
func (o *GetIndustryFacilitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIndustryFacilitiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesOK  %+v", 200, o.Payload)
}

func (o *GetIndustryFacilitiesOK) String() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesOK  %+v", 200, o.Payload)
}

func (o *GetIndustryFacilitiesOK) GetPayload() []*GetIndustryFacilitiesOKBodyItems0 {
	return o.Payload
}

func (o *GetIndustryFacilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustryFacilitiesNotModified creates a GetIndustryFacilitiesNotModified with default headers values
func NewGetIndustryFacilitiesNotModified() *GetIndustryFacilitiesNotModified {
	return &GetIndustryFacilitiesNotModified{}
}

/*
GetIndustryFacilitiesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetIndustryFacilitiesNotModified struct {

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

// IsSuccess returns true when this get industry facilities not modified response has a 2xx status code
func (o *GetIndustryFacilitiesNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry facilities not modified response has a 3xx status code
func (o *GetIndustryFacilitiesNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get industry facilities not modified response has a 4xx status code
func (o *GetIndustryFacilitiesNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry facilities not modified response has a 5xx status code
func (o *GetIndustryFacilitiesNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry facilities not modified response a status code equal to that given
func (o *GetIndustryFacilitiesNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetIndustryFacilitiesNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesNotModified ", 304)
}

func (o *GetIndustryFacilitiesNotModified) String() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesNotModified ", 304)
}

func (o *GetIndustryFacilitiesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustryFacilitiesBadRequest creates a GetIndustryFacilitiesBadRequest with default headers values
func NewGetIndustryFacilitiesBadRequest() *GetIndustryFacilitiesBadRequest {
	return &GetIndustryFacilitiesBadRequest{}
}

/*
GetIndustryFacilitiesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetIndustryFacilitiesBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get industry facilities bad request response has a 2xx status code
func (o *GetIndustryFacilitiesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry facilities bad request response has a 3xx status code
func (o *GetIndustryFacilitiesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry facilities bad request response has a 4xx status code
func (o *GetIndustryFacilitiesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get industry facilities bad request response has a 5xx status code
func (o *GetIndustryFacilitiesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry facilities bad request response a status code equal to that given
func (o *GetIndustryFacilitiesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIndustryFacilitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIndustryFacilitiesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIndustryFacilitiesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetIndustryFacilitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustryFacilitiesEnhanceYourCalm creates a GetIndustryFacilitiesEnhanceYourCalm with default headers values
func NewGetIndustryFacilitiesEnhanceYourCalm() *GetIndustryFacilitiesEnhanceYourCalm {
	return &GetIndustryFacilitiesEnhanceYourCalm{}
}

/*
GetIndustryFacilitiesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetIndustryFacilitiesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get industry facilities enhance your calm response has a 2xx status code
func (o *GetIndustryFacilitiesEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry facilities enhance your calm response has a 3xx status code
func (o *GetIndustryFacilitiesEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry facilities enhance your calm response has a 4xx status code
func (o *GetIndustryFacilitiesEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get industry facilities enhance your calm response has a 5xx status code
func (o *GetIndustryFacilitiesEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry facilities enhance your calm response a status code equal to that given
func (o *GetIndustryFacilitiesEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetIndustryFacilitiesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetIndustryFacilitiesEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetIndustryFacilitiesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetIndustryFacilitiesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustryFacilitiesInternalServerError creates a GetIndustryFacilitiesInternalServerError with default headers values
func NewGetIndustryFacilitiesInternalServerError() *GetIndustryFacilitiesInternalServerError {
	return &GetIndustryFacilitiesInternalServerError{}
}

/*
GetIndustryFacilitiesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetIndustryFacilitiesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get industry facilities internal server error response has a 2xx status code
func (o *GetIndustryFacilitiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry facilities internal server error response has a 3xx status code
func (o *GetIndustryFacilitiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry facilities internal server error response has a 4xx status code
func (o *GetIndustryFacilitiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry facilities internal server error response has a 5xx status code
func (o *GetIndustryFacilitiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get industry facilities internal server error response a status code equal to that given
func (o *GetIndustryFacilitiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIndustryFacilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIndustryFacilitiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIndustryFacilitiesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetIndustryFacilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustryFacilitiesServiceUnavailable creates a GetIndustryFacilitiesServiceUnavailable with default headers values
func NewGetIndustryFacilitiesServiceUnavailable() *GetIndustryFacilitiesServiceUnavailable {
	return &GetIndustryFacilitiesServiceUnavailable{}
}

/*
GetIndustryFacilitiesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetIndustryFacilitiesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get industry facilities service unavailable response has a 2xx status code
func (o *GetIndustryFacilitiesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry facilities service unavailable response has a 3xx status code
func (o *GetIndustryFacilitiesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry facilities service unavailable response has a 4xx status code
func (o *GetIndustryFacilitiesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry facilities service unavailable response has a 5xx status code
func (o *GetIndustryFacilitiesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get industry facilities service unavailable response a status code equal to that given
func (o *GetIndustryFacilitiesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIndustryFacilitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIndustryFacilitiesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIndustryFacilitiesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetIndustryFacilitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustryFacilitiesGatewayTimeout creates a GetIndustryFacilitiesGatewayTimeout with default headers values
func NewGetIndustryFacilitiesGatewayTimeout() *GetIndustryFacilitiesGatewayTimeout {
	return &GetIndustryFacilitiesGatewayTimeout{}
}

/*
GetIndustryFacilitiesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetIndustryFacilitiesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get industry facilities gateway timeout response has a 2xx status code
func (o *GetIndustryFacilitiesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry facilities gateway timeout response has a 3xx status code
func (o *GetIndustryFacilitiesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry facilities gateway timeout response has a 4xx status code
func (o *GetIndustryFacilitiesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry facilities gateway timeout response has a 5xx status code
func (o *GetIndustryFacilitiesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get industry facilities gateway timeout response a status code equal to that given
func (o *GetIndustryFacilitiesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIndustryFacilitiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIndustryFacilitiesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/industry/facilities/][%d] getIndustryFacilitiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIndustryFacilitiesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetIndustryFacilitiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetIndustryFacilitiesOKBodyItems0 get_industry_facilities_200_ok
//
// 200 ok object
swagger:model GetIndustryFacilitiesOKBodyItems0
*/
type GetIndustryFacilitiesOKBodyItems0 struct {

	// get_industry_facilities_facility_id
	//
	// ID of the facility
	// Required: true
	FacilityID *int64 `json:"facility_id"`

	// get_industry_facilities_owner_id
	//
	// Owner of the facility
	// Required: true
	OwnerID *int32 `json:"owner_id"`

	// get_industry_facilities_region_id
	//
	// Region ID where the facility is
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_industry_facilities_solar_system_id
	//
	// Solar system ID where the facility is
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_industry_facilities_tax
	//
	// Tax imposed by the facility
	Tax float32 `json:"tax,omitempty"`

	// get_industry_facilities_type_id
	//
	// Type ID of the facility
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get industry facilities o k body items0
func (o *GetIndustryFacilitiesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
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

func (o *GetIndustryFacilitiesOKBodyItems0) validateFacilityID(formats strfmt.Registry) error {

	if err := validate.Required("facility_id", "body", o.FacilityID); err != nil {
		return err
	}

	return nil
}

func (o *GetIndustryFacilitiesOKBodyItems0) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("owner_id", "body", o.OwnerID); err != nil {
		return err
	}

	return nil
}

func (o *GetIndustryFacilitiesOKBodyItems0) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", o.RegionID); err != nil {
		return err
	}

	return nil
}

func (o *GetIndustryFacilitiesOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetIndustryFacilitiesOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get industry facilities o k body items0 based on context it is used
func (o *GetIndustryFacilitiesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIndustryFacilitiesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIndustryFacilitiesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetIndustryFacilitiesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
