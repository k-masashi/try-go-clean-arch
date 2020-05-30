package controller

import (
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port/server"
)

type CourseController struct {
	InputPort server.CourseInputPort
}

func NewCourseController() *CourseController {
	return nil
}

func (controller *CourseController) GetCourses() (*server.GetCoursesResponse, error) {
	return controller.InputPort.GetCourses()
}

func (controller *CourseController) GetCourse(request *server.GetCourseRequest) (*server.GetCourseResponse, error) {
	return controller.InputPort.GetCourse(request)
}
