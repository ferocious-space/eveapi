// Code generated by go-swagger; DO NOT EDIT.

package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// GetCorporationsCorporationIDRolesHistoryReader is a Reader for the GetCorporationsCorporationIDRolesHistory structure.
type GetCorporationsCorporationIDRolesHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDRolesHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDRolesHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDRolesHistoryNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDRolesHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDRolesHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDRolesHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDRolesHistoryEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDRolesHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDRolesHistoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDRolesHistoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDRolesHistoryOK creates a GetCorporationsCorporationIDRolesHistoryOK with default headers values
func NewGetCorporationsCorporationIDRolesHistoryOK() *GetCorporationsCorporationIDRolesHistoryOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDRolesHistoryOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDRolesHistoryOK describes a response with status code 200, with default header values.

List of role changes
*/
type GetCorporationsCorporationIDRolesHistoryOK struct {

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

	/* Maximum page number

	   Format: int32
	   Default: 1
	*/
	XPages int32

	Payload []*GetCorporationsCorporationIDRolesHistoryOKBodyItems0
}

func (o *GetCorporationsCorporationIDRolesHistoryOK) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryOK) GetPayload() []*GetCorporationsCorporationIDRolesHistoryOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// hydrates response header X-Pages
	hdrXPages := response.GetHeader("X-Pages")

	if hdrXPages != "" {
		valxPages, err := swag.ConvertInt32(hdrXPages)
		if err != nil {
			return errors.InvalidType("X-Pages", "header", "int32", hdrXPages)
		}
		o.XPages = valxPages
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesHistoryNotModified creates a GetCorporationsCorporationIDRolesHistoryNotModified with default headers values
func NewGetCorporationsCorporationIDRolesHistoryNotModified() *GetCorporationsCorporationIDRolesHistoryNotModified {
	return &GetCorporationsCorporationIDRolesHistoryNotModified{}
}

/* GetCorporationsCorporationIDRolesHistoryNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDRolesHistoryNotModified struct {

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

func (o *GetCorporationsCorporationIDRolesHistoryNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryNotModified ", 304)
}

func (o *GetCorporationsCorporationIDRolesHistoryNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDRolesHistoryBadRequest creates a GetCorporationsCorporationIDRolesHistoryBadRequest with default headers values
func NewGetCorporationsCorporationIDRolesHistoryBadRequest() *GetCorporationsCorporationIDRolesHistoryBadRequest {
	return &GetCorporationsCorporationIDRolesHistoryBadRequest{}
}

/* GetCorporationsCorporationIDRolesHistoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDRolesHistoryBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDRolesHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesHistoryUnauthorized creates a GetCorporationsCorporationIDRolesHistoryUnauthorized with default headers values
func NewGetCorporationsCorporationIDRolesHistoryUnauthorized() *GetCorporationsCorporationIDRolesHistoryUnauthorized {
	return &GetCorporationsCorporationIDRolesHistoryUnauthorized{}
}

/* GetCorporationsCorporationIDRolesHistoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDRolesHistoryUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDRolesHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesHistoryForbidden creates a GetCorporationsCorporationIDRolesHistoryForbidden with default headers values
func NewGetCorporationsCorporationIDRolesHistoryForbidden() *GetCorporationsCorporationIDRolesHistoryForbidden {
	return &GetCorporationsCorporationIDRolesHistoryForbidden{}
}

/* GetCorporationsCorporationIDRolesHistoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDRolesHistoryForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDRolesHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesHistoryEnhanceYourCalm creates a GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDRolesHistoryEnhanceYourCalm() *GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm {
	return &GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesHistoryInternalServerError creates a GetCorporationsCorporationIDRolesHistoryInternalServerError with default headers values
func NewGetCorporationsCorporationIDRolesHistoryInternalServerError() *GetCorporationsCorporationIDRolesHistoryInternalServerError {
	return &GetCorporationsCorporationIDRolesHistoryInternalServerError{}
}

/* GetCorporationsCorporationIDRolesHistoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDRolesHistoryInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDRolesHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesHistoryServiceUnavailable creates a GetCorporationsCorporationIDRolesHistoryServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDRolesHistoryServiceUnavailable() *GetCorporationsCorporationIDRolesHistoryServiceUnavailable {
	return &GetCorporationsCorporationIDRolesHistoryServiceUnavailable{}
}

/* GetCorporationsCorporationIDRolesHistoryServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDRolesHistoryServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDRolesHistoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDRolesHistoryGatewayTimeout creates a GetCorporationsCorporationIDRolesHistoryGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDRolesHistoryGatewayTimeout() *GetCorporationsCorporationIDRolesHistoryGatewayTimeout {
	return &GetCorporationsCorporationIDRolesHistoryGatewayTimeout{}
}

/* GetCorporationsCorporationIDRolesHistoryGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDRolesHistoryGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDRolesHistoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/corporations/{corporation_id}/roles/history/][%d] getCorporationsCorporationIdRolesHistoryGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDRolesHistoryGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDRolesHistoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDRolesHistoryOKBodyItems0 get_corporations_corporation_id_roles_history_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDRolesHistoryOKBodyItems0
*/
type GetCorporationsCorporationIDRolesHistoryOKBodyItems0 struct {

	// get_corporations_corporation_id_roles_history_changed_at
	//
	// changed_at string
	// Required: true
	// Format: date-time
	ChangedAt *strfmt.DateTime `json:"changed_at"`

	// get_corporations_corporation_id_roles_history_character_id
	//
	// The character whose roles are changed
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporations_corporation_id_roles_history_issuer_id
	//
	// ID of the character who issued this change
	// Required: true
	IssuerID *int32 `json:"issuer_id"`

	// get_corporations_corporation_id_roles_history_new_roles
	//
	// new_roles array
	// Required: true
	// Max Items: 50
	NewRoles []string `json:"new_roles"`

	// get_corporations_corporation_id_roles_history_old_roles
	//
	// old_roles array
	// Required: true
	// Max Items: 50
	OldRoles []string `json:"old_roles"`

	// get_corporations_corporation_id_roles_history_role_type
	//
	// role_type string
	// Required: true
	// Enum: [grantable_roles grantable_roles_at_base grantable_roles_at_hq grantable_roles_at_other roles roles_at_base roles_at_hq roles_at_other]
	RoleType *string `json:"role_type"`
}

// Validate validates this get corporations corporation ID roles history o k body items0
func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChangedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIssuerID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNewRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOldRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoleType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateChangedAt(formats strfmt.Registry) error {

	if err := validate.Required("changed_at", "body", o.ChangedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("changed_at", "body", "date-time", o.ChangedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", o.CharacterID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateIssuerID(formats strfmt.Registry) error {

	if err := validate.Required("issuer_id", "body", o.IssuerID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdRolesHistoryOKBodyItems0NewRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesHistoryOKBodyItems0NewRolesItemsEnum = append(getCorporationsCorporationIdRolesHistoryOKBodyItems0NewRolesItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateNewRolesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesHistoryOKBodyItems0NewRolesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateNewRoles(formats strfmt.Registry) error {

	if err := validate.Required("new_roles", "body", o.NewRoles); err != nil {
		return err
	}

	iNewRolesSize := int64(len(o.NewRoles))

	if err := validate.MaxItems("new_roles", "body", iNewRolesSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.NewRoles); i++ {

		// value enum
		if err := o.validateNewRolesItemsEnum("new_roles"+"."+strconv.Itoa(i), "body", o.NewRoles[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesHistoryOKBodyItems0OldRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesHistoryOKBodyItems0OldRolesItemsEnum = append(getCorporationsCorporationIdRolesHistoryOKBodyItems0OldRolesItemsEnum, v)
	}
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateOldRolesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesHistoryOKBodyItems0OldRolesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateOldRoles(formats strfmt.Registry) error {

	if err := validate.Required("old_roles", "body", o.OldRoles); err != nil {
		return err
	}

	iOldRolesSize := int64(len(o.OldRoles))

	if err := validate.MaxItems("old_roles", "body", iOldRolesSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.OldRoles); i++ {

		// value enum
		if err := o.validateOldRolesItemsEnum("old_roles"+"."+strconv.Itoa(i), "body", o.OldRoles[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCorporationsCorporationIdRolesHistoryOKBodyItems0TypeRoleTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["grantable_roles","grantable_roles_at_base","grantable_roles_at_hq","grantable_roles_at_other","roles","roles_at_base","roles_at_hq","roles_at_other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdRolesHistoryOKBodyItems0TypeRoleTypePropEnum = append(getCorporationsCorporationIdRolesHistoryOKBodyItems0TypeRoleTypePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRoles captures enum value "grantable_roles"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRoles string = "grantable_roles"

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRolesAtBase captures enum value "grantable_roles_at_base"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRolesAtBase string = "grantable_roles_at_base"

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRolesAtHq captures enum value "grantable_roles_at_hq"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRolesAtHq string = "grantable_roles_at_hq"

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRolesAtOther captures enum value "grantable_roles_at_other"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeGrantableRolesAtOther string = "grantable_roles_at_other"

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRoles captures enum value "roles"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRoles string = "roles"

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRolesAtBase captures enum value "roles_at_base"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRolesAtBase string = "roles_at_base"

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRolesAtHq captures enum value "roles_at_hq"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRolesAtHq string = "roles_at_hq"

	// GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRolesAtOther captures enum value "roles_at_other"
	GetCorporationsCorporationIDRolesHistoryOKBodyItems0RoleTypeRolesAtOther string = "roles_at_other"
)

// prop value enum
func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateRoleTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdRolesHistoryOKBodyItems0TypeRoleTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) validateRoleType(formats strfmt.Registry) error {

	if err := validate.Required("role_type", "body", o.RoleType); err != nil {
		return err
	}

	// value enum
	if err := o.validateRoleTypeEnum("role_type", "body", *o.RoleType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID roles history o k body items0 based on context it is used
func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDRolesHistoryOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDRolesHistoryOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
