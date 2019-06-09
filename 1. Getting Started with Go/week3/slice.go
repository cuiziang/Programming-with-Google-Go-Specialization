package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	i := -1

	for true {
		var s string
		fmt.Print("Enter a num, exit if \"X\":\n")
		_, _ = fmt.Scanln(&s)
		ne, _ := strconv.Atoi(s)
		if s == "X" {
			break
		}

		if i < 2 {
			i++

			sli[i] = ne
		} else {
			sli = append(sli, ne)
		}
	}

	sort.Ints(sli)

	for i, v := range sli {
		fmt.Printf("A[%d]: %d\n", i, v)
	}
}
