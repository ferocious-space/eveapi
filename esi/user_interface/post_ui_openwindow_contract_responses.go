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

// PostUIOpenwindowContractReader is a Reader for the PostUIOpenwindowContract structure.
type PostUIOpenwindowContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIOpenwindowContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUIOpenwindowContractNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUIOpenwindowContractBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUIOpenwindowContractUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUIOpenwindowContractForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostUIOpenwindowContractEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUIOpenwindowContractInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUIOpenwindowContractServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUIOpenwindowContractGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUIOpenwindowContractNoContent creates a PostUIOpenwindowContractNoContent with default headers values
func NewPostUIOpenwindowContractNoContent() *PostUIOpenwindowContractNoContent {
	return &PostUIOpenwindowContractNoContent{}
}

/*
PostUIOpenwindowContractNoContent describes a response with status code 204, with default header values.

Open window request received
*/
type PostUIOpenwindowContractNoContent struct {
}

// IsSuccess returns true when this post Ui openwindow contract no content response has a 2xx status code
func (o *PostUIOpenwindowContractNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Ui openwindow contract no content response has a 3xx status code
func (o *PostUIOpenwindowContractNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract no content response has a 4xx status code
func (o *PostUIOpenwindowContractNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow contract no content response has a 5xx status code
func (o *PostUIOpenwindowContractNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow contract no content response a status code equal to that given
func (o *PostUIOpenwindowContractNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PostUIOpenwindowContractNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractNoContent ", 204)
}

func (o *PostUIOpenwindowContractNoContent) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractNoContent ", 204)
}

func (o *PostUIOpenwindowContractNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIOpenwindowContractBadRequest creates a PostUIOpenwindowContractBadRequest with default headers values
func NewPostUIOpenwindowContractBadRequest() *PostUIOpenwindowContractBadRequest {
	return &PostUIOpenwindowContractBadRequest{}
}

/*
PostUIOpenwindowContractBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostUIOpenwindowContractBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this post Ui openwindow contract bad request response has a 2xx status code
func (o *PostUIOpenwindowContractBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow contract bad request response has a 3xx status code
func (o *PostUIOpenwindowContractBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract bad request response has a 4xx status code
func (o *PostUIOpenwindowContractBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow contract bad request response has a 5xx status code
func (o *PostUIOpenwindowContractBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow contract bad request response a status code equal to that given
func (o *PostUIOpenwindowContractBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUIOpenwindowContractBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIOpenwindowContractBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIOpenwindowContractBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostUIOpenwindowContractBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowContractUnauthorized creates a PostUIOpenwindowContractUnauthorized with default headers values
func NewPostUIOpenwindowContractUnauthorized() *PostUIOpenwindowContractUnauthorized {
	return &PostUIOpenwindowContractUnauthorized{}
}

/*
PostUIOpenwindowContractUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUIOpenwindowContractUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this post Ui openwindow contract unauthorized response has a 2xx status code
func (o *PostUIOpenwindowContractUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow contract unauthorized response has a 3xx status code
func (o *PostUIOpenwindowContractUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract unauthorized response has a 4xx status code
func (o *PostUIOpenwindowContractUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow contract unauthorized response has a 5xx status code
func (o *PostUIOpenwindowContractUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow contract unauthorized response a status code equal to that given
func (o *PostUIOpenwindowContractUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUIOpenwindowContractUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIOpenwindowContractUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIOpenwindowContractUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostUIOpenwindowContractUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowContractForbidden creates a PostUIOpenwindowContractForbidden with default headers values
func NewPostUIOpenwindowContractForbidden() *PostUIOpenwindowContractForbidden {
	return &PostUIOpenwindowContractForbidden{}
}

/*
PostUIOpenwindowContractForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUIOpenwindowContractForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this post Ui openwindow contract forbidden response has a 2xx status code
func (o *PostUIOpenwindowContractForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow contract forbidden response has a 3xx status code
func (o *PostUIOpenwindowContractForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract forbidden response has a 4xx status code
func (o *PostUIOpenwindowContractForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow contract forbidden response has a 5xx status code
func (o *PostUIOpenwindowContractForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow contract forbidden response a status code equal to that given
func (o *PostUIOpenwindowContractForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUIOpenwindowContractForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowContractForbidden) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowContractForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostUIOpenwindowContractForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowContractEnhanceYourCalm creates a PostUIOpenwindowContractEnhanceYourCalm with default headers values
func NewPostUIOpenwindowContractEnhanceYourCalm() *PostUIOpenwindowContractEnhanceYourCalm {
	return &PostUIOpenwindowContractEnhanceYourCalm{}
}

/*
PostUIOpenwindowContractEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostUIOpenwindowContractEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this post Ui openwindow contract enhance your calm response has a 2xx status code
func (o *PostUIOpenwindowContractEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow contract enhance your calm response has a 3xx status code
func (o *PostUIOpenwindowContractEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract enhance your calm response has a 4xx status code
func (o *PostUIOpenwindowContractEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow contract enhance your calm response has a 5xx status code
func (o *PostUIOpenwindowContractEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow contract enhance your calm response a status code equal to that given
func (o *PostUIOpenwindowContractEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

func (o *PostUIOpenwindowContractEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIOpenwindowContractEnhanceYourCalm) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIOpenwindowContractEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostUIOpenwindowContractEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowContractInternalServerError creates a PostUIOpenwindowContractInternalServerError with default headers values
func NewPostUIOpenwindowContractInternalServerError() *PostUIOpenwindowContractInternalServerError {
	return &PostUIOpenwindowContractInternalServerError{}
}

/*
PostUIOpenwindowContractInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostUIOpenwindowContractInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this post Ui openwindow contract internal server error response has a 2xx status code
func (o *PostUIOpenwindowContractInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow contract internal server error response has a 3xx status code
func (o *PostUIOpenwindowContractInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract internal server error response has a 4xx status code
func (o *PostUIOpenwindowContractInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow contract internal server error response has a 5xx status code
func (o *PostUIOpenwindowContractInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow contract internal server error response a status code equal to that given
func (o *PostUIOpenwindowContractInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostUIOpenwindowContractInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowContractInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowContractInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostUIOpenwindowContractInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowContractServiceUnavailable creates a PostUIOpenwindowContractServiceUnavailable with default headers values
func NewPostUIOpenwindowContractServiceUnavailable() *PostUIOpenwindowContractServiceUnavailable {
	return &PostUIOpenwindowContractServiceUnavailable{}
}

/*
PostUIOpenwindowContractServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostUIOpenwindowContractServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this post Ui openwindow contract service unavailable response has a 2xx status code
func (o *PostUIOpenwindowContractServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow contract service unavailable response has a 3xx status code
func (o *PostUIOpenwindowContractServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract service unavailable response has a 4xx status code
func (o *PostUIOpenwindowContractServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow contract service unavailable response has a 5xx status code
func (o *PostUIOpenwindowContractServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow contract service unavailable response a status code equal to that given
func (o *PostUIOpenwindowContractServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostUIOpenwindowContractServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIOpenwindowContractServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIOpenwindowContractServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostUIOpenwindowContractServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowContractGatewayTimeout creates a PostUIOpenwindowContractGatewayTimeout with default headers values
func NewPostUIOpenwindowContractGatewayTimeout() *PostUIOpenwindowContractGatewayTimeout {
	return &PostUIOpenwindowContractGatewayTimeout{}
}

/*
PostUIOpenwindowContractGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostUIOpenwindowContractGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this post Ui openwindow contract gateway timeout response has a 2xx status code
func (o *PostUIOpenwindowContractGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow contract gateway timeout response has a 3xx status code
func (o *PostUIOpenwindowContractGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow contract gateway timeout response has a 4xx status code
func (o *PostUIOpenwindowContractGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow contract gateway timeout response has a 5xx status code
func (o *PostUIOpenwindowContractGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow contract gateway timeout response a status code equal to that given
func (o *PostUIOpenwindowContractGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostUIOpenwindowContractGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIOpenwindowContractGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/contract/][%d] postUiOpenwindowContractGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIOpenwindowContractGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostUIOpenwindowContractGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
