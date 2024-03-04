package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Jai Guru Dev !!")
	log.Println("Jai Ho Ji")
	app := echo.New()
	app.GET("/", MyHandler)
	app.Logger.Fatal(app.Start(":9999"))
}

func MyHandler(c echo.Context) error {
	c.Response().Writer.WriteHeader(http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return My().Render(c.Request().Context(), c.Response().Writer)
}
