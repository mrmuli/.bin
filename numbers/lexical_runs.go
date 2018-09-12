package main

import (
	"fmt"
)

var dogName string = "buddy"

func main(){
	call()
	require()
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

	fmt.Println("output: ", output)
}