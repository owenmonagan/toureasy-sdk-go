// Code generated by go-swagger; DO NOT EDIT.

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/owenmonagan/toureasy-sdk-go/gen/models"
)

// TournamentServiceDeleteMixin1Reader is a Reader for the TournamentServiceDeleteMixin1 structure.
type TournamentServiceDeleteMixin1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TournamentServiceDeleteMixin1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTournamentServiceDeleteMixin1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTournamentServiceDeleteMixin1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTournamentServiceDeleteMixin1OK creates a TournamentServiceDeleteMixin1OK with default headers values
func NewTournamentServiceDeleteMixin1OK() *TournamentServiceDeleteMixin1OK {
	return &TournamentServiceDeleteMixin1OK{}
}

/*TournamentServiceDeleteMixin1OK handles this case with default header values.

A successful response.
*/
type TournamentServiceDeleteMixin1OK struct {
	Payload models.APIDeleteResponse
}

func (o *TournamentServiceDeleteMixin1OK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/entries/{uuid}][%d] tournamentServiceDeleteMixin1OK  %+v", 200, o.Payload)
}

func (o *TournamentServiceDeleteMixin1OK) GetPayload() models.APIDeleteResponse {
	return o.Payload
}

func (o *TournamentServiceDeleteMixin1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTournamentServiceDeleteMixin1Default creates a TournamentServiceDeleteMixin1Default with default headers values
func NewTournamentServiceDeleteMixin1Default(code int) *TournamentServiceDeleteMixin1Default {
	return &TournamentServiceDeleteMixin1Default{
		_statusCode: code,
	}
}

/*TournamentServiceDeleteMixin1Default handles this case with default header values.

An unexpected error response.
*/
type TournamentServiceDeleteMixin1Default struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the tournament service delete mixin1 default response
func (o *TournamentServiceDeleteMixin1Default) Code() int {
	return o._statusCode
}

func (o *TournamentServiceDeleteMixin1Default) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/entries/{uuid}][%d] TournamentService_DeleteMixin1 default  %+v", o._statusCode, o.Payload)
}

func (o *TournamentServiceDeleteMixin1Default) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *TournamentServiceDeleteMixin1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}