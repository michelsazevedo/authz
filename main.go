package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michelsazevedo/authz/api"
	"github.com/michelsazevedo/authz/config"
	"github.com/michelsazevedo/authz/domain"
	m "github.com/michelsazevedo/authz/middleware"
	"github.com/michelsazevedo/authz/repository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	conf, _ := config.NewConfig()
	userRepository := repository.NewUserRepository(conf)
	service := domain.NewUserService(userRepository)
	handler := api.NewHandler(service)
	settings := m.NewSettings(conf.Settings)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(settings.Settings)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("Method=" + c.Request().Method)

			return nil
		},
	}))

	users := e.Group("/users")
	users.POST("/signup", handler.SignUp)
	users.POST("/signin", handler.SignIn)

	auth := e.Group("/auth")
	auth.GET("/", handler.Auth)
	auth.POST("/refresh", handler.Refresh)

	e.GET("/healthz", func(c echo.Context) error {
		healthz := map[string]int{"status": 200}
		return c.JSON(http.StatusOK, healthz)
	})

	e.Logger.Fatal(e.Start(conf.Settings.Server.Port))
}
