package main

import (
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yizenghui/reader"
	"gopkg.in/russross/blackfriday.v2"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func Article(c echo.Context) error {
	a := &Data
	return c.Render(http.StatusOK, "article", a)
}

var Data struct {
	Title string
	Info  string
	Media string
	URL   string
	PubAt string
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e := echo.New()

	e.Renderer = t
	// Middleware

	e.GET("/hello", Hello)
	// Route => handler
	e.GET("/", func(c echo.Context) error {

		urlStr := c.QueryParam("url")

		if urlStr == "" {
			return c.String(http.StatusOK, "url")
		}
		md, _ := reader.Read(urlStr)

		input := []byte(md)

		unsafe := blackfriday.Run(input)
		content := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

		html := fmt.Sprintf(`<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<link rel="preload" href="https://yize.gitlab.io/css/main.css" as="style" />
			%v`, string(content[:]))

		return c.HTML(http.StatusOK, html)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
