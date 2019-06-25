package parsers

import (
	"fmt"
	"github.com/gocolly/colly"
	"gopkg.in/xmlpath.v2"
	"log"
	"net/http"
	"strings"
	"time"
)

// DOM树索引配置
type IndexResp struct {
	Url   string
	Xpath string
}

//DOM结构配置
type ArticleResp struct {
	Name  string
	Xpath string
}

type BaseParser struct {
	IndexResp   IndexResp
	ArticleList []ArticleResp
	ToDoList    []string
}

//解析索引URL
func (p *BaseParser) ParseIndex() {
	resp, err := http.Get(p.IndexResp.Url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	root, err := xmlpath.ParseHTML(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	path := xmlpath.MustCompile(p.IndexResp.Xpath)
	iter := path.Iter(root)

	var todoUrls []string
	for iter.Next() {
		xUrl := iter.Node().String()
		//TODO: 需要解决相对路径的拼接问题
		if strings.HasPrefix(xUrl, "http") {
			todoUrls = append(todoUrls, xUrl)
			//fmt.Println(xUrl)
		}
	}
	p.ToDoList = todoUrls

}

/**
解析字段
*/
func (p *BaseParser) ParserArticle() {
	for _, xUrl := range p.ToDoList {
		go func() {
			// 解析Article 详细配置
			for _, articleConf := range p.ArticleList {
				c := colly.NewCollector(colly.MaxDepth(1))

				//TODO: Article parser
				c.OnHTML(articleConf.Xpath, func(element *colly.HTMLElement) {
					articleText := element.Text
					fmt.Print(articleText)
				})

				c.OnRequest(func(request *colly.Request) {
					fmt.Println("Visiting", xUrl)
				})

				err := c.Visit(xUrl)
				if err != nil {
					//log.Fatalf("Raise a exception when visit %v", err)
					fmt.Printf("Raise a exception when visit %v", err)
				}
			}
		}()
		time.Sleep(5 * time.Second)
	}
}
