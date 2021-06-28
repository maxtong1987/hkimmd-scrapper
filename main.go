package main

import (
	"fmt"
	"io"
	"net/http"
)

func getUrl(year, month, day int) string {
	return fmt.Sprintf("https://www.immd.gov.hk/hkt/stat_%02d%02d%02d.html", year, month, day)
}

func main() {
	url := getUrl(2021, 5, 1)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(string(body))
}
