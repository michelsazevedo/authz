package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michelsazevedo/authz/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, _ := config.NewConfig()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	e := echo.New()
	e.Use(middleware.Recover())
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

	e.GET("/healthz", func(c echo.Context) error {
		healthz := map[string]int{"status": 200}
		return c.JSON(200, healthz)
	})

	e.Logger.Fatal(e.Start(conf.Settings.Server.Port))
}
