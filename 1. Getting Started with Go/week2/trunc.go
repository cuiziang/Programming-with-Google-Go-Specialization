package main

import (
	"fmt"
)

func main() {
	var f float32
	fmt.Print("Enter a floating point number:\n")
	_, _ = fmt.Scan(&f)
	fmt.Println(int(f))
}
