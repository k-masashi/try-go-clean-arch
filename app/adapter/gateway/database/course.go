package database

import (
	"github.com/k-masashi/try-go-clean-arch/app/domain"
)

type CourseRepositoryGateWay struct {
	SqlHandler
}

func NewCourseRepository() *CourseRepositoryGateWay {
	return &CourseRepositoryGateWay{}
}

func (repository *CourseRepositoryGateWay) FindAll() ([]domain.Course, error) {
	row, error := repository.Query("SELECT id, name, description, created_at FROM course")
	defer row.Close()
	if error != nil {
		return nil, error
	}

	var (
		courseID    int
		name        string
		description string
		createdAt   string
	)
	courses := []domain.Course{ }
	for row.Next() {
		error = row.Scan(
			&courseID,
			&name,
			&description,
			&createdAt
		)

		if error != nil {
			return nil, error
		}

		courses = append(
			courses, 
			domain.Course{
				ID: domain.CourseID(courseID),
				Name: name,
				Description: description,
				CreatedAt: createdAt,
			}
		)
	}
	return courses, nil
}

func (repository *CourseRepositoryGateWay) Find(id domain.CourseID) (*domain.Course, error) {
	row, error := repository.Query("SELECT id, name, description, created_at FROM course WHERE id = ?", id)
	defer row.Close()
	if error != nil {
		return nil, error
	}

	if (!row.HasNext()) {
		// TODO: Error Handling
		return nil, nil
	}

	var (
		courseID    int
		name        string
		description string
		createdAt   string
	)
	error = row.Scan(
		&courseID,
		&name,
		&description,
		&createdAt
	)

	if error != nil {
		return nil, error
	}

	course := domain.Course {
		ID: domain.CourseID(courseID),
		Name: name,
		Description: description,
		CreatedAt: createdAt,
	}
	return &course, nil
}