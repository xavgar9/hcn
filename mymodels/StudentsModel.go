package mymodels

// Student bla bla...
type Student struct {
	ID    *int    `json:"ID"`
	Name  *string `json:"Name"`
	Email *string `json:"Email"`
}

// AllStudents bla bla...
type AllStudents []Student

// StudentTuition bla bla...
type StudentTuition struct {
	ID        *int `json:"ID"`
	CourseID  *int `json:"CourseID"`
	StudentID *int `json:"StudentID"`
}

// AllStudentTuition bla bla...
type AllStudentTuition []StudentTuition
