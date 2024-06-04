package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthzHandler interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
	Auth(c echo.Context) error
	Refresh(c echo.Context) error
}

type handler struct {
}

func NewHandler() AuthzHandler {
	return &handler{}
}

func (h *handler) SignUp(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]bool{"success": true})
}

func (h *handler) SignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func (h *handler) Auth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func (h *handler) Refresh(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]bool{"success": true})
}
