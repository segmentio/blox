package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ContainerInstanceDetailRemainingResourceModel container instance detail remaining resource model

swagger:model ContainerInstanceDetailRemainingResourceModel
*/
type ContainerInstanceDetailRemainingResourceModel struct {

	/* name

	Required: true
	*/
	Name *string `json:"name"`

	/* type

	Required: true
	*/
	Type *string `json:"type"`

	/* value

	Required: true
	*/
	Value *string `json:"value"`
}

// Validate validates this container instance detail remaining resource model
func (m *ContainerInstanceDetailRemainingResourceModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerInstanceDetailRemainingResourceModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstanceDetailRemainingResourceModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ContainerInstanceDetailRemainingResourceModel) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}