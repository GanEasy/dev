package main

import (
	"github.com/sundy-li/html2article"
)

func main() {
	urlStr := "https://www.leiphone.com/news/201602/DsiQtR6c1jCu7iwA.html"
	urlStr = "http://book.zongheng.com/chapter/688697/38691248.html"
	urlStr = "http://www.76wx.com/book/190/3202125.html"
	urlStr = "http://book.zongheng.com/chapter/688697/38712592.html"

	urlStr = "https://mp.weixin.qq.com/s?__biz=MzA5OTgyOTcyNw==&mid=2650309323&idx=2&sn=e304ad167e2f7027a9e8db61d54260ce&chksm=88f0707bbf87f96dd2e3a949edd288e9cb42729d1c61f30a26adb6aff9d1ebc7602189b41a01&scene=27#wechat_redirect"
	ext, err := html2article.NewFromUrl(urlStr)
	if err != nil {
		panic(err)
	}
	article, err := ext.ToArticle()
	if err != nil {
		panic(err)
	}
	println("article title is =>", article.Title)
	println("article publishtime is =>", article.Publishtime)
	println("article content is =>", article.Content)

	//parse the article to be readability
	// article.Readable(urlStr)
	// println("read=>", article.ReadContent)
}
