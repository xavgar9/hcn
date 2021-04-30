package mymodels

import helper "hcn/myhelpers/structValidationHelper"

// ClinicalCase struct
type ClinicalCase struct {
	ID          *int    `json:"ID,omitempty"`
	Title       *string `json:"Title,omitempty"`
	Description *string `json:"Description,omitempty"`
	Media       *string `json:"Media,omitempty"`
	TeacherID   *int    `json:"TeacherID,omitempty"`
}

// AllClinicalCases bslice of Clinical Cases
type AllClinicalCases []ClinicalCase

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model ClinicalCase) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model ClinicalCase) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}
