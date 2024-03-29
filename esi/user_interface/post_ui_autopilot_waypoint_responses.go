// Code generated by go-swagger; DO NOT EDIT.

package user_interface

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/models"
)

// PostUIAutopilotWaypointReader is a Reader for the PostUIAutopilotWaypoint structure.
type PostUIAutopilotWaypointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIAutopilotWaypointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUIAutopilotWaypointNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUIAutopilotWaypointBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUIAutopilotWaypointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUIAutopilotWaypointForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostUIAutopilotWaypointEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUIAutopilotWaypointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUIAutopilotWaypointServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUIAutopilotWaypointGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v2/ui/autopilot/waypoint/] post_ui_autopilot_waypoint", response, response.Code())
	}
}

// NewPostUIAutopilotWaypointNoContent creates a PostUIAutopilotWaypointNoContent with default headers values
func NewPostUIAutopilotWaypointNoContent() *PostUIAutopilotWaypointNoContent {
	return &PostUIAutopilotWaypointNoContent{}
}

/*
PostUIAutopilotWaypointNoContent describes a response with status code 204, with default header values.

Open window request received
*/
type PostUIAutopilotWaypointNoContent struct {
}

// IsSuccess returns true when this post Ui autopilot waypoint no content response has a 2xx status code
func (o *PostUIAutopilotWaypointNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Ui autopilot waypoint no content response has a 3xx status code
func (o *PostUIAutopilotWaypointNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint no content response has a 4xx status code
func (o *PostUIAutopilotWaypointNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui autopilot waypoint no content response has a 5xx status code
func (o *PostUIAutopilotWaypointNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui autopilot waypoint no content response a status code equal to that given
func (o *PostUIAutopilotWaypointNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post Ui autopilot waypoint no content response
func (o *PostUIAutopilotWaypointNoContent) Code() int {
	return 204
}

func (o *PostUIAutopilotWaypointNoContent) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointNoContent ", 204)
}

func (o *PostUIAutopilotWaypointNoContent) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointNoContent ", 204)
}

func (o *PostUIAutopilotWaypointNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIAutopilotWaypointBadRequest creates a PostUIAutopilotWaypointBadRequest with default headers values
func NewPostUIAutopilotWaypointBadRequest() *PostUIAutopilotWaypointBadRequest {
	return &PostUIAutopilotWaypointBadRequest{}
}

/*
PostUIAutopilotWaypointBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostUIAutopilotWaypointBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this post Ui autopilot waypoint bad request response has a 2xx status code
func (o *PostUIAutopilotWaypointBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui autopilot waypoint bad request response has a 3xx status code
func (o *PostUIAutopilotWaypointBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint bad request response has a 4xx status code
func (o *PostUIAutopilotWaypointBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui autopilot waypoint bad request response has a 5xx status code
func (o *PostUIAutopilotWaypointBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui autopilot waypoint bad request response a status code equal to that given
func (o *PostUIAutopilotWaypointBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post Ui autopilot waypoint bad request response
func (o *PostUIAutopilotWaypointBadRequest) Code() int {
	return 400
}

func (o *PostUIAutopilotWaypointBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIAutopilotWaypointBadRequest) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIAutopilotWaypointBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostUIAutopilotWaypointBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIAutopilotWaypointUnauthorized creates a PostUIAutopilotWaypointUnauthorized with default headers values
func NewPostUIAutopilotWaypointUnauthorized() *PostUIAutopilotWaypointUnauthorized {
	return &PostUIAutopilotWaypointUnauthorized{}
}

/*
PostUIAutopilotWaypointUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUIAutopilotWaypointUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this post Ui autopilot waypoint unauthorized response has a 2xx status code
func (o *PostUIAutopilotWaypointUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui autopilot waypoint unauthorized response has a 3xx status code
func (o *PostUIAutopilotWaypointUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint unauthorized response has a 4xx status code
func (o *PostUIAutopilotWaypointUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui autopilot waypoint unauthorized response has a 5xx status code
func (o *PostUIAutopilotWaypointUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui autopilot waypoint unauthorized response a status code equal to that given
func (o *PostUIAutopilotWaypointUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post Ui autopilot waypoint unauthorized response
func (o *PostUIAutopilotWaypointUnauthorized) Code() int {
	return 401
}

func (o *PostUIAutopilotWaypointUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIAutopilotWaypointUnauthorized) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIAutopilotWaypointUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostUIAutopilotWaypointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIAutopilotWaypointForbidden creates a PostUIAutopilotWaypointForbidden with default headers values
func NewPostUIAutopilotWaypointForbidden() *PostUIAutopilotWaypointForbidden {
	return &PostUIAutopilotWaypointForbidden{}
}

/*
PostUIAutopilotWaypointForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUIAutopilotWaypointForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this post Ui autopilot waypoint forbidden response has a 2xx status code
func (o *PostUIAutopilotWaypointForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui autopilot waypoint forbidden response has a 3xx status code
func (o *PostUIAutopilotWaypointForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint forbidden response has a 4xx status code
func (o *PostUIAutopilotWaypointForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui autopilot waypoint forbidden response has a 5xx status code
func (o *PostUIAutopilotWaypointForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui autopilot waypoint forbidden response a status code equal to that given
func (o *PostUIAutopilotWaypointForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post Ui autopilot waypoint forbidden response
func (o *PostUIAutopilotWaypointForbidden) Code() int {
	return 403
}

func (o *PostUIAutopilotWaypointForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointForbidden  %+v", 403, o.Payload)
}

func (o *PostUIAutopilotWaypointForbidden) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointForbidden  %+v", 403, o.Payload)
}

func (o *PostUIAutopilotWaypointForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostUIAutopilotWaypointForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIAutopilotWaypointEnhanceYourCalm creates a PostUIAutopilotWaypointEnhanceYourCalm with default headers values
func NewPostUIAutopilotWaypointEnhanceYourCalm() *PostUIAutopilotWaypointEnhanceYourCalm {
	return &PostUIAutopilotWaypointEnhanceYourCalm{}
}

/*
PostUIAutopilotWaypointEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostUIAutopilotWaypointEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this post Ui autopilot waypoint enhance your calm response has a 2xx status code
func (o *PostUIAutopilotWaypointEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui autopilot waypoint enhance your calm response has a 3xx status code
func (o *PostUIAutopilotWaypointEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint enhance your calm response has a 4xx status code
func (o *PostUIAutopilotWaypointEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui autopilot waypoint enhance your calm response has a 5xx status code
func (o *PostUIAutopilotWaypointEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui autopilot waypoint enhance your calm response a status code equal to that given
func (o *PostUIAutopilotWaypointEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the post Ui autopilot waypoint enhance your calm response
func (o *PostUIAutopilotWaypointEnhanceYourCalm) Code() int {
	return 420
}

func (o *PostUIAutopilotWaypointEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIAutopilotWaypointEnhanceYourCalm) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIAutopilotWaypointEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostUIAutopilotWaypointEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIAutopilotWaypointInternalServerError creates a PostUIAutopilotWaypointInternalServerError with default headers values
func NewPostUIAutopilotWaypointInternalServerError() *PostUIAutopilotWaypointInternalServerError {
	return &PostUIAutopilotWaypointInternalServerError{}
}

/*
PostUIAutopilotWaypointInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostUIAutopilotWaypointInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this post Ui autopilot waypoint internal server error response has a 2xx status code
func (o *PostUIAutopilotWaypointInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui autopilot waypoint internal server error response has a 3xx status code
func (o *PostUIAutopilotWaypointInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint internal server error response has a 4xx status code
func (o *PostUIAutopilotWaypointInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui autopilot waypoint internal server error response has a 5xx status code
func (o *PostUIAutopilotWaypointInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui autopilot waypoint internal server error response a status code equal to that given
func (o *PostUIAutopilotWaypointInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post Ui autopilot waypoint internal server error response
func (o *PostUIAutopilotWaypointInternalServerError) Code() int {
	return 500
}

func (o *PostUIAutopilotWaypointInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIAutopilotWaypointInternalServerError) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIAutopilotWaypointInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostUIAutopilotWaypointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIAutopilotWaypointServiceUnavailable creates a PostUIAutopilotWaypointServiceUnavailable with default headers values
func NewPostUIAutopilotWaypointServiceUnavailable() *PostUIAutopilotWaypointServiceUnavailable {
	return &PostUIAutopilotWaypointServiceUnavailable{}
}

/*
PostUIAutopilotWaypointServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostUIAutopilotWaypointServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this post Ui autopilot waypoint service unavailable response has a 2xx status code
func (o *PostUIAutopilotWaypointServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui autopilot waypoint service unavailable response has a 3xx status code
func (o *PostUIAutopilotWaypointServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint service unavailable response has a 4xx status code
func (o *PostUIAutopilotWaypointServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui autopilot waypoint service unavailable response has a 5xx status code
func (o *PostUIAutopilotWaypointServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui autopilot waypoint service unavailable response a status code equal to that given
func (o *PostUIAutopilotWaypointServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the post Ui autopilot waypoint service unavailable response
func (o *PostUIAutopilotWaypointServiceUnavailable) Code() int {
	return 503
}

func (o *PostUIAutopilotWaypointServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIAutopilotWaypointServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIAutopilotWaypointServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostUIAutopilotWaypointServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIAutopilotWaypointGatewayTimeout creates a PostUIAutopilotWaypointGatewayTimeout with default headers values
func NewPostUIAutopilotWaypointGatewayTimeout() *PostUIAutopilotWaypointGatewayTimeout {
	return &PostUIAutopilotWaypointGatewayTimeout{}
}

/*
PostUIAutopilotWaypointGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostUIAutopilotWaypointGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this post Ui autopilot waypoint gateway timeout response has a 2xx status code
func (o *PostUIAutopilotWaypointGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui autopilot waypoint gateway timeout response has a 3xx status code
func (o *PostUIAutopilotWaypointGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui autopilot waypoint gateway timeout response has a 4xx status code
func (o *PostUIAutopilotWaypointGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui autopilot waypoint gateway timeout response has a 5xx status code
func (o *PostUIAutopilotWaypointGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui autopilot waypoint gateway timeout response a status code equal to that given
func (o *PostUIAutopilotWaypointGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the post Ui autopilot waypoint gateway timeout response
func (o *PostUIAutopilotWaypointGatewayTimeout) Code() int {
	return 504
}

func (o *PostUIAutopilotWaypointGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIAutopilotWaypointGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v2/ui/autopilot/waypoint/][%d] postUiAutopilotWaypointGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIAutopilotWaypointGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostUIAutopilotWaypointGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
