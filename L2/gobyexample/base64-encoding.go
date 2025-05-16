package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	s := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(s)

	decodeString, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(decodeString))

	toString := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(toString)
	bytes, _ := base64.URLEncoding.DecodeString(toString)
	fmt.Println(string(bytes))
}

//YWJjMTIzIT8kKiYoKSctPUB+
//abc123!?$*&()'-=@~
//YWJjMTIzIT8kKiYoKSctPUB-
//abc123!?$*&()'-=@~
