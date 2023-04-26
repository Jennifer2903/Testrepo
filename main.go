package main

import (
	"fmt"
	"os"
)

func main() {
	//read output from file
	file, _ := os.ReadFile("./file.txt")
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
