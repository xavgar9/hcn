package mymodels

// Teacher bla bla...
type Teacher struct {
	ID    *int    `json:"ID"`
	Name  *string `json:"Name"`
	Email *string `json:"Email"`
}

// AllTeachers bla bla...
type AllTeachers []Teacher
