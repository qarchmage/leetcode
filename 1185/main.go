package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
}
func dayOfTheWeek(day int, month int, year int) string {
	format := "01/02/2006"
	t := fmt.Sprintf("%.2d/%.2d/%d", month, day, year)
	r, _ := time.Parse(format, t)
	return r.Weekday().String()
}
