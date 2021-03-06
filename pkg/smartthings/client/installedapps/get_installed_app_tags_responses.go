// Code generated by go-swagger; DO NOT EDIT.

package installedapps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// GetInstalledAppTagsReader is a Reader for the GetInstalledAppTags structure.
type GetInstalledAppTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstalledAppTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstalledAppTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInstalledAppTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetInstalledAppTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetInstalledAppTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetInstalledAppTagsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewGetInstalledAppTagsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInstalledAppTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInstalledAppTagsOK creates a GetInstalledAppTagsOK with default headers values
func NewGetInstalledAppTagsOK() *GetInstalledAppTagsOK {
	return &GetInstalledAppTagsOK{}
}

/*GetInstalledAppTagsOK handles this case with default header values.

A response containing an InstalledApp's tags.
*/
type GetInstalledAppTagsOK struct {
	Payload *models.GetTagsResponse
}

func (o *GetInstalledAppTagsOK) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/tags][%d] getInstalledAppTagsOK  %+v", 200, o.Payload)
}

func (o *GetInstalledAppTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstalledAppTagsBadRequest creates a GetInstalledAppTagsBadRequest with default headers values
func NewGetInstalledAppTagsBadRequest() *GetInstalledAppTagsBadRequest {
	return &GetInstalledAppTagsBadRequest{}
}

/*GetInstalledAppTagsBadRequest handles this case with default header values.

Bad request
*/
type GetInstalledAppTagsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetInstalledAppTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/tags][%d] getInstalledAppTagsBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstalledAppTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstalledAppTagsUnauthorized creates a GetInstalledAppTagsUnauthorized with default headers values
func NewGetInstalledAppTagsUnauthorized() *GetInstalledAppTagsUnauthorized {
	return &GetInstalledAppTagsUnauthorized{}
}

/*GetInstalledAppTagsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetInstalledAppTagsUnauthorized struct {
}

func (o *GetInstalledAppTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/tags][%d] getInstalledAppTagsUnauthorized ", 401)
}

func (o *GetInstalledAppTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstalledAppTagsForbidden creates a GetInstalledAppTagsForbidden with default headers values
func NewGetInstalledAppTagsForbidden() *GetInstalledAppTagsForbidden {
	return &GetInstalledAppTagsForbidden{}
}

/*GetInstalledAppTagsForbidden handles this case with default header values.

Forbidden
*/
type GetInstalledAppTagsForbidden struct {
}

func (o *GetInstalledAppTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/tags][%d] getInstalledAppTagsForbidden ", 403)
}

func (o *GetInstalledAppTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstalledAppTagsUnprocessableEntity creates a GetInstalledAppTagsUnprocessableEntity with default headers values
func NewGetInstalledAppTagsUnprocessableEntity() *GetInstalledAppTagsUnprocessableEntity {
	return &GetInstalledAppTagsUnprocessableEntity{}
}

/*GetInstalledAppTagsUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type GetInstalledAppTagsUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *GetInstalledAppTagsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/tags][%d] getInstalledAppTagsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetInstalledAppTagsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstalledAppTagsTooManyRequests creates a GetInstalledAppTagsTooManyRequests with default headers values
func NewGetInstalledAppTagsTooManyRequests() *GetInstalledAppTagsTooManyRequests {
	return &GetInstalledAppTagsTooManyRequests{}
}

/*GetInstalledAppTagsTooManyRequests handles this case with default header values.

Too many requests
*/
type GetInstalledAppTagsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetInstalledAppTagsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/tags][%d] getInstalledAppTagsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInstalledAppTagsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstalledAppTagsDefault creates a GetInstalledAppTagsDefault with default headers values
func NewGetInstalledAppTagsDefault(code int) *GetInstalledAppTagsDefault {
	return &GetInstalledAppTagsDefault{
		_statusCode: code,
	}
}

/*GetInstalledAppTagsDefault handles this case with default header values.

Unexpected error
*/
type GetInstalledAppTagsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get installed app tags default response
func (o *GetInstalledAppTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetInstalledAppTagsDefault) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/tags][%d] getInstalledAppTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstalledAppTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
