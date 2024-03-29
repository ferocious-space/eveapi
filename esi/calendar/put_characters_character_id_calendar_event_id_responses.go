// Code generated by go-swagger; DO NOT EDIT.

package calendar

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

// PutCharactersCharacterIDCalendarEventIDReader is a Reader for the PutCharactersCharacterIDCalendarEventID structure.
type PutCharactersCharacterIDCalendarEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCharactersCharacterIDCalendarEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutCharactersCharacterIDCalendarEventIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCharactersCharacterIDCalendarEventIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutCharactersCharacterIDCalendarEventIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutCharactersCharacterIDCalendarEventIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPutCharactersCharacterIDCalendarEventIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutCharactersCharacterIDCalendarEventIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutCharactersCharacterIDCalendarEventIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutCharactersCharacterIDCalendarEventIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v3/characters/{character_id}/calendar/{event_id}/] put_characters_character_id_calendar_event_id", response, response.Code())
	}
}

// NewPutCharactersCharacterIDCalendarEventIDNoContent creates a PutCharactersCharacterIDCalendarEventIDNoContent with default headers values
func NewPutCharactersCharacterIDCalendarEventIDNoContent() *PutCharactersCharacterIDCalendarEventIDNoContent {
	return &PutCharactersCharacterIDCalendarEventIDNoContent{}
}

/*
PutCharactersCharacterIDCalendarEventIDNoContent describes a response with status code 204, with default header values.

Event updated
*/
type PutCharactersCharacterIDCalendarEventIDNoContent struct {
}

// IsSuccess returns true when this put characters character Id calendar event Id no content response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put characters character Id calendar event Id no content response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id no content response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put characters character Id calendar event Id no content response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put characters character Id calendar event Id no content response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put characters character Id calendar event Id no content response
func (o *PutCharactersCharacterIDCalendarEventIDNoContent) Code() int {
	return 204
}

func (o *PutCharactersCharacterIDCalendarEventIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdNoContent ", 204)
}

func (o *PutCharactersCharacterIDCalendarEventIDNoContent) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdNoContent ", 204)
}

