package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Test out returning a struct as JSON

	m := &MathResponse{ // Create a dummy instance of MathResponse
		Error:  "No error",
		Result: 1.34,
	}

	enc := json.NewEncoder(w) // Create an encoder over the response writer

	err := enc.Encode(m) // Encode the MathResponse instance
	if err != nil {
		fmt.Fprintf(w, "ERROR: %s\n", err)
	}

	fmt.Printf("OUT: %v\n", r.Body)
}

func mathHandler(writer http.ResponseWriter, request *http.Request) {
	// Accept a JSON encoded MathRequest object, decode it, carry out the calculation
	// with, encode the result (or a error) in JSON and return vie the http response writer

	// Decode request
	math_request := &MathRequest{}       // Create an empty instance of a MathRequest struct
	dec := json.NewDecoder(request.Body) // Create a new json decoder over the request body
	err := dec.Decode(math_request)      // Decode the JSON request and put the result in the math_request struct instance
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	// Do the calculation and store results in a new math_response struct
	math_response := &MathResponse{}
	op := math_request.Op
	left := math_request.Left
	right := math_request.Right

	switch op {
	case "+":
		math_response.Result = left + right
	case "-":
		math_response.Result = left - right
	case "*":
		math_response.Result = left * right
	case "/":
		if right == 0.0 {
			math_response.Error = "ERROR: Division by zero"
		} else {
			math_response.Result = left / right
		}
	default:
		math_response.Error = fmt.Sprintf("ERROR: unknown operator %s", op)
	}

	// Encode the math_response struct into JSON and send it over http
	enc := json.NewEncoder(writer) // Create a JSON encoder over the http writer
	writer.Header().Set("Content-Type", "application/json")
	if math_response.Error != "" {
		writer.WriteHeader(http.StatusBadRequest)
	}
	enc.Encode(math_response) // Encode the struct and sednd
}

func main() {
	// Register handlers
	http.HandleFunc("/math", mathHandler)
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
