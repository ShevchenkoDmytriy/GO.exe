package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello World!");
	// var a int = 10
	// fmt.Println(a)
	// var num float64 = 22.22
	// fmt.Println(num)
	// b := 22
	// fmt.Println(b)
	// name := "asfaf"
	// fmt.Println(name)
	// fmt.Println(len(name))
	// name1 := "qfqfgo run main.go"
	// fmt.Println("Whats your name?")
	// fmt.Scan(&name1)
	// fmt.Println("Hello " + name1)
	// var age int
	// fmt.Println("How old are you?")
	// fmt.Scan(&age)
	// fmt.Println("You " + fmt.Sprint(age))
	// num := 0
	// if num > 0 {
	// 	fmt.Print(">")
	// } else if num < 0 {
	// 	fmt.Println("<")
	// } else {
	// 	fmt.Println("=")
	// }
	var a, b, c float64
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	fmt.Println("Enter c")
	fmt.Scan(&c)
	D := (b * b) - 4*a*c
	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / 2 * a
		x2 := (-b - math.Sqrt(D)) / 2 * a
		fmt.Println(fmt.Sprint(x1) + " " + fmt.Sprint(x2))
	} else if D < 0 {
		fmt.Println("Error")
	} else {
		x1 := (-b) / (2 * a)
		fmt.Println(x1)
	}
}
