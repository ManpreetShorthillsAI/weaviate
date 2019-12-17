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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ActionsValidateReader is a Reader for the ActionsValidate structure.
type ActionsValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionsValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewActionsValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActionsValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewActionsValidateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewActionsValidateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionsValidateOK creates a ActionsValidateOK with default headers values
func NewActionsValidateOK() *ActionsValidateOK {
	return &ActionsValidateOK{}
}

/*ActionsValidateOK handles this case with default header values.

Successfully validated.
*/
type ActionsValidateOK struct {
}

func (o *ActionsValidateOK) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] actionsValidateOK ", 200)
}

func (o *ActionsValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsValidateUnauthorized creates a ActionsValidateUnauthorized with default headers values
func NewActionsValidateUnauthorized() *ActionsValidateUnauthorized {
	return &ActionsValidateUnauthorized{}
}

/*ActionsValidateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ActionsValidateUnauthorized struct {
}

func (o *ActionsValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] actionsValidateUnauthorized ", 401)
}

func (o *ActionsValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsValidateForbidden creates a ActionsValidateForbidden with default headers values
func NewActionsValidateForbidden() *ActionsValidateForbidden {
	return &ActionsValidateForbidden{}
}

/*ActionsValidateForbidden handles this case with default header values.

Forbidden
*/
type ActionsValidateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ActionsValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] actionsValidateForbidden  %+v", 403, o.Payload)
}

func (o *ActionsValidateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActionsValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsValidateUnprocessableEntity creates a ActionsValidateUnprocessableEntity with default headers values
func NewActionsValidateUnprocessableEntity() *ActionsValidateUnprocessableEntity {
	return &ActionsValidateUnprocessableEntity{}
}

/*ActionsValidateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type ActionsValidateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ActionsValidateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] actionsValidateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ActionsValidateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActionsValidateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionsValidateInternalServerError creates a ActionsValidateInternalServerError with default headers values
func NewActionsValidateInternalServerError() *ActionsValidateInternalServerError {
	return &ActionsValidateInternalServerError{}
}

/*ActionsValidateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ActionsValidateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ActionsValidateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /actions/validate][%d] actionsValidateInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionsValidateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ActionsValidateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}