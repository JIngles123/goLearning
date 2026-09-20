package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println("current time:")
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println("\ncurrent time in UTC:")
	fmt.Println(t)

	week = 60 * 60 * 24 * 7 * 1e9 // 1000000000纳秒
	week_from_now := t.Add(week)
	fmt.Println("\nweek from now:")
	fmt.Println(week_from_now)

	fmt.Println("\ntime.RFC822:", t.Format(time.RFC822))
	fmt.Println("time.ANSIC:", t.Format(time.ANSIC))
	fmt.Println("\n21 Dec 2011 08:52:", t.Format("21 Dec 2011 08:52"))
	fmt.Println(t.Format("21 Dec 2011 08:52"))

	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}