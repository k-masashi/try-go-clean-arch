package domain

type CourseID string

type Course struct {
	ID          CourseID
	Name        string
	Description string
	CreatedAt   string
}
