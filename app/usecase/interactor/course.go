package interactor

import (
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port/repository"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port/server"
)

type CourseInteractor struct {
	OutputPort       server.CourseOutputPort
	CourseRepository repository.CourseRepository
}

func NewCourseInteractor(
	outputPort server.CourseOutputPort,
	courseRepository repository.CourseRepository,
) *CourseInteractor {
	return &CourseInteractor{
		OutputPort:       outputPort,
		CourseRepository: courseRepository,
	}
}

func (interactor *CourseInteractor) GetCourses() (*server.GetCoursesResponse, error) {
	response, error := interactor.CourseRepository.FindAll()

	if error != nil {
		return nil, error
	}

	return interactor.OutputPort.GetCourses(response)
}

func (interactor *CourseInteractor) GetCourse(request *server.GetCourseRequest) (*server.GetCourseResponse, error) {
	response, error := interactor.CourseRepository.FindByID(request.CourseID)

	if error != nil {
		return nil, error
	}

	return interactor.OutputPort.GetCourse(response)
}
