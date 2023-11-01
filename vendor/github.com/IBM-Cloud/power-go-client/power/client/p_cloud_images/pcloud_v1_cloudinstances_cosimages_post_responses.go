// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV1CloudinstancesCosimagesPostReader is a Reader for the PcloudV1CloudinstancesCosimagesPost structure.
type PcloudV1CloudinstancesCosimagesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV1CloudinstancesCosimagesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudV1CloudinstancesCosimagesPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV1CloudinstancesCosimagesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV1CloudinstancesCosimagesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV1CloudinstancesCosimagesPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudV1CloudinstancesCosimagesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudV1CloudinstancesCosimagesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV1CloudinstancesCosimagesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images] pcloud.v1.cloudinstances.cosimages.post", response, response.Code())
	}
}

// NewPcloudV1CloudinstancesCosimagesPostAccepted creates a PcloudV1CloudinstancesCosimagesPostAccepted with default headers values
func NewPcloudV1CloudinstancesCosimagesPostAccepted() *PcloudV1CloudinstancesCosimagesPostAccepted {
	return &PcloudV1CloudinstancesCosimagesPostAccepted{}
}

/*
PcloudV1CloudinstancesCosimagesPostAccepted describes a response with status code 202, with default header values.

Accepted, cos-image import successfully added to the jobs queue
*/
type PcloudV1CloudinstancesCosimagesPostAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud v1 cloudinstances cosimages post accepted response has a 2xx status code
func (o *PcloudV1CloudinstancesCosimagesPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v1 cloudinstances cosimages post accepted response has a 3xx status code
func (o *PcloudV1CloudinstancesCosimagesPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v1 cloudinstances cosimages post accepted response has a 4xx status code
func (o *PcloudV1CloudinstancesCosimagesPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v1 cloudinstances cosimages post accepted response has a 5xx status code
func (o *PcloudV1CloudinstancesCosimagesPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v1 cloudinstances cosimages post accepted response a status code equal to that given
func (o *PcloudV1CloudinstancesCosimagesPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud v1 cloudinstances cosimages post accepted response
func (o *PcloudV1CloudinstancesCosimagesPostAccepted) Code() int {
	return 202
}

func (o *PcloudV1CloudinstancesCosimagesPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudV1CloudinstancesCosimagesPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV1CloudinstancesCosimagesPostBadRequest creates a PcloudV1CloudinstancesCosimagesPostBadRequest with default headers values
func NewPcloudV1CloudinstancesCosimagesPostBadRequest() *PcloudV1CloudinstancesCosimagesPostBadRequest {
	return &PcloudV1CloudinstancesCosimagesPostBadRequest{}
}

/*
PcloudV1CloudinstancesCosimagesPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV1CloudinstancesCosimagesPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v1 cloudinstances cosimages post bad request response has a 2xx status code
func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v1 cloudinstances cosimages post bad request response has a 3xx status code
func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v1 cloudinstances cosimages post bad request response has a 4xx status code
func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v1 cloudinstances cosimages post bad request response has a 5xx status code
func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v1 cloudinstances cosimages post bad request response a status code equal to that given
func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud v1 cloudinstances cosimages post bad request response
func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) Code() int {
	return 400
}

func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV1CloudinstancesCosimagesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV1CloudinstancesCosimagesPostUnauthorized creates a PcloudV1CloudinstancesCosimagesPostUnauthorized with default headers values
func NewPcloudV1CloudinstancesCosimagesPostUnauthorized() *PcloudV1CloudinstancesCosimagesPostUnauthorized {
	return &PcloudV1CloudinstancesCosimagesPostUnauthorized{}
}

/*
PcloudV1CloudinstancesCosimagesPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV1CloudinstancesCosimagesPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v1 cloudinstances cosimages post unauthorized response has a 2xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v1 cloudinstances cosimages post unauthorized response has a 3xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v1 cloudinstances cosimages post unauthorized response has a 4xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v1 cloudinstances cosimages post unauthorized response has a 5xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v1 cloudinstances cosimages post unauthorized response a status code equal to that given
func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud v1 cloudinstances cosimages post unauthorized response
func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV1CloudinstancesCosimagesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV1CloudinstancesCosimagesPostForbidden creates a PcloudV1CloudinstancesCosimagesPostForbidden with default headers values
func NewPcloudV1CloudinstancesCosimagesPostForbidden() *PcloudV1CloudinstancesCosimagesPostForbidden {
	return &PcloudV1CloudinstancesCosimagesPostForbidden{}
}

/*
PcloudV1CloudinstancesCosimagesPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV1CloudinstancesCosimagesPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v1 cloudinstances cosimages post forbidden response has a 2xx status code
func (o *PcloudV1CloudinstancesCosimagesPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v1 cloudinstances cosimages post forbidden response has a 3xx status code
func (o *PcloudV1CloudinstancesCosimagesPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v1 cloudinstances cosimages post forbidden response has a 4xx status code
func (o *PcloudV1CloudinstancesCosimagesPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v1 cloudinstances cosimages post forbidden response has a 5xx status code
func (o *PcloudV1CloudinstancesCosimagesPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v1 cloudinstances cosimages post forbidden response a status code equal to that given
func (o *PcloudV1CloudinstancesCosimagesPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud v1 cloudinstances cosimages post forbidden response
func (o *PcloudV1CloudinstancesCosimagesPostForbidden) Code() int {
	return 403
}

func (o *PcloudV1CloudinstancesCosimagesPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV1CloudinstancesCosimagesPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV1CloudinstancesCosimagesPostConflict creates a PcloudV1CloudinstancesCosimagesPostConflict with default headers values
func NewPcloudV1CloudinstancesCosimagesPostConflict() *PcloudV1CloudinstancesCosimagesPostConflict {
	return &PcloudV1CloudinstancesCosimagesPostConflict{}
}

/*
PcloudV1CloudinstancesCosimagesPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudV1CloudinstancesCosimagesPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v1 cloudinstances cosimages post conflict response has a 2xx status code
func (o *PcloudV1CloudinstancesCosimagesPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v1 cloudinstances cosimages post conflict response has a 3xx status code
func (o *PcloudV1CloudinstancesCosimagesPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v1 cloudinstances cosimages post conflict response has a 4xx status code
func (o *PcloudV1CloudinstancesCosimagesPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v1 cloudinstances cosimages post conflict response has a 5xx status code
func (o *PcloudV1CloudinstancesCosimagesPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v1 cloudinstances cosimages post conflict response a status code equal to that given
func (o *PcloudV1CloudinstancesCosimagesPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud v1 cloudinstances cosimages post conflict response
func (o *PcloudV1CloudinstancesCosimagesPostConflict) Code() int {
	return 409
}

func (o *PcloudV1CloudinstancesCosimagesPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostConflict) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostConflict  %+v", 409, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV1CloudinstancesCosimagesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV1CloudinstancesCosimagesPostUnprocessableEntity creates a PcloudV1CloudinstancesCosimagesPostUnprocessableEntity with default headers values
func NewPcloudV1CloudinstancesCosimagesPostUnprocessableEntity() *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity {
	return &PcloudV1CloudinstancesCosimagesPostUnprocessableEntity{}
}

/*
PcloudV1CloudinstancesCosimagesPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudV1CloudinstancesCosimagesPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v1 cloudinstances cosimages post unprocessable entity response has a 2xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v1 cloudinstances cosimages post unprocessable entity response has a 3xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v1 cloudinstances cosimages post unprocessable entity response has a 4xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v1 cloudinstances cosimages post unprocessable entity response has a 5xx status code
func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v1 cloudinstances cosimages post unprocessable entity response a status code equal to that given
func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud v1 cloudinstances cosimages post unprocessable entity response
func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV1CloudinstancesCosimagesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV1CloudinstancesCosimagesPostInternalServerError creates a PcloudV1CloudinstancesCosimagesPostInternalServerError with default headers values
func NewPcloudV1CloudinstancesCosimagesPostInternalServerError() *PcloudV1CloudinstancesCosimagesPostInternalServerError {
	return &PcloudV1CloudinstancesCosimagesPostInternalServerError{}
}

/*
PcloudV1CloudinstancesCosimagesPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV1CloudinstancesCosimagesPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v1 cloudinstances cosimages post internal server error response has a 2xx status code
func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v1 cloudinstances cosimages post internal server error response has a 3xx status code
func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v1 cloudinstances cosimages post internal server error response has a 4xx status code
func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v1 cloudinstances cosimages post internal server error response has a 5xx status code
func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v1 cloudinstances cosimages post internal server error response a status code equal to that given
func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud v1 cloudinstances cosimages post internal server error response
func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/cos-images][%d] pcloudV1CloudinstancesCosimagesPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV1CloudinstancesCosimagesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
