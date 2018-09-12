package main

import (
	"fmt"
)

var dogName string = "buddy"

func main(){
	call()
	//require()
	celcius()
}

func call(){
	fmt.Println("Heeey", dogName)
}

func require(){
	fmt.Println("give me a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println("confirming... ", input)
	output := input * 2
	fmt.Println("doubled input: ", output)
}

func celcius(){
	fmt.Println("Enter temp in Fahrenheit: ")
	var input float32
	fmt.Scanf("%f", &input)
	c := (input - 32) * 5/9

	fmt.Println("temp in celcius: ", c)
}