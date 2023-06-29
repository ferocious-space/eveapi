// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// PutFleetsFleetIDMembersMemberIDReader is a Reader for the PutFleetsFleetIDMembersMemberID structure.
type PutFleetsFleetIDMembersMemberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetsFleetIDMembersMemberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutFleetsFleetIDMembersMemberIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutFleetsFleetIDMembersMemberIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutFleetsFleetIDMembersMemberIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutFleetsFleetIDMembersMemberIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutFleetsFleetIDMembersMemberIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPutFleetsFleetIDMembersMemberIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutFleetsFleetIDMembersMemberIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutFleetsFleetIDMembersMemberIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutFleetsFleetIDMembersMemberIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutFleetsFleetIDMembersMemberIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/fleets/{fleet_id}/members/{member_id}/] put_fleets_fleet_id_members_member_id", response, response.Code())
	}
}

// NewPutFleetsFleetIDMembersMemberIDNoContent creates a PutFleetsFleetIDMembersMemberIDNoContent with default headers values
func NewPutFleetsFleetIDMembersMemberIDNoContent() *PutFleetsFleetIDMembersMemberIDNoContent {
	return &PutFleetsFleetIDMembersMemberIDNoContent{}
}

/*
PutFleetsFleetIDMembersMemberIDNoContent describes a response with status code 204, with default header values.

Fleet invitation sent
*/
type PutFleetsFleetIDMembersMemberIDNoContent struct {
}

// IsSuccess returns true when this put fleets fleet Id members member Id no content response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put fleets fleet Id members member Id no content response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id no content response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id members member Id no content response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id members member Id no content response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the put fleets fleet Id members member Id no content response
func (o *PutFleetsFleetIDMembersMemberIDNoContent) Code() int {
	return 204
}

func (o *PutFleetsFleetIDMembersMemberIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdNoContent ", 204)
}

func (o *PutFleetsFleetIDMembersMemberIDNoContent) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdNoContent ", 204)
}

func (o *PutFleetsFleetIDMembersMemberIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDBadRequest creates a PutFleetsFleetIDMembersMemberIDBadRequest with default headers values
func NewPutFleetsFleetIDMembersMemberIDBadRequest() *PutFleetsFleetIDMembersMemberIDBadRequest {
	return &PutFleetsFleetIDMembersMemberIDBadRequest{}
}

/*
PutFleetsFleetIDMembersMemberIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutFleetsFleetIDMembersMemberIDBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this put fleets fleet Id members member Id bad request response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id bad request response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id bad request response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id members member Id bad request response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id members member Id bad request response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put fleets fleet Id members member Id bad request response
func (o *PutFleetsFleetIDMembersMemberIDBadRequest) Code() int {
	return 400
}

func (o *PutFleetsFleetIDMembersMemberIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDUnauthorized creates a PutFleetsFleetIDMembersMemberIDUnauthorized with default headers values
func NewPutFleetsFleetIDMembersMemberIDUnauthorized() *PutFleetsFleetIDMembersMemberIDUnauthorized {
	return &PutFleetsFleetIDMembersMemberIDUnauthorized{}
}

/*
PutFleetsFleetIDMembersMemberIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutFleetsFleetIDMembersMemberIDUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this put fleets fleet Id members member Id unauthorized response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id unauthorized response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id unauthorized response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id members member Id unauthorized response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id members member Id unauthorized response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put fleets fleet Id members member Id unauthorized response
func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) Code() int {
	return 401
}

func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDForbidden creates a PutFleetsFleetIDMembersMemberIDForbidden with default headers values
func NewPutFleetsFleetIDMembersMemberIDForbidden() *PutFleetsFleetIDMembersMemberIDForbidden {
	return &PutFleetsFleetIDMembersMemberIDForbidden{}
}

/*
PutFleetsFleetIDMembersMemberIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutFleetsFleetIDMembersMemberIDForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this put fleets fleet Id members member Id forbidden response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id forbidden response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id forbidden response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id members member Id forbidden response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id members member Id forbidden response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put fleets fleet Id members member Id forbidden response
func (o *PutFleetsFleetIDMembersMemberIDForbidden) Code() int {
	return 403
}

func (o *PutFleetsFleetIDMembersMemberIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDNotFound creates a PutFleetsFleetIDMembersMemberIDNotFound with default headers values
func NewPutFleetsFleetIDMembersMemberIDNotFound() *PutFleetsFleetIDMembersMemberIDNotFound {
	return &PutFleetsFleetIDMembersMemberIDNotFound{}
}

/*
PutFleetsFleetIDMembersMemberIDNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type PutFleetsFleetIDMembersMemberIDNotFound struct {
	Payload *PutFleetsFleetIDMembersMemberIDNotFoundBody
}

// IsSuccess returns true when this put fleets fleet Id members member Id not found response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id not found response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id not found response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id members member Id not found response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id members member Id not found response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put fleets fleet Id members member Id not found response
func (o *PutFleetsFleetIDMembersMemberIDNotFound) Code() int {
	return 404
}

func (o *PutFleetsFleetIDMembersMemberIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDNotFound) GetPayload() *PutFleetsFleetIDMembersMemberIDNotFoundBody {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutFleetsFleetIDMembersMemberIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDEnhanceYourCalm creates a PutFleetsFleetIDMembersMemberIDEnhanceYourCalm with default headers values
func NewPutFleetsFleetIDMembersMemberIDEnhanceYourCalm() *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm {
	return &PutFleetsFleetIDMembersMemberIDEnhanceYourCalm{}
}

/*
PutFleetsFleetIDMembersMemberIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PutFleetsFleetIDMembersMemberIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this put fleets fleet Id members member Id enhance your calm response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id enhance your calm response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id enhance your calm response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id members member Id enhance your calm response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id members member Id enhance your calm response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the put fleets fleet Id members member Id enhance your calm response
func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) Code() int {
	return 420
}

func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDUnprocessableEntity creates a PutFleetsFleetIDMembersMemberIDUnprocessableEntity with default headers values
func NewPutFleetsFleetIDMembersMemberIDUnprocessableEntity() *PutFleetsFleetIDMembersMemberIDUnprocessableEntity {
	return &PutFleetsFleetIDMembersMemberIDUnprocessableEntity{}
}

/*
PutFleetsFleetIDMembersMemberIDUnprocessableEntity describes a response with status code 422, with default header values.

Errors in invitation
*/
type PutFleetsFleetIDMembersMemberIDUnprocessableEntity struct {
	Payload *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody
}

// IsSuccess returns true when this put fleets fleet Id members member Id unprocessable entity response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id unprocessable entity response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id unprocessable entity response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this put fleets fleet Id members member Id unprocessable entity response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this put fleets fleet Id members member Id unprocessable entity response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the put fleets fleet Id members member Id unprocessable entity response
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) Code() int {
	return 422
}

