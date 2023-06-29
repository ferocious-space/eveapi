// Code generated by go-swagger; DO NOT EDIT.

package meta

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
)

// GetVerifyReader is a Reader for the GetVerify structure.
type GetVerifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVerifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVerifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVerifyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVerifyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVerifyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /verify/] get_verify", response, response.Code())
	}
}

// NewGetVerifyOK creates a GetVerifyOK with default headers values
func NewGetVerifyOK() *GetVerifyOK {
	return &GetVerifyOK{}
}

/*
GetVerifyOK describes a response with status code 200, with default header values.

SSO /verify JSON response
*/
type GetVerifyOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *GetVerifyOKBody
}

// IsSuccess returns true when this get verify o k response has a 2xx status code
func (o *GetVerifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get verify o k response has a 3xx status code
func (o *GetVerifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get verify o k response has a 4xx status code
func (o *GetVerifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get verify o k response has a 5xx status code
func (o *GetVerifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get verify o k response a status code equal to that given
func (o *GetVerifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get verify o k response
func (o *GetVerifyOK) Code() int {
	return 200
}

func (o *GetVerifyOK) Error() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyOK  %+v", 200, o.Payload)
}

func (o *GetVerifyOK) String() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyOK  %+v", 200, o.Payload)
}

func (o *GetVerifyOK) GetPayload() *GetVerifyOKBody {
	return o.Payload
}

func (o *GetVerifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Expires
	hdrExpires := response.GetHeader("Expires")

	if hdrExpires != "" {
		o.Expires = hdrExpires
	}

	// hydrates response header Last-Modified
	hdrLastModified := response.GetHeader("Last-Modified")

	if hdrLastModified != "" {
		o.LastModified = hdrLastModified
	}

	o.Payload = new(GetVerifyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVerifyBadRequest creates a GetVerifyBadRequest with default headers values
func NewGetVerifyBadRequest() *GetVerifyBadRequest {
	return &GetVerifyBadRequest{}
}

/*
GetVerifyBadRequest describes a response with status code 400, with default header values.

SSO /verify JSON error
*/
type GetVerifyBadRequest struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload *GetVerifyBadRequestBody
}

// IsSuccess returns true when this get verify bad request response has a 2xx status code
func (o *GetVerifyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get verify bad request response has a 3xx status code
func (o *GetVerifyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get verify bad request response has a 4xx status code
func (o *GetVerifyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get verify bad request response has a 5xx status code
func (o *GetVerifyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get verify bad request response a status code equal to that given
func (o *GetVerifyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get verify bad request response
func (o *GetVerifyBadRequest) Code() int {
	return 400
}

func (o *GetVerifyBadRequest) Error() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyBadRequest  %+v", 400, o.Payload)
}

func (o *GetVerifyBadRequest) String() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyBadRequest  %+v", 400, o.Payload)
}

func (o *GetVerifyBadRequest) GetPayload() *GetVerifyBadRequestBody {
	return o.Payload
}

func (o *GetVerifyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Expires
	hdrExpires := response.GetHeader("Expires")

	if hdrExpires != "" {
		o.Expires = hdrExpires
	}

	// hydrates response header Last-Modified
	hdrLastModified := response.GetHeader("Last-Modified")

	if hdrLastModified != "" {
		o.LastModified = hdrLastModified
	}

	o.Payload = new(GetVerifyBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVerifyUnauthorized creates a GetVerifyUnauthorized with default headers values
func NewGetVerifyUnauthorized() *GetVerifyUnauthorized {
	return &GetVerifyUnauthorized{}
}

/*
GetVerifyUnauthorized describes a response with status code 401, with default header values.

Authorization not provided
*/
type GetVerifyUnauthorized struct {
	Payload *GetVerifyUnauthorizedBody
}

// IsSuccess returns true when this get verify unauthorized response has a 2xx status code
func (o *GetVerifyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get verify unauthorized response has a 3xx status code
func (o *GetVerifyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get verify unauthorized response has a 4xx status code
func (o *GetVerifyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get verify unauthorized response has a 5xx status code
func (o *GetVerifyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get verify unauthorized response a status code equal to that given
func (o *GetVerifyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get verify unauthorized response
func (o *GetVerifyUnauthorized) Code() int {
	return 401
}

func (o *GetVerifyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVerifyUnauthorized) String() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVerifyUnauthorized) GetPayload() *GetVerifyUnauthorizedBody {
	return o.Payload
}

func (o *GetVerifyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVerifyUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVerifyNotFound creates a GetVerifyNotFound with default headers values
func NewGetVerifyNotFound() *GetVerifyNotFound {
	return &GetVerifyNotFound{}
}

/*
GetVerifyNotFound describes a response with status code 404, with default header values.

Unsupported datasource
*/
type GetVerifyNotFound struct {
	Payload *GetVerifyNotFoundBody
}

// IsSuccess returns true when this get verify not found response has a 2xx status code
func (o *GetVerifyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get verify not found response has a 3xx status code
func (o *GetVerifyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get verify not found response has a 4xx status code
func (o *GetVerifyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get verify not found response has a 5xx status code
func (o *GetVerifyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get verify not found response a status code equal to that given
func (o *GetVerifyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get verify not found response
func (o *GetVerifyNotFound) Code() int {
	return 404
}

func (o *GetVerifyNotFound) Error() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyNotFound  %+v", 404, o.Payload)
}

func (o *GetVerifyNotFound) String() string {
	return fmt.Sprintf("[GET /verify/][%d] getVerifyNotFound  %+v", 404, o.Payload)
}

func (o *GetVerifyNotFound) GetPayload() *GetVerifyNotFoundBody {
	return o.Payload
}

func (o *GetVerifyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVerifyNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetVerifyBadRequestBody get_verify_error
//
// SSO /verify JSON
swagger:model GetVerifyBadRequestBody
*/
type GetVerifyBadRequestBody struct {

	// get_verify_error_error
	//
	// Generic error returned by SSO
	// Required: true
	Error *string `json:"error"`

	// get_verify_error_description
	//
	// Detailed error description from SSO
	ErrorDescription string `json:"error_description,omitempty"`
}

// Validate validates this get verify bad request body
func (o *GetVerifyBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetVerifyBadRequestBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getVerifyBadRequest"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get verify bad request body based on context it is used
func (o *GetVerifyBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetVerifyBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVerifyBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVerifyBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetVerifyNotFoundBody get_verify_not_found
//
// Unsupported datasource
swagger:model GetVerifyNotFoundBody
*/
type GetVerifyNotFoundBody struct {

	// get_verify_not_found_error
	//
	// Unsupported datasource
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get verify not found body
func (o *GetVerifyNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetVerifyNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getVerifyNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get verify not found body based on context it is used
func (o *GetVerifyNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetVerifyNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVerifyNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetVerifyNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetVerifyOKBody get_verify_ok
//
// SSO /verify JSON
swagger:model GetVerifyOKBody
*/
type GetVerifyOKBody struct {

	// get_verify_character_id
	//
	// Token owner's character ID
	// Required: true
	CharacterID *int32 `json:"CharacterID"`

	// get_verify_character_name
	//
	// Token owner's character name
	CharacterName string `json:"CharacterName,omitempty"`

	// get_verify_character_owner_hash
	//
	// Hash of the character's owner. If the character is sold or otherwise transferred, this will change
	CharacterOwnerHash string `json:"CharacterOwnerHash,omitempty"`

	// get_verify_expires_on
	//
	// Expiry time of the token (not RFC3339)
	ExpiresOn string `json:"ExpiresOn,omitempty"`

	// get_verify_intellectual_property
	//
	// The IP which generated the token
	IntellectualProperty string `json:"IntellectualProperty,omitempty"`

	// get_verify_scopes
	//
	// Space separated list of scopes the token is valid for
	Scopes string `json:"Scopes,omitempty"`

	// get_verify_token_type
	//
	// Type of access token
	TokenType string `json:"TokenType,omitempty"`
}

// Validate validates this get verify o k body
func (o *GetVerifyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetVerifyOKBody) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("getVerifyOK"+"."+"CharacterID", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get verify o k body based on context it is used
func (o *GetVerifyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetVerifyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVerifyOKBody) UnmarshalBinary(b []byte) error {
	var res GetVerifyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetVerifyUnauthorizedBody get_verify_unauthorized
//
// Authorization not provided
swagger:model GetVerifyUnauthorizedBody
*/
type GetVerifyUnauthorizedBody struct {

	// get_verify_unauthorized_error
	//
	// Authorization not provided
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this get verify unauthorized body
func (o *GetVerifyUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetVerifyUnauthorizedBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("getVerifyUnauthorized"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get verify unauthorized body based on context it is used
func (o *GetVerifyUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetVerifyUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVerifyUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetVerifyUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
