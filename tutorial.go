package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to my game!\n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)

	if age < 10 {
		fmt.Println("You can play!")
	} else {
		fmt.Println("You can't play :(")
		return
	}

	fmt.Printf("Cats or dogs?")
	var answer string
	fmt.Scan(&answer)

	fmt.Println(answer)
}
