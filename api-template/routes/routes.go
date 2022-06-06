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

	courseGroup := apiV1Group.Group(("/class"))
	{
		courseGroup.GET("/all", h.GetAllClass)
		courseGroup.POST("/add", h.AddClass)
	}
}
