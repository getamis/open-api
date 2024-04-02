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

// GetSiteAssetPublicSignatureReader is a Reader for the GetSiteAssetPublicSignature structure.
type GetSiteAssetPublicSignatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteAssetPublicSignatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteAssetPublicSignatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSiteAssetPublicSignatureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteAssetPublicSignatureOK creates a GetSiteAssetPublicSignatureOK with default headers values
func NewGetSiteAssetPublicSignatureOK() *GetSiteAssetPublicSignatureOK {
	return &GetSiteAssetPublicSignatureOK{}
}

/*
GetSiteAssetPublicSignatureOK handles this case with default header values.

OK
*/
type GetSiteAssetPublicSignatureOK struct {
	Payload *models.AssetPublicSignature
}

func (o *GetSiteAssetPublicSignatureOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/assets/{asset_id}/public_signature][%d] getSiteAssetPublicSignatureOK  %+v", 200, o.Payload)
}

func (o *GetSiteAssetPublicSignatureOK) GetPayload() *models.AssetPublicSignature {
	return o.Payload
}

func (o *GetSiteAssetPublicSignatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssetPublicSignature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteAssetPublicSignatureDefault creates a GetSiteAssetPublicSignatureDefault with default headers values
func NewGetSiteAssetPublicSignatureDefault(code int) *GetSiteAssetPublicSignatureDefault {
	return &GetSiteAssetPublicSignatureDefault{
		_statusCode: code,
	}
}

/*
GetSiteAssetPublicSignatureDefault handles this case with default header values.

error
*/
type GetSiteAssetPublicSignatureDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get site asset public signature default response
func (o *GetSiteAssetPublicSignatureDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteAssetPublicSignatureDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/assets/{asset_id}/public_signature][%d] getSiteAssetPublicSignature default  %+v", o._statusCode, o.Payload)
}

func (o *GetSiteAssetPublicSignatureDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSiteAssetPublicSignatureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
