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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaThingsPropertiesDeleteReader is a Reader for the WeaviateSchemaThingsPropertiesDelete structure.
type WeaviateSchemaThingsPropertiesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaThingsPropertiesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaThingsPropertiesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaThingsPropertiesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaThingsPropertiesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaThingsPropertiesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaThingsPropertiesDeleteOK creates a WeaviateSchemaThingsPropertiesDeleteOK with default headers values
func NewWeaviateSchemaThingsPropertiesDeleteOK() *WeaviateSchemaThingsPropertiesDeleteOK {
	return &WeaviateSchemaThingsPropertiesDeleteOK{}
}

/*WeaviateSchemaThingsPropertiesDeleteOK handles this case with default header values.

Removed the property from the ontology.
*/
type WeaviateSchemaThingsPropertiesDeleteOK struct {
}

func (o *WeaviateSchemaThingsPropertiesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesDeleteOK ", 200)
}

func (o *WeaviateSchemaThingsPropertiesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesDeleteUnauthorized creates a WeaviateSchemaThingsPropertiesDeleteUnauthorized with default headers values
func NewWeaviateSchemaThingsPropertiesDeleteUnauthorized() *WeaviateSchemaThingsPropertiesDeleteUnauthorized {
	return &WeaviateSchemaThingsPropertiesDeleteUnauthorized{}
}

/*WeaviateSchemaThingsPropertiesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaThingsPropertiesDeleteUnauthorized struct {
}

func (o *WeaviateSchemaThingsPropertiesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesDeleteUnauthorized ", 401)
}

func (o *WeaviateSchemaThingsPropertiesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesDeleteForbidden creates a WeaviateSchemaThingsPropertiesDeleteForbidden with default headers values
func NewWeaviateSchemaThingsPropertiesDeleteForbidden() *WeaviateSchemaThingsPropertiesDeleteForbidden {
	return &WeaviateSchemaThingsPropertiesDeleteForbidden{}
}

/*WeaviateSchemaThingsPropertiesDeleteForbidden handles this case with default header values.

Could not find the Thing class or property.
*/
type WeaviateSchemaThingsPropertiesDeleteForbidden struct {
}

func (o *WeaviateSchemaThingsPropertiesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesDeleteForbidden ", 403)
}

func (o *WeaviateSchemaThingsPropertiesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesDeleteInternalServerError creates a WeaviateSchemaThingsPropertiesDeleteInternalServerError with default headers values
func NewWeaviateSchemaThingsPropertiesDeleteInternalServerError() *WeaviateSchemaThingsPropertiesDeleteInternalServerError {
	return &WeaviateSchemaThingsPropertiesDeleteInternalServerError{}
}

/*WeaviateSchemaThingsPropertiesDeleteInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaThingsPropertiesDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsPropertiesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaThingsPropertiesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
