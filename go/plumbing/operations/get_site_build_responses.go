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

// GetSiteBuildReader is a Reader for the GetSiteBuild structure.
type GetSiteBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSiteBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteBuildOK creates a GetSiteBuildOK with default headers values
func NewGetSiteBuildOK() *GetSiteBuildOK {
	return &GetSiteBuildOK{}
}

/*
GetSiteBuildOK handles this case with default header values.

OK
*/
type GetSiteBuildOK struct {
	Payload *models.Build
}

func (o *GetSiteBuildOK) Error() string {
	return fmt.Sprintf("[GET /builds/{build_id}][%d] getSiteBuildOK  %+v", 200, o.Payload)
}

func (o *GetSiteBuildOK) GetPayload() *models.Build {
	return o.Payload
}

func (o *GetSiteBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Build)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteBuildDefault creates a GetSiteBuildDefault with default headers values
func NewGetSiteBuildDefault(code int) *GetSiteBuildDefault {
	return &GetSiteBuildDefault{
		_statusCode: code,
	}
}

/*
GetSiteBuildDefault handles this case with default header values.

error
*/
type GetSiteBuildDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get site build default response
func (o *GetSiteBuildDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteBuildDefault) Error() string {
	return fmt.Sprintf("[GET /builds/{build_id}][%d] getSiteBuild default  %+v", o._statusCode, o.Payload)
}

func (o *GetSiteBuildDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSiteBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
