package spiders

import (
	uuid "github.com/satori/go.uuid"
	"yespider-go/spider/parsers"
)

type BaseSpider struct {
	TaskId        string
	TaskName      string
	Parser        string
	StartUrl      string
	IndexXpath    string
	ArticleXpaths []ArticleXpath
}

type ArticleXpath struct {
	Xpath string
	Name  string
}

func NewSpider(parser, taskName, startUrl, indexXpath string) *BaseSpider {

	uid := uuid.Must(uuid.NewV4()).String()

	return &BaseSpider{
		TaskId:     uid,
		TaskName:   taskName,
		Parser:     parser,
		StartUrl:   startUrl,
		IndexXpath: indexXpath,
	}
}

func (s *BaseSpider) parserArticle() {

}

/**
处理任务
*/
func (s *BaseSpider) DealTask() {
	parser := parsers.BaseParser{
		IndexResp: struct {
			Url   string
			Xpath string
		}{Url: s.StartUrl, Xpath: s.IndexXpath},
	}
	parser.ParseIndex()
	parser.ParserArticle()
}
