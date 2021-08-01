package main

import (
	"youtube-manager-go/web/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
