package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Row stores one Row of passenger traffic data
type Row struct {
	Year    int
	Month   int
	Day     int
	Numbers []int
}

// ToCSV converts Row into csv format string
func (r *Row) ToCSV() string {
	s := make([]string, 0, 9)
	date := fmt.Sprintf("%d-%02d-%02d", r.Year, r.Month, r.Day)
	s = append(s, date)
	for _, num := range r.Numbers {
		numStr := strconv.Itoa(num)
		s = append(s, numStr)
	}
	return strings.Join(s, ",")
}
