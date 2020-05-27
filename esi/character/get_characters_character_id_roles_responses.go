// Code generated by go-swagger; DO NOT EDIT.

package character

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

// GetCharactersCharacterIDRolesReader is a Reader for the GetCharactersCharacterIDRoles structure.
type GetCharactersCharacterIDRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDRolesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDRolesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDRolesOK creates a GetCharactersCharacterIDRolesOK with default headers values
func NewGetCharactersCharacterIDRolesOK() *GetCharactersCharacterIDRolesOK {
	return &GetCharactersCharacterIDRolesOK{}
}

/* GetCharactersCharacterIDRolesOK describes a response with status code 200, with default header values.

The character's roles in thier corporation
*/
type GetCharactersCharacterIDRolesOK struct {

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

	Payload *GetCharactersCharacterIDRolesOKBody
}

func (o *GetCharactersCharacterIDRolesOK) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDRolesOK) GetPayload() *GetCharactersCharacterIDRolesOKBody {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(GetCharactersCharacterIDRolesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesNotModified creates a GetCharactersCharacterIDRolesNotModified with default headers values
func NewGetCharactersCharacterIDRolesNotModified() *GetCharactersCharacterIDRolesNotModified {
	return &GetCharactersCharacterIDRolesNotModified{}
}

/* GetCharactersCharacterIDRolesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDRolesNotModified struct {

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

func (o *GetCharactersCharacterIDRolesNotModified) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesNotModified ", 304)
}

func (o *GetCharactersCharacterIDRolesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDRolesBadRequest creates a GetCharactersCharacterIDRolesBadRequest with default headers values
func NewGetCharactersCharacterIDRolesBadRequest() *GetCharactersCharacterIDRolesBadRequest {
	return &GetCharactersCharacterIDRolesBadRequest{}
}

/* GetCharactersCharacterIDRolesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDRolesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDRolesBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesUnauthorized creates a GetCharactersCharacterIDRolesUnauthorized with default headers values
func NewGetCharactersCharacterIDRolesUnauthorized() *GetCharactersCharacterIDRolesUnauthorized {
	return &GetCharactersCharacterIDRolesUnauthorized{}
}

/* GetCharactersCharacterIDRolesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDRolesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDRolesUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesForbidden creates a GetCharactersCharacterIDRolesForbidden with default headers values
func NewGetCharactersCharacterIDRolesForbidden() *GetCharactersCharacterIDRolesForbidden {
	return &GetCharactersCharacterIDRolesForbidden{}
}

/* GetCharactersCharacterIDRolesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDRolesForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDRolesForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesEnhanceYourCalm creates a GetCharactersCharacterIDRolesEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDRolesEnhanceYourCalm() *GetCharactersCharacterIDRolesEnhanceYourCalm {
	return &GetCharactersCharacterIDRolesEnhanceYourCalm{}
}

/* GetCharactersCharacterIDRolesEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDRolesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDRolesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDRolesEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesInternalServerError creates a GetCharactersCharacterIDRolesInternalServerError with default headers values
func NewGetCharactersCharacterIDRolesInternalServerError() *GetCharactersCharacterIDRolesInternalServerError {
	return &GetCharactersCharacterIDRolesInternalServerError{}
}

/* GetCharactersCharacterIDRolesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDRolesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDRolesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesServiceUnavailable creates a GetCharactersCharacterIDRolesServiceUnavailable with default headers values
func NewGetCharactersCharacterIDRolesServiceUnavailable() *GetCharactersCharacterIDRolesServiceUnavailable {
	return &GetCharactersCharacterIDRolesServiceUnavailable{}
}

/* GetCharactersCharacterIDRolesServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDRolesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDRolesServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDRolesGatewayTimeout creates a GetCharactersCharacterIDRolesGatewayTimeout with default headers values
func NewGetCharactersCharacterIDRolesGatewayTimeout() *GetCharactersCharacterIDRolesGatewayTimeout {
	return &GetCharactersCharacterIDRolesGatewayTimeout{}
}

/* GetCharactersCharacterIDRolesGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDRolesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v2/characters/{character_id}/roles/][%d] getCharactersCharacterIdRolesGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDRolesGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDRolesOKBody get_characters_character_id_roles_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDRolesOKBody
*/
type GetCharactersCharacterIDRolesOKBody struct {

	// get_characters_character_id_roles_roles
	//
	// roles array
	// Max Items: 50
	Roles []string `json:"roles"`

	// get_characters_character_id_roles_roles_at_base
	//
	// roles_at_base array
	// Max Items: 50
	RolesAtBase []string `json:"roles_at_base"`

	// get_characters_character_id_roles_roles_at_hq
	//
	// roles_at_hq array
	// Max Items: 50
	RolesAtHq []string `json:"roles_at_hq"`

	// get_characters_character_id_roles_roles_at_other
	//
	// roles_at_other array
	// Max Items: 50
	RolesAtOther []string `json:"roles_at_other"`
}

// Validate validates this get characters character ID roles o k body
func (o *GetCharactersCharacterIDRolesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRolesAtBase(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRolesAtHq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRolesAtOther(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getCharactersCharacterIdRolesOKBodyRolesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdRolesOKBodyRolesItemsEnum = append(getCharactersCharacterIdRolesOKBodyRolesItemsEnum, v)
	}
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRolesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdRolesOKBodyRolesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(o.Roles) { // not required
		return nil
	}

	iRolesSize := int64(len(o.Roles))

	if err := validate.MaxItems("getCharactersCharacterIdRolesOK"+"."+"roles", "body", iRolesSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.Roles); i++ {

		// value enum
		if err := o.validateRolesItemsEnum("getCharactersCharacterIdRolesOK"+"."+"roles"+"."+strconv.Itoa(i), "body", o.Roles[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCharactersCharacterIdRolesOKBodyRolesAtBaseItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdRolesOKBodyRolesAtBaseItemsEnum = append(getCharactersCharacterIdRolesOKBodyRolesAtBaseItemsEnum, v)
	}
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRolesAtBaseItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdRolesOKBodyRolesAtBaseItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRolesAtBase(formats strfmt.Registry) error {
	if swag.IsZero(o.RolesAtBase) { // not required
		return nil
	}

	iRolesAtBaseSize := int64(len(o.RolesAtBase))

	if err := validate.MaxItems("getCharactersCharacterIdRolesOK"+"."+"roles_at_base", "body", iRolesAtBaseSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.RolesAtBase); i++ {

		// value enum
		if err := o.validateRolesAtBaseItemsEnum("getCharactersCharacterIdRolesOK"+"."+"roles_at_base"+"."+strconv.Itoa(i), "body", o.RolesAtBase[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCharactersCharacterIdRolesOKBodyRolesAtHqItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdRolesOKBodyRolesAtHqItemsEnum = append(getCharactersCharacterIdRolesOKBodyRolesAtHqItemsEnum, v)
	}
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRolesAtHqItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdRolesOKBodyRolesAtHqItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRolesAtHq(formats strfmt.Registry) error {
	if swag.IsZero(o.RolesAtHq) { // not required
		return nil
	}

	iRolesAtHqSize := int64(len(o.RolesAtHq))

	if err := validate.MaxItems("getCharactersCharacterIdRolesOK"+"."+"roles_at_hq", "body", iRolesAtHqSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.RolesAtHq); i++ {

		// value enum
		if err := o.validateRolesAtHqItemsEnum("getCharactersCharacterIdRolesOK"+"."+"roles_at_hq"+"."+strconv.Itoa(i), "body", o.RolesAtHq[i]); err != nil {
			return err
		}

	}

	return nil
}

var getCharactersCharacterIdRolesOKBodyRolesAtOtherItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Account_Take_1","Account_Take_2","Account_Take_3","Account_Take_4","Account_Take_5","Account_Take_6","Account_Take_7","Accountant","Auditor","Communications_Officer","Config_Equipment","Config_Starbase_Equipment","Container_Take_1","Container_Take_2","Container_Take_3","Container_Take_4","Container_Take_5","Container_Take_6","Container_Take_7","Contract_Manager","Diplomat","Director","Factory_Manager","Fitting_Manager","Hangar_Query_1","Hangar_Query_2","Hangar_Query_3","Hangar_Query_4","Hangar_Query_5","Hangar_Query_6","Hangar_Query_7","Hangar_Take_1","Hangar_Take_2","Hangar_Take_3","Hangar_Take_4","Hangar_Take_5","Hangar_Take_6","Hangar_Take_7","Junior_Accountant","Personnel_Manager","Rent_Factory_Facility","Rent_Office","Rent_Research_Facility","Security_Officer","Starbase_Defense_Operator","Starbase_Fuel_Technician","Station_Manager","Trader"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdRolesOKBodyRolesAtOtherItemsEnum = append(getCharactersCharacterIdRolesOKBodyRolesAtOtherItemsEnum, v)
	}
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRolesAtOtherItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdRolesOKBodyRolesAtOtherItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDRolesOKBody) validateRolesAtOther(formats strfmt.Registry) error {
	if swag.IsZero(o.RolesAtOther) { // not required
		return nil
	}

	iRolesAtOtherSize := int64(len(o.RolesAtOther))

	if err := validate.MaxItems("getCharactersCharacterIdRolesOK"+"."+"roles_at_other", "body", iRolesAtOtherSize, 50); err != nil {
		return err
	}

	for i := 0; i < len(o.RolesAtOther); i++ {

		// value enum
		if err := o.validateRolesAtOtherItemsEnum("getCharactersCharacterIdRolesOK"+"."+"roles_at_other"+"."+strconv.Itoa(i), "body", o.RolesAtOther[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this get characters character ID roles o k body based on context it is used
func (o *GetCharactersCharacterIDRolesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDRolesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDRolesOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDRolesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
