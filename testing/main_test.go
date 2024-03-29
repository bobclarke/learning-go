package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"invalid-data", 200.0, 0.0, 0.0, true},
}

func TestFromTable(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but didn't get one")
			}
		} else {
			if err != nil {
				t.Error("Got an error but didn't expect one")
			}
		}
		if got != tt.expected {
			t.Errorf("Got %f but expected %f", got, tt.expected)
		}
	}
}
