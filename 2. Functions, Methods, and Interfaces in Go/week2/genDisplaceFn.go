package main

import "fmt"

func main() {

	var acceleration, initialVelocity, initialDisplacement, time float32
	fmt.Print("Enter acceleration: ")
	_, _ = fmt.Scanln(&acceleration)
	fmt.Print("Enter initial velocity: ")
	_, _ = fmt.Scanln(&initialVelocity)
	fmt.Print("Enter initial displacement: ")
	_, _ = fmt.Scanln(&initialDisplacement)

	fn := GenDisplaceFn(10, 2, 1)

	fmt.Print("Enter time: ")
	_, _ = fmt.Scanln(&time)
	fmt.Printf("%f sec:\n", time)
	fmt.Println(fn(time))

}

func GenDisplaceFn(acceleration float32, initialVelocity float32, initialDisplacement float32) func(float32) float32 {
	fn := func(time float32) float32 {
		return (acceleration*time*time)/2 + initialVelocity*time + initialDisplacement
	}
	return fn
}
