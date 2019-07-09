package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

func main() {

	cow := Cow{"grass", "walk", "moo"}
	bird := Bird{"worms", "fly", "peep"}
	Snake := Snake{"mice", "slither", "hsss"}

	var input1, input2, input3 string
	animalMap := make(map[string]Animal)

	for true {
		fmt.Printf("> ")
		_, _ = fmt.Scan(&input1)
		_, _ = fmt.Scan(&input2)
		_, _ = fmt.Scan(&input3)

		if input1 == "newanimal" {
			if input3 == "cow" {
				animalMap[input2] = cow
			} else if input3 == "bird" {
				animalMap[input2] = bird
			} else if input3 == "snake" {
				animalMap[input2] = Snake
			} else {
				fmt.Printf("type error\n")
				continue
			}

		} else if input1 == "query" {
			animal, ok := animalMap[input2]
			if !ok {
				fmt.Println("Incorrect input. Try again")
			} else {
				if input3 == "eat" {
					animal.Eat()
				} else if input3 == "move" {
					animal.Move()
				} else if input3 == "speak" {
					animal.Speak()
				} else {
					fmt.Println("Incorrect input. Try again")
				}
			}

		}
	}
}
