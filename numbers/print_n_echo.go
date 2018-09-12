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

	var x string 
	x = "Hello"
	fmt.Println(x)
	x = "why now"
	x += " ?"
	fmt.Println(x)

	var y string = "hello"
	var z string = "hello"
	var k string = "acdce"
	fmt.Println(y == z) 
	fmt.Println(y == k)

	j := "peanut"
	l := "butter"
	pbj := j + " " + l + " n jelly"
	fmt.Println(j + " " + l)
	fmt.Println(pbj)
}