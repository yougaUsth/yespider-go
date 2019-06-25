package api

import (
	"github.com/gin-gonic/gin"
	"yespider-go/spider/parsers"
	"yespider-go/spider/spiders"
	//"yespider-go/spider/spiders"
)

func GetTaskInfo(c *gin.Context) {
	//c.Params
	taskName := c.DefaultQuery("task_name", "")
	//
	articleConfigs := make([]parsers.ArticleResp, 2)
	articleConfigs[0] = parsers.ArticleResp{Name: "title", Xpath: "//div[@class='section-cnt-tit']"}
	articleConfigs[1] = parsers.ArticleResp{Name: "text", Xpath: "//div[@class='article-main']"}
	//
	newSpider := spiders.NewSpider("news", "test", "http://news.iqilu.com/shandong/", "//div[@class='mod-list']//a/@href", articleConfigs)
	newSpider.DealTask()

	c.JSON(200, gin.H{"msg": "Test get targets!" + taskName})
}
