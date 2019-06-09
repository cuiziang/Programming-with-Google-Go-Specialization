package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, add string

	fmt.Print("Enter a name:\n")
	_, _ = fmt.Scanln(&name)

	fmt.Print("Enter a add:\n")
	_, _ = fmt.Scanln(&add)

	addMap := map[string]string{
		name: add,
	}

	addJosn, _ := json.Marshal(addMap)
	fmt.Print(string(addJosn))
}
