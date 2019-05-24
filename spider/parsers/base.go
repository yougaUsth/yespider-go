package parsers

import (
	"fmt"
	"github.com/xmlpath"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Url   string
	Xpath string
}

type BaseParser struct {
	IndexResp    Response
	ArticleXpath []Response
	ToDoList     []string
}

/**
解析索引url
*/
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
			fmt.Println(xUrl)
		}
	}
	p.ToDoList = todoUrls

}

/**
解析字段
*/
func (p *BaseParser) ParserArticle() {
	for _, xUrl := range p.ToDoList {
		resp, err := http.Get(xUrl)
		if err != nil {
			fmt.Print(err)
		}
		root, _ := xmlpath.Parse(resp.Body)
		for _, article := range p.ArticleXpath {
			path := xmlpath.MustCompile(article.Xpath)
			iter := path.Iter(root)
			//iter := path.Iter()
			for iter.Next() {
				articleString := iter.Node().String()
				fmt.Print(articleString)
			}
		}
		resp.Body.Close()
		fmt.Print(xUrl)
	}
}
