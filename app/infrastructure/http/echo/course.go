package echo

import (
	"net/http"
	"strconv"

	"github.com/k-masashi/try-go-clean-arch/app/adapter/controller"
	"github.com/k-masashi/try-go-clean-arch/app/domain"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port"
	"github.com/labstack/echo"
)

type CourseRequestParameters struct {
	ID int
}

type Course struct {
	ID          domain.CourseID
	Name        string
	Description string
	CreatedAt   string
}

func (server *Server) GetCourses(controller *controller.CourseController) echo.HandlerFunc {
	return func(context echo.Context) error {
		response, error := controller.GetCourses()
		if error != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		coursesResponse, httpStatus := convertToCourses(response)
		return context.JSON(httpStatus, coursesResponse)
	}
}

func (server *Server) GetCourse(controller *controller.CourseController) echo.HandlerFunc {
	return func(context echo.Context) error {
		id, _ := strconv.Atoi(context.Param("id"))
		request := CourseRequestParameters{
			ID: id,
		}
		getCourseRequest := port.GetCourseRequest{
			CourseID: domain.CourseID(request.ID),
		}
		response, error := controller.GetCourse(&getCourseRequest)
		if error != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		coursesResponse, httpStatus := convertToCourseResponse(response)
		return context.JSON(httpStatus, coursesResponse)
	}
}

type CourseResponse struct {
	Course *Course `json:"course"`
}

type CoursesResponse struct {
	Courses []Course `json:"courses"`
}

func convertToCourseResponse(response *port.GetCourseResponse) (CourseResponse, int) {
	var jsonResponse CourseResponse

	if response.Course == nil {
		jsonResponse = CourseResponse{
			Course: nil,
		}
		return jsonResponse, http.StatusOK
	}

	jsonResponse = CourseResponse{
		Course: &Course{
			ID:          response.Course.ID,
			Name:        response.Course.Name,
			Description: response.Course.Description,
			CreatedAt:   response.Course.CreatedAt,
		},
	}
	return jsonResponse, http.StatusOK
}

func convertToCourses(response *port.GetCoursesResponse) (CoursesResponse, int) {
	jsonResponse := CoursesResponse{
		Courses: []Course{},
	}

	for _, course := range response.Courses {
		jsonResponse.Courses = append(
			jsonResponse.Courses,
			Course{
				ID:          course.ID,
				Name:        course.Name,
				Description: course.Description,
				CreatedAt:   course.CreatedAt,
			},
		)
	}
	return jsonResponse, http.StatusOK
}
