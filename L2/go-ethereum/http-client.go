package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	resp, err := http.Get("https://gobyexample.com/http-client")
	check(err)
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

//Response status: 200 OK
//<!DOCTYPE html>
//<html>
//<head>
//<meta charset="utf-8">
//<title>Go by Example: HTTP Client</title>
