// Code generated by go-swagger; DO NOT EDIT.

package killmails

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

// GetCorporationsCorporationIDKillmailsRecentReader is a Reader for the GetCorporationsCorporationIDKillmailsRecent structure.
type GetCorporationsCorporationIDKillmailsRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDKillmailsRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDKillmailsRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDKillmailsRecentNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDKillmailsRecentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDKillmailsRecentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDKillmailsRecentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDKillmailsRecentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDKillmailsRecentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDKillmailsRecentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDKillmailsRecentOK creates a GetCorporationsCorporationIDKillmailsRecentOK with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentOK() *GetCorporationsCorporationIDKillmailsRecentOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDKillmailsRecentOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDKillmailsRecentOK describes a response with status code 200, with default header values.

A list of killmail IDs and hashes
*/
type GetCorporationsCorporationIDKillmailsRecentOK struct {

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

	Payload []*GetCorporationsCorporationIDKillmailsRecentOKBodyItems0
}

func (o *GetCorporationsCorporationIDKillmailsRecentOK) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentOK) GetPayload() []*GetCorporationsCorporationIDKillmailsRecentOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDKillmailsRecentNotModified creates a GetCorporationsCorporationIDKillmailsRecentNotModified with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentNotModified() *GetCorporationsCorporationIDKillmailsRecentNotModified {
	return &GetCorporationsCorporationIDKillmailsRecentNotModified{}
}

/* GetCorporationsCorporationIDKillmailsRecentNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDKillmailsRecentNotModified struct {

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

func (o *GetCorporationsCorporationIDKillmailsRecentNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentNotModified ", 304)
}

func (o *GetCorporationsCorporationIDKillmailsRecentNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDKillmailsRecentBadRequest creates a GetCorporationsCorporationIDKillmailsRecentBadRequest with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentBadRequest() *GetCorporationsCorporationIDKillmailsRecentBadRequest {
	return &GetCorporationsCorporationIDKillmailsRecentBadRequest{}
}

/* GetCorporationsCorporationIDKillmailsRecentBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDKillmailsRecentBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDKillmailsRecentBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentUnauthorized creates a GetCorporationsCorporationIDKillmailsRecentUnauthorized with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentUnauthorized() *GetCorporationsCorporationIDKillmailsRecentUnauthorized {
	return &GetCorporationsCorporationIDKillmailsRecentUnauthorized{}
}

/* GetCorporationsCorporationIDKillmailsRecentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDKillmailsRecentUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDKillmailsRecentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentForbidden creates a GetCorporationsCorporationIDKillmailsRecentForbidden with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentForbidden() *GetCorporationsCorporationIDKillmailsRecentForbidden {
	return &GetCorporationsCorporationIDKillmailsRecentForbidden{}
}

/* GetCorporationsCorporationIDKillmailsRecentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDKillmailsRecentForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDKillmailsRecentForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm creates a GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm() *GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm {
	return &GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentInternalServerError creates a GetCorporationsCorporationIDKillmailsRecentInternalServerError with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentInternalServerError() *GetCorporationsCorporationIDKillmailsRecentInternalServerError {
	return &GetCorporationsCorporationIDKillmailsRecentInternalServerError{}
}

/* GetCorporationsCorporationIDKillmailsRecentInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDKillmailsRecentInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDKillmailsRecentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentServiceUnavailable creates a GetCorporationsCorporationIDKillmailsRecentServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentServiceUnavailable() *GetCorporationsCorporationIDKillmailsRecentServiceUnavailable {
	return &GetCorporationsCorporationIDKillmailsRecentServiceUnavailable{}
}

/* GetCorporationsCorporationIDKillmailsRecentServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDKillmailsRecentServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDKillmailsRecentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentGatewayTimeout creates a GetCorporationsCorporationIDKillmailsRecentGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentGatewayTimeout() *GetCorporationsCorporationIDKillmailsRecentGatewayTimeout {
	return &GetCorporationsCorporationIDKillmailsRecentGatewayTimeout{}
}

/* GetCorporationsCorporationIDKillmailsRecentGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDKillmailsRecentGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDKillmailsRecentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDKillmailsRecentGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDKillmailsRecentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDKillmailsRecentOKBodyItems0 get_corporations_corporation_id_killmails_recent_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDKillmailsRecentOKBodyItems0
*/
type GetCorporationsCorporationIDKillmailsRecentOKBodyItems0 struct {

	// get_corporations_corporation_id_killmails_recent_killmail_hash
	//
	// A hash of this killmail
	// Required: true
	KillmailHash *string `json:"killmail_hash"`

	// get_corporations_corporation_id_killmails_recent_killmail_id
	//
	// ID of this killmail
	// Required: true
	KillmailID *int32 `json:"killmail_id"`
}

// Validate validates this get corporations corporation ID killmails recent o k body items0
func (o *GetCorporationsCorporationIDKillmailsRecentOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKillmailHash(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKillmailID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDKillmailsRecentOKBodyItems0) validateKillmailHash(formats strfmt.Registry) error {

	if err := validate.Required("killmail_hash", "body", o.KillmailHash); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDKillmailsRecentOKBodyItems0) validateKillmailID(formats strfmt.Registry) error {

	if err := validate.Required("killmail_id", "body", o.KillmailID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID killmails recent o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDKillmailsRecentOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDKillmailsRecentOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDKillmailsRecentOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDKillmailsRecentOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
