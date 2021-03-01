package mymodels

// Activity bla bla...
type Activity struct {
	ID             *int    `json:"ID"`
	Title          *string `json:"Title"`
	Description    *string `json:"Description"`
	Type           *string `json:"Type"`
	CreationDate   *string `json:"CreationDate"`
	LimitDate      *string `json:"LimitDate"`
	CourseID       *int    `json:"CourseID"`
	ClinicalCaseID *int    `json:"ClinicalCaseID"`
	HCNID          *int    `json:"HCNID"`
	Difficulty     *int    `json:"Difficulty"`
	TeacherID      *int    `json:"TeacherID"`
}

// AllActivities bla bla...
type AllActivities []Activity
