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

// GetUniverseAncestriesReader is a Reader for the GetUniverseAncestries structure.
type GetUniverseAncestriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseAncestriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseAncestriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseAncestriesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseAncestriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseAncestriesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseAncestriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseAncestriesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseAncestriesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseAncestriesOK creates a GetUniverseAncestriesOK with default headers values
func NewGetUniverseAncestriesOK() *GetUniverseAncestriesOK {
	return &GetUniverseAncestriesOK{}
}

/*
GetUniverseAncestriesOK describes a response with status code 200, with default header values.

A list of ancestries
*/
type GetUniverseAncestriesOK struct {

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

	Payload []*GetUniverseAncestriesOKBodyItems0
}

// IsSuccess returns true when this get universe ancestries o k response has a 2xx status code
func (o *GetUniverseAncestriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe ancestries o k response has a 3xx status code
func (o *GetUniverseAncestriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe ancestries o k response has a 4xx status code
func (o *GetUniverseAncestriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe ancestries o k response has a 5xx status code
func (o *GetUniverseAncestriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe ancestries o k response a status code equal to that given
func (o *GetUniverseAncestriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUniverseAncestriesOK) Error() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseAncestriesOK) String() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseAncestriesOK) GetPayload() []*GetUniverseAncestriesOKBodyItems0 {
	return o.Payload
}

func (o *GetUniverseAncestriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseAncestriesNotModified creates a GetUniverseAncestriesNotModified with default headers values
func NewGetUniverseAncestriesNotModified() *GetUniverseAncestriesNotModified {
	return &GetUniverseAncestriesNotModified{}
}

/*
GetUniverseAncestriesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseAncestriesNotModified struct {

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

// IsSuccess returns true when this get universe ancestries not modified response has a 2xx status code
func (o *GetUniverseAncestriesNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe ancestries not modified response has a 3xx status code
func (o *GetUniverseAncestriesNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe ancestries not modified response has a 4xx status code
func (o *GetUniverseAncestriesNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe ancestries not modified response has a 5xx status code
func (o *GetUniverseAncestriesNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe ancestries not modified response a status code equal to that given
func (o *GetUniverseAncestriesNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetUniverseAncestriesNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesNotModified ", 304)
}

func (o *GetUniverseAncestriesNotModified) String() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesNotModified ", 304)
}

func (o *GetUniverseAncestriesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseAncestriesBadRequest creates a GetUniverseAncestriesBadRequest with default headers values
func NewGetUniverseAncestriesBadRequest() *GetUniverseAncestriesBadRequest {
	return &GetUniverseAncestriesBadRequest{}
}

/*
GetUniverseAncestriesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseAncestriesBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe ancestries bad request response has a 2xx status code
func (o *GetUniverseAncestriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe ancestries bad request response has a 3xx status code
func (o *GetUniverseAncestriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe ancestries bad request response has a 4xx status code
func (o *GetUniverseAncestriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe ancestries bad request response has a 5xx status code
func (o *GetUniverseAncestriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe ancestries bad request response a status code equal to that given
func (o *GetUniverseAncestriesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUniverseAncestriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseAncestriesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseAncestriesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseAncestriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesEnhanceYourCalm creates a GetUniverseAncestriesEnhanceYourCalm with default headers values
func NewGetUniverseAncestriesEnhanceYourCalm() *GetUniverseAncestriesEnhanceYourCalm {
	return &GetUniverseAncestriesEnhanceYourCalm{}
}

/*
GetUniverseAncestriesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseAncestriesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe ancestries enhance your calm response has a 2xx status code
func (o *GetUniverseAncestriesEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe ancestries enhance your calm response has a 3xx status code
func (o *GetUniverseAncestriesEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe ancestries enhance your calm response has a 4xx status code
func (o *GetUniverseAncestriesEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe ancestries enhance your calm response has a 5xx status code
func (o *GetUniverseAncestriesEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe ancestries enhance your calm response a status code equal to that given
func (o *GetUniverseAncestriesEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *GetUniverseAncestriesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseAncestriesEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseAncestriesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseAncestriesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesInternalServerError creates a GetUniverseAncestriesInternalServerError with default headers values
func NewGetUniverseAncestriesInternalServerError() *GetUniverseAncestriesInternalServerError {
	return &GetUniverseAncestriesInternalServerError{}
}

/*
GetUniverseAncestriesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseAncestriesInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe ancestries internal server error response has a 2xx status code
func (o *GetUniverseAncestriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe ancestries internal server error response has a 3xx status code
func (o *GetUniverseAncestriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe ancestries internal server error response has a 4xx status code
func (o *GetUniverseAncestriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe ancestries internal server error response has a 5xx status code
func (o *GetUniverseAncestriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe ancestries internal server error response a status code equal to that given
func (o *GetUniverseAncestriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUniverseAncestriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseAncestriesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseAncestriesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseAncestriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesServiceUnavailable creates a GetUniverseAncestriesServiceUnavailable with default headers values
func NewGetUniverseAncestriesServiceUnavailable() *GetUniverseAncestriesServiceUnavailable {
	return &GetUniverseAncestriesServiceUnavailable{}
}

/*
GetUniverseAncestriesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseAncestriesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe ancestries service unavailable response has a 2xx status code
func (o *GetUniverseAncestriesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe ancestries service unavailable response has a 3xx status code
func (o *GetUniverseAncestriesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe ancestries service unavailable response has a 4xx status code
func (o *GetUniverseAncestriesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe ancestries service unavailable response has a 5xx status code
func (o *GetUniverseAncestriesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe ancestries service unavailable response a status code equal to that given
func (o *GetUniverseAncestriesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUniverseAncestriesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseAncestriesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseAncestriesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseAncestriesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesGatewayTimeout creates a GetUniverseAncestriesGatewayTimeout with default headers values
func NewGetUniverseAncestriesGatewayTimeout() *GetUniverseAncestriesGatewayTimeout {
	return &GetUniverseAncestriesGatewayTimeout{}
}

/*
GetUniverseAncestriesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseAncestriesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe ancestries gateway timeout response has a 2xx status code
func (o *GetUniverseAncestriesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe ancestries gateway timeout response has a 3xx status code
func (o *GetUniverseAncestriesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe ancestries gateway timeout response has a 4xx status code
func (o *GetUniverseAncestriesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe ancestries gateway timeout response has a 5xx status code
func (o *GetUniverseAncestriesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe ancestries gateway timeout response a status code equal to that given
func (o *GetUniverseAncestriesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUniverseAncestriesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseAncestriesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/universe/ancestries/][%d] getUniverseAncestriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseAncestriesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseAncestriesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseAncestriesOKBodyItems0 get_universe_ancestries_200_ok
//
// 200 ok object
swagger:model GetUniverseAncestriesOKBodyItems0
*/
type GetUniverseAncestriesOKBodyItems0 struct {

	// get_universe_ancestries_bloodline_id
	//
	// The bloodline associated with this ancestry
	// Required: true
	BloodlineID *int32 `json:"bloodline_id"`

	// get_universe_ancestries_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_universe_ancestries_icon_id
	//
	// icon_id integer
	IconID int32 `json:"icon_id,omitempty"`

	// get_universe_ancestries_id
	//
	// id integer
	// Required: true
	ID *int32 `json:"id"`

	// get_universe_ancestries_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_ancestries_short_description
	//
	// short_description string
	ShortDescription string `json:"short_description,omitempty"`
}

// Validate validates this get universe ancestries o k body items0
func (o *GetUniverseAncestriesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBloodlineID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUniverseAncestriesOKBodyItems0) validateBloodlineID(formats strfmt.Registry) error {

	if err := validate.Required("bloodline_id", "body", o.BloodlineID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseAncestriesOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseAncestriesOKBodyItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseAncestriesOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe ancestries o k body items0 based on context it is used
func (o *GetUniverseAncestriesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseAncestriesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseAncestriesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseAncestriesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