func (o *PutCharactersCharacterIDCalendarEventIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCharactersCharacterIDCalendarEventIDBadRequest creates a PutCharactersCharacterIDCalendarEventIDBadRequest with default headers values
func NewPutCharactersCharacterIDCalendarEventIDBadRequest() *PutCharactersCharacterIDCalendarEventIDBadRequest {
	return &PutCharactersCharacterIDCalendarEventIDBadRequest{}
}

/*
PutCharactersCharacterIDCalendarEventIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutCharactersCharacterIDCalendarEventIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this put characters character Id calendar event Id bad request response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put characters character Id calendar event Id bad request response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id bad request response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put characters character Id calendar event Id bad request response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put characters character Id calendar event Id bad request response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put characters character Id calendar event Id bad request response
func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) Code() int {
	return 400
}

func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PutCharactersCharacterIDCalendarEventIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDCalendarEventIDUnauthorized creates a PutCharactersCharacterIDCalendarEventIDUnauthorized with default headers values
func NewPutCharactersCharacterIDCalendarEventIDUnauthorized() *PutCharactersCharacterIDCalendarEventIDUnauthorized {
	return &PutCharactersCharacterIDCalendarEventIDUnauthorized{}
}

/*
PutCharactersCharacterIDCalendarEventIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutCharactersCharacterIDCalendarEventIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this put characters character Id calendar event Id unauthorized response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put characters character Id calendar event Id unauthorized response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id unauthorized response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put characters character Id calendar event Id unauthorized response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put characters character Id calendar event Id unauthorized response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put characters character Id calendar event Id unauthorized response
func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) Code() int {
	return 401
}

func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PutCharactersCharacterIDCalendarEventIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDCalendarEventIDForbidden creates a PutCharactersCharacterIDCalendarEventIDForbidden with default headers values
func NewPutCharactersCharacterIDCalendarEventIDForbidden() *PutCharactersCharacterIDCalendarEventIDForbidden {
	return &PutCharactersCharacterIDCalendarEventIDForbidden{}
}

/*
PutCharactersCharacterIDCalendarEventIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutCharactersCharacterIDCalendarEventIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this put characters character Id calendar event Id forbidden response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put characters character Id calendar event Id forbidden response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id forbidden response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put characters character Id calendar event Id forbidden response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put characters character Id calendar event Id forbidden response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put characters character Id calendar event Id forbidden response
func (o *PutCharactersCharacterIDCalendarEventIDForbidden) Code() int {
	return 403
}

func (o *PutCharactersCharacterIDCalendarEventIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdForbidden  %+v", 403, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDForbidden) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdForbidden  %+v", 403, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PutCharactersCharacterIDCalendarEventIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDCalendarEventIDEnhanceYourCalm creates a PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm with default headers values
func NewPutCharactersCharacterIDCalendarEventIDEnhanceYourCalm() *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm {
	return &PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm{}
}

/*
PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this put characters character Id calendar event Id enhance your calm response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put characters character Id calendar event Id enhance your calm response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id enhance your calm response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this put characters character Id calendar event Id enhance your calm response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this put characters character Id calendar event Id enhance your calm response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the put characters character Id calendar event Id enhance your calm response
func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PutCharactersCharacterIDCalendarEventIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDCalendarEventIDInternalServerError creates a PutCharactersCharacterIDCalendarEventIDInternalServerError with default headers values
func NewPutCharactersCharacterIDCalendarEventIDInternalServerError() *PutCharactersCharacterIDCalendarEventIDInternalServerError {
	return &PutCharactersCharacterIDCalendarEventIDInternalServerError{}
}

/*
PutCharactersCharacterIDCalendarEventIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PutCharactersCharacterIDCalendarEventIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this put characters character Id calendar event Id internal server error response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put characters character Id calendar event Id internal server error response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id internal server error response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put characters character Id calendar event Id internal server error response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put characters character Id calendar event Id internal server error response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put characters character Id calendar event Id internal server error response
func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) Code() int {
	return 500
}

func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PutCharactersCharacterIDCalendarEventIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDCalendarEventIDServiceUnavailable creates a PutCharactersCharacterIDCalendarEventIDServiceUnavailable with default headers values
func NewPutCharactersCharacterIDCalendarEventIDServiceUnavailable() *PutCharactersCharacterIDCalendarEventIDServiceUnavailable {
	return &PutCharactersCharacterIDCalendarEventIDServiceUnavailable{}
}

/*
PutCharactersCharacterIDCalendarEventIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PutCharactersCharacterIDCalendarEventIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this put characters character Id calendar event Id service unavailable response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put characters character Id calendar event Id service unavailable response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id service unavailable response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put characters character Id calendar event Id service unavailable response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put characters character Id calendar event Id service unavailable response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the put characters character Id calendar event Id service unavailable response
func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) Code() int {
	return 503
}

func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PutCharactersCharacterIDCalendarEventIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCharactersCharacterIDCalendarEventIDGatewayTimeout creates a PutCharactersCharacterIDCalendarEventIDGatewayTimeout with default headers values
func NewPutCharactersCharacterIDCalendarEventIDGatewayTimeout() *PutCharactersCharacterIDCalendarEventIDGatewayTimeout {
	return &PutCharactersCharacterIDCalendarEventIDGatewayTimeout{}
}

/*
PutCharactersCharacterIDCalendarEventIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PutCharactersCharacterIDCalendarEventIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this put characters character Id calendar event Id gateway timeout response has a 2xx status code
func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put characters character Id calendar event Id gateway timeout response has a 3xx status code
func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put characters character Id calendar event Id gateway timeout response has a 4xx status code
func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put characters character Id calendar event Id gateway timeout response has a 5xx status code
func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put characters character Id calendar event Id gateway timeout response a status code equal to that given
func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the put characters character Id calendar event Id gateway timeout response
func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) Code() int {
	return 504
}

func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v3/characters/{character_id}/calendar/{event_id}/][%d] putCharactersCharacterIdCalendarEventIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PutCharactersCharacterIDCalendarEventIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutCharactersCharacterIDCalendarEventIDBody put_characters_character_id_calendar_event_id_response
//
// response object
swagger:model PutCharactersCharacterIDCalendarEventIDBody
*/
type PutCharactersCharacterIDCalendarEventIDBody struct {

	// put_characters_character_id_calendar_event_id_response_response
	//
	// response string
	// Required: true
	// Enum: [accepted declined tentative]
	Response *string `json:"response"`
}

// Validate validates this put characters character ID calendar event ID body
func (o *PutCharactersCharacterIDCalendarEventIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var putCharactersCharacterIdCalendarEventIdBodyTypeResponsePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["accepted","declined","tentative"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putCharactersCharacterIdCalendarEventIdBodyTypeResponsePropEnum = append(putCharactersCharacterIdCalendarEventIdBodyTypeResponsePropEnum, v)
	}
}

const (

	// PutCharactersCharacterIDCalendarEventIDBodyResponseAccepted captures enum value "accepted"
	PutCharactersCharacterIDCalendarEventIDBodyResponseAccepted string = "accepted"

	// PutCharactersCharacterIDCalendarEventIDBodyResponseDeclined captures enum value "declined"
	PutCharactersCharacterIDCalendarEventIDBodyResponseDeclined string = "declined"

	// PutCharactersCharacterIDCalendarEventIDBodyResponseTentative captures enum value "tentative"
	PutCharactersCharacterIDCalendarEventIDBodyResponseTentative string = "tentative"
)

// prop value enum
func (o *PutCharactersCharacterIDCalendarEventIDBody) validateResponseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, putCharactersCharacterIdCalendarEventIdBodyTypeResponsePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PutCharactersCharacterIDCalendarEventIDBody) validateResponse(formats strfmt.Registry) error {

	if err := validate.Required("response"+"."+"response", "body", o.Response); err != nil {
		return err
	}

	// value enum
	if err := o.validateResponseEnum("response"+"."+"response", "body", *o.Response); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put characters character ID calendar event ID body based on context it is used
func (o *PutCharactersCharacterIDCalendarEventIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCharactersCharacterIDCalendarEventIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCharactersCharacterIDCalendarEventIDBody) UnmarshalBinary(b []byte) error {
	var res PutCharactersCharacterIDCalendarEventIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
