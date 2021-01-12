package myhandlers

import (
	"database/sql"
)

// ProductModel bla bla...
type ProductModel struct {
	Db *sql.DB
}

// Student bla bla...
type Student struct {
	ID    *int    `json:"ID"`
	Name  *string `json:"Name"`
	Email *string `json:"Email"`
}

// AllStudents bla bla...
type AllStudents []Student

// Teacher bla bla...
type Teacher struct {
	ID    *int    `json:"ID"`
	Name  *string `json:"Name"`
	Email *string `json:"Email"`
}

// AllTeachers bla bla...
type AllTeachers []Teacher

// Course bla bla...
type Course struct {
	ID           *int    `json:"ID"`
	TeacherID    *int    `json:"TeacherID"`
	Name         *string `json:"Name"`
	CreationDate *string `json:"CreationDate"`
}

// AllCourses bla bla...
type AllCourses []Course

// StudentTuition bla bla...
type StudentTuition struct {
	ID        *int `json:"ID"`
	CourseID  *int `json:"CourseID"`
	StudentID *int `json:"StudentID"`
}

// AllStudentTuition bla bla...
type AllStudentTuition []StudentTuition

// Announcement bla bla...
type Announcement struct {
	ID           *int    `json:"ID"`
	CourseID     *int    `json:"CourseID"`
	Title        *string `json:"Title"`
	Description  *string `json:"Description"`
	CreationDate *string `json:"CreationDate"`
}

// AllAnnouncements bla bla...
type AllAnnouncements []Announcement

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
}

// AllActivities bla bla...
type AllActivities []Activity

// Feedback bla bla...
type Feedback struct {
	ID         *int `json:"ID"`
	ActivityID *int `json:"ActivityID"`
	StudentID  *int `json:"StudentID"`
}

// AllFeedbacks bla bla...
type AllFeedbacks []Feedback

// HCN (Historia Clínica Nutricional)...
type HCN struct {
	ID        *int `json:"ID"`
	TeacherID *int `json:"TeacherID"`
}

// HCNVinculation bla bla...
type HCNVinculation struct {
	ID             *int `json:"ID"`
	ClinicalCaseID *int `json:"ClinicalCaseID"`
	HCNID          *int `json:"HCNID"`
}

// AllHCN bla bla...
type AllHCN []HCN

/*
// HCNCourse struct that represents the new relationship between a HCN and a Course ...
type HCNCourse struct {
	ID          *int `json:"ID"`
	CoursesID   *int `json:"CoursesID"`
	HCNID       *int `json:"HCNID"`
	Displayable *int `json:"Displayable"`
}

// AllHCNsCourses bla bla...
type AllHCNsCourses []HCNCourse

// HCNCCase struct that represents the new relationship between a HCN and a Clinical Case ...
type HCNCCase struct {
	ID              *int `json:"ID"`
	ClinicalCasesID *int `json:"ClinicalCasesID"`
	HCNID           *int `json:"HCNID"`
}

// AllHCNsCCases bla bla...
type AllHCNsCCases []HCNCCase
*/
