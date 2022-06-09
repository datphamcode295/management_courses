package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

// GetAllClassesHandler
// @Summary get all classes
// @Description get all classes
// @Tags Class
// @Accept	json
// @Produce  json
// @Success 200 {object} model.Class	"ok"
// @Router /api/v1/class/all [get]
func (h *Handler) GetAllClass(c echo.Context) error {

	sliceClasses, err := h.repo.Class.GetAllClasses(h.store)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, sliceClasses)
}

// AddClass
// @Summary add class
// @Description add class
// @Tags Class
// @Accept	json
// @Produce  json
// @Param name query string true "class name"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Class	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/class/add [post]
func (h *Handler) AddClass(c echo.Context) error {
	name := c.QueryParam("name")

	username := c.QueryParam("username")
	password := c.QueryParam("password")
	role, checkerr := h.GetRole(c, username, password)
	if checkerr != nil {
		return util.HandleError(c, checkerr)
	}
	if role != "admin" {
		return util.HandleError(c, errors.ErrUnAuthorized)
	}

	classParams := model.Class{Name: name}
	sliceClasses, err := h.repo.Class.PostClass(h.store, classParams)

	c.JSON(http.StatusOK, sliceClasses)
	if err != nil {
		return util.HandleError(c, err)
	}

	// h.store.DB().Create(&(classParams))
	return nil
}

// FindClassByID
// @Summary find class by id
// @Description find class by id
// @Tags Class
// @Accept	json
// @Produce  json
// @Param id path string true "id of class"
// @Success 200 {object} model.Class	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/class/id/{id} [get]
func (h *Handler) FindClassByID(c echo.Context) error {
	id := c.Param("id")
	class, err := h.repo.Class.GetClassByID(h.store, id)
	c.JSON(http.StatusOK, class)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return nil
}

// FindClassByName
// @Summary find class by name
// @Description find class by name
// @Tags Class
// @Accept	json
// @Produce  json
// @Param name path string true "class name"
// @Success 200 {object} model.Class	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/class/name/{name} [get]
func (h *Handler) FindClassByName(c echo.Context) error {
	name := c.Param("name")
	class, err := h.repo.Class.GetClassByName(h.store, name)
	c.JSON(http.StatusOK, class)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return nil
}

// FindAllStudents
// @Summary find all students in a class
// @Description find all students in a class
// @Tags Class
// @Accept	json
// @Produce  json
// @Param name query string true "class name"
// @Success 200 {object} []model.Client	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/class/students/{name} [get]
func (h *Handler) FindAllStudent(c echo.Context) error {
	name := c.Param("name")
	class, err := h.repo.Class.GetClassByName(h.store, name)
	if err != nil {
		return util.HandleError(c, err)
	}
	students, err := h.repo.Class.GetAllStudent(h.store, class.ID)

	c.JSON(http.StatusOK, students)
	if err != nil {
		return util.HandleError(c, err)
	}
	return nil

}

// FindAllCourses
// @Summary find all courses of a class
// @Description find all courses of a class
// @Tags Class
// @Accept	json
// @Produce  json
// @Param name path string true "class name"
// @Success 200 {object} []model.Course	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/class/course/{name} [post]
func (h *Handler) FindAllCourses(c echo.Context) error {
	name := c.Param("name")
	//get class id from name
	class, err := h.repo.Class.GetClassByName(h.store, name)
	if err != nil {
		return util.HandleError(c, err)
	}
	//get records(course_id, class_id)
	records, err := h.repo.CourseAtClass.GetByClassID(h.store, class.ID)

	if err != nil {
		return util.HandleError(c, err)
	}

	var classes []model.Course

	for _, e := range records {
		temp, err := h.repo.Course.GetCourseById(h.store, e.CourseId)
		if err != nil {
			return util.HandleError(c, err)
		}
		classes = append(classes, *temp)
	}

	c.JSON(http.StatusOK, classes)

	return nil

}

// DeleteClassByID
// @Summary delete class by id
// @Description delete class by id
// @Tags Class
// @Accept	json
// @Produce  json
// @Param id query string true "id of class"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Class	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/class/id [delete]
func (h *Handler) DeleteClassByID(c echo.Context) error {
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
	class, _ := h.repo.Class.GetClassByID(h.store, id)
	if class.Name == "" {
		return util.HandleError(c, errors.ErrClassNotfound)
	}

	_, err := h.repo.Class.DeleteByID(h.store, id)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, class)
}
