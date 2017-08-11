package main

import (
	"fmt"
	"log"
	//"io/ioutil"
	"net/http"
	//	"net/http/cookiejar"
	"net/url"
	"sync"
	// "bytes"

	"github.com/PuerkitoBio/goquery"

	// "golang.org/x/text/encoding/simplifiedchinese"
	// "golang.org/x/text/transform"
	"github.com/axgle/mahonia"
)

type Jar struct {
	lk      sync.Mutex
	cookies map[string][]*http.Cookie
}

func NewJar() *Jar {
	jar := new(Jar)
	jar.cookies = make(map[string][]*http.Cookie)
	return jar
}

// SetCookies handles the receipt of the cookies in a reply for the
// given URL.  It may or may not choose to save the cookies, depending
// on the jar's policy and implementation.
func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.lk.Lock()
	jar.cookies[u.Host] = cookies
	jar.lk.Unlock()
}

// Cookies returns the cookies to send in a request for the given URL.
// It is up to the implementation to honor the standard cookie use
// restrictions such as in RFC 6265.
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies[u.Host]
}

func main() {

	username := "zenghui"
	password := "000000"
	// 编码转换器 解释 gbk 编码内容
	enc := mahonia.NewDecoder("gbk")

	jar := NewJar()
	client := http.Client{Jar: jar}

	resp, err := client.Get("http://192.168.1.121:8088/admin_common/login")
	if err != nil {
		//	 handle error
		fmt.Printf("发出请求失败!")
	}

	form, err := goquery.NewDocumentFromReader(resp.Body)
	//	defer
	resp.Body.Close()

	if err != nil {
		fmt.Printf("获取登录页信息失败!")
	}

	//	fmt.Printf("获取登录页Token")
	token, _ := form.Find(".login").Find("input[name='DA__FORMTOKEN']").Eq(0).Attr("value")
	//	fmt.Printf("Token %s\n", token)

	data := make(url.Values)
	data["DA__FORMTOKEN"] = []string{token}
	data["username"] = []string{username}
	data["password"] = []string{password}
	resp, err = client.PostForm("http://192.168.1.121:8088/admin_common/login", data)
	if err != nil {
		fmt.Printf("登录时提交数据异常")
	}
	formPost, err := goquery.NewDocumentFromReader(resp.Body)

	resp.Body.Close()

	postMsg := formPost.Find(".box_bd").Eq(0).Text()
	fmt.Printf("使用 %s&%s 进行登录 \n", username, password)

	// 返回信息是GBK 需要进行转码

	fmt.Println("Login Msg ", enc.ConvertString(postMsg))

	resp, err = client.Get("http://192.168.1.121:8088/performance_staff/daka")
	if err != nil {
		//	 handle error
		fmt.Printf("获取打卡页面信息")
	}
	//	defer
	Doc, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	Time := Doc.Find("td#time").Eq(0).Text()

	fmt.Printf("Time %s\n", Time)

	// work_situation
	// TODO 这里需要做一些逻辑判断、如果已经打卡，那么提示已经打卡上班，无需重复操作，否则打卡成功

	sign_token, _ := Doc.Find("input[name='DA__FORMTOKEN']").Eq(0).Attr("value")
	// fmt.Printf("Sign Token %s\n", sign_token)

	//	fmt.Printf("打卡描述：(上班按 Enter、下班输入内容加Enter 换行使用<br />标签)\n")

	signData := make(url.Values)
	signData["DA__FORMTOKEN"] = []string{sign_token}
	signData["work_situation"] = []string{"1"}
	signData["work_time_type"] = []string{"1"}

	// 打卡日报
	//	daily_report := ""
	//	_, err3 := fmt.Scanln(&daily_report)
	//	if nil == err3 {
	//		//fmt.Printf("daily_report=%s", daily_report)
	//		signData["daily_report"] = []string{daily_report}
	//	}

	resp, err = client.PostForm("http://192.168.1.121:8088/performance_staff/daka", signData)
	signPost, err := goquery.NewDocumentFromReader(resp.Body)

	resp.Body.Close()

	signMsg := signPost.Find(".box_bd").Eq(0).Text()
	// 返回信息是GBK 需要进行转码 TODO
	fmt.Printf("Sign Msg %s\n", enc.ConvertString(signMsg))

	// 防空洞
	tips := "ok"
	_, err4 := fmt.Scanln(&tips)
	if nil == err4 {
		fmt.Println("tips=%", tips)
	}

}
