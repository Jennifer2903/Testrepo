package main

import (
	"fmt"
	"os"
)

func main() {
	//read output from file
	//read output from file
	file, err := os.ReadFile("./file.txt")
	if err != nil {
		// error handling goes here, print it out and exit
	}

	fmt.Printf("Print message: %s, \n", file)

	//read input from file
	var name string

	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("Hello, %s! You are %d year old \n ", name, age)

}
