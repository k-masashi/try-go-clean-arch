package server

import (
	"github.com/k-masashi/try-go-clean-arch/app/domain"
)

type CourseInputPort interface {
	GetCourses() (GetCoursesResponse, error)
	GetCourse(*GetCourseRequestParameters) (GetCourseResponse, error)
}

type CourseOutputPort interface {
	GetCourses([]domain.Course) (*GetCoursesResponse, error)
	GetCourse(*domain.Course) (*GetCourseResponse, error)
}

type GetCourseRequestParameters struct {
	CourseID domain.CourseID
}

type GetCourseResponse struct {
	Course *domain.Course
}

type GetCoursesResponse struct {
	Courses []domain.Course
}
