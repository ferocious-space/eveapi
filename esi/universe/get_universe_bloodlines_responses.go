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

// GetUniverseBloodlinesReader is a Reader for the GetUniverseBloodlines structure.
type GetUniverseBloodlinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseBloodlinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseBloodlinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseBloodlinesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseBloodlinesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseBloodlinesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseBloodlinesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseBloodlinesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseBloodlinesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseBloodlinesOK creates a GetUniverseBloodlinesOK with default headers values
func NewGetUniverseBloodlinesOK() *GetUniverseBloodlinesOK {
	return &GetUniverseBloodlinesOK{}
}

/*
GetUniverseBloodlinesOK describes a response with status code 200, with default header values.

A list of bloodlines
*/
type GetUniverseBloodlinesOK struct {

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

	Payload []*GetUniverseBloodlinesOKBodyItems0
}

// IsSuccess returns true when this get universe bloodlines o k response has a 2xx status code
func (o *GetUniverseBloodlinesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe bloodlines o k response has a 3xx status code
func (o *GetUniverseBloodlinesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe bloodlines o k response has a 4xx status code
func (o *GetUniverseBloodlinesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe bloodlines o k response has a 5xx status code
func (o *GetUniverseBloodlinesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe bloodlines o k response a status code equal to that given
func (o *GetUniverseBloodlinesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe bloodlines o k response
func (o *GetUniverseBloodlinesOK) Code() int {
	return 200
}

func (o *GetUniverseBloodlinesOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseBloodlinesOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseBloodlinesOK) GetPayload() []*GetUniverseBloodlinesOKBodyItems0 {
	return o.Payload
}

func (o *GetUniverseBloodlinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseBloodlinesNotModified creates a GetUniverseBloodlinesNotModified with default headers values
func NewGetUniverseBloodlinesNotModified() *GetUniverseBloodlinesNotModified {
	return &GetUniverseBloodlinesNotModified{}
}

/*
GetUniverseBloodlinesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseBloodlinesNotModified struct {

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

// IsSuccess returns true when this get universe bloodlines not modified response has a 2xx status code
func (o *GetUniverseBloodlinesNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe bloodlines not modified response has a 3xx status code
func (o *GetUniverseBloodlinesNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe bloodlines not modified response has a 4xx status code
func (o *GetUniverseBloodlinesNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe bloodlines not modified response has a 5xx status code
func (o *GetUniverseBloodlinesNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe bloodlines not modified response a status code equal to that given
func (o *GetUniverseBloodlinesNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe bloodlines not modified response
func (o *GetUniverseBloodlinesNotModified) Code() int {
	return 304
}

func (o *GetUniverseBloodlinesNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesNotModified ", 304)
}

func (o *GetUniverseBloodlinesNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesNotModified ", 304)
}

func (o *GetUniverseBloodlinesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseBloodlinesBadRequest creates a GetUniverseBloodlinesBadRequest with default headers values
func NewGetUniverseBloodlinesBadRequest() *GetUniverseBloodlinesBadRequest {
	return &GetUniverseBloodlinesBadRequest{}
}

/*
GetUniverseBloodlinesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseBloodlinesBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe bloodlines bad request response has a 2xx status code
func (o *GetUniverseBloodlinesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe bloodlines bad request response has a 3xx status code
func (o *GetUniverseBloodlinesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe bloodlines bad request response has a 4xx status code
func (o *GetUniverseBloodlinesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe bloodlines bad request response has a 5xx status code
func (o *GetUniverseBloodlinesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe bloodlines bad request response a status code equal to that given
func (o *GetUniverseBloodlinesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe bloodlines bad request response
func (o *GetUniverseBloodlinesBadRequest) Code() int {
	return 400
}

func (o *GetUniverseBloodlinesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseBloodlinesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseBloodlinesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseBloodlinesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseBloodlinesEnhanceYourCalm creates a GetUniverseBloodlinesEnhanceYourCalm with default headers values
func NewGetUniverseBloodlinesEnhanceYourCalm() *GetUniverseBloodlinesEnhanceYourCalm {
	return &GetUniverseBloodlinesEnhanceYourCalm{}
}

/*
GetUniverseBloodlinesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseBloodlinesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe bloodlines enhance your calm response has a 2xx status code
func (o *GetUniverseBloodlinesEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe bloodlines enhance your calm response has a 3xx status code
func (o *GetUniverseBloodlinesEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe bloodlines enhance your calm response has a 4xx status code
func (o *GetUniverseBloodlinesEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe bloodlines enhance your calm response has a 5xx status code
func (o *GetUniverseBloodlinesEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe bloodlines enhance your calm response a status code equal to that given
func (o *GetUniverseBloodlinesEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe bloodlines enhance your calm response
func (o *GetUniverseBloodlinesEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseBloodlinesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseBloodlinesEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseBloodlinesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseBloodlinesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseBloodlinesInternalServerError creates a GetUniverseBloodlinesInternalServerError with default headers values
func NewGetUniverseBloodlinesInternalServerError() *GetUniverseBloodlinesInternalServerError {
	return &GetUniverseBloodlinesInternalServerError{}
}

/*
GetUniverseBloodlinesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseBloodlinesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe bloodlines internal server error response has a 2xx status code
func (o *GetUniverseBloodlinesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe bloodlines internal server error response has a 3xx status code
func (o *GetUniverseBloodlinesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe bloodlines internal server error response has a 4xx status code
func (o *GetUniverseBloodlinesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe bloodlines internal server error response has a 5xx status code
func (o *GetUniverseBloodlinesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe bloodlines internal server error response a status code equal to that given
func (o *GetUniverseBloodlinesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe bloodlines internal server error response
func (o *GetUniverseBloodlinesInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseBloodlinesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseBloodlinesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseBloodlinesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseBloodlinesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseBloodlinesServiceUnavailable creates a GetUniverseBloodlinesServiceUnavailable with default headers values
func NewGetUniverseBloodlinesServiceUnavailable() *GetUniverseBloodlinesServiceUnavailable {
	return &GetUniverseBloodlinesServiceUnavailable{}
}

/*
GetUniverseBloodlinesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseBloodlinesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe bloodlines service unavailable response has a 2xx status code
func (o *GetUniverseBloodlinesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe bloodlines service unavailable response has a 3xx status code
func (o *GetUniverseBloodlinesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe bloodlines service unavailable response has a 4xx status code
func (o *GetUniverseBloodlinesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe bloodlines service unavailable response has a 5xx status code
func (o *GetUniverseBloodlinesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe bloodlines service unavailable response a status code equal to that given
func (o *GetUniverseBloodlinesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe bloodlines service unavailable response
func (o *GetUniverseBloodlinesServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseBloodlinesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseBloodlinesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseBloodlinesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseBloodlinesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseBloodlinesGatewayTimeout creates a GetUniverseBloodlinesGatewayTimeout with default headers values
func NewGetUniverseBloodlinesGatewayTimeout() *GetUniverseBloodlinesGatewayTimeout {
	return &GetUniverseBloodlinesGatewayTimeout{}
}

/*
GetUniverseBloodlinesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseBloodlinesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe bloodlines gateway timeout response has a 2xx status code
func (o *GetUniverseBloodlinesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe bloodlines gateway timeout response has a 3xx status code
func (o *GetUniverseBloodlinesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe bloodlines gateway timeout response has a 4xx status code
func (o *GetUniverseBloodlinesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe bloodlines gateway timeout response has a 5xx status code
func (o *GetUniverseBloodlinesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe bloodlines gateway timeout response a status code equal to that given
func (o *GetUniverseBloodlinesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe bloodlines gateway timeout response
func (o *GetUniverseBloodlinesGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseBloodlinesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseBloodlinesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/bloodlines/][%d] getUniverseBloodlinesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseBloodlinesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseBloodlinesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseBloodlinesOKBodyItems0 get_universe_bloodlines_200_ok
//
// 200 ok object
swagger:model GetUniverseBloodlinesOKBodyItems0
*/
type GetUniverseBloodlinesOKBodyItems0 struct {

	// get_universe_bloodlines_bloodline_id
	//
	// bloodline_id integer
	// Required: true
	BloodlineID *int32 `json:"bloodline_id"`

	// get_universe_bloodlines_charisma
	//
	// charisma integer
	// Required: true
	Charisma *int32 `json:"charisma"`

	// get_universe_bloodlines_corporation_id
	//
	// corporation_id integer
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_universe_bloodlines_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_universe_bloodlines_intelligence
	//
	// intelligence integer
	// Required: true
	Intelligence *int32 `json:"intelligence"`

	// get_universe_bloodlines_memory
	//
	// memory integer
	// Required: true
	Memory *int32 `json:"memory"`

	// get_universe_bloodlines_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_bloodlines_perception
	//
	// perception integer
	// Required: true
	Perception *int32 `json:"perception"`

	// get_universe_bloodlines_race_id
	//
	// race_id integer
	// Required: true
	RaceID *int32 `json:"race_id"`

	// get_universe_bloodlines_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`

	// get_universe_bloodlines_willpower
	//
	// willpower integer
	// Required: true
	Willpower *int32 `json:"willpower"`
}

// Validate validates this get universe bloodlines o k body items0
func (o *GetUniverseBloodlinesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBloodlineID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCharisma(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCorporationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIntelligence(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePerception(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWillpower(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateBloodlineID(formats strfmt.Registry) error {

	if err := validate.Required("bloodline_id", "body", o.BloodlineID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateCharisma(formats strfmt.Registry) error {

	if err := validate.Required("charisma", "body", o.Charisma); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", o.CorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateIntelligence(formats strfmt.Registry) error {

	if err := validate.Required("intelligence", "body", o.Intelligence); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", o.Memory); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validatePerception(formats strfmt.Registry) error {

	if err := validate.Required("perception", "body", o.Perception); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateRaceID(formats strfmt.Registry) error {

	if err := validate.Required("race_id", "body", o.RaceID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("ship_type_id", "body", o.ShipTypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseBloodlinesOKBodyItems0) validateWillpower(formats strfmt.Registry) error {

	if err := validate.Required("willpower", "body", o.Willpower); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe bloodlines o k body items0 based on context it is used
func (o *GetUniverseBloodlinesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseBloodlinesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseBloodlinesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseBloodlinesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
