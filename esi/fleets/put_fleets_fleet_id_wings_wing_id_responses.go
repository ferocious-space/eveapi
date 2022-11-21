// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// PutFleetsFleetIDWingsWingIDReader is a Reader for the PutFleetsFleetIDWingsWingID structure.
type PutFleetsFleetIDWingsWingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetsFleetIDWingsWingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutFleetsFleetIDWingsWingIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutFleetsFleetIDWingsWingIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutFleetsFleetIDWingsWingIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutFleetsFleetIDWingsWingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutFleetsFleetIDWingsWingIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPutFleetsFleetIDWingsWingIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutFleetsFleetIDWingsWingIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutFleetsFleetIDWingsWingIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutFleetsFleetIDWingsWingIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutFleetsFleetIDWingsWingIDNoContent creates a PutFleetsFleetIDWingsWingIDNoContent with default headers values
func NewPutFleetsFleetIDWingsWingIDNoContent() *PutFleetsFleetIDWingsWingIDNoContent {
	return &PutFleetsFleetIDWingsWingIDNoContent{}
}

/*
PutFleetsFleetIDWingsWingIDNoContent describes a response with status code 204, with default header values.

Wing renamed
*/
type PutFleetsFleetIDWingsWingIDNoContent struct {
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id no content response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id no content response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id no content response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id wings wing Id no content response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id wings wing Id no content response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PutFleetsFleetIDWingsWingIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdNoContent ", 204)
}

func (o *PutFleetsFleetIDWingsWingIDNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdNoContent ", 204)
}

func (o *PutFleetsFleetIDWingsWingIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetsFleetIDWingsWingIDBadRequest creates a PutFleetsFleetIDWingsWingIDBadRequest with default headers values
func NewPutFleetsFleetIDWingsWingIDBadRequest() *PutFleetsFleetIDWingsWingIDBadRequest {
	return &PutFleetsFleetIDWingsWingIDBadRequest{}
}

/*
PutFleetsFleetIDWingsWingIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutFleetsFleetIDWingsWingIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id bad request response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id bad request response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id bad request response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id wings wing Id bad request response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id wings wing Id bad request response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutFleetsFleetIDWingsWingIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDUnauthorized creates a PutFleetsFleetIDWingsWingIDUnauthorized with default headers values
func NewPutFleetsFleetIDWingsWingIDUnauthorized() *PutFleetsFleetIDWingsWingIDUnauthorized {
	return &PutFleetsFleetIDWingsWingIDUnauthorized{}
}

/*
PutFleetsFleetIDWingsWingIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutFleetsFleetIDWingsWingIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id unauthorized response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id unauthorized response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id unauthorized response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id wings wing Id unauthorized response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id wings wing Id unauthorized response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutFleetsFleetIDWingsWingIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDForbidden creates a PutFleetsFleetIDWingsWingIDForbidden with default headers values
func NewPutFleetsFleetIDWingsWingIDForbidden() *PutFleetsFleetIDWingsWingIDForbidden {
	return &PutFleetsFleetIDWingsWingIDForbidden{}
}

/*
PutFleetsFleetIDWingsWingIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutFleetsFleetIDWingsWingIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id forbidden response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id forbidden response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id forbidden response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id wings wing Id forbidden response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id wings wing Id forbidden response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutFleetsFleetIDWingsWingIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDNotFound creates a PutFleetsFleetIDWingsWingIDNotFound with default headers values
func NewPutFleetsFleetIDWingsWingIDNotFound() *PutFleetsFleetIDWingsWingIDNotFound {
	return &PutFleetsFleetIDWingsWingIDNotFound{}
}

/*
PutFleetsFleetIDWingsWingIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type PutFleetsFleetIDWingsWingIDNotFound struct {
	Payload *PutFleetsFleetIDWingsWingIDNotFoundBody
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id not found response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id not found response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id not found response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id wings wing Id not found response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id wings wing Id not found response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutFleetsFleetIDWingsWingIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDNotFound) GetPayload() *PutFleetsFleetIDWingsWingIDNotFoundBody {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutFleetsFleetIDWingsWingIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDEnhanceYourCalm creates a PutFleetsFleetIDWingsWingIDEnhanceYourCalm with default headers values
func NewPutFleetsFleetIDWingsWingIDEnhanceYourCalm() *PutFleetsFleetIDWingsWingIDEnhanceYourCalm {
	return &PutFleetsFleetIDWingsWingIDEnhanceYourCalm{}
}

/*
PutFleetsFleetIDWingsWingIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PutFleetsFleetIDWingsWingIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id enhance your calm response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id enhance your calm response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id enhance your calm response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id wings wing Id enhance your calm response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id wings wing Id enhance your calm response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDInternalServerError creates a PutFleetsFleetIDWingsWingIDInternalServerError with default headers values
func NewPutFleetsFleetIDWingsWingIDInternalServerError() *PutFleetsFleetIDWingsWingIDInternalServerError {
	return &PutFleetsFleetIDWingsWingIDInternalServerError{}
}

/*
PutFleetsFleetIDWingsWingIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PutFleetsFleetIDWingsWingIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id internal server error response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id internal server error response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id internal server error response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id wings wing Id internal server error response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put fleets fleet Id wings wing Id internal server error response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutFleetsFleetIDWingsWingIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDServiceUnavailable creates a PutFleetsFleetIDWingsWingIDServiceUnavailable with default headers values
func NewPutFleetsFleetIDWingsWingIDServiceUnavailable() *PutFleetsFleetIDWingsWingIDServiceUnavailable {
	return &PutFleetsFleetIDWingsWingIDServiceUnavailable{}
}

/*
PutFleetsFleetIDWingsWingIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PutFleetsFleetIDWingsWingIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id service unavailable response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id service unavailable response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id service unavailable response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id wings wing Id service unavailable response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put fleets fleet Id wings wing Id service unavailable response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDWingsWingIDGatewayTimeout creates a PutFleetsFleetIDWingsWingIDGatewayTimeout with default headers values
func NewPutFleetsFleetIDWingsWingIDGatewayTimeout() *PutFleetsFleetIDWingsWingIDGatewayTimeout {
	return &PutFleetsFleetIDWingsWingIDGatewayTimeout{}
}

/*
PutFleetsFleetIDWingsWingIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PutFleetsFleetIDWingsWingIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this put fleets fleet Id wings wing Id gateway timeout response has a 2xx status code
func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id wings wing Id gateway timeout response has a 3xx status code
func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id wings wing Id gateway timeout response has a 4xx status code
func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id wings wing Id gateway timeout response has a 5xx status code
func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put fleets fleet Id wings wing Id gateway timeout response a status code equal to that given
func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] putFleetsFleetIdWingsWingIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PutFleetsFleetIDWingsWingIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutFleetsFleetIDWingsWingIDBody put_fleets_fleet_id_wings_wing_id_naming
//
// naming object
swagger:model PutFleetsFleetIDWingsWingIDBody
*/
type PutFleetsFleetIDWingsWingIDBody struct {

	// put_fleets_fleet_id_wings_wing_id_name
	//
	// name string
	// Required: true
	// Max Length: 10
	Name *string `json:"name"`
}

// Validate validates this put fleets fleet ID wings wing ID body
func (o *PutFleetsFleetIDWingsWingIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutFleetsFleetIDWingsWingIDBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("naming"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("naming"+"."+"name", "body", *o.Name, 10); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put fleets fleet ID wings wing ID body based on context it is used
func (o *PutFleetsFleetIDWingsWingIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDWingsWingIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutFleetsFleetIDWingsWingIDNotFoundBody put_fleets_fleet_id_wings_wing_id_not_found
//
// Not found
swagger:model PutFleetsFleetIDWingsWingIDNotFoundBody
*/
type PutFleetsFleetIDWingsWingIDNotFoundBody struct {

	// put_fleets_fleet_id_wings_wing_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this put fleets fleet ID wings wing ID not found body
func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put fleets fleet ID wings wing ID not found body based on context it is used
func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutFleetsFleetIDWingsWingIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDWingsWingIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
