package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yizenghui/reader"
)

func main() {

	e := echo.New()
	// Middleware

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		urlStr := c.QueryParam("url")

		if urlStr == "" {
			return c.String(http.StatusOK, "url")
		}
		md, _ := reader.Read(urlStr)

		return c.String(http.StatusOK, md)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
