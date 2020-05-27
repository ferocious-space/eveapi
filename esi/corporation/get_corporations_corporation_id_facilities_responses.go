// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// GetCorporationsCorporationIDFacilitiesReader is a Reader for the GetCorporationsCorporationIDFacilities structure.
type GetCorporationsCorporationIDFacilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDFacilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDFacilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDFacilitiesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDFacilitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDFacilitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDFacilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDFacilitiesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDFacilitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDFacilitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDFacilitiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDFacilitiesOK creates a GetCorporationsCorporationIDFacilitiesOK with default headers values
func NewGetCorporationsCorporationIDFacilitiesOK() *GetCorporationsCorporationIDFacilitiesOK {
	return &GetCorporationsCorporationIDFacilitiesOK{}
}

/* GetCorporationsCorporationIDFacilitiesOK describes a response with status code 200, with default header values.

List of corporation facilities
*/
type GetCorporationsCorporationIDFacilitiesOK struct {

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

	Payload []*GetCorporationsCorporationIDFacilitiesOKBodyItems0
}

func (o *GetCorporationsCorporationIDFacilitiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesOK) GetPayload() []*GetCorporationsCorporationIDFacilitiesOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDFacilitiesNotModified creates a GetCorporationsCorporationIDFacilitiesNotModified with default headers values
func NewGetCorporationsCorporationIDFacilitiesNotModified() *GetCorporationsCorporationIDFacilitiesNotModified {
	return &GetCorporationsCorporationIDFacilitiesNotModified{}
}

/* GetCorporationsCorporationIDFacilitiesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDFacilitiesNotModified struct {

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

func (o *GetCorporationsCorporationIDFacilitiesNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesNotModified ", 304)
}

func (o *GetCorporationsCorporationIDFacilitiesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDFacilitiesBadRequest creates a GetCorporationsCorporationIDFacilitiesBadRequest with default headers values
func NewGetCorporationsCorporationIDFacilitiesBadRequest() *GetCorporationsCorporationIDFacilitiesBadRequest {
	return &GetCorporationsCorporationIDFacilitiesBadRequest{}
}

/* GetCorporationsCorporationIDFacilitiesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDFacilitiesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDFacilitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesUnauthorized creates a GetCorporationsCorporationIDFacilitiesUnauthorized with default headers values
func NewGetCorporationsCorporationIDFacilitiesUnauthorized() *GetCorporationsCorporationIDFacilitiesUnauthorized {
	return &GetCorporationsCorporationIDFacilitiesUnauthorized{}
}

/* GetCorporationsCorporationIDFacilitiesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDFacilitiesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesForbidden creates a GetCorporationsCorporationIDFacilitiesForbidden with default headers values
func NewGetCorporationsCorporationIDFacilitiesForbidden() *GetCorporationsCorporationIDFacilitiesForbidden {
	return &GetCorporationsCorporationIDFacilitiesForbidden{}
}

/* GetCorporationsCorporationIDFacilitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDFacilitiesForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDFacilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesEnhanceYourCalm creates a GetCorporationsCorporationIDFacilitiesEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDFacilitiesEnhanceYourCalm() *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm {
	return &GetCorporationsCorporationIDFacilitiesEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDFacilitiesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDFacilitiesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesInternalServerError creates a GetCorporationsCorporationIDFacilitiesInternalServerError with default headers values
func NewGetCorporationsCorporationIDFacilitiesInternalServerError() *GetCorporationsCorporationIDFacilitiesInternalServerError {
	return &GetCorporationsCorporationIDFacilitiesInternalServerError{}
}

/* GetCorporationsCorporationIDFacilitiesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDFacilitiesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesServiceUnavailable creates a GetCorporationsCorporationIDFacilitiesServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDFacilitiesServiceUnavailable() *GetCorporationsCorporationIDFacilitiesServiceUnavailable {
	return &GetCorporationsCorporationIDFacilitiesServiceUnavailable{}
}

/* GetCorporationsCorporationIDFacilitiesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDFacilitiesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDFacilitiesGatewayTimeout creates a GetCorporationsCorporationIDFacilitiesGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDFacilitiesGatewayTimeout() *GetCorporationsCorporationIDFacilitiesGatewayTimeout {
	return &GetCorporationsCorporationIDFacilitiesGatewayTimeout{}
}

/* GetCorporationsCorporationIDFacilitiesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDFacilitiesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/facilities/][%d] getCorporationsCorporationIdFacilitiesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDFacilitiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDFacilitiesOKBodyItems0 get_corporations_corporation_id_facilities_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDFacilitiesOKBodyItems0
*/
type GetCorporationsCorporationIDFacilitiesOKBodyItems0 struct {

	// get_corporations_corporation_id_facilities_facility_id
	//
	// facility_id integer
	// Required: true
	FacilityID *int64 `json:"facility_id"`

	// get_corporations_corporation_id_facilities_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_corporations_corporation_id_facilities_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporations corporation ID facilities o k body items0
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFacilityID(formats); err != nil {
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

func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) validateFacilityID(formats strfmt.Registry) error {

	if err := validate.Required("facility_id", "body", o.FacilityID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID facilities o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDFacilitiesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDFacilitiesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
