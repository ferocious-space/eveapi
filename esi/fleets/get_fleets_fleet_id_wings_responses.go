// Code generated by go-swagger; DO NOT EDIT.

package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
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

// GetFleetsFleetIDWingsReader is a Reader for the GetFleetsFleetIDWings structure.
type GetFleetsFleetIDWingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetsFleetIDWingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFleetsFleetIDWingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetFleetsFleetIDWingsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetFleetsFleetIDWingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFleetsFleetIDWingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFleetsFleetIDWingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFleetsFleetIDWingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetFleetsFleetIDWingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFleetsFleetIDWingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFleetsFleetIDWingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFleetsFleetIDWingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFleetsFleetIDWingsOK creates a GetFleetsFleetIDWingsOK with default headers values
func NewGetFleetsFleetIDWingsOK() *GetFleetsFleetIDWingsOK {
	return &GetFleetsFleetIDWingsOK{}
}

/* GetFleetsFleetIDWingsOK describes a response with status code 200, with default header values.

A list of fleet wings
*/
type GetFleetsFleetIDWingsOK struct {

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

	Payload []*GetFleetsFleetIDWingsOKBodyItems0
}

func (o *GetFleetsFleetIDWingsOK) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsOK  %+v", 200, o.Payload)
}
func (o *GetFleetsFleetIDWingsOK) GetPayload() []*GetFleetsFleetIDWingsOKBodyItems0 {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFleetsFleetIDWingsNotModified creates a GetFleetsFleetIDWingsNotModified with default headers values
func NewGetFleetsFleetIDWingsNotModified() *GetFleetsFleetIDWingsNotModified {
	return &GetFleetsFleetIDWingsNotModified{}
}

/* GetFleetsFleetIDWingsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetFleetsFleetIDWingsNotModified struct {

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

func (o *GetFleetsFleetIDWingsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsNotModified ", 304)
}

func (o *GetFleetsFleetIDWingsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFleetsFleetIDWingsBadRequest creates a GetFleetsFleetIDWingsBadRequest with default headers values
func NewGetFleetsFleetIDWingsBadRequest() *GetFleetsFleetIDWingsBadRequest {
	return &GetFleetsFleetIDWingsBadRequest{}
}

/* GetFleetsFleetIDWingsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFleetsFleetIDWingsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetFleetsFleetIDWingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsBadRequest  %+v", 400, o.Payload)
}
func (o *GetFleetsFleetIDWingsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDWingsUnauthorized creates a GetFleetsFleetIDWingsUnauthorized with default headers values
func NewGetFleetsFleetIDWingsUnauthorized() *GetFleetsFleetIDWingsUnauthorized {
	return &GetFleetsFleetIDWingsUnauthorized{}
}

/* GetFleetsFleetIDWingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFleetsFleetIDWingsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetFleetsFleetIDWingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFleetsFleetIDWingsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDWingsForbidden creates a GetFleetsFleetIDWingsForbidden with default headers values
func NewGetFleetsFleetIDWingsForbidden() *GetFleetsFleetIDWingsForbidden {
	return &GetFleetsFleetIDWingsForbidden{}
}

/* GetFleetsFleetIDWingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFleetsFleetIDWingsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetFleetsFleetIDWingsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsForbidden  %+v", 403, o.Payload)
}
func (o *GetFleetsFleetIDWingsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDWingsNotFound creates a GetFleetsFleetIDWingsNotFound with default headers values
func NewGetFleetsFleetIDWingsNotFound() *GetFleetsFleetIDWingsNotFound {
	return &GetFleetsFleetIDWingsNotFound{}
}

/* GetFleetsFleetIDWingsNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type GetFleetsFleetIDWingsNotFound struct {
	Payload *GetFleetsFleetIDWingsNotFoundBody
}

func (o *GetFleetsFleetIDWingsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsNotFound  %+v", 404, o.Payload)
}
func (o *GetFleetsFleetIDWingsNotFound) GetPayload() *GetFleetsFleetIDWingsNotFoundBody {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetFleetsFleetIDWingsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDWingsEnhanceYourCalm creates a GetFleetsFleetIDWingsEnhanceYourCalm with default headers values
func NewGetFleetsFleetIDWingsEnhanceYourCalm() *GetFleetsFleetIDWingsEnhanceYourCalm {
	return &GetFleetsFleetIDWingsEnhanceYourCalm{}
}

/* GetFleetsFleetIDWingsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetFleetsFleetIDWingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetFleetsFleetIDWingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetFleetsFleetIDWingsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDWingsInternalServerError creates a GetFleetsFleetIDWingsInternalServerError with default headers values
func NewGetFleetsFleetIDWingsInternalServerError() *GetFleetsFleetIDWingsInternalServerError {
	return &GetFleetsFleetIDWingsInternalServerError{}
}

/* GetFleetsFleetIDWingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetFleetsFleetIDWingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetFleetsFleetIDWingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFleetsFleetIDWingsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDWingsServiceUnavailable creates a GetFleetsFleetIDWingsServiceUnavailable with default headers values
func NewGetFleetsFleetIDWingsServiceUnavailable() *GetFleetsFleetIDWingsServiceUnavailable {
	return &GetFleetsFleetIDWingsServiceUnavailable{}
}

/* GetFleetsFleetIDWingsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetFleetsFleetIDWingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetFleetsFleetIDWingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetFleetsFleetIDWingsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetsFleetIDWingsGatewayTimeout creates a GetFleetsFleetIDWingsGatewayTimeout with default headers values
func NewGetFleetsFleetIDWingsGatewayTimeout() *GetFleetsFleetIDWingsGatewayTimeout {
	return &GetFleetsFleetIDWingsGatewayTimeout{}
}

/* GetFleetsFleetIDWingsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetFleetsFleetIDWingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetFleetsFleetIDWingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/fleets/{fleet_id}/wings/][%d] getFleetsFleetIdWingsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetFleetsFleetIDWingsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetFleetsFleetIDWingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetFleetsFleetIDWingsNotFoundBody get_fleets_fleet_id_wings_not_found
//
// Not found
swagger:model GetFleetsFleetIDWingsNotFoundBody
*/
type GetFleetsFleetIDWingsNotFoundBody struct {

	// get_fleets_fleet_id_wings_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get fleets fleet ID wings not found body
func (o *GetFleetsFleetIDWingsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get fleets fleet ID wings not found body based on context it is used
func (o *GetFleetsFleetIDWingsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDWingsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDWingsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDWingsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFleetsFleetIDWingsOKBodyItems0 get_fleets_fleet_id_wings_200_ok
//
// 200 ok object
swagger:model GetFleetsFleetIDWingsOKBodyItems0
*/
type GetFleetsFleetIDWingsOKBodyItems0 struct {

	// get_fleets_fleet_id_wings_id
	//
	// id integer
	// Required: true
	ID *int64 `json:"id"`

	// get_fleets_fleet_id_wings_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_fleets_fleet_id_wings_squads
	//
	// squads array
	// Required: true
	// Max Items: 25
	Squads []*GetFleetsFleetIDWingsOKBodyItems0SquadsItems0 `json:"squads"`
}

// Validate validates this get fleets fleet ID wings o k body items0
func (o *GetFleetsFleetIDWingsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSquads(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFleetsFleetIDWingsOKBodyItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDWingsOKBodyItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDWingsOKBodyItems0) validateSquads(formats strfmt.Registry) error {

	if err := validate.Required("squads", "body", o.Squads); err != nil {
		return err
	}

	iSquadsSize := int64(len(o.Squads))

	if err := validate.MaxItems("squads", "body", iSquadsSize, 25); err != nil {
		return err
	}

	for i := 0; i < len(o.Squads); i++ {
		if swag.IsZero(o.Squads[i]) { // not required
			continue
		}

		if o.Squads[i] != nil {
			if err := o.Squads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("squads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get fleets fleet ID wings o k body items0 based on the context it is used
func (o *GetFleetsFleetIDWingsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSquads(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFleetsFleetIDWingsOKBodyItems0) contextValidateSquads(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Squads); i++ {

		if o.Squads[i] != nil {
			if err := o.Squads[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("squads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDWingsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDWingsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDWingsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetFleetsFleetIDWingsOKBodyItems0SquadsItems0 get_fleets_fleet_id_wings_squad
//
// squad object
swagger:model GetFleetsFleetIDWingsOKBodyItems0SquadsItems0
*/
type GetFleetsFleetIDWingsOKBodyItems0SquadsItems0 struct {

	// get_fleets_fleet_id_wings_squad_id
	//
	// id integer
	// Required: true
	ID *int64 `json:"id"`

	// get_fleets_fleet_id_wings_squad_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this get fleets fleet ID wings o k body items0 squads items0
func (o *GetFleetsFleetIDWingsOKBodyItems0SquadsItems0) Validate(formats strfmt.Registry) error {
	var res []error

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

func (o *GetFleetsFleetIDWingsOKBodyItems0SquadsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetFleetsFleetIDWingsOKBodyItems0SquadsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get fleets fleet ID wings o k body items0 squads items0 based on context it is used
func (o *GetFleetsFleetIDWingsOKBodyItems0SquadsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetFleetsFleetIDWingsOKBodyItems0SquadsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFleetsFleetIDWingsOKBodyItems0SquadsItems0) UnmarshalBinary(b []byte) error {
	var res GetFleetsFleetIDWingsOKBodyItems0SquadsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
