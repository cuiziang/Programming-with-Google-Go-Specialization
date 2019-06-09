package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Enter a string:\n")
	_, _ = fmt.Scanln(&s)
	strings.ToLower(s)
	if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
		fmt.Println("Found!" +
			"")
	} else {
		fmt.Println("Not Found!")
	}
}
