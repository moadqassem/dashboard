// Code generated by go-swagger; DO NOT EDIT.

package mlaadminsetting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// UpdateMLAAdminSettingReader is a Reader for the UpdateMLAAdminSetting structure.
type UpdateMLAAdminSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMLAAdminSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMLAAdminSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMLAAdminSettingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMLAAdminSettingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateMLAAdminSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMLAAdminSettingOK creates a UpdateMLAAdminSettingOK with default headers values
func NewUpdateMLAAdminSettingOK() *UpdateMLAAdminSettingOK {
	return &UpdateMLAAdminSettingOK{}
}

/*
UpdateMLAAdminSettingOK describes a response with status code 200, with default header values.

MLAAdminSetting
*/
type UpdateMLAAdminSettingOK struct {
	Payload *models.MLAAdminSetting
}

// IsSuccess returns true when this update m l a admin setting o k response has a 2xx status code
func (o *UpdateMLAAdminSettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update m l a admin setting o k response has a 3xx status code
func (o *UpdateMLAAdminSettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m l a admin setting o k response has a 4xx status code
func (o *UpdateMLAAdminSettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update m l a admin setting o k response has a 5xx status code
func (o *UpdateMLAAdminSettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update m l a admin setting o k response a status code equal to that given
func (o *UpdateMLAAdminSettingOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateMLAAdminSettingOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSettingOK  %+v", 200, o.Payload)
}

func (o *UpdateMLAAdminSettingOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSettingOK  %+v", 200, o.Payload)
}

func (o *UpdateMLAAdminSettingOK) GetPayload() *models.MLAAdminSetting {
	return o.Payload
}

func (o *UpdateMLAAdminSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MLAAdminSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMLAAdminSettingUnauthorized creates a UpdateMLAAdminSettingUnauthorized with default headers values
func NewUpdateMLAAdminSettingUnauthorized() *UpdateMLAAdminSettingUnauthorized {
	return &UpdateMLAAdminSettingUnauthorized{}
}

/*
UpdateMLAAdminSettingUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpdateMLAAdminSettingUnauthorized struct {
}

// IsSuccess returns true when this update m l a admin setting unauthorized response has a 2xx status code
func (o *UpdateMLAAdminSettingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m l a admin setting unauthorized response has a 3xx status code
func (o *UpdateMLAAdminSettingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m l a admin setting unauthorized response has a 4xx status code
func (o *UpdateMLAAdminSettingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m l a admin setting unauthorized response has a 5xx status code
func (o *UpdateMLAAdminSettingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update m l a admin setting unauthorized response a status code equal to that given
func (o *UpdateMLAAdminSettingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateMLAAdminSettingUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSettingUnauthorized ", 401)
}

func (o *UpdateMLAAdminSettingUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSettingUnauthorized ", 401)
}

func (o *UpdateMLAAdminSettingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMLAAdminSettingForbidden creates a UpdateMLAAdminSettingForbidden with default headers values
func NewUpdateMLAAdminSettingForbidden() *UpdateMLAAdminSettingForbidden {
	return &UpdateMLAAdminSettingForbidden{}
}

/*
UpdateMLAAdminSettingForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpdateMLAAdminSettingForbidden struct {
}

// IsSuccess returns true when this update m l a admin setting forbidden response has a 2xx status code
func (o *UpdateMLAAdminSettingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update m l a admin setting forbidden response has a 3xx status code
func (o *UpdateMLAAdminSettingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update m l a admin setting forbidden response has a 4xx status code
func (o *UpdateMLAAdminSettingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update m l a admin setting forbidden response has a 5xx status code
func (o *UpdateMLAAdminSettingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update m l a admin setting forbidden response a status code equal to that given
func (o *UpdateMLAAdminSettingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateMLAAdminSettingForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSettingForbidden ", 403)
}

func (o *UpdateMLAAdminSettingForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSettingForbidden ", 403)
}

func (o *UpdateMLAAdminSettingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMLAAdminSettingDefault creates a UpdateMLAAdminSettingDefault with default headers values
func NewUpdateMLAAdminSettingDefault(code int) *UpdateMLAAdminSettingDefault {
	return &UpdateMLAAdminSettingDefault{
		_statusCode: code,
	}
}

/*
UpdateMLAAdminSettingDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateMLAAdminSettingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update m l a admin setting default response
func (o *UpdateMLAAdminSettingDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this update m l a admin setting default response has a 2xx status code
func (o *UpdateMLAAdminSettingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update m l a admin setting default response has a 3xx status code
func (o *UpdateMLAAdminSettingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update m l a admin setting default response has a 4xx status code
func (o *UpdateMLAAdminSettingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update m l a admin setting default response has a 5xx status code
func (o *UpdateMLAAdminSettingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update m l a admin setting default response a status code equal to that given
func (o *UpdateMLAAdminSettingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UpdateMLAAdminSettingDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSetting default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMLAAdminSettingDefault) String() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/mlaadminsetting][%d] updateMLAAdminSetting default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMLAAdminSettingDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateMLAAdminSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}