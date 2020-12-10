package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println(url)
	resp, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(len(body))
}

func main() {
	go responseSize("https://www.inhatc.ac.kr")
	go responseSize("https://www.nate.com")
	go responseSize("https://www.google.com")
	time.Sleep(time.Second * 5)
}
