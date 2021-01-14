package mymodels

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
