package main

import (
	"fmt"

	"github.com/bobclarke/learning01/helpers"
)

func main() {
	fmt.Println("Hello")

	var st helpers.SomeType
	st.Name = "Test"
	st.Description = "This is a test"

	fmt.Println(st.Name)

}
