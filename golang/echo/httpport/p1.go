package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "m1 Hello, World!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start("zj.93zp.dev:1323"))
}
