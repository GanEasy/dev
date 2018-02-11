package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	mpoauth2 "github.com/chanxuehong/wechat.v2/mp/oauth2"
	"github.com/chanxuehong/wechat.v2/oauth2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	wxAppId           = "wxdd75ceb6bd54c179"               // 填上自己的参数
	wxAppSecret       = "3a71070c1ba4e56f47ee4657e917a4ef" // 填上自己的参数
	oauth2RedirectURI = "http://wdt.readfollow.com/sign"   // 填上自己的参数
	oauth2Scope       = "snsapi_userinfo"                  // 填上自己的参数
)

var (
	oauth2Endpoint oauth2.Endpoint = mpoauth2.NewEndpoint(wxAppId, wxAppSecret)
)

func main() {
	fmt.Println(http.ListenAndServe(":80", nil))

	e := echo.New()

	e.Use(middleware.CORS())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/login", Login)
	e.GET("/sign", Sign)
	e.Logger.Fatal(e.Start(":3345"))
}

//Sign 答卷完成记录提交 这里判断提交不太好用。直接用get了
func Sign(c echo.Context) (err error) {

	code := c.QueryParam("code")
	if code == "" {
		return errors.New("用户禁止授权")
	}

	queryState := c.QueryParam("state")
	if queryState == "" {
		return errors.New("state 参数为空")
	}

	oauth2Client := oauth2.Client{
		Endpoint: oauth2Endpoint,
	}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("token: %+v\r\n", token)

	userinfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		log.Println(err)
		return
	}
	return c.JSON(http.StatusOK, userinfo)
}

//Login 答卷完成记录提交 这里判断提交不太好用。直接用get了
func Login(c echo.Context) error {
	AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, "snsapi_base", "start_sign")
	return c.Redirect(http.StatusMovedPermanently, AuthCodeURL)
}
