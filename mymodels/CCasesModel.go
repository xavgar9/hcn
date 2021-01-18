package mymodels

// ClinicalCase bla bla...
type ClinicalCase struct {
	ID          *int    `json:"ID"`
	Title       *string `json:"Title"`
	Description *string `json:"Description"`
	Media       *string `json:"Media"`
	TeacherID   *int    `json:"TeacherID"`
}

// AllClinicalCases bla bla...
type AllClinicalCases []ClinicalCase
