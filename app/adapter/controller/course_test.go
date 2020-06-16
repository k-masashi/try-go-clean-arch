package controller

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/k-masashi/try-go-clean-arch/app/usecase/port"
	mock "github.com/k-masashi/try-go-clean-arch/app/usecase/port/mock"
)

func TestCourseController_GetCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected = &port.GetCoursesResponse{}
	var err error

	mockInputPort := mock.NewMockCourseInputPort(ctrl)
	mockInputPort.EXPECT().GetCourses().Return(expected, err)

	courseCtrl := CourseController{
		InputPort: mockInputPort,
	}
	actual, err := courseCtrl.GetCourses()

	if err != nil {
		t.Error("Actual GetCourses() is not same as expected")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual GetCourses() is not same as expected")
	}
}
