package spiders

import (
	uuid "github.com/satori/go.uuid"
	"yespider-go/spider/parsers"
)

type BaseSpider struct {
	TaskId        string                // UUID标识任务
	TaskName      string                // 任务名
	ParserName    string                // 解析器名
	StartUrl      string                // 起始路径
	IndexXpath    string                // 索引list Xpath
	ArticleXpaths []parsers.ArticleResp // 字段 Xpath
}

func NewSpider(parser, taskName, startUrl, indexXpath string, articleConfig []parsers.ArticleResp) *BaseSpider {

	uid := uuid.Must(uuid.NewV4()).String()

	return &BaseSpider{
		TaskId:        uid,
		TaskName:      taskName,
		ParserName:    parser,
		StartUrl:      startUrl,
		IndexXpath:    indexXpath,
		ArticleXpaths: articleConfig,
	}
}

func (s *BaseSpider) parserArticle() {

}

//开心就好
func (s *BaseSpider) DealTask() {
	parser := parsers.BaseParser{
		IndexResp: struct {
			Url   string
			Xpath string
		}{Url: s.StartUrl, Xpath: s.IndexXpath},
		ArticleList: s.ArticleXpaths,
	}
	parser.ParseIndex()
	parser.ParserArticle()
}
