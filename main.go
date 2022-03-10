package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://example.com/x/y%2Fz")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.ForceQuery)
	fmt.Println(u.Fragment)
	fmt.Println(u.Opaque)
	fmt.Println(u.RawFragment)
	fmt.Println("Path:", u.Path)
	fmt.Println("RawPath:", u.RawPath)
	fmt.Println("EscapedPath:", u.EscapedPath())
}
