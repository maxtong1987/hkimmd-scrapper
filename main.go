package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getUrl(year, month, day int) string {
	return fmt.Sprintf("https://www.immd.gov.hk/hkt/stat_%02d%02d%02d.html", year, month, day)
}

func main() {
	url := getUrl(2021, 5, 31)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	doc.Find(".a").ChildrenFiltered(".hRight").Each(
		func(i int, element *goquery.Selection) {
			txt := element.Text()
			fmt.Println(txt)
		})
}
