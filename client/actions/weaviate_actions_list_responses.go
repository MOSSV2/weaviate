/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

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

// WeaviateActionsListReader is a Reader for the WeaviateActionsList structure.
type WeaviateActionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateActionsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsListOK creates a WeaviateActionsListOK with default headers values
func NewWeaviateActionsListOK() *WeaviateActionsListOK {
	return &WeaviateActionsListOK{}
}

/*WeaviateActionsListOK handles this case with default header values.

Successful response.
*/
type WeaviateActionsListOK struct {
	Payload *models.ActionsListResponse
}

func (o *WeaviateActionsListOK) Error() string {
	return fmt.Sprintf("[GET /actions][%d] weaviateActionsListOK  %+v", 200, o.Payload)
}

func (o *WeaviateActionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsListUnauthorized creates a WeaviateActionsListUnauthorized with default headers values
func NewWeaviateActionsListUnauthorized() *WeaviateActionsListUnauthorized {
	return &WeaviateActionsListUnauthorized{}
}

/*WeaviateActionsListUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsListUnauthorized struct {
}

func (o *WeaviateActionsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /actions][%d] weaviateActionsListUnauthorized ", 401)
}

func (o *WeaviateActionsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsListForbidden creates a WeaviateActionsListForbidden with default headers values
func NewWeaviateActionsListForbidden() *WeaviateActionsListForbidden {
	return &WeaviateActionsListForbidden{}
}

/*WeaviateActionsListForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateActionsListForbidden struct {
}

func (o *WeaviateActionsListForbidden) Error() string {
	return fmt.Sprintf("[GET /actions][%d] weaviateActionsListForbidden ", 403)
}

func (o *WeaviateActionsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsListNotFound creates a WeaviateActionsListNotFound with default headers values
func NewWeaviateActionsListNotFound() *WeaviateActionsListNotFound {
	return &WeaviateActionsListNotFound{}
}

/*WeaviateActionsListNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateActionsListNotFound struct {
}

func (o *WeaviateActionsListNotFound) Error() string {
	return fmt.Sprintf("[GET /actions][%d] weaviateActionsListNotFound ", 404)
}

func (o *WeaviateActionsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsListInternalServerError creates a WeaviateActionsListInternalServerError with default headers values
func NewWeaviateActionsListInternalServerError() *WeaviateActionsListInternalServerError {
	return &WeaviateActionsListInternalServerError{}
}

/*WeaviateActionsListInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateActionsListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /actions][%d] weaviateActionsListInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
