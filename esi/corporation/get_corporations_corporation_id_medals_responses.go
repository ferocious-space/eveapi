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

// GetCorporationsCorporationIDMedalsReader is a Reader for the GetCorporationsCorporationIDMedals structure.
type GetCorporationsCorporationIDMedalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMedalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDMedalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDMedalsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDMedalsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDMedalsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDMedalsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDMedalsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDMedalsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDMedalsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDMedalsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMedalsOK creates a GetCorporationsCorporationIDMedalsOK with default headers values
func NewGetCorporationsCorporationIDMedalsOK() *GetCorporationsCorporationIDMedalsOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDMedalsOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDMedalsOK describes a response with status code 200, with default header values.

A list of medals
*/
type GetCorporationsCorporationIDMedalsOK struct {

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

	Payload []*GetCorporationsCorporationIDMedalsOKBodyItems0
}

func (o *GetCorporationsCorporationIDMedalsOK) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsOK) GetPayload() []*GetCorporationsCorporationIDMedalsOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsNotModified creates a GetCorporationsCorporationIDMedalsNotModified with default headers values
func NewGetCorporationsCorporationIDMedalsNotModified() *GetCorporationsCorporationIDMedalsNotModified {
	return &GetCorporationsCorporationIDMedalsNotModified{}
}

/* GetCorporationsCorporationIDMedalsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDMedalsNotModified struct {

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

func (o *GetCorporationsCorporationIDMedalsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMedalsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsBadRequest creates a GetCorporationsCorporationIDMedalsBadRequest with default headers values
func NewGetCorporationsCorporationIDMedalsBadRequest() *GetCorporationsCorporationIDMedalsBadRequest {
	return &GetCorporationsCorporationIDMedalsBadRequest{}
}

/* GetCorporationsCorporationIDMedalsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDMedalsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDMedalsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsUnauthorized creates a GetCorporationsCorporationIDMedalsUnauthorized with default headers values
func NewGetCorporationsCorporationIDMedalsUnauthorized() *GetCorporationsCorporationIDMedalsUnauthorized {
	return &GetCorporationsCorporationIDMedalsUnauthorized{}
}

/* GetCorporationsCorporationIDMedalsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMedalsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDMedalsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsForbidden creates a GetCorporationsCorporationIDMedalsForbidden with default headers values
func NewGetCorporationsCorporationIDMedalsForbidden() *GetCorporationsCorporationIDMedalsForbidden {
	return &GetCorporationsCorporationIDMedalsForbidden{}
}

/* GetCorporationsCorporationIDMedalsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMedalsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDMedalsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsEnhanceYourCalm creates a GetCorporationsCorporationIDMedalsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMedalsEnhanceYourCalm() *GetCorporationsCorporationIDMedalsEnhanceYourCalm {
	return &GetCorporationsCorporationIDMedalsEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDMedalsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDMedalsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDMedalsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsInternalServerError creates a GetCorporationsCorporationIDMedalsInternalServerError with default headers values
func NewGetCorporationsCorporationIDMedalsInternalServerError() *GetCorporationsCorporationIDMedalsInternalServerError {
	return &GetCorporationsCorporationIDMedalsInternalServerError{}
}

/* GetCorporationsCorporationIDMedalsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMedalsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDMedalsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsServiceUnavailable creates a GetCorporationsCorporationIDMedalsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMedalsServiceUnavailable() *GetCorporationsCorporationIDMedalsServiceUnavailable {
	return &GetCorporationsCorporationIDMedalsServiceUnavailable{}
}

/* GetCorporationsCorporationIDMedalsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMedalsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDMedalsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsGatewayTimeout creates a GetCorporationsCorporationIDMedalsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMedalsGatewayTimeout() *GetCorporationsCorporationIDMedalsGatewayTimeout {
	return &GetCorporationsCorporationIDMedalsGatewayTimeout{}
}

/* GetCorporationsCorporationIDMedalsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMedalsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDMedalsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/medals/][%d] getCorporationsCorporationIdMedalsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDMedalsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDMedalsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDMedalsOKBodyItems0 get_corporations_corporation_id_medals_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDMedalsOKBodyItems0
*/
type GetCorporationsCorporationIDMedalsOKBodyItems0 struct {

	// get_corporations_corporation_id_medals_created_at
	//
	// created_at string
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// get_corporations_corporation_id_medals_creator_id
	//
	// ID of the character who created this medal
	// Required: true
	CreatorID *int32 `json:"creator_id"`

	// get_corporations_corporation_id_medals_description
	//
	// description string
	// Required: true
	// Max Length: 1000
	Description *string `json:"description"`

	// get_corporations_corporation_id_medals_medal_id
	//
	// medal_id integer
	// Required: true
	MedalID *int32 `json:"medal_id"`

	// get_corporations_corporation_id_medals_title
	//
	// title string
	// Required: true
	// Max Length: 100
	Title *string `json:"title"`
}

// Validate validates this get corporations corporation ID medals o k body items0
func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMedalID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) validateCreatorID(formats strfmt.Registry) error {

	if err := validate.Required("creator_id", "body", o.CreatorID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", o.Description); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *o.Description, 1000); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) validateMedalID(formats strfmt.Registry) error {

	if err := validate.Required("medal_id", "body", o.MedalID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", o.Title); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *o.Title, 100); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID medals o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDMedalsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDMedalsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
