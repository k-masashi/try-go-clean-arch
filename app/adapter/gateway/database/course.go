package database

import (
	"github.com/k-masashi/try-go-clean-arch/app/domain"
)

type CourseRepositoryGateWay struct {
	SqlHandler SqlHandler
}

func NewCourseRepository(sqlHandler SqlHandler) *CourseRepositoryGateWay {
	return &CourseRepositoryGateWay{
		SqlHandler: sqlHandler,
	}
}

func (repository *CourseRepositoryGateWay) FindAll() ([]domain.Course, error) {
	row, error := repository.SqlHandler.Query("SELECT id, name, description, created_at FROM course")
	if error != nil {
		return nil, error
	}
	defer row.Close()

	var (
		courseID    int
		name        string
		description string
		createdAt   string
	)
	courses := []domain.Course{}
	for row.HasNext() {
		error = row.Scan(
			&courseID,
			&name,
			&description,
			&createdAt,
		)

		if error != nil {
			return nil, error
		}

		courses = append(
			courses,
			domain.Course{
				ID:          domain.CourseID(courseID),
				Name:        name,
				Description: description,
				CreatedAt:   createdAt,
			},
		)
	}
	return courses, nil
}

func (repository *CourseRepositoryGateWay) FindByID(id domain.CourseID) (*domain.Course, error) {
	row, error := repository.SqlHandler.Query("SELECT id, name, description, created_at FROM course WHERE id = ?", id)
	if error != nil {
		return nil, error
	}
	defer row.Close()

	if !row.HasNext() {
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
		&createdAt,
	)

	if error != nil {
		return nil, error
	}

	course := domain.Course{
		ID:          domain.CourseID(courseID),
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
	}
	return &course, nil
}
