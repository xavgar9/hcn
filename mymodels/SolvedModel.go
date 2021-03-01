package mymodels

// SolvedHCN bla bla...
type SolvedHCN struct {
	ID          *int    `json:"ID"`
	OriginalHCN *int    `json:"OriginalHCN"`
	CourseID    *int    `json:"CourseID"`
	MongoID     *string `json:"MongoID"`
	Solver      *int    `json:"Solver"`
	Reviewed    *int    `json:"Reviewed"`
	TeacherID   *int    `json:"TeacherID"`
}

// AllSolvedHCN bla bla...
type AllSolvedHCN []SolvedHCN
