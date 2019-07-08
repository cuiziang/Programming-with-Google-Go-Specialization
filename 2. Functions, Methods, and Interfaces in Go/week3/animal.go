package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var input1 string
	var input2 string
	var animal Animal

	for true {
		fmt.Printf("> ")
		_, _ = fmt.Scan(&input1)
		_, _ = fmt.Scan(&input2)

		if input1 == "cow" {
			animal = Animal{"grass", "walk", "moo"}
		} else if input1 == "bird" {
			animal = Animal{"worms", "fly", "peep"}
		} else if input1 == "snake" {
			animal = Animal{"mice", "slither", "hsss"}
		} else {
			fmt.Printf("type error\n")
			continue
		}

		if input2 == "eat" {
			animal.Eat()
		} else if input2 == "move" {
			animal.Move()
		} else if input2 == "speak" {
			animal.Speak()
		} else {
			fmt.Printf("type error\n")
			continue
		}
	}
}
