// Code generated by go-swagger; DO NOT EDIT.

package fittings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ferocious-space/eveapi/models"
)

// DeleteCharactersCharacterIDFittingsFittingIDReader is a Reader for the DeleteCharactersCharacterIDFittingsFittingID structure.
type DeleteCharactersCharacterIDFittingsFittingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCharactersCharacterIDFittingsFittingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCharactersCharacterIDFittingsFittingIDNoContent creates a DeleteCharactersCharacterIDFittingsFittingIDNoContent with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDNoContent() *DeleteCharactersCharacterIDFittingsFittingIDNoContent {
	return &DeleteCharactersCharacterIDFittingsFittingIDNoContent{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDNoContent describes a response with status code 204, with default header values.

Fitting deleted
*/
type DeleteCharactersCharacterIDFittingsFittingIDNoContent struct {
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdNoContent ", 204)
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCharactersCharacterIDFittingsFittingIDBadRequest creates a DeleteCharactersCharacterIDFittingsFittingIDBadRequest with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDBadRequest() *DeleteCharactersCharacterIDFittingsFittingIDBadRequest {
	return &DeleteCharactersCharacterIDFittingsFittingIDBadRequest{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteCharactersCharacterIDFittingsFittingIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCharactersCharacterIDFittingsFittingIDBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDFittingsFittingIDUnauthorized creates a DeleteCharactersCharacterIDFittingsFittingIDUnauthorized with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDUnauthorized() *DeleteCharactersCharacterIDFittingsFittingIDUnauthorized {
	return &DeleteCharactersCharacterIDFittingsFittingIDUnauthorized{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteCharactersCharacterIDFittingsFittingIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCharactersCharacterIDFittingsFittingIDUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDFittingsFittingIDForbidden creates a DeleteCharactersCharacterIDFittingsFittingIDForbidden with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDForbidden() *DeleteCharactersCharacterIDFittingsFittingIDForbidden {
	return &DeleteCharactersCharacterIDFittingsFittingIDForbidden{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteCharactersCharacterIDFittingsFittingIDForbidden struct {
	Payload *models.Forbidden
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteCharactersCharacterIDFittingsFittingIDForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm creates a DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm() *DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm {
	return &DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDFittingsFittingIDInternalServerError creates a DeleteCharactersCharacterIDFittingsFittingIDInternalServerError with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDInternalServerError() *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError {
	return &DeleteCharactersCharacterIDFittingsFittingIDInternalServerError{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteCharactersCharacterIDFittingsFittingIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable creates a DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable() *DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable {
	return &DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdServiceUnavailable  %+v", 503, o.Payload)
}
func (o *DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout creates a DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout with default headers values
func NewDeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout() *DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout {
	return &DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout{}
}

/* DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/characters/{character_id}/fittings/{fitting_id}/][%d] deleteCharactersCharacterIdFittingsFittingIdGatewayTimeout  %+v", 504, o.Payload)
}
func (o *DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *DeleteCharactersCharacterIDFittingsFittingIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
