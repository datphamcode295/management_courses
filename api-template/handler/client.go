package handler

import (
	"net/http"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
)

// GetAllClients
// @Summary Get all clients information
// @Description Get all clients information
// @Tags Client
// @Accept	json
// @Produce  json
// @Success 200 {object} model.Client	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/client/all [get]
func (h *Handler) GetAllClient(c echo.Context) error {

	classes, err := h.repo.Client.GetAllClients(h.store)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, classes)
}

// AddClient
// @Summary Add all clients information
// @Description Add all clients information
// @Tags Client
// @Accept	json
// @Produce  json
// @Param name query string true "name of user"
// @Param role query string true "role of user"
// @Param classid query string true "id of class"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Client	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/client/ [post]
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

// DeleteClient
// @Summary Delete clients information
// @Description Delete clients information
// @Tags Client
// @Accept	json
// @Produce  json
// @Param id query string true "if of client"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Client	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/client/id [delete]
func (h *Handler) DeleteClient(c echo.Context) error {
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

	//check user exist
	user, _ := h.repo.Client.GetByID(h.store, id)

	if user.Username == "" {
		return util.HandleError(c, errors.ErrClientDoesNotExits)
	}

	_, err := h.repo.Client.Delete(h.store, id)

	if err != nil {
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}

// UpdatePassword
// @Summary update password
// @Description update password
// @Tags Client
// @Accept	json
// @Produce  json
// @Param id query string true "id of user need to change password"
// @Param newpass query string true "new password"
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Client	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/client/login [get]
func (h *Handler) UpdateClientPasswordByID(c echo.Context) error {
	id := c.QueryParam("id")
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	newpassword := c.QueryParam("newpassword")

	//check role of the user
	role, checkerr := h.GetRole(c, username, password)
	if checkerr != nil {
		return util.HandleError(c, checkerr)
	}

	u, _ := h.repo.Client.GetByID(h.store, id)
	if role != "admin" {

		if u.Username != username {
			return util.HandleError(c, errors.ErrUnAuthorized)
		}
	}

	_, err := h.repo.Client.UpdatePassword(h.store, id, newpassword)
	if err != nil {
		return util.HandleError(c, err)
	}
	u.Password = newpassword
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) GetRole(c echo.Context, username string, password string) (string, error) {
	client, err := h.repo.Client.CheckingLogin(h.store, username, password)
	if err != nil {
		return "", errors.ErrAccountNotfound
	}

	return client.Role, nil
}

// Login
// @Summary Login
// @Description Login
// @Tags Client
// @Accept	json
// @Produce  json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {object} model.Client	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/client/login [post]
func (h *Handler) Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	client, err := h.repo.Client.CheckingLogin(h.store, username, password)
	if err != nil {
		return util.HandleError(c, errors.ErrAccountNotfound)
	}
	return c.JSON(http.StatusOK, client)
}
