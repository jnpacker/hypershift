// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2VolumescloneStartPostReader is a Reader for the PcloudV2VolumescloneStartPost structure.
type PcloudV2VolumescloneStartPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneStartPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudV2VolumescloneStartPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPcloudV2VolumescloneStartPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2VolumescloneStartPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2VolumescloneStartPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumescloneStartPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start] pcloud.v2.volumesclone.start.post", response, response.Code())
	}
}

// NewPcloudV2VolumescloneStartPostOK creates a PcloudV2VolumescloneStartPostOK with default headers values
func NewPcloudV2VolumescloneStartPostOK() *PcloudV2VolumescloneStartPostOK {
	return &PcloudV2VolumescloneStartPostOK{}
}

/*
PcloudV2VolumescloneStartPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudV2VolumescloneStartPostOK struct {
	Payload *models.VolumesClone
}

// IsSuccess returns true when this pcloud v2 volumesclone start post o k response has a 2xx status code
func (o *PcloudV2VolumescloneStartPostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 volumesclone start post o k response has a 3xx status code
func (o *PcloudV2VolumescloneStartPostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone start post o k response has a 4xx status code
func (o *PcloudV2VolumescloneStartPostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone start post o k response has a 5xx status code
func (o *PcloudV2VolumescloneStartPostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone start post o k response a status code equal to that given
func (o *PcloudV2VolumescloneStartPostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud v2 volumesclone start post o k response
func (o *PcloudV2VolumescloneStartPostOK) Code() int {
	return 200
}

func (o *PcloudV2VolumescloneStartPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostOK) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostOK  %+v", 200, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostOK) GetPayload() *models.VolumesClone {
	return o.Payload
}

func (o *PcloudV2VolumescloneStartPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneStartPostUnauthorized creates a PcloudV2VolumescloneStartPostUnauthorized with default headers values
func NewPcloudV2VolumescloneStartPostUnauthorized() *PcloudV2VolumescloneStartPostUnauthorized {
	return &PcloudV2VolumescloneStartPostUnauthorized{}
}

/*
PcloudV2VolumescloneStartPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumescloneStartPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone start post unauthorized response has a 2xx status code
func (o *PcloudV2VolumescloneStartPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone start post unauthorized response has a 3xx status code
func (o *PcloudV2VolumescloneStartPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone start post unauthorized response has a 4xx status code
func (o *PcloudV2VolumescloneStartPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone start post unauthorized response has a 5xx status code
func (o *PcloudV2VolumescloneStartPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone start post unauthorized response a status code equal to that given
func (o *PcloudV2VolumescloneStartPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud v2 volumesclone start post unauthorized response
func (o *PcloudV2VolumescloneStartPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudV2VolumescloneStartPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneStartPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneStartPostForbidden creates a PcloudV2VolumescloneStartPostForbidden with default headers values
func NewPcloudV2VolumescloneStartPostForbidden() *PcloudV2VolumescloneStartPostForbidden {
	return &PcloudV2VolumescloneStartPostForbidden{}
}

/*
PcloudV2VolumescloneStartPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2VolumescloneStartPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone start post forbidden response has a 2xx status code
func (o *PcloudV2VolumescloneStartPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone start post forbidden response has a 3xx status code
func (o *PcloudV2VolumescloneStartPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone start post forbidden response has a 4xx status code
func (o *PcloudV2VolumescloneStartPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone start post forbidden response has a 5xx status code
func (o *PcloudV2VolumescloneStartPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone start post forbidden response a status code equal to that given
func (o *PcloudV2VolumescloneStartPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud v2 volumesclone start post forbidden response
func (o *PcloudV2VolumescloneStartPostForbidden) Code() int {
	return 403
}

func (o *PcloudV2VolumescloneStartPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneStartPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneStartPostNotFound creates a PcloudV2VolumescloneStartPostNotFound with default headers values
func NewPcloudV2VolumescloneStartPostNotFound() *PcloudV2VolumescloneStartPostNotFound {
	return &PcloudV2VolumescloneStartPostNotFound{}
}

/*
PcloudV2VolumescloneStartPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2VolumescloneStartPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone start post not found response has a 2xx status code
func (o *PcloudV2VolumescloneStartPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone start post not found response has a 3xx status code
func (o *PcloudV2VolumescloneStartPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone start post not found response has a 4xx status code
func (o *PcloudV2VolumescloneStartPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone start post not found response has a 5xx status code
func (o *PcloudV2VolumescloneStartPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone start post not found response a status code equal to that given
func (o *PcloudV2VolumescloneStartPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud v2 volumesclone start post not found response
func (o *PcloudV2VolumescloneStartPostNotFound) Code() int {
	return 404
}

func (o *PcloudV2VolumescloneStartPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostNotFound) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostNotFound  %+v", 404, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneStartPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneStartPostInternalServerError creates a PcloudV2VolumescloneStartPostInternalServerError with default headers values
func NewPcloudV2VolumescloneStartPostInternalServerError() *PcloudV2VolumescloneStartPostInternalServerError {
	return &PcloudV2VolumescloneStartPostInternalServerError{}
}

/*
PcloudV2VolumescloneStartPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneStartPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone start post internal server error response has a 2xx status code
func (o *PcloudV2VolumescloneStartPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone start post internal server error response has a 3xx status code
func (o *PcloudV2VolumescloneStartPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone start post internal server error response has a 4xx status code
func (o *PcloudV2VolumescloneStartPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone start post internal server error response has a 5xx status code
func (o *PcloudV2VolumescloneStartPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 volumesclone start post internal server error response a status code equal to that given
func (o *PcloudV2VolumescloneStartPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud v2 volumesclone start post internal server error response
func (o *PcloudV2VolumescloneStartPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudV2VolumescloneStartPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/start][%d] pcloudV2VolumescloneStartPostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumescloneStartPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneStartPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
