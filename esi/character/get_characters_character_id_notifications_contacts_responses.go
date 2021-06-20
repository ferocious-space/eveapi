// Code generated by go-swagger; DO NOT EDIT.

package character

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

// GetCharactersCharacterIDNotificationsContactsReader is a Reader for the GetCharactersCharacterIDNotificationsContacts structure.
type GetCharactersCharacterIDNotificationsContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDNotificationsContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDNotificationsContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDNotificationsContactsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDNotificationsContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDNotificationsContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDNotificationsContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDNotificationsContactsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDNotificationsContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDNotificationsContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDNotificationsContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDNotificationsContactsOK creates a GetCharactersCharacterIDNotificationsContactsOK with default headers values
func NewGetCharactersCharacterIDNotificationsContactsOK() *GetCharactersCharacterIDNotificationsContactsOK {
	return &GetCharactersCharacterIDNotificationsContactsOK{}
}

/* GetCharactersCharacterIDNotificationsContactsOK describes a response with status code 200, with default header values.

A list of contact notifications
*/
type GetCharactersCharacterIDNotificationsContactsOK struct {

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

	Payload []*GetCharactersCharacterIDNotificationsContactsOKBodyItems0
}

func (o *GetCharactersCharacterIDNotificationsContactsOK) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsOK) GetPayload() []*GetCharactersCharacterIDNotificationsContactsOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDNotificationsContactsNotModified creates a GetCharactersCharacterIDNotificationsContactsNotModified with default headers values
func NewGetCharactersCharacterIDNotificationsContactsNotModified() *GetCharactersCharacterIDNotificationsContactsNotModified {
	return &GetCharactersCharacterIDNotificationsContactsNotModified{}
}

/* GetCharactersCharacterIDNotificationsContactsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDNotificationsContactsNotModified struct {

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

func (o *GetCharactersCharacterIDNotificationsContactsNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsNotModified ", 304)
}

func (o *GetCharactersCharacterIDNotificationsContactsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDNotificationsContactsBadRequest creates a GetCharactersCharacterIDNotificationsContactsBadRequest with default headers values
func NewGetCharactersCharacterIDNotificationsContactsBadRequest() *GetCharactersCharacterIDNotificationsContactsBadRequest {
	return &GetCharactersCharacterIDNotificationsContactsBadRequest{}
}

/* GetCharactersCharacterIDNotificationsContactsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDNotificationsContactsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDNotificationsContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDNotificationsContactsUnauthorized creates a GetCharactersCharacterIDNotificationsContactsUnauthorized with default headers values
func NewGetCharactersCharacterIDNotificationsContactsUnauthorized() *GetCharactersCharacterIDNotificationsContactsUnauthorized {
	return &GetCharactersCharacterIDNotificationsContactsUnauthorized{}
}

/* GetCharactersCharacterIDNotificationsContactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDNotificationsContactsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDNotificationsContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDNotificationsContactsForbidden creates a GetCharactersCharacterIDNotificationsContactsForbidden with default headers values
func NewGetCharactersCharacterIDNotificationsContactsForbidden() *GetCharactersCharacterIDNotificationsContactsForbidden {
	return &GetCharactersCharacterIDNotificationsContactsForbidden{}
}

/* GetCharactersCharacterIDNotificationsContactsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDNotificationsContactsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDNotificationsContactsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDNotificationsContactsEnhanceYourCalm creates a GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDNotificationsContactsEnhanceYourCalm() *GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm {
	return &GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm{}
}

/* GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDNotificationsContactsInternalServerError creates a GetCharactersCharacterIDNotificationsContactsInternalServerError with default headers values
func NewGetCharactersCharacterIDNotificationsContactsInternalServerError() *GetCharactersCharacterIDNotificationsContactsInternalServerError {
	return &GetCharactersCharacterIDNotificationsContactsInternalServerError{}
}

/* GetCharactersCharacterIDNotificationsContactsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDNotificationsContactsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDNotificationsContactsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDNotificationsContactsServiceUnavailable creates a GetCharactersCharacterIDNotificationsContactsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDNotificationsContactsServiceUnavailable() *GetCharactersCharacterIDNotificationsContactsServiceUnavailable {
	return &GetCharactersCharacterIDNotificationsContactsServiceUnavailable{}
}

/* GetCharactersCharacterIDNotificationsContactsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDNotificationsContactsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDNotificationsContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDNotificationsContactsGatewayTimeout creates a GetCharactersCharacterIDNotificationsContactsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDNotificationsContactsGatewayTimeout() *GetCharactersCharacterIDNotificationsContactsGatewayTimeout {
	return &GetCharactersCharacterIDNotificationsContactsGatewayTimeout{}
}

/* GetCharactersCharacterIDNotificationsContactsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDNotificationsContactsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDNotificationsContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/notifications/contacts/][%d] getCharactersCharacterIdNotificationsContactsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDNotificationsContactsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDNotificationsContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDNotificationsContactsOKBodyItems0 get_characters_character_id_notifications_contacts_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDNotificationsContactsOKBodyItems0
*/
type GetCharactersCharacterIDNotificationsContactsOKBodyItems0 struct {

	// get_characters_character_id_notifications_contacts_message
	//
	// message string
	// Required: true
	Message *string `json:"message"`

	// get_characters_character_id_notifications_contacts_notification_id
	//
	// notification_id integer
	// Required: true
	NotificationID *int32 `json:"notification_id"`

	// get_characters_character_id_notifications_contacts_send_date
	//
	// send_date string
	// Required: true
	// Format: date-time
	SendDate *strfmt.DateTime `json:"send_date"`

	// get_characters_character_id_notifications_contacts_sender_character_id
	//
	// sender_character_id integer
	// Required: true
	SenderCharacterID *int32 `json:"sender_character_id"`

	// get_characters_character_id_notifications_contacts_standing_level
	//
	// A number representing the standing level the receiver has been added at by the sender. The standing levels are as follows: -10 -> Terrible | -5 -> Bad |  0 -> Neutral |  5 -> Good |  10 -> Excellent
	// Required: true
	StandingLevel *float32 `json:"standing_level"`
}

// Validate validates this get characters character ID notifications contacts o k body items0
func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNotificationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSendDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSenderCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStandingLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) validateNotificationID(formats strfmt.Registry) error {

	if err := validate.Required("notification_id", "body", o.NotificationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) validateSendDate(formats strfmt.Registry) error {

	if err := validate.Required("send_date", "body", o.SendDate); err != nil {
		return err
	}

	if err := validate.FormatOf("send_date", "body", "date-time", o.SendDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) validateSenderCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("sender_character_id", "body", o.SenderCharacterID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) validateStandingLevel(formats strfmt.Registry) error {

	if err := validate.Required("standing_level", "body", o.StandingLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID notifications contacts o k body items0 based on context it is used
func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDNotificationsContactsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDNotificationsContactsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
