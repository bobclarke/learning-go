package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func getContentType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("ERROR: %+v", err)
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	fmt.Printf("URL: %s CTYPE: %s\n", url, ct)

}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://golang.cafe/",
		"https://yourbasic.org/golang/unused-local-variables/",
		"https://www.tutorialkart.com/golang-tutorial/golang-division-operator/",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://google.co.uk",
		"https://golang.org",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://bbc.co.uk",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://api.github.com",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
		"https://www.includehelp.com/golang/math-max-function-with-examples.aspx",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			getContentType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
