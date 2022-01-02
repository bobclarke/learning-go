package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var db map[string]interface{} // define an in memory database

type Entry struct { // Define struct for a database entry - this will be used for queryig and updating
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	key := r.URL.Path[4:] // parse the key from the url
	value, ok := db[key]  // use the key to find the value from the database
	if !ok {
		http.Error(w, fmt.Sprintf("key %q not found", key), http.StatusNotFound)
	}

	entry := &Entry{
		Key:   key,
		Value: value,
	}

	enc := json.NewEncoder(w)                 // Creat a JSON encoder over the http.ResponseWriter
	if err := enc.Encode(entry); err != nil { // encode our query object
		fmt.Fprintf(w, "ERROR: ", err)
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	entry := &Entry{}                         // Create an empty Entry object
	enc := json.NewDecoder(r.Body)            // Create a JSON decoder around the request body
	if err := enc.Decode(entry); err != nil { // Load our json request into our entry object
		fmt.Fprintf(w, "ERROR: %s", err)
	}
	db[entry.Key] = entry.Value // Update the database
}

func main() {
	db = make(map[string]interface{})     // init database
	http.HandleFunc("/db/", queryHandler) // register handler(s)
	http.HandleFunc("/db", updateHandler)

	err := http.ListenAndServe("0.0.0.0:8080", nil) // start an http server
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
