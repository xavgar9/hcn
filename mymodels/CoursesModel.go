package mymodels

// Course bla bla...
type Course struct {
	ID           *int    `json:"ID"`
	TeacherID    *int    `json:"TeacherID"`
	Name         *string `json:"Name"`
	CreationDate *string `json:"CreationDate"`
}

// AllCourses bla bla...
type AllCourses []Course

// CourseHCN struct that represents the new relationship between a HCN and a Course ...
type CourseHCN struct {
	ID          *int `json:"ID"`
	CourseID    *int `json:"CourseID"`
	HCNID       *int `json:"HCNID"`
	Displayable *int `json:"Displayable"`
}

// AllCourseHCN bla bla...
type AllCourseHCN []CourseHCN

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
