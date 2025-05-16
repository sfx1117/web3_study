package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()
	p(t)
	p(t.Format(time.RFC3339))

	t1, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if err != nil {
		panic(err)
	}
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, err := time.Parse(form, "8 41 PM")
	if err != nil {
		panic(err)
	}
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	parse, err := time.Parse(ansic, "8:41PM")
	p(err)
	p(parse)
}

//2025-05-16 10:29:01.6273423 +0800 CST m=+0.000543001
//2025-05-16T10:29:01+08:00
//2012-11-01 22:08:41 +0000 +0000
//10:29AM
//Fri May 16 10:29:01 2025
//2025-05-16T10:29:01.627342+08:00
//0000-01-01 20:41:00 +0000 UTC
//2025-05-16T10:29:01-00:00
//parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
//0001-01-01 00:00:00 +0000 UTC
