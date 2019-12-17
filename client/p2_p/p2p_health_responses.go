//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// P2pHealthReader is a Reader for the P2pHealth structure.
type P2pHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *P2pHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewP2pHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewP2pHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewP2pHealthOK creates a P2pHealthOK with default headers values
func NewP2pHealthOK() *P2pHealthOK {
	return &P2pHealthOK{}
}

/*P2pHealthOK handles this case with default header values.

Alive and kicking!
*/
type P2pHealthOK struct {
}

func (o *P2pHealthOK) Error() string {
	return fmt.Sprintf("[GET /p2p/health][%d] p2pHealthOK ", 200)
}

func (o *P2pHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewP2pHealthInternalServerError creates a P2pHealthInternalServerError with default headers values
func NewP2pHealthInternalServerError() *P2pHealthInternalServerError {
	return &P2pHealthInternalServerError{}
}

/*P2pHealthInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type P2pHealthInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *P2pHealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /p2p/health][%d] p2pHealthInternalServerError  %+v", 500, o.Payload)
}

func (o *P2pHealthInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *P2pHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}