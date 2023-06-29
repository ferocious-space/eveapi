// Code generated by go-swagger; DO NOT EDIT.

package sovereignty

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

// GetSovereigntyCampaignsReader is a Reader for the GetSovereigntyCampaigns structure.
type GetSovereigntyCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSovereigntyCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetSovereigntyCampaignsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetSovereigntyCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetSovereigntyCampaignsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSovereigntyCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSovereigntyCampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSovereigntyCampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/sovereignty/campaigns/] get_sovereignty_campaigns", response, response.Code())
	}
}

// NewGetSovereigntyCampaignsOK creates a GetSovereigntyCampaignsOK with default headers values
func NewGetSovereigntyCampaignsOK() *GetSovereigntyCampaignsOK {
	return &GetSovereigntyCampaignsOK{}
}

/*
GetSovereigntyCampaignsOK describes a response with status code 200, with default header values.

A list of sovereignty campaigns
*/
type GetSovereigntyCampaignsOK struct {

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

	Payload []*GetSovereigntyCampaignsOKBodyItems0
}

// IsSuccess returns true when this get sovereignty campaigns o k response has a 2xx status code
func (o *GetSovereigntyCampaignsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sovereignty campaigns o k response has a 3xx status code
func (o *GetSovereigntyCampaignsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty campaigns o k response has a 4xx status code
func (o *GetSovereigntyCampaignsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty campaigns o k response has a 5xx status code
func (o *GetSovereigntyCampaignsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty campaigns o k response a status code equal to that given
func (o *GetSovereigntyCampaignsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sovereignty campaigns o k response
func (o *GetSovereigntyCampaignsOK) Code() int {
	return 200
}

func (o *GetSovereigntyCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyCampaignsOK) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyCampaignsOK) GetPayload() []*GetSovereigntyCampaignsOKBodyItems0 {
	return o.Payload
}

func (o *GetSovereigntyCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyCampaignsNotModified creates a GetSovereigntyCampaignsNotModified with default headers values
func NewGetSovereigntyCampaignsNotModified() *GetSovereigntyCampaignsNotModified {
	return &GetSovereigntyCampaignsNotModified{}
}

/*
GetSovereigntyCampaignsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetSovereigntyCampaignsNotModified struct {

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

// IsSuccess returns true when this get sovereignty campaigns not modified response has a 2xx status code
func (o *GetSovereigntyCampaignsNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty campaigns not modified response has a 3xx status code
func (o *GetSovereigntyCampaignsNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get sovereignty campaigns not modified response has a 4xx status code
func (o *GetSovereigntyCampaignsNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty campaigns not modified response has a 5xx status code
func (o *GetSovereigntyCampaignsNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty campaigns not modified response a status code equal to that given
func (o *GetSovereigntyCampaignsNotModified) IsCode(code int) bool {
	return code == 304
}

// Code gets the status code for the get sovereignty campaigns not modified response
func (o *GetSovereigntyCampaignsNotModified) Code() int {
	return 304
}

func (o *GetSovereigntyCampaignsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsNotModified ", 304)
}

func (o *GetSovereigntyCampaignsNotModified) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsNotModified ", 304)
}

func (o *GetSovereigntyCampaignsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyCampaignsBadRequest creates a GetSovereigntyCampaignsBadRequest with default headers values
func NewGetSovereigntyCampaignsBadRequest() *GetSovereigntyCampaignsBadRequest {
	return &GetSovereigntyCampaignsBadRequest{}
}

/*
GetSovereigntyCampaignsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSovereigntyCampaignsBadRequest struct {
	Payload *models.BadRequest
}

// IsSuccess returns true when this get sovereignty campaigns bad request response has a 2xx status code
func (o *GetSovereigntyCampaignsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty campaigns bad request response has a 3xx status code
func (o *GetSovereigntyCampaignsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty campaigns bad request response has a 4xx status code
func (o *GetSovereigntyCampaignsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sovereignty campaigns bad request response has a 5xx status code
func (o *GetSovereigntyCampaignsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty campaigns bad request response a status code equal to that given
func (o *GetSovereigntyCampaignsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get sovereignty campaigns bad request response
func (o *GetSovereigntyCampaignsBadRequest) Code() int {
	return 400
}

func (o *GetSovereigntyCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyCampaignsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyCampaignsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetSovereigntyCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsEnhanceYourCalm creates a GetSovereigntyCampaignsEnhanceYourCalm with default headers values
func NewGetSovereigntyCampaignsEnhanceYourCalm() *GetSovereigntyCampaignsEnhanceYourCalm {
	return &GetSovereigntyCampaignsEnhanceYourCalm{}
}

/*
GetSovereigntyCampaignsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetSovereigntyCampaignsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

// IsSuccess returns true when this get sovereignty campaigns enhance your calm response has a 2xx status code
func (o *GetSovereigntyCampaignsEnhanceYourCalm) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty campaigns enhance your calm response has a 3xx status code
func (o *GetSovereigntyCampaignsEnhanceYourCalm) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty campaigns enhance your calm response has a 4xx status code
func (o *GetSovereigntyCampaignsEnhanceYourCalm) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sovereignty campaigns enhance your calm response has a 5xx status code
func (o *GetSovereigntyCampaignsEnhanceYourCalm) IsServerError() bool {
	return false
}

// IsCode returns true when this get sovereignty campaigns enhance your calm response a status code equal to that given
func (o *GetSovereigntyCampaignsEnhanceYourCalm) IsCode(code int) bool {
	return code == 420
}

// Code gets the status code for the get sovereignty campaigns enhance your calm response
func (o *GetSovereigntyCampaignsEnhanceYourCalm) Code() int {
	return 420
}

func (o *GetSovereigntyCampaignsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyCampaignsEnhanceYourCalm) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyCampaignsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetSovereigntyCampaignsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsInternalServerError creates a GetSovereigntyCampaignsInternalServerError with default headers values
func NewGetSovereigntyCampaignsInternalServerError() *GetSovereigntyCampaignsInternalServerError {
	return &GetSovereigntyCampaignsInternalServerError{}
}

/*
GetSovereigntyCampaignsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSovereigntyCampaignsInternalServerError struct {
	Payload *models.InternalServerError
}

// IsSuccess returns true when this get sovereignty campaigns internal server error response has a 2xx status code
func (o *GetSovereigntyCampaignsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty campaigns internal server error response has a 3xx status code
func (o *GetSovereigntyCampaignsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty campaigns internal server error response has a 4xx status code
func (o *GetSovereigntyCampaignsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty campaigns internal server error response has a 5xx status code
func (o *GetSovereigntyCampaignsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty campaigns internal server error response a status code equal to that given
func (o *GetSovereigntyCampaignsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get sovereignty campaigns internal server error response
func (o *GetSovereigntyCampaignsInternalServerError) Code() int {
	return 500
}

func (o *GetSovereigntyCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyCampaignsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyCampaignsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetSovereigntyCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsServiceUnavailable creates a GetSovereigntyCampaignsServiceUnavailable with default headers values
func NewGetSovereigntyCampaignsServiceUnavailable() *GetSovereigntyCampaignsServiceUnavailable {
	return &GetSovereigntyCampaignsServiceUnavailable{}
}

/*
GetSovereigntyCampaignsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetSovereigntyCampaignsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

// IsSuccess returns true when this get sovereignty campaigns service unavailable response has a 2xx status code
func (o *GetSovereigntyCampaignsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty campaigns service unavailable response has a 3xx status code
func (o *GetSovereigntyCampaignsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty campaigns service unavailable response has a 4xx status code
func (o *GetSovereigntyCampaignsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty campaigns service unavailable response has a 5xx status code
func (o *GetSovereigntyCampaignsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty campaigns service unavailable response a status code equal to that given
func (o *GetSovereigntyCampaignsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

// Code gets the status code for the get sovereignty campaigns service unavailable response
func (o *GetSovereigntyCampaignsServiceUnavailable) Code() int {
	return 503
}

func (o *GetSovereigntyCampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyCampaignsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyCampaignsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetSovereigntyCampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsGatewayTimeout creates a GetSovereigntyCampaignsGatewayTimeout with default headers values
func NewGetSovereigntyCampaignsGatewayTimeout() *GetSovereigntyCampaignsGatewayTimeout {
	return &GetSovereigntyCampaignsGatewayTimeout{}
}

/*
GetSovereigntyCampaignsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetSovereigntyCampaignsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

// IsSuccess returns true when this get sovereignty campaigns gateway timeout response has a 2xx status code
func (o *GetSovereigntyCampaignsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sovereignty campaigns gateway timeout response has a 3xx status code
func (o *GetSovereigntyCampaignsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sovereignty campaigns gateway timeout response has a 4xx status code
func (o *GetSovereigntyCampaignsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sovereignty campaigns gateway timeout response has a 5xx status code
func (o *GetSovereigntyCampaignsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get sovereignty campaigns gateway timeout response a status code equal to that given
func (o *GetSovereigntyCampaignsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the get sovereignty campaigns gateway timeout response
func (o *GetSovereigntyCampaignsGatewayTimeout) Code() int {
	return 504
}

func (o *GetSovereigntyCampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyCampaignsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sovereignty/campaigns/][%d] getSovereigntyCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyCampaignsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetSovereigntyCampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetSovereigntyCampaignsOKBodyItems0 get_sovereignty_campaigns_200_ok
//
// 200 ok object
swagger:model GetSovereigntyCampaignsOKBodyItems0
*/
type GetSovereigntyCampaignsOKBodyItems0 struct {

	// get_sovereignty_campaigns_attackers_score
	//
	// Score for all attacking parties, only present in Defense Events.
	//
	AttackersScore float32 `json:"attackers_score,omitempty"`

	// get_sovereignty_campaigns_campaign_id
	//
	// Unique ID for this campaign.
	// Required: true
	CampaignID *int32 `json:"campaign_id"`

	// get_sovereignty_campaigns_constellation_id
	//
	// The constellation in which the campaign will take place.
	//
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_sovereignty_campaigns_defender_id
	//
	// Defending alliance, only present in Defense Events
	//
	DefenderID int32 `json:"defender_id,omitempty"`

	// get_sovereignty_campaigns_defender_score
	//
	// Score for the defending alliance, only present in Defense Events.
	//
	DefenderScore float32 `json:"defender_score,omitempty"`

	// get_sovereignty_campaigns_event_type
	//
	// Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as "Defense Events", station_freeport as "Freeport Events".
	//
	// Required: true
	// Enum: [tcu_defense ihub_defense station_defense station_freeport]
	EventType *string `json:"event_type"`

	// get_sovereignty_campaigns_participants
	//
	// Alliance participating and their respective scores, only present in Freeport Events.
	//
	// Max Items: 5000
	Participants []*GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0 `json:"participants"`

	// get_sovereignty_campaigns_solar_system_id
	//
	// The solar system the structure is located in.
	//
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_sovereignty_campaigns_start_time
	//
	// Time the event is scheduled to start.
	//
	// Required: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time"`

	// get_sovereignty_campaigns_structure_id
	//
	// The structure item ID that is related to this campaign.
	//
	// Required: true
	StructureID *int64 `json:"structure_id"`
}

// Validate validates this get sovereignty campaigns o k body items0
func (o *GetSovereigntyCampaignsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCampaignID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConstellationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParticipants(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStructureID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateCampaignID(formats strfmt.Registry) error {

	if err := validate.Required("campaign_id", "body", o.CampaignID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("constellation_id", "body", o.ConstellationID); err != nil {
		return err
	}

	return nil
}

var getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcu_defense","ihub_defense","station_defense","station_freeport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum = append(getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum, v)
	}
}

const (

	// GetSovereigntyCampaignsOKBodyItems0EventTypeTcuDefense captures enum value "tcu_defense"
	GetSovereigntyCampaignsOKBodyItems0EventTypeTcuDefense string = "tcu_defense"

	// GetSovereigntyCampaignsOKBodyItems0EventTypeIhubDefense captures enum value "ihub_defense"
	GetSovereigntyCampaignsOKBodyItems0EventTypeIhubDefense string = "ihub_defense"

	// GetSovereigntyCampaignsOKBodyItems0EventTypeStationDefense captures enum value "station_defense"
	GetSovereigntyCampaignsOKBodyItems0EventTypeStationDefense string = "station_defense"

	// GetSovereigntyCampaignsOKBodyItems0EventTypeStationFreeport captures enum value "station_freeport"
	GetSovereigntyCampaignsOKBodyItems0EventTypeStationFreeport string = "station_freeport"
)

// prop value enum
func (o *GetSovereigntyCampaignsOKBodyItems0) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getSovereigntyCampaignsOKBodyItems0TypeEventTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("event_type", "body", o.EventType); err != nil {
		return err
	}

	// value enum
	if err := o.validateEventTypeEnum("event_type", "body", *o.EventType); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateParticipants(formats strfmt.Registry) error {
	if swag.IsZero(o.Participants) { // not required
		return nil
	}

	iParticipantsSize := int64(len(o.Participants))

	if err := validate.MaxItems("participants", "body", iParticipantsSize, 5000); err != nil {
		return err
	}

	for i := 0; i < len(o.Participants); i++ {
		if swag.IsZero(o.Participants[i]) { // not required
			continue
		}

		if o.Participants[i] != nil {
			if err := o.Participants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", o.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", o.StartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("start_time", "body", "date-time", o.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", o.StructureID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get sovereignty campaigns o k body items0 based on the context it is used
func (o *GetSovereigntyCampaignsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParticipants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0) contextValidateParticipants(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Participants); i++ {

		if o.Participants[i] != nil {

			if swag.IsZero(o.Participants[i]) { // not required
				return nil
			}

			if err := o.Participants[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSovereigntyCampaignsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSovereigntyCampaignsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyCampaignsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0 get_sovereignty_campaigns_participant
//
// participant object
swagger:model GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0
*/
type GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0 struct {

	// get_sovereignty_campaigns_alliance_id
	//
	// alliance_id integer
	// Required: true
	AllianceID *int32 `json:"alliance_id"`

	// get_sovereignty_campaigns_score
	//
	// score number
	// Required: true
	Score *float32 `json:"score"`
}

// Validate validates this get sovereignty campaigns o k body items0 participants items0
func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllianceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) validateAllianceID(formats strfmt.Registry) error {

	if err := validate.Required("alliance_id", "body", o.AllianceID); err != nil {
		return err
	}

	return nil
}

func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) validateScore(formats strfmt.Registry) error {

	if err := validate.Required("score", "body", o.Score); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get sovereignty campaigns o k body items0 participants items0 based on context it is used
func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyCampaignsOKBodyItems0ParticipantsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
