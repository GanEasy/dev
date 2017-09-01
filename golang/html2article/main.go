package main

import (
	"github.com/sundy-li/html2article"
)

func main() {
	urlStr := "https://www.leiphone.com/news/201602/DsiQtR6c1jCu7iwA.html"
	urlStr = "http://book.zongheng.com/chapter/688697/38691248.html"
	urlStr = "http://www.76wx.com/book/190/3202125.html"
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
	article.Readable(urlStr)
	println("read=>", article.ReadContent)
}