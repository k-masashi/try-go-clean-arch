package presenter

import (
	"github.com/k-masashi/try-go-clean-arch/app/domain"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port/server"
)

type CourseHTTPPresenter struct {
	server.CourseOutputPort
}

func NewCourseHTTPPresenter() *CourseHTTPPresenter {
	return &CourseHTTPPresenter{}
}

func (presenter *CourseHTTPPresenter) GetCourses(courses []domain.Course) (*server.GetCoursesResponse, error) {
	response := &server.GetCoursesResponse{
		Courses: courses,
	}
	return response, nil
}

func (presenter *CourseHTTPPresenter) GetCourse(course *domain.Course) (*server.GetCourseResponse, error) {
	response := &server.GetCourseResponse{
		Course: course,
	}
	return response, nil
}
