// Code generated by go-swagger; DO NOT EDIT.

package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/ferocious-space/eveapi/models"
)

// GetCharactersCharacterIDIndustryJobsReader is a Reader for the GetCharactersCharacterIDIndustryJobs structure.
type GetCharactersCharacterIDIndustryJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDIndustryJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCharactersCharacterIDIndustryJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetCharactersCharacterIDIndustryJobsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetCharactersCharacterIDIndustryJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCharactersCharacterIDIndustryJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCharactersCharacterIDIndustryJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 420:
		result := NewGetCharactersCharacterIDIndustryJobsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCharactersCharacterIDIndustryJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCharactersCharacterIDIndustryJobsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCharactersCharacterIDIndustryJobsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCharactersCharacterIDIndustryJobsOK creates a GetCharactersCharacterIDIndustryJobsOK with default headers values
func NewGetCharactersCharacterIDIndustryJobsOK() *GetCharactersCharacterIDIndustryJobsOK {
	return &GetCharactersCharacterIDIndustryJobsOK{}
}

/* GetCharactersCharacterIDIndustryJobsOK describes a response with status code 200, with default header values.

Industry jobs placed by a character
*/
type GetCharactersCharacterIDIndustryJobsOK struct {

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

	Payload []*GetCharactersCharacterIDIndustryJobsOKBodyItems0
}

func (o *GetCharactersCharacterIDIndustryJobsOK) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsOK  %+v", 200, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsOK) GetPayload() []*GetCharactersCharacterIDIndustryJobsOKBodyItems0 {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDIndustryJobsNotModified creates a GetCharactersCharacterIDIndustryJobsNotModified with default headers values
func NewGetCharactersCharacterIDIndustryJobsNotModified() *GetCharactersCharacterIDIndustryJobsNotModified {
	return &GetCharactersCharacterIDIndustryJobsNotModified{}
}

/* GetCharactersCharacterIDIndustryJobsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type GetCharactersCharacterIDIndustryJobsNotModified struct {

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

func (o *GetCharactersCharacterIDIndustryJobsNotModified) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsNotModified ", 304)
}

func (o *GetCharactersCharacterIDIndustryJobsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDIndustryJobsBadRequest creates a GetCharactersCharacterIDIndustryJobsBadRequest with default headers values
func NewGetCharactersCharacterIDIndustryJobsBadRequest() *GetCharactersCharacterIDIndustryJobsBadRequest {
	return &GetCharactersCharacterIDIndustryJobsBadRequest{}
}

/* GetCharactersCharacterIDIndustryJobsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCharactersCharacterIDIndustryJobsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDIndustryJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsBadRequest  %+v", 400, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsBadRequest) GetPayload() *models.BadRequest {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDIndustryJobsUnauthorized creates a GetCharactersCharacterIDIndustryJobsUnauthorized with default headers values
func NewGetCharactersCharacterIDIndustryJobsUnauthorized() *GetCharactersCharacterIDIndustryJobsUnauthorized {
	return &GetCharactersCharacterIDIndustryJobsUnauthorized{}
}

/* GetCharactersCharacterIDIndustryJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCharactersCharacterIDIndustryJobsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDIndustryJobsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsUnauthorized) GetPayload() *models.Unauthorized {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDIndustryJobsForbidden creates a GetCharactersCharacterIDIndustryJobsForbidden with default headers values
func NewGetCharactersCharacterIDIndustryJobsForbidden() *GetCharactersCharacterIDIndustryJobsForbidden {
	return &GetCharactersCharacterIDIndustryJobsForbidden{}
}

/* GetCharactersCharacterIDIndustryJobsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCharactersCharacterIDIndustryJobsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDIndustryJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsForbidden  %+v", 403, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsForbidden) GetPayload() *models.Forbidden {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDIndustryJobsEnhanceYourCalm creates a GetCharactersCharacterIDIndustryJobsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDIndustryJobsEnhanceYourCalm() *GetCharactersCharacterIDIndustryJobsEnhanceYourCalm {
	return &GetCharactersCharacterIDIndustryJobsEnhanceYourCalm{}
}

/* GetCharactersCharacterIDIndustryJobsEnhanceYourCalm describes a response with status code 420, with default header values.

Error limited
*/
type GetCharactersCharacterIDIndustryJobsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDIndustryJobsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsEnhanceYourCalm  %+v", 420, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsEnhanceYourCalm) GetPayload() *models.ErrorLimited {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDIndustryJobsInternalServerError creates a GetCharactersCharacterIDIndustryJobsInternalServerError with default headers values
func NewGetCharactersCharacterIDIndustryJobsInternalServerError() *GetCharactersCharacterIDIndustryJobsInternalServerError {
	return &GetCharactersCharacterIDIndustryJobsInternalServerError{}
}

/* GetCharactersCharacterIDIndustryJobsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetCharactersCharacterIDIndustryJobsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDIndustryJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDIndustryJobsServiceUnavailable creates a GetCharactersCharacterIDIndustryJobsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDIndustryJobsServiceUnavailable() *GetCharactersCharacterIDIndustryJobsServiceUnavailable {
	return &GetCharactersCharacterIDIndustryJobsServiceUnavailable{}
}

/* GetCharactersCharacterIDIndustryJobsServiceUnavailable describes a response with status code 503, with default header values.

Service unavailable
*/
type GetCharactersCharacterIDIndustryJobsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDIndustryJobsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsServiceUnavailable) GetPayload() *models.ServiceUnavailable {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDIndustryJobsGatewayTimeout creates a GetCharactersCharacterIDIndustryJobsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDIndustryJobsGatewayTimeout() *GetCharactersCharacterIDIndustryJobsGatewayTimeout {
	return &GetCharactersCharacterIDIndustryJobsGatewayTimeout{}
}

/* GetCharactersCharacterIDIndustryJobsGatewayTimeout describes a response with status code 504, with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDIndustryJobsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDIndustryJobsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/characters/{character_id}/industry/jobs/][%d] getCharactersCharacterIdIndustryJobsGatewayTimeout  %+v", 504, o.Payload)
}
func (o *GetCharactersCharacterIDIndustryJobsGatewayTimeout) GetPayload() *models.GatewayTimeout {
	return o.Payload
}

func (o *GetCharactersCharacterIDIndustryJobsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetCharactersCharacterIDIndustryJobsOKBodyItems0 get_characters_character_id_industry_jobs_200_ok
//
// 200 ok object
swagger:model GetCharactersCharacterIDIndustryJobsOKBodyItems0
*/
type GetCharactersCharacterIDIndustryJobsOKBodyItems0 struct {

	// get_characters_character_id_industry_jobs_activity_id
	//
	// Job activity ID
	// Required: true
	ActivityID *int32 `json:"activity_id"`

	// get_characters_character_id_industry_jobs_blueprint_id
	//
	// blueprint_id integer
	// Required: true
	BlueprintID *int64 `json:"blueprint_id"`

	// get_characters_character_id_industry_jobs_blueprint_location_id
	//
	// Location ID of the location from which the blueprint was installed. Normally a station ID, but can also be an asset (e.g. container) or corporation facility
	// Required: true
	BlueprintLocationID *int64 `json:"blueprint_location_id"`

	// get_characters_character_id_industry_jobs_blueprint_type_id
	//
	// blueprint_type_id integer
	// Required: true
	BlueprintTypeID *int32 `json:"blueprint_type_id"`

	// get_characters_character_id_industry_jobs_completed_character_id
	//
	// ID of the character which completed this job
	CompletedCharacterID int32 `json:"completed_character_id,omitempty"`

	// get_characters_character_id_industry_jobs_completed_date
	//
	// Date and time when this job was completed
	// Format: date-time
	CompletedDate strfmt.DateTime `json:"completed_date,omitempty"`

	// get_characters_character_id_industry_jobs_cost
	//
	// The sume of job installation fee and industry facility tax
	Cost float64 `json:"cost,omitempty"`

	// get_characters_character_id_industry_jobs_duration
	//
	// Job duration in seconds
	// Required: true
	Duration *int32 `json:"duration"`

	// get_characters_character_id_industry_jobs_end_date
	//
	// Date and time when this job finished
	// Required: true
	// Format: date-time
	EndDate *strfmt.DateTime `json:"end_date"`

	// get_characters_character_id_industry_jobs_facility_id
	//
	// ID of the facility where this job is running
	// Required: true
	FacilityID *int64 `json:"facility_id"`

	// get_characters_character_id_industry_jobs_installer_id
	//
	// ID of the character which installed this job
	// Required: true
	InstallerID *int32 `json:"installer_id"`

	// get_characters_character_id_industry_jobs_job_id
	//
	// Unique job ID
	// Required: true
	JobID *int32 `json:"job_id"`

	// get_characters_character_id_industry_jobs_licensed_runs
	//
	// Number of runs blueprint is licensed for
	LicensedRuns int32 `json:"licensed_runs,omitempty"`

	// get_characters_character_id_industry_jobs_output_location_id
	//
	// Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility
	// Required: true
	OutputLocationID *int64 `json:"output_location_id"`

	// get_characters_character_id_industry_jobs_pause_date
	//
	// Date and time when this job was paused (i.e. time when the facility where this job was installed went offline)
	// Format: date-time
	PauseDate strfmt.DateTime `json:"pause_date,omitempty"`

	// get_characters_character_id_industry_jobs_probability
	//
	// Chance of success for invention
	Probability float32 `json:"probability,omitempty"`

	// get_characters_character_id_industry_jobs_product_type_id
	//
	// Type ID of product (manufactured, copied or invented)
	ProductTypeID int32 `json:"product_type_id,omitempty"`

	// get_characters_character_id_industry_jobs_runs
	//
	// Number of runs for a manufacturing job, or number of copies to make for a blueprint copy
	// Required: true
	Runs *int32 `json:"runs"`

	// get_characters_character_id_industry_jobs_start_date
	//
	// Date and time when this job started
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"start_date"`

	// get_characters_character_id_industry_jobs_station_id
	//
	// ID of the station where industry facility is located
	// Required: true
	StationID *int64 `json:"station_id"`

	// get_characters_character_id_industry_jobs_status
	//
	// status string
	// Required: true
	// Enum: [active cancelled delivered paused ready reverted]
	Status *string `json:"status"`

	// get_characters_character_id_industry_jobs_successful_runs
	//
	// Number of successful runs for this job. Equal to runs unless this is an invention job
	SuccessfulRuns int32 `json:"successful_runs,omitempty"`
}

// Validate validates this get characters character ID industry jobs o k body items0
func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActivityID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBlueprintID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBlueprintLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBlueprintTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCompletedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInstallerID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOutputLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePauseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateActivityID(formats strfmt.Registry) error {

	if err := validate.Required("activity_id", "body", o.ActivityID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateBlueprintID(formats strfmt.Registry) error {

	if err := validate.Required("blueprint_id", "body", o.BlueprintID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateBlueprintLocationID(formats strfmt.Registry) error {

	if err := validate.Required("blueprint_location_id", "body", o.BlueprintLocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateBlueprintTypeID(formats strfmt.Registry) error {

	if err := validate.Required("blueprint_type_id", "body", o.BlueprintTypeID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateCompletedDate(formats strfmt.Registry) error {
	if swag.IsZero(o.CompletedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_date", "body", "date-time", o.CompletedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", o.Duration); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateEndDate(formats strfmt.Registry) error {

	if err := validate.Required("end_date", "body", o.EndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("end_date", "body", "date-time", o.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateFacilityID(formats strfmt.Registry) error {

	if err := validate.Required("facility_id", "body", o.FacilityID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateInstallerID(formats strfmt.Registry) error {

	if err := validate.Required("installer_id", "body", o.InstallerID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", o.JobID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateOutputLocationID(formats strfmt.Registry) error {

	if err := validate.Required("output_location_id", "body", o.OutputLocationID); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validatePauseDate(formats strfmt.Registry) error {
	if swag.IsZero(o.PauseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("pause_date", "body", "date-time", o.PauseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateRuns(formats strfmt.Registry) error {

	if err := validate.Required("runs", "body", o.Runs); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("start_date", "body", o.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("start_date", "body", "date-time", o.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("station_id", "body", o.StationID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdIndustryJobsOKBodyItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","cancelled","delivered","paused","ready","reverted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdIndustryJobsOKBodyItems0TypeStatusPropEnum = append(getCharactersCharacterIdIndustryJobsOKBodyItems0TypeStatusPropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusActive captures enum value "active"
	GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusActive string = "active"

	// GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusCancelled captures enum value "cancelled"
	GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusCancelled string = "cancelled"

	// GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusDelivered captures enum value "delivered"
	GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusDelivered string = "delivered"

	// GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusPaused captures enum value "paused"
	GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusPaused string = "paused"

	// GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusReady captures enum value "ready"
	GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusReady string = "ready"

	// GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusReverted captures enum value "reverted"
	GetCharactersCharacterIDIndustryJobsOKBodyItems0StatusReverted string = "reverted"
)

// prop value enum
func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getCharactersCharacterIdIndustryJobsOKBodyItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", o.Status); err != nil {
		return err
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get characters character ID industry jobs o k body items0 based on context it is used
func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCharactersCharacterIDIndustryJobsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDIndustryJobsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
