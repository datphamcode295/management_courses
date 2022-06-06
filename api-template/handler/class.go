package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/repo/pg"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllClass(c echo.Context) error {

	sliceClasses, err := pg.GetAllClasses(h.store)

	c.JSON(http.StatusOK, sliceClasses)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return nil
}

func (h *Handler) AddClass(c echo.Context) error {
	// id := c.QueryString()

	classParams := model.Class{ID: "3"}
	sliceClasses, err := h.repo.Class.PostClass(h.store, classParams)

	c.JSON(http.StatusOK, sliceClasses)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}

	// h.store.DB().Create(&(classParams))
	return nil
}
