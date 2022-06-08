package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllClass(c echo.Context) error {

	sliceClasses, err := h.repo.Class.GetAllClasses(h.store)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, sliceClasses)
}

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
func (h *Handler) FindClassByID(c echo.Context) error {
	id := c.Param("id")
	class, err := h.repo.Class.GetClassByID(h.store, id)
	c.JSON(http.StatusOK, class)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return nil
}

func (h *Handler) FindClassByName(c echo.Context) error {
	name := c.Param("name")
	class, err := h.repo.Class.GetClassByName(h.store, name)
	c.JSON(http.StatusOK, class)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return nil
}

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
	c.JSON(http.StatusOK, "delete successfully the class with id: "+id)

	return nil
}
