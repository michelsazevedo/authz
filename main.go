package main

import (
	"github.com/labstack/echo/v4"
	"github.com/michelsazevedo/authz/config"
)

func main() {
	e := echo.New()
	conf, _ := config.NewConfig()

	e.GET("/healthz", func(c echo.Context) error {
		healthz := map[string]int{"status": 200}
		return c.JSON(200, healthz)
	})

	e.Logger.Fatal(e.Start(conf.Settings.Server.Port))
}
