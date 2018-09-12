package main

import (
	"fmt"
)

var dogName string = "buddy"

func main(){
	Call()
	//Require()
	//Celcius()
	//FeetToMeters()
	//ActualAge()
}

func Call() {
	fmt.Println("Heeey", dogName)
}

func Require(){
	fmt.Println("give me a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println("confirming... ", input)
	output := input * 2
	fmt.Println("doubled input: ", output)
}

func Celcius() {
	fmt.Println("Enter temp in Fahrenheit: ")
	var input float32
	fmt.Scanf("%f", &input)
	c := (input - 32) * 5/9

	fmt.Println("temp in celcius: ", c)
}

func FeetToMeters() {
	fmt.Println("Enter feet for conv: ")
	var input float32
	fmt.Scanf("%f", &input)
	meters := input * 0.3048

	fmt.Println("meters: ", meters)
}

func ActualAge() {
	fmt.Println("Enter number of seconds: ")
	var input float64
	fmt.Scanf("%f", &input)
	age := input / 31557600

	fmt.Println("actual age in years is: ", age)
}