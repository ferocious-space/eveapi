// Code generated by go-swagger; DO NOT EDIT.

package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ferocious-space/eveapi/models"
)

// DeleteFleetsFleetIDWingsWingIDReader is a Reader for the DeleteFleetsFleetIDWingsWingID structure.
type DeleteFleetsFleetIDWingsWingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetsFleetIDWingsWingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFleetsFleetIDWingsWingIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFleetsFleetIDWingsWingIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFleetsFleetIDWingsWingIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFleetsFleetIDWingsWingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFleetsFleetIDWingsWingIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewDeleteFleetsFleetIDWingsWingIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFleetsFleetIDWingsWingIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFleetsFleetIDWingsWingIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFleetsFleetIDWingsWingIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFleetsFleetIDWingsWingIDNoContent creates a DeleteFleetsFleetIDWingsWingIDNoContent with default headers values
func NewDeleteFleetsFleetIDWingsWingIDNoContent() *DeleteFleetsFleetIDWingsWingIDNoContent {
	return &DeleteFleetsFleetIDWingsWingIDNoContent{}
}

/*
DeleteFleetsFleetIDWingsWingIDNoContent describes a response with status code 204, with default header values.

Wing deleted
*/
type DeleteFleetsFleetIDWingsWingIDNoContent struct {
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id no content response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id no content response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id no content response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id no content response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id wings wing Id no content response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete fleets fleet Id wings wing Id no content response
func (o *DeleteFleetsFleetIDWingsWingIDNoContent) Code() int {
	return 204
}

func (o *DeleteFleetsFleetIDWingsWingIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDWingsWingIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDWingsWingIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDBadRequest creates a DeleteFleetsFleetIDWingsWingIDBadRequest with default headers values
func NewDeleteFleetsFleetIDWingsWingIDBadRequest() *DeleteFleetsFleetIDWingsWingIDBadRequest {
	return &DeleteFleetsFleetIDWingsWingIDBadRequest{}
}

/*
DeleteFleetsFleetIDWingsWingIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteFleetsFleetIDWingsWingIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id bad request response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id bad request response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id bad request response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id bad request response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id wings wing Id bad request response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete fleets fleet Id wings wing Id bad request response
func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) Code() int {
	return 400
}

func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDUnauthorized creates a DeleteFleetsFleetIDWingsWingIDUnauthorized with default headers values
func NewDeleteFleetsFleetIDWingsWingIDUnauthorized() *DeleteFleetsFleetIDWingsWingIDUnauthorized {
	return &DeleteFleetsFleetIDWingsWingIDUnauthorized{}
}

/*
DeleteFleetsFleetIDWingsWingIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteFleetsFleetIDWingsWingIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id unauthorized response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id unauthorized response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id unauthorized response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id unauthorized response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id wings wing Id unauthorized response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete fleets fleet Id wings wing Id unauthorized response
func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDForbidden creates a DeleteFleetsFleetIDWingsWingIDForbidden with default headers values
func NewDeleteFleetsFleetIDWingsWingIDForbidden() *DeleteFleetsFleetIDWingsWingIDForbidden {
	return &DeleteFleetsFleetIDWingsWingIDForbidden{}
}

/*
DeleteFleetsFleetIDWingsWingIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteFleetsFleetIDWingsWingIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id forbidden response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id forbidden response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id forbidden response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id forbidden response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id wings wing Id forbidden response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete fleets fleet Id wings wing Id forbidden response
func (o *DeleteFleetsFleetIDWingsWingIDForbidden) Code() int {
	return 403
}

func (o *DeleteFleetsFleetIDWingsWingIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDNotFound creates a DeleteFleetsFleetIDWingsWingIDNotFound with default headers values
func NewDeleteFleetsFleetIDWingsWingIDNotFound() *DeleteFleetsFleetIDWingsWingIDNotFound {
	return &DeleteFleetsFleetIDWingsWingIDNotFound{}
}

/*
DeleteFleetsFleetIDWingsWingIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type DeleteFleetsFleetIDWingsWingIDNotFound struct {
	Payload *DeleteFleetsFleetIDWingsWingIDNotFoundBody
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id not found response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id not found response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id not found response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id not found response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id wings wing Id not found response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete fleets fleet Id wings wing Id not found response
func (o *DeleteFleetsFleetIDWingsWingIDNotFound) Code() int {
	return 404
}

func (o *DeleteFleetsFleetIDWingsWingIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDNotFound) GetPayload() *DeleteFleetsFleetIDWingsWingIDNotFoundBody {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteFleetsFleetIDWingsWingIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDEnhanceYourCalm creates a DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm with default headers values
func NewDeleteFleetsFleetIDWingsWingIDEnhanceYourCalm() *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm {
	return &DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm{}
}

/*
DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id enhance your calm response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id enhance your calm response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id enhance your calm response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id enhance your calm response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id wings wing Id enhance your calm response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the delete fleets fleet Id wings wing Id enhance your calm response
func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDInternalServerError creates a DeleteFleetsFleetIDWingsWingIDInternalServerError with default headers values
func NewDeleteFleetsFleetIDWingsWingIDInternalServerError() *DeleteFleetsFleetIDWingsWingIDInternalServerError {
	return &DeleteFleetsFleetIDWingsWingIDInternalServerError{}
}

/*
DeleteFleetsFleetIDWingsWingIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteFleetsFleetIDWingsWingIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id internal server error response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id internal server error response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id internal server error response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id internal server error response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fleets fleet Id wings wing Id internal server error response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete fleets fleet Id wings wing Id internal server error response
func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDServiceUnavailable creates a DeleteFleetsFleetIDWingsWingIDServiceUnavailable with default headers values
func NewDeleteFleetsFleetIDWingsWingIDServiceUnavailable() *DeleteFleetsFleetIDWingsWingIDServiceUnavailable {
	return &DeleteFleetsFleetIDWingsWingIDServiceUnavailable{}
}

/*
DeleteFleetsFleetIDWingsWingIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type DeleteFleetsFleetIDWingsWingIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id service unavailable response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id service unavailable response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id service unavailable response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id service unavailable response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fleets fleet Id wings wing Id service unavailable response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the delete fleets fleet Id wings wing Id service unavailable response
func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) Code() int {
	return 503
}

func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDWingsWingIDGatewayTimeout creates a DeleteFleetsFleetIDWingsWingIDGatewayTimeout with default headers values
func NewDeleteFleetsFleetIDWingsWingIDGatewayTimeout() *DeleteFleetsFleetIDWingsWingIDGatewayTimeout {
	return &DeleteFleetsFleetIDWingsWingIDGatewayTimeout{}
}

/*
DeleteFleetsFleetIDWingsWingIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type DeleteFleetsFleetIDWingsWingIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this delete fleets fleet Id wings wing Id gateway timeout response has a 2xx status code
func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id wings wing Id gateway timeout response has a 3xx status code
func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id wings wing Id gateway timeout response has a 4xx status code
func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id wings wing Id gateway timeout response has a 5xx status code
func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fleets fleet Id wings wing Id gateway timeout response a status code equal to that given
func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the delete fleets fleet Id wings wing Id gateway timeout response
func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) Code() int {
	return 504
}

func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/wings/{wing_id}/][%d] deleteFleetsFleetIdWingsWingIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *DeleteFleetsFleetIDWingsWingIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DeleteFleetsFleetIDWingsWingIDNotFoundBody delete_fleets_fleet_id_wings_wing_id_not_found
//
// Not found
swagger:model DeleteFleetsFleetIDWingsWingIDNotFoundBody
*/
type DeleteFleetsFleetIDWingsWingIDNotFoundBody struct {

	// delete_fleets_fleet_id_wings_wing_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet ID wings wing ID not found body
func (o *DeleteFleetsFleetIDWingsWingIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete fleets fleet ID wings wing ID not found body based on context it is used
func (o *DeleteFleetsFleetIDWingsWingIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteFleetsFleetIDWingsWingIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteFleetsFleetIDWingsWingIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDWingsWingIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
