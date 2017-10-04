package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/GanEasy/InsectFigure/core"
	"github.com/disintegration/imaging"
	"github.com/labstack/echo"
)

// 接入微信接口服务
func api(c echo.Context) error {
	input := c.Param("url")
	uDec, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		log.Fatalln(err)
	}
	PrintHandler(string(uDec), c.Response().Writer, c.Request())

	var err2 error
	return err2
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func PrintHandler(u string, w http.ResponseWriter, r *http.Request) {

	imgname := GetMd5String(u)

	imgpath := fmt.Sprintf("file/%v.jpg", imgname)

	// 如果本地服务器不存在缓存，再去拿
	_, err := os.Stat(imgpath)
	if os.IsNotExist(err) {
		core.SaveImg(u, imgpath)
	}

	src, err := imaging.Open(imgpath)
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}
	// src = imaging.Resize(src, 256, 0, imaging.Lanczos)
	src = imaging.Resize(src, 350, 0, imaging.Lanczos)
	src = imaging.CropAnchor(src, 350, 200, imaging.Center)

	// http.ServeFile(w, r, imgpath)

	// w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Content-Type", "image/jpeg")

	jpeg.Encode(w, src, &jpeg.Options{Quality: 80})
	// gif.Encode(w, src, &gif.Options{NumColors: 256})
	// png.Encode(w, src)

	// imgpath := "./13.jpg"
	// core.SaveImg("http://localhost:1323/2015052700230736104.jpg", imgpath)

	// // if !FileExist(imgpath) {
	// // 	w.Write([]byte("Error:Image Not Found."))
	// // 	return
	// // }
	// http.ServeFile(w, r, imgpath)
	// os.Remove(imgpath)
}

func main() {
	e := echo.New()

	// Handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/:url", api)

	// Handler
	e.GET("/:url/:param", func(c echo.Context) error {
		input := c.Param("url")

		// input = "http://mmbiz.qpic.cn/mmbiz_jpg/Z8SUoc8pJqdBfxCtd51ibGNr7IOXNI4DuUVbpToIqdhZUibOYDmW0S8nCGchoExiaMIPJ8oaMsXB7KSyKNcsVjibBg/0?wx_fmt=jpeg"
		// uEnc := base64.URLEncoding.EncodeToString([]byte(input))
		// aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL1o4U1VvYzhwSnFkQmZ4Q3RkNTFpYkdOcjdJT1hOSTREdVVWYnBUb0lxZGhaVWliT1lEbVcwUzhuQ0djaG9FeGlhTUlQSjhvYU1zWEI3S1N5S05jc1ZqaWJCZy8wP3d4X2ZtdD1qcGVn

		// fmt.Println(string(uEnc))

		uDec, err := base64.URLEncoding.DecodeString(input)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(uDec))
		// fmt.Println(string(uEnc))
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.File("/favicon.ico", "images/favicon.ico")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
