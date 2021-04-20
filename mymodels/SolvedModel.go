package mymodels

// SolvedHCN bla bla...
type SolvedHCN struct {
	ID          *int    `json:"ID"`
	ActivityID  *int    `json:"ActivityID"`
	OriginalHCN *int    `json:"OriginalHCN"`
	CourseID    *int    `json:"CourseID,omitempty"`
	MongoID     *string `json:"MongoID"`
	Solver      *int    `json:"Solver"`
	Reviewed    *int    `json:"Reviewed"`
	TeacherID   *int    `json:"TeacherID,omitempty"`
}

// AllSolvedHCN bla bla...
type AllSolvedHCN []SolvedHCN
