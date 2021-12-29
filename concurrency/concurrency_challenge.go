package main

import (
	"fmt"
	"net/http"
)

func getContentType(url string, result chan string) {
	resp, err := http.Get(url)
	if err != nil {
		result <- fmt.Sprintf("ERROR: %+v", err)
	}
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	result <- fmt.Sprintf("%s ---> %s", url, ct)
}

func main() {

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

	ch := make(chan string)
	for _, url := range urls {
		go getContentType(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
}
