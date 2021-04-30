package mymodels

import (
	helper "hcn/myhelpers/structValidationHelper"
)

// Student struct
type Student struct {
	ID    *int    `json:"ID,omitempty"`
	Name  *string `json:"Name,omitempty"`
	Email *string `json:"Email,omitempty"`
}

// AllStudents slice of students
type AllStudents []Student

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model Student) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model Student) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}

// StudentTuition struct
type StudentTuition struct {
	ID        *int `json:"ID,omitempty"`
	CourseID  *int `json:"CourseID,omitempty"`
	StudentID *int `json:"StudentID,omitempty"`
}

// AllStudentTuition slice of student tuitions.
type AllStudentTuition []StudentTuition

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model StudentTuition) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model StudentTuition) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}
