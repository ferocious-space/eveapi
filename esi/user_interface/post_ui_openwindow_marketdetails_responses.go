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

// PostUIOpenwindowMarketdetailsReader is a Reader for the PostUIOpenwindowMarketdetails structure.
type PostUIOpenwindowMarketdetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIOpenwindowMarketdetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUIOpenwindowMarketdetailsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUIOpenwindowMarketdetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUIOpenwindowMarketdetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUIOpenwindowMarketdetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostUIOpenwindowMarketdetailsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUIOpenwindowMarketdetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUIOpenwindowMarketdetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUIOpenwindowMarketdetailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUIOpenwindowMarketdetailsNoContent creates a PostUIOpenwindowMarketdetailsNoContent with default headers values
func NewPostUIOpenwindowMarketdetailsNoContent() *PostUIOpenwindowMarketdetailsNoContent {
	return &PostUIOpenwindowMarketdetailsNoContent{}
}

/* PostUIOpenwindowMarketdetailsNoContent describes a response with status code 204, with default header values.

Open window request received
*/
type PostUIOpenwindowMarketdetailsNoContent struct {
}

func (o *PostUIOpenwindowMarketdetailsNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsNoContent ", 204)
}

func (o *PostUIOpenwindowMarketdetailsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIOpenwindowMarketdetailsBadRequest creates a PostUIOpenwindowMarketdetailsBadRequest with default headers values
func NewPostUIOpenwindowMarketdetailsBadRequest() *PostUIOpenwindowMarketdetailsBadRequest {
	return &PostUIOpenwindowMarketdetailsBadRequest{}
}

/* PostUIOpenwindowMarketdetailsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostUIOpenwindowMarketdetailsBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostUIOpenwindowMarketdetailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsBadRequest  %+v", 400, o.Payload)
}
func (o *PostUIOpenwindowMarketdetailsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostUIOpenwindowMarketdetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowMarketdetailsUnauthorized creates a PostUIOpenwindowMarketdetailsUnauthorized with default headers values
func NewPostUIOpenwindowMarketdetailsUnauthorized() *PostUIOpenwindowMarketdetailsUnauthorized {
	return &PostUIOpenwindowMarketdetailsUnauthorized{}
}

/* PostUIOpenwindowMarketdetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUIOpenwindowMarketdetailsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostUIOpenwindowMarketdetailsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsUnauthorized  %+v", 401, o.Payload)
}
func (o *PostUIOpenwindowMarketdetailsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostUIOpenwindowMarketdetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowMarketdetailsForbidden creates a PostUIOpenwindowMarketdetailsForbidden with default headers values
func NewPostUIOpenwindowMarketdetailsForbidden() *PostUIOpenwindowMarketdetailsForbidden {
	return &PostUIOpenwindowMarketdetailsForbidden{}
}

/* PostUIOpenwindowMarketdetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUIOpenwindowMarketdetailsForbidden struct {
	Payload *models.Forbidden
}

func (o *PostUIOpenwindowMarketdetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsForbidden  %+v", 403, o.Payload)
}
func (o *PostUIOpenwindowMarketdetailsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostUIOpenwindowMarketdetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowMarketdetailsEnhanceYourCalm creates a PostUIOpenwindowMarketdetailsEnhanceYourCalm with default headers values
func NewPostUIOpenwindowMarketdetailsEnhanceYourCalm() *PostUIOpenwindowMarketdetailsEnhanceYourCalm {
	return &PostUIOpenwindowMarketdetailsEnhanceYourCalm{}
}

/* PostUIOpenwindowMarketdetailsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostUIOpenwindowMarketdetailsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostUIOpenwindowMarketdetailsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *PostUIOpenwindowMarketdetailsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostUIOpenwindowMarketdetailsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowMarketdetailsInternalServerError creates a PostUIOpenwindowMarketdetailsInternalServerError with default headers values
func NewPostUIOpenwindowMarketdetailsInternalServerError() *PostUIOpenwindowMarketdetailsInternalServerError {
	return &PostUIOpenwindowMarketdetailsInternalServerError{}
}

/* PostUIOpenwindowMarketdetailsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostUIOpenwindowMarketdetailsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostUIOpenwindowMarketdetailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsInternalServerError  %+v", 500, o.Payload)
}
func (o *PostUIOpenwindowMarketdetailsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostUIOpenwindowMarketdetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowMarketdetailsServiceUnavailable creates a PostUIOpenwindowMarketdetailsServiceUnavailable with default headers values
func NewPostUIOpenwindowMarketdetailsServiceUnavailable() *PostUIOpenwindowMarketdetailsServiceUnavailable {
	return &PostUIOpenwindowMarketdetailsServiceUnavailable{}
}

/* PostUIOpenwindowMarketdetailsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostUIOpenwindowMarketdetailsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostUIOpenwindowMarketdetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostUIOpenwindowMarketdetailsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostUIOpenwindowMarketdetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowMarketdetailsGatewayTimeout creates a PostUIOpenwindowMarketdetailsGatewayTimeout with default headers values
func NewPostUIOpenwindowMarketdetailsGatewayTimeout() *PostUIOpenwindowMarketdetailsGatewayTimeout {
	return &PostUIOpenwindowMarketdetailsGatewayTimeout{}
}

/* PostUIOpenwindowMarketdetailsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostUIOpenwindowMarketdetailsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostUIOpenwindowMarketdetailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/marketdetails/][%d] postUiOpenwindowMarketdetailsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *PostUIOpenwindowMarketdetailsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostUIOpenwindowMarketdetailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
