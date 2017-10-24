package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DebugConfig debug config
// swagger:model DebugConfig
type DebugConfig struct {

	// active
	Active bool `json:"active,omitempty"`

	// attached
	Attached bool `json:"attached,omitempty"`

	// attachment
	// Required: true
	Attachment *Attachment `json:"attachment"`

	// breakpoints
	Breakpoints []*Breakpoint `json:"breakpoints"`

	// debugger
	Debugger string `json:"debugger,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// image
	// Required: true
	Image *string `json:"image"`

	// immediately
	Immediately bool `json:"immediately,omitempty"`
}

// Validate validates this debug config
func (m *DebugConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachment(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBreakpoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebugConfig) validateAttachment(formats strfmt.Registry) error {

	if err := validate.Required("attachment", "body", m.Attachment); err != nil {
		return err
	}

	if m.Attachment != nil {

		if err := m.Attachment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment")
			}
			return err
		}
	}

	return nil
}

func (m *DebugConfig) validateBreakpoints(formats strfmt.Registry) error {

	if swag.IsZero(m.Breakpoints) { // not required
		return nil
	}

	for i := 0; i < len(m.Breakpoints); i++ {

		if swag.IsZero(m.Breakpoints[i]) { // not required
			continue
		}

		if m.Breakpoints[i] != nil {

			if err := m.Breakpoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("breakpoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DebugConfig) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DebugConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebugConfig) UnmarshalBinary(b []byte) error {
	var res DebugConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
