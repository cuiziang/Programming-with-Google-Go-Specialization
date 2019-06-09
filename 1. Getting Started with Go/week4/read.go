package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	file, _ := os.Open("1. Getting Started with Go/week4/foo.txt")
	defer file.Close()
	names := make([]Name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name Name
		line := strings.Split(scanner.Text(), " ")
		name.fname = line[0]
		name.lname = line[1]
		names = append(names, name)
	}
	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}
}
