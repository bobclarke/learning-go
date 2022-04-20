package main

import (
	"encoding/json"
	"fmt"
)

type Superhero struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Alias     string `json:"alias"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clarke",
		"last_name": "Kent",
		"alias": "Superman",
		"has_dog": 	true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"alias": "Batman",
		"has_dog": false
	}
]`

	readJsonIntoStruct(myJson)

}

func readJsonIntoStruct(s string) {
	var superHeros []Superhero

	err := json.Unmarshal([]byte(s), &superHeros)
	if err != nil {
		fmt.Println("Error unmarhalling JSON", err)
	}

	for _, val := range superHeros {
		fmt.Println(val.Alias)
	}
}

func jsonFromStruct() {

	var sh1 Superhero
	sh1.FirstName = "Peter"
	sh1.LastName = "Parker"
	sh1.Alias = "Spiderman"
	sh1.HasDog = true

	json.Marshal()

}
