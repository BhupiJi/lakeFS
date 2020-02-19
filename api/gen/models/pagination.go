// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Pagination pagination
// swagger:model pagination
type Pagination struct {

	// has more
	// Required: true
	HasMore *bool `json:"has_more"`

	// max per page
	// Minimum: 0
	MaxPerPage *int64 `json:"max_per_page,omitempty"`

	// next offset
	NextOffset string `json:"next_offset,omitempty"`

	// results
	// Minimum: 0
	Results *int64 `json:"results,omitempty"`
}

// Validate validates this pagination
func (m *Pagination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasMore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxPerPage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pagination) validateHasMore(formats strfmt.Registry) error {

	if err := validate.Required("has_more", "body", m.HasMore); err != nil {
		return err
	}

	return nil
}

func (m *Pagination) validateMaxPerPage(formats strfmt.Registry) error {

	if swag.IsZero(m.MaxPerPage) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_per_page", "body", int64(*m.MaxPerPage), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Pagination) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	if err := validate.MinimumInt("results", "body", int64(*m.Results), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Pagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pagination) UnmarshalBinary(b []byte) error {
	var res Pagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}