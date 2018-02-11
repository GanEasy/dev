package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	i()
}

func i() {
	// Server
	e := echo.New()

	e.Any("/api/*", func(c echo.Context) (err error) {

		req := c.Request()
		res := c.Response()

		cli := &http.Client{}
		body := make([]byte, 0)
		n, err := io.ReadFull(req.Body, body)
		if err != nil {
			io.WriteString(res, "Request Data Error")
			return
		}
		reqUrl := "http://mm.93zp.dev" + req.URL.Path

		fmt.Println(reqUrl)
		req2, err := http.NewRequest(req.Method, reqUrl, strings.NewReader(string(body)))
		if err != nil {
			io.WriteString(res, "Request Error")
			return
		}
		// set request content type
		contentType := req.Header.Get("Content-Type")
		req2.Header.Set("Content-Type", contentType)
		// request
		rep2, err := cli.Do(req2)
		if err != nil {
			io.WriteString(res, "Not Found!")
			return
		}
		defer rep2.Body.Close()
		n, err = io.ReadFull(rep2.Body, body)
		if err != nil {
			io.WriteString(res, "Request Error")
			return
		}
		// set response header
		for k, v := range rep2.Header {
			res.Header().Set(k, v[0])
		}
		io.WriteString(res, string(body[:n]))

		return
	})

	e.Any("/*", func(c echo.Context) (err error) {

		return c.String(http.StatusOK, "Website")
		// req := c.Request()
		// res := c.Response()
		// host := hosts[req.Host]

		// if host == nil {
		// 	err = echo.ErrNotFound
		// } else {
		// 	host.Echo.ServeHTTP(res, req)
		// }

		// return
	})
	e.Logger.Fatal(e.Start(":1323"))
}
