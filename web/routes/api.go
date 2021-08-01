package routes

import (
	"youtube-manager-go/web/api"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	apiG := e.Group("/api")
	{
		apiG.GET("/popular", api.FetchMostPopularVideos())
	}
}
