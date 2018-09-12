package main

import (
	"fmt" // str literal 
)


func main() {
	fmt.Println("go" + "lang")

	fmt.Println("321325 × 424521 = ", 321325*424521)

	fmt.Println((true && false) || (false && true) || !(false && false))
	fmt.Println(!(false && false) || (false && true))
	fmt.Println(true || false)
	fmt.Println("choc chip"[1])
}