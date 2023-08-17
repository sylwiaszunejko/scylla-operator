// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v2/models"
)

// FindConfigCompactionRowsCountWarningThresholdReader is a Reader for the FindConfigCompactionRowsCountWarningThreshold structure.
type FindConfigCompactionRowsCountWarningThresholdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCompactionRowsCountWarningThresholdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCompactionRowsCountWarningThresholdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCompactionRowsCountWarningThresholdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCompactionRowsCountWarningThresholdOK creates a FindConfigCompactionRowsCountWarningThresholdOK with default headers values
func NewFindConfigCompactionRowsCountWarningThresholdOK() *FindConfigCompactionRowsCountWarningThresholdOK {
	return &FindConfigCompactionRowsCountWarningThresholdOK{}
}

/*
FindConfigCompactionRowsCountWarningThresholdOK handles this case with default header values.

Config value
*/
type FindConfigCompactionRowsCountWarningThresholdOK struct {
	Payload int64
}

func (o *FindConfigCompactionRowsCountWarningThresholdOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCompactionRowsCountWarningThresholdOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCompactionRowsCountWarningThresholdDefault creates a FindConfigCompactionRowsCountWarningThresholdDefault with default headers values
func NewFindConfigCompactionRowsCountWarningThresholdDefault(code int) *FindConfigCompactionRowsCountWarningThresholdDefault {
	return &FindConfigCompactionRowsCountWarningThresholdDefault{
		_statusCode: code,
	}
}

/*
FindConfigCompactionRowsCountWarningThresholdDefault handles this case with default header values.

unexpected error
*/
type FindConfigCompactionRowsCountWarningThresholdDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config compaction rows count warning threshold default response
func (o *FindConfigCompactionRowsCountWarningThresholdDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCompactionRowsCountWarningThresholdDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCompactionRowsCountWarningThresholdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCompactionRowsCountWarningThresholdDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}