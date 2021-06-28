package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.immd.gov.hk/hkt/stat_20210531.html")
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
