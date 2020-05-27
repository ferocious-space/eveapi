// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// GetCorporationsCorporationIDWalletsDivisionTransactionsReader is a Reader for the GetCorporationsCorporationIDWalletsDivisionTransactions structure.
type GetCorporationsCorporationIDWalletsDivisionTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsOK creates a GetCorporationsCorporationIDWalletsDivisionTransactionsOK with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsOK() *GetCorporationsCorporationIDWalletsDivisionTransactionsOK {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsOK{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsOK describes a response with status code 200, with default header values.

Wallet transactions
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOK) GetPayload() []*GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsNotModified creates a GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsNotModified() *GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* RFC7232 compliant entity tag
	 */
	ETag string

	/* RFC7231 formatted datetime string
	 */
	Expires string

	/* RFC7231 formatted datetime string
	 */
	LastModified string
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsNotModified ", 304)
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header ETag
	hdrETag := response.GetHeader("ETag")

	if hdrETag != "" {
		o.ETag = hdrETag
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

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest creates a GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest() *GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized creates a GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized() *GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsForbidden creates a GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsForbidden() *GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm creates a GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm() *GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError creates a GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError() *GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable creates a GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable() *GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout creates a GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout() *GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout {
	return &GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout{}
}

/* GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/corporations/{corporation_id}/wallets/{division}/transactions/][%d] getCorporationsCorporationIdWalletsDivisionTransactionsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0 get_corporations_corporation_id_wallets_division_transactions_200_ok
//
// wallet transaction
swagger:model GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0
*/
type GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0 struct {

	// get_corporations_corporation_id_wallets_division_transactions_client_id
	//
	// client_id integer
	// Required: true
	ClientID *int32 `json:"client_id"`

	// get_corporations_corporation_id_wallets_division_transactions_date
	//
	// Date and time of transaction
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// get_corporations_corporation_id_wallets_division_transactions_is_buy
	//
	// is_buy boolean
	// Required: true
	IsBuy *bool `json:"is_buy"`

	// get_corporations_corporation_id_wallets_division_transactions_journal_ref_id
	//
	// -1 if there is no corresponding wallet journal entry
	// Required: true
	JournalRefID *int64 `json:"journal_ref_id"`

	// get_corporations_corporation_id_wallets_division_transactions_location_id
	//
	// location_id integer
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_corporations_corporation_id_wallets_division_transactions_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_corporations_corporation_id_wallets_division_transactions_transaction_id
	//
	// Unique transaction ID
	// Required: true
	TransactionID *int64 `json:"transaction_id"`

	// get_corporations_corporation_id_wallets_division_transactions_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_corporations_corporation_id_wallets_division_transactions_unit_price
	//
	// Amount paid per unit
	// Required: true
	UnitPrice *float64 `json:"unit_price"`
}

// Validate validates this get corporations corporation ID wallets division transactions o k body items0
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIsBuy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateJournalRefID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnitPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", o.ClientID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", o.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", o.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateIsBuy(formats strfmt.Registry) error {

	if err := validate.Required("is_buy", "body", o.IsBuy); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateJournalRefID(formats strfmt.Registry) error {

	if err := validate.Required("journal_ref_id", "body", o.JournalRefID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", o.LocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", o.Quantity); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transaction_id", "body", o.TransactionID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) validateUnitPrice(formats strfmt.Registry) error {

	if err := validate.Required("unit_price", "body", o.UnitPrice); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID wallets division transactions o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDWalletsDivisionTransactionsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
