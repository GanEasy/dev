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
	wxAppId           = "wx702b93aef72f3549"               // 填上自己的参数
	wxAppSecret       = "8b69f45fc737a938cbaaffc05b192394" // 填上自己的参数
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
	e.GET("/user", User)
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

	if state == "start_sign" {
		// wdt
		url := fmt.Sprintf(`http://wdt.readfollow.com/?token=%v&openid=%v`, token.AccessToken, token.OpenId)
		return c.Redirect(http.StatusMovedPermanently, url)
	}
	//
	return c.JSON(http.StatusOK, token)
}

//Login 答卷完成记录提交 这里判断提交不太好用。直接用get了
func Login(c echo.Context) error {
	AuthCodeURL := mpoauth2.AuthCodeURL(wxAppId, oauth2RedirectURI, "snsapi_base", "start_sign")
	return c.Redirect(http.StatusMovedPermanently, AuthCodeURL)
}

//User 获取用户信息
func User(c echo.Context) (err error) {

	token := c.QueryParam("token")
	if token == "" {
		return errors.New("token 参数为空")
	}

	openid := c.QueryParam("openid")
	if openid == "" {
		return errors.New("openid 参数为空")
	}

	userinfo, err := mpoauth2.GetUserInfo(token, openid, "", nil)
	if err != nil {
		log.Println(err)
		return
	}
	return c.JSON(http.StatusOK, userinfo)
}
