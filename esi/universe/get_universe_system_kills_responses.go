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

// GetUniverseSystemKillsReader is a Reader for the GetUniverseSystemKills structure.
type GetUniverseSystemKillsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseSystemKillsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUniverseSystemKillsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetUniverseSystemKillsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetUniverseSystemKillsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetUniverseSystemKillsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUniverseSystemKillsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUniverseSystemKillsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUniverseSystemKillsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUniverseSystemKillsOK creates a GetUniverseSystemKillsOK with default headers values
func NewGetUniverseSystemKillsOK() *GetUniverseSystemKillsOK {
	return &GetUniverseSystemKillsOK{}
}

/*
GetUniverseSystemKillsOK describes a response with status code 200, with default header values.

A list of systems and number of ship, pod and NPC kills
*/
type GetUniverseSystemKillsOK struct {

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

	Payload []*GetUniverseSystemKillsOKBodyItems0
}

// IsSuccess returns true when this get universe system kills o k response has a 2xx status code
func (o *GetUniverseSystemKillsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get universe system kills o k response has a 3xx status code
func (o *GetUniverseSystemKillsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe system kills o k response has a 4xx status code
func (o *GetUniverseSystemKillsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe system kills o k response has a 5xx status code
func (o *GetUniverseSystemKillsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe system kills o k response a status code equal to that given
func (o *GetUniverseSystemKillsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get universe system kills o k response
func (o *GetUniverseSystemKillsOK) Code() int {
	return 200
}

func (o *GetUniverseSystemKillsOK) Error() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseSystemKillsOK) String() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsOK  %+v", 200, o.Payload)
}

func (o *GetUniverseSystemKillsOK) GetPayload() []*GetUniverseSystemKillsOKBodyItems0 {
	return o.Payload
}

func (o *GetUniverseSystemKillsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseSystemKillsNotModified creates a GetUniverseSystemKillsNotModified with default headers values
func NewGetUniverseSystemKillsNotModified() *GetUniverseSystemKillsNotModified {
	return &GetUniverseSystemKillsNotModified{}
}

/*
GetUniverseSystemKillsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetUniverseSystemKillsNotModified struct {

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

// IsSuccess returns true when this get universe system kills not modified response has a 2xx status code
func (o *GetUniverseSystemKillsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe system kills not modified response has a 3xx status code
func (o *GetUniverseSystemKillsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get universe system kills not modified response has a 4xx status code
func (o *GetUniverseSystemKillsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe system kills not modified response has a 5xx status code
func (o *GetUniverseSystemKillsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe system kills not modified response a status code equal to that given
func (o *GetUniverseSystemKillsNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get universe system kills not modified response
func (o *GetUniverseSystemKillsNotModified) Code() int {
	return 304
}

func (o *GetUniverseSystemKillsNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsNotModified ", 304)
}

func (o *GetUniverseSystemKillsNotModified) String() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsNotModified ", 304)
}

func (o *GetUniverseSystemKillsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseSystemKillsBadRequest creates a GetUniverseSystemKillsBadRequest with default headers values
func NewGetUniverseSystemKillsBadRequest() *GetUniverseSystemKillsBadRequest {
	return &GetUniverseSystemKillsBadRequest{}
}

/*
GetUniverseSystemKillsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetUniverseSystemKillsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get universe system kills bad request response has a 2xx status code
func (o *GetUniverseSystemKillsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe system kills bad request response has a 3xx status code
func (o *GetUniverseSystemKillsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe system kills bad request response has a 4xx status code
func (o *GetUniverseSystemKillsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe system kills bad request response has a 5xx status code
func (o *GetUniverseSystemKillsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe system kills bad request response a status code equal to that given
func (o *GetUniverseSystemKillsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get universe system kills bad request response
func (o *GetUniverseSystemKillsBadRequest) Code() int {
	return 400
}

func (o *GetUniverseSystemKillsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseSystemKillsBadRequest) String() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseSystemKillsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetUniverseSystemKillsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemKillsEnhanceYourCalm creates a GetUniverseSystemKillsEnhanceYourCalm with default headers values
func NewGetUniverseSystemKillsEnhanceYourCalm() *GetUniverseSystemKillsEnhanceYourCalm {
	return &GetUniverseSystemKillsEnhanceYourCalm{}
}

/*
GetUniverseSystemKillsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetUniverseSystemKillsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get universe system kills enhance your calm response has a 2xx status code
func (o *GetUniverseSystemKillsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe system kills enhance your calm response has a 3xx status code
func (o *GetUniverseSystemKillsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe system kills enhance your calm response has a 4xx status code
func (o *GetUniverseSystemKillsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get universe system kills enhance your calm response has a 5xx status code
func (o *GetUniverseSystemKillsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get universe system kills enhance your calm response a status code equal to that given
func (o *GetUniverseSystemKillsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get universe system kills enhance your calm response
func (o *GetUniverseSystemKillsEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetUniverseSystemKillsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseSystemKillsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseSystemKillsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetUniverseSystemKillsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemKillsInternalServerError creates a GetUniverseSystemKillsInternalServerError with default headers values
func NewGetUniverseSystemKillsInternalServerError() *GetUniverseSystemKillsInternalServerError {
	return &GetUniverseSystemKillsInternalServerError{}
}

/*
GetUniverseSystemKillsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetUniverseSystemKillsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get universe system kills internal server error response has a 2xx status code
func (o *GetUniverseSystemKillsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe system kills internal server error response has a 3xx status code
func (o *GetUniverseSystemKillsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe system kills internal server error response has a 4xx status code
func (o *GetUniverseSystemKillsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe system kills internal server error response has a 5xx status code
func (o *GetUniverseSystemKillsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe system kills internal server error response a status code equal to that given
func (o *GetUniverseSystemKillsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get universe system kills internal server error response
func (o *GetUniverseSystemKillsInternalServerError) Code() int {
	return 500
}

func (o *GetUniverseSystemKillsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseSystemKillsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseSystemKillsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetUniverseSystemKillsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemKillsServiceUnavailable creates a GetUniverseSystemKillsServiceUnavailable with default headers values
func NewGetUniverseSystemKillsServiceUnavailable() *GetUniverseSystemKillsServiceUnavailable {
	return &GetUniverseSystemKillsServiceUnavailable{}
}

/*
GetUniverseSystemKillsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetUniverseSystemKillsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get universe system kills service unavailable response has a 2xx status code
func (o *GetUniverseSystemKillsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe system kills service unavailable response has a 3xx status code
func (o *GetUniverseSystemKillsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe system kills service unavailable response has a 4xx status code
func (o *GetUniverseSystemKillsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe system kills service unavailable response has a 5xx status code
func (o *GetUniverseSystemKillsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe system kills service unavailable response a status code equal to that given
func (o *GetUniverseSystemKillsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get universe system kills service unavailable response
func (o *GetUniverseSystemKillsServiceUnavailable) Code() int {
	return 503
}

func (o *GetUniverseSystemKillsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseSystemKillsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseSystemKillsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetUniverseSystemKillsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseSystemKillsGatewayTimeout creates a GetUniverseSystemKillsGatewayTimeout with default headers values
func NewGetUniverseSystemKillsGatewayTimeout() *GetUniverseSystemKillsGatewayTimeout {
	return &GetUniverseSystemKillsGatewayTimeout{}
}

/*
GetUniverseSystemKillsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetUniverseSystemKillsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get universe system kills gateway timeout response has a 2xx status code
func (o *GetUniverseSystemKillsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get universe system kills gateway timeout response has a 3xx status code
func (o *GetUniverseSystemKillsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get universe system kills gateway timeout response has a 4xx status code
func (o *GetUniverseSystemKillsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get universe system kills gateway timeout response has a 5xx status code
func (o *GetUniverseSystemKillsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get universe system kills gateway timeout response a status code equal to that given
func (o *GetUniverseSystemKillsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get universe system kills gateway timeout response
func (o *GetUniverseSystemKillsGatewayTimeout) Code() int {
	return 504
}

func (o *GetUniverseSystemKillsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseSystemKillsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v2/universe/system_kills/][%d] getUniverseSystemKillsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseSystemKillsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetUniverseSystemKillsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetUniverseSystemKillsOKBodyItems0 get_universe_system_kills_200_ok
//
// 200 ok object
swagger:model GetUniverseSystemKillsOKBodyItems0
*/
type GetUniverseSystemKillsOKBodyItems0 struct {

	// get_universe_system_kills_npc_kills
	//
	// Number of NPC ships killed in this system
	// Required: true
	NpcKills *int32 `json:"npc_kills"`

	// get_universe_system_kills_pod_kills
	//
	// Number of pods killed in this system
	// Required: true
	PodKills *int32 `json:"pod_kills"`

	// get_universe_system_kills_ship_kills
	//
	// Number of player ships killed in this system
	// Required: true
	ShipKills *int32 `json:"ship_kills"`

	// get_universe_system_kills_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe system kills o k body items0
func (o *GetUniverseSystemKillsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNpcKills(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePodKills(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateShipKills(formats); err != nil {
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

func (o *GetUniverseSystemKillsOKBodyItems0) validateNpcKills(formats strfmt.Registry) error {

	if err := validate.Required("npc_kills", "body", o.NpcKills); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemKillsOKBodyItems0) validatePodKills(formats strfmt.Registry) error {

	if err := validate.Required("pod_kills", "body", o.PodKills); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemKillsOKBodyItems0) validateShipKills(formats strfmt.Registry) error {

	if err := validate.Required("ship_kills", "body", o.ShipKills); err != nil {
		return err
	}

	return nil
}

func (o *GetUniverseSystemKillsOKBodyItems0) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get universe system kills o k body items0 based on context it is used
func (o *GetUniverseSystemKillsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUniverseSystemKillsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUniverseSystemKillsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemKillsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
