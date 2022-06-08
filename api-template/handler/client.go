package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllClient(c echo.Context) error {

	classes, err := h.repo.Client.GetAllClients(h.store)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, classes)
}

func (h *Handler) AddClient(c echo.Context) error {
	password := c.QueryParam("password")
	username := c.QueryParam("username")
	classid := c.QueryParam("classid")
	role := c.QueryParam("role")
	name := c.QueryParam("name")

	//check double username
	user, _ := h.repo.Client.GetByUsername(h.store, username)

	if user.Password != "" {
		return util.HandleError(c, errors.ErrClientAlreadyExits)
	}

	param := model.Client{Username: username, Password: password, ClassID: classid, Role: role, Name: name}
	client, err := h.repo.Client.PostClient(h.store, param)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, client)
}

func (h *Handler) DeleteClient(c echo.Context) error {
	id := c.Param("id")

	//check user exist
	user, _ := h.repo.Client.GetByID(h.store, id)

	if user.Username == "" {
		return util.HandleError(c, errors.ErrClientDoesNotExits)
	}

	_, err := h.repo.Client.Delete(h.store, id)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, "Delete successfully the clinent with id"+id)
}

func (h *Handler) UpdateClientPasswordByID(c echo.Context) error {
	id := c.QueryParam("id")
	password := c.QueryParam("password")

	client, err := h.repo.Client.UpdatePassword(h.store, id, password)
	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, client)
}

func (h *Handler) GetRole(c echo.Context, username string, password string) (string, error) {
	client, err := h.repo.Client.CheckingLogin(h.store, username, password)
	if err != nil {
		return "", errors.ErrAccountNotfound
	}

	return client.Role, nil
}

func (h *Handler) Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	client, err := h.repo.Client.CheckingLogin(h.store, username, password)
	if err != nil {
		return util.HandleError(c, errors.ErrAccountNotfound)
	}
	return c.JSON(http.StatusOK, client)
}
