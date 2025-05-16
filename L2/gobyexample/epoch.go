package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))

}

//2025-05-16 09:54:57.5610495 +0800 CST m=+0.000501601
//1747360497
//1747360497561
//1747360497561049500
//2025-05-16 09:54:57 +0800 CST
//2025-05-16 09:54:57.5610495 +0800 CST
