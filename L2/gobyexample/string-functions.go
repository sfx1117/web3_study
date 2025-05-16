package main

import (
	"fmt"
	"strings"
)

/*
	订阅区块
*/

func main() {
	fmt.Println("contains", strings.Contains("string", "s"))
	fmt.Println("Count", strings.Count("strings", "s"))
	fmt.Println("HasPrefix", strings.HasPrefix("string", "s"))
	fmt.Println("HasSuffix", strings.HasSuffix("string", "ing"))
	fmt.Println("Index", strings.Index("string", "s"))
	fmt.Println("Join", strings.Join([]string{"te", "da"}, ","))
	fmt.Println("Repeat", strings.Repeat("adf", 3))
	fmt.Println("Replace", strings.Replace("string", "tr", "bb", 3))
	fmt.Println("Replace", strings.Replace("string", "st", "bb", -1))
	fmt.Println("Split", strings.Split("string,add,dbb", ","))
	fmt.Println("ToLower", strings.ToLower("STRTRW"))
	fmt.Println("ToUpper", strings.ToUpper("string"))
}

//contains true
//Count 2
//HasPrefix true
//HasSuffix true
//Index 0
//Join te,da
//Repeat adfadfadf
//Replace sbbing
//Replace bbring
//Split [string add dbb]
//ToLower strtrw
//ToUpper STRING
