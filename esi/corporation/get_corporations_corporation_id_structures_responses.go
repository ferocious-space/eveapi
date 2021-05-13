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

// GetCorporationsCorporationIDStructuresReader is a Reader for the GetCorporationsCorporationIDStructures structure.
type GetCorporationsCorporationIDStructuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDStructuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorporationsCorporationIDStructuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCorporationsCorporationIDStructuresNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCorporationsCorporationIDStructuresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCorporationsCorporationIDStructuresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCorporationsCorporationIDStructuresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCorporationsCorporationIDStructuresEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorporationsCorporationIDStructuresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCorporationsCorporationIDStructuresServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCorporationsCorporationIDStructuresGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDStructuresOK creates a GetCorporationsCorporationIDStructuresOK with default headers values
func NewGetCorporationsCorporationIDStructuresOK() *GetCorporationsCorporationIDStructuresOK {
	var (
		// initialize headers with default values
		xPagesDefault = int32(1)
	)

	return &GetCorporationsCorporationIDStructuresOK{

		XPages: xPagesDefault,
	}
}

/* GetCorporationsCorporationIDStructuresOK describes a response with status code 200, with default header values.

List of corporation structures' information
*/
type GetCorporationsCorporationIDStructuresOK struct {

	/* The caching mechanism used
	 */
	CacheControl string

	/* The language used in the response
	 */
	ContentLanguage string

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

	Payload []*GetCorporationsCorporationIDStructuresOKBodyItems0
}

