package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func readHtml() *goquery.Document {
	doc, _ := goquery.NewDocument(os.Args[2])
	return doc
}

func readHtmlLocal() *goquery.Document {
	fileInfos, _ := ioutil.ReadFile("./sample.html")
	stringReader := strings.NewReader(string(fileInfos))
	doc, _ := goquery.NewDocumentFromReader(stringReader)
	return doc
}

func download(url string, path string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("ファイルダウンロードエラー:", url)
	}
	body, _ := ioutil.ReadAll(response.Body)
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	file.Write(body)
}

func main() {
	num := os.Args[1]
	os.Mkdir(num, 0755)
	fmt.Println("mkdir " + num)

	doc := readHtml()
	// doc := readHtmlLocal()

	// 調整するときにはここのマッチングを変える
	doc.Find("section img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("data-src")

		path := num + "/" + filepath.Base(url)
		fmt.Println(path)
		download(url, path)
	})
}
