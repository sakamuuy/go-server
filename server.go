package main

import (
	"youtube-manager-go/web/routes"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter((&logrus.JSONFormatter{}))
}

func main() {
	app := echo.New()

	// logger := middleware.LoggerWithConfig(middleware.LoggerConfig{})
	// app.Use(logger)

	routes.Init(app)

	app.Logger.Fatal(app.Start(":8080"))
}
