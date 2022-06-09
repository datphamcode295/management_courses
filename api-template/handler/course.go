package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

// GetAllCourses
// @Summary get all courses
// @Description get all courses
// @Tags Course
// @Accept	json
// @Produce  json
// @Success 200 {object} model.Course	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/course/all [get]
func (h *Handler) GetAllCourse(c echo.Context) error {
	sliceClasses, err := h.repo.Course.GetAllCourses(h.store)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, sliceClasses)

}

// AddCourse
// @Summary add course
// @Description add course
// @Tags Course
// @Accept	json
// @Produce  json
// @Param name query string true "name of course"
// @Param lecturer query string true "lecturer of course"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Course	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/course/add [post]
func (h *Handler) AddCourse(c echo.Context) error {

	name := c.QueryParam("name")
	lecturer := c.QueryParam("lecturer")

	username := c.QueryParam("username")
	password := c.QueryParam("password")
	role, checkerr := h.GetRole(c, username, password)
	if checkerr != nil {
		return util.HandleError(c, checkerr)
	}
	if role != "admin" {
		return util.HandleError(c, errors.ErrUnAuthorized)
	}

	param := model.Course{Name: name, Lecturer: lecturer}
	course, err := h.repo.Course.PostCourse(h.store, param)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, course)
}

// DeleteCourseByID
// @Summary delete course by id
// @Description delete course by id
// @Tags Course
// @Accept	json
// @Produce  json
// @Param id query string true "id of course"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Course	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/course/id [delete]
func (h *Handler) DeleteCourseByID(c echo.Context) error {
	id := c.QueryParam("id")

	username := c.QueryParam("username")
	password := c.QueryParam("password")
	//check role of the user
	role, checkerr := h.GetRole(c, username, password)
	if checkerr != nil {
		return util.HandleError(c, checkerr)
	}
	if role != "admin" {
		return util.HandleError(c, errors.ErrUnAuthorized)
	}

	//check course id exits
	course, _ := h.repo.Course.GetCourseById(h.store, id)
	if course.Name == "" {
		return util.HandleError(c, errors.ErrCourseNotfound)
	}

	_, err := h.repo.Course.DeleteByID(h.store, id)

	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, course)
}

// FindAllClasses
// @Summary find all classes of a course
// @Description find all classes of a course
// @Tags Course
// @Accept	json
// @Produce  json
// @Param id path string true "id of course"
// @Success 200 {object} []model.Class	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/course/allClasses/{id} [get]
func (h *Handler) FindAllClasses(c echo.Context) error {
	id := c.Param("id")
	//get records(classid, userid)
	records, err := h.repo.CourseAtClass.GetByCourseID(h.store, id)

	if err != nil {
		return util.HandleError(c, err)
	}
	var classes []model.Class

	for _, e := range records {
		temp, err := h.repo.Class.GetClassByID(h.store, e.ClassId)
		if err != nil {
			return util.HandleError(c, err)
		}
		classes = append(classes, *temp)
	}

	return c.JSON(http.StatusOK, classes)
}
