package mymodels

import helper "hcn/myhelpers/structValidationHelper"

// SolvedHCN struct
type SolvedHCN struct {
	ID          *int    `json:"ID"`
	ActivityID  *int    `json:"ActivityID,omitempty"`
	Solver      *int    `json:"Solver,omitempty"`
	OriginalHCN *int    `json:"OriginalHCN,omitempty"`
	CourseID    *int    `json:"CourseID,omitempty"`
	MongoID     *string `json:"MongoID,omitempty"`
	Reviewed    *int    `json:"Reviewed,omitempty"`
	TeacherID   *int    `json:"TeacherID,omitempty"`
}

// AllSolvedHCN slice of SolvedHCN
type AllSolvedHCN []SolvedHCN

// ValidateFields checks the fields of the struct.
//
// If not struct fields are given, will check all fields.
func (model SolvedHCN) ValidateFields(structFields ...[]string) (bool, error) {
	if len(structFields) == 0 {
		return helper.ValidateFields(model)
	}
	return helper.ValidateFields(model, structFields[0])
}

// GetFields return the names and fields values of the struct.
func (model SolvedHCN) GetFields() (string, []string, []string, error) {
	return helper.GetFields(model)
}
