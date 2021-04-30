package mymodels

import (
	helper "hcn/myhelpers/structValidationHelper"
)

// Announcement struct
type Announcement struct {
	ID           *int    `json:"ID,omitempty"`
	CourseID     *int    `json:"CourseID,omitempty"`
	Title        *string `json:"Title,omitempty"`
	Description  *string `json:"Description,omitempty"`
	CreationDate *string `json:"CreationDate,omitempty"`
}

// AllAnnouncements slice of announcements
type AllAnnouncements []Announcement

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model Announcement) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model Announcement) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}
