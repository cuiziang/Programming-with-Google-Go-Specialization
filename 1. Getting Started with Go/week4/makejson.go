package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, add string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a name:\n")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Print("Enter a add:\n")
	if scanner.Scan() {
		add = scanner.Text()
	}
	addMap := map[string]string{
		"name":    name,
		"address": add,
	}

	addJosn, _ := json.Marshal(addMap)
	fmt.Print(string(addJosn))
}
