package mymodels

import helper "hcn/myhelpers/structValidationHelper"

// Teacher struct
type Teacher struct {
	ID       *int    `json:"ID,omitempty"`
	Name     *string `json:"Name,omitempty"`
	Email    *string `json:"Email,omitempty"`
	Password *string `json:"Password,omitempty"`
}

// AllTeachers slice of teachers
type AllTeachers []Teacher

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model Teacher) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model Teacher) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}
