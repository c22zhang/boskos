// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudTenantsSshkeysDeleteReader is a Reader for the PcloudTenantsSshkeysDelete structure.
type PcloudTenantsSshkeysDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsSshkeysDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsSshkeysDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsSshkeysDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsSshkeysDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudTenantsSshkeysDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsSshkeysDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}] pcloud.tenants.sshkeys.delete", response, response.Code())
	}
}

// NewPcloudTenantsSshkeysDeleteOK creates a PcloudTenantsSshkeysDeleteOK with default headers values
func NewPcloudTenantsSshkeysDeleteOK() *PcloudTenantsSshkeysDeleteOK {
	return &PcloudTenantsSshkeysDeleteOK{}
}

/*
PcloudTenantsSshkeysDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsSshkeysDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud tenants sshkeys delete o k response has a 2xx status code
func (o *PcloudTenantsSshkeysDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud tenants sshkeys delete o k response has a 3xx status code
func (o *PcloudTenantsSshkeysDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys delete o k response has a 4xx status code
func (o *PcloudTenantsSshkeysDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants sshkeys delete o k response has a 5xx status code
func (o *PcloudTenantsSshkeysDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys delete o k response a status code equal to that given
func (o *PcloudTenantsSshkeysDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud tenants sshkeys delete o k response
func (o *PcloudTenantsSshkeysDeleteOK) Code() int {
	return 200
}

func (o *PcloudTenantsSshkeysDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteBadRequest creates a PcloudTenantsSshkeysDeleteBadRequest with default headers values
func NewPcloudTenantsSshkeysDeleteBadRequest() *PcloudTenantsSshkeysDeleteBadRequest {
	return &PcloudTenantsSshkeysDeleteBadRequest{}
}

/*
PcloudTenantsSshkeysDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsSshkeysDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys delete bad request response has a 2xx status code
func (o *PcloudTenantsSshkeysDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys delete bad request response has a 3xx status code
func (o *PcloudTenantsSshkeysDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys delete bad request response has a 4xx status code
func (o *PcloudTenantsSshkeysDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys delete bad request response has a 5xx status code
func (o *PcloudTenantsSshkeysDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys delete bad request response a status code equal to that given
func (o *PcloudTenantsSshkeysDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud tenants sshkeys delete bad request response
func (o *PcloudTenantsSshkeysDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudTenantsSshkeysDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteUnauthorized creates a PcloudTenantsSshkeysDeleteUnauthorized with default headers values
func NewPcloudTenantsSshkeysDeleteUnauthorized() *PcloudTenantsSshkeysDeleteUnauthorized {
	return &PcloudTenantsSshkeysDeleteUnauthorized{}
}

/*
PcloudTenantsSshkeysDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsSshkeysDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys delete unauthorized response has a 2xx status code
func (o *PcloudTenantsSshkeysDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys delete unauthorized response has a 3xx status code
func (o *PcloudTenantsSshkeysDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys delete unauthorized response has a 4xx status code
func (o *PcloudTenantsSshkeysDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys delete unauthorized response has a 5xx status code
func (o *PcloudTenantsSshkeysDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys delete unauthorized response a status code equal to that given
func (o *PcloudTenantsSshkeysDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud tenants sshkeys delete unauthorized response
func (o *PcloudTenantsSshkeysDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudTenantsSshkeysDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteGone creates a PcloudTenantsSshkeysDeleteGone with default headers values
func NewPcloudTenantsSshkeysDeleteGone() *PcloudTenantsSshkeysDeleteGone {
	return &PcloudTenantsSshkeysDeleteGone{}
}

/*
PcloudTenantsSshkeysDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudTenantsSshkeysDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys delete gone response has a 2xx status code
func (o *PcloudTenantsSshkeysDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys delete gone response has a 3xx status code
func (o *PcloudTenantsSshkeysDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys delete gone response has a 4xx status code
func (o *PcloudTenantsSshkeysDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tenants sshkeys delete gone response has a 5xx status code
func (o *PcloudTenantsSshkeysDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tenants sshkeys delete gone response a status code equal to that given
func (o *PcloudTenantsSshkeysDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud tenants sshkeys delete gone response
func (o *PcloudTenantsSshkeysDeleteGone) Code() int {
	return 410
}

func (o *PcloudTenantsSshkeysDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteGone) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteGone  %+v", 410, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteInternalServerError creates a PcloudTenantsSshkeysDeleteInternalServerError with default headers values
func NewPcloudTenantsSshkeysDeleteInternalServerError() *PcloudTenantsSshkeysDeleteInternalServerError {
	return &PcloudTenantsSshkeysDeleteInternalServerError{}
}

/*
PcloudTenantsSshkeysDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsSshkeysDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tenants sshkeys delete internal server error response has a 2xx status code
func (o *PcloudTenantsSshkeysDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tenants sshkeys delete internal server error response has a 3xx status code
func (o *PcloudTenantsSshkeysDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tenants sshkeys delete internal server error response has a 4xx status code
func (o *PcloudTenantsSshkeysDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tenants sshkeys delete internal server error response has a 5xx status code
func (o *PcloudTenantsSshkeysDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud tenants sshkeys delete internal server error response a status code equal to that given
func (o *PcloudTenantsSshkeysDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud tenants sshkeys delete internal server error response
func (o *PcloudTenantsSshkeysDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudTenantsSshkeysDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTenantsSshkeysDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
