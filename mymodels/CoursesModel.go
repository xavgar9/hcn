package mymodels

import helper "hcn/myhelpers/structValidationHelper"

// Course struct
type Course struct {
	ID           *int    `json:"ID"`
	TeacherID    *int    `json:"TeacherID"`
	Name         *string `json:"Name"`
	CreationDate *string `json:"CreationDate"`
}

// AllCourses slice of courses
type AllCourses []Course

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model Course) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model Course) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}

// CourseHCN struct that represents the new relationship between a HCN and a Course ...
type CourseHCN struct {
	ID          *int `json:"ID"`
	CourseID    *int `json:"CourseID"`
	HCNID       *int `json:"HCNID"`
	Displayable *int `json:"Displayable"`
}

// AllCourseHCN bla bla...
type AllCourseHCN []CourseHCN

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model CourseHCN) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model CourseHCN) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}

// CourseClinicalCase struct that represents the new relationship between a Clinical Case and a Course ...
type CourseClinicalCase struct {
	ID             *int `json:"ID"`
	CourseID       *int `json:"CourseID"`
	ClinicalCaseID *int `json:"ClinicalCaseID"`
	Displayable    *int `json:"Displayable"`
}

// AllCourseClinicalCase bla bla...
type AllCourseClinicalCase []CourseClinicalCase

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model CourseClinicalCase) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model CourseClinicalCase) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}

/*
// HCNCCase struct that represents the new relationship between a HCN and a Clinical Case ...
type HCNCCase struct {
	ID              *int `json:"ID"`
	ClinicalCasesID *int `json:"ClinicalCasesID"`
	HCNID           *int `json:"HCNID"`
}

// AllHCNsCCases bla bla...
type AllHCNsCCases []HCNCCase
*/
