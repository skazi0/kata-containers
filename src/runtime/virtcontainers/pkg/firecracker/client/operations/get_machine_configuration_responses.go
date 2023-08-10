// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/pkg/firecracker/client/models"
)

// GetMachineConfigurationReader is a Reader for the GetMachineConfiguration structure.
type GetMachineConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMachineConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMachineConfigurationOK creates a GetMachineConfigurationOK with default headers values
func NewGetMachineConfigurationOK() *GetMachineConfigurationOK {
	return &GetMachineConfigurationOK{}
}

/*
GetMachineConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetMachineConfigurationOK struct {
	Payload *models.MachineConfiguration
}

// IsSuccess returns true when this get machine configuration o k response has a 2xx status code
func (o *GetMachineConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get machine configuration o k response has a 3xx status code
func (o *GetMachineConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get machine configuration o k response has a 4xx status code
func (o *GetMachineConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get machine configuration o k response has a 5xx status code
func (o *GetMachineConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get machine configuration o k response a status code equal to that given
func (o *GetMachineConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetMachineConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /machine-config][%d] getMachineConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetMachineConfigurationOK) String() string {
	return fmt.Sprintf("[GET /machine-config][%d] getMachineConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetMachineConfigurationOK) GetPayload() *models.MachineConfiguration {
	return o.Payload
}

func (o *GetMachineConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineConfigurationDefault creates a GetMachineConfigurationDefault with default headers values
func NewGetMachineConfigurationDefault(code int) *GetMachineConfigurationDefault {
	return &GetMachineConfigurationDefault{
		_statusCode: code,
	}
}

/*
GetMachineConfigurationDefault describes a response with status code -1, with default header values.

Internal server error
*/
type GetMachineConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get machine configuration default response
func (o *GetMachineConfigurationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get machine configuration default response has a 2xx status code
func (o *GetMachineConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get machine configuration default response has a 3xx status code
func (o *GetMachineConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get machine configuration default response has a 4xx status code
func (o *GetMachineConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get machine configuration default response has a 5xx status code
func (o *GetMachineConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get machine configuration default response a status code equal to that given
func (o *GetMachineConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetMachineConfigurationDefault) Error() string {
	return fmt.Sprintf("[GET /machine-config][%d] getMachineConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *GetMachineConfigurationDefault) String() string {
	return fmt.Sprintf("[GET /machine-config][%d] getMachineConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *GetMachineConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMachineConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