func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) GetPayload() *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDInternalServerError creates a PutFleetsFleetIDMembersMemberIDInternalServerError with default headers values
func NewPutFleetsFleetIDMembersMemberIDInternalServerError() *PutFleetsFleetIDMembersMemberIDInternalServerError {
	return &PutFleetsFleetIDMembersMemberIDInternalServerError{}
}

/*
PutFleetsFleetIDMembersMemberIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PutFleetsFleetIDMembersMemberIDInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this put fleets fleet Id members member Id internal server error response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id internal server error response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id internal server error response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id members member Id internal server error response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put fleets fleet Id members member Id internal server error response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put fleets fleet Id members member Id internal server error response
func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) Code() int {
	return 500
}

func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDServiceUnavailable creates a PutFleetsFleetIDMembersMemberIDServiceUnavailable with default headers values
func NewPutFleetsFleetIDMembersMemberIDServiceUnavailable() *PutFleetsFleetIDMembersMemberIDServiceUnavailable {
	return &PutFleetsFleetIDMembersMemberIDServiceUnavailable{}
}

/*
PutFleetsFleetIDMembersMemberIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PutFleetsFleetIDMembersMemberIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this put fleets fleet Id members member Id service unavailable response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id service unavailable response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id service unavailable response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id members member Id service unavailable response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put fleets fleet Id members member Id service unavailable response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the put fleets fleet Id members member Id service unavailable response
func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) Code() int {
	return 503
}

func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDMembersMemberIDGatewayTimeout creates a PutFleetsFleetIDMembersMemberIDGatewayTimeout with default headers values
func NewPutFleetsFleetIDMembersMemberIDGatewayTimeout() *PutFleetsFleetIDMembersMemberIDGatewayTimeout {
	return &PutFleetsFleetIDMembersMemberIDGatewayTimeout{}
}

/*
PutFleetsFleetIDMembersMemberIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PutFleetsFleetIDMembersMemberIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this put fleets fleet Id members member Id gateway timeout response has a 2xx status code
func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put fleets fleet Id members member Id gateway timeout response has a 3xx status code
func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put fleets fleet Id members member Id gateway timeout response has a 4xx status code
func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put fleets fleet Id members member Id gateway timeout response has a 5xx status code
func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put fleets fleet Id members member Id gateway timeout response a status code equal to that given
func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the put fleets fleet Id members member Id gateway timeout response
func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) Code() int {
	return 504
}

func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/fleets/{fleet_id}/members/{member_id}/][%d] putFleetsFleetIdMembersMemberIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PutFleetsFleetIDMembersMemberIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutFleetsFleetIDMembersMemberIDBody put_fleets_fleet_id_members_member_id_movement
//
// movement object
swagger:model PutFleetsFleetIDMembersMemberIDBody
*/
type PutFleetsFleetIDMembersMemberIDBody struct {

	// put_fleets_fleet_id_members_member_id_role
	//
	// If a character is moved to the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified. If a character is moved to the `wing_commander` role, only `wing_id` should be specified. If a character is moved to the `squad_commander` role, both `wing_id` and `squad_id` should be specified. If a character is moved to the `squad_member` role, both `wing_id` and `squad_id` should be specified.
	// Required: true
	// Enum: [fleet_commander wing_commander squad_commander squad_member]
	Role *string `json:"role"`

	// put_fleets_fleet_id_members_member_id_squad_id
	//
	// squad_id integer
	// Minimum: 0
	SquadID *int64 `json:"squad_id,omitempty"`

	// put_fleets_fleet_id_members_member_id_wing_id
	//
	// wing_id integer
	// Minimum: 0
	WingID *int64 `json:"wing_id,omitempty"`
}

