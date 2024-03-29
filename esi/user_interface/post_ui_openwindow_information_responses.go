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

// PostUIOpenwindowInformationReader is a Reader for the PostUIOpenwindowInformation structure.
type PostUIOpenwindowInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIOpenwindowInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUIOpenwindowInformationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUIOpenwindowInformationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUIOpenwindowInformationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUIOpenwindowInformationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostUIOpenwindowInformationEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUIOpenwindowInformationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUIOpenwindowInformationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUIOpenwindowInformationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/ui/openwindow/information/] post_ui_openwindow_information", response, response.Code())
	}
}

// NewPostUIOpenwindowInformationNoContent creates a PostUIOpenwindowInformationNoContent with default headers values
func NewPostUIOpenwindowInformationNoContent() *PostUIOpenwindowInformationNoContent {
	return &PostUIOpenwindowInformationNoContent{}
}

/*
PostUIOpenwindowInformationNoContent describes a response with status code 204, with default header values.

Open window request received
*/
type PostUIOpenwindowInformationNoContent struct {
}

// IsSuccess returns true when this post Ui openwindow information no content response has a 2xx status code
func (o *PostUIOpenwindowInformationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Ui openwindow information no content response has a 3xx status code
func (o *PostUIOpenwindowInformationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information no content response has a 4xx status code
func (o *PostUIOpenwindowInformationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow information no content response has a 5xx status code
func (o *PostUIOpenwindowInformationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow information no content response a status code equal to that given
func (o *PostUIOpenwindowInformationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post Ui openwindow information no content response
func (o *PostUIOpenwindowInformationNoContent) Code() int {
	return 204
}

func (o *PostUIOpenwindowInformationNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationNoContent ", 204)
}

func (o *PostUIOpenwindowInformationNoContent) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationNoContent ", 204)
}

func (o *PostUIOpenwindowInformationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIOpenwindowInformationBadRequest creates a PostUIOpenwindowInformationBadRequest with default headers values
func NewPostUIOpenwindowInformationBadRequest() *PostUIOpenwindowInformationBadRequest {
	return &PostUIOpenwindowInformationBadRequest{}
}

/*
PostUIOpenwindowInformationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostUIOpenwindowInformationBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this post Ui openwindow information bad request response has a 2xx status code
func (o *PostUIOpenwindowInformationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow information bad request response has a 3xx status code
func (o *PostUIOpenwindowInformationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information bad request response has a 4xx status code
func (o *PostUIOpenwindowInformationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow information bad request response has a 5xx status code
func (o *PostUIOpenwindowInformationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow information bad request response a status code equal to that given
func (o *PostUIOpenwindowInformationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post Ui openwindow information bad request response
func (o *PostUIOpenwindowInformationBadRequest) Code() int {
	return 400
}

func (o *PostUIOpenwindowInformationBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIOpenwindowInformationBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIOpenwindowInformationBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostUIOpenwindowInformationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowInformationUnauthorized creates a PostUIOpenwindowInformationUnauthorized with default headers values
func NewPostUIOpenwindowInformationUnauthorized() *PostUIOpenwindowInformationUnauthorized {
	return &PostUIOpenwindowInformationUnauthorized{}
}

/*
PostUIOpenwindowInformationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUIOpenwindowInformationUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this post Ui openwindow information unauthorized response has a 2xx status code
func (o *PostUIOpenwindowInformationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow information unauthorized response has a 3xx status code
func (o *PostUIOpenwindowInformationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information unauthorized response has a 4xx status code
func (o *PostUIOpenwindowInformationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow information unauthorized response has a 5xx status code
func (o *PostUIOpenwindowInformationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow information unauthorized response a status code equal to that given
func (o *PostUIOpenwindowInformationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post Ui openwindow information unauthorized response
func (o *PostUIOpenwindowInformationUnauthorized) Code() int {
	return 401
}

func (o *PostUIOpenwindowInformationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIOpenwindowInformationUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIOpenwindowInformationUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostUIOpenwindowInformationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowInformationForbidden creates a PostUIOpenwindowInformationForbidden with default headers values
func NewPostUIOpenwindowInformationForbidden() *PostUIOpenwindowInformationForbidden {
	return &PostUIOpenwindowInformationForbidden{}
}

/*
PostUIOpenwindowInformationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUIOpenwindowInformationForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this post Ui openwindow information forbidden response has a 2xx status code
func (o *PostUIOpenwindowInformationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow information forbidden response has a 3xx status code
func (o *PostUIOpenwindowInformationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information forbidden response has a 4xx status code
func (o *PostUIOpenwindowInformationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow information forbidden response has a 5xx status code
func (o *PostUIOpenwindowInformationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow information forbidden response a status code equal to that given
func (o *PostUIOpenwindowInformationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post Ui openwindow information forbidden response
func (o *PostUIOpenwindowInformationForbidden) Code() int {
	return 403
}

func (o *PostUIOpenwindowInformationForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowInformationForbidden) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowInformationForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostUIOpenwindowInformationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowInformationEnhanceYourCalm creates a PostUIOpenwindowInformationEnhanceYourCalm with default headers values
func NewPostUIOpenwindowInformationEnhanceYourCalm() *PostUIOpenwindowInformationEnhanceYourCalm {
	return &PostUIOpenwindowInformationEnhanceYourCalm{}
}

/*
PostUIOpenwindowInformationEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostUIOpenwindowInformationEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this post Ui openwindow information enhance your calm response has a 2xx status code
func (o *PostUIOpenwindowInformationEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow information enhance your calm response has a 3xx status code
func (o *PostUIOpenwindowInformationEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information enhance your calm response has a 4xx status code
func (o *PostUIOpenwindowInformationEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow information enhance your calm response has a 5xx status code
func (o *PostUIOpenwindowInformationEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow information enhance your calm response a status code equal to that given
func (o *PostUIOpenwindowInformationEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the post Ui openwindow information enhance your calm response
func (o *PostUIOpenwindowInformationEnhanceYourCalm) Code() int {
	return 420
}

func (o *PostUIOpenwindowInformationEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIOpenwindowInformationEnhanceYourCalm) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIOpenwindowInformationEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostUIOpenwindowInformationEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowInformationInternalServerError creates a PostUIOpenwindowInformationInternalServerError with default headers values
func NewPostUIOpenwindowInformationInternalServerError() *PostUIOpenwindowInformationInternalServerError {
	return &PostUIOpenwindowInformationInternalServerError{}
}

/*
PostUIOpenwindowInformationInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostUIOpenwindowInformationInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this post Ui openwindow information internal server error response has a 2xx status code
func (o *PostUIOpenwindowInformationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow information internal server error response has a 3xx status code
func (o *PostUIOpenwindowInformationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information internal server error response has a 4xx status code
func (o *PostUIOpenwindowInformationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow information internal server error response has a 5xx status code
func (o *PostUIOpenwindowInformationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow information internal server error response a status code equal to that given
func (o *PostUIOpenwindowInformationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post Ui openwindow information internal server error response
func (o *PostUIOpenwindowInformationInternalServerError) Code() int {
	return 500
}

func (o *PostUIOpenwindowInformationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowInformationInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowInformationInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostUIOpenwindowInformationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowInformationServiceUnavailable creates a PostUIOpenwindowInformationServiceUnavailable with default headers values
func NewPostUIOpenwindowInformationServiceUnavailable() *PostUIOpenwindowInformationServiceUnavailable {
	return &PostUIOpenwindowInformationServiceUnavailable{}
}

/*
PostUIOpenwindowInformationServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostUIOpenwindowInformationServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this post Ui openwindow information service unavailable response has a 2xx status code
func (o *PostUIOpenwindowInformationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow information service unavailable response has a 3xx status code
func (o *PostUIOpenwindowInformationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information service unavailable response has a 4xx status code
func (o *PostUIOpenwindowInformationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow information service unavailable response has a 5xx status code
func (o *PostUIOpenwindowInformationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow information service unavailable response a status code equal to that given
func (o *PostUIOpenwindowInformationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the post Ui openwindow information service unavailable response
func (o *PostUIOpenwindowInformationServiceUnavailable) Code() int {
	return 503
}

func (o *PostUIOpenwindowInformationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIOpenwindowInformationServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIOpenwindowInformationServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostUIOpenwindowInformationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowInformationGatewayTimeout creates a PostUIOpenwindowInformationGatewayTimeout with default headers values
func NewPostUIOpenwindowInformationGatewayTimeout() *PostUIOpenwindowInformationGatewayTimeout {
	return &PostUIOpenwindowInformationGatewayTimeout{}
}

/*
PostUIOpenwindowInformationGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostUIOpenwindowInformationGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this post Ui openwindow information gateway timeout response has a 2xx status code
func (o *PostUIOpenwindowInformationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow information gateway timeout response has a 3xx status code
func (o *PostUIOpenwindowInformationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow information gateway timeout response has a 4xx status code
func (o *PostUIOpenwindowInformationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow information gateway timeout response has a 5xx status code
func (o *PostUIOpenwindowInformationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow information gateway timeout response a status code equal to that given
func (o *PostUIOpenwindowInformationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the post Ui openwindow information gateway timeout response
func (o *PostUIOpenwindowInformationGatewayTimeout) Code() int {
	return 504
}

func (o *PostUIOpenwindowInformationGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIOpenwindowInformationGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/information/][%d] postUiOpenwindowInformationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIOpenwindowInformationGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostUIOpenwindowInformationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
