package controller

import (
	"github.com/k-masashi/try-go-clean-arch/app/adapter/presenter"
	"github.com/k-masashi/try-go-clean-arch/app/adapter/gateway/database"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port/server"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/interactor"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/interactor"
)

type CourseController struct {
	InputPort server.CourseInputPort
}

func NewCourseController() *CourseController {
	return &CourseController{
		InputPort: interactor.NewCourseInteractor(
			presenter.NewCourseHTTPPresenter(),
			database.NewCourseRepository()
		)
	}
}

func (controller *CourseController) GetCourses() (*server.GetCoursesResponse, error) {
	return controller.InputPort.GetCourses()
}

func (controller *CourseController) GetCourse(request *server.GetCourseRequest) (*server.GetCourseResponse, error) {
	return controller.InputPort.GetCourse(request)
}
