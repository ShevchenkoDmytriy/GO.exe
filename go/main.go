package main

import "fmt"

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
	// var a, b, c float64
	// fmt.Println("Enter a")
	// fmt.Scan(&a)
	// fmt.Println("Enter b")
	// fmt.Scan(&b)
	// fmt.Println("Enter c")
	// fmt.Scan(&c)
	// D := (b * b) - 4*a*c
	// if D > 0 {
	// 	x1 := (-b + math.Sqrt(D)) / 2 * a
	// 	x2 := (-b - math.Sqrt(D)) / 2 * a
	// 	fmt.Println(fmt.Sprint(x1) + " " + fmt.Sprint(x2))
	// } else if D < 0 {
	// 	fmt.Println("Error")
	// } else {
	// 	x1 := (-b) / (2 * a)
	// 	fmt.Println(x1)
	// }
	// fmt.Scan(&a)
	// for i := 1; i <= 100; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// for _, row := range matrix {
	// 	for _, col := range row {
	// 		fmt.Printf("%d", col)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// world := []string{"H", "E", "L", "L", "O", "W", "O", "R", "L", "D"}
	// mas := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// for i := 0; i < len(world); i++ {
	// 	for j := 0; j < len(mas); j++ {
	// 		if world[i] == mas[j] {
	// 			fmt.Printf(mas[j])
	// 		}
	// 	}
	// }
	num := 6
	switch {
	case num%2 == 0:
		fmt.Println("/2")
	case num%2 != 0:
		fmt.Println("!/2")

	}
}
