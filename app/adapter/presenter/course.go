package presenter

import (
	"github.com/k-masashi/try-go-clean-arch/app/domain"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port"
)

type CourseHTTPPresenter struct {
	port.CourseOutputPort
}

func NewCourseHTTPPresenter() *CourseHTTPPresenter {
	return &CourseHTTPPresenter{}
}

func (presenter *CourseHTTPPresenter) GetCourses(courses []domain.Course) (*port.GetCoursesResponse, error) {
	response := &port.GetCoursesResponse{
		Courses: courses,
	}
	return response, nil
}

func (presenter *CourseHTTPPresenter) GetCourse(course *domain.Course) (*port.GetCourseResponse, error) {
	response := &port.GetCourseResponse{
		Course: course,
	}
	return response, nil
}