// Validate validates this put fleets fleet ID members member ID body
func (o *PutFleetsFleetIDMembersMemberIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSquadID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var putFleetsFleetIdMembersMemberIdBodyTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fleet_commander","wing_commander","squad_commander","squad_member"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putFleetsFleetIdMembersMemberIdBodyTypeRolePropEnum = append(putFleetsFleetIdMembersMemberIdBodyTypeRolePropEnum, v)
	}
}

const (

	// PutFleetsFleetIDMembersMemberIDBodyRoleFleetCommander captures enum value "fleet_commander"
	PutFleetsFleetIDMembersMemberIDBodyRoleFleetCommander string = "fleet_commander"

	// PutFleetsFleetIDMembersMemberIDBodyRoleWingCommander captures enum value "wing_commander"
	PutFleetsFleetIDMembersMemberIDBodyRoleWingCommander string = "wing_commander"

	// PutFleetsFleetIDMembersMemberIDBodyRoleSquadCommander captures enum value "squad_commander"
	PutFleetsFleetIDMembersMemberIDBodyRoleSquadCommander string = "squad_commander"

	// PutFleetsFleetIDMembersMemberIDBodyRoleSquadMember captures enum value "squad_member"
	PutFleetsFleetIDMembersMemberIDBodyRoleSquadMember string = "squad_member"
)

// prop value enum
func (o *PutFleetsFleetIDMembersMemberIDBody) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, putFleetsFleetIdMembersMemberIdBodyTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PutFleetsFleetIDMembersMemberIDBody) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("movement"+"."+"role", "body", o.Role); err != nil {
		return err
	}

	// value enum
	if err := o.validateRoleEnum("movement"+"."+"role", "body", *o.Role); err != nil {
		return err
	}

	return nil
}

func (o *PutFleetsFleetIDMembersMemberIDBody) validateSquadID(formats strfmt.Registry) error {
	if swag.IsZero(o.SquadID) { // not required
		return nil
	}

	if err := validate.MinimumInt("movement"+"."+"squad_id", "body", *o.SquadID, 0, false); err != nil {
		return err
	}

	return nil
}

func (o *PutFleetsFleetIDMembersMemberIDBody) validateWingID(formats strfmt.Registry) error {
	if swag.IsZero(o.WingID) { // not required
		return nil
	}

	if err := validate.MinimumInt("movement"+"."+"wing_id", "body", *o.WingID, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put fleets fleet ID members member ID body based on context it is used
func (o *PutFleetsFleetIDMembersMemberIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutFleetsFleetIDMembersMemberIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutFleetsFleetIDMembersMemberIDBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDMembersMemberIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutFleetsFleetIDMembersMemberIDNotFoundBody put_fleets_fleet_id_members_member_id_not_found
//
// Not found
swagger:model PutFleetsFleetIDMembersMemberIDNotFoundBody
*/
type PutFleetsFleetIDMembersMemberIDNotFoundBody struct {

	// put_fleets_fleet_id_members_member_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this put fleets fleet ID members member ID not found body
func (o *PutFleetsFleetIDMembersMemberIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put fleets fleet ID members member ID not found body based on context it is used
func (o *PutFleetsFleetIDMembersMemberIDNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutFleetsFleetIDMembersMemberIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutFleetsFleetIDMembersMemberIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDMembersMemberIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody put_fleets_fleet_id_members_member_id_unprocessable_entity
//
// 422 unprocessable entity object
swagger:model PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody
*/
type PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody struct {

	// put_fleets_fleet_id_members_member_id_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this put fleets fleet ID members member ID unprocessable entity body
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put fleets fleet ID members member ID unprocessable entity body based on context it is used
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
