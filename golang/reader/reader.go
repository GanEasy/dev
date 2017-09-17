package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
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

		input := []byte(md)
		unsafe := blackfriday.MarkdownCommon(input)
		content := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

		html := fmt.Sprintf(`<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<link rel="preload" href="https://yize.gitlab.io/css/main.css" as="style" />
			%v`, string(content[:]))

		return c.HTML(http.StatusOK, html)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
