package main

import (
	"fmt"
	"log"
	"net/http"
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

func toLocalDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func getDataByDate(year, month, day int) ([]int, error) {
	url := getUrl(year, month, day)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
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
	return numbers, nil
}

type RowData struct {
	Year    int
	Month   int
	Day     int
	Numbers []int
}

func main() {
	beginDate := toLocalDate(2021, 6, 1)
	endDate := toLocalDate(2021, 6, 27)

	numTable := make([]RowData, 0, 365)

	for date := beginDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
		row := RowData{
			Year:  date.Year(),
			Month: int(date.Month()),
			Day:   date.Day(),
		}

		numbers, err := getDataByDate(row.Year, row.Month, row.Day)
		if err != nil {
			log.Fatal(err)
		}
		row.Numbers = numbers
		numTable = append(numTable, row)
	}

	fmt.Print(numTable)
}
