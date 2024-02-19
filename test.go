package main

import (
	"fmt"
)

func main() {
	var testmap map[string]string

	if testmap == nil {
		fmt.Println("Nil map")
	} else {
		fmt.Println("Not nil map")
	}
}
