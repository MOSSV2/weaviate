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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SingleRef single ref
// swagger:model SingleRef
type SingleRef struct {

	// URI to point to the cross-ref. Should be in the form of weaviate://localhost/things/<uuid> for the example of a local cross-ref to a thing
	// Format: uri
	NrDollarCref strfmt.URI `json:"$cref,omitempty"`
}

// Validate validates this single ref
func (m *SingleRef) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNrDollarCref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SingleRef) validateNrDollarCref(formats strfmt.Registry) error {

	if swag.IsZero(m.NrDollarCref) { // not required
		return nil
	}

	if err := validate.FormatOf("$cref", "body", "uri", m.NrDollarCref.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SingleRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SingleRef) UnmarshalBinary(b []byte) error {
	var res SingleRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
