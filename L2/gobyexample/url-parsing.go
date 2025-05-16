package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	query, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(query)
	fmt.Println(query["k"])
	fmt.Println(query["k"][0])
}

//postgres://user:pass@host.com:5432/path?k=v#f
//postgres
//user:pass
//user
//pass true
//host.com:5432
//host.com
//5432
///path
//f
//k=v
//map[k:[v]]
//[v]
//v
