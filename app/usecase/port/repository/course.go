package repository

import (
	"github.com/k-masashi/try-go-clean-arch/app/domain"
)

type CourseRepository interface {
	FindAll() ([]domain.Course, error)
	FindByID(domain.CourseID) (*domain.Course, error)
}
