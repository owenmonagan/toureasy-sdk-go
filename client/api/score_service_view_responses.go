// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owenmonagan/toureasy-sdk-go/models"
)

// ScoreServiceViewReader is a Reader for the ScoreServiceView structure.
type ScoreServiceViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScoreServiceViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScoreServiceViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScoreServiceViewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScoreServiceViewOK creates a ScoreServiceViewOK with default headers values
func NewScoreServiceViewOK() *ScoreServiceViewOK {
	return &ScoreServiceViewOK{}
}

/*ScoreServiceViewOK handles this case with default header values.

A successful response.
*/
type ScoreServiceViewOK struct {
	Payload *models.APIScoreResponse
}

func (o *ScoreServiceViewOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/score/{id}][%d] scoreServiceViewOK  %+v", 200, o.Payload)
}

func (o *ScoreServiceViewOK) GetPayload() *models.APIScoreResponse {
	return o.Payload
}

func (o *ScoreServiceViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIScoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScoreServiceViewDefault creates a ScoreServiceViewDefault with default headers values
func NewScoreServiceViewDefault(code int) *ScoreServiceViewDefault {
	return &ScoreServiceViewDefault{
		_statusCode: code,
	}
}

/*ScoreServiceViewDefault handles this case with default header values.

An unexpected error response.
*/
type ScoreServiceViewDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the score service view default response
func (o *ScoreServiceViewDefault) Code() int {
	return o._statusCode
}

func (o *ScoreServiceViewDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/score/{id}][%d] ScoreService_View default  %+v", o._statusCode, o.Payload)
}

func (o *ScoreServiceViewDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ScoreServiceViewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}