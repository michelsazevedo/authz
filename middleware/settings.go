package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/michelsazevedo/authz/config"
)

type Settings struct {
	settings config.Settings
}

func NewSettings(settings config.Settings) *Settings {
	return &Settings{settings: settings}
}

func (m *Settings) Settings(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetRequest(c.Request().WithContext(config.ToContext(c.Request().Context(), m.settings)))

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
