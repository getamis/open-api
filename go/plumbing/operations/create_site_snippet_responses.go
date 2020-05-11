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

// CreateSiteSnippetReader is a Reader for the CreateSiteSnippet structure.
type CreateSiteSnippetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSiteSnippetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSiteSnippetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateSiteSnippetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSiteSnippetCreated creates a CreateSiteSnippetCreated with default headers values
func NewCreateSiteSnippetCreated() *CreateSiteSnippetCreated {
	return &CreateSiteSnippetCreated{}
}

/*CreateSiteSnippetCreated handles this case with default header values.

OK
*/
type CreateSiteSnippetCreated struct {
	Payload *models.Snippet
}

func (o *CreateSiteSnippetCreated) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/snippets][%d] createSiteSnippetCreated  %+v", 201, o.Payload)
}

func (o *CreateSiteSnippetCreated) GetPayload() *models.Snippet {
	return o.Payload
}

func (o *CreateSiteSnippetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snippet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteSnippetDefault creates a CreateSiteSnippetDefault with default headers values
func NewCreateSiteSnippetDefault(code int) *CreateSiteSnippetDefault {
	return &CreateSiteSnippetDefault{
		_statusCode: code,
	}
}

/*CreateSiteSnippetDefault handles this case with default header values.

error
*/
type CreateSiteSnippetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create site snippet default response
func (o *CreateSiteSnippetDefault) Code() int {
	return o._statusCode
}

func (o *CreateSiteSnippetDefault) Error() string {
	return fmt.Sprintf("[POST /sites/{site_id}/snippets][%d] createSiteSnippet default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteSnippetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSiteSnippetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
