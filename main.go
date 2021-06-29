package main

import (
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

func toLocalDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
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

// Row stores one Row of passenger traffic data
type Row struct {
	Year    int
	Month   int
	Day     int
	Numbers []int
}

func (r *Row) toCSV() string {
	s := make([]string, 0, 9)
	date := fmt.Sprintf("%d-%02d-%02d", r.Year, r.Month, r.Day)
	s = append(s, date)
	for _, num := range r.Numbers {
		numStr := strconv.Itoa(num)
		s = append(s, numStr)
	}
	return strings.Join(s, ",")
}

func main() {
	beginDate := toLocalDate(2021, 6, 1)
	endDate := toLocalDate(2021, 6, 27)
	csvFile := "data.csv"

	f, err := os.Create(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows, err := getRows(beginDate, endDate)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range rows {
		f.WriteString(row.toCSV() + "\n")
	}

	fmt.Print("Complete!")
}
