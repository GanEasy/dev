package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feed.williamlong.info/")
	// feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)
	for k, item := range feed.Items {
		fmt.Println(k, item.Title, item.Link)
	}
}
