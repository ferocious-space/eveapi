// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// PostCharactersCharacterIDAssetsLocationsReader is a Reader for the PostCharactersCharacterIDAssetsLocations structure.
type PostCharactersCharacterIDAssetsLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCharactersCharacterIDAssetsLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCharactersCharacterIDAssetsLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCharactersCharacterIDAssetsLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCharactersCharacterIDAssetsLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCharactersCharacterIDAssetsLocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewPostCharactersCharacterIDAssetsLocationsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCharactersCharacterIDAssetsLocationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCharactersCharacterIDAssetsLocationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCharactersCharacterIDAssetsLocationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v2/characters/{character_id}/assets/locations/] post_characters_character_id_assets_locations", response, response.Code())
	}
}

// NewPostCharactersCharacterIDAssetsLocationsOK creates a PostCharactersCharacterIDAssetsLocationsOK with default headers values
func NewPostCharactersCharacterIDAssetsLocationsOK() *PostCharactersCharacterIDAssetsLocationsOK {
	return &PostCharactersCharacterIDAssetsLocationsOK{}
}

/*
PostCharactersCharacterIDAssetsLocationsOK describes a response with status code 200, with default header values.

List of asset locations
*/
type PostCharactersCharacterIDAssetsLocationsOK struct {
	Payload []*PostCharactersCharacterIDAssetsLocationsOKBodyItems0
}

// IsSuccess returns true when this post characters character Id assets locations o k response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post characters character Id assets locations o k response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations o k response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post characters character Id assets locations o k response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post characters character Id assets locations o k response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post characters character Id assets locations o k response
func (o *PostCharactersCharacterIDAssetsLocationsOK) Code() int {
	return 200
}

