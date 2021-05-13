// Code generated by go-swagger; DO NOT EDIT.

package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// GetCorporationsCorporationIDStandingsReader is a Reader for the GetCorporationsCorporationIDStandings structure.
type GetCorporationsCorporationIDStandingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDStandingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDStandingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDStandingsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDStandingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDStandingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDStandingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDStandingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDStandingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDStandingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDStandingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDStandingsOK creates a GetCorporationsCorporationIDStandingsOK with default headers values
func NewGetCorporationsCorporationIDStandingsOK() *GetCorporationsCorporationIDStandingsOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDStandingsOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDStandingsOK describes a response with status code 200, with default header values.

A list of standings
*/
type GetCorporationsCorporationIDStandingsOK struct {

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

	/* Maximum page number

	   Format: int32
	   Default: 1
	*/
	XPages int32

	Payload []*GetCorporationsCorporationIDStandingsOKBodyItems0
}

func (o *GetCorporationsCorporationIDStandingsOK) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsOK) GetPayload() []*GetCorporationsCorporationIDStandingsOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// hydrates response header X-Pages
	hdrXPages := response.GetHeader("X-Pages")

	if hdrXPages != "" {
		valxPages, err := swag.ConvertInt32(hdrXPages)
		if err != nil {
			return errors.InvalidType("X-Pages", "header", "int32", hdrXPages)
		}
		o.XPages = valxPages
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStandingsNotModified creates a GetCorporationsCorporationIDStandingsNotModified with default headers values
func NewGetCorporationsCorporationIDStandingsNotModified() *GetCorporationsCorporationIDStandingsNotModified {
	return &GetCorporationsCorporationIDStandingsNotModified{}
}

/* GetCorporationsCorporationIDStandingsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDStandingsNotModified struct {

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

func (o *GetCorporationsCorporationIDStandingsNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDStandingsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDStandingsBadRequest creates a GetCorporationsCorporationIDStandingsBadRequest with default headers values
func NewGetCorporationsCorporationIDStandingsBadRequest() *GetCorporationsCorporationIDStandingsBadRequest {
	return &GetCorporationsCorporationIDStandingsBadRequest{}
}

/* GetCorporationsCorporationIDStandingsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDStandingsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDStandingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStandingsUnauthorized creates a GetCorporationsCorporationIDStandingsUnauthorized with default headers values
func NewGetCorporationsCorporationIDStandingsUnauthorized() *GetCorporationsCorporationIDStandingsUnauthorized {
	return &GetCorporationsCorporationIDStandingsUnauthorized{}
}

/* GetCorporationsCorporationIDStandingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDStandingsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDStandingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStandingsForbidden creates a GetCorporationsCorporationIDStandingsForbidden with default headers values
func NewGetCorporationsCorporationIDStandingsForbidden() *GetCorporationsCorporationIDStandingsForbidden {
	return &GetCorporationsCorporationIDStandingsForbidden{}
}

/* GetCorporationsCorporationIDStandingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDStandingsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDStandingsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStandingsEnhanceYourCalm creates a GetCorporationsCorporationIDStandingsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDStandingsEnhanceYourCalm() *GetCorporationsCorporationIDStandingsEnhanceYourCalm {
	return &GetCorporationsCorporationIDStandingsEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDStandingsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDStandingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDStandingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStandingsInternalServerError creates a GetCorporationsCorporationIDStandingsInternalServerError with default headers values
func NewGetCorporationsCorporationIDStandingsInternalServerError() *GetCorporationsCorporationIDStandingsInternalServerError {
	return &GetCorporationsCorporationIDStandingsInternalServerError{}
}

/* GetCorporationsCorporationIDStandingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDStandingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDStandingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStandingsServiceUnavailable creates a GetCorporationsCorporationIDStandingsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDStandingsServiceUnavailable() *GetCorporationsCorporationIDStandingsServiceUnavailable {
	return &GetCorporationsCorporationIDStandingsServiceUnavailable{}
}

/* GetCorporationsCorporationIDStandingsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDStandingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDStandingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStandingsGatewayTimeout creates a GetCorporationsCorporationIDStandingsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDStandingsGatewayTimeout() *GetCorporationsCorporationIDStandingsGatewayTimeout {
	return &GetCorporationsCorporationIDStandingsGatewayTimeout{}
}

/* GetCorporationsCorporationIDStandingsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDStandingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDStandingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/standings/][%d] getCorporationsCorporationIdStandingsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDStandingsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStandingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDStandingsOKBodyItems0 get_corporations_corporation_id_standings_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDStandingsOKBodyItems0
*/
type GetCorporationsCorporationIDStandingsOKBodyItems0 struct {

	// get_corporations_corporation_id_standings_from_id
	//
	// from_id integer
	// Required: true
	FromID *int32 `json:"from_id"`

	// get_corporations_corporation_id_standings_from_type
	//
	// from_type string
	// Required: true
	// Enum: [agent npc_corp faction]
	FromType *string `json:"from_type"`

	// get_corporations_corporation_id_standings_standing
	//
	// standing number
	// Required: true
	Standing *float32 `json:"standing"`
}

// Validate validates this get corporations corporation ID standings o k body items0
func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFromID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFromType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStanding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", o.FromID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStandingsOKBodyItems0TypeFromTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["agent","npc_corp","faction"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStandingsOKBodyItems0TypeFromTypePropEnum = append(getCorporationsCorporationIdStandingsOKBodyItems0TypeFromTypePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStandingsOKBodyItems0FromTypeAgent captures enum value "agent"
	GetCorporationsCorporationIDStandingsOKBodyItems0FromTypeAgent string = "agent"

	// GetCorporationsCorporationIDStandingsOKBodyItems0FromTypeNpcCorp captures enum value "npc_corp"
	GetCorporationsCorporationIDStandingsOKBodyItems0FromTypeNpcCorp string = "npc_corp"

	// GetCorporationsCorporationIDStandingsOKBodyItems0FromTypeFaction captures enum value "faction"
	GetCorporationsCorporationIDStandingsOKBodyItems0FromTypeFaction string = "faction"
)

// prop value enum
func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) validateFromTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdStandingsOKBodyItems0TypeFromTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) validateFromType(formats strfmt.Registry) error {

	if err := validate.Required("from_type", "body", o.FromType); err != nil {
		return err
	}

	// value enum
	if err := o.validateFromTypeEnum("from_type", "body", *o.FromType); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) validateStanding(formats strfmt.Registry) error {

	if err := validate.Required("standing", "body", o.Standing); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID standings o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDStandingsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDStandingsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
