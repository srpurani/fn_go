// Code generated by go-swagger; DO NOT EDIT.

package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	modelsv2 "github.com/fnproject/fn_go/modelsv2"
)

// GetCallsCallIDReader is a Reader for the GetCallsCallID structure.
type GetCallsCallIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCallsCallIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCallsCallIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCallsCallIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCallsCallIDOK creates a GetCallsCallIDOK with default headers values
func NewGetCallsCallIDOK() *GetCallsCallIDOK {
	return &GetCallsCallIDOK{}
}

/*GetCallsCallIDOK handles this case with default header values.

Call found
*/
type GetCallsCallIDOK struct {
	Payload *modelsv2.Call
}

func (o *GetCallsCallIDOK) Error() string {
	return fmt.Sprintf("[GET /calls/{callID}][%d] getCallsCallIdOK  %+v", 200, o.Payload)
}

func (o *GetCallsCallIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Call)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCallsCallIDNotFound creates a GetCallsCallIDNotFound with default headers values
func NewGetCallsCallIDNotFound() *GetCallsCallIDNotFound {
	return &GetCallsCallIDNotFound{}
}

/*GetCallsCallIDNotFound handles this case with default header values.

Call not found.
*/
type GetCallsCallIDNotFound struct {
	Payload *modelsv2.Error
}

func (o *GetCallsCallIDNotFound) Error() string {
	return fmt.Sprintf("[GET /calls/{callID}][%d] getCallsCallIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCallsCallIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modelsv2.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}