package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	 Needles and pins
	 needles and pins 
	 Sew me a sail
	 to catch the wind
	`
	lc_text := strings.ToLower(text)
	f := strings.Fields(lc_text)

	fmt.Printf("f is %q (%T) \n", f, f)

	for k, v := range f {
		fmt.Printf("key is %d value is %s\n", k, v)
	}
}