func (o *PostCharactersCharacterIDAssetsLocationsOK) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsOK  %+v", 200, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsOK) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsOK  %+v", 200, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsOK) GetPayload() []*PostCharactersCharacterIDAssetsLocationsOKBodyItems0 {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsBadRequest creates a PostCharactersCharacterIDAssetsLocationsBadRequest with default headers values
func NewPostCharactersCharacterIDAssetsLocationsBadRequest() *PostCharactersCharacterIDAssetsLocationsBadRequest {
	return &PostCharactersCharacterIDAssetsLocationsBadRequest{}
}

/*
PostCharactersCharacterIDAssetsLocationsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostCharactersCharacterIDAssetsLocationsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this post characters character Id assets locations bad request response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post characters character Id assets locations bad request response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations bad request response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post characters character Id assets locations bad request response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post characters character Id assets locations bad request response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post characters character Id assets locations bad request response
func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) Code() int {
	return 400
}

func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsUnauthorized creates a PostCharactersCharacterIDAssetsLocationsUnauthorized with default headers values
func NewPostCharactersCharacterIDAssetsLocationsUnauthorized() *PostCharactersCharacterIDAssetsLocationsUnauthorized {
	return &PostCharactersCharacterIDAssetsLocationsUnauthorized{}
}

/*
PostCharactersCharacterIDAssetsLocationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostCharactersCharacterIDAssetsLocationsUnauthorized struct {
	Payload *models.Unauthorized
}

// IsSuccess returns true when this post characters character Id assets locations unauthorized response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post characters character Id assets locations unauthorized response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations unauthorized response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post characters character Id assets locations unauthorized response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post characters character Id assets locations unauthorized response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post characters character Id assets locations unauthorized response
func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) Code() int {
	return 401
}

func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsForbidden creates a PostCharactersCharacterIDAssetsLocationsForbidden with default headers values
func NewPostCharactersCharacterIDAssetsLocationsForbidden() *PostCharactersCharacterIDAssetsLocationsForbidden {
	return &PostCharactersCharacterIDAssetsLocationsForbidden{}
}

/*
PostCharactersCharacterIDAssetsLocationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostCharactersCharacterIDAssetsLocationsForbidden struct {
	Payload *models.Forbidden
}

// IsSuccess returns true when this post characters character Id assets locations forbidden response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post characters character Id assets locations forbidden response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations forbidden response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post characters character Id assets locations forbidden response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post characters character Id assets locations forbidden response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post characters character Id assets locations forbidden response
func (o *PostCharactersCharacterIDAssetsLocationsForbidden) Code() int {
	return 403
}

func (o *PostCharactersCharacterIDAssetsLocationsForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsForbidden) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsForbidden  %+v", 403, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsEnhanceYourCalm creates a PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm with default headers values
func NewPostCharactersCharacterIDAssetsLocationsEnhanceYourCalm() *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm {
	return &PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm{}
}

/*
PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this post characters character Id assets locations enhance your calm response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post characters character Id assets locations enhance your calm response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations enhance your calm response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this post characters character Id assets locations enhance your calm response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this post characters character Id assets locations enhance your calm response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the post characters character Id assets locations enhance your calm response
func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) Code() int {
	return 420
}

func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsInternalServerError creates a PostCharactersCharacterIDAssetsLocationsInternalServerError with default headers values
func NewPostCharactersCharacterIDAssetsLocationsInternalServerError() *PostCharactersCharacterIDAssetsLocationsInternalServerError {
	return &PostCharactersCharacterIDAssetsLocationsInternalServerError{}
}

/*
PostCharactersCharacterIDAssetsLocationsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostCharactersCharacterIDAssetsLocationsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this post characters character Id assets locations internal server error response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post characters character Id assets locations internal server error response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations internal server error response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post characters character Id assets locations internal server error response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post characters character Id assets locations internal server error response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post characters character Id assets locations internal server error response
func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) Code() int {
	return 500
}

func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsServiceUnavailable creates a PostCharactersCharacterIDAssetsLocationsServiceUnavailable with default headers values
func NewPostCharactersCharacterIDAssetsLocationsServiceUnavailable() *PostCharactersCharacterIDAssetsLocationsServiceUnavailable {
	return &PostCharactersCharacterIDAssetsLocationsServiceUnavailable{}
}

/*
PostCharactersCharacterIDAssetsLocationsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type PostCharactersCharacterIDAssetsLocationsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this post characters character Id assets locations service unavailable response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post characters character Id assets locations service unavailable response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations service unavailable response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post characters character Id assets locations service unavailable response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post characters character Id assets locations service unavailable response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the post characters character Id assets locations service unavailable response
func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) Code() int {
	return 503
}

func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCharactersCharacterIDAssetsLocationsGatewayTimeout creates a PostCharactersCharacterIDAssetsLocationsGatewayTimeout with default headers values
func NewPostCharactersCharacterIDAssetsLocationsGatewayTimeout() *PostCharactersCharacterIDAssetsLocationsGatewayTimeout {
	return &PostCharactersCharacterIDAssetsLocationsGatewayTimeout{}
}

/*
PostCharactersCharacterIDAssetsLocationsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type PostCharactersCharacterIDAssetsLocationsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this post characters character Id assets locations gateway timeout response has a 2xx status code
func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post characters character Id assets locations gateway timeout response has a 3xx status code
func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post characters character Id assets locations gateway timeout response has a 4xx status code
func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post characters character Id assets locations gateway timeout response has a 5xx status code
func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post characters character Id assets locations gateway timeout response a status code equal to that given
func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the post characters character Id assets locations gateway timeout response
func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) Code() int {
	return 504
}

func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v2/characters/{character_id}/assets/locations/][%d] postCharactersCharacterIdAssetsLocationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *PostCharactersCharacterIDAssetsLocationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostCharactersCharacterIDAssetsLocationsOKBodyItems0 post_characters_character_id_assets_locations_200_ok
//
// 200 ok object
swagger:model PostCharactersCharacterIDAssetsLocationsOKBodyItems0
*/
type PostCharactersCharacterIDAssetsLocationsOKBodyItems0 struct {

	// post_characters_character_id_assets_locations_item_id
	//
	// item_id integer
	// Required: true
	ItemID *int64 `json:"item_id"`

	// position
	// Required: true
	Position *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position `json:"position"`
}

// Validate validates this post characters character ID assets locations o k body items0
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", o.ItemID); err != nil {
		return err
	}

	return nil
}

func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", o.Position); err != nil {
		return err
	}

	if o.Position != nil {
		if err := o.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post characters character ID assets locations o k body items0 based on the context it is used
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0) contextValidatePosition(ctx context.Context, formats strfmt.Registry) error {

	if o.Position != nil {

		if err := o.Position.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDAssetsLocationsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position post_characters_character_id_assets_locations_position
//
// position object
swagger:model PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position
*/
type PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position struct {

	// post_characters_character_id_assets_locations_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// post_characters_character_id_assets_locations_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// post_characters_character_id_assets_locations_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this post characters character ID assets locations o k body items0 position
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateX(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateY(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateZ(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position) validateX(formats strfmt.Registry) error {

	if err := validate.Required("position"+"."+"x", "body", o.X); err != nil {
		return err
	}

	return nil
}

func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position) validateY(formats strfmt.Registry) error {

	if err := validate.Required("position"+"."+"y", "body", o.Y); err != nil {
		return err
	}

	return nil
}

func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("position"+"."+"z", "body", o.Z); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post characters character ID assets locations o k body items0 position based on context it is used
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDAssetsLocationsOKBodyItems0Position
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
