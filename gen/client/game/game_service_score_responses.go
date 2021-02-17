// Code generated by go-swagger; DO NOT EDIT.

package game

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owenmonagan/toureasy-sdk-go/gen/models"
)

// GameServiceScoreReader is a Reader for the GameServiceScore structure.
type GameServiceScoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GameServiceScoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGameServiceScoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGameServiceScoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGameServiceScoreOK creates a GameServiceScoreOK with default headers values
func NewGameServiceScoreOK() *GameServiceScoreOK {
	return &GameServiceScoreOK{}
}

/*GameServiceScoreOK handles this case with default header values.

A successful response.
*/
type GameServiceScoreOK struct {
	Payload *models.APIGames
}

func (o *GameServiceScoreOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/game/score][%d] gameServiceScoreOK  %+v", 200, o.Payload)
}

func (o *GameServiceScoreOK) GetPayload() *models.APIGames {
	return o.Payload
}

func (o *GameServiceScoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIGames)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGameServiceScoreDefault creates a GameServiceScoreDefault with default headers values
func NewGameServiceScoreDefault(code int) *GameServiceScoreDefault {
	return &GameServiceScoreDefault{
		_statusCode: code,
	}
}

/*GameServiceScoreDefault handles this case with default header values.

An unexpected error response.
*/
type GameServiceScoreDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the game service score default response
func (o *GameServiceScoreDefault) Code() int {
	return o._statusCode
}

func (o *GameServiceScoreDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/game/score][%d] GameService_Score default  %+v", o._statusCode, o.Payload)
}

func (o *GameServiceScoreDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GameServiceScoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}