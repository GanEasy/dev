package reader

import "github.com/PuerkitoBio/goquery"
import "fmt"

// GetContent 获取前100粉丝级别
func GetContent(url string) {
	// id
	g, e := goquery.NewDocument(url)
	if e != nil {

		fmt.Println(e)
	} else {
		fmt.Println("txt", g.Text())
	}
}
