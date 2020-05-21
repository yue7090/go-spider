package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("hell go scrapy!!")
	res, err := http.Get("http://www.mangocity.com")
	if err != nil {
		fmt.Println("http error ", err)
		return 
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read error", err)
		return 
	}
	fmt.Println(string(body))
}