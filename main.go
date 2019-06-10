package main

import (
	"time"
	"yespider-go/spider/spiders"
)


func main() {
	s := spiders.NewSpider("news", "test", "http://news.iqilu.com/shandong/", "//div[@class='mod-list']//a/@href")
	s.DealTask()
	//fmt.Print(s)
	time.Sleep(5 * time.Second)

}
