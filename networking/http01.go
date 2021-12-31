package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// Get request
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal("ERROR %s", err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("-----")

	// Post request

	j := &Job{ // Instantiate a new instance of Job
		User:   "Saitama",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer // Create an IO buffer (this is needed for http Post)

	enc := json.NewEncoder(&buf) // Create a new JSON encoder that can write to this IO buffer

	err = enc.Encode(j) // Encode our Job struct as json into our IO buffer

	if err != nil {
		log.Fatalf("ERROR unable to encode our Job stuct data into JSON: %+v\n", err)
	}

	resp, err = http.Post("http://httpbin.org/post", "application/json", &buf) // Post it
	if err != nil {
		log.Fatalf("ERROR unable to call httpbin.org, %+v", err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}
