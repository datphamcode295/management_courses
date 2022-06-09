package routes

import (
	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/handler"
	"github.com/MStation-io/api/repo"
	"github.com/labstack/echo/v4"
)

func NewRoutes(e *echo.Echo, h *handler.Handler, cfg config.Config, s repo.DBRepo) {
	apiV1Group := e.Group("/api/v1")

	authGroup := apiV1Group.Group(("/auth"))
	{
		authGroup.POST("/upsert-wallet", h.UpsertWalletHandler)
		authGroup.GET("/challenge", h.GetChallengeHandler)
	}

	classGroup := apiV1Group.Group(("/class"))
	{
		classGroup.GET("/all", h.GetAllClass)
		classGroup.GET("/id/:id", h.FindClassByID)
		classGroup.GET("/name/:name", h.FindClassByName)
		classGroup.POST("/add", h.AddClass)
		classGroup.GET("/students/:name", h.FindAllStudent)
		classGroup.GET("/course/:name", h.FindAllCourses)
		classGroup.DELETE("/id", h.DeleteClassByID)
	}
	courseGroup := apiV1Group.Group("/course")
	{
		courseGroup.GET("/all", h.GetAllCourse)
		courseGroup.POST("/add", h.AddCourse)
		courseGroup.DELETE("/id", h.DeleteCourseByID)
		courseGroup.GET("/allClasses/:id", h.FindAllClasses)
	}
	atsGroup := apiV1Group.Group("/courseatclass")
	{
		atsGroup.GET("/all", h.GetAllAts)
		atsGroup.GET("/:id", h.FindAtsByID)
		atsGroup.POST("/add", h.AddAts)
		atsGroup.DELETE("/id", h.DeleteAtsByID)

	}
	clientGroup := apiV1Group.Group("/client")
	{
		clientGroup.GET("/all", h.GetAllClient)
		clientGroup.POST("/", h.AddClient)
		clientGroup.DELETE("/id", h.DeleteClient)
		clientGroup.PUT("/", h.UpdateClientPasswordByID)
		clientGroup.GET("/login", h.Login)

	}

}
