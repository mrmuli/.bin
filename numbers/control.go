package main

import (
	"fmt"
)



func main() {
	// buddy()
	// loopNumerals()
	getSteps()
}


func buddy() {
	fmt.Println("Heeeey buddy")
}

func loopNumbers() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}
}

func loopNumerals() {
	for i := 0; i <=10; i++ {
		fmt.Println(i)
	}
}

func getSteps() {
	// fmt.Println("Enter number: ")
	// var input int
	// fmt.Scanf("%f", &input)

	step := 1
	// for input >= 1 {
	// 	if input % 2 == 0 {
	// 		fmt.Println("even")
	// 		step += 1
	// 		input = input / 2
	// 	} else {
	// 		fmt.Println("odd")
	// 		step += 1
	// 		input = (input * 3) + 1
	// 	}
	// }
	for i := 12; i >= 1; i++ {
		if i % 2 == 0 {
			step += 1
			fmt.Println(step)
			i = i / 2
			fmt.Println(i)
		} else {
			step += 1
			fmt.Println(step)
			i = (i * 3) + 1
			fmt.Println(i)
		}
	}
}