package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getUrl(year, month, day int) string {
	return fmt.Sprintf("https://www.immd.gov.hk/hkt/stat_%02d%02d%02d.html", year, month, day)
}

func text2Int(text string) (int, error) {
	str := strings.ReplaceAll(text, ",", "")
	return strconv.Atoi(str)
}

func getDocFromUrl(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return goquery.NewDocumentFromReader(resp.Body)
}

func getRow(year, month, day int) (*Row, error) {
	url := getUrl(year, month, day)
	doc, err := getDocFromUrl(url)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	numbers := make([]int, 0, 8)
	doc.Find(".a").ChildrenFiltered(".hRight").Each(
		func(i int, element *goquery.Selection) {
			txt := element.Text()
			num, err := text2Int(txt)
			if err != nil {
				fmt.Print(err)
				return
			}
			numbers = append(numbers, num)
		})
	return &Row{
		Year:    year,
		Month:   month,
		Day:     day,
		Numbers: numbers,
	}, nil
}

func getRows(begin, end time.Time) ([]*Row, error) {
	rows := make([]*Row, 0, 1024)
	for date := begin; date.Before(end); date = date.AddDate(0, 0, 1) {
		row, err := getRow(date.Year(), int(date.Month()), date.Day())
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

var (
	beginDate dateVar
	endDate   dateVar
	csvFile   string
)

func init() {
	flag.Var(&beginDate, "b", "begin date (e.g. 2020-01-21)")
	flag.Var(&endDate, "e", "end date (e.g. 2020-01-27)")
	flag.StringVar(&csvFile, "f", "data.csv", "csv file in which you want to save passenger traffic data")
}

func main() {
	flag.Parse()
	f, err := os.Create(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows, err := getRows(beginDate.Time, endDate.Time)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		f.WriteString(row.toCSV() + "\n")
	}

	fmt.Println("Complete!")
}
