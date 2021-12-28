package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://google.co.uk"
	ct, err := contentType(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	fmt.Printf("Content type for %s is %s", url, ct)
}

func contentType(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "foo", err
	} else {
		ct := resp.Header.Get("Content-Type")
		return ct, nil
	}
}
