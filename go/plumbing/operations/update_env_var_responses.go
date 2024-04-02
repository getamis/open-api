// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// UpdateEnvVarReader is a Reader for the UpdateEnvVar structure.
type UpdateEnvVarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEnvVarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEnvVarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateEnvVarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateEnvVarOK creates a UpdateEnvVarOK with default headers values
func NewUpdateEnvVarOK() *UpdateEnvVarOK {
	return &UpdateEnvVarOK{}
}

/*
UpdateEnvVarOK handles this case with default header values.

OK
*/
type UpdateEnvVarOK struct {
	Payload *models.EnvVar
}

func (o *UpdateEnvVarOK) Error() string {
	return fmt.Sprintf("[PUT /accounts/{account_id}/env/{key}][%d] updateEnvVarOK  %+v", 200, o.Payload)
}

func (o *UpdateEnvVarOK) GetPayload() *models.EnvVar {
	return o.Payload
}

func (o *UpdateEnvVarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvVar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEnvVarDefault creates a UpdateEnvVarDefault with default headers values
func NewUpdateEnvVarDefault(code int) *UpdateEnvVarDefault {
	return &UpdateEnvVarDefault{
		_statusCode: code,
	}
}

/*
UpdateEnvVarDefault handles this case with default header values.

error
*/
type UpdateEnvVarDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update env var default response
func (o *UpdateEnvVarDefault) Code() int {
	return o._statusCode
}

func (o *UpdateEnvVarDefault) Error() string {
	return fmt.Sprintf("[PUT /accounts/{account_id}/env/{key}][%d] updateEnvVar default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateEnvVarDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateEnvVarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
