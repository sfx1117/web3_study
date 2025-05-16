package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}

//2025-05-16 09:50:25.3105641 +0800 CST m=+0.000537401
//2009-11-17 20:34:58.651387237 +0000 UTC
//2009
//November
//17
//20
//34
//58
//651387237
//UTC
//Tuesday
//true
//false
//false
//135797h15m26.659176863s
//135797.2574053269
//8.147835444319614e+06
//4.888701266591769e+08
//488870126659176863
//2025-05-16 01:50:25.3105641 +0000 UTC
//1994-05-22 15:19:31.992210374 +0000 UTC
