// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SiteSetup site setup
// swagger:model siteSetup

type SiteSetup struct {
	Site

	// repo
	Repo *RepoInfo `json:"repo,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SiteSetup) UnmarshalJSON(raw []byte) error {

	var aO0 Site
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Site = aO0

	var data struct {
		Repo *RepoInfo `json:"repo,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Repo = data.Repo

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SiteSetup) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.Site)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var data struct {
		Repo *RepoInfo `json:"repo,omitempty"`
	}

	data.Repo = m.Repo

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this site setup
func (m *SiteSetup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Site.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SiteSetup) validateRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.Repo) { // not required
		return nil
	}

	if m.Repo != nil {

		if err := m.Repo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SiteSetup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SiteSetup) UnmarshalBinary(b []byte) error {
	var res SiteSetup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}