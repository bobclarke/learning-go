package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	url := "https://api.github.com/users/" + login

	resp, err := http.Get(url) //Get the JSON data
	if err != nil {
		log.Fatalf("ERROR: unable to call %s - %s", url, err)
	}

	defer resp.Body.Close()

	//var buf bytes.Buffer // Create a buffer

	//io.Copy(&buf, resp.Body) // copy the http response to this buffer

	u := &User{} // Create an empty User instance - note, u is a pointer to this User instance

	enc := json.NewDecoder(resp.Body) // Create a JSON decoder to decode our HTTP response

	err = enc.Decode(u) // Decode the JSON response into the User struct
	if err != nil {
		return u, fmt.Errorf("Error")
	} else {
		return u, nil
	}
}

func main() {
	userDetails, err := userInfo("bobclarke")
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	fmt.Printf("User: %s - public repos: %d\n", userDetails.Name, userDetails.PublicRepos)
}
