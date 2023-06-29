// Code generated by go-swagger; DO NOT EDIT.

package user_interface

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

// PostUIOpenwindowNewmailReader is a Reader for the PostUIOpenwindowNewmail structure.
type PostUIOpenwindowNewmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUIOpenwindowNewmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUIOpenwindowNewmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUIOpenwindowNewmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUIOpenwindowNewmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUIOpenwindowNewmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostUIOpenwindowNewmailEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostUIOpenwindowNewmailUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUIOpenwindowNewmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUIOpenwindowNewmailServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUIOpenwindowNewmailGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/ui/openwindow/newmail/] post_ui_openwindow_newmail", response, response.Code())
	}
}

// NewPostUIOpenwindowNewmailNoContent creates a PostUIOpenwindowNewmailNoContent with default headers values
func NewPostUIOpenwindowNewmailNoContent() *PostUIOpenwindowNewmailNoContent {
	return &PostUIOpenwindowNewmailNoContent{}
}

/*
PostUIOpenwindowNewmailNoContent describes a response with status code 204, with default header values.

Open window request received
*/
type PostUIOpenwindowNewmailNoContent struct {
}

// IsSuccess returns true when this post Ui openwindow newmail no content response has a 2xx status code
func (o *PostUIOpenwindowNewmailNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Ui openwindow newmail no content response has a 3xx status code
func (o *PostUIOpenwindowNewmailNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail no content response has a 4xx status code
func (o *PostUIOpenwindowNewmailNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow newmail no content response has a 5xx status code
func (o *PostUIOpenwindowNewmailNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow newmail no content response a status code equal to that given
func (o *PostUIOpenwindowNewmailNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post Ui openwindow newmail no content response
func (o *PostUIOpenwindowNewmailNoContent) Code() int {
	return 204
}

func (o *PostUIOpenwindowNewmailNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailNoContent ", 204)
}

func (o *PostUIOpenwindowNewmailNoContent) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailNoContent ", 204)
}

func (o *PostUIOpenwindowNewmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUIOpenwindowNewmailBadRequest creates a PostUIOpenwindowNewmailBadRequest with default headers values
func NewPostUIOpenwindowNewmailBadRequest() *PostUIOpenwindowNewmailBadRequest {
	return &PostUIOpenwindowNewmailBadRequest{}
}

/*
PostUIOpenwindowNewmailBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostUIOpenwindowNewmailBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this post Ui openwindow newmail bad request response has a 2xx status code
func (o *PostUIOpenwindowNewmailBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail bad request response has a 3xx status code
func (o *PostUIOpenwindowNewmailBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail bad request response has a 4xx status code
func (o *PostUIOpenwindowNewmailBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow newmail bad request response has a 5xx status code
func (o *PostUIOpenwindowNewmailBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow newmail bad request response a status code equal to that given
func (o *PostUIOpenwindowNewmailBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post Ui openwindow newmail bad request response
func (o *PostUIOpenwindowNewmailBadRequest) Code() int {
	return 400
}

func (o *PostUIOpenwindowNewmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIOpenwindowNewmailBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailBadRequest  %+v", 400, o.Payload)
}

func (o *PostUIOpenwindowNewmailBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowNewmailUnauthorized creates a PostUIOpenwindowNewmailUnauthorized with default headers values
func NewPostUIOpenwindowNewmailUnauthorized() *PostUIOpenwindowNewmailUnauthorized {
	return &PostUIOpenwindowNewmailUnauthorized{}
}

/*
PostUIOpenwindowNewmailUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUIOpenwindowNewmailUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this post Ui openwindow newmail unauthorized response has a 2xx status code
func (o *PostUIOpenwindowNewmailUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail unauthorized response has a 3xx status code
func (o *PostUIOpenwindowNewmailUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail unauthorized response has a 4xx status code
func (o *PostUIOpenwindowNewmailUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow newmail unauthorized response has a 5xx status code
func (o *PostUIOpenwindowNewmailUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow newmail unauthorized response a status code equal to that given
func (o *PostUIOpenwindowNewmailUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post Ui openwindow newmail unauthorized response
func (o *PostUIOpenwindowNewmailUnauthorized) Code() int {
	return 401
}

func (o *PostUIOpenwindowNewmailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIOpenwindowNewmailUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUIOpenwindowNewmailUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowNewmailForbidden creates a PostUIOpenwindowNewmailForbidden with default headers values
func NewPostUIOpenwindowNewmailForbidden() *PostUIOpenwindowNewmailForbidden {
	return &PostUIOpenwindowNewmailForbidden{}
}

/*
PostUIOpenwindowNewmailForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUIOpenwindowNewmailForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this post Ui openwindow newmail forbidden response has a 2xx status code
func (o *PostUIOpenwindowNewmailForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail forbidden response has a 3xx status code
func (o *PostUIOpenwindowNewmailForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail forbidden response has a 4xx status code
func (o *PostUIOpenwindowNewmailForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow newmail forbidden response has a 5xx status code
func (o *PostUIOpenwindowNewmailForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow newmail forbidden response a status code equal to that given
func (o *PostUIOpenwindowNewmailForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post Ui openwindow newmail forbidden response
func (o *PostUIOpenwindowNewmailForbidden) Code() int {
	return 403
}

func (o *PostUIOpenwindowNewmailForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowNewmailForbidden) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailForbidden  %+v", 403, o.Payload)
}

func (o *PostUIOpenwindowNewmailForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowNewmailEnhanceYourCalm creates a PostUIOpenwindowNewmailEnhanceYourCalm with default headers values
func NewPostUIOpenwindowNewmailEnhanceYourCalm() *PostUIOpenwindowNewmailEnhanceYourCalm {
	return &PostUIOpenwindowNewmailEnhanceYourCalm{}
}

/*
PostUIOpenwindowNewmailEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostUIOpenwindowNewmailEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this post Ui openwindow newmail enhance your calm response has a 2xx status code
func (o *PostUIOpenwindowNewmailEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail enhance your calm response has a 3xx status code
func (o *PostUIOpenwindowNewmailEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail enhance your calm response has a 4xx status code
func (o *PostUIOpenwindowNewmailEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow newmail enhance your calm response has a 5xx status code
func (o *PostUIOpenwindowNewmailEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow newmail enhance your calm response a status code equal to that given
func (o *PostUIOpenwindowNewmailEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the post Ui openwindow newmail enhance your calm response
func (o *PostUIOpenwindowNewmailEnhanceYourCalm) Code() int {
	return 420
}

func (o *PostUIOpenwindowNewmailEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIOpenwindowNewmailEnhanceYourCalm) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostUIOpenwindowNewmailEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowNewmailUnprocessableEntity creates a PostUIOpenwindowNewmailUnprocessableEntity with default headers values
func NewPostUIOpenwindowNewmailUnprocessableEntity() *PostUIOpenwindowNewmailUnprocessableEntity {
	return &PostUIOpenwindowNewmailUnprocessableEntity{}
}

/*
PostUIOpenwindowNewmailUnprocessableEntity describes a response with status code 422, with default header values.

Invalid request
*/
type PostUIOpenwindowNewmailUnprocessableEntity struct {
	Payload *PostUIOpenwindowNewmailUnprocessableEntityBody
}

// IsSuccess returns true when this post Ui openwindow newmail unprocessable entity response has a 2xx status code
func (o *PostUIOpenwindowNewmailUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail unprocessable entity response has a 3xx status code
func (o *PostUIOpenwindowNewmailUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail unprocessable entity response has a 4xx status code
func (o *PostUIOpenwindowNewmailUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post Ui openwindow newmail unprocessable entity response has a 5xx status code
func (o *PostUIOpenwindowNewmailUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post Ui openwindow newmail unprocessable entity response a status code equal to that given
func (o *PostUIOpenwindowNewmailUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the post Ui openwindow newmail unprocessable entity response
func (o *PostUIOpenwindowNewmailUnprocessableEntity) Code() int {
	return 422
}

func (o *PostUIOpenwindowNewmailUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostUIOpenwindowNewmailUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostUIOpenwindowNewmailUnprocessableEntity) GetPayload() *PostUIOpenwindowNewmailUnprocessableEntityBody {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUIOpenwindowNewmailUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowNewmailInternalServerError creates a PostUIOpenwindowNewmailInternalServerError with default headers values
func NewPostUIOpenwindowNewmailInternalServerError() *PostUIOpenwindowNewmailInternalServerError {
	return &PostUIOpenwindowNewmailInternalServerError{}
}

/*
PostUIOpenwindowNewmailInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostUIOpenwindowNewmailInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this post Ui openwindow newmail internal server error response has a 2xx status code
func (o *PostUIOpenwindowNewmailInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail internal server error response has a 3xx status code
func (o *PostUIOpenwindowNewmailInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail internal server error response has a 4xx status code
func (o *PostUIOpenwindowNewmailInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow newmail internal server error response has a 5xx status code
func (o *PostUIOpenwindowNewmailInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow newmail internal server error response a status code equal to that given
func (o *PostUIOpenwindowNewmailInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post Ui openwindow newmail internal server error response
func (o *PostUIOpenwindowNewmailInternalServerError) Code() int {
	return 500
}

func (o *PostUIOpenwindowNewmailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowNewmailInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUIOpenwindowNewmailInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowNewmailServiceUnavailable creates a PostUIOpenwindowNewmailServiceUnavailable with default headers values
func NewPostUIOpenwindowNewmailServiceUnavailable() *PostUIOpenwindowNewmailServiceUnavailable {
	return &PostUIOpenwindowNewmailServiceUnavailable{}
}

/*
PostUIOpenwindowNewmailServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostUIOpenwindowNewmailServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this post Ui openwindow newmail service unavailable response has a 2xx status code
func (o *PostUIOpenwindowNewmailServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail service unavailable response has a 3xx status code
func (o *PostUIOpenwindowNewmailServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail service unavailable response has a 4xx status code
func (o *PostUIOpenwindowNewmailServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow newmail service unavailable response has a 5xx status code
func (o *PostUIOpenwindowNewmailServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow newmail service unavailable response a status code equal to that given
func (o *PostUIOpenwindowNewmailServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the post Ui openwindow newmail service unavailable response
func (o *PostUIOpenwindowNewmailServiceUnavailable) Code() int {
	return 503
}

func (o *PostUIOpenwindowNewmailServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIOpenwindowNewmailServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUIOpenwindowNewmailServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUIOpenwindowNewmailGatewayTimeout creates a PostUIOpenwindowNewmailGatewayTimeout with default headers values
func NewPostUIOpenwindowNewmailGatewayTimeout() *PostUIOpenwindowNewmailGatewayTimeout {
	return &PostUIOpenwindowNewmailGatewayTimeout{}
}

/*
PostUIOpenwindowNewmailGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostUIOpenwindowNewmailGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this post Ui openwindow newmail gateway timeout response has a 2xx status code
func (o *PostUIOpenwindowNewmailGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post Ui openwindow newmail gateway timeout response has a 3xx status code
func (o *PostUIOpenwindowNewmailGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Ui openwindow newmail gateway timeout response has a 4xx status code
func (o *PostUIOpenwindowNewmailGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Ui openwindow newmail gateway timeout response has a 5xx status code
func (o *PostUIOpenwindowNewmailGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post Ui openwindow newmail gateway timeout response a status code equal to that given
func (o *PostUIOpenwindowNewmailGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the post Ui openwindow newmail gateway timeout response
func (o *PostUIOpenwindowNewmailGatewayTimeout) Code() int {
	return 504
}

func (o *PostUIOpenwindowNewmailGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIOpenwindowNewmailGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/ui/openwindow/newmail/][%d] postUiOpenwindowNewmailGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUIOpenwindowNewmailGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostUIOpenwindowNewmailGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostUIOpenwindowNewmailBody post_ui_openwindow_newmail_new_mail
//
// new_mail object
swagger:model PostUIOpenwindowNewmailBody
*/
type PostUIOpenwindowNewmailBody struct {

	// post_ui_openwindow_newmail_body
	//
	// body string
	// Required: true
	// Max Length: 10000
	Body *string `json:"body"`

	// post_ui_openwindow_newmail_recipients
	//
	// recipients array
	// Required: true
	// Max Items: 50
	// Min Items: 1
	Recipients []int32 `json:"recipients"`

	// post_ui_openwindow_newmail_subject
	//
	// subject string
	// Required: true
	// Max Length: 1000
	Subject *string `json:"subject"`

	// post_ui_openwindow_newmail_to_corp_or_alliance_id
	//
	// to_corp_or_alliance_id integer
	ToCorpOrAllianceID int32 `json:"to_corp_or_alliance_id,omitempty"`

	// post_ui_openwindow_newmail_to_mailing_list_id
	//
	// Corporations, alliances and mailing lists are all types of mailing groups. You may only send to one mailing group, at a time, so you may fill out either this field or the to_corp_or_alliance_ids field
	ToMailingListID int32 `json:"to_mailing_list_id,omitempty"`
}

// Validate validates this post UI openwindow newmail body
func (o *PostUIOpenwindowNewmailBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostUIOpenwindowNewmailBody) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("new_mail"+"."+"body", "body", o.Body); err != nil {
		return err
	}

	if err := validate.MaxLength("new_mail"+"."+"body", "body", *o.Body, 10000); err != nil {
		return err
	}

	return nil
}

func (o *PostUIOpenwindowNewmailBody) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("new_mail"+"."+"recipients", "body", o.Recipients); err != nil {
		return err
	}

	iRecipientsSize := int64(len(o.Recipients))

	if err := validate.MinItems("new_mail"+"."+"recipients", "body", iRecipientsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("new_mail"+"."+"recipients", "body", iRecipientsSize, 50); err != nil {
		return err
	}

	return nil
}

func (o *PostUIOpenwindowNewmailBody) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("new_mail"+"."+"subject", "body", o.Subject); err != nil {
		return err
	}

	if err := validate.MaxLength("new_mail"+"."+"subject", "body", *o.Subject, 1000); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post UI openwindow newmail body based on context it is used
func (o *PostUIOpenwindowNewmailBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostUIOpenwindowNewmailBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUIOpenwindowNewmailBody) UnmarshalBinary(b []byte) error {
	var res PostUIOpenwindowNewmailBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostUIOpenwindowNewmailUnprocessableEntityBody post_ui_openwindow_newmail_unprocessable_entity
//
// Unprocessable entity
swagger:model PostUIOpenwindowNewmailUnprocessableEntityBody
*/
type PostUIOpenwindowNewmailUnprocessableEntityBody struct {

	// post_ui_openwindow_newmail_422_unprocessable_entity
	//
	// Unprocessable entity message
	Error string `json:"error,omitempty"`
}

// Validate validates this post UI openwindow newmail unprocessable entity body
func (o *PostUIOpenwindowNewmailUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post UI openwindow newmail unprocessable entity body based on context it is used
func (o *PostUIOpenwindowNewmailUnprocessableEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostUIOpenwindowNewmailUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUIOpenwindowNewmailUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res PostUIOpenwindowNewmailUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
