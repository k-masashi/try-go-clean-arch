package domain

type CourseID int

type Course struct {
	ID          CourseID
	Name        string
	Description string
	CreatedAt   string
}
