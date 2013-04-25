package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get_mirrors() []byte {
	response, err := http.Get("https://www.archlinux.org/mirrors/status/json/")
	if err != nil {
		log.Fatal(err)
	}
	read, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return read
}

type Mirror struct {
	Protocol, Url, Country, Country_code string
	Completion_pct, Score                float64
}

type Result struct {
	Urls []Mirror
}

func main() {
	mirrors := get_mirrors()
	var result Result
	err := json.Unmarshal(mirrors, &result)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(len(result.Urls))
	}
}
