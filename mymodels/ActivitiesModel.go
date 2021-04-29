package mymodels

import (
	helper "hcn/myhelpers/structValidationHelper"
)

// Activity struct
type Activity struct {
	ID             *int    `json:"ID,omitempty"`
	Title          *string `json:"Title,omitempty"`
	Description    *string `json:"Description,omitempty"`
	Type           *string `json:"Type,omitempty"`
	CreationDate   *string `json:"CreationDate,omitempty"`
	LimitDate      *string `json:"LimitDate,omitempty"`
	CourseID       *int    `json:"CourseID,omitempty"`
	ClinicalCaseID *int    `json:"ClinicalCaseID,omitempty"`
	HCNID          *int    `json:"HCNID,omitempty"`
	Difficulty     *int    `json:"Difficulty,omitempty"`
}

// AllActivities slice of structs
type AllActivities []Activity

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model Activity) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model Activity) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}
