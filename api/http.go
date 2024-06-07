package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michelsazevedo/authz/config/jwt"
	"github.com/michelsazevedo/authz/domain"
)

type AuthzHandler interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
	Auth(c echo.Context) error
	Refresh(c echo.Context) error
}

type handler struct {
	userService domain.UserService
}

func NewHandler(userService domain.UserService) AuthzHandler {
	return &handler{userService: userService}
}

func (h *handler) SignUp(c echo.Context) error {
	user := &domain.User{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	if err := h.userService.SignUp(c.Request().Context(), user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *handler) SignIn(c echo.Context) error {
	signInParams := &domain.SignInParams{}

	if err := c.Bind(signInParams); err != nil {
		return c.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	user, err := h.userService.SignIn(c.Request().Context(), signInParams)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *handler) Auth(c echo.Context) error {
	token := jwt.FromContext(c.Request().Context())

	return c.JSON(http.StatusOK, token)
}

func (h *handler) Refresh(c echo.Context) error {
	token := jwt.FromContext(c.Request().Context())
	user, err := h.userService.Refresh(c.Request().Context(), token)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity))
	}

	return c.JSON(http.StatusCreated, user)
}
