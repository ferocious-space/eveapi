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

// GetSovereigntyMapReader is a Reader for the GetSovereigntyMap structure.
type GetSovereigntyMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSovereigntyMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetSovereigntyMapNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetSovereigntyMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetSovereigntyMapEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSovereigntyMapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSovereigntyMapServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSovereigntyMapGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/sovereignty/map/] get_sovereignty_map", response, response.Code())
	}
}

// NewGetSovereigntyMapOK creates a GetSovereigntyMapOK with default headers values
func NewGetSovereigntyMapOK() *GetSovereigntyMapOK {
	return &GetSovereigntyMapOK{}
}

/*
GetSovereigntyMapOK describes a response with status code 200, with default header values.

A list of sovereignty information for solar systems in New Eden
*/
type GetSovereigntyMapOK struct {

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

	Payload []*GetSovereigntyMapOKBodyItems0
}

// IsSuccess returns true when this get sovereignty map o k response has a 2xx status code
func (o *GetSovereigntyMapOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sovereignty map o k response has a 3xx status code
func (o *GetSovereigntyMapOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty map o k response has a 4xx status code
func (o *GetSovereigntyMapOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty map o k response has a 5xx status code
func (o *GetSovereigntyMapOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty map o k response a status code equal to that given
func (o *GetSovereigntyMapOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sovereignty map o k response
func (o *GetSovereigntyMapOK) Code() int {
	return 200
}

func (o *GetSovereigntyMapOK) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyMapOK) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyMapOK) GetPayload() []*GetSovereigntyMapOKBodyItems0 {
	return o.Payload
}

func (o *GetSovereigntyMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyMapNotModified creates a GetSovereigntyMapNotModified with default headers values
func NewGetSovereigntyMapNotModified() *GetSovereigntyMapNotModified {
	return &GetSovereigntyMapNotModified{}
}

/*
GetSovereigntyMapNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetSovereigntyMapNotModified struct {

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

// IsSuccess returns true when this get sovereignty map not modified response has a 2xx status code
func (o *GetSovereigntyMapNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty map not modified response has a 3xx status code
func (o *GetSovereigntyMapNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get sovereignty map not modified response has a 4xx status code
func (o *GetSovereigntyMapNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty map not modified response has a 5xx status code
func (o *GetSovereigntyMapNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty map not modified response a status code equal to that given
func (o *GetSovereigntyMapNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get sovereignty map not modified response
func (o *GetSovereigntyMapNotModified) Code() int {
	return 304
}

func (o *GetSovereigntyMapNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapNotModified ", 304)
}

func (o *GetSovereigntyMapNotModified) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapNotModified ", 304)
}

func (o *GetSovereigntyMapNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyMapBadRequest creates a GetSovereigntyMapBadRequest with default headers values
func NewGetSovereigntyMapBadRequest() *GetSovereigntyMapBadRequest {
	return &GetSovereigntyMapBadRequest{}
}

/*
GetSovereigntyMapBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSovereigntyMapBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get sovereignty map bad request response has a 2xx status code
func (o *GetSovereigntyMapBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty map bad request response has a 3xx status code
func (o *GetSovereigntyMapBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty map bad request response has a 4xx status code
func (o *GetSovereigntyMapBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sovereignty map bad request response has a 5xx status code
func (o *GetSovereigntyMapBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty map bad request response a status code equal to that given
func (o *GetSovereigntyMapBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get sovereignty map bad request response
func (o *GetSovereigntyMapBadRequest) Code() int {
	return 400
}

func (o *GetSovereigntyMapBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyMapBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyMapBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetSovereigntyMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapEnhanceYourCalm creates a GetSovereigntyMapEnhanceYourCalm with default headers values
func NewGetSovereigntyMapEnhanceYourCalm() *GetSovereigntyMapEnhanceYourCalm {
	return &GetSovereigntyMapEnhanceYourCalm{}
}

/*
GetSovereigntyMapEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetSovereigntyMapEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get sovereignty map enhance your calm response has a 2xx status code
func (o *GetSovereigntyMapEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty map enhance your calm response has a 3xx status code
func (o *GetSovereigntyMapEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty map enhance your calm response has a 4xx status code
func (o *GetSovereigntyMapEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sovereignty map enhance your calm response has a 5xx status code
func (o *GetSovereigntyMapEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty map enhance your calm response a status code equal to that given
func (o *GetSovereigntyMapEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get sovereignty map enhance your calm response
func (o *GetSovereigntyMapEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetSovereigntyMapEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyMapEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyMapEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetSovereigntyMapEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapInternalServerError creates a GetSovereigntyMapInternalServerError with default headers values
func NewGetSovereigntyMapInternalServerError() *GetSovereigntyMapInternalServerError {
	return &GetSovereigntyMapInternalServerError{}
}

/*
GetSovereigntyMapInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSovereigntyMapInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get sovereignty map internal server error response has a 2xx status code
func (o *GetSovereigntyMapInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty map internal server error response has a 3xx status code
func (o *GetSovereigntyMapInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty map internal server error response has a 4xx status code
func (o *GetSovereigntyMapInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty map internal server error response has a 5xx status code
func (o *GetSovereigntyMapInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty map internal server error response a status code equal to that given
func (o *GetSovereigntyMapInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get sovereignty map internal server error response
func (o *GetSovereigntyMapInternalServerError) Code() int {
	return 500
}

func (o *GetSovereigntyMapInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyMapInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyMapInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetSovereigntyMapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapServiceUnavailable creates a GetSovereigntyMapServiceUnavailable with default headers values
func NewGetSovereigntyMapServiceUnavailable() *GetSovereigntyMapServiceUnavailable {
	return &GetSovereigntyMapServiceUnavailable{}
}

/*
GetSovereigntyMapServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetSovereigntyMapServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get sovereignty map service unavailable response has a 2xx status code
func (o *GetSovereigntyMapServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty map service unavailable response has a 3xx status code
func (o *GetSovereigntyMapServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty map service unavailable response has a 4xx status code
func (o *GetSovereigntyMapServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty map service unavailable response has a 5xx status code
func (o *GetSovereigntyMapServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty map service unavailable response a status code equal to that given
func (o *GetSovereigntyMapServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get sovereignty map service unavailable response
func (o *GetSovereigntyMapServiceUnavailable) Code() int {
	return 503
}

func (o *GetSovereigntyMapServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyMapServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyMapServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetSovereigntyMapServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapGatewayTimeout creates a GetSovereigntyMapGatewayTimeout with default headers values
func NewGetSovereigntyMapGatewayTimeout() *GetSovereigntyMapGatewayTimeout {
	return &GetSovereigntyMapGatewayTimeout{}
}

/*
GetSovereigntyMapGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetSovereigntyMapGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get sovereignty map gateway timeout response has a 2xx status code
func (o *GetSovereigntyMapGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty map gateway timeout response has a 3xx status code
func (o *GetSovereigntyMapGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty map gateway timeout response has a 4xx status code
func (o *GetSovereigntyMapGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty map gateway timeout response has a 5xx status code
func (o *GetSovereigntyMapGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty map gateway timeout response a status code equal to that given
func (o *GetSovereigntyMapGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get sovereignty map gateway timeout response
func (o *GetSovereigntyMapGatewayTimeout) Code() int {
	return 504
}

func (o *GetSovereigntyMapGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyMapGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/map/][%d] getSovereigntyMapGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyMapGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetSovereigntyMapGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetSovereigntyMapOKBodyItems0 get_sovereignty_map_200_ok
//
// 200 ok object
swagger:model GetSovereigntyMapOKBodyItems0
*/
type GetSovereigntyMapOKBodyItems0 struct {

	// get_sovereignty_map_alliance_id
	//
	// alliance_id integer
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_sovereignty_map_corporation_id
	//
	// corporation_id integer
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_sovereignty_map_faction_id
	//
	// faction_id integer
	FactionID int32 `json:"faction_id,omitempty"`

	// get_sovereignty_map_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get sovereignty map o k body items0
func (o *GetSovereigntyMapOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyMapOKBodyItems0) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get sovereignty map o k body items0 based on context it is used
func (o *GetSovereigntyMapOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSovereigntyMapOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSovereigntyMapOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyMapOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
