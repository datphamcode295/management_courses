package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

// GetAllCourseAtClass
// @Summary get all record in table courseatclass
// @Description get all record in table courseatclass
// @Tags Course At Class
// @Accept	json
// @Produce  json
// @Success 200 {object} model.CourseAtClass	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/courseatclass/all [get]
func (h *Handler) GetAllAts(c echo.Context) error {
	ats, err := h.repo.CourseAtClass.GetAll(h.store)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, ats)
}

// FindAtsByID
// @Summary find CourseAtClass record by id
// @Description find CourseAtClass record by id
// @Tags Course At Class
// @Accept	json
// @Produce  json
// @Param id path string true "id of row in CourseAtClass"
// @Success 200 {object} model.CourseAtClass	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/courseatclass/{id} [get]
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

// AddRelationshipCourseAtClass
// @Summary add new relationship between course and class
// @Description add new relationship between course and class
// @Tags Course At Class
// @Accept	json
// @Produce  json
// @Param course_id query string true "id of course"
// @Param class_id query string true "id of class"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.CourseAtClass	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/courseatclass/add [post]
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

// DeleteRelationshipByID
// @Summary delete relationship between course and class
// @Description delete relationship between course and class
// @Tags Course At Class
// @Accept	json
// @Produce  json
// @Param id query string true "id of row in table"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.CourseAtClass	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/courseatclass/id [delete]
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

	return c.JSON(http.StatusOK, ats)
}
