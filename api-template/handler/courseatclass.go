package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllAts(c echo.Context) error {
	ats, err := h.repo.CourseAtClass.GetAll(h.store)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, ats)
}
func (h *Handler) FindAtsByID(c echo.Context) error {
	id := c.Param("id")
	at, err := h.repo.CourseAtClass.GetByID(h.store, id)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	if at.ClassId == "" {
		return util.HandleError(c, errors.ErrAtNotfound)
	}

	return c.JSON(http.StatusOK, at)
}
func (h *Handler) AddAts(c echo.Context) error {
	classid := c.QueryParam("classid")
	courseid := c.QueryParam("courseid")

	username := c.QueryParam("username")
	password := c.QueryParam("password")
	role, checkerr := h.GetRole(c, username, password)
	if checkerr != nil {
		return util.HandleError(c, checkerr)
	}
	if role != "admin" {
		return util.HandleError(c, errors.ErrUnAuthorized)
	}

	param := model.CourseAtClass{ClassId: classid, CourseId: courseid}

	ats, err := h.repo.CourseAtClass.PostAts(h.store, param)

	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, ats)
}

func (h *Handler) DeleteAtsByID(c echo.Context) error {
	id := c.QueryParam("id")

	username := c.QueryParam("username")
	password := c.QueryParam("password")
	role, checkerr := h.GetRole(c, username, password)
	if checkerr != nil {
		return util.HandleError(c, checkerr)
	}
	if role != "admin" {
		return util.HandleError(c, errors.ErrUnAuthorized)
	}

	//check records exist
	ats, _ := h.repo.CourseAtClass.GetByID(h.store, id)
	if ats.ClassId == "" {
		return util.HandleError(c, errors.ErrAtsNotfound)
	}

	_, err := h.repo.CourseAtClass.DeleteByID(h.store, id)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, "delete successfully the course_at_class with id: "+id)
}
