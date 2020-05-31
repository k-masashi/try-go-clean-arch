package interactor

import (
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port/repository"
)

type CourseInteractor struct {
	OutputPort       port.CourseOutputPort
	CourseRepository repository.CourseRepository
}

func NewCourseInteractor(
	outputPort port.CourseOutputPort,
	courseRepository repository.CourseRepository,
) *CourseInteractor {
	return &CourseInteractor{
		OutputPort:       outputPort,
		CourseRepository: courseRepository,
	}
}

func (interactor *CourseInteractor) GetCourses() (*port.GetCoursesResponse, error) {
	response, error := interactor.CourseRepository.FindAll()

	if error != nil {
		return nil, error
	}

	return interactor.OutputPort.GetCourses(response)
}

func (interactor *CourseInteractor) GetCourse(request *port.GetCourseRequest) (*port.GetCourseResponse, error) {
	response, error := interactor.CourseRepository.FindByID(request.CourseID)

	if error != nil {
		return nil, error
	}

	return interactor.OutputPort.GetCourse(response)
}