func (o *GetCorporationsCorporationIDStructuresOK) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresOK  %+v", 200, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresOK) GetPayload() []*GetCorporationsCorporationIDStructuresOKBodyItems0 {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Cache-Control
	hdrCacheControl := response.GetHeader("Cache-Control")

	if hdrCacheControl != "" {
		o.CacheControl = hdrCacheControl
	}

	// hydrates response header Content-Language
	hdrContentLanguage := response.GetHeader("Content-Language")

	if hdrContentLanguage != "" {
		o.ContentLanguage = hdrContentLanguage
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

// NewGetCorporationsCorporationIDStructuresNotModified creates a GetCorporationsCorporationIDStructuresNotModified with default headers values
func NewGetCorporationsCorporationIDStructuresNotModified() *GetCorporationsCorporationIDStructuresNotModified {
	return &GetCorporationsCorporationIDStructuresNotModified{}
}

/* GetCorporationsCorporationIDStructuresNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCorporationsCorporationIDStructuresNotModified struct {

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

func (o *GetCorporationsCorporationIDStructuresNotModified) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresNotModified ", 304)
}

func (o *GetCorporationsCorporationIDStructuresNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDStructuresBadRequest creates a GetCorporationsCorporationIDStructuresBadRequest with default headers values
func NewGetCorporationsCorporationIDStructuresBadRequest() *GetCorporationsCorporationIDStructuresBadRequest {
	return &GetCorporationsCorporationIDStructuresBadRequest{}
}

/* GetCorporationsCorporationIDStructuresBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCorporationsCorporationIDStructuresBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDStructuresBadRequest) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresBadRequest  %+v", 400, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresUnauthorized creates a GetCorporationsCorporationIDStructuresUnauthorized with default headers values
func NewGetCorporationsCorporationIDStructuresUnauthorized() *GetCorporationsCorporationIDStructuresUnauthorized {
	return &GetCorporationsCorporationIDStructuresUnauthorized{}
}

/* GetCorporationsCorporationIDStructuresUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDStructuresUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDStructuresUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresForbidden creates a GetCorporationsCorporationIDStructuresForbidden with default headers values
func NewGetCorporationsCorporationIDStructuresForbidden() *GetCorporationsCorporationIDStructuresForbidden {
	return &GetCorporationsCorporationIDStructuresForbidden{}
}

/* GetCorporationsCorporationIDStructuresForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCorporationsCorporationIDStructuresForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDStructuresForbidden) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresForbidden  %+v", 403, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresEnhanceYourCalm creates a GetCorporationsCorporationIDStructuresEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDStructuresEnhanceYourCalm() *GetCorporationsCorporationIDStructuresEnhanceYourCalm {
	return &GetCorporationsCorporationIDStructuresEnhanceYourCalm{}
}

/* GetCorporationsCorporationIDStructuresEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCorporationsCorporationIDStructuresEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDStructuresEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresInternalServerError creates a GetCorporationsCorporationIDStructuresInternalServerError with default headers values
func NewGetCorporationsCorporationIDStructuresInternalServerError() *GetCorporationsCorporationIDStructuresInternalServerError {
	return &GetCorporationsCorporationIDStructuresInternalServerError{}
}

/* GetCorporationsCorporationIDStructuresInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCorporationsCorporationIDStructuresInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDStructuresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresServiceUnavailable creates a GetCorporationsCorporationIDStructuresServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDStructuresServiceUnavailable() *GetCorporationsCorporationIDStructuresServiceUnavailable {
	return &GetCorporationsCorporationIDStructuresServiceUnavailable{}
}

/* GetCorporationsCorporationIDStructuresServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDStructuresServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDStructuresServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresGatewayTimeout creates a GetCorporationsCorporationIDStructuresGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDStructuresGatewayTimeout() *GetCorporationsCorporationIDStructuresGatewayTimeout {
	return &GetCorporationsCorporationIDStructuresGatewayTimeout{}
}

/* GetCorporationsCorporationIDStructuresGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDStructuresGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDStructuresGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v4/corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCorporationsCorporationIDStructuresGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCorporationsCorporationIDStructuresGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCorporationsCorporationIDStructuresOKBodyItems0 get_corporations_corporation_id_structures_200_ok
//
// 200 ok object
swagger:model GetCorporationsCorporationIDStructuresOKBodyItems0
*/
type GetCorporationsCorporationIDStructuresOKBodyItems0 struct {

	// get_corporations_corporation_id_structures_corporation_id
	//
	// ID of the corporation that owns the structure
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_corporations_corporation_id_structures_fuel_expires
	//
	// Date on which the structure will run out of fuel
	// Format: date-time
	FuelExpires strfmt.DateTime `json:"fuel_expires,omitempty"`

	// get_corporations_corporation_id_structures_name
	//
	// The structure name
	Name string `json:"name,omitempty"`

	// get_corporations_corporation_id_structures_next_reinforce_apply
	//
	// The date and time when the structure's newly requested reinforcement times (e.g. next_reinforce_hour and next_reinforce_day) will take effect
	// Format: date-time
	NextReinforceApply strfmt.DateTime `json:"next_reinforce_apply,omitempty"`

	// get_corporations_corporation_id_structures_next_reinforce_hour
	//
	// The requested change to reinforce_hour that will take effect at the time shown by next_reinforce_apply
	// Maximum: 23
	// Minimum: 0
	NextReinforceHour *int32 `json:"next_reinforce_hour,omitempty"`

	// get_corporations_corporation_id_structures_profile_id
	//
	// The id of the ACL profile for this citadel
	// Required: true
	ProfileID *int32 `json:"profile_id"`

	// get_corporations_corporation_id_structures_reinforce_hour
	//
	// The hour of day that determines the four hour window when the structure will randomly exit its reinforcement periods and become vulnerable to attack against its armor and/or hull. The structure will become vulnerable at a random time that is +/- 2 hours centered on the value of this property
	// Maximum: 23
	// Minimum: 0
	ReinforceHour *int32 `json:"reinforce_hour,omitempty"`

	// get_corporations_corporation_id_structures_services
	//
	// Contains a list of service upgrades, and their state
	// Max Items: 10
	Services []*GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0 `json:"services"`

	// get_corporations_corporation_id_structures_state
	//
	// state string
	// Required: true
	// Enum: [anchor_vulnerable anchoring armor_reinforce armor_vulnerable deploy_vulnerable fitting_invulnerable hull_reinforce hull_vulnerable online_deprecated onlining_vulnerable shield_vulnerable unanchored unknown]
	State *string `json:"state"`

	// get_corporations_corporation_id_structures_state_timer_end
	//
	// Date at which the structure will move to it's next state
	// Format: date-time
	StateTimerEnd strfmt.DateTime `json:"state_timer_end,omitempty"`

	// get_corporations_corporation_id_structures_state_timer_start
	//
	// Date at which the structure entered it's current state
	// Format: date-time
	StateTimerStart strfmt.DateTime `json:"state_timer_start,omitempty"`

	// get_corporations_corporation_id_structures_structure_id
	//
	// The Item ID of the structure
	// Required: true
	StructureID *int64 `json:"structure_id"`

	// get_corporations_corporation_id_structures_system_id
	//
	// The solar system the structure is in
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_corporations_corporation_id_structures_type_id
	//
	// The type id of the structure
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_corporations_corporation_id_structures_unanchors_at
	//
	// Date at which the structure will unanchor
	// Format: date-time
	UnanchorsAt strfmt.DateTime `json:"unanchors_at,omitempty"`
}

// Validate validates this get corporations corporation ID structures o k body items0
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCorporationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFuelExpires(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNextReinforceApply(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNextReinforceHour(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReinforceHour(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStateTimerEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStateTimerStart(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStructureID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnanchorsAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", o.CorporationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateFuelExpires(formats strfmt.Registry) error {
	if swag.IsZero(o.FuelExpires) { // not required
		return nil
	}

	if err := validate.FormatOf("fuel_expires", "body", "date-time", o.FuelExpires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateNextReinforceApply(formats strfmt.Registry) error {
	if swag.IsZero(o.NextReinforceApply) { // not required
		return nil
	}

	if err := validate.FormatOf("next_reinforce_apply", "body", "date-time", o.NextReinforceApply.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateNextReinforceHour(formats strfmt.Registry) error {
	if swag.IsZero(o.NextReinforceHour) { // not required
		return nil
	}

	if err := validate.MinimumInt("next_reinforce_hour", "body", int64(*o.NextReinforceHour), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("next_reinforce_hour", "body", int64(*o.NextReinforceHour), 23, false); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateProfileID(formats strfmt.Registry) error {

	if err := validate.Required("profile_id", "body", o.ProfileID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateReinforceHour(formats strfmt.Registry) error {
	if swag.IsZero(o.ReinforceHour) { // not required
		return nil
	}

	if err := validate.MinimumInt("reinforce_hour", "body", int64(*o.ReinforceHour), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("reinforce_hour", "body", int64(*o.ReinforceHour), 23, false); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(o.Services) { // not required
		return nil
	}

	iServicesSize := int64(len(o.Services))

	if err := validate.MaxItems("services", "body", iServicesSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(o.Services); i++ {
		if swag.IsZero(o.Services[i]) { // not required
			continue
		}

		if o.Services[i] != nil {
			if err := o.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getCorporationsCorporationIdStructuresOKBodyItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["anchor_vulnerable","anchoring","armor_reinforce","armor_vulnerable","deploy_vulnerable","fitting_invulnerable","hull_reinforce","hull_vulnerable","online_deprecated","onlining_vulnerable","shield_vulnerable","unanchored","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStructuresOKBodyItems0TypeStatePropEnum = append(getCorporationsCorporationIdStructuresOKBodyItems0TypeStatePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateAnchorVulnerable captures enum value "anchor_vulnerable"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateAnchorVulnerable string = "anchor_vulnerable"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateAnchoring captures enum value "anchoring"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateAnchoring string = "anchoring"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateArmorReinforce captures enum value "armor_reinforce"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateArmorReinforce string = "armor_reinforce"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateArmorVulnerable captures enum value "armor_vulnerable"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateArmorVulnerable string = "armor_vulnerable"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateDeployVulnerable captures enum value "deploy_vulnerable"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateDeployVulnerable string = "deploy_vulnerable"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateFittingInvulnerable captures enum value "fitting_invulnerable"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateFittingInvulnerable string = "fitting_invulnerable"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateHullReinforce captures enum value "hull_reinforce"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateHullReinforce string = "hull_reinforce"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateHullVulnerable captures enum value "hull_vulnerable"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateHullVulnerable string = "hull_vulnerable"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateOnlineDeprecated captures enum value "online_deprecated"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateOnlineDeprecated string = "online_deprecated"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateOnliningVulnerable captures enum value "onlining_vulnerable"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateOnliningVulnerable string = "onlining_vulnerable"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateShieldVulnerable captures enum value "shield_vulnerable"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateShieldVulnerable string = "shield_vulnerable"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateUnanchored captures enum value "unanchored"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateUnanchored string = "unanchored"

	// GetCorporationsCorporationIDStructuresOKBodyItems0StateUnknown captures enum value "unknown"
	GetCorporationsCorporationIDStructuresOKBodyItems0StateUnknown string = "unknown"
)

// prop value enum
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdStructuresOKBodyItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateStateTimerEnd(formats strfmt.Registry) error {
	if swag.IsZero(o.StateTimerEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("state_timer_end", "body", "date-time", o.StateTimerEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateStateTimerStart(formats strfmt.Registry) error {
	if swag.IsZero(o.StateTimerStart) { // not required
		return nil
	}

	if err := validate.FormatOf("state_timer_start", "body", "date-time", o.StateTimerStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", o.StructureID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", o.SystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", o.TypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) validateUnanchorsAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UnanchorsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("unanchors_at", "body", "date-time", o.UnanchorsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get corporations corporation ID structures o k body items0 based on the context it is used
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Services); i++ {

		if o.Services[i] != nil {
			if err := o.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDStructuresOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0 get_corporations_corporation_id_structures_service
//
// service object
swagger:model GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0
*/
type GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0 struct {

	// get_corporations_corporation_id_structures_service_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_corporations_corporation_id_structures_service_state
	//
	// state string
	// Required: true
	// Enum: [online offline cleanup]
	State *string `json:"state"`
}

// Validate validates this get corporations corporation ID structures o k body items0 services items0
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStructuresOKBodyItems0ServicesItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline","cleanup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStructuresOKBodyItems0ServicesItems0TypeStatePropEnum = append(getCorporationsCorporationIdStructuresOKBodyItems0ServicesItems0TypeStatePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0StateOnline captures enum value "online"
	GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0StateOnline string = "online"

	// GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0StateOffline captures enum value "offline"
	GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0StateOffline string = "offline"

	// GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0StateCleanup captures enum value "cleanup"
	GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0StateCleanup string = "cleanup"
)

// prop value enum
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCorporationsCorporationIdStructuresOKBodyItems0ServicesItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get corporations corporation ID structures o k body items0 services items0 based on context it is used
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDStructuresOKBodyItems0ServicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
