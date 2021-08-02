package api

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {

	godotenv.Load()

	return func(c echo.Context) error {
		key := os.Getenv("API_KEY")
		logrus.Debugf("KEY: %v", key)
		context := context.Background()
		service, err := youtube.NewService(context, option.WithAPIKey(key))
		if err != nil {
			logrus.Fatalf("Error creating new Youtube service: %v", err)
		}

		call := service.Videos.List([]string{"id", "snippet"}).Chart("mostPopular").MaxResults(3)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
