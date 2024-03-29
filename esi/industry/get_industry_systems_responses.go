// Code generated by go-swagger; DO NOT EDIT.

package industry

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

// GetIndustrySystemsReader is a Reader for the GetIndustrySystems structure.
type GetIndustrySystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIndustrySystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIndustrySystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetIndustrySystemsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetIndustrySystemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetIndustrySystemsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIndustrySystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIndustrySystemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIndustrySystemsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/industry/systems/] get_industry_systems", response, response.Code())
	}
}

// NewGetIndustrySystemsOK creates a GetIndustrySystemsOK with default headers values
func NewGetIndustrySystemsOK() *GetIndustrySystemsOK {
	return &GetIndustrySystemsOK{}
}

/*
GetIndustrySystemsOK describes a response with status code 200, with default header values.

A list of cost indicies
*/
type GetIndustrySystemsOK struct {

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

	Payload []*GetIndustrySystemsOKBodyItems0
}

// IsSuccess returns true when this get industry systems o k response has a 2xx status code
func (o *GetIndustrySystemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get industry systems o k response has a 3xx status code
func (o *GetIndustrySystemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry systems o k response has a 4xx status code
func (o *GetIndustrySystemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry systems o k response has a 5xx status code
func (o *GetIndustrySystemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry systems o k response a status code equal to that given
func (o *GetIndustrySystemsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get industry systems o k response
func (o *GetIndustrySystemsOK) Code() int {
	return 200
}

func (o *GetIndustrySystemsOK) Error() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsOK  %+v", 200, o.Payload)
}

func (o *GetIndustrySystemsOK) String() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsOK  %+v", 200, o.Payload)
}

func (o *GetIndustrySystemsOK) GetPayload() []*GetIndustrySystemsOKBodyItems0 {
	return o.Payload
}

func (o *GetIndustrySystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustrySystemsNotModified creates a GetIndustrySystemsNotModified with default headers values
func NewGetIndustrySystemsNotModified() *GetIndustrySystemsNotModified {
	return &GetIndustrySystemsNotModified{}
}

/*
GetIndustrySystemsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetIndustrySystemsNotModified struct {

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

// IsSuccess returns true when this get industry systems not modified response has a 2xx status code
func (o *GetIndustrySystemsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry systems not modified response has a 3xx status code
func (o *GetIndustrySystemsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get industry systems not modified response has a 4xx status code
func (o *GetIndustrySystemsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry systems not modified response has a 5xx status code
func (o *GetIndustrySystemsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry systems not modified response a status code equal to that given
func (o *GetIndustrySystemsNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get industry systems not modified response
func (o *GetIndustrySystemsNotModified) Code() int {
	return 304
}

func (o *GetIndustrySystemsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsNotModified ", 304)
}

func (o *GetIndustrySystemsNotModified) String() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsNotModified ", 304)
}

func (o *GetIndustrySystemsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustrySystemsBadRequest creates a GetIndustrySystemsBadRequest with default headers values
func NewGetIndustrySystemsBadRequest() *GetIndustrySystemsBadRequest {
	return &GetIndustrySystemsBadRequest{}
}

/*
GetIndustrySystemsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetIndustrySystemsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get industry systems bad request response has a 2xx status code
func (o *GetIndustrySystemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry systems bad request response has a 3xx status code
func (o *GetIndustrySystemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry systems bad request response has a 4xx status code
func (o *GetIndustrySystemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get industry systems bad request response has a 5xx status code
func (o *GetIndustrySystemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry systems bad request response a status code equal to that given
func (o *GetIndustrySystemsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get industry systems bad request response
func (o *GetIndustrySystemsBadRequest) Code() int {
	return 400
}

func (o *GetIndustrySystemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIndustrySystemsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIndustrySystemsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetIndustrySystemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsEnhanceYourCalm creates a GetIndustrySystemsEnhanceYourCalm with default headers values
func NewGetIndustrySystemsEnhanceYourCalm() *GetIndustrySystemsEnhanceYourCalm {
	return &GetIndustrySystemsEnhanceYourCalm{}
}

/*
GetIndustrySystemsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetIndustrySystemsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get industry systems enhance your calm response has a 2xx status code
func (o *GetIndustrySystemsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry systems enhance your calm response has a 3xx status code
func (o *GetIndustrySystemsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry systems enhance your calm response has a 4xx status code
func (o *GetIndustrySystemsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get industry systems enhance your calm response has a 5xx status code
func (o *GetIndustrySystemsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get industry systems enhance your calm response a status code equal to that given
func (o *GetIndustrySystemsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get industry systems enhance your calm response
func (o *GetIndustrySystemsEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetIndustrySystemsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetIndustrySystemsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetIndustrySystemsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetIndustrySystemsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsInternalServerError creates a GetIndustrySystemsInternalServerError with default headers values
func NewGetIndustrySystemsInternalServerError() *GetIndustrySystemsInternalServerError {
	return &GetIndustrySystemsInternalServerError{}
}

/*
GetIndustrySystemsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetIndustrySystemsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get industry systems internal server error response has a 2xx status code
func (o *GetIndustrySystemsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry systems internal server error response has a 3xx status code
func (o *GetIndustrySystemsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry systems internal server error response has a 4xx status code
func (o *GetIndustrySystemsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry systems internal server error response has a 5xx status code
func (o *GetIndustrySystemsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get industry systems internal server error response a status code equal to that given
func (o *GetIndustrySystemsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get industry systems internal server error response
func (o *GetIndustrySystemsInternalServerError) Code() int {
	return 500
}

func (o *GetIndustrySystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIndustrySystemsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIndustrySystemsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetIndustrySystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsServiceUnavailable creates a GetIndustrySystemsServiceUnavailable with default headers values
func NewGetIndustrySystemsServiceUnavailable() *GetIndustrySystemsServiceUnavailable {
	return &GetIndustrySystemsServiceUnavailable{}
}

/*
GetIndustrySystemsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetIndustrySystemsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get industry systems service unavailable response has a 2xx status code
func (o *GetIndustrySystemsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry systems service unavailable response has a 3xx status code
func (o *GetIndustrySystemsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry systems service unavailable response has a 4xx status code
func (o *GetIndustrySystemsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry systems service unavailable response has a 5xx status code
func (o *GetIndustrySystemsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get industry systems service unavailable response a status code equal to that given
func (o *GetIndustrySystemsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get industry systems service unavailable response
func (o *GetIndustrySystemsServiceUnavailable) Code() int {
	return 503
}

func (o *GetIndustrySystemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIndustrySystemsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIndustrySystemsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetIndustrySystemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsGatewayTimeout creates a GetIndustrySystemsGatewayTimeout with default headers values
func NewGetIndustrySystemsGatewayTimeout() *GetIndustrySystemsGatewayTimeout {
	return &GetIndustrySystemsGatewayTimeout{}
}

/*
GetIndustrySystemsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetIndustrySystemsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get industry systems gateway timeout response has a 2xx status code
func (o *GetIndustrySystemsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get industry systems gateway timeout response has a 3xx status code
func (o *GetIndustrySystemsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get industry systems gateway timeout response has a 4xx status code
func (o *GetIndustrySystemsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get industry systems gateway timeout response has a 5xx status code
func (o *GetIndustrySystemsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get industry systems gateway timeout response a status code equal to that given
func (o *GetIndustrySystemsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get industry systems gateway timeout response
func (o *GetIndustrySystemsGatewayTimeout) Code() int {
	return 504
}

func (o *GetIndustrySystemsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIndustrySystemsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/industry/systems/][%d] getIndustrySystemsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIndustrySystemsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetIndustrySystemsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetIndustrySystemsOKBodyItems0 get_industry_systems_200_ok
//
// 200 ok object
swagger:model GetIndustrySystemsOKBodyItems0
*/
type GetIndustrySystemsOKBodyItems0 struct {

	// get_industry_systems_cost_indices
	//
	// cost_indices array
	// Required: true
	// Max Items: 10
	CostIndices []*GetIndustrySystemsOKBodyItems0CostIndicesItems0 `json:"cost_indices"`

	// get_industry_systems_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`
}

// Validate validates this get industry systems o k body items0
func (o *GetIndustrySystemsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCostIndices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIndustrySystemsOKBodyItems0) validateCostIndices(formats strfmt.Registry) error {

	if err := validate.Required("cost_indices", "body", o.CostIndices); err != nil {
		return err
	}

	iCostIndicesSize := int64(len(o.CostIndices))

	if err := validate.MaxItems("cost_indices", "body", iCostIndicesSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(o.CostIndices); i++ {
		if swag.IsZero(o.CostIndices[i]) { // not required
			continue
		}

		if o.CostIndices[i] != nil {
			if err := o.CostIndices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_indices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_indices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetIndustrySystemsOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get industry systems o k body items0 based on the context it is used
func (o *GetIndustrySystemsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCostIndices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIndustrySystemsOKBodyItems0) contextValidateCostIndices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.CostIndices); i++ {

		if o.CostIndices[i] != nil {

			if swag.IsZero(o.CostIndices[i]) { // not required
				return nil
			}

			if err := o.CostIndices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_indices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_indices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetIndustrySystemsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetIndustrySystemsOKBodyItems0CostIndicesItems0 get_industry_systems_cost_indice
//
// cost_indice object
swagger:model GetIndustrySystemsOKBodyItems0CostIndicesItems0
*/
type GetIndustrySystemsOKBodyItems0CostIndicesItems0 struct {

	// get_industry_systems_activity
	//
	// activity string
	// Required: true
	// Enum: [copying duplicating invention manufacturing none reaction researching_material_efficiency researching_technology researching_time_efficiency reverse_engineering]
	Activity *string `json:"activity"`

	// get_industry_systems_cost_index
	//
	// cost_index number
	// Required: true
	CostIndex *float32 `json:"cost_index"`
}

// Validate validates this get industry systems o k body items0 cost indices items0
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCostIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["copying","duplicating","invention","manufacturing","none","reaction","researching_material_efficiency","researching_technology","researching_time_efficiency","reverse_engineering"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum = append(getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum, v)
	}
}

const (

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityCopying captures enum value "copying"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityCopying string = "copying"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityDuplicating captures enum value "duplicating"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityDuplicating string = "duplicating"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityInvention captures enum value "invention"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityInvention string = "invention"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityManufacturing captures enum value "manufacturing"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityManufacturing string = "manufacturing"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityNone captures enum value "none"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityNone string = "none"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityReaction captures enum value "reaction"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityReaction string = "reaction"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingMaterialEfficiency captures enum value "researching_material_efficiency"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingMaterialEfficiency string = "researching_material_efficiency"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTechnology captures enum value "researching_technology"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTechnology string = "researching_technology"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTimeEfficiency captures enum value "researching_time_efficiency"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityResearchingTimeEfficiency string = "researching_time_efficiency"

	// GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityReverseEngineering captures enum value "reverse_engineering"
	GetIndustrySystemsOKBodyItems0CostIndicesItems0ActivityReverseEngineering string = "reverse_engineering"
)

// prop value enum
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) validateActivityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getIndustrySystemsOKBodyItems0CostIndicesItems0TypeActivityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) validateActivity(formats strfmt.Registry) error {

	if err := validate.Required("activity", "body", o.Activity); err != nil {
		return err
	}

	// value enum
	if err := o.validateActivityEnum("activity", "body", *o.Activity); err != nil {
		return err
	}

	return nil
}

func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) validateCostIndex(formats strfmt.Registry) error {

	if err := validate.Required("cost_index", "body", o.CostIndex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get industry systems o k body items0 cost indices items0 based on context it is used
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetIndustrySystemsOKBodyItems0CostIndicesItems0) UnmarshalBinary(b []byte) error {
	var res GetIndustrySystemsOKBodyItems0CostIndicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
