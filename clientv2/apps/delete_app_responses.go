// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fnproject/fn_go/modelsv2"
)

// DeleteAppReader is a Reader for the DeleteApp structure.
type DeleteAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAppNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAppNoContent creates a DeleteAppNoContent with default headers values
func NewDeleteAppNoContent() *DeleteAppNoContent {
	return &DeleteAppNoContent{}
}

/*DeleteAppNoContent handles this case with default header values.

Application successfully deleted.
*/
type DeleteAppNoContent struct {
}

func (o *DeleteAppNoContent) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appID}][%d] deleteAppNoContent ", 204)
}

func (o *DeleteAppNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAppNotFound creates a DeleteAppNotFound with default headers values
func NewDeleteAppNotFound() *DeleteAppNotFound {
	return &DeleteAppNotFound{}
}

/*DeleteAppNotFound handles this case with default header values.

Application does not exist.
*/
type DeleteAppNotFound struct {
	Payload *modelsv2.Error
}

func (o *DeleteAppNotFound) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appID}][%d] deleteAppNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAppNotFound) GetPayload() *modelsv2.Error {
	return o.Payload
}

func (o *DeleteAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAppDefault creates a DeleteAppDefault with default headers values
func NewDeleteAppDefault(code int) *DeleteAppDefault {
	return &DeleteAppDefault{
		_statusCode: code,
	}
}

/*DeleteAppDefault handles this case with default header values.

An unexpected error occurred.
*/
type DeleteAppDefault struct {
	_statusCode int

	Payload *modelsv2.Error
}

// Code gets the status code for the delete app default response
func (o *DeleteAppDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAppDefault) Error() string {
	return fmt.Sprintf("[DELETE /apps/{appID}][%d] DeleteApp default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAppDefault) GetPayload() *modelsv2.Error {
	return o.Payload
}

func (o *DeleteAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
