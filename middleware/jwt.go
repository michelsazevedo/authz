package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/michelsazevedo/authz/config"
	auth "github.com/michelsazevedo/authz/config/jwt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/michelsazevedo/authz/domain"
	"github.com/rs/zerolog/log"
)

type JwtAuth struct {
	JwtClaims domain.JwtToken
}

func NewJwtJwtAuth() *JwtAuth {
	return &JwtAuth{}
}

func (m *JwtAuth) JwtAuthenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := strings.TrimSpace(c.Request().Header.Get("access-token"))

		if accessToken == "" {
			if err := c.JSON(http.StatusForbidden, http.StatusText(http.StatusBadRequest)); err != nil {
				return err
			}
		}

		_, err := jwt.ParseWithClaims(accessToken, &m.JwtClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.FromContext(c.Request().Context()).Secret), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				log.Error().Str("URI", "/auth/").Int("status", 403).Msg(err.Error())
				if err := c.JSON(http.StatusForbidden, http.StatusText(http.StatusForbidden)); err != nil {
					return err
				}

				return echo.NewHTTPError(http.StatusUnauthorized, err.Error()).SetInternal(err)
			}

			if errors.Is(err, jwt.ErrSignatureInvalid) {
				log.Error().Str("URI", "/auth/").Int("status", 401).Msg(err.Error())
				if err := c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)); err != nil {
					return err
				}

				return echo.NewHTTPError(http.StatusUnauthorized, err.Error()).SetInternal(err)
			}

			log.Error().Str("URI", "/auth/").Int("status", 401).Msg(err.Error())
			if err := c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)); err != nil {
				return err
			}

			return echo.NewHTTPError(http.StatusUnauthorized, err.Error()).SetInternal(err)
		}

		c.SetRequest(c.Request().WithContext(auth.ToContext(c.Request().Context(), m.JwtClaims)))

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
