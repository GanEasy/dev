package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Logger.Fatal(e.Start(":8005"))
}
