package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//read output from file
	prg2, _ := ioutil.ReadFile("prg2.txt")
	fmt.Printf("Print message: %s/n", prg2)

	//read input from file
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("Hello, %s! You are %d year old \n ", name, age)

}
