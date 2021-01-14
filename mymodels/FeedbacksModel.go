package mymodels

// Feedback bla bla...
type Feedback struct {
	ID         *int `json:"ID"`
	ActivityID *int `json:"ActivityID"`
	StudentID  *int `json:"StudentID"`
}

// AllFeedbacks bla bla...
type AllFeedbacks []Feedback
