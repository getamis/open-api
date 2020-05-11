// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// GetSiteReader is a Reader for the GetSite structure.
type GetSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteOK creates a GetSiteOK with default headers values
func NewGetSiteOK() *GetSiteOK {
	return &GetSiteOK{}
}

/*GetSiteOK handles this case with default header values.

OK
*/
type GetSiteOK struct {
	Payload *models.Site
}

func (o *GetSiteOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}][%d] getSiteOK  %+v", 200, o.Payload)
}

func (o *GetSiteOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *GetSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteDefault creates a GetSiteDefault with default headers values
func NewGetSiteDefault(code int) *GetSiteDefault {
	return &GetSiteDefault{
		_statusCode: code,
	}
}

/*GetSiteDefault handles this case with default header values.

error
*/
type GetSiteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get site default response
func (o *GetSiteDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}][%d] getSite default  %+v", o._statusCode, o.Payload)
}

func (o *GetSiteDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
