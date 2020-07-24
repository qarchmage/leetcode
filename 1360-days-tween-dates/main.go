package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31"))
}

func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse("2006-01-02", date1)
	d2, _ := time.Parse("2006-01-02", date2)
	r := 0
	r = int(d2.Sub(d1).Hours()) / 24
	if r < 0 {
		r = int(d1.Sub(d2).Hours()) / 24
	}
	return int(r)

}
