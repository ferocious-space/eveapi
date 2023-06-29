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

// DeleteFleetsFleetIDMembersMemberIDReader is a Reader for the DeleteFleetsFleetIDMembersMemberID structure.
type DeleteFleetsFleetIDMembersMemberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetsFleetIDMembersMemberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFleetsFleetIDMembersMemberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFleetsFleetIDMembersMemberIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFleetsFleetIDMembersMemberIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFleetsFleetIDMembersMemberIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFleetsFleetIDMembersMemberIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewDeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFleetsFleetIDMembersMemberIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFleetsFleetIDMembersMemberIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFleetsFleetIDMembersMemberIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/] delete_fleets_fleet_id_members_member_id", response, response.Code())
	}
}

// NewDeleteFleetsFleetIDMembersMemberIDNoContent creates a DeleteFleetsFleetIDMembersMemberIDNoContent with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDNoContent() *DeleteFleetsFleetIDMembersMemberIDNoContent {
	return &DeleteFleetsFleetIDMembersMemberIDNoContent{}
}

/*
DeleteFleetsFleetIDMembersMemberIDNoContent describes a response with status code 204, with default header values.

Fleet member kicked
*/
type DeleteFleetsFleetIDMembersMemberIDNoContent struct {
}

// IsSuccess returns true when this delete fleets fleet Id members member Id no content response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete fleets fleet Id members member Id no content response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id no content response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id members member Id no content response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id members member Id no content response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete fleets fleet Id members member Id no content response
func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) Code() int {
	return 204
}

func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNoContent ", 204)
}

func (o *DeleteFleetsFleetIDMembersMemberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDBadRequest creates a DeleteFleetsFleetIDMembersMemberIDBadRequest with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDBadRequest() *DeleteFleetsFleetIDMembersMemberIDBadRequest {
	return &DeleteFleetsFleetIDMembersMemberIDBadRequest{}
}

/*
DeleteFleetsFleetIDMembersMemberIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteFleetsFleetIDMembersMemberIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this delete fleets fleet Id members member Id bad request response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id bad request response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id bad request response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id members member Id bad request response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id members member Id bad request response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete fleets fleet Id members member Id bad request response
func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) Code() int {
	return 400
}

func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDUnauthorized creates a DeleteFleetsFleetIDMembersMemberIDUnauthorized with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDUnauthorized() *DeleteFleetsFleetIDMembersMemberIDUnauthorized {
	return &DeleteFleetsFleetIDMembersMemberIDUnauthorized{}
}

/*
DeleteFleetsFleetIDMembersMemberIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteFleetsFleetIDMembersMemberIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this delete fleets fleet Id members member Id unauthorized response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id unauthorized response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id unauthorized response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id members member Id unauthorized response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id members member Id unauthorized response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete fleets fleet Id members member Id unauthorized response
func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDForbidden creates a DeleteFleetsFleetIDMembersMemberIDForbidden with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDForbidden() *DeleteFleetsFleetIDMembersMemberIDForbidden {
	return &DeleteFleetsFleetIDMembersMemberIDForbidden{}
}

/*
DeleteFleetsFleetIDMembersMemberIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteFleetsFleetIDMembersMemberIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this delete fleets fleet Id members member Id forbidden response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id forbidden response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id forbidden response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id members member Id forbidden response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id members member Id forbidden response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete fleets fleet Id members member Id forbidden response
func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) Code() int {
	return 403
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDNotFound creates a DeleteFleetsFleetIDMembersMemberIDNotFound with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDNotFound() *DeleteFleetsFleetIDMembersMemberIDNotFound {
	return &DeleteFleetsFleetIDMembersMemberIDNotFound{}
}

/*
DeleteFleetsFleetIDMembersMemberIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type DeleteFleetsFleetIDMembersMemberIDNotFound struct {
	Payload *DeleteFleetsFleetIDMembersMemberIDNotFoundBody
}

// IsSuccess returns true when this delete fleets fleet Id members member Id not found response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id not found response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id not found response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id members member Id not found response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id members member Id not found response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete fleets fleet Id members member Id not found response
func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) Code() int {
	return 404
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) GetPayload() *DeleteFleetsFleetIDMembersMemberIDNotFoundBody {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DeleteFleetsFleetIDMembersMemberIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm creates a DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm() *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm {
	return &DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm{}
}

/*
DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this delete fleets fleet Id members member Id enhance your calm response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id enhance your calm response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id enhance your calm response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete fleets fleet Id members member Id enhance your calm response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this delete fleets fleet Id members member Id enhance your calm response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the delete fleets fleet Id members member Id enhance your calm response
func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDInternalServerError creates a DeleteFleetsFleetIDMembersMemberIDInternalServerError with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDInternalServerError() *DeleteFleetsFleetIDMembersMemberIDInternalServerError {
	return &DeleteFleetsFleetIDMembersMemberIDInternalServerError{}
}

/*
DeleteFleetsFleetIDMembersMemberIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteFleetsFleetIDMembersMemberIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this delete fleets fleet Id members member Id internal server error response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id internal server error response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id internal server error response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id members member Id internal server error response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fleets fleet Id members member Id internal server error response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete fleets fleet Id members member Id internal server error response
func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDServiceUnavailable creates a DeleteFleetsFleetIDMembersMemberIDServiceUnavailable with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDServiceUnavailable() *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable {
	return &DeleteFleetsFleetIDMembersMemberIDServiceUnavailable{}
}

/*
DeleteFleetsFleetIDMembersMemberIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type DeleteFleetsFleetIDMembersMemberIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this delete fleets fleet Id members member Id service unavailable response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id service unavailable response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id service unavailable response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id members member Id service unavailable response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fleets fleet Id members member Id service unavailable response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the delete fleets fleet Id members member Id service unavailable response
func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) Code() int {
	return 503
}

func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetsFleetIDMembersMemberIDGatewayTimeout creates a DeleteFleetsFleetIDMembersMemberIDGatewayTimeout with default headers values
func NewDeleteFleetsFleetIDMembersMemberIDGatewayTimeout() *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout {
	return &DeleteFleetsFleetIDMembersMemberIDGatewayTimeout{}
}

/*
DeleteFleetsFleetIDMembersMemberIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type DeleteFleetsFleetIDMembersMemberIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this delete fleets fleet Id members member Id gateway timeout response has a 2xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete fleets fleet Id members member Id gateway timeout response has a 3xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete fleets fleet Id members member Id gateway timeout response has a 4xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete fleets fleet Id members member Id gateway timeout response has a 5xx status code
func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete fleets fleet Id members member Id gateway timeout response a status code equal to that given
func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the delete fleets fleet Id members member Id gateway timeout response
func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) Code() int {
	return 504
}

func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/fleets/{fleet_id}/members/{member_id}/][%d] deleteFleetsFleetIdMembersMemberIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *DeleteFleetsFleetIDMembersMemberIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DeleteFleetsFleetIDMembersMemberIDNotFoundBody delete_fleets_fleet_id_members_member_id_not_found
//
// Not found
swagger:model DeleteFleetsFleetIDMembersMemberIDNotFoundBody
*/
type DeleteFleetsFleetIDMembersMemberIDNotFoundBody struct {

	// delete_fleets_fleet_id_members_member_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet ID members member ID not found body
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete fleets fleet ID members member ID not found body based on context it is used
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDMembersMemberIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
