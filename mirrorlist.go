package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetMirrors() []byte {
	response, err := http.Get("https://www.archlinux.org/mirrors/status/json/")
	if err != nil {
		log.Fatal(err)
	}
	read, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
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

func SetOut(f string) (w io.Writer) {
	if strings.TrimSpace(f) == "" {
		w = os.Stdout
	} else {
		file, err := os.Create(f)
		if err != nil {
			log.Fatal(err)
		}
		w = file
	}
	return
}

func main() {
	var f string
	flag.StringVar(&f, "o", "", "Output to file.")
	flag.Parse()
	out := SetOut(f)
	mirrors := GetMirrors()
	var result Result
	err := json.Unmarshal(mirrors, &result)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range result.Urls {
		fmt.Fprintf(out, "## Score: %v, %v\n", v.Score, v.Country)
		fmt.Fprintf(out, "#Server = %v$repo/os/$arch\n\n", v.Url)
	}
}
