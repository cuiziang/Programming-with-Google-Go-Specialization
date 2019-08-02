package week2

import "fmt"

func mainShadow() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go mainShadow()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
