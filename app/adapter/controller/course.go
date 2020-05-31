package controller

import (
	"github.com/k-masashi/try-go-clean-arch/app/adapter/gateway/database"
	"github.com/k-masashi/try-go-clean-arch/app/adapter/presenter"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/interactor"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port"
)

type CourseController struct {
	InputPort port.CourseInputPort
}

func NewCourseController(sqlHandler database.SqlHandler) *CourseController {
	return &CourseController{
		InputPort: interactor.NewCourseInteractor(
			presenter.NewCourseHTTPPresenter(),
			database.NewCourseRepository(sqlHandler),
		),
	}
}

func (controller *CourseController) GetCourses() (*port.GetCoursesResponse, error) {
	return controller.InputPort.GetCourses()
}

func (controller *CourseController) GetCourse(request *port.GetCourseRequest) (*port.GetCourseResponse, error) {
	return controller.InputPort.GetCourse(request)
}
