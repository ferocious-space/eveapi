// Code generated by go-swagger; DO NOT EDIT.

package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/models"
)

// DeleteCharactersCharacterIDMailMailIDReader is a Reader for the DeleteCharactersCharacterIDMailMailID structure.
type DeleteCharactersCharacterIDMailMailIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCharactersCharacterIDMailMailIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCharactersCharacterIDMailMailIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCharactersCharacterIDMailMailIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCharactersCharacterIDMailMailIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCharactersCharacterIDMailMailIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewDeleteCharactersCharacterIDMailMailIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCharactersCharacterIDMailMailIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteCharactersCharacterIDMailMailIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteCharactersCharacterIDMailMailIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCharactersCharacterIDMailMailIDNoContent creates a DeleteCharactersCharacterIDMailMailIDNoContent with default headers values
func NewDeleteCharactersCharacterIDMailMailIDNoContent() *DeleteCharactersCharacterIDMailMailIDNoContent {
	return &DeleteCharactersCharacterIDMailMailIDNoContent{}
}

/*
DeleteCharactersCharacterIDMailMailIDNoContent describes a response with status code 204, with default header values.

Mail deleted
*/
type DeleteCharactersCharacterIDMailMailIDNoContent struct {
}

// IsSuccess returns true when this delete characters character Id mail mail Id no content response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete characters character Id mail mail Id no content response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id no content response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete characters character Id mail mail Id no content response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete characters character Id mail mail Id no content response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete characters character Id mail mail Id no content response
func (o *DeleteCharactersCharacterIDMailMailIDNoContent) Code() int {
	return 204
}

func (o *DeleteCharactersCharacterIDMailMailIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdNoContent ", 204)
}

func (o *DeleteCharactersCharacterIDMailMailIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdNoContent ", 204)
}

func (o *DeleteCharactersCharacterIDMailMailIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDBadRequest creates a DeleteCharactersCharacterIDMailMailIDBadRequest with default headers values
func NewDeleteCharactersCharacterIDMailMailIDBadRequest() *DeleteCharactersCharacterIDMailMailIDBadRequest {
	return &DeleteCharactersCharacterIDMailMailIDBadRequest{}
}

/*
DeleteCharactersCharacterIDMailMailIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteCharactersCharacterIDMailMailIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this delete characters character Id mail mail Id bad request response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete characters character Id mail mail Id bad request response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id bad request response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete characters character Id mail mail Id bad request response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete characters character Id mail mail Id bad request response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete characters character Id mail mail Id bad request response
func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) Code() int {
	return 400
}

func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDMailMailIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDUnauthorized creates a DeleteCharactersCharacterIDMailMailIDUnauthorized with default headers values
func NewDeleteCharactersCharacterIDMailMailIDUnauthorized() *DeleteCharactersCharacterIDMailMailIDUnauthorized {
	return &DeleteCharactersCharacterIDMailMailIDUnauthorized{}
}

/*
DeleteCharactersCharacterIDMailMailIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteCharactersCharacterIDMailMailIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this delete characters character Id mail mail Id unauthorized response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete characters character Id mail mail Id unauthorized response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id unauthorized response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete characters character Id mail mail Id unauthorized response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete characters character Id mail mail Id unauthorized response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete characters character Id mail mail Id unauthorized response
func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDMailMailIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDForbidden creates a DeleteCharactersCharacterIDMailMailIDForbidden with default headers values
func NewDeleteCharactersCharacterIDMailMailIDForbidden() *DeleteCharactersCharacterIDMailMailIDForbidden {
	return &DeleteCharactersCharacterIDMailMailIDForbidden{}
}

/*
DeleteCharactersCharacterIDMailMailIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCharactersCharacterIDMailMailIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this delete characters character Id mail mail Id forbidden response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete characters character Id mail mail Id forbidden response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id forbidden response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete characters character Id mail mail Id forbidden response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete characters character Id mail mail Id forbidden response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete characters character Id mail mail Id forbidden response
func (o *DeleteCharactersCharacterIDMailMailIDForbidden) Code() int {
	return 403
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDMailMailIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDEnhanceYourCalm creates a DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm with default headers values
func NewDeleteCharactersCharacterIDMailMailIDEnhanceYourCalm() *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm {
	return &DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm{}
}

/*
DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this delete characters character Id mail mail Id enhance your calm response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete characters character Id mail mail Id enhance your calm response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id enhance your calm response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete characters character Id mail mail Id enhance your calm response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this delete characters character Id mail mail Id enhance your calm response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the delete characters character Id mail mail Id enhance your calm response
func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDMailMailIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDInternalServerError creates a DeleteCharactersCharacterIDMailMailIDInternalServerError with default headers values
func NewDeleteCharactersCharacterIDMailMailIDInternalServerError() *DeleteCharactersCharacterIDMailMailIDInternalServerError {
	return &DeleteCharactersCharacterIDMailMailIDInternalServerError{}
}

/*
DeleteCharactersCharacterIDMailMailIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteCharactersCharacterIDMailMailIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this delete characters character Id mail mail Id internal server error response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete characters character Id mail mail Id internal server error response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id internal server error response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete characters character Id mail mail Id internal server error response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete characters character Id mail mail Id internal server error response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete characters character Id mail mail Id internal server error response
func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDMailMailIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDServiceUnavailable creates a DeleteCharactersCharacterIDMailMailIDServiceUnavailable with default headers values
func NewDeleteCharactersCharacterIDMailMailIDServiceUnavailable() *DeleteCharactersCharacterIDMailMailIDServiceUnavailable {
	return &DeleteCharactersCharacterIDMailMailIDServiceUnavailable{}
}

/*
DeleteCharactersCharacterIDMailMailIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type DeleteCharactersCharacterIDMailMailIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this delete characters character Id mail mail Id service unavailable response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete characters character Id mail mail Id service unavailable response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id service unavailable response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete characters character Id mail mail Id service unavailable response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete characters character Id mail mail Id service unavailable response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the delete characters character Id mail mail Id service unavailable response
func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) Code() int {
	return 503
}

func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDMailMailIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDMailMailIDGatewayTimeout creates a DeleteCharactersCharacterIDMailMailIDGatewayTimeout with default headers values
func NewDeleteCharactersCharacterIDMailMailIDGatewayTimeout() *DeleteCharactersCharacterIDMailMailIDGatewayTimeout {
	return &DeleteCharactersCharacterIDMailMailIDGatewayTimeout{}
}

/*
DeleteCharactersCharacterIDMailMailIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type DeleteCharactersCharacterIDMailMailIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this delete characters character Id mail mail Id gateway timeout response has a 2xx status code
func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete characters character Id mail mail Id gateway timeout response has a 3xx status code
func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete characters character Id mail mail Id gateway timeout response has a 4xx status code
func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete characters character Id mail mail Id gateway timeout response has a 5xx status code
func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete characters character Id mail mail Id gateway timeout response a status code equal to that given
func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the delete characters character Id mail mail Id gateway timeout response
func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) Code() int {
	return 504
}

func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/mail/{mail_id}/][%d] deleteCharactersCharacterIdMailMailIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDMailMailIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
