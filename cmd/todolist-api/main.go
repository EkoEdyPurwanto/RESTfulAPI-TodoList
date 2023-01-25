package main

import (
	"LearnECHO/internal/config"
	"LearnECHO/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	log := e.Logger

	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed when parsing config: %v", err)
	}

	database.ConnectDB(&cfg)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
