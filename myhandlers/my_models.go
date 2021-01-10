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

// Announcement bla bla...
type Announcement struct {
	ID           *int    `json:"ID"`
	CoursesID    *int    `json:"CoursesID"`
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
	TeachersID  *int    `json:"TeachersID"`
}

// AllClinicalCases bla bla...
type AllClinicalCases []ClinicalCase

// Course bla bla...
type Course struct {
	ID           *int    `json:"ID"`
	Teacher      *int    `json:"Teacher"`
	Name         *string `json:"Name"`
	CreationDate *string `json:"CreationDate"`
}

// AllCourses bla bla...
type AllCourses []Course

// Activity bla bla...
type Activity struct {
	ID              *int    `json:"ID"`
	Title           *string `json:"Title"`
	Description     *string `json:"Description"`
	Type            *string `json:"Type"`
	CreationDate    *string `json:"CreationDate"`
	LimitDate       *string `json:"LimitDate"`
	CoursesID       *int    `json:"CoursesID"`
	ClinicalCasesID *int    `json:"ClinicalCasesID"`
	Difficulty      *int    `json:"Difficulty"`
}

// AllActivities bla bla...
type AllActivities []Activity

// Feedback bla bla...
type Feedback struct {
	ID           *int `json:"ID"`
	ActivitiesID *int `json:"ActivitiesID"`
	StudentsID   *int `json:"StudentsID"`
}

// AllFeedbacks bla bla...
type AllFeedbacks []Feedback
