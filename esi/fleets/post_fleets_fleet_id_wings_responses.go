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

// PostFleetsFleetIDWingsReader is a Reader for the PostFleetsFleetIDWings structure.
type PostFleetsFleetIDWingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetsFleetIDWingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostFleetsFleetIDWingsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostFleetsFleetIDWingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostFleetsFleetIDWingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostFleetsFleetIDWingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostFleetsFleetIDWingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostFleetsFleetIDWingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostFleetsFleetIDWingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostFleetsFleetIDWingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostFleetsFleetIDWingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/fleets/{fleet_id}/wings/] post_fleets_fleet_id_wings", response, response.Code())
	}
}

// NewPostFleetsFleetIDWingsCreated creates a PostFleetsFleetIDWingsCreated with default headers values
func NewPostFleetsFleetIDWingsCreated() *PostFleetsFleetIDWingsCreated {
	return &PostFleetsFleetIDWingsCreated{}
}

/*
PostFleetsFleetIDWingsCreated describes a response with status code 201, with default header values.

Wing created
*/
type PostFleetsFleetIDWingsCreated struct {
	Payload *PostFleetsFleetIDWingsCreatedBody
}

// IsSuccess returns true when this post fleets fleet Id wings created response has a 2xx status code
func (o *PostFleetsFleetIDWingsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post fleets fleet Id wings created response has a 3xx status code
func (o *PostFleetsFleetIDWingsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings created response has a 4xx status code
func (o *PostFleetsFleetIDWingsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post fleets fleet Id wings created response has a 5xx status code
func (o *PostFleetsFleetIDWingsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post fleets fleet Id wings created response a status code equal to that given
func (o *PostFleetsFleetIDWingsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post fleets fleet Id wings created response
func (o *PostFleetsFleetIDWingsCreated) Code() int {
	return 201
}

func (o *PostFleetsFleetIDWingsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsCreated  %+v", 201, o.Payload)
}

func (o *PostFleetsFleetIDWingsCreated) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsCreated  %+v", 201, o.Payload)
}

func (o *PostFleetsFleetIDWingsCreated) GetPayload() *PostFleetsFleetIDWingsCreatedBody {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostFleetsFleetIDWingsCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsBadRequest creates a PostFleetsFleetIDWingsBadRequest with default headers values
func NewPostFleetsFleetIDWingsBadRequest() *PostFleetsFleetIDWingsBadRequest {
	return &PostFleetsFleetIDWingsBadRequest{}
}

/*
PostFleetsFleetIDWingsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostFleetsFleetIDWingsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this post fleets fleet Id wings bad request response has a 2xx status code
func (o *PostFleetsFleetIDWingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings bad request response has a 3xx status code
func (o *PostFleetsFleetIDWingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings bad request response has a 4xx status code
func (o *PostFleetsFleetIDWingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post fleets fleet Id wings bad request response has a 5xx status code
func (o *PostFleetsFleetIDWingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post fleets fleet Id wings bad request response a status code equal to that given
func (o *PostFleetsFleetIDWingsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post fleets fleet Id wings bad request response
func (o *PostFleetsFleetIDWingsBadRequest) Code() int {
	return 400
}

func (o *PostFleetsFleetIDWingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostFleetsFleetIDWingsBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostFleetsFleetIDWingsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsUnauthorized creates a PostFleetsFleetIDWingsUnauthorized with default headers values
func NewPostFleetsFleetIDWingsUnauthorized() *PostFleetsFleetIDWingsUnauthorized {
	return &PostFleetsFleetIDWingsUnauthorized{}
}

/*
PostFleetsFleetIDWingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostFleetsFleetIDWingsUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this post fleets fleet Id wings unauthorized response has a 2xx status code
func (o *PostFleetsFleetIDWingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings unauthorized response has a 3xx status code
func (o *PostFleetsFleetIDWingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings unauthorized response has a 4xx status code
func (o *PostFleetsFleetIDWingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post fleets fleet Id wings unauthorized response has a 5xx status code
func (o *PostFleetsFleetIDWingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post fleets fleet Id wings unauthorized response a status code equal to that given
func (o *PostFleetsFleetIDWingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post fleets fleet Id wings unauthorized response
func (o *PostFleetsFleetIDWingsUnauthorized) Code() int {
	return 401
}

func (o *PostFleetsFleetIDWingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFleetsFleetIDWingsUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFleetsFleetIDWingsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsForbidden creates a PostFleetsFleetIDWingsForbidden with default headers values
func NewPostFleetsFleetIDWingsForbidden() *PostFleetsFleetIDWingsForbidden {
	return &PostFleetsFleetIDWingsForbidden{}
}

/*
PostFleetsFleetIDWingsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostFleetsFleetIDWingsForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this post fleets fleet Id wings forbidden response has a 2xx status code
func (o *PostFleetsFleetIDWingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings forbidden response has a 3xx status code
func (o *PostFleetsFleetIDWingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings forbidden response has a 4xx status code
func (o *PostFleetsFleetIDWingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post fleets fleet Id wings forbidden response has a 5xx status code
func (o *PostFleetsFleetIDWingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post fleets fleet Id wings forbidden response a status code equal to that given
func (o *PostFleetsFleetIDWingsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post fleets fleet Id wings forbidden response
func (o *PostFleetsFleetIDWingsForbidden) Code() int {
	return 403
}

func (o *PostFleetsFleetIDWingsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsForbidden  %+v", 403, o.Payload)
}

func (o *PostFleetsFleetIDWingsForbidden) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsForbidden  %+v", 403, o.Payload)
}

func (o *PostFleetsFleetIDWingsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsNotFound creates a PostFleetsFleetIDWingsNotFound with default headers values
func NewPostFleetsFleetIDWingsNotFound() *PostFleetsFleetIDWingsNotFound {
	return &PostFleetsFleetIDWingsNotFound{}
}

/*
PostFleetsFleetIDWingsNotFound describes a response with status code 404, with default header values.

The fleet does not exist or you don't have access to it
*/
type PostFleetsFleetIDWingsNotFound struct {
	Payload *PostFleetsFleetIDWingsNotFoundBody
}

// IsSuccess returns true when this post fleets fleet Id wings not found response has a 2xx status code
func (o *PostFleetsFleetIDWingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings not found response has a 3xx status code
func (o *PostFleetsFleetIDWingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings not found response has a 4xx status code
func (o *PostFleetsFleetIDWingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post fleets fleet Id wings not found response has a 5xx status code
func (o *PostFleetsFleetIDWingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post fleets fleet Id wings not found response a status code equal to that given
func (o *PostFleetsFleetIDWingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post fleets fleet Id wings not found response
func (o *PostFleetsFleetIDWingsNotFound) Code() int {
	return 404
}

func (o *PostFleetsFleetIDWingsNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsNotFound  %+v", 404, o.Payload)
}

func (o *PostFleetsFleetIDWingsNotFound) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsNotFound  %+v", 404, o.Payload)
}

func (o *PostFleetsFleetIDWingsNotFound) GetPayload() *PostFleetsFleetIDWingsNotFoundBody {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostFleetsFleetIDWingsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsEnhanceYourCalm creates a PostFleetsFleetIDWingsEnhanceYourCalm with default headers values
func NewPostFleetsFleetIDWingsEnhanceYourCalm() *PostFleetsFleetIDWingsEnhanceYourCalm {
	return &PostFleetsFleetIDWingsEnhanceYourCalm{}
}

/*
PostFleetsFleetIDWingsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostFleetsFleetIDWingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this post fleets fleet Id wings enhance your calm response has a 2xx status code
func (o *PostFleetsFleetIDWingsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings enhance your calm response has a 3xx status code
func (o *PostFleetsFleetIDWingsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings enhance your calm response has a 4xx status code
func (o *PostFleetsFleetIDWingsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this post fleets fleet Id wings enhance your calm response has a 5xx status code
func (o *PostFleetsFleetIDWingsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this post fleets fleet Id wings enhance your calm response a status code equal to that given
func (o *PostFleetsFleetIDWingsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the post fleets fleet Id wings enhance your calm response
func (o *PostFleetsFleetIDWingsEnhanceYourCalm) Code() int {
	return 420
}

func (o *PostFleetsFleetIDWingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostFleetsFleetIDWingsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostFleetsFleetIDWingsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsInternalServerError creates a PostFleetsFleetIDWingsInternalServerError with default headers values
func NewPostFleetsFleetIDWingsInternalServerError() *PostFleetsFleetIDWingsInternalServerError {
	return &PostFleetsFleetIDWingsInternalServerError{}
}

/*
PostFleetsFleetIDWingsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostFleetsFleetIDWingsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this post fleets fleet Id wings internal server error response has a 2xx status code
func (o *PostFleetsFleetIDWingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings internal server error response has a 3xx status code
func (o *PostFleetsFleetIDWingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings internal server error response has a 4xx status code
func (o *PostFleetsFleetIDWingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post fleets fleet Id wings internal server error response has a 5xx status code
func (o *PostFleetsFleetIDWingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post fleets fleet Id wings internal server error response a status code equal to that given
func (o *PostFleetsFleetIDWingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post fleets fleet Id wings internal server error response
func (o *PostFleetsFleetIDWingsInternalServerError) Code() int {
	return 500
}

func (o *PostFleetsFleetIDWingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFleetsFleetIDWingsInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFleetsFleetIDWingsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsServiceUnavailable creates a PostFleetsFleetIDWingsServiceUnavailable with default headers values
func NewPostFleetsFleetIDWingsServiceUnavailable() *PostFleetsFleetIDWingsServiceUnavailable {
	return &PostFleetsFleetIDWingsServiceUnavailable{}
}

/*
PostFleetsFleetIDWingsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostFleetsFleetIDWingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this post fleets fleet Id wings service unavailable response has a 2xx status code
func (o *PostFleetsFleetIDWingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings service unavailable response has a 3xx status code
func (o *PostFleetsFleetIDWingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings service unavailable response has a 4xx status code
func (o *PostFleetsFleetIDWingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post fleets fleet Id wings service unavailable response has a 5xx status code
func (o *PostFleetsFleetIDWingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post fleets fleet Id wings service unavailable response a status code equal to that given
func (o *PostFleetsFleetIDWingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the post fleets fleet Id wings service unavailable response
func (o *PostFleetsFleetIDWingsServiceUnavailable) Code() int {
	return 503
}

func (o *PostFleetsFleetIDWingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFleetsFleetIDWingsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFleetsFleetIDWingsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsGatewayTimeout creates a PostFleetsFleetIDWingsGatewayTimeout with default headers values
func NewPostFleetsFleetIDWingsGatewayTimeout() *PostFleetsFleetIDWingsGatewayTimeout {
	return &PostFleetsFleetIDWingsGatewayTimeout{}
}

/*
PostFleetsFleetIDWingsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostFleetsFleetIDWingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this post fleets fleet Id wings gateway timeout response has a 2xx status code
func (o *PostFleetsFleetIDWingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post fleets fleet Id wings gateway timeout response has a 3xx status code
func (o *PostFleetsFleetIDWingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post fleets fleet Id wings gateway timeout response has a 4xx status code
func (o *PostFleetsFleetIDWingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post fleets fleet Id wings gateway timeout response has a 5xx status code
func (o *PostFleetsFleetIDWingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post fleets fleet Id wings gateway timeout response a status code equal to that given
func (o *PostFleetsFleetIDWingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the post fleets fleet Id wings gateway timeout response
func (o *PostFleetsFleetIDWingsGatewayTimeout) Code() int {
	return 504
}

func (o *PostFleetsFleetIDWingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFleetsFleetIDWingsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFleetsFleetIDWingsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostFleetsFleetIDWingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostFleetsFleetIDWingsCreatedBody post_fleets_fleet_id_wings_created
//
// 201 created object
swagger:model PostFleetsFleetIDWingsCreatedBody
*/
type PostFleetsFleetIDWingsCreatedBody struct {

	// post_fleets_fleet_id_wings_wing_id
	//
	// The wing_id of the newly created wing
	// Required: true
	WingID *int64 `json:"wing_id"`
}

// Validate validates this post fleets fleet ID wings created body
func (o *PostFleetsFleetIDWingsCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateWingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostFleetsFleetIDWingsCreatedBody) validateWingID(formats strfmt.Registry) error {

	if err := validate.Required("postFleetsFleetIdWingsCreated"+"."+"wing_id", "body", o.WingID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post fleets fleet ID wings created body based on context it is used
func (o *PostFleetsFleetIDWingsCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostFleetsFleetIDWingsCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFleetsFleetIDWingsCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostFleetsFleetIDWingsCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostFleetsFleetIDWingsNotFoundBody post_fleets_fleet_id_wings_not_found
//
// Not found
swagger:model PostFleetsFleetIDWingsNotFoundBody
*/
type PostFleetsFleetIDWingsNotFoundBody struct {

	// post_fleets_fleet_id_wings_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this post fleets fleet ID wings not found body
func (o *PostFleetsFleetIDWingsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post fleets fleet ID wings not found body based on context it is used
func (o *PostFleetsFleetIDWingsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostFleetsFleetIDWingsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostFleetsFleetIDWingsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostFleetsFleetIDWingsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
