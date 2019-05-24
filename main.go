package main

import (
	"yspider/spider/spiders"
)

func Fibnacci() func() (ret int) {
	a, b := 0, 1
	return func() (ret int) {
		ret = b
		a, b = b, a+b
		return ret
	}
}

func main() {
	s := spiders.NewSpider("news", "test", "http://news.iqilu.com/shandong/", "//div[@class='mod-list']//a/@href")
	s.DealTask()
	//fmt.Print(s)
	//time.Sleep(5 * time.Second)
}
