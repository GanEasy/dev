package main

import "github.com/labstack/echo"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Route => handler
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!\n")
	// })
	e.Static("/", "src")
	// Start server
	e.Logger.Fatal(e.Start(":8552"))
}
