package echo

import (
	"fmt"

	"github.com/k-masashi/try-go-clean-arch/app/adapter/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type Server struct {
	*echo.Echo
}

func Run() {
	server, error := createServer()

	if error != nil {
		panic(error)
	}

	server.setRouter()
	server.run()
}

func createServer() (*Server, error) {
	return &Server{
		Echo: echo.New(),
	}, nil
}

func (server *Server) setRouter() {
	server.Echo.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	api := server.Echo.Group(server.getEndpointRoot())
	{
		courseController := controller.NewCourseController()
		api.GET("/courses", server.GetCourses(courseController))
		api.GET("/course/:id", server.GetCourse(courseController))
	}
}

func (server *Server) getEndpointRoot() string {
	return fmt.Sprintf("/api/%s", viper.GetString("version"))
}

func (server *Server) run() {
	server.Echo.Logger.Fatal(server.Echo.Start(viper.GetString("server.address")))
}
